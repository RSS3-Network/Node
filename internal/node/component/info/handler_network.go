package info

import (
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/schema/worker/rss"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
)

type NetworkConfigResponse struct {
	Data NetworkConfig `json:"data"`
}

type NetworkConfig struct {
	RSS           NetworkConfigDetailForRSS `json:"rss,omitempty"`
	Decentralized []NetworkConfigDetail     `json:"decentralized,omitempty"`
	Federated     []NetworkConfigDetail     `json:"federated,omitempty"`
}

type NetworkConfigDetail struct {
	ID             string         `json:"id,omitempty"`
	EndpointConfig *Endpoint      `json:"endpoint_configs,omitempty"`
	WorkerConfig   []workerConfig `json:"worker_configs,omitempty"`
}

type NetworkConfigDetailForRSS struct {
	ID             string       `json:"id,omitempty"`
	EndpointConfig *Endpoint    `json:"endpoint_configs,omitempty"`
	WorkerConfig   workerConfig `json:"worker_configs,omitempty"`
}

// GetNetworkConfig GetNetworksConfig returns the configuration for all supported networks.
func (c *Component) GetNetworkConfig(ctx echo.Context) error {
	go c.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, "config")

	config := NetworkConfig{
		RSS:           getNetworkConfigDetailForRSS(network.RSSSource),
		Decentralized: getNetworkConfigDetail(network.ArweaveSource, network.EthereumSource, network.FarcasterSource, network.NearSource),
		Federated:     getNetworkConfigDetail(network.ActivityPubSource),
	}

	return ctx.JSON(http.StatusOK, NetworkConfigResponse{
		Data: config,
	})
}

func getNetworkConfigDetailForRSS(source network.Source) NetworkConfigDetailForRSS {
	n := source.Networks()[0]

	networkDetail := NetworkConfigDetailForRSS{
		ID: n.String(),
	}

	worker := rss.RSSHub
	config := WorkerToConfigMap[source][worker]

	workerConfig := config
	workerConfig.ID.Value = n.String() + "-" + worker.Name()
	workerConfig.Network.Value = n.String()
	workerConfig.MinimumResource = calculateMinimumResources(n, worker)
	networkDetail.WorkerConfig = workerConfig

	return networkDetail
}

func getNetworkConfigDetail(sources ...network.Source) []NetworkConfigDetail {
	var details []NetworkConfigDetail

	for _, s := range sources {
		for _, n := range s.Networks() {
			if n == network.Unknown || n == network.SatoshiVM || n == network.Bitcoin {
				continue
			}

			networkDetail := NetworkConfigDetail{
				ID:           n.String(),
				WorkerConfig: []workerConfig{},
			}

			if s != network.RSSSource {
				endpointConfig := getEndpointConfig(n)
				networkDetail.EndpointConfig = &endpointConfig
			}

			var workerConfigs []workerConfig

			for worker, config := range WorkerToConfigMap[s] {
				if lo.Contains(NetworkToWorkersMap[n], worker) {
					workerConfig := config
					workerConfig.ID.Value = n.String() + "-" + worker.Name()
					workerConfig.Network.Value = n.String()
					workerConfig.MinimumResource = calculateMinimumResources(n, worker)

					// Change endpoint type to string for Mastodon (ActivityPub)
					if s == network.ActivityPubSource {
						workerConfig.EndpointID.Type = StringType
					}

					workerConfigs = append(workerConfigs, workerConfig)
				}
			}

			// Sort the worker configs
			sort.Slice(workerConfigs, func(i, j int) bool {
				if workerConfigs[i].Worker.Value == "core" {
					return true
				}
				if workerConfigs[j].Worker.Value == "core" {
					return false
				}
				return workerConfigs[i].Worker.Value.(string) < workerConfigs[j].Worker.Value.(string)
			})

			networkDetail.WorkerConfig = workerConfigs

			if len(networkDetail.WorkerConfig) > 0 {
				details = append(details, networkDetail)
			}
		}
	}

	return details
}
