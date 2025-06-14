-- name: CreateAISession :one
INSERT INTO ai_sessions (
  user_id, expires_at, prompt, status
) VALUES (
  $1, $2, $3, 'completed'
) 
RETURNING *;

-- name: GetAISessionsByUserID :many
SELECT * FROM ai_sessions
WHERE user_id = $1
AND expires_at > NOW()
ORDER BY created_at DESC;

-- name: DeleteExpiredAISessions :exec
DELETE FROM ai_sessions
WHERE user_id = $1
AND expires_at <= NOW();