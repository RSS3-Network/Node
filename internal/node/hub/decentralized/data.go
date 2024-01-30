package decentralized

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/samber/lo"
)

func (h *Hub) getFeed(ctx context.Context, request model.FeedQuery) (*schema.Feed, *int, error) {
	return h.databaseClient.FindFeed(ctx, request)
}

func (h *Hub) getFeeds(ctx context.Context, request model.FeedsQuery) ([]*schema.Feed, string, error) {
	feeds, err := h.databaseClient.FindFeeds(ctx, request)
	if err != nil {
		return nil, "", fmt.Errorf("failed to find feeds: %w", err)
	}

	last, err := lo.Last(feeds)
	if err == nil {
		return feeds, h.transformCursor(ctx, last), nil
	}

	return nil, "", nil
}

func (h *Hub) getCursor(ctx context.Context, cursor *string) (*schema.Feed, error) {
	if cursor == nil {
		return nil, nil
	}

	str := strings.Split(*cursor, ":")
	if len(str) != 2 {
		return nil, fmt.Errorf("invalid cursor")
	}

	network, err := filter.NetworkString(str[1])
	if err != nil {
		return nil, fmt.Errorf("invalid cursor: %w", err)
	}

	data, _, err := h.getFeed(ctx, model.FeedQuery{ID: lo.ToPtr(str[0]), Network: lo.ToPtr(network)})
	if err != nil {
		return nil, fmt.Errorf("failed to get cursor: %w", err)
	}

	return data, nil
}

func (h *Hub) transformCursor(_ context.Context, feed *schema.Feed) string {
	if feed == nil {
		return ""
	}

	return fmt.Sprintf("%s:%s", feed.ID, feed.Network)
}

func (h *Hub) getIndexCount(ctx context.Context) (int64, *time.Time, error) {
	var (
		count      int64
		updateTime *time.Time
	)

	checkpoints, err := h.databaseClient.LoadCheckpoints(ctx, "", filter.NetworkUnknown, "")
	if err != nil {
		return count, nil, fmt.Errorf("failed to find index count: %w", err)
	}

	for _, checkpoint := range checkpoints {
		count += checkpoint.IndexCount

		if updateTime == nil || checkpoint.UpdatedAt.After(*updateTime) {
			updateTime = lo.ToPtr(checkpoint.UpdatedAt)
		}
	}

	return count, updateTime, nil
}
