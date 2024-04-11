package engine

import (
	"time"

	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"go.opentelemetry.io/otel/attribute"
)

type Task interface {
	ID() string
	GetNetwork() filter.Network
	GetTimestamp() uint64
	Validate() error
	BuildFeed(options ...schema.FeedOption) (*schema.Feed, error)
}

func BuildTaskTraceAttributes(task Task) []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("task.id", task.ID()),
		attribute.Stringer("task.network", task.GetNetwork()),
		attribute.Stringer("task.timestamp", time.Unix(int64(task.GetTimestamp()), 0)),
	}
}
