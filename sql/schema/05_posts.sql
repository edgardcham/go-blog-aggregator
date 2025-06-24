-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    description TEXT NOT NULL,
    published_at TIMESTAMP WITH TIME ZONE NOT NULL,
    feed_id UUID NOT NULL REFERENCES feeds(id)
);

-- +goose Down
DROP TABLE posts;