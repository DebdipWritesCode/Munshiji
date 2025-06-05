-- name: CreateDelegate :one
INSERT INTO delegates (
  score_sheet_id, name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetDelegateByID :one
SELECT * FROM delegates
WHERE id = $1;

-- name: GetDelegatesByScoreSheetID :many
SELECT * FROM delegates
WHERE score_sheet_id = $1;

-- name: UpdateDelegateName :one
UPDATE delegates
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteDelegate :exec
DELETE FROM delegates
WHERE id = $1;