package stream

import (
	"context"

	"github.com/rss3-network/protocol-go/schema/activity"
)

type Client interface {
	PushActivities(ctx context.Context, activities []*activity.Activity) error
}
