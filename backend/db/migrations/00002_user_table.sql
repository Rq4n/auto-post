-- +goose Up
SELECT 'up SQL query';

-- +goose Down
DROP TABLE IF NOT EXISTS users;
