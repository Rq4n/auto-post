-- name: CreatePublisherJob :many
INSERT INTO publisher (post_id, user_id, social_connection_id, status)
SELECT $1, $2, unnest(@social_connection_ids::uuid[]), 'pending'
RETURNING *;

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
