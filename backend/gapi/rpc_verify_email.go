package gapi

import (
	"context"

	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/pb"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	violations := validateVerifyEmailRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	err := server.store.VerifyUserEmail(ctx, req.GetToken())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to verify email: %s", err.Error())
	}

	return &pb.VerifyEmailResponse{
		Message: "Email verified successfully, you can now log in.",
	}, nil
}

func validateVerifyEmailRequest(req *pb.VerifyEmailRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetToken()); err != nil {
		violations = append(violations, fieldViolation("token", err))
	}

	return violations
}
