package info

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/node/schema/worker/rss"
	"github.com/rss3-network/protocol-go/schema/network"
)

type AssetsResponse struct {
	Data struct {
		Networks map[string]AssetDetail `json:"networks"`
		Workers  map[string]AssetDetail `json:"workers"`
	} `json:"data"`
}

type AssetDetail struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Platform string `json:"platform,omitempty"`
	IconURL  string `json:"icon_url"`
}

// GetNodeAssets returns the network, workers, and assets info.
func (c *Component) GetNodeAssets(ctx echo.Context) error {
	response := AssetsResponse{}
	response.Data.Networks = make(map[string]AssetDetail)
	response.Data.Workers = make(map[string]AssetDetail)

	// Aggregate networks
	for _, n := range network.NetworkValues() {
		if n == network.Unknown || n == network.Bitcoin || n == network.Mastodon || n == network.RSS || n == network.SatoshiVM {
			continue
		}

		response.Data.Networks[n.String()] = AssetDetail{
			Type:    string(n.Source()),
			Name:    NetworkToNameMap[n],
			IconURL: NetworkToIconURLMap[n],
		}
	}

	// Aggregate decentralized workers
	for _, w := range decentralized.WorkerValues() {
		if w.Name() == decentralized.SAVM.Name() {
			continue
		}

		platform := decentralized.ToPlatformMap[w]
		response.Data.Workers[w.Name()] = AssetDetail{
			Name:     w.Name(),
			Platform: platform.String(),
			Type:     w.Component(),
			IconURL:  decentralized.ToIconURLMap[w],
		}
	}

	// Aggregate RSS workers
	for _, w := range rss.WorkerValues() {
		response.Data.Workers[w.Name()] = AssetDetail{
			Name:     w.Name(),
			Platform: "RSSHub",
			Type:     w.Component(),
			IconURL:  rss.ToIconURLMap[w],
		}
	}

	return ctx.JSON(http.StatusOK, response)
}

var NetworkToNameMap = map[network.Network]string{
	network.Arbitrum:          "Arbitrum",
	network.Arweave:           "Arweave",
	network.Avalanche:         "Avalanche",
	network.Base:              "Base",
	network.BinanceSmartChain: "Binance Smart Chain",
	network.Bitcoin:           "Bitcoin",
	network.Crossbell:         "Crossbell",
	network.Ethereum:          "Ethereum",
	network.Farcaster:         "Farcaster",
	network.Gnosis:            "Gnosis",
	network.Linea:             "Linea",
	network.Mastodon:          "Mastodon",
	network.Optimism:          "Optimism",
	network.Polygon:           "Polygon",
	network.SatoshiVM:         "SatoshiVM",
	network.VSL:               "VSL",
	network.XLayer:            "X Layer",
}

var NetworkToIconURLMap = map[network.Network]string{
	network.Arbitrum:          "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/arbitrum.svg",
	network.Arweave:           "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/arweave-mono.svg",
	network.Avalanche:         "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/avalanche.svg",
	network.Base:              "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/base.svg",
	network.BinanceSmartChain: "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/binance-smart-chain.svg",
	network.Crossbell:         "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/crossbell.svg",
	network.Ethereum:          "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/ethereum-mono.svg",
	network.Farcaster:         "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/farcaster.svg",
	network.Gnosis:            "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/gnosis-chain-mono.svg",
	network.Linea:             "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/linea-mono.svg",
	network.Optimism:          "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/optimism.svg",
	network.Polygon:           "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/polygon.svg",
	network.VSL:               "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/rss3-vsl.svg",
	network.XLayer:            "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/x-layer-mono.svg",
}
