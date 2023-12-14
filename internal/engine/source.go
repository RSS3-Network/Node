package engine

import (
	"context"
	"encoding/json"

	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

type Source interface {
	Chain() filter.Chain
	State() json.RawMessage
	Start(ctx context.Context, tasksChan chan<- []Task, errorChan chan<- error)
}

type SourceFilter interface{}
