package info

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
)

type NetworkConfigResponse struct {
	Data NetworkConfig `json:"data"`
}

type NetworkConfig struct {
	RSS           []NetworkConfigDetail `json:"rss,omitempty"`
	Decentralized []NetworkConfigDetail `json:"decentralized,omitempty"`
	Federated     []NetworkConfigDetail `json:"federated,omitempty"`
}

type NetworkConfigDetail struct {
	ID             string         `json:"id,omitempty"`
	EndpointConfig *Endpoint      `json:"endpoint_configs,omitempty"`
	WorkerConfig   []workerConfig `json:"worker_configs,omitempty"`
}

// GetNetworkConfig GetNetworksConfig returns the configuration for all supported networks.
func (c *Component) GetNetworkConfig(ctx echo.Context) error {
	go c.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, "config")

	config := NetworkConfig{
		RSS:           getNetworkConfigDetail(network.RSSSource),
		Decentralized: getNetworkConfigDetail(network.ArweaveSource, network.EthereumSource, network.FarcasterSource, network.NearSource),
		Federated:     getNetworkConfigDetail(network.ActivityPubSource),
	}

	return ctx.JSON(http.StatusOK, NetworkConfigResponse{
		Data: config,
	})
}

func getNetworkConfigDetail(sources ...network.Source) []NetworkConfigDetail {
	var details []NetworkConfigDetail

	for _, s := range sources {
		for _, n := range s.Networks() {
			if n == network.SatoshiVM {
				continue
			}

			networkDetail := NetworkConfigDetail{
				ID:           n.String(),
				WorkerConfig: []workerConfig{},
			}

			if s != network.RSSSource {
				endpointConfig := getEndpointConfig(s)
				networkDetail.EndpointConfig = &endpointConfig
			}

			for worker, config := range WorkerToConfigMap[s] {
				if lo.Contains(NetworkToWorkersMap[n], worker) {
					workerConfig := config
					workerConfig.ID.Value = n.String() + "-" + worker.Name()
					workerConfig.Network.Value = n.String()
					workerConfig.MinimumResource = calculateMinimumResources(n, worker)
					networkDetail.WorkerConfig = append(networkDetail.WorkerConfig, workerConfig)
				}
			}

			if len(networkDetail.WorkerConfig) > 0 {
				details = append(details, networkDetail)
			}
		}
	}

	return details
}
