package engine

import (
	"context"

	"github.com/naturalselectionlabs/rss3-node/schema"
)

// Name represents a worker name.
//
//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Name --linecomment --output worker_string.go --json --sql
type Name int

const (
	Fallback Name = iota + 1 // fallback
)

type Worker interface {
	Name() string
	Match(ctx context.Context, task Task) (bool, error)
	Transform(ctx context.Context, task Task) (*schema.Feed, error)
}
