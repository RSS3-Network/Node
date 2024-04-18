package stream

import (
	"context"

	activityx "github.com/rss3-network/protocol-go/schema/activity"
)

type Client interface {
	PushActivities(ctx context.Context, activities []*activityx.Activity) error
}
