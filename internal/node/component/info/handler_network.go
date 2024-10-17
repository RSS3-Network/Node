package info

import (
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/schema/worker"
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
		RSS:           getNetworkConfigDetailForRSS(network.RSSProtocol),
		Decentralized: getNetworkConfigDetail(network.ArweaveProtocol, network.EthereumProtocol, network.FarcasterProtocol, network.NearProtocol),
		Federated:     getNetworkConfigDetail(network.ActivityPubProtocol),
	}

	return ctx.JSON(http.StatusOK, NetworkConfigResponse{
		Data: config,
	})
}

func getNetworkConfigDetailForRSS(protocol network.Protocol) NetworkConfigDetailForRSS {
	n := protocol.Networks()[0]

	networkDetail := NetworkConfigDetailForRSS{
		ID: n.String(),
	}

	worker := rss.Core
	config := WorkerToConfigMap[protocol][worker]

	workerConfig := config
	workerConfig.ID.Value = n.String() + "-" + worker.Name()
	workerConfig.Network.Value = n.String()
	workerConfig.MinimumResource = calculateMinimumResources(n, worker)

	if workerConfig.EndpointID != nil {
		workerConfig.EndpointID.Type = URLType
	}

	networkDetail.WorkerConfig = workerConfig

	return networkDetail
}

func getNetworkConfigDetail(protocols ...network.Protocol) []NetworkConfigDetail {
	var details []NetworkConfigDetail

	for _, protocol := range protocols {
		for _, n := range protocol.Networks() {
			if shouldSkipNetwork(n) {
				continue
			}

			networkDetail := createNetworkDetail(protocol, n)
			workerConfigs := getWorkerConfigs(protocol, n)
			sortWorkerConfigs(workerConfigs)

			networkDetail.WorkerConfig = workerConfigs

			if len(networkDetail.WorkerConfig) > 0 {
				details = append(details, networkDetail)
			}
		}
	}

	return details
}

func shouldSkipNetwork(n network.Network) bool {
	return n == network.Unknown || n == network.SatoshiVM || n == network.Bitcoin
}

func createNetworkDetail(protocol network.Protocol, n network.Network) NetworkConfigDetail {
	networkDetail := NetworkConfigDetail{
		ID:           n.String(),
		WorkerConfig: []workerConfig{},
	}

	if protocol != network.RSSProtocol {
		endpointConfig := getEndpointConfig(n)
		networkDetail.EndpointConfig = &endpointConfig
	}

	return networkDetail
}

func getWorkerConfigs(protocol network.Protocol, n network.Network) []workerConfig {
	var workerConfigs []workerConfig

	for w, config := range WorkerToConfigMap[protocol] {
		if lo.Contains(NetworkToWorkersMap[n], w) {
			workerConfig := createWorkerConfig(n, w, config)
			workerConfigs = append(workerConfigs, workerConfig)
		}
	}

	return workerConfigs
}

func createWorkerConfig(n network.Network, worker worker.Worker, config workerConfig) workerConfig {
	workerConfig := config
	workerConfig.ID.Value = n.String() + "-" + worker.Name()
	workerConfig.Network.Value = n.String()
	workerConfig.MinimumResource = calculateMinimumResources(n, worker)

	if workerConfig.EndpointID != nil {
		workerConfig.EndpointID.Type = StringType
		workerConfig.EndpointID.Value = n.String()
	}

	return workerConfig
}

func sortWorkerConfigs(workerConfigs []workerConfig) {
	sort.Slice(workerConfigs, func(i, j int) bool {
		if workerConfigs[i].Worker.Value == "core" {
			return true
		}

		if workerConfigs[j].Worker.Value == "core" {
			return false
		}

		return workerConfigs[i].Worker.Value.(string) < workerConfigs[j].Worker.Value.(string)
	})
}
