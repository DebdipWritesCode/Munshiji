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

func (server *Server) CreateScore(ctx context.Context, req *pb.CreateScoreRequest) (*pb.CreateScoreResponse, error) {
	violations := validateCreateScoreRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	_, err := server.store.GetParameterByID(ctx, req.GetParameterId())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "parameter with ID %d not found", req.GetParameterId())
		}
		return nil, status.Errorf(codes.Internal, "failed to get parameter: %v", err)
	}

	_, err = server.store.GetDelegateByID(ctx, req.GetDelegateId())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "delegate with ID %d not found", req.GetDelegateId())
		}
		return nil, status.Errorf(codes.Internal, "failed to get delegate: %v", err)
	}

	arg := db.CreateScoreParams{
		DelegateID:  req.GetDelegateId(),
		ParameterID: req.GetParameterId(),
		Value:       req.GetValue(),
		Note: sql.NullString{
			String: req.GetNote(),
			Valid:  req.GetNote() != "",
		},
	}

	score, err := server.store.CreateScore(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create score: %v", err)
	}

	rsp := &pb.CreateScoreResponse{
		Score: ConvertScoreToProto(score),
	}

	return rsp, nil
}

func validateCreateScoreRequest(req *pb.CreateScoreRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetDelegateId()); err != nil {
		violations = append(violations, fieldViolation("delegate_id", err))
	}
	if err := val.ValidateID(req.GetParameterId()); err != nil {
		violations = append(violations, fieldViolation("parameter_id", err))
	}
	if req.GetNote() != "" {
		if err := val.ValidateNote(req.GetNote()); err != nil {
			violations = append(violations, fieldViolation("note", err))
		}
	}

	return violations
}
