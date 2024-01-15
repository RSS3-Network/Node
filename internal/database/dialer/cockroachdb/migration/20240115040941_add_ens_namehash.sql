-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "dataset_ens_namehashes"
(
    "hash_hex" text NOT NULL,
    "name"     text NOT NULL,

    CONSTRAINT "pk_dataset_ens_namehashs" PRIMARY KEY ("hash_hex")
);

CREATE INDEX IF NOT EXISTS idx_ensnamehash_name ON dataset_ens_namehashes (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_ensnamehash_name;
DROP TABLE IF EXISTS "dataset_ens_namehashes";
-- +goose StatementEnd
