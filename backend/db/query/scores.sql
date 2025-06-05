-- name: CreateScore :one
INSERT INTO scores (
  parameter_id, delegate_id, value, note 
)
VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetScoreByID :one
SELECT * FROM scores
WHERE id = $1;

-- name: GetScoresByParameterID :many
SELECT * FROM scores
WHERE parameter_id = $1
ORDER BY id;

-- name: GetScoresByDelegateID :many
SELECT * FROM scores
WHERE delegate_id = $1
ORDER BY id;

-- name: UpdateScore :one
UPDATE scores
SET
  value = COALESCE(sqlc.narg(value), value),
  note = COALESCE(sqlc.narg(note), note)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteScore :exec
DELETE FROM scores
WHERE id = $1;