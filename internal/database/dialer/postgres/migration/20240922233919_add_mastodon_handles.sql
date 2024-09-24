-- +goose Up
CREATE TABLE IF NOT EXISTS dataset_mastodon_handles (
                                          handle VARCHAR(255) PRIMARY KEY,
                                          last_updated TIMESTAMP NOT NULL
);

CREATE INDEX idx_mastodon_handles_last_updated ON dataset_mastodon_handles(last_updated);

-- +goose Down
-- +goose StatementBegin
DROP TABLE  IF EXISTS dataset_mastodon_handles;
-- +goose StatementEnd