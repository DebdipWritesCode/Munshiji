package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/DebdipWritesCode/Munshiji/backend/constants"
	"github.com/DebdipWritesCode/Munshiji/backend/util"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskVerifySendMail = "task:send_verify_mail"

type PayloadSendVerifyMail struct {
	UserID string `json:"user_id"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyMail(
	ctx context.Context,
	payload *PayloadSendVerifyMail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	task := asynq.NewTask(TaskVerifySendMail, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("Task enqueued successfully")
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(
	ctx context.Context,
	task *asynq.Task,
) error {
	var payload PayloadSendVerifyMail

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		log.Error().Err(err).Str("type", task.Type()).Bytes("payload", task.Payload()).Msg("Failed to unmarshal payload")
		return asynq.SkipRetry
	}
	userID, err := strconv.Atoi(payload.UserID)
	if err != nil {
		log.Error().Err(err).Str("type", task.Type()).Bytes("payload", task.Payload()).Msg("Invalid user ID in payload")
		return asynq.SkipRetry
	}

	user, err := processor.store.GetUserByID(ctx, int32(userID))
	if err != nil {
		if err == sql.ErrNoRows {
			return asynq.SkipRetry
		}
		return fmt.Errorf("failed to get user by ID: %w", err)
	}

	config, err := util.LoadConfig("..")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load config")
		return asynq.SkipRetry
	}

	subject := constants.EmailSubject
	verifyUrl := fmt.Sprintf("%s?token=%s", config.FrontendVerifyEmailURL, payload.UserID)
	content := constants.CreateEmailBody(verifyUrl)
	to := []string{user.Email}

	if err := processor.mailer.SendEmail(subject, content, to, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to send verification email: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).Str("user_email", user.Email).Msg("Verification email sent successfully")
	return nil
}
