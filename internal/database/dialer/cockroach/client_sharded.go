package cockroach

import (
	"context"
	"fmt"
	"math"

	"github.com/naturalselectionlabs/rss3-node/internal/database/table"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/samber/lo"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm/clause"
)

// saveFeedsSharded saves feeds to the database in sharded tables.
func (c *client) saveFeedsSharded(ctx context.Context, feeds []schema.Feed) error {
	shards := make(map[string][]schema.Feed)

	// Group feeds by sharded name.
	for _, feed := range feeds {
		var feedSharded table.Feed

		shardedName := feedSharded.ShardedName(lo.ToPtr(feed))

		if _, exists := shards[shardedName]; !exists {
			shards[shardedName] = make([]schema.Feed, 0)
		}

		shards[shardedName] = append(shards[shardedName], feed)
	}

	errorGroup, ctx := errgroup.WithContext(ctx)

	// Save feeds and indexes in parallel.
	for shardedName, feeds := range shards {
		shardedName, feeds := shardedName, feeds

		errorGroup.Go(func() error {
			tableFeeds := make(table.Feeds, 0)

			if err := tableFeeds.Import(feeds); err != nil {
				return err
			}

			if err := c.createShardedTable(ctx, shardedName, tableFeeds[0].TableName()); err != nil {
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

			if err := c.database.WithContext(ctx).Table(shardedName).Clauses(onConflict).CreateInBatches(tableFeeds, math.MaxUint8).Error; err != nil {
				return err
			}

			return c.saveIndexesSharded(ctx, feeds)
		})
	}

	return errorGroup.Wait()
}

// saveIndexesSharded saves indexes to the database in a same sharded table.
func (c *client) saveIndexesSharded(ctx context.Context, feeds []schema.Feed) error {
	indexes := make(table.Indexes, 0)

	if err := indexes.Import(feeds); err != nil {
		return err
	}

	if err := c.createShardedTable(ctx, indexes[0].ShardedName(), indexes[0].TableName()); err != nil {
		return fmt.Errorf("create sharded table: %w", err)
	}

	// Delete indexes with the same feed id and chain.
	pkIndexes := make(map[string][]string)

	for _, index := range indexes {
		if _, ok := pkIndexes[index.ID+index.Chain]; ok {
			continue
		}

		pkIndexes[index.ID+index.Chain] = []string{
			index.ID,
			index.Chain,
		}
	}

	err := c.database.WithContext(ctx).Table(indexes[0].ShardedName()).
		Where("(id, chain) IN (?)", lo.MapToSlice(pkIndexes, func(_ string, value []string) []string {
			return value
		})).Delete(&table.Indexes{}).Error
	if err != nil {
		return fmt.Errorf("delete indexes: %w", err)
	}

	// Save indexes.
	onConflict := clause.OnConflict{
		Columns: []clause.Column{
			{
				Name: "id",
			},
			{
				Name: "chain",
			},
			{
				Name: "owner",
			},
		},
		UpdateAll: true,
	}

	return c.database.WithContext(ctx).Table(indexes[0].ShardedName()).Clauses(onConflict).CreateInBatches(indexes, math.MaxUint8).Error
}
