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
var maxRecentHandles = 100000

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

// FindActivity finds an Activity by id.
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

func (c *client) FindFederatedActivities(ctx context.Context, query model.FederatedActivitiesQuery) ([]*activityx.Activity, error) {
	if c.partition {
		return c.findFederatedActivitiesPartitioned(ctx, query)
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

func (c *client) UpdateRecentMastodonHandles(ctx context.Context, handles []*model.MastodonHandle) error {
	logger := zap.L()
	logger.Info("Starting UpdateRecentMastodonHandles", zap.Int("handleCount", len(handles)))

	// ToDo: remove hardcoded db table name and fix the function logic
	tableName := "dataset_mastodon_update_handles"
	// Construct the SQL query
	sql, valueArgs := buildInsertQuery(handles, tableName)

	logger.Debug("Generated SQL", zap.String("sql", sql), zap.Any("args", valueArgs))

	// Execute the update query
	result := c.database.WithContext(ctx).Exec(sql, valueArgs...)
	if result.Error != nil {
		logger.Error("Error updating recent handles", zap.Error(result.Error))
		return fmt.Errorf("failed to update recent handles: %w", result.Error)
	}

	// Trim old handles
	if err := trimOldHandles(ctx, c, tableName); err != nil {
		logger.Error("Error trimming old handles", zap.Error(err))
		return err
	}

	logger.Info("Completed UpdateRecentMastodonHandles",
		zap.Int64("updatedHandles", result.RowsAffected))

	return nil
}

// buildInsertQuery generates the SQL insert query and corresponding values for Mastodon handles
func buildInsertQuery(handles []*model.MastodonHandle, tableName string) (string, []interface{}) {
	valueStrings := make([]string, len(handles))
	valueArgs := make([]interface{}, 0, len(handles)*2)

	for i, h := range handles {
		valueStrings[i] = fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2)

		valueArgs = append(valueArgs, h.Handle, h.LastUpdated)
	}

	// Insert SQL with ON CONFLICT clause
	sql := fmt.Sprintf(`
        INSERT INTO %s (handle, last_updated)
        VALUES %s
        ON CONFLICT (handle) DO UPDATE SET last_updated = EXCLUDED.last_updated;`,
		tableName, strings.Join(valueStrings, ","),
	)

	return sql, valueArgs
}

// trimOldHandles removes outdated handles from the dataset
func trimOldHandles(ctx context.Context, c *client, tableName string) error {
	trimResult := c.database.WithContext(ctx).Exec(fmt.Sprintf(`
        DELETE FROM %s
        WHERE handle NOT IN (
            SELECT handle FROM %s
            ORDER BY last_updated DESC
            LIMIT $1
        )`, tableName, tableName), maxRecentHandles)

	if trimResult.Error != nil {
		return fmt.Errorf("failed to trim old handles: %w", trimResult.Error)
	}

	return nil
}

func (c *client) GetUpdatedMastodonHandles(ctx context.Context, since uint64, limit int, cursor string) (*model.PaginatedMastodonHandles, error) {
	var result []model.MastodonHandle

	var totalCount int64

	// ToDo: remove hardcoded db table name
	// Query for the handles table
	query := c.database.WithContext(ctx).
		Table("dataset_mastodon_update_handles")

	// If 'since' is greater than 0, we're fetching updated handles; otherwise, fetch all handles
	if since > 0 {
		query = query.Where("last_updated > ?", since)
	}

	// Count total handles (or updated handles)
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, fmt.Errorf("failed to count handles: %w", err)
	}

	// Apply cursor if provided
	if cursor != "" {
		var cursorLastUpdated uint64

		var cursorHandle string

		// Decode the cursor format: "<last_updated>:<handle>"
		_, err := fmt.Sscanf(cursor, "%d:%s", &cursorLastUpdated, &cursorHandle)
		if err != nil {
			return nil, fmt.Errorf("invalid cursor format: %w", err)
		}

		query = query.Where("last_updated < ? OR (last_updated = ? AND handle > ?)", cursorLastUpdated, cursorLastUpdated, cursorHandle)
	}

	// Fetch handles with their last_updated timestamps
	err := query.
		Select("handle, last_updated").
		Order("last_updated DESC, handle ASC").
		Limit(limit + 1).
		Find(&result).
		Error

	if err != nil {
		return nil, fmt.Errorf("failed to get handles: %w", err)
	}

	// Determine the next cursor if there are more than 'limit' handles
	var nextCursor string

	if len(result) > limit {
		lastHandle := result[limit-1]

		// Create the next cursor from the 'last_updated' and 'handle'
		nextCursor = fmt.Sprintf("%d:%s", lastHandle.LastUpdated, lastHandle.Handle)

		// Remove the extra handle for pagination
		result = result[:limit]
	}

	// Extract just the handles for the response
	handles := make([]string, len(result))
	for i, r := range result {
		handles[i] = r.Handle
	}

	return &model.PaginatedMastodonHandles{
		Handles:    handles,
		TotalCount: totalCount,
		NextCursor: nextCursor,
	}, nil
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
