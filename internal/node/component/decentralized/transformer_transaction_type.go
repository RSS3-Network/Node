package decentralized

import (
	"context"
	"fmt"

	"github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"go.uber.org/zap"
)

func (c *Component) TransformTransactionType(ctx context.Context, network network.Network, action activity.Action) (activity.Action, error) {
	switch action.Type {
	case typex.TransactionTransfer, typex.TransactionMint, typex.TransactionBurn:
		return c.transformTransactionTransfer(ctx, network, action)
	case typex.TransactionApproval:
		return c.transformTransactionApproval(ctx, network, action)
	case typex.TransactionBridge:
		return c.transformTransactionBridge(ctx, network, action)
	}

	return action, nil
}

func (c *Component) transformTransactionTransfer(ctx context.Context, network network.Network, action activity.Action) (activity.Action, error) {
	transfer, ok := action.Metadata.(*metadata.TransactionTransfer)
	if !ok {
		zap.L().Error("invalid metadata type", zap.Any("metadata", action.Metadata))

		return activity.Action{}, fmt.Errorf("invalid metadata type: %T", action.Metadata)
	}

	action.Metadata = metadata.TransactionTransfer(c.buildTransactionTransferMetadata(ctx, network, metadata.Token(*transfer)))

	return action, nil
}

func (c *Component) transformTransactionApproval(ctx context.Context, network network.Network, action activity.Action) (activity.Action, error) {
	approval, ok := action.Metadata.(*metadata.TransactionApproval)
	if !ok {
		zap.L().Error("invalid metadata type", zap.Any("metadata", action.Metadata))

		return activity.Action{}, fmt.Errorf("invalid metadata type: %T", action.Metadata)
	}

	approval.Token = c.buildTransactionTransferMetadata(ctx, network, approval.Token)

	action.Metadata = approval

	return action, nil
}

func (c *Component) transformTransactionBridge(ctx context.Context, network network.Network, action activity.Action) (activity.Action, error) {
	bridge, ok := action.Metadata.(*metadata.TransactionBridge)
	if !ok {
		zap.L().Error("invalid metadata type", zap.Any("metadata", action.Metadata))

		return activity.Action{}, fmt.Errorf("invalid metadata type: %T", action.Metadata)
	}

	bridge.Token = c.buildTransactionTransferMetadata(ctx, network, bridge.Token)

	action.Metadata = bridge

	return action, nil
}

func (c *Component) buildTransactionTransferMetadata(_ context.Context, _ network.Network, token metadata.Token) metadata.Token {
	// TODO should we implement a function to get token logo image url as related url in api?
	return token
}
