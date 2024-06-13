package decentralized

import (
	"context"
	"fmt"

	"github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

// TransformCollectibleType adds collectible token URLs in opensea, looksrare and other platforms
func (c *Component) TransformCollectibleType(ctx context.Context, n network.Network, action activity.Action) (activity.Action, error) {
	switch action.Type {
	case typex.CollectibleTransfer, typex.CollectibleMint, typex.CollectibleBurn:
		return c.transformCollectibleTransfer(ctx, n, action)
	case typex.CollectibleApproval:
		return c.transformCollectibleApproval(ctx, n, action)
	case typex.CollectibleTrade:
		return c.transformCollectibleTrade(ctx, n, action)
	}

	return action, nil
}

// transformCollectibleTransfer adds URLs to collectible transfer action
func (c *Component) transformCollectibleTransfer(ctx context.Context, network network.Network, action activity.Action) (activity.Action, error) {
	transfer, ok := action.Metadata.(*metadata.CollectibleTransfer)
	if !ok {
		zap.L().Error("invalid metadata type", zap.Any("metadata", action.Metadata))

		return activity.Action{}, fmt.Errorf("invalid metadata type: %T", action.Metadata)
	}

	if transfer.Address != nil && transfer.ID != nil {
		action.RelatedURLs = append(action.RelatedURLs, c.buildCollectibleChainURL(ctx, network, lo.FromPtr(transfer.Address), transfer.ID.String())...)
	}

	return action, nil
}

// transformCollectibleApproval adds URLs to collectible approval action
func (c *Component) transformCollectibleApproval(ctx context.Context, network network.Network, action activity.Action) (activity.Action, error) {
	approval, ok := action.Metadata.(*metadata.CollectibleApproval)
	if !ok {
		zap.L().Error("invalid metadata type", zap.Any("metadata", action.Metadata))

		return activity.Action{}, fmt.Errorf("invalid metadata type: %T", action.Metadata)
	}

	if approval.Address != nil && approval.ID != nil {
		action.RelatedURLs = append(action.RelatedURLs, c.buildCollectibleChainURL(ctx, network, lo.FromPtr(approval.Address), approval.ID.String())...)
	}

	return action, nil
}

// transformCollectibleTrade adds URLs to collectible trade action
func (c *Component) transformCollectibleTrade(ctx context.Context, network network.Network, action activity.Action) (activity.Action, error) {
	trade, ok := action.Metadata.(*metadata.CollectibleTrade)
	if !ok {
		zap.L().Error("invalid metadata type", zap.Any("metadata", action.Metadata))

		return activity.Action{}, fmt.Errorf("invalid metadata type: %T", action.Metadata)
	}

	if trade.Address != nil && trade.ID != nil {
		action.RelatedURLs = append(action.RelatedURLs, c.buildCollectibleChainURL(ctx, network, lo.FromPtr(trade.Address), trade.ID.String())...)
	}

	return action, nil
}

// buildCollectibleChainURL builds the collectible chain URL based on network
func (c *Component) buildCollectibleChainURL(_ context.Context, n network.Network, address, id string) []string {
	var result []string

	switch n {
	case network.Ethereum:
		result = append(result, fmt.Sprintf("https://opensea.io/assets/%s/%s", address, id),
			fmt.Sprintf("https://looksrare.org/collections/%s/%s", address, id))
	case network.Polygon:
		result = append(result, fmt.Sprintf("https://opensea.io/assets/matic/%s/%s", address, id))
	case network.BinanceSmartChain:
		result = append(result, fmt.Sprintf("https://tofunft.com/nft/bsc/%s/%s", address, id))
	case network.Arbitrum:
		result = append(result, fmt.Sprintf("https://opensea.io/assets/arbitrum/%s/%s", address, id))
	case network.Optimism:
		result = append(result, fmt.Sprintf("https://qx.app/asset/%s/%s", address, id))
	case network.Avalanche:
		result = append(result, fmt.Sprintf("https://opensea.io/assets/avalanche/%s/%s", address, id))
	case network.Base:
		result = append(result, fmt.Sprintf("https://opensea.io/assets/base/%s/%s", address, id))
	}

	return result
}
