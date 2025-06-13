package transform

import (
	"database/sql"

	db "github.com/DebdipWritesCode/MUN_Scoresheet/backend/db/sqlc"
)

type ScoreJSON struct {
	ID          int32   `json:"id"`
	DelegateID  int32   `json:"delegate_id"`
	ParameterID int32   `json:"parameter_id"`
	Value       float64 `json:"value"`
	Note        *string `json:"note"`
}

func ConvertScoresToDB(jsonScores []ScoreJSON) []db.Score {
	scores := make([]db.Score, 0, len(jsonScores))
	for _, s := range jsonScores {
		scores = append(scores, db.Score{
			ID:          s.ID,
			DelegateID:  s.DelegateID,
			ParameterID: s.ParameterID,
			Value:       s.Value,
			Note:        sql.NullString{String: derefString(s.Note), Valid: s.Note != nil},
		})
	}
	return scores
}
