package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/protocol-go/schema/network"
	"go.uber.org/zap"
)

func (c *client) UpdateRecentMastodonHandles(ctx context.Context, handles []*model.MastodonHandle, network network.Network) error {
	logger := zap.L()
	logger.Info("Starting UpdateRecentMastodonHandles", zap.Int("handleCount", len(handles)))

	// build the mastodon update handle table
	tableName := c.buildTableName(network)

	// Construct and execute the SQL query
	err := c.executeInsertQuery(ctx, handles, tableName)
	if err != nil {
		logger.Error("Error in executing insert query", zap.Error(err))
		return err
	}

	// Trim old handles
	if err := c.trimOldHandles(ctx, tableName); err != nil {
		logger.Error("Error trimming old handles", zap.Error(err))
		return err
	}

	logger.Info("Completed UpdateRecentMastodonHandles")

	return nil
}

// executeInsertQuery generates and executes the SQL insert query for Mastodon handles
func (c *client) executeInsertQuery(ctx context.Context, handles []*model.MastodonHandle, tableName string) error {
	valueStrings := make([]string, len(handles))
	valueArgs := make([]interface{}, 0, len(handles)*2)

	for i, h := range handles {
		valueStrings[i] = fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2)

		valueArgs = append(valueArgs, h.Handle, h.LastUpdated)
	}

	statement := fmt.Sprintf(`
		INSERT INTO %s (handle, last_updated)
		VALUES %s
		ON CONFLICT (handle)
		DO UPDATE SET last_updated = EXCLUDED.last_updated;`,
		tableName, strings.Join(valueStrings, ","),
	)

	// Execute the query
	result := c.database.WithContext(ctx).Exec(statement, valueArgs...)
	if result.Error != nil {
		return fmt.Errorf("failed to execute insert query: %w", result.Error)
	}

	return nil
}

// trimOldHandles removes outdated handles from the dataset
func (c *client) trimOldHandles(ctx context.Context, tableName string) error {
	statement := fmt.Sprintf(`
		DELETE FROM %s
		WHERE handle NOT IN (
			SELECT handle
			FROM %s
			ORDER BY last_updated DESC
			LIMIT $1
		);`, tableName, tableName)

	trimResult := c.database.WithContext(ctx).Exec(statement, maxRecentHandles)

	if trimResult.Error != nil {
		return fmt.Errorf("failed to trim old handles: %w", trimResult.Error)
	}

	return nil
}

func (c *client) GetUpdatedMastodonHandles(ctx context.Context, since uint64, limit int, cursor string) (*model.PaginatedMastodonHandles, error) {
	var result []model.MastodonHandle

	var totalCount int64

	//ToDo: fix the hardcode network value
	tableName := c.buildTableName(network.Mastodon)

	// Build the query
	query := c.database.WithContext(ctx).Table(tableName).Where("last_updated > ?", since)

	// Count total handles
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, fmt.Errorf("failed to count handles: %w", err)
	}

	// Apply cursor if provided
	if cursor != "" {
		var cursorLastUpdated uint64

		var cursorHandle string

		_, err := fmt.Sscanf(cursor, "%d:%s", &cursorLastUpdated, &cursorHandle)
		if err != nil {
			return nil, fmt.Errorf("invalid cursor format: %w", err)
		}

		query = query.Where("last_updated < ? OR (last_updated = ? AND handle > ?)", cursorLastUpdated, cursorLastUpdated, cursorHandle)
	}

	// Fetch handles with their last_updated timestamps
	if err := query.Select("handle, last_updated").Order("last_updated DESC, handle ASC").Limit(limit + 1).Find(&result).Error; err != nil {
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

// buildTableName generates a table name for storing updated handles based on the federated network
func (c *client) buildTableName(network network.Network) string {
	return fmt.Sprintf("dataset_%s_update_handles", network)
}
