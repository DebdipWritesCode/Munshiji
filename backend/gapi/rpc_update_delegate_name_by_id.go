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

func (server *Server) UpdateDelegateNameByID(ctx context.Context, req *pb.UpdateDelegateNameByIDRequest) (*pb.UpdateDelegateNameByIDResponse, error) {
	violations := validateUpdateDelegateNameByIDRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
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
