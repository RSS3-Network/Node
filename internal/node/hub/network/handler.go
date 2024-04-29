package network

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/protocol-go/schema/network"
)

type Request struct {
	Network string `param:"network" validate:"required"`
}

type Response struct {
	Data []string `json:"data"`
}

// GetNetworksHandler get networks handler
func (h *Hub) GetNetworksHandler(c echo.Context) error {
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

	return c.JSON(http.StatusOK, Response{
		Data: result,
	})
}

// GetWorkersByNetwork returns workers by Network.
func (h *Hub) GetWorkersByNetwork(c echo.Context) error {
	var request Request

	if err := c.Bind(&request); err != nil {
		return fmt.Errorf("bind failed: %w", err)
	}

	if err := c.Validate(&request); err != nil {
		return fmt.Errorf("validate failed: %w", err)
	}

	go h.countRequest(context.TODO(), request.Network)

	nid, err := network.NetworkString(request.Network)

	if err != nil {
		return fmt.Errorf("network string failed: %w", err)
	}

	return c.JSON(http.StatusOK, Response{
		Data: NetworkToWorkersMap[nid],
	})
}

// NetworkToWorkersMap is a map of network to workers.
var NetworkToWorkersMap = map[network.Network][]string{
	network.Ethereum: {
		worker.Core.String(),
		worker.RSS3.String(),
		worker.OpenSea.String(),
		worker.Uniswap.String(),
		worker.Optimism.String(),
		worker.Looksrare.String(),
		worker.Highlight.String(),
		worker.Aave.String(),
		worker.Lido.String(),
		worker.ENS.String(),
		worker.Oneinch.String(),
		worker.Stargate.String(),
		worker.Curve.String(),
	},
	network.Arweave: {
		worker.Mirror.String(),
		worker.Paragraph.String(),
		worker.Momoka.String(),
	},
	network.Farcaster: {
		worker.Core.String(),
	},
	network.Polygon: {
		worker.Aavegotchi.String(),
		worker.Lens.String(),
		worker.Matters.String(),
		worker.Aave.String(),
		worker.IQWiki.String(),
		worker.Core.String(),
		worker.Highlight.String(),
		worker.Stargate.String(),
		worker.Curve.String(),
	},
	network.Crossbell: {
		worker.Crossbell.String(),
	},
	network.Avalanche: {
		worker.Aave.String(),
		worker.Stargate.String(),
		worker.Curve.String(),
	},
	network.Base: {
		worker.Aave.String(),
		worker.Core.String(),
		worker.Stargate.String(),
	},
	network.Optimism: {
		worker.Aave.String(),
		worker.Core.String(),
		worker.Highlight.String(),
		worker.KiwiStand.String(),
		worker.Stargate.String(),
		worker.Curve.String(),
	},
	network.Arbitrum: {
		worker.Aave.String(),
		worker.Core.String(),
		worker.Highlight.String(),
		worker.Stargate.String(),
		worker.Curve.String(),
	},
	network.VSL: {
		worker.Core.String(),
		worker.RSS3.String(),
		worker.VSL.String(),
	},
	network.SatoshiVM: {
		worker.Core.String(),
		worker.Uniswap.String(),
		worker.SAVM.String(),
	},
	network.BinanceSmartChain: {
		worker.Core.String(),
		worker.Stargate.String(),
	},
	network.Gnosis: {
		worker.Core.String(),
		worker.Curve.String(),
	},
	network.Linea: {
		worker.Core.String(),
		worker.Uniswap.String(),
		worker.Stargate.String(),
	},
}
