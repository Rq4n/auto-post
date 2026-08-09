-- +goose Up
CREATE TABLE IF NOT EXISTS social_connections (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_ID UUID REFERENCES user(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  
  CONSTRAINT unique_user_provider UNIQUE (user_id, provider),
  CONSTRAINT valid_provider CHECK (provider IN ('linkedin', 'twitter'))

);

-- +goose Down
DROP TABLE IF NOT EXISTS social_connections;
  
