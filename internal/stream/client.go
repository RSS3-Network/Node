package stream

import (
	"context"

	"github.com/rss3-network/protocol-go/schema/activity"
)

type Client interface {
	PushFeeds(ctx context.Context, feeds []*activity.Activity) error
}
