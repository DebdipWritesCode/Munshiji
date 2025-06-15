package gapi

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	db "github.com/DebdipWritesCode/MUN_Scoresheet/backend/db/sqlc"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/pb"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/tools"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/transform"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetFeedbackByLLM(ctx context.Context, req *pb.GetFeedbackByLLMRequest) (*pb.GetFeedbackByLLMResponse, error) {
	violations := validateGetFeedbackByLLM(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	aiSessions, err := server.store.GetAISessionsByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get AI sessions by user ID: %v", err)
	}

	if len(aiSessions) != 0 {
		expiryTime := aiSessions[0].ExpiresAt
		remainingTime := expiryTime.Sub(time.Now().UTC()).Seconds()

		return nil, status.Errorf(codes.FailedPrecondition, "You can request feedback again after %f seconds", remainingTime)
	}

	delegates := req.GetDelegates()

	promptBatches := transform.BatchAndGeneratePrompts(delegates, 5)

	type Result struct {
		Delegates []*pb.DelegateInput
		Response  string
		Err       error
	}

	resultCh := make(chan Result, len(promptBatches))
	var wg sync.WaitGroup

	for i, batch := range promptBatches {
		wg.Add(1)

		go func(i int, b transform.BatchedPrompt) {
			defer wg.Done()
			resp, err := tools.CallLLM(server.config.OpenAIAPIKey, b.Prompt)
			resultCh <- Result{
				Delegates: b.Delegates,
				Response:  resp,
				Err:       err,
			}
		}(i, batch)
	}

	wg.Wait()
	close(resultCh)

	var feedbacks []*pb.DelegateFeedback

	for res := range resultCh {
		if res.Err != nil {
			return nil, status.Errorf(codes.Internal, "LLM batch failed: %v", res.Err)
		}

		parsed := transform.ParseLLMFeedback(res.Response)
		feedbacks = append(feedbacks, parsed...)
	}

	err = server.store.DeleteExpiredAISessions(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete expired AI sessions: %v", err)
	}

	expiration := time.Now().UTC().Add(5 * time.Minute)

	var fullPrompt strings.Builder
	for _, b := range promptBatches {
		fullPrompt.WriteString(b.Prompt)
		fullPrompt.WriteString("\n\n")
	}

	_, err = server.store.CreateAISession(ctx, db.CreateAISessionParams{
		UserID:    req.GetUserId(),
		Prompt:    fullPrompt.String(),
		ExpiresAt: expiration,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create AI session: %v", err)
	}

	rsp := &pb.GetFeedbackByLLMResponse{
		Feedbacks: feedbacks,
	}

	return rsp, nil
}

func validateGetFeedbackByLLM(req *pb.GetFeedbackByLLMRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetUserId()); err != nil {
		violations = append(violations, fieldViolation("user_id", err))
	}

	delegates := req.GetDelegates()

	for i, delegate := range delegates {
		if err := val.ValidateDelegateInput(delegate); err != nil {
			violations = append(violations, fieldViolation(fmt.Sprintf("delegates[%d]", i), err))
		}
	}

	return violations
}
