-- +goose Up
-- +goose StatementBegin
ALTER TABLE "activities"
    ADD "calldata" json;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "activities"
DROP COLUMN "calldata";
-- +goose StatementEnd
