package gapi

import (
	"context"
	"database/sql"

	"github.com/DebdipWritesCode/Munshiji/backend/pb"
	"github.com/DebdipWritesCode/Munshiji/backend/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) DeleteScoreSheet(ctx context.Context, req *pb.DeleteScoreSheetRequest) (*emptypb.Empty, error) {
	violations := validateDeleteScoreSheetRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	err := server.store.DeleteSheet(ctx, req.GetScoreSheetId())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "score sheet with ID %d not found", req.GetScoreSheetId())
		}
		return nil, status.Errorf(codes.Internal, "failed to delete score sheet: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func validateDeleteScoreSheetRequest(req *pb.DeleteScoreSheetRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetScoreSheetId()); err != nil {
		violations = append(violations, fieldViolation("score_sheet_id", err))
	}

	return violations
}
