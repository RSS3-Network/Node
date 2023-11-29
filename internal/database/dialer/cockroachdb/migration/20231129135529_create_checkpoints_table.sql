-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "checkpoints"
(
    "id"         text        NOT NULL,
    "chain"      text        NOT NULL,
    "worker"     text        NOT NULL,
    "state"      jsonb       NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),

    CONSTRAINT "pk_checkpoints" PRIMARY KEY ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "checkpoints";
-- +goose StatementEnd
