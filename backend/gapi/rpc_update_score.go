package gapi

import (
	"context"
	"database/sql"

	db "github.com/DebdipWritesCode/MUN_Scoresheet/backend/db/sqlc"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/pb"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateScore(ctx context.Context, req *pb.UpdateScoreRequest) (*pb.UpdateScoreResponse, error) {
	violations := validateUpdateScoreRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	value := sql.NullFloat64{}
	if req.Value != nil {
		value = sql.NullFloat64{
			Float64: *req.Value,
			Valid:   true,
		}
	}

	note := sql.NullString{}
	if req.Note != nil {
		note = sql.NullString{
			String: *req.Note,
			Valid:  true,
		}
	}

	arg := db.UpdateScoreParams{
		ID:    req.GetScoreId(),
		Value: value,
		Note:  note,
	}

	score, err := server.store.UpdateScore(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "score with ID %d not found", req.GetScoreId())
		}
		return nil, status.Errorf(codes.Internal, "failed to update score: %v", err)
	}

	rsp := &pb.UpdateScoreResponse{
		Score: ConvertScoreToProto(score),
	}

	return rsp, nil
}

func validateUpdateScoreRequest(req *pb.UpdateScoreRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetScoreId()); err != nil {
		violations = append(violations, fieldViolation("score_id", err))
	}
	if req.GetNote() != "" {
		if err := val.ValidateNote(req.GetNote()); err != nil {
			violations = append(violations, fieldViolation("note", err))
		}
	}

	return violations
}
