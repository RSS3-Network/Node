package postgres

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sort"
	"sync"
	"time"

	"github.com/rss3-network/node/internal/database/dialer/postgres/table"
	"github.com/rss3-network/node/internal/database/model"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"github.com/sourcegraph/conc/pool"
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

// findPartitionTableExists check if a partition table exists.
func (c *client) findPartitionTableExists(ctx context.Context, name string) (bool, error) {
	statement := fmt.Sprintf(`SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = '%s')`, name)

	var exists bool

	if err := c.database.WithContext(ctx).Raw(statement).Scan(&exists).Error; err != nil {
		return false, err
	}

	return exists, nil
}

// findIndexesPartitionTable finds partition table names of indexes in the past year.
func (c *client) findIndexesPartitionTables(_ context.Context, index table.Index) []string {
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

	return partitionedNames
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

// saveActivitiesPartitioned saves Activities in partitioned tables.
func (c *client) saveActivitiesPartitioned(ctx context.Context, activities []*activityx.Activity) error {
	partitions := make(map[string][]*activityx.Activity)

	// Group activities by partition name.
	for _, a := range activities {
		var activityTable table.Activity

		name := activityTable.PartitionName(a)

		if _, exists := partitions[name]; !exists {
			partitions[name] = make([]*activityx.Activity, 0)
		}

		partitions[name] = append(partitions[name], a)
	}

	errorGroup, ctx := errgroup.WithContext(ctx)

	// Save activities and indexes in parallel.
	for name, activities := range partitions {
		name, activities := name, activities

		errorGroup.Go(func() error {
			tableActivities := make(table.Activities, 0)

			if err := tableActivities.Import(activities); err != nil {
				return err
			}

			if len(tableActivities) == 0 {
				return nil
			}

			// #nosec
			if err := c.createPartitionTable(ctx, name, tableActivities[0].TableName()); err != nil {
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
				CreateInBatches(tableActivities, math.MaxUint8).
				Error; err != nil {
				return err
			}

			return c.saveIndexesPartitioned(ctx, activities)
		})
	}

	return errorGroup.Wait()
}

// findActivityPartitioned finds an activity  by id.
func (c *client) findActivityPartitioned(ctx context.Context, query model.ActivityQuery) (*activityx.Activity, *int, error) {
	matchedActivity, err := c.findIndexPartitioned(ctx, query)
	if err != nil {
		return nil, nil, fmt.Errorf("first matchedActivity: %w", err)
	}

	if matchedActivity == nil {
		return nil, nil, nil
	}

	activity := &table.Activity{
		ID:        matchedActivity.ID,
		Network:   matchedActivity.Network,
		Timestamp: matchedActivity.Timestamp,
	}

	if err := c.database.WithContext(ctx).Table(activity.PartitionName(nil)).Where("id = ?", matchedActivity.ID).Limit(1).Find(&activity).Error; err != nil {
		return nil, nil, fmt.Errorf("find activity: %w", err)
	}

	result, err := activity.Export(matchedActivity)
	if err != nil {
		return nil, nil, fmt.Errorf("export activity: %w", err)
	}

	page := math.Ceil(float64(len(activity.Actions)) / float64(query.ActionLimit))

	activity.Actions = lo.Slice(activity.Actions, query.ActionLimit*(query.ActionPage-1), query.ActionLimit*query.ActionPage)

	return result, lo.ToPtr(int(page)), nil
}

// findActivitiesPartitioned finds activities.
func (c *client) findActivitiesPartitioned(ctx context.Context, query model.ActivitiesQuery) ([]*activityx.Activity, error) {
	indexes, err := c.findIndexesPartitioned(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("find indexes: %w", err)
	}

	partition := lo.GroupBy(indexes, func(query *table.Index) string {
		activity := table.Activity{
			ID:        query.ID,
			Network:   query.Network,
			Timestamp: query.Timestamp,
		}

		return activity.PartitionName(nil)
	})

	var (
		result = make([]*activityx.Activity, 0)
		locker sync.Mutex
	)

	errorGroup, errorCtx := errgroup.WithContext(ctx)

	for tableName, index := range partition {
		tableName, index := tableName, index

		ids := lo.Map(index, func(index *table.Index, _ int) string {
			return index.ID
		})

		errorGroup.Go(func() error {
			tableActivities := make(table.Activities, 0)

			if err := c.database.WithContext(errorCtx).Table(tableName).Where("id IN ?", lo.Uniq(ids)).Find(&tableActivities).Error; err != nil {
				zap.L().Error("failed to find activities", zap.Error(err), zap.String("tableName", tableName))

				return err
			}

			locker.Lock()
			defer locker.Unlock()

			activities, err := tableActivities.Export(index)
			if err != nil {
				return err
			}

			result = append(result, activities...)

			return nil
		})
	}

	if err := errorGroup.Wait(); err != nil {
		return nil, err
	}

	lo.ForEach(result, func(activity *activityx.Activity, i int) {
		result[i].Actions = lo.Slice(activity.Actions, 0, query.ActionLimit)
	})

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Timestamp > result[j].Timestamp
	})

	return result, nil
}

