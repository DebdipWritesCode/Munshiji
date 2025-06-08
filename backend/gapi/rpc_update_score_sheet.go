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

func (server *Server) UpdateScoreSheet(ctx context.Context, req *pb.UpdateScoreSheetRequest) (*pb.UpdateScoreSheetResponse, error) {
	violations := validateUpdateScoreSheetRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	arg := db.UpdateSheetParams{
		ID: req.GetScoreSheetId(),
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  req.GetName() != "",
		},
		CommitteeName: sql.NullString{
			String: req.GetCommitteeName(),
			Valid:  req.GetCommitteeName() != "",
		},
		Chair: sql.NullString{
			String: req.GetChair(),
			Valid:  req.GetChair() != "",
		},
		ViceChair: sql.NullString{
			String: req.GetViceChair(),
			Valid:  req.GetViceChair() != "",
		},
		Rapporteur: sql.NullString{
			String: req.GetRapporteur(),
			Valid:  req.GetRapporteur() != "",
		},
	}

	scoreSheet, err := server.store.UpdateSheet(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "score sheet with ID %d not found", req.GetScoreSheetId())
		}
		return nil, status.Errorf(codes.Internal, "failed to update score sheet: %v", err)
	}

	rsp := &pb.UpdateScoreSheetResponse{
		ScoreSheet: convertScoreSheetToProto(scoreSheet),
	}

	return rsp, nil
}

func validateUpdateScoreSheetRequest(req *pb.UpdateScoreSheetRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetScoreSheetId()); err != nil {
		violations = append(violations, fieldViolation("score_sheet_id", err))
	}

	if req.GetName() != "" {
		if err := val.ValidateName(req.GetName()); err != nil {
			violations = append(violations, fieldViolation("name", err))
		}
	}

	if req.GetCommitteeName() != "" {
		if err := val.ValidateName(req.GetCommitteeName()); err != nil {
			violations = append(violations, fieldViolation("committee_name", err))
		}
	}

	if req.GetChair() != "" {
		if err := val.ValidateName(req.GetChair()); err != nil {
			violations = append(violations, fieldViolation("chair", err))
		}
	}

	if req.GetViceChair() != "" {
		if err := val.ValidateName(req.GetViceChair()); err != nil {
			violations = append(violations, fieldViolation("vice_chair", err))
		}
	}

	if req.GetRapporteur() != "" {
		if err := val.ValidateName(req.GetRapporteur()); err != nil {
			violations = append(violations, fieldViolation("rapporteur", err))
		}
	}

	return violations
}
