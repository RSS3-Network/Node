-- +goose Up

CREATE TABLE IF NOT EXISTS dataset_mastodon_update_handles (
                                           handle VARCHAR(255) PRIMARY KEY,
                                            last_updated BIGINT NOT NULL
);

CREATE INDEX idx_mastodon_recent_update_handles_last_updated ON dataset_mastodon_update_handles(last_updated DESC);

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS dataset_mastodon_recent_update_handles;

-- +goose StatementEnd
