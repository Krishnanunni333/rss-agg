-- +goose Up
CREATE TABLE
    feeds (
        id UUID PRIMARY KEY,
        created_at TIMESTAMPTZ NOT NULL DEFAULT now (),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT now (),
        name TEXT NOT NULL,
        url TEXT UNIQUE NOT NULL,
        user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE
    );

-- +goose Down
DROP TABLE feeds;