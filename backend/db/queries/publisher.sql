-- name: FetchPendingJobs :many
SELECT * FROM publishers
WHERE status = 'pending'
  AND scheduled_at <= NOW()
ORDER BY scheduled_at ASC
FOR UPDATE SKIP LOCKED
LIMIT 50;

-- name: UpdateJobAsProcessing :exec
UPDATE publishers
SET status = 'processing',
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateJobAsCompleted :exec
UPDATE publishers
SET status = 'completed',
    published_at = NOW(),
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateJobAsFailed :exec
UPDATE publishers
SET status = 'failed',
    updated_at = NOW()
WHERE id = $1;
