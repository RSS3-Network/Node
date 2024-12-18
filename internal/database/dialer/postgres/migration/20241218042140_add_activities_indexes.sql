-- +goose Up
-- +goose StatementBegin
CREATE INDEX CONCURRENTLY IF NOT EXISTS activities_id_platform_index ON public.activities (platform);
CREATE INDEX CONCURRENTLY IF NOT EXISTS activities_id_tag_and_type_index ON public.activities (tag, type);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX CONCURRENTLY IF EXISTS public.activities_id_platform_index;
DROP INDEX CONCURRENTLY IF EXISTS public.activities_id_tag_and_type_index;
-- +goose StatementEnd
