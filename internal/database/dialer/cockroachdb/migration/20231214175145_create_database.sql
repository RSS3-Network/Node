-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "indexes"
(
    "id"         text        NOT NULL,
    "owner"      text        NOT NULL,
    "network"    text        NOT NULL,
    "index"      int         NOT NULL,
    "platform"   text,
    "tag"        text        NOT NULL,
    "type"       text        NOT NULL,
    "status"     bool        NOT NULL,
    "direction"  int         NOT NULL,
    "timestamp"  timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),

    CONSTRAINT "pk_indexes" PRIMARY KEY ("id", "network", "owner")
);

CREATE INDEX IF NOT EXISTS "idx_indexes_platform" ON "indexes" ("platform", "timestamp" DESC, "index" DESC);
CREATE INDEX IF NOT EXISTS "idx_indexes_filter" ON "indexes" ("tag", "timestamp" DESC, "index" DESC, "type");
CREATE INDEX IF NOT EXISTS "idx_indexes_network" ON "indexes" ("network", "timestamp" DESC, "index" DESC);
CREATE INDEX IF NOT EXISTS "idx_indexes_owner" ON "indexes" ("owner", "timestamp" DESC, "index" DESC, "direction");
CREATE INDEX IF NOT EXISTS "idx_indexes_timestamp" ON "indexes" ("timestamp" DESC, "index" DESC);

CREATE TABLE IF NOT EXISTS "feeds"
(
    "id"            text        NOT NULL,
    "network"       text        NOT NULL,
    "platform"      text,
    "index"         int         NOT NULL,
    "from"          text        NOT NULL,
    "to"            text        NOT NULL,
    "tag"           text        NOT NULL,
    "type"          text        NOT NULL,
    "status"        bool        NOT NULL,
    "actions"       json        NOT NULL,
    "total_actions" bigint      NOT NULL,
    "fee"           json,
    "timestamp"     timestamptz NOT NULL,
    "created_at"    timestamptz NOT NULL DEFAULT now(),
    "updated_at"    timestamptz NOT NULL DEFAULT now(),

    CONSTRAINT "pk_feeds" PRIMARY KEY ("id")
    );

CREATE INDEX IF NOT EXISTS "idx_feeds_platform" ON "feeds" ("platform");
CREATE INDEX IF NOT EXISTS "idx_feeds_timestamp" ON "feeds" ("timestamp" DESC);
CREATE INDEX IF NOT EXISTS "idx_feeds_tag_type" ON "feeds" ("tag", "type");
CREATE INDEX IF NOT EXISTS "idx_feeds_total_actions" ON "feeds" ("total_actions");

CREATE TABLE IF NOT EXISTS "checkpoints"
(
    "id"         text        NOT NULL,
    "network"    text        NOT NULL,
    "worker"     text        NOT NULL,
    "state"      jsonb       NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),

    CONSTRAINT "pk_checkpoints" PRIMARY KEY ("id")
    );

CREATE TABLE IF NOT EXISTS "dataset_mirror_posts"
(
    "id"                   text        NOT NULL,
    "origin_content_digest" text     NOT NULL,

    CONSTRAINT "pk_dataset_mirror_posts" PRIMARY KEY ("id")
    );

CREATE INDEX IF NOT EXISTS idx_origin_content_digest ON dataset_mirror_posts (origin_content_digest);


CREATE TABLE IF NOT EXISTS "dataset_farcaster_profiles"
(
    "fid"             bigint NOT NULL,
    "username"        text,
    "custody_address" text,
    "eth_addresses"   text[],

    CONSTRAINT "pk_dataset_farcaster_profiles" PRIMARY KEY ("fid")
    );


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "indexes";
DROP TABLE IF EXISTS "feeds";
DROP TABLE IF EXISTS "checkpoints";
DROP TABLE IF EXISTS "dataset_mirror_posts";
DROP TABLE IF EXISTS "dataset_farcaster_profiles";
-- +goose StatementEnd
