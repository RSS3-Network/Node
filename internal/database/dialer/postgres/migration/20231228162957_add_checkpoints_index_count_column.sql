-- +goose Up
-- +goose StatementBegin
ALTER TABLE "checkpoints"
    ADD "index_count" bigint NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "checkpoints"
DROP COLUMN "index_count";
-- +goose StatementEnd
