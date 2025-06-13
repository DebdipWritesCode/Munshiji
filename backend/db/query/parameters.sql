-- name: CreateParameter :one
INSERT INTO parameters (
  score_sheet_id, name, rule_type, is_special_parameter,
  special_scores_rule, special_length_rule,
  score_weight, length_weight
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetParameterByID :one
SELECT * FROM parameters
WHERE id = $1;

-- name: GetParameterByScoreSheetIDAndName :one
SELECT * FROM parameters
WHERE score_sheet_id = $1 AND name = $2;

-- name: GetParametersByScoreSheetID :many
SELECT * FROM parameters
WHERE score_sheet_id = $1
ORDER BY id;

-- name: UpdateParameter :one
UPDATE parameters
SET
  name = COALESCE(sqlc.narg(name), name),
  rule_type = COALESCE(sqlc.narg(rule_type), rule_type),
  is_special_parameter = COALESCE(sqlc.narg(is_special_parameter), is_special_parameter),
  special_scores_rule = COALESCE(sqlc.narg(special_scores_rule), special_scores_rule),
  special_length_rule = COALESCE(sqlc.narg(special_length_rule), special_length_rule),
  score_weight = COALESCE(sqlc.narg(score_weight), score_weight),
  length_weight = COALESCE(sqlc.narg(length_weight), length_weight)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteParameter :exec
DELETE FROM parameters
WHERE id = $1;