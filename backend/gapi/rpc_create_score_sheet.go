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

func (server *Server) CreateScoreSheet(ctx context.Context, req *pb.CreateScoreSheetRequest) (*pb.CreateScoreSheetResponse, error) {
	violations := validateCreateScoreSheetRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	_, err := server.store.GetUserByID(ctx, req.GetCreatedBy())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user with ID %d not found", req.GetCreatedBy())
		}
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	arg := db.CreateSheetParams{
		Name:          req.GetName(),
		CommitteeName: req.GetCommitteeName(),
		Chair:         req.GetChair(),
		ViceChair: sql.NullString{
			String: req.GetViceChair(),
			Valid:  req.GetViceChair() != "",
		},
		Rapporteur: sql.NullString{
			String: req.GetRapporteur(),
			Valid:  req.GetRapporteur() != "",
		},
		CreatedBy: req.GetCreatedBy(),
	}

	scoreSheet, err := server.store.CreateSheet(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create score sheet: %v", err)
	}

	rsp := &pb.CreateScoreSheetResponse{
		ScoreSheet: convertScoreSheetToProto(scoreSheet),
	}

	return rsp, nil
}

func validateCreateScoreSheetRequest(req *pb.CreateScoreSheetRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}
	if err := val.ValidateName(req.GetCommitteeName()); err != nil {
		violations = append(violations, fieldViolation("committee_name", err))
	}
	if err := val.ValidateName(req.GetChair()); err != nil {
		violations = append(violations, fieldViolation("chair", err))
	}
	if err := val.ValidateID(req.GetCreatedBy()); err != nil {
		violations = append(violations, fieldViolation("created_by", err))
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
