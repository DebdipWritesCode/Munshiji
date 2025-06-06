-- name: CreateSession :one
INSERT INTO sessions (
  user_id, jwt_token, refresh_token, expires_at, refresh_expires_at
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetSessionsByUserID :one
SELECT * FROM sessions
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: GetSessionByID :one
SELECT * FROM sessions
WHERE id = $1;

-- name: GetSessionsByRefreshToken :one
SELECT * FROM sessions
WHERE refresh_token = $1;

-- name: DeleteSessionByID :exec
DELETE FROM sessions
WHERE id = $1;

-- name: DeleteSessionsByUserID :exec
DELETE FROM sessions
WHERE user_id = $1;

-- name: UpdateSessionExpiration :one
UPDATE sessions
SET expires_at = $2,
  jwt_token = $3
WHERE id = $1
RETURNING *;