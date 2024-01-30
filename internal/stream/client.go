package stream

import (
	"context"

	"github.com/rss3-network/protocol-go/schema"
)

type Client interface {
	PushFeeds(ctx context.Context, feeds []*schema.Feed) error
}
