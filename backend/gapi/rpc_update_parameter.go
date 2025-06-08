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

func (server *Server) UpdateParameter(ctx context.Context, req *pb.UpdateParameterRequest) (*pb.UpdateParameterResponse, error) {
	violations := validateUpdateParameterRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	arg := db.UpdateParameterParams{
		ID: req.GetParameterId(),
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
		RuleType: sql.NullString{
			String: req.GetRuleType(),
			Valid:  req.RuleType != nil,
		},
		IsSpecialParameter: sql.NullBool{
			Bool:  req.GetIsSpecialParameter(),
			Valid: req.IsSpecialParameter != nil,
		},
		SpecialScoresRule: sql.NullString{
			String: req.GetSpecialScoresRule(),
			Valid:  req.SpecialScoresRule != nil,
		},
		SpecialLengthRule: sql.NullString{
			String: req.GetSpecialLengthRule(),
			Valid:  req.SpecialLengthRule != nil,
		},
		ScoreWeight: sql.NullFloat64{
			Float64: req.GetScoreWeight(),
			Valid:   req.ScoreWeight != nil,
		},
		LengthWeight: sql.NullFloat64{
			Float64: req.GetLengthWeight(),
			Valid:   req.LengthWeight != nil,
		},
	}

	parameter, err := server.store.UpdateParameter(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "parameter with ID %d not found", req.GetParameterId())
		}
		return nil, status.Errorf(codes.Internal, "failed to update parameter: %v", err)
	}

	rsp := &pb.UpdateParameterResponse{
		Parameter: ConvertParameterToProto(parameter),
	}

	return rsp, nil
}

func validateUpdateParameterRequest(req *pb.UpdateParameterRequest) []*errdetails.BadRequest_FieldViolation {
	var violations []*errdetails.BadRequest_FieldViolation

	violations = append(violations, validateRequiredUpdateFields(req)...)
	violations = append(violations, validateSpecialUpdateRules(req)...)
	violations = append(violations, validateUpdateWeights(req)...)

	return violations
}

func validateRequiredUpdateFields(req *pb.UpdateParameterRequest) []*errdetails.BadRequest_FieldViolation {
	var violations []*errdetails.BadRequest_FieldViolation

	if err := val.ValidateID(req.GetParameterId()); err != nil {
		violations = append(violations, fieldViolation("parameter_id", err))
	}

	if req.GetName() != "" {
		if err := val.ValidateName(req.GetName()); err != nil {
			violations = append(violations, fieldViolation("name", err))
		}
	}

	if req.GetRuleType() != "" {
		if err := val.ValidateRule(req.GetRuleType(), []string{"absolute", "average", "special"}); err != nil {
			violations = append(violations, fieldViolation("rule_type", err))
		}
	}

	return violations
}

func validateSpecialUpdateRules(req *pb.UpdateParameterRequest) []*errdetails.BadRequest_FieldViolation {
	var violations []*errdetails.BadRequest_FieldViolation

	if err := val.ValidateSpecialConditionRule(req.GetRuleType(), req.IsSpecialParameter); err != nil {
		violations = append(violations, fieldViolation("is_special_parameter", err))
	}

	if req.IsSpecialParameter != nil {
		validRules := []string{"absolute", "average"}
		if *req.IsSpecialParameter {
			validRules = []string{"special"}
		}
		if err := val.ValidateRule(req.GetRuleType(), validRules); err != nil {
			violations = append(violations, fieldViolation("rule_type", err))
		}
	}

	if req.SpecialScoresRule != nil {
		if err := val.ValidateRule(*req.SpecialScoresRule, []string{"absolute", "average"}); err != nil {
			violations = append(violations, fieldViolation("special_scores_rule", err))
		}
	}

	if req.SpecialLengthRule != nil {
		if err := val.ValidateRule(*req.SpecialLengthRule, []string{"absolute", "average"}); err != nil {
			violations = append(violations, fieldViolation("special_length_rule", err))
		}
	}

	return violations
}

func validateUpdateWeights(req *pb.UpdateParameterRequest) []*errdetails.BadRequest_FieldViolation {
	var violations []*errdetails.BadRequest_FieldViolation

	if err := val.ValidateWeight("score_weight", req.ScoreWeight); err != nil {
		violations = append(violations, fieldViolation("score_weight", err))
	}

	if err := val.ValidateWeight("length_weight", req.LengthWeight); err != nil {
		violations = append(violations, fieldViolation("length_weight", err))
	}

	return violations
}
