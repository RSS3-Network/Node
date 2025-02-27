-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS dataset_bluesky_profiles
(
    "did"        text PRIMARY KEY,
    "handle"     text,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX idx_dataset_bluesky_profiles_cursor ON dataset_bluesky_profiles (updated_at DESC, created_at DESC);

CREATE INDEX idx_dataset_bluesky_profiles_handle ON dataset_bluesky_profiles (handle);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS dataset_bluesky_profiles;
-- +goose StatementEnd
