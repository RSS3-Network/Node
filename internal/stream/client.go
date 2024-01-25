package stream

import (
	"context"

	"github.com/rss3-network/serving-node/schema"
)

type Client interface {
	PushFeeds(ctx context.Context, feeds []*schema.Feed) error
}
