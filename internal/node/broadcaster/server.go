package broadcaster

import (
	"context"
	"fmt"

	"github.com/naturalselectionlabs/rss3-global-indexer/client"
	"github.com/naturalselectionlabs/rss3-node/config"
	"github.com/robfig/cron/v3"
)

type Broadcaster struct {
	client *client.Client
	config *config.File
	cron   *cron.Cron
}

func (b *Broadcaster) Run(ctx context.Context) error {
	if err := b.Register(ctx); err != nil {
		return fmt.Errorf("register: %w", err)
	}

	_, err := b.cron.AddFunc("@every 1m", func() {
		if err := b.Heartbeat(ctx); err != nil {
			return
		}
	})
	if err != nil {
		return fmt.Errorf("add heartbeat cron job: %w", err)
	}

	b.cron.Start()

	return nil
}

func NewBroadcaster(_ context.Context, config *config.File) (*Broadcaster, error) {
	// Dial global indexer client.
	globalIndexerClient, err := client.NewClient(config.Discovery.Server.GlobalIndexerEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to dial global indexer client: %w", err)
	}

	return &Broadcaster{
		client: globalIndexerClient,
		config: config,
		cron:   cron.New(),
	}, nil
}
