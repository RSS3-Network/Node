-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS dataset_mastodon_update_handles (
    "handle" text PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX idx_mastodon_update_handles_time_at ON dataset_mastodon_update_handles(updated_at DESC,created_at DESC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS dataset_mastodon_update_handles;
-- +goose StatementEnd
