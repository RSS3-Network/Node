package engine

import (
	"context"

	"github.com/naturalselectionlabs/rss3-node/schema"
)

type Name string

const (
	Fallback Name = "fallback"
)

type Worker interface {
	Match(ctx context.Context, task Task) (bool, error)
	Transform(ctx context.Context, task Task) (*schema.Feed, error)
}
