package gapi

import (
	"context"
	"strconv"
	"time"

	db "github.com/DebdipWritesCode/MUN_Scoresheet/backend/db/sqlc"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/pb"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/util"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/val"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/worker"
	"github.com/hibiken/asynq"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	violations := validateCreateUserRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %v", err)
	}

	arg := db.CreateUserTxParams{
		CreateUserParams: db.CreateUserParams{
			Name:         req.GetName(),
			Email:        req.GetEmail(),
			PasswordHash: hashedPassword,
		},
		AfterCreate: func(user db.User) error {
			taskPayload := &worker.PayloadSendVerifyMail{
				UserID: strconv.Itoa(int(user.ID)),
			}

			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.QueueCritical),
			}

			return server.taskDistributor.DistributeTaskSendVerifyMail(ctx, taskPayload, opts...)
		},
	}

	txResult, err := server.store.CreateUserTx(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			return nil, status.Errorf(codes.AlreadyExists, "user with email %s already exists: %s", req.GetEmail(), err.Error())
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	rsp := &pb.CreateUserResponse{
		User: converUserToProto(txResult.User),
	}

	return rsp, nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}
	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}
	return violations
}
