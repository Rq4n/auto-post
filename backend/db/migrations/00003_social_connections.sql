-- +goose Up
CREATE TABLE social_connections (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    provider TEXT NOT NULL,
    provider_user_id TEXT NOT NULL,

    access_token TEXT NOT NULL,
    refresh_token TEXT,

    expires_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (user_id, provider, provider_user_id)
);

-- +goose Down
DROP TABLE IF EXISTS social_connections;
