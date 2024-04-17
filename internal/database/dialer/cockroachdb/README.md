# Migration with Atlas

`atlas.hcl` is the configuration file for Atlas.
Think it as a terraform but for databases.
It contains the configurations and environments for:
1. database connection
2. where to store migration files
3. where to look for schema files (gorm models)
4. other configurations

## Generate an Atlas schema.hcl
Inspect without generation
```
    atlas schema inspect -u "postgres://root:@localhost:26257?sslmode=disable"
```
Or generate a schema.hcl file (we do not need it)
```
    atlas schema inspect -u "postgres://root:@localhost:26257?sslmode=disable" > ./internal/database/dialer/cockroachdb/table/schema.hcl
```

## Convert goose migrations to Atlas migrations
If you have existing goose migrations under `migration` folder (relative to your current directory), you can convert it to an Atlas migration.
```
    atlas migrate import --from file://migration?format=goose --to file://atlas
```

## Init database from an Atlas schema.hcl
If you have an existing schema.hcl file, you can initialize the database with it.
```
    atlas schema apply --env dev --to file://internal/database/dialer/cockroachdb/table/schema.hcl
```

## Create a migration
After modifying any gorm models, create a migration file against your local development database.
```
    atlas migrate diff migration_name --env dev
```

## Apply migrations
Apply the migration to any database via `--url` flag.
```
    atlas migrate apply --dir file://migration --url "postgres://root:@localhost:26257?sslmode=disable"
```

Or your local development database.
```
    atlas migrate apply --env dev
```

For more information, please refer to:
1. [Atlas Quick Introduction](https://atlasgo.io/getting-started/)
2. [Automatic migration planning for GORM](https://atlasgo.io/guides/orms/gorm)
