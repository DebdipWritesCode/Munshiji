package gapi

import (
	"context"

	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/pb"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetScoreSheetByUserId(ctx context.Context, req *pb.GetScoreSheetByUserIdRequest) (*pb.GetScoreSheetByUserIdResponse, error) {
	violations := validateGetScoreSheetsByUserID(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	sheets, err := server.store.GetSheetsByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get score sheets by user ID", err)
	}

	pbSheets := make([]*pb.ScoreSheet, len(sheets))
	for i, sheet := range sheets {
		pbSheets[i] = convertScoreSheetToProto(sheet)
	}

	rsp := &pb.GetScoreSheetByUserIdResponse{
		ScoreSheets: pbSheets,
	}

	return rsp, nil
}

func validateGetScoreSheetsByUserID(req *pb.GetScoreSheetByUserIdRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetUserId()); err != nil {
		violations = append(violations, fieldViolation("user_id", err))
	}

	return violations
}
