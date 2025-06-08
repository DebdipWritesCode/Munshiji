package gapi

import (
	"context"
	"database/sql"

	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/pb"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) DeleteDelegate(ctx context.Context, req *pb.DeleteDelegateRequest) (*emptypb.Empty, error) {
	violations := validateDeleteDelegateRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	err := server.store.DeleteDelegate(ctx, req.GetDelegateId())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "delegate with ID %d not found", req.GetDelegateId())
		}
		return nil, status.Errorf(codes.Internal, "failed to delete delegate: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func validateDeleteDelegateRequest(req *pb.DeleteDelegateRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetDelegateId()); err != nil {
		violations = append(violations, fieldViolation("delegate_id", err))
	}

	return violations
}
