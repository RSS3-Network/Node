package cockroach

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var _ database.Client = (*client)(nil)

type client struct {
	mode     database.Mode
	database *gorm.DB
}

// Migrate migrates the database.
func (c *client) Migrate(_ context.Context) error {
	goose.SetBaseFS(database.MigrationFS)
	goose.SetTableName("versions")
	goose.SetLogger(&database.SugaredLogger{Logger: zap.L().Sugar()})

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("set migration dialect: %w", err)
	}

	connector, err := c.database.DB()
	if err != nil {
		return fmt.Errorf("get database connector: %w", err)
	}

	return goose.Up(connector, "migration")
}

// WithTransaction executes a transaction.
func (c *client) WithTransaction(ctx context.Context, transactionFunction func(ctx context.Context, client database.Client) error, transactionOptions ...*sql.TxOptions) error {
	transaction, err := c.Begin(ctx, transactionOptions...)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	if err := transactionFunction(ctx, transaction); err != nil {
		_ = transaction.Rollback()

		return fmt.Errorf("execute transaction: %w", err)
	}

	if err := transaction.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

// Begin begins a transaction.
func (c *client) Begin(ctx context.Context, transactionOptions ...*sql.TxOptions) (database.Client, error) {
	transaction := c.database.WithContext(ctx).Begin(transactionOptions...)
	if err := transaction.Error; err != nil {
		return nil, err
	}

	return &client{
		database: transaction,
	}, nil
}

// Rollback rolls back a transaction.
func (c *client) Rollback() error {
	return c.database.Rollback().Error
}

// Commit commits a transaction.
func (c *client) Commit() error {
	return c.database.Commit().Error
}

// createShardedTable creates a sharded table.
func (c *client) createShardedTable(ctx context.Context, name, template string) error {
	statement := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS "%s" (LIKE "%s" INCLUDING ALL);`, name, template)

	return c.database.WithContext(ctx).Exec(statement).Error
}

// SaveFeeds saves feeds and indexes to the database.
func (c *client) SaveFeeds(ctx context.Context, feeds []schema.Feed) error {
	switch c.mode {
	case database.ModeSharded:
		return c.saveFeedsSharded(ctx, feeds)
	case database.ModeSingle:
		// TODO
		return fmt.Errorf("not implemented")
	}

	return nil
}

// Dial dials a database.
func Dial(_ context.Context, dataSourceName string, mode database.Mode) (database.Client, error) {
	var err error

	instance := client{
		mode: mode,
	}

	logger := zapgorm2.New(zap.L())
	logger.SetAsDefault()

	config := gorm.Config{
		Logger: logger,
	}

	if instance.database, err = gorm.Open(postgres.Open(dataSourceName), &config); err != nil {
		return nil, err
	}

	return &instance, nil
}
