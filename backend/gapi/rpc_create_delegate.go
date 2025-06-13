package gapi

import (
	"context"
	"database/sql"

	db "github.com/DebdipWritesCode/MUN_Scoresheet/backend/db/sqlc"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/pb"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/val"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateDelegate(ctx context.Context, req *pb.CreateDelegateRequest) (*pb.CreateDelegateResponse, error) {
	violations := validateCreateDelegateRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	_, err := server.store.GetSheetByID(ctx, req.GetScoreSheetId())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "score sheet with ID %d not found", req.GetScoreSheetId())
		}
		return nil, status.Errorf(codes.Internal, "failed to get score sheet: %v", err)
	}

	arg := db.CreateDelegateParams{
		Name:         req.GetName(),
		ScoreSheetID: req.GetScoreSheetId(),
	}

	delegate, err := server.store.CreateDelegate(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			return nil, status.Errorf(codes.AlreadyExists, "delegate with name %s already exists for score sheet ID %d", req.GetName(), req.GetScoreSheetId())
		}
		return nil, status.Errorf(codes.Internal, "failed to create delegate: %v", err)
	}

	rsp := &pb.CreateDelegateResponse{
		Delegate: ConvertDelegateToProto(delegate),
	}

	return rsp, nil
}

func validateCreateDelegateRequest(req *pb.CreateDelegateRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}
	if err := val.ValidateID(req.GetScoreSheetId()); err != nil {
		violations = append(violations, fieldViolation("score_sheet_id", err))
	}

	return violations
}
