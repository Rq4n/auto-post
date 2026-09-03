-- name: CreateNewPosts :one
INSERT INTO posts (user_id, title, body)
VALUES ($1, $2, $3)
RETURNING id, user_id, title, body, created_at, updated_at;;


-- name: GetPublishersByPostID :many
SELECT * FROM publisher
WHERE post_id = $1
ORDER BY created_at ASC;

