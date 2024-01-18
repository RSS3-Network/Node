package broadcaster

import (
	"context"
	"fmt"

	"github.com/naturalselectionlabs/rss3-global-indexer/client"
	"go.uber.org/zap"
)

func (b *Broadcaster) Register(ctx context.Context) error {
	req := client.RegisterNodeRequest{
		Address:   b.config.Discovery.Maintainer.EvmAddress,
		Signature: b.config.Discovery.Maintainer.Signature,
		Endpoint:  b.config.Discovery.Server.Endpoint,
	}

	zap.L().Info("registering node", zap.Any("request", req))

	if err := b.client.RegisterNode(ctx, req); err != nil {
		zap.L().Error("failed to register node", zap.Error(err))

		return fmt.Errorf("register node: %w", err)
	}

	return nil
}
