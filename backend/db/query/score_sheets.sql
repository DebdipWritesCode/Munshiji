-- name: CreateSheet :one
INSERT INTO score_sheets (
  name, committee_name, chair, vice_chair, rapporteur, created_by
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetSheetByID :one
SELECT * FROM score_sheets
WHERE id = $1;

-- name: GetSheetsByUserID :many
SELECT * FROM score_sheets
WHERE created_by = $1;

-- name: GetSheetWithDetailsByID :one
SELECT
  ss.*,
  d.delegates,
  p.parameters,
  s.scores
FROM score_sheets ss

LEFT JOIN LATERAL (
  SELECT json_agg(d) AS delegates
  FROM delegates d
  WHERE d.score_sheet_id = ss.id
) d ON true

LEFT JOIN LATERAL (
  SELECT json_agg(p) AS parameters
  FROM parameters p
  WHERE p.score_sheet_id = ss.id
) p ON true

LEFT JOIN LATERAL (
  SELECT json_agg(s) AS scores
  FROM scores s
  JOIN delegates d2 ON s.delegate_id = d2.id
  JOIN parameters p2 ON s.parameter_id = p2.id
  WHERE d2.score_sheet_id = ss.id
) s ON true

WHERE ss.id = $1;

-- name: UpdateSheet :one
UPDATE score_sheets
SET name = COALESCE(sqlc.narg(name), name),
    committee_name = COALESCE(sqlc.narg(committee_name), committee_name),
    chair = COALESCE(sqlc.narg(chair), chair),
    vice_chair = COALESCE(sqlc.narg(vice_chair), vice_chair),
    rapporteur = COALESCE(sqlc.narg(rapporteur), rapporteur),
    updated_at = now()
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteSheet :exec
DELETE FROM score_sheets
WHERE id = $1;

-- name: DeleteSheetsByUserID :exec
DELETE FROM score_sheets
WHERE created_by = $1;