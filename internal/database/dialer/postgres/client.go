package postgres

import (
	"context"
	"database/sql"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pressly/goose/v3"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/database/dialer/postgres/table"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/node/internal/engine"
	mirror_model "github.com/rss3-network/node/internal/engine/worker/decentralized/contract/mirror/model"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	networkx "github.com/rss3-network/protocol-go/schema/network"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"moul.io/zapgorm2"
)

var _ database.Client = (*client)(nil)

const (
	DefaultKVClosedTimestampTargetDuration = 3 * time.Second
)

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
	transactionFunc := func() error {
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

	retryIfFunc := func(err error) bool {
		// https://www.cockroachlabs.com/docs/stable/transaction-retry-error-reference#retry_serializable
		return strings.Contains(err.Error(), "TransactionRetryWithProtoRefreshError: TransactionRetryError:")
	}

	onRetryFunc := func(n uint, err error) {
		zap.L().Error("execute transaction", zap.Uint("retry", n), zap.Error(err))
	}

	return retry.Do(
		transactionFunc,
		retry.RetryIf(retryIfFunc), retry.MaxJitter(DefaultKVClosedTimestampTargetDuration), retry.DelayType(retry.RandomDelay), retry.OnRetry(onRetryFunc), retry.Attempts(0),
	)
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

func (c *client) LoadCheckpoint(ctx context.Context, id string, network networkx.Network, worker string) (*engine.Checkpoint, error) {
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

func (c *client) LoadCheckpoints(ctx context.Context, id string, network networkx.Network, worker string) ([]*engine.Checkpoint, error) {
	databaseStatement := c.database.WithContext(ctx)

	var checkpoints []*table.Checkpoint

	zap.L().Info("load checkpoints", zap.String("id", id), zap.String("network", network.String()), zap.String("worker", worker))

	if id != "" {
		databaseStatement = databaseStatement.Where("id = ?", id)
	}

	if network != networkx.Unknown {
		databaseStatement = databaseStatement.Where("network = ?", network)
	}

	if worker != "" {
		databaseStatement = databaseStatement.Where("worker= ?", worker)
	}

	if err := databaseStatement.Find(&checkpoints).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	result := make([]*engine.Checkpoint, 0, len(checkpoints))

	for _, checkpoint := range checkpoints {
		data, err := checkpoint.Export()
		if err != nil {
			return nil, err
		}

		result = append(result, data)
	}

	return result, nil
}

func (c *client) SaveCheckpoint(ctx context.Context, checkpoint *engine.Checkpoint) error {
	spanStartOptions := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("checkpoint.id", checkpoint.ID),
		),
	}

	ctx, span := otel.Tracer("").Start(ctx, "Database saveCheckpoint", spanStartOptions...)
	defer span.End()

	clauses := []clause.Expression{
		clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"state":       checkpoint.State,
				"updated_at":  time.Now(),
				"index_count": gorm.Expr("checkpoints.index_count + ?", checkpoint.IndexCount),
			}),
		},
	}

	var value table.Checkpoint
	if err := value.Import(checkpoint); err != nil {
		return err
	}

	return c.database.WithContext(ctx).Clauses(clauses...).Create(&value).Error
}

// SaveActivities saves activities and indexes to the database.
func (c *client) SaveActivities(ctx context.Context, activities []*activityx.Activity) error {
	spanStartOptions := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.Int("activities", len(activities)),
		),
	}

	ctx, span := otel.Tracer("").Start(ctx, "Database saveActivities", spanStartOptions...)
	defer span.End()

	if c.partition {
		return c.saveActivitiesPartitioned(ctx, activities)
	}

	return fmt.Errorf("not implemented")
}

// LoadDatasetMirrorPost LoadMirrorPost loads a mirror post with tx id and origin content digest.
func (c *client) LoadDatasetMirrorPost(ctx context.Context, originContentDigest string) (*mirror_model.DatasetMirrorPost, error) {
	var value table.DatasetMirrorPost

	// Get first mirror dataset post record using origin content digest
	if err := c.database.WithContext(ctx).
		Where("origin_content_digest = ?", originContentDigest).
		First(&value).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		// Initialize a default post.
		value = table.DatasetMirrorPost{
			OriginContentDigest: originContentDigest,
		}
	}

	return value.Export()
}

// SaveDatasetMirrorPost saves a mirror post's tx id and origin content digest to the database.
func (c *client) SaveDatasetMirrorPost(ctx context.Context, post *mirror_model.DatasetMirrorPost) error {
	clauses := []clause.Expression{
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		},
	}

	var value table.DatasetMirrorPost
	if err := value.Import(post); err != nil {
		return err
	}

	return c.database.WithContext(ctx).Clauses(clauses...).Create(&value).Error
}

// FindActivity finds a Activity by id.
func (c *client) FindActivity(ctx context.Context, query model.ActivityQuery) (*activityx.Activity, *int, error) {
	if c.partition {
		return c.findActivityPartitioned(ctx, query)
	}

	return nil, nil, fmt.Errorf("not implemented")
}

// FindActivities finds Activities.
func (c *client) FindActivities(ctx context.Context, query model.ActivitiesQuery) ([]*activityx.Activity, error) {
	if c.partition {
		return c.findActivitiesPartitioned(ctx, query)
	}

	return nil, fmt.Errorf("not implemented")
}

// DeleteExpiredActivities deletes expired activities.
func (c *client) DeleteExpiredActivities(ctx context.Context, network networkx.Network, timestamp time.Time) error {
	if c.partition {
		return c.deleteExpiredActivitiesPartitioned(ctx, network, timestamp)
	}

	return fmt.Errorf("not implemented")
}

// LoadDatasetFarcasterProfile loads a profile.
func (c *client) LoadDatasetFarcasterProfile(ctx context.Context, fid int64) (*model.Profile, error) {
	var value table.DatasetFarcasterProfile

	if err := c.database.WithContext(ctx).
		Where("fid = ?", fid).
		First(&value).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		// Initialize a default profile.
		value = table.DatasetFarcasterProfile{
			Fid: fid,
		}
	}

	return value.Export()
}

// SaveDatasetFarcasterProfile saves a profile.
func (c *client) SaveDatasetFarcasterProfile(ctx context.Context, profile *model.Profile) error {
	clauses := []clause.Expression{
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "fid"}},
			UpdateAll: true,
		},
	}

	var value table.DatasetFarcasterProfile
	if err := value.Import(profile); err != nil {
		return err
	}

	return c.database.WithContext(ctx).Clauses(clauses...).Create(&value).Error
}

// LoadDatasetENSNamehash accepts an ENS namehash, and response with a record containing its original name string
func (c *client) LoadDatasetENSNamehash(ctx context.Context, hash common.Hash) (*model.ENSNamehash, error) {
	var value table.DatasetENSNamehash

	if err := c.database.WithContext(ctx).
		Where("hash = ?", hash).
		First(&value).
		Error; err != nil {
		// No such ENS
		return nil, err
	}

	return value.Export()
}

// SaveDatasetENSNamehash save an ENS namehash into database for further queries
func (c *client) SaveDatasetENSNamehash(ctx context.Context, namehash *model.ENSNamehash) error {
	clauses := []clause.Expression{
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "hash"}},
			UpdateAll: true,
		},
	}

	var value table.DatasetENSNamehash
	if err := value.Import(namehash); err != nil {
		return err
	}

	return c.database.WithContext(ctx).Clauses(clauses...).Create(&value).Error
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
		Logger:                 logger,
		SkipDefaultTransaction: true,
	}

	if instance.database, err = gorm.Open(postgres.Open(dataSourceName), &config); err != nil {
		return nil, err
	}

	if instance.partition {
		instance.loadIndexesPartitionTables(ctx)
	}

	return &instance, nil
}
