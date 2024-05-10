package decentralized

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/protocol-go/schema/network"
)

type Request struct {
	Network string `param:"network" validate:"required"`
}

type NetworkResponse struct {
	Data []string `json:"data"`
}

type ListWorkerResponse struct {
	Data []worker.Worker `json:"data"`
}

// GetNetworksHandler get networks handler
func (c *Component) GetNetworksHandler(ctx echo.Context) error {
	networkList := network.NetworkValues()

	result := make([]string, 0)

	for _, item := range networkList {
		networkStr := item.String()
		// do not add unknown network
		if networkStr == "unknown" {
			continue
		}

		result = append(result, networkStr)
	}

	return ctx.JSON(http.StatusOK, NetworkResponse{
		Data: result,
	})
}

// GetWorkersByNetwork returns workers by Network.
func (c *Component) GetWorkersByNetwork(ctx echo.Context) error {
	var request Request

	if err := ctx.Bind(&request); err != nil {
		return fmt.Errorf("bind failed: %w", err)
	}

	if err := ctx.Validate(&request); err != nil {
		return fmt.Errorf("validate failed: %w", err)
	}

	nid, err := network.NetworkString(request.Network)

	if err != nil {
		return fmt.Errorf("network string failed: %w", err)
	}

	return ctx.JSON(http.StatusOK, ListWorkerResponse{
		Data: NetworkToWorkersMap[nid],
	})
}

// NetworkToWorkersMap is a map of network to workers.
var NetworkToWorkersMap = map[network.Network][]worker.Worker{
	network.Ethereum: {
		worker.Core,
		worker.RSS3,
		worker.OpenSea,
		worker.Uniswap,
		worker.Optimism,
		worker.Looksrare,
		worker.Highlight,
		worker.Aave,
		worker.Lido,
		worker.ENS,
		worker.Oneinch,
		worker.Stargate,
		worker.Curve,
	},
	network.Arweave: {
		worker.Mirror,
		worker.Paragraph,
		worker.Momoka,
	},
	network.Farcaster: {
		worker.Core,
	},
	network.Polygon: {
		worker.Aavegotchi,
		worker.Lens,
		worker.Matters,
		worker.Aave,
		worker.IQWiki,
		worker.Core,
		worker.Highlight,
		worker.Stargate,
		worker.Curve,
	},
	network.Crossbell: {
		worker.Crossbell,
	},
	network.Avalanche: {
		worker.Aave,
		worker.Stargate,
		worker.Curve,
	},
	network.Base: {
		worker.Aave,
		worker.Core,
		worker.Stargate,
	},
	network.Optimism: {
		worker.Aave,
		worker.Core,
		worker.Highlight,
		worker.KiwiStand,
		worker.Stargate,
		worker.Curve,
	},
	network.Arbitrum: {
		worker.Aave,
		worker.Core,
		worker.Highlight,
		worker.Stargate,
		worker.Curve,
	},
	network.VSL: {
		worker.Core,
		worker.RSS3,
		worker.VSL,
	},
	network.SatoshiVM: {
		worker.Core,
		worker.Uniswap,
		worker.SAVM,
	},
	network.BinanceSmartChain: {
		worker.Core,
		worker.Stargate,
	},
	network.Gnosis: {
		worker.Core,
		worker.Curve,
	},
	network.Linea: {
		worker.Core,
		worker.Uniswap,
		worker.Stargate,
	},
}
