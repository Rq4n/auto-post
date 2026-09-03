-- name: ConnectNewProvider :one
INSERT INTO social_connections (
  user_id,
  provider,
  provider_user_id,
  access_token,
  refresh_token,
  expires_at
)
VALUES($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetProviderByUserID :many
SELECT
   id,
   user_id,
   provider_user_id,
   created_at
FROM social_connections
WHERE user_id = $1
ORDER BY created_at ASC;
