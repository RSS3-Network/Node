package cockroachdb

import (
	"context"
	"database/sql"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/database/dialer/cockroachdb/table"
	"github.com/naturalselectionlabs/rss3-node/internal/database/model"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"moul.io/zapgorm2"
)

var _ database.Client = (*client)(nil)

//go:embed migration/*.sql
var migrationFS embed.FS

type client struct {
	partition bool
	database  *gorm.DB
}

// Migrate migrates the database.
func (c *client) Migrate(ctx context.Context) error {
	goose.SetBaseFS(migrationFS)
	goose.SetTableName("versions")
	goose.SetLogger(&database.SugaredLogger{Logger: zap.L().Sugar()})

	if err := goose.SetDialect(new(postgres.Dialector).Name()); err != nil {
		return fmt.Errorf("set migration dialect: %w", err)
	}

	connector, err := c.database.DB()
	if err != nil {
		return fmt.Errorf("get database connector: %w", err)
	}

	return goose.UpContext(ctx, connector, "migration")
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
		partition: c.partition,
		database:  transaction,
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

func (c *client) LoadCheckpoint(ctx context.Context, id string, network filter.Network, worker string) (*engine.Checkpoint, error) {
	var value table.Checkpoint

	zap.L().Info("load checkpoint", zap.String("id", id), zap.String("network", network.String()), zap.String("worker", worker))

	if err := c.database.WithContext(ctx).
		Where("id = ? AND network = ? AND worker = ?", id, network, worker).
		First(&value).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		// Initialize a default checkpoint.
		value = table.Checkpoint{
			ID:        id,
			Network:   network,
			Worker:    worker,
			State:     json.RawMessage("{}"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	}

	return value.Export()
}

func (c *client) SaveCheckpoint(ctx context.Context, checkpoint *engine.Checkpoint) error {
	clauses := []clause.Expression{
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"state", "updated_at"}),
		},
	}

	var value table.Checkpoint
	if err := value.Import(checkpoint); err != nil {
		return err
	}

	return c.database.WithContext(ctx).Clauses(clauses...).Create(&value).Error
}

// SaveFeeds saves feeds and indexes to the database.
func (c *client) SaveFeeds(ctx context.Context, feeds []*schema.Feed) error {
	if c.partition {
		return c.saveFeedsPartitioned(ctx, feeds)
	}

	return fmt.Errorf("not implemented")
}

// FirstFeed finds a feed by id.
func (c *client) FirstFeed(ctx context.Context, query model.FeedQuery) (*schema.Feed, *int, error) {
	if c.partition {
		return c.firstFeedPartitioned(ctx, query)
	}

	return nil, nil, fmt.Errorf("not implemented")
}

// FindFeeds finds feeds.
func (c *client) FindFeeds(ctx context.Context, query model.FeedsQuery) ([]*schema.Feed, error) {
	if c.partition {
		return c.findFeedsPartitioned(ctx, query)
	}

	return nil, fmt.Errorf("not implemented")
}

// Dial dials a database.
func Dial(ctx context.Context, dataSourceName string, partition bool) (database.Client, error) {
	var err error

	instance := client{
		partition: partition,
	}

	logger := zapgorm2.New(zap.L())
	logger.SetAsDefault()

	config := gorm.Config{
		Logger: logger,
	}

	if instance.database, err = gorm.Open(postgres.Open(dataSourceName), &config); err != nil {
		return nil, err
	}

	if instance.partition {
		go instance.autoLoadIndexesPartitionTables(ctx)
	}

	return &instance, nil
}
