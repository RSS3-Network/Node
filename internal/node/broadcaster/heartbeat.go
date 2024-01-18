package broadcaster

import (
	"context"
	"fmt"
	"time"

	"github.com/naturalselectionlabs/rss3-global-indexer/client"
	"go.uber.org/zap"
)

func (b *Broadcaster) Heartbeat(ctx context.Context) error {
	req := client.HeartbeatRequest{
		Address:   b.config.Discovery.Maintainer.EvmAddress,
		Signature: b.config.Discovery.Maintainer.Signature,
		Timestamp: time.Now().Unix(),
	}

	zap.L().Info("sending heartbeat", zap.Any("request", req))

	if err := b.client.NodeHeartbeat(ctx, req); err != nil {
		zap.L().Error("failed to send heartbeat", zap.Error(err))

		return fmt.Errorf("send heartbeat: %w", err)
	}

	return nil
}
