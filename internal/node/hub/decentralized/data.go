package decentralized

import (
	"context"
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/database/model"
	"github.com/naturalselectionlabs/rss3-node/schema"
)

func (h *Hub) getFeed(ctx context.Context, request model.FeedQuery) (*schema.Feed, *int, error) {
	return h.databaseClient.FirstFeed(ctx, request)
}

func (h *Hub) getCursor(ctx context.Context, cursor *string) (*schema.Feed, error) {
	if cursor == nil {
		return nil, nil
	}

	// TODO resolve network

	data, _, err := h.getFeed(ctx, model.FeedQuery{ID: cursor})
	if err != nil {
		return nil, fmt.Errorf("invalid cursor: %w", err)
	}

	return data, nil
}

func (h *Hub) getFeeds(ctx context.Context, request model.FeedsQuery) ([]*schema.Feed, error) {
	return h.databaseClient.FindFeeds(ctx, request)
}
