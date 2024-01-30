package engine

import (
	"context"
	"encoding/json"

	"github.com/rss3-network/protocol-go/schema/filter"
)

type Source interface {
	Network() filter.Network
	State() json.RawMessage
	Start(ctx context.Context, tasksChan chan<- []Task, errorChan chan<- error)
}

type SourceFilter interface{}
