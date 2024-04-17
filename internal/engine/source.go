package engine

import (
	"context"
	"encoding/json"

	"github.com/rss3-network/protocol-go/schema/network"
)

type Source interface {
	Network() network.Network
	State() json.RawMessage
	Start(ctx context.Context, tasksChan chan<- *Tasks, errorChan chan<- error)
}

type SourceFilter interface{}
