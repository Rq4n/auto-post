-- name: CreateUser :one
INSERT INTO users (google_id, email)
VALUES($1, $2)
RETURNING *;

-- name: GetUserByGoogleID :one
SELECT * FROM USERS
WHERE google_id = $1;

