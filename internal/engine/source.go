package engine

import (
	"context"
	"encoding/json"

	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

type Source interface {
	Network() filter.Network
	State() json.RawMessage
	Start(ctx context.Context, tasksChan chan<- []Task, errorChan chan<- error)
}
