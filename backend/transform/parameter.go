package transform

import (
	"database/sql"

	db "github.com/DebdipWritesCode/Munshiji/backend/db/sqlc"
)

type ParameterJSON struct {
	ID                 int32    `json:"id"`
	ScoreSheetID       int32    `json:"score_sheet_id"`
	Name               string   `json:"name"`
	RuleType           string   `json:"rule_type"`
	IsSpecialParameter bool     `json:"is_special_parameter"`
	SpecialScoresRule  *string  `json:"special_scores_rule"`
	SpecialLengthRule  *string  `json:"special_length_rule"`
	ScoreWeight        *float64 `json:"score_weight"`
	LengthWeight       *float64 `json:"length_weight"`
}

func ConvertParametersToDB(jsonParams []ParameterJSON) []db.Parameter {
	params := make([]db.Parameter, 0, len(jsonParams))
	for _, p := range jsonParams {
		params = append(params, db.Parameter{
			ID:                 p.ID,
			ScoreSheetID:       p.ScoreSheetID,
			Name:               p.Name,
			RuleType:           p.RuleType,
			IsSpecialParameter: p.IsSpecialParameter,
			SpecialScoresRule: sql.NullString{
				String: derefString(p.SpecialScoresRule),
				Valid:  p.SpecialScoresRule != nil,
			},
			SpecialLengthRule: sql.NullString{
				String: derefString(p.SpecialLengthRule),
				Valid:  p.SpecialLengthRule != nil,
			},
			ScoreWeight: sql.NullFloat64{
				Float64: derefFloat64(p.ScoreWeight),
				Valid:   p.ScoreWeight != nil,
			},
			LengthWeight: sql.NullFloat64{
				Float64: derefFloat64(p.LengthWeight),
				Valid:   p.LengthWeight != nil,
			},
		})
	}
	return params
}