// saveIndexesPartitioned saves indexes in partitioned tables.
func (c *client) saveIndexesPartitioned(ctx context.Context, activities []*activityx.Activity) error {
	indexes := make(table.Indexes, 0)

	if err := indexes.Import(activities); err != nil {
		return err
	}

	if len(indexes) == 0 {
		return nil
	}

	// #nosec
	if err := c.createPartitionTable(ctx, indexes[0].PartitionName(), indexes[0].TableName()); err != nil {
		return fmt.Errorf("create partition table: %w", err)
	}

	// Delete indexes with the same activity id and network.
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

	errorPool := pool.New().WithContext(ctx).WithMaxGoroutines(10).WithCancelOnError().WithFirstError()

	for _, condition := range lo.Chunk(conditions, math.MaxUint8) {
		condition := condition

		errorPool.Go(func(ctx context.Context) error {
			return c.database.WithContext(ctx).
				Table(indexes[0].PartitionName()).
				Where("(id, network) IN (?)", condition).
				Delete(&table.Index{}).
				Error
		})
	}

	if err := errorPool.Wait(); err != nil {
		return fmt.Errorf("failed to delete indexes: %w", err)
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

	errorPool = pool.New().WithContext(ctx).WithMaxGoroutines(10).WithCancelOnError().WithFirstError()

	for _, index := range lo.Chunk(indexes, math.MaxUint8) {
		index := index

		errorPool.Go(func(ctx context.Context) error {
			return c.database.WithContext(ctx).
				Table(indexes[0].PartitionName()).
				Clauses(onConflict).
				Create(index).
				Error
		})
	}

	if err := errorPool.Wait(); err != nil {
		return fmt.Errorf("failed to save indexes: %w", err)
	}

	return nil
}

// findIndexPartitioned finds an activity  by id.
func (c *client) findIndexPartitioned(ctx context.Context, query model.ActivityQuery) (*table.Index, error) {
	index := table.Index{
		Timestamp: time.Now(),
	}

	tables := c.findIndexesPartitionTables(ctx, index)

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

			if data != nil && lo.IsNotEmpty(data.ID) {
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
//
//nolint:gocognit
func (c *client) findIndexesPartitioned(ctx context.Context, query model.ActivitiesQuery) ([]*table.Index, error) {
	index := table.Index{
		Timestamp: time.Now(),
	}

	if query.EndTimestamp != nil && *query.EndTimestamp > 0 && *query.EndTimestamp < uint64(index.Timestamp.Unix()) {
		index.Timestamp = time.Unix(int64(lo.FromPtr(query.EndTimestamp)), 0)
	}

	if query.Cursor != nil && query.Cursor.Timestamp < uint64(index.Timestamp.Unix()) {
		index.Timestamp = time.Unix(int64(query.Cursor.Timestamp), 0)
	}

	partitionedNames := c.findIndexesPartitionTables(ctx, index)

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

// deleteExpiredActivitiesPartitioned deletes expired activities.
func (c *client) deleteExpiredActivitiesPartitioned(ctx context.Context, network network.Network, timestamp time.Time) error {
	var (
		batchSize            = 1000
		dropActivitiesTables = make([]string, 0)
		checkTablesTimestamp = []time.Time{timestamp}
	)

	for i := 1; i <= 4; i++ {
		dropActivitiesTables = append(dropActivitiesTables, c.buildActivitiesTableNames(network, timestamp.AddDate(0, -3*i, 0)))
		checkTablesTimestamp = append(checkTablesTimestamp, timestamp.AddDate(0, -3*i, 0))
	}

	// Drop expired activities tables.
	for _, name := range dropActivitiesTables {
		zap.L().Info("dropping table", zap.String("table", name))

		if err := c.database.WithContext(ctx).Exec(fmt.Sprintf(`DROP TABLE IF EXISTS "%s"`, name)).Error; err != nil {
			zap.L().Error("failed to drop table", zap.Error(err), zap.String("table", name))

			return fmt.Errorf("drop table: %w", err)
		}
	}

	for _, checkTimestamp := range checkTablesTimestamp {
		activityTable := c.buildActivitiesTableNames(network, checkTimestamp)

		indexTable := c.buildIndexesTableNames(checkTimestamp)

		activityTableExists, err := c.findPartitionTableExists(ctx, activityTable)
		if err != nil {
			return fmt.Errorf("find partition table exists: %w", err)
		}

		indexTableExists, err := c.findPartitionTableExists(ctx, indexTable)
		if err != nil {
			return fmt.Errorf("find partition table exists: %w", err)
		}

		if !indexTableExists {
			continue
		}

		zap.L().Info("deleting expired activities", zap.String("table", activityTable), zap.String("indexTable", indexTable))

		for {
			done, err := c.batchDeleteExpiredActivities(ctx, network, timestamp, batchSize, &indexTable, lo.Ternary(activityTableExists, &activityTable, nil))
			if err != nil {
				return fmt.Errorf("batch delete expired activities: %w", err)
			}

			if done {
				break
			}
		}
	}

	return nil
}

func (c *client) batchDeleteExpiredActivities(ctx context.Context, network network.Network, timestamp time.Time, batchSize int, indexTable *string, activityTable *string) (bool, error) {
	databaseTransaction := c.database.WithContext(ctx).Begin()
	defer func() {
		_ = databaseTransaction.Rollback().Error
	}()

	var transactionIDs []string

	// Find expired activities.
	if err := c.database.WithContext(ctx).Table(*indexTable).Select("id").Where("network = ?", network.String()).
		Where("timestamp < ?", timestamp).Limit(batchSize).
		Pluck("id", &transactionIDs).Error; err != nil {
		zap.L().Error("failed to find expired activities", zap.Error(err), zap.String("table", *indexTable))

		return false, fmt.Errorf("find expired activities: %w", err)
	}

	if len(transactionIDs) == 0 {
		return true, nil
	}

	// Delete expired indexes.
	if err := databaseTransaction.Table(*indexTable).Where("id IN ?", transactionIDs).Delete(&table.Index{}).Error; err != nil {
		zap.L().Error("failed to delete expired indexes", zap.Error(err), zap.String("table", *indexTable))

		return false, fmt.Errorf("delete expired indexes: %w", err)
	}

	if activityTable != nil {
		// Delete expired activities.
		if err := databaseTransaction.Table(*activityTable).Where("id IN ?", transactionIDs).Delete(&table.Activity{}).Error; err != nil {
			zap.L().Error("failed to delete expired activities", zap.Error(err), zap.String("table", *activityTable))

			return false, fmt.Errorf("delete expired activities: %w", err)
		}
	}

	if err := databaseTransaction.Commit().Error; err != nil {
		zap.L().Error("failed to commit transaction", zap.Error(err))

		return false, fmt.Errorf("commit transaction: %w", err)
	}

	return false, nil
}

// buildFindIndexStatement builds the query index statement.
func (c *client) buildFindIndexStatement(ctx context.Context, partitionedName string, query model.ActivityQuery) *gorm.DB {
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
func (c *client) buildFindIndexesStatement(ctx context.Context, partition string, query model.ActivitiesQuery) *gorm.DB {
	databaseStatement := c.database.WithContext(ctx).Table(partition)

	if query.Distinct != nil && lo.FromPtr(query.Distinct) {
		databaseStatement = databaseStatement.Select("DISTINCT (id) id, timestamp, index, network")
	}

	if query.Owner != nil {
		databaseStatement = databaseStatement.Where("owner = ?", query.Owner)
	}

	if len(query.Owners) > 0 {
		databaseStatement = databaseStatement.Where("owner IN ?", query.Owners)
	}

	if query.Status != nil {
		databaseStatement = databaseStatement.Where("status = ?", query.Status)
	}

	if query.Direction != nil {
		databaseStatement = databaseStatement.Where("direction = ?", query.Direction)
	}

	if query.StartTimestamp != nil && *query.StartTimestamp > 0 {
		databaseStatement = databaseStatement.Where("timestamp >= ?", time.Unix(int64(*query.StartTimestamp), 0))
	}

	if query.EndTimestamp != nil && *query.EndTimestamp > 0 {
		databaseStatement = databaseStatement.Where("timestamp <= ?", time.Unix(int64(*query.EndTimestamp), 0))
	}

	if query.Platform != "" {
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

	return databaseStatement.Order("timestamp DESC, index DESC").Limit(query.Limit)
}

// buildActivitiesTableNames builds the activities table names.
func (c *client) buildActivitiesTableNames(network network.Network, timestamp time.Time) string {
	return fmt.Sprintf("%s_%s_%d_q%d", (*table.Activity).TableName(nil), network, timestamp.Year(), int(timestamp.Month()-1)/3+1)
}

// buildIndexesTableNames builds the indexes table names.
func (c *client) buildIndexesTableNames(timestamp time.Time) string {
	return fmt.Sprintf("%s_%d_q%d", (*table.Index).TableName(nil), timestamp.Year(), int(timestamp.Month()-1)/3+1)
}
