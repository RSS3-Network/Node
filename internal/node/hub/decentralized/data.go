package decentralized

import (
	"context"
	"fmt"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/samber/lo"
	"strings"

	"github.com/naturalselectionlabs/rss3-node/internal/database/model"
	"github.com/naturalselectionlabs/rss3-node/schema"
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

func (h *Hub) transformCursor(ctx context.Context, feed *schema.Feed) string {
	if feed == nil {
		return ""
	}

	return fmt.Sprintf("%s:%s", feed.ID, feed.Network)
}
