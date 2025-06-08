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

func (server *Server) DeleteScore(ctx context.Context, req *pb.DeleteScoreRequest) (*emptypb.Empty, error) {
	violations := validateDeleteScoreRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	err := server.store.DeleteScore(ctx, req.GetScoreId())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "score with ID %d not found", req.GetScoreId())
		}
		return nil, status.Errorf(codes.Internal, "failed to delete score: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func validateDeleteScoreRequest(req *pb.DeleteScoreRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetScoreId()); err != nil {
		violations = append(violations, fieldViolation("score_id", err))
	}

	return violations
}
