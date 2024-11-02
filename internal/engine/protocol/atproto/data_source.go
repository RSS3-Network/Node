package atproto

import (
	"context"
	"encoding/json"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
)

var _ engine.DataSource = (*dataSource)(nil)

type dataSource struct {
	config         *config.Module
	state          State
	databaseClient database.Client
}

func (s *dataSource) Network() network.Network {
	return s.config.Network
}

func (s *dataSource) State() json.RawMessage {
	return lo.Must(json.Marshal(s.state))
}

func (s *dataSource) Start(ctx context.Context, tasksChan chan<- *engine.Tasks, errorChan chan<- error) {
	//TODO implement me
	panic("implement me")
}

func (s *dataSource) initialize() error {
	// TODO implement me
	panic("implement me")
}
