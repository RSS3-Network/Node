package info

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/node/monitor"
	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/node/schema/worker/rss"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/samber/lo"
)

type WorkerResponse struct {
	Data ComponentInfo `json:"data"`
}

type ComponentInfo struct {
	Decentralized []*WorkerInfo `json:"decentralized"`
	RSS           *WorkerInfo   `json:"rss"`
	Federated     []*WorkerInfo `json:"federated"`
}

type WorkerInfo struct {
	WorkerID string                 `json:"worker_id"`
	Worker   worker.Worker          `json:"worker"`
	Network  network.Network        `json:"network"`
	Tags     []tag.Tag              `json:"tags"`
	Platform decentralized.Platform `json:"platform"`
	Status   worker.Status          `json:"status"`
	monitor.WorkerProgress
}

// GetWorkersStatus returns the status of all workers.
func (c *Component) GetWorkersStatus(ctx echo.Context) error {
	go c.CollectTrace(ctx.Request().Context(), ctx.Request().RequestURI, "status")

	workerCount := len(c.config.Component.Decentralized) + lo.Ternary(c.config.Component.RSS != nil, 1, 0) + len(c.config.Component.Federated)
	workerInfoChan := make(chan *WorkerInfo, workerCount)

	var response *WorkerResponse

	if c.redisClient != nil {
		// Fetch all worker info concurrently.
		c.fetchAllWorkerInfo(ctx, workerInfoChan)

		// Build the worker response.
		response = c.buildWorkerResponse(workerInfoChan)
	} else if c.config.Component.RSS != nil {
		m := c.config.Component.RSS

		response = &WorkerResponse{
			Data: ComponentInfo{
				RSS: &WorkerInfo{
					WorkerID: m.ID,
					Network:  m.Network,
					Worker:   m.Worker,
					Tags:     []tag.Tag{tag.RSS},
					Platform: decentralized.PlatformUnknown,
					Status:   worker.StatusReady},
			},
		}
	}

	return ctx.JSON(http.StatusOK, response)
}

// fetchAllWorkerInfo fetches the status of all workers concurrently.
func (c *Component) fetchAllWorkerInfo(ctx echo.Context, workerInfoChan chan<- *WorkerInfo) {
	var wg sync.WaitGroup

	fetchWorkerInfo := func(w *config.Module, fetchFunc func(context.Context, *config.Module) *WorkerInfo) {
		wg.Add(1)

		go func(module *config.Module) {
			defer wg.Done()

			workerInfoChan <- fetchFunc(ctx.Request().Context(), module)
		}(w)
	}

	modules := append(append(c.config.Component.Decentralized, c.config.Component.RSS), c.config.Component.Federated...)

	for _, m := range modules {
		fetchWorkerInfo(m, c.fetchWorkerInfo)
	}

	go func() {
		wg.Wait()
		close(workerInfoChan)
	}()
}

// buildWorkerResponse builds the worker response from the worker info channel
func (c *Component) buildWorkerResponse(workerInfoChan <-chan *WorkerInfo) *WorkerResponse {
	response := &WorkerResponse{
		Data: ComponentInfo{
			Decentralized: []*WorkerInfo{},
			RSS:           &WorkerInfo{},
			Federated:     []*WorkerInfo{},
		},
	}

	for workerInfo := range workerInfoChan {
		switch workerInfo.Network.Source() {
		case network.RSSSource:
			response.Data.RSS = workerInfo
		case network.EthereumSource, network.FarcasterSource, network.ArweaveSource:
			response.Data.Decentralized = append(response.Data.Decentralized, workerInfo)
		default:
			response.Data.Federated = append(response.Data.Federated, workerInfo)
		}
	}

	return response
}

// fetchWorkerInfo fetches the worker info with the different network source.
func (c *Component) fetchWorkerInfo(ctx context.Context, module *config.Module) *WorkerInfo {
	// Fetch status and progress from a specific worker by id.
	status, workerProgress := c.getWorkerStatusAndProgressByID(ctx, module.ID)

	workerInfo := &WorkerInfo{
		WorkerID: module.ID,
		Network:  module.Network,
		Worker:   module.Worker,
		Status:   status,
		WorkerProgress: monitor.WorkerProgress{
			RemoteState:  workerProgress.RemoteState,
			IndexedState: workerProgress.IndexedState,
			IndexCount:   workerProgress.IndexCount,
		},
	}

	switch module.Network.Source() {
	case network.EthereumSource, network.ArweaveSource, network.FarcasterSource:
		workerInfo.Platform = decentralized.ToPlatformMap[module.Worker.(decentralized.Worker)]
		workerInfo.Tags = decentralized.ToTagsMap[module.Worker.(decentralized.Worker)]

		// Handle special tags for decentralized core workers.
		if module.Worker == decentralized.Core {
			switch module.Network {
			case network.Farcaster:
				workerInfo.Tags = []tag.Tag{tag.Social}
			case network.Arweave:
				workerInfo.Tags = []tag.Tag{tag.Transaction}
			case network.VSL:
				workerInfo.Tags = append(workerInfo.Tags, tag.Exchange)
			default:
			}
		}
	case network.RSSSource:
		workerInfo.Tags = rss.ToTagsMap[module.Worker.(rss.Worker)]
	}

	return workerInfo
}

// getWorkerStatusAndProgressByID gets both worker status and progress from Redis cache by worker ID.
func (c *Component) getWorkerStatusAndProgressByID(ctx context.Context, workerID string) (worker.Status, monitor.WorkerProgress) {
	if c.redisClient == nil {
		return worker.StatusUnknown, monitor.WorkerProgress{}
	}

	statusKey := c.buildWorkerIDStatusCacheKey(workerID)
	progressKey := c.buildWorkerProgressCacheKey(workerID)

	command := c.redisClient.B().Mget().Key(statusKey, progressKey).Build()

	result := c.redisClient.Do(ctx, command)
	if err := result.Error(); err != nil {
		return worker.StatusUnknown, monitor.WorkerProgress{}
	}

	values, err := result.ToArray()
	if err != nil || len(values) < 2 {
		return worker.StatusUnknown, monitor.WorkerProgress{}
	}

	// Parse the status
	statusValue, err := c.parseRedisJSONValue(values[0].String())
	if err != nil {
		return worker.StatusUnknown, monitor.WorkerProgress{}
	}

	status, err := worker.StatusString(statusValue)
	if err != nil {
		status = worker.StatusUnknown
	}

	// Parse the progress
	progressValue, err := c.parseRedisJSONValue(values[1].String())
	if err != nil {
		return status, monitor.WorkerProgress{}
	}

	var workerProgress monitor.WorkerProgress

	if progressValue != "" {
		err = json.Unmarshal([]byte(progressValue), &workerProgress)
		if err != nil {
			return status, monitor.WorkerProgress{}
		}
	}

	return status, workerProgress
}

// extract the value field from the redis result string
func (c *Component) parseRedisJSONValue(jsonStr string) (string, error) {
	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", err
	}

	value, ok := data["Value"].(string)
	if !ok {
		return "", fmt.Errorf("value field is not a string")
	}

	return value, nil
}

// buildWorkerIDStatusCacheKey builds the cache key for the worker status by id.
func (c *Component) buildWorkerIDStatusCacheKey(workerID string) string {
	return fmt.Sprintf("worker:status:id:%s", workerID)
}

// buildWorkerProgressCacheKey builds the cache key for the worker progress by id.
func (c *Component) buildWorkerProgressCacheKey(workerID string) string {
	return fmt.Sprintf("worker:progress:%s", workerID)
}
