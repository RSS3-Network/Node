package cockroachdb

import (
	"context"
	"fmt"
	"math"

	"github.com/naturalselectionlabs/rss3-node/internal/database/dialer/cockroachdb/table"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/samber/lo"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm/clause"
)

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

// saveIndexesPartitioned saves indexes in partitioned tables.
func (c *client) saveIndexesPartitioned(ctx context.Context, feeds []*schema.Feed) error {
	indexes := make(table.Indexes, 0)

	if err := indexes.Import(feeds); err != nil {
		return err
	}

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

	return c.database.WithContext(ctx).
		Table(indexes[0].PartitionName()).
		Clauses(onConflict).
		CreateInBatches(indexes, math.MaxUint8).
		Error
}
