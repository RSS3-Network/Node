package stream

import (
	"context"

	"github.com/naturalselectionlabs/rss3-node/schema"
)

type Client interface {
	PushFeeds(ctx context.Context, feeds []*schema.Feed) error
}
