package engine

import (
	"context"
	"encoding/json"

	"github.com/rss3-network/protocol-go/schema/network"
)

// DataSource is the interface that wraps the basic methods of a data source.
type DataSource interface {
	Network() network.Network
	State() json.RawMessage
	Start(ctx context.Context, tasksChan chan<- *Tasks, errorChan chan<- error)
}

type DataSourceFilter interface{}
