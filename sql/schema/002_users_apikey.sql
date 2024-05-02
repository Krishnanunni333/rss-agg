-- +goose Up
ALTER TABLE users ADD COLUMN IF NOT EXISTS api_key TEXT UNIQUE NOT NULL DEFAULT encode(sha256(random()::text::bytea), 'hex');

-- +goose Down
ALTER TABLE users DROP COLUMN IF EXISTS api_key;