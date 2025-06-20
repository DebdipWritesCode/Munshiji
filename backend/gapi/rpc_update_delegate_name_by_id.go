package gapi

import (
	"context"
	"database/sql"

	db "github.com/DebdipWritesCode/Munshiji/backend/db/sqlc"
	"github.com/DebdipWritesCode/Munshiji/backend/pb"
	"github.com/DebdipWritesCode/Munshiji/backend/val"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateDelegateNameByID(ctx context.Context, req *pb.UpdateDelegateNameByIDRequest) (*pb.UpdateDelegateNameByIDResponse, error) {
	violations := validateUpdateDelegateNameByIDRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	dl, err := server.store.GetDelegateByID(ctx, req.GetDelegateId())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "delegate with ID %d not found", req.GetDelegateId())
		}
		return nil, status.Errorf(codes.Internal, "failed to retrieve delegate: %v", err)
	}

	existing, err := server.store.GetDelegateByScoreSheetIDAndName(ctx, db.GetDelegateByScoreSheetIDAndNameParams{
		ScoreSheetID: dl.ScoreSheetID,
		Name:         req.GetName(),
	})

	if err == nil && existing.ID != dl.ID {
		return nil, status.Errorf(codes.AlreadyExists, "delegate with name %s already exists in score sheet %d", req.GetName(), dl.ScoreSheetID)
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, status.Errorf(codes.Internal, "failed to check for existing delegate: %v", err)
	}

	arg := db.UpdateDelegateNameParams{
		ID:   req.GetDelegateId(),
		Name: req.GetName(),
	}

	delegate, err := server.store.UpdateDelegateName(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "delegate with ID %d not found", req.GetDelegateId())
		}
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			return nil, status.Errorf(codes.AlreadyExists, "delegate with name %s already exists", req.GetName())
		}
		return nil, status.Errorf(codes.Internal, "failed to update delegate: %v", err)
	}

	rsp := &pb.UpdateDelegateNameByIDResponse{
		Delegate: ConvertDelegateToProto(delegate),
	}

	return rsp, nil
}

func validateUpdateDelegateNameByIDRequest(req *pb.UpdateDelegateNameByIDRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetDelegateId()); err != nil {
		violations = append(violations, fieldViolation("delegate_id", err))
	}
	if err := val.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}

	return violations
}
