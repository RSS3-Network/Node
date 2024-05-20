package network

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

type WorkerRequest struct {
	Network string `param:"network" validate:"required"`
	Worker  string `param:"worker" validate:"required"`
}

type Response struct {
	Data []string `json:"data"`
}

type ListWorkerResponse struct {
	Data []worker.Worker `json:"data"`
}

type WorkerConfigResponse struct {
	Data workerConfig `json:"data"`
}

type EndpointConfigResponse struct {
	Data Endpoint `json:"data"`
}

// GetNetworksHandler returns the list of all networks that the node can support.
func (c *Component) GetNetworksHandler(ctx echo.Context) error {
	go c.CollectMetric(ctx.Request().Context(), "networks")

	networkList := network.NetworkValues()

	result := make([]string, 0)

	for _, item := range networkList {
		networkStr := item.String()
		// skip unknown and bitcoin network
		if networkStr == network.Unknown.String() || networkStr == network.Bitcoin.String() {
			continue
		}

		result = append(result, networkStr)
	}

	return ctx.JSON(http.StatusOK, Response{
		Data: result,
	})
}

// GetWorkersByNetwork returns list of all workers for the specified network.
func (c *Component) GetWorkersByNetwork(ctx echo.Context) error {
	var request Request

	if err := ctx.Bind(&request); err != nil {
		return fmt.Errorf("bind failed: %w", err)
	}

	if err := ctx.Validate(&request); err != nil {
		return fmt.Errorf("validate failed: %w", err)
	}

	go c.CollectMetric(ctx.Request().Context(), request.Network)

	nid, err := network.NetworkString(request.Network)

	if err != nil {
		return fmt.Errorf("network string failed: %w", err)
	}

	return ctx.JSON(http.StatusOK, ListWorkerResponse{
		Data: NetworkToWorkersMap[nid],
	})
}

// GetWorkerConfig returns the recommended configuration for the specified worker.
func (c *Component) GetWorkerConfig(ctx echo.Context) error {
	var request WorkerRequest

	if err := ctx.Bind(&request); err != nil {
		return fmt.Errorf("bind failed: %w", err)
	}

	if err := ctx.Validate(&request); err != nil {
		return fmt.Errorf("validate failed: %w", err)
	}

	go c.CollectMetric(ctx.Request().Context(), fmt.Sprintf("%s-%s", request.Network, request.Worker))

	wid, err := worker.WorkerString(request.Worker)
	if err != nil {
		return fmt.Errorf("worker string failed: %w", err)
	}

	nid, err := network.NetworkString(request.Network)
	if err != nil {
		return fmt.Errorf("network string failed: %w", err)
	}

	// set default values for specific network architecture worker
	config := WorkerToConfigMap[nid.Source()][wid]
	config.Network.Value = nid

	// calculate the minimum resources required for the worker
	config.MinimumResource = calculateMinimumResources(nid, wid)

	return ctx.JSON(http.StatusOK, WorkerConfigResponse{
		Data: config,
	})
}

// GetEndpointConfig returns the endpoint configuration for the node.
func (c *Component) GetEndpointConfig(ctx echo.Context) error {
	config := getEndpointConfig()

	return ctx.JSON(http.StatusOK, EndpointConfigResponse{
		Data: config,
	})
}
