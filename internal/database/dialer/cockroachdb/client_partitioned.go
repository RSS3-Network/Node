package cockroachdb

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sort"
	"sync"
	"time"

	"github.com/rss3-network/node/internal/database/dialer/cockroachdb/table"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var indexesTables sync.Map

// createPartitionTable creates a partition table.
func (c *client) createPartitionTable(ctx context.Context, name, template string) error {
	statement := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS "%s" (LIKE "%s" INCLUDING ALL);`, name, template)

	err := c.database.WithContext(ctx).Exec(statement).Error
	if err != nil {
		return err
	}

	if template == (*table.Index).TableName(nil) {
		indexesTables.Store(name, struct{}{})
	}

	return nil
}

// findIndexesPartitionTable finds partition table names of indexes in the past year.
func (c *client) findIndexesPartitionTables(_ context.Context, index table.Index) ([]string, error) {
	partitionedNames := make([]string, 0)

	for i := 0; i <= 4; i++ {
		if index.Timestamp.Unix() < time.Now().AddDate(-1, 0, 0).Unix() {
			break
		}

		if _, exists := indexesTables.Load(index.PartitionName()); exists {
			partitionedNames = append(partitionedNames, index.PartitionName())
		}

		year := index.Timestamp.Year()
		month := index.Timestamp.Month()

		index.Timestamp = time.Date(lo.Ternary(month < 3, year-1, year), lo.Ternary(month < 3, month+9, month-3), index.Timestamp.Day(), 23, 59, 59, 1e9-1, time.Local)
	}

	return partitionedNames, nil
}

// loadIndexesPartitionTables loads indexes partition tables.
func (c *client) loadIndexesPartitionTables(ctx context.Context) {
	result := make([]string, 0)

	if err := c.database.WithContext(ctx).Table("pg_tables").Where("tablename LIKE ?", fmt.Sprintf("%s_%%", (*table.Index).TableName(nil))).Pluck("tablename", &result).Error; err != nil {
		zap.L().Error("failed to load indexes partition tables", zap.Error(err))
	}

	for _, tableName := range result {
		indexesTables.Store(tableName, struct{}{})
	}
}

// saveFeedsPartitioned saves feeds in partitioned tables.
func (c *client) saveFeedsPartitioned(ctx context.Context, feeds []*schema.Feed) error {
	partitions := make(map[string][]*schema.Feed)

	// Group feeds by partition name.
	for _, feed := range feeds {
		var feedTable table.Feed

		name := feedTable.PartitionName(feed)

		if _, exists := partitions[name]; !exists {
			partitions[name] = make([]*schema.Feed, 0)
		}

		partitions[name] = append(partitions[name], feed)
	}

	errorGroup, ctx := errgroup.WithContext(ctx)

	// Save feeds and indexes in parallel.
	for name, feeds := range partitions {
		name, feeds := name, feeds

		errorGroup.Go(func() error {
			tableFeeds := make(table.Feeds, 0)

			if err := tableFeeds.Import(feeds); err != nil {
				return err
			}

			if len(tableFeeds) == 0 {
				return nil
			}

			// #nosec
			if err := c.createPartitionTable(ctx, name, tableFeeds[0].TableName()); err != nil {
				return err
			}

			onConflict := clause.OnConflict{
				Columns: []clause.Column{
					{
						Name: "id",
					},
				},
				UpdateAll: true,
			}

			if err := c.database.WithContext(ctx).
				Table(name).
				Clauses(onConflict).
				CreateInBatches(tableFeeds, math.MaxUint8).
				Error; err != nil {
				return err
			}

			return c.saveIndexesPartitioned(ctx, feeds)
		})
	}

	return errorGroup.Wait()
}

// findFeedPartitioned finds a feed by id.
func (c *client) findFeedPartitioned(ctx context.Context, query model.FeedQuery) (*schema.Feed, *int, error) {
	index, err := c.findIndexPartitioned(ctx, query)
	if err != nil {
		return nil, nil, fmt.Errorf("first index: %w", err)
	}

	if index == nil {
		return nil, nil, nil
	}

	feed := &table.Feed{
		ID:        index.ID,
		Network:   index.Network,
		Timestamp: index.Timestamp,
	}

	if err := c.database.WithContext(ctx).Table(feed.PartitionName(nil)).Where("id = ?", index.ID).Limit(1).Find(&feed).Error; err != nil {
		return nil, nil, fmt.Errorf("find feed: %w", err)
	}

	result, err := feed.Export(index)
	if err != nil {
		return nil, nil, fmt.Errorf("export feed: %w", err)
	}

	page := math.Ceil(float64(len(feed.Actions)) / float64(query.ActionLimit))

	feed.Actions = lo.Slice(feed.Actions, query.ActionLimit*(query.ActionPage-1), query.ActionLimit*query.ActionPage)

	return result, lo.ToPtr(int(page)), nil
}

// findFeedsPartitioned finds feeds.
func (c *client) findFeedsPartitioned(ctx context.Context, query model.FeedsQuery) ([]*schema.Feed, error) {
	indexes, err := c.findIndexesPartitioned(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("find indexes: %w", err)
	}

	partition := lo.GroupBy(indexes, func(query *table.Index) string {
		feed := table.Feed{
			ID:        query.ID,
			Network:   query.Network,
			Timestamp: query.Timestamp,
		}

		return feed.PartitionName(nil)
	})

	var (
		result = make([]*schema.Feed, 0)
		locker sync.Mutex
	)

	errorGroup, errorCtx := errgroup.WithContext(ctx)

	for tableName, index := range partition {
		tableName, index := tableName, index

		ids := lo.Map(index, func(index *table.Index, _ int) string {
			return index.ID
		})

		errorGroup.Go(func() error {
			tableFeeds := make(table.Feeds, 0)

			if err := c.database.WithContext(errorCtx).Table(tableName).Where("id IN ?", lo.Uniq(ids)).Find(&tableFeeds).Error; err != nil {
				zap.L().Error("failed to find feeds", zap.Error(err), zap.String("tableName", tableName))

				return err
			}

			locker.Lock()
			defer locker.Unlock()

			feeds, err := tableFeeds.Export(index)
			if err != nil {
				return err
			}

			result = append(result, feeds...)

			return nil
		})
	}

	if err := errorGroup.Wait(); err != nil {
		return nil, err
	}

	lo.ForEach(result, func(feed *schema.Feed, i int) {
		result[i].Actions = lo.Slice(feed.Actions, 0, query.ActionLimit)
	})

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Timestamp > result[j].Timestamp
	})

	return result, nil
}

// saveIndexesPartitioned saves indexes in partitioned tables.
func (c *client) saveIndexesPartitioned(ctx context.Context, feeds []*schema.Feed) error {
	indexes := make(table.Indexes, 0)

	if err := indexes.Import(feeds); err != nil {
		return err
	}

	if len(indexes) == 0 {
		return nil
	}

	// #nosec
	if err := c.createPartitionTable(ctx, indexes[0].PartitionName(), indexes[0].TableName()); err != nil {
		return fmt.Errorf("create partition table: %w", err)
	}

	// Delete indexes with the same feed id and network.
	pkIndexes := make(map[string][]string)

	for _, index := range indexes {
		if _, ok := pkIndexes[index.ID+index.Network.String()]; ok {
			continue
		}

		pkIndexes[index.ID+index.Network.String()] = []string{
			index.ID,
			index.Network.String(),
		}
	}

	conditions := lo.MapToSlice(pkIndexes, func(_ string, value []string) []string {
		return value
	})

	// #nosec
	if err := c.database.WithContext(ctx).
		Table(indexes[0].PartitionName()).
		Where("(id, network) IN (?)", conditions).
		Delete(&table.Indexes{}).
		Error; err != nil {
		return fmt.Errorf("delete indexes: %w", err)
	}

	// Save indexes.
	onConflict := clause.OnConflict{
		Columns: []clause.Column{
			{
				Name: "id",
			},
			{
				Name: "network",
			},
			{
				Name: "owner",
			},
		},
		UpdateAll: true,
	}

	// #nosec
	return c.database.WithContext(ctx).
		Table(indexes[0].PartitionName()).
		Clauses(onConflict).
		CreateInBatches(indexes, math.MaxUint8).
		Error
}

// findIndexPartitioned finds a feed by id.
func (c *client) findIndexPartitioned(ctx context.Context, query model.FeedQuery) (*table.Index, error) {
	index := table.Index{
		Timestamp: time.Now(),
	}

	tables, err := c.findIndexesPartitionTables(ctx, index)
	if err != nil {
		return nil, fmt.Errorf("find indexes partition tables: %w", err)
	}

	ctx, cancel := context.WithCancel(ctx)
	errorGroup, errorContext := errgroup.WithContext(ctx)
	resultChan := make(chan *table.Index, len(tables))
	errorChan := make(chan error)
	stopChan := make(chan struct{})

	for _, tableName := range tables {
		tableName := tableName

		errorGroup.Go(func() error {
			var result table.Index

			if err := c.buildFindIndexStatement(errorContext, tableName, query).Limit(1).Find(&result).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				zap.L().Error("failed to find first index", zap.Error(err), zap.String("partition table", tableName))

				return fmt.Errorf("find first index: %w", err)
			}

			if lo.IsEmpty(result.ID) {
				return nil
			}

			select {
			case <-stopChan:
				return nil
			case resultChan <- &result:
				return nil
			}
		})
	}

	go func() {
		defer close(errorChan)

		errorChan <- errorGroup.Wait()
	}()

	defer cancel()

	count := 0

	for {
		select {
		case err := <-errorChan:
			if err != nil {
				return nil, fmt.Errorf("failed to wait result: %w", err)
			}
		case data := <-resultChan:
			count++

			if data != nil {
				close(stopChan)
				return data, nil
			}

			if count == len(tables) {
				return nil, nil
			}
		}
	}
}

// findIndexesPartitioned finds indexes.
func (c *client) findIndexesPartitioned(ctx context.Context, query model.FeedsQuery) ([]*table.Index, error) {
	index := table.Index{
		Timestamp: time.Now(),
	}

	if query.EndTimestamp != nil && lo.FromPtr(query.EndTimestamp) < uint64(index.Timestamp.Unix()) {
		index.Timestamp = time.Unix(int64(lo.FromPtr(query.EndTimestamp)), 0)
	}

	if query.Cursor != nil && query.Cursor.Timestamp < uint64(index.Timestamp.Unix()) {
		index.Timestamp = time.Unix(int64(query.Cursor.Timestamp), 0)
	}

	partitionedNames, err := c.findIndexesPartitionTables(ctx, index)
	if err != nil {
		zap.L().Error("failed to build partitioned indexes past year", zap.Error(err))

		return nil, err
	}

	if len(partitionedNames) == 0 {
		return nil, nil
	}

	ctx, cancel := context.WithCancel(ctx)
	errorGroup, errorContext := errgroup.WithContext(ctx)
	indexes := make([][]*table.Index, len(partitionedNames))
	resultChan := make(chan int, len(partitionedNames))
	errorChan := make(chan error)
	stopChan := make(chan struct{})

	var mutex sync.RWMutex

	for partitionedIndex, partitionedName := range partitionedNames {
		partitionedIndex, partitionedName := partitionedIndex, partitionedName

		errorGroup.Go(func() error {
			var result []*table.Index

			databaseStatement := c.buildFindIndexesStatement(errorContext, partitionedName, query)

			if err := databaseStatement.Find(&result).Error; err != nil {
				return fmt.Errorf("failed to find indexes: %w", err)
			}

			mutex.Lock()
			indexes[partitionedIndex] = result
			mutex.Unlock()

			select {
			case <-stopChan:
				return nil
			case resultChan <- partitionedIndex:
				return nil
			}
		})
	}

	go func() {
		defer close(errorChan)

		errorChan <- errorGroup.Wait()
	}()

	defer func() {
		cancel()
	}()

	for {
		select {
		case err := <-errorChan:
			if err != nil {
				return nil, fmt.Errorf("failed to wait result: %w", err)
			}
		case <-resultChan:
			result := make([]*table.Index, 0, query.Limit)
			flag := true

			mutex.RLock()

			for _, data := range indexes {
				data := data

				if data == nil {
					flag = false
					break
				}

				result = append(result, data...)

				if len(result) >= query.Limit {
					close(stopChan)
					mutex.RUnlock()

					return result[:query.Limit], nil
				}
			}

			mutex.RUnlock()

			if flag {
				return result, nil
			}
		}
	}
}

// buildFindIndexStatement builds the query index statement.
func (c *client) buildFindIndexStatement(ctx context.Context, partitionedName string, query model.FeedQuery) *gorm.DB {
	databaseStatement := c.database.WithContext(ctx).Table(partitionedName)

	if query.ID != nil {
		databaseStatement = databaseStatement.Where("id = ?", lo.FromPtr(query.ID))
	}

	if query.Network != nil {
		databaseStatement = databaseStatement.Where("network = ?", lo.FromPtr(query.Network))
	}

	if query.Owner != nil {
		databaseStatement = databaseStatement.Where("owner = ?", lo.FromPtr(query.Owner))
	}

	return databaseStatement
}

// buildFindIndexesStatement builds the query indexes statement.
func (c *client) buildFindIndexesStatement(ctx context.Context, partition string, query model.FeedsQuery) *gorm.DB {
	databaseStatement := c.database.WithContext(ctx).Table(partition)

	var table *string

	if query.Distinct != nil && lo.FromPtr(query.Distinct) {
		databaseStatement = databaseStatement.Select("DISTINCT (id) id, timestamp, index, network")
	}

	if query.Owner != nil {
		table = lo.ToPtr(fmt.Sprintf(`%s@idx_indexes_owner`, partition))
		databaseStatement = databaseStatement.Where("owner = ?", query.Owner)
	}

	if len(query.Owners) > 0 {
		table = lo.ToPtr(fmt.Sprintf(`%s@idx_indexes_owner`, partition))
		databaseStatement = databaseStatement.Where("owner IN ?", query.Owners)
	}

	if query.Status != nil {
		databaseStatement = databaseStatement.Where("status = ?", query.Status)
	}

	if query.Direction != nil {
		databaseStatement = databaseStatement.Where("direction = ?", query.Direction)
	}

	if query.StartTimestamp != nil {
		databaseStatement = databaseStatement.Where("timestamp >= ?", time.Unix(int64(*query.StartTimestamp), 0))
	}

	if query.EndTimestamp != nil {
		databaseStatement = databaseStatement.Where("timestamp <= ?", time.Unix(int64(*query.EndTimestamp), 0))
	}

	if query.Platform != nil {
		databaseStatement = databaseStatement.Where("platform = ?", query.Platform)
	}

	if len(query.Platforms) > 0 {
		databaseStatement = databaseStatement.Where("platform IN ?", query.Platforms)
	}

	if len(query.Tags) > 0 {
		databaseStatement = databaseStatement.Where("tag IN ?", query.Tags)
	}

	if len(query.Types) > 0 {
		databaseStatement = databaseStatement.Where("type IN ?", query.Types)
	}

	if len(query.Network) > 0 {
		databaseStatement = databaseStatement.Where("network IN ?", query.Network)
	}

	if query.Cursor != nil && query.Cursor.Timestamp > 0 {
		databaseStatement = databaseStatement.Where("timestamp < ? OR (timestamp = ? AND index < ?)", time.Unix(int64(query.Cursor.Timestamp), 0), time.Unix(int64(query.Cursor.Timestamp), 0), query.Cursor.Index)
	}

	if table != nil {
		databaseStatement.Statement.TableExpr.SQL = *table
	}

	return databaseStatement.Order("timestamp DESC, index DESC").Limit(query.Limit)
}
