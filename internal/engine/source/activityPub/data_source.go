package activitypub

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/activityPub/mastodon"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
)

var _ engine.DataSource = (*dataSource)(nil)

type dataSource struct {
	config         *config.Module
	mastodonClient mastodon.Client
	databaseClient database.Client
	state          State
}

func (s *dataSource) Network() network.Network {
	return s.config.Network
}

func (s *dataSource) State() json.RawMessage {
	return lo.Must(json.Marshal(s.state))
}

func (s *dataSource) Start(_ context.Context, _ chan<- *engine.Tasks, errorChan chan<- error) {
	if err := s.initialize(); err != nil {
		errorChan <- fmt.Errorf("initialize dataSource: %w", err)
		return
	}
}

func (s *dataSource) initialize() (err error) {
	client, err := mastodon.NewClient(s.config.Endpoint.URL)
	if err != nil {
		return fmt.Errorf("create activityPub client: %w", err)
	}

	s.mastodonClient = client

	return nil
}

//  func retryOperation(ctx context.Context, operation func(ctx context.Context) error) error {
//	return retry.Do(
//		func() error {
//			return operation(ctx)
//		},
//		retry.Attempts(0),
//		retry.Delay(1*time.Second),
//		retry.DelayType(retry.BackOffDelay),
//		retry.OnRetry(func(n uint, err error) {
//			zap.L().Warn("retry farcaster dataSource", zap.Uint("retry", n), zap.Error(err))
//		}),
//	)
//}

func NewSource(config *config.Module, checkpoint *engine.Checkpoint, databaseClient database.Client) (engine.DataSource, error) {
	var (
		state State
	)

	// Initialize state from checkpoint.
	if checkpoint != nil {
		if err := json.Unmarshal(checkpoint.State, &state); err != nil {
			return nil, err
		}
	}

	instance := dataSource{
		databaseClient: databaseClient,
		config:         config,
		state:          state,
	}

	return &instance, nil
}
