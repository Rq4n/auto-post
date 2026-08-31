-- +goose Up
CREATE TABLE IF NOT EXISTS publisher (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    post_id UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    social_connection_id UUID NOT NULL REFERENCES social_connections(id) ON DELETE CASCADE,

    status TEXT NOT NULL DEFAULT 'pending'
        CHECK (status IN (
            'pending',
            -- 'processing',
            'completed',
            'failed'
            -- 'cancelled'
        )),

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (post_id, social_connection_id)
);

-- +goose Down
DROP TABLE IF EXISTS publisher;
