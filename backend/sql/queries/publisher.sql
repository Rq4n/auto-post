
-- name: CreatePublisherJob :many
INSERT INTO publisher (post_id, user_id, social_connection_id, status)
SELECT $1, $2, sc_id, 'pending'
FROM unnest($3::uuid[]) AS sc_id
RETURNING *;

-- name: GetPublishersByPostID :many
SELECT * FROM publisher
WHERE post_id = $1
ORDER BY created_at ASC;

-- name: UpdatePublisherAsFailed :exec
UPDATE publisher
SET status = 'failed',
    updated_at = NOW()
WHERE id = $1;

-- name: UpdatePublisherAsCompleted :exec
UPDATE publisher
SET status = 'completed',
    publisher_at = NOW(),
    updated_at = NOW()
WHERE id = $1;
