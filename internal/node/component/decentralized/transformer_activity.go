package decentralized

import (
	"context"
	"fmt"

	"github.com/rss3-network/node/schema/worker/decentralized"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"go.uber.org/zap"
)

// TransformActivity should add related URLs to the activity based on action tag, network and platform
func (c *Component) TransformActivity(ctx context.Context, activity *activityx.Activity) (*activityx.Activity, error) {
	if activity == nil {
		return nil, nil
	}

	// iterate over actions and transform them based on tag, network and platform
	lop.ForEach(activity.Actions, func(actionPtr *activityx.Action, index int) {
		action := *actionPtr

		var err error

		switch action.Tag {
		case tag.Collectible:
			*activity.Actions[index], err = c.TransformCollectibleType(ctx, activity.Network, *actionPtr)
		case tag.Social:
			*activity.Actions[index], err = c.TransformSocialType(ctx, activity.Network, activity.Platform, *actionPtr)
		default:
			activity.Actions[index] = actionPtr
		}

		if err != nil {
			zap.L().Error("failed to transform action", zap.Error(err), zap.String("id", activity.ID))
		}

		activity.Actions[index].RelatedURLs = append(activity.Actions[index].RelatedURLs, c.AddTransactionChainURL(ctx, activity.Network, activity.Platform, activity.ID)...)
	})

	return activity, nil
}

func (c *Component) AddTransactionChainURL(_ context.Context, n network.Network, platform string, id string) []string {
	urls := make([]string, 0)

	switch n {
	case network.Ethereum:
		urls = append(urls, fmt.Sprintf("https://etherscan.io/tx/%s", id))
	case network.Polygon:
		urls = append(urls, fmt.Sprintf("https://polygonscan.com/tx/%s", id))
	case network.BinanceSmartChain:
		urls = append(urls, fmt.Sprintf("https://bscscan.com/tx/%s", id))
	case network.Gnosis:
		urls = append(urls, fmt.Sprintf("https://blockscout.com/xdai/mainnet/tx/%s", id))
	case network.Crossbell:
		urls = append(urls, fmt.Sprintf("https://scan.crossbell.io/tx/%s", id))
	case network.Arbitrum:
		urls = append(urls, fmt.Sprintf("https://arbiscan.io/tx/%s", id))
	case network.Optimism:
		urls = append(urls, fmt.Sprintf("https://optimistic.etherscan.io/tx/%s", id))
	case network.Avalanche:
		urls = append(urls, fmt.Sprintf("https://snowtrace.io/tx/%s", id))
	case network.Base:
		urls = append(urls, fmt.Sprintf("https://basescan.org/tx/%s", id))
	case network.Arweave:
		urls = append(urls, fmt.Sprintf("https://arweave.app/tx/%s", id))

		if lo.IsNotEmpty(platform) && platform == decentralized.PlatformLens.String() {
			urls = append(urls, fmt.Sprintf("https://momoka.lens.xyz/tx/%s", id))
		}
	case network.Linea:
		urls = append(urls, fmt.Sprintf("https://lineascan.build/tx/%s", id))
	case network.SatoshiVM:
		urls = append(urls, fmt.Sprintf("https://svmscan.io/tx/%s", id))
	case network.VSL:
		urls = append(urls, fmt.Sprintf("https://scan.rss3.io/tx/%s", id))

	default:
		return nil
	}

	return urls
}
