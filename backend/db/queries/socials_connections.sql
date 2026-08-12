-- name: ConnectNewProvider :one
INSERT INTO social_connections (user_id, provider)
VALUES($1, $2)
RETURNING id, user_id, provider, created_at, updated_at;



