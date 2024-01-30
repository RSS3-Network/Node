package engine

import (
	"context"

	"github.com/rss3-network/protocol-go/schema"
)

type Worker interface {
	Name() string
	Filter() SourceFilter
	Match(ctx context.Context, task Task) (bool, error)
	Transform(ctx context.Context, task Task) (*schema.Feed, error)
}
