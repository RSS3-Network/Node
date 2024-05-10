package decentralized

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
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
	go c.CollectMetric(ctx.Request().Context(), "networks")

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

	go c.CollectMetric(ctx.Request().Context(), common.HexToAddress(request.Network).String())

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
		worker.Aave,
		worker.Core,
		worker.Curve,
		worker.ENS,
		worker.Highlight,
		worker.Lido,
		worker.Looksrare,
		worker.Oneinch,
		worker.OpenSea,
		worker.Optimism,
		worker.RSS3,
		worker.Stargate,
		worker.Uniswap,
	},
	network.Arweave: {
		worker.Mirror,
		worker.Momoka,
		worker.Paragraph,
	},
	network.Farcaster: {
		worker.Core,
	},
	network.Polygon: {
		worker.Aave,
		worker.Aavegotchi,
		worker.Core,
		worker.Curve,
		worker.Highlight,
		worker.IQWiki,
		worker.Lens,
		worker.Matters,
		worker.Stargate,
	},
	network.Crossbell: {
		worker.Crossbell,
	},
	network.Avalanche: {
		worker.Aave,
		worker.Curve,
		worker.Stargate,
	},
	network.Base: {
		worker.Aave,
		worker.Core,
		worker.Stargate,
	},
	network.Optimism: {
		worker.Aave,
		worker.Core,
		worker.Curve,
		worker.Highlight,
		worker.KiwiStand,
		worker.Stargate,
	},
	network.Arbitrum: {
		worker.Aave,
		worker.Core,
		worker.Curve,
		worker.Highlight,
		worker.Stargate,
	},
	network.VSL: {
		worker.Core,
		worker.RSS3,
		worker.VSL,
	},
	network.SatoshiVM: {
		worker.Core,
		worker.SAVM,
		worker.Uniswap,
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
		worker.Stargate,
		worker.Uniswap,
	},
}
