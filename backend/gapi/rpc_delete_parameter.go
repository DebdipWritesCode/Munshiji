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

func (server *Server) DeleteParameter(ctx context.Context, req *pb.DeleteParameterRequest) (*emptypb.Empty, error) {
	violations := validateDeleteParameterRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	err := server.store.DeleteParameter(ctx, req.GetParameterId())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "parameter with ID %d not found", req.GetParameterId())
		}
		return nil, status.Errorf(codes.Internal, "failed to delete parameter: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func validateDeleteParameterRequest(req *pb.DeleteParameterRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetParameterId()); err != nil {
		violations = append(violations, fieldViolation("parameter_id", err))
	}

	return violations
}
