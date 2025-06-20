package gapi

import (
	"context"
	"encoding/json"

	db "github.com/DebdipWritesCode/Munshiji/backend/db/sqlc"
	"github.com/DebdipWritesCode/Munshiji/backend/pb"
	"github.com/DebdipWritesCode/Munshiji/backend/transform"
	"github.com/DebdipWritesCode/Munshiji/backend/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetScoreSheetDetails(ctx context.Context, req *pb.GetScoreSheetDetailsRequest) (*pb.GetScoreSheetDetailsResponse, error) {
	violations := validateGetScoreSheetDetailsRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	details, err := server.store.GetSheetWithDetailsByID(ctx, req.GetScoreSheetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get score sheets by user ID: %v", err)
	}

	scoreSheet := convertScoreSheetToProto(db.ScoreSheet{
		ID:            details.ID,
		Name:          details.Name,
		CommitteeName: details.CommitteeName,
		CreatedAt:     details.CreatedAt,
		UpdatedAt:     details.UpdatedAt,
		CreatedBy:     details.CreatedBy,
		Chair:         details.Chair,
		ViceChair:     details.ViceChair,
		Rapporteur:    details.Rapporteur,
	})

	var (
		delegates  []db.Delegate
		parameters []db.Parameter
		scores     []db.Score
	)

	if details.Delegates != nil {
		if err := json.Unmarshal(details.Delegates, &delegates); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to unmarshal delegates: %v", err)
		}
	}

	var paramJSON []transform.ParameterJSON
	if details.Parameters != nil {
		if err := json.Unmarshal(details.Parameters, &paramJSON); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to unmarshal parameters: %v", err)
		}
		parameters = transform.ConvertParametersToDB(paramJSON)
	}

	var scoreJSON []transform.ScoreJSON
	if details.Scores != nil {
		if err := json.Unmarshal(details.Scores, &scoreJSON); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to unmarshal scores: %v", err)
		}
		scores = transform.ConvertScoresToDB(scoreJSON)
	}

	pbDelegates := make([]*pb.Delegate, 0, len(delegates))
	for _, d := range delegates {
		pbDelegates = append(pbDelegates, ConvertDelegateToProto(d))
	}

	pbParameters := make([]*pb.Parameter, 0, len(parameters))
	for _, p := range parameters {
		pbParameters = append(pbParameters, ConvertParameterToProto(p))
	}

	pbScores := make([]*pb.Score, 0, len(scores))
	for _, s := range scores {
		pbScores = append(pbScores, ConvertScoreToProto(s))
	}

	rsp := &pb.GetScoreSheetDetailsResponse{
		ScoreSheet: scoreSheet,
		Delegates:  pbDelegates,
		Parameters: pbParameters,
		Scores:     pbScores,
	}

	return rsp, nil
}

func validateGetScoreSheetDetailsRequest(req *pb.GetScoreSheetDetailsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetScoreSheetId()); err != nil {
		violations = append(violations, fieldViolation("score_sheet_id", err))
	}

	return violations
}
