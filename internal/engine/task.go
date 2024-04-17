package engine

import (
	"time"

	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
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

var _ propagation.TextMapCarrier = (*Tasks)(nil)

type Tasks struct {
	Tasks []Task

	// metadata is used to store OpenTelemetry trace context.
	metadata map[string]string
}

func (t *Tasks) Get(key string) string {
	return t.metadata[key]
}

func (t *Tasks) Set(key string, value string) {
	if t.metadata == nil {
		t.metadata = make(map[string]string)
	}

	t.metadata[key] = value
}

func (t *Tasks) Keys() []string {
	return lo.Keys(t.metadata)
}

func (t *Tasks) Len() int {
	return len(t.Tasks)
}
