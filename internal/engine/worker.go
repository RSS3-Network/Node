package engine

import (
	"context"

	"github.com/naturalselectionlabs/rss3-node/schema"
)

type Worker interface {
	Name() string
	Match(ctx context.Context, task Task) (bool, error)
	Transform(ctx context.Context, task Task) (*schema.Feed, error)
}
