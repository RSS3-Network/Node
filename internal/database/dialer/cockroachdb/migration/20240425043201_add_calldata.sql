-- +goose Up
-- +goose StatementBegin
ALTER TABLE "feeds"
    ADD "calldata" json;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "feeds"
DROP COLUMN "calldata";
-- +goose StatementEnd
