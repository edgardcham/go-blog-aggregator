-- +goose Up
CREATE TABLE feeds
(
    id         UUID PRIMARY KEY,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL,
    name       VARCHAR(255) NOT NULL,
    url TEXT NOT NULL,
    user_id UUID NOT NULL,
    UNIQUE(url),
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE
);


-- +goose Down
DROP TABLE feeds;