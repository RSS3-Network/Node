package engine

import (
	"context"

	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

type Source interface {
	Chain() filter.Chain
	Start(ctx context.Context, tasksChan chan<- []Task, errorChan chan<- error)
}
