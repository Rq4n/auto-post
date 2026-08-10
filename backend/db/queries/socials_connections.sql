-- name: ConnectNewProvider :one
INSERT INTO social_connections (provider)
VALUES($1)
RETURNING *;



