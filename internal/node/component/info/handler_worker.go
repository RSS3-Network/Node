package info

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/node/schema/worker/rss"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
)

type WorkerResponse struct {
	Data ComponentInfo `json:"data"`
}

type ComponentInfo struct {
	Decentralized []*WorkerInfo `json:"decentralized"`
	RSS           []*WorkerInfo `json:"rss"`
	Federated     []*WorkerInfo `json:"federated"`
}

type WorkerInfo struct {
	Network  network.Network        `json:"network"`
	Worker   worker.Worker          `json:"worker"`
	Tags     []tag.Tag              `json:"tags"`
	Platform decentralized.Platform `json:"platform"`
	Status   worker.Status          `json:"status"`
}

// WorkerKey is the key for the worker status aggregator.
type WorkerKey struct {
	Network network.Network
	Worker  worker.Worker
}

// WorkerStatusAggregator aggregates the statuses of workers with the same Network+Worker.
type WorkerStatusAggregator struct {
	Statuses []worker.Status
}

// GetWorkersStatus returns the status of all workers.
func (c *Component) GetWorkersStatus(ctx echo.Context) error {
	workerCount := len(c.config.Component.Decentralized) + len(c.config.Component.RSS) + len(c.config.Component.Federated)
	workerInfoChan := make(chan *WorkerInfo, workerCount)

	// Fetch all worker info concurrently.
	c.fetchAllWorkerInfo(ctx, workerInfoChan)

	// Aggregate the statuses of workers.
	aggregatedWorkers := c.aggregateWorkers(workerInfoChan)

	// Build the worker response.
	response := c.buildWorkerResponse(aggregatedWorkers)

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

	for _, w := range c.config.Component.Decentralized {
		fetchWorkerInfo(w, c.fetchDecentralizedWorkerInfo)
	}

	for _, w := range c.config.Component.RSS {
		fetchWorkerInfo(w, c.fetchRssWorkerInfo)
	}

	go func() {
		wg.Wait()
		close(workerInfoChan)
	}()
}

// aggregateWorkers aggregates the same workers through network + worker.
func (c *Component) aggregateWorkers(workerInfoChan <-chan *WorkerInfo) map[WorkerKey]*WorkerStatusAggregator {
	aggregatedWorkers := make(map[WorkerKey]*WorkerStatusAggregator)

	for workerInfo := range workerInfoChan {
		// construct the key for the worker status aggregator.
		key := WorkerKey{Network: workerInfo.Network, Worker: workerInfo.Worker}

		if _, exists := aggregatedWorkers[key]; !exists {
			aggregatedWorkers[key] = &WorkerStatusAggregator{Statuses: []worker.Status{}}
		}

		aggregatedWorkers[key].Statuses = append(aggregatedWorkers[key].Statuses, workerInfo.Status)
	}

	return aggregatedWorkers
}

// buildWorkerResponse aggregates statuses to determine the final status of a worker
func (c *Component) buildWorkerResponse(aggregatedWorkers map[WorkerKey]*WorkerStatusAggregator) *WorkerResponse {
	decentralizedWorkers := make([]*WorkerInfo, 0, len(aggregatedWorkers))
	rssWorkers := make([]*WorkerInfo, 0, len(aggregatedWorkers))

	for key, aggregator := range aggregatedWorkers {
		workerInfo := &WorkerInfo{
			Network: key.Network,
			Worker:  key.Worker,
			Status:  determineFinalStatus(aggregator.Statuses),
		}

		switch key.Network.Source() {
		case network.RSSSource:
			workerInfo.Tags = rss.ToTagsMap[key.Worker.(rss.Worker)]
			rssWorkers = append(rssWorkers, workerInfo)
		case network.EthereumSource, network.FarcasterSource, network.ArweaveSource:
			workerInfo.Tags = decentralized.ToTagsMap[key.Worker.(decentralized.Worker)]
			workerInfo.Platform = decentralized.ToPlatformMap[key.Worker.(decentralized.Worker)]
			decentralizedWorkers = append(decentralizedWorkers, workerInfo)
		}
	}

	return &WorkerResponse{
		Data: ComponentInfo{
			Decentralized: decentralizedWorkers,
			RSS:           rssWorkers,
		},
	}
}

func (c *Component) fetchDecentralizedWorkerInfo(ctx context.Context, module *config.Module) *WorkerInfo {
	// Fetch status from a specific worker by id.
	status := c.getWorkerStatusByID(ctx, module.ID)

	return &WorkerInfo{
		Network:  module.Network,
		Worker:   module.Worker,
		Platform: decentralized.ToPlatformMap[module.Worker.(decentralized.Worker)],
		Tags:     decentralized.ToTagsMap[module.Worker.(decentralized.Worker)],
		Status:   status,
	}
}

func (c *Component) fetchRssWorkerInfo(ctx context.Context, module *config.Module) *WorkerInfo {
	// Fetch status from a specific worker by id.
	status := c.getWorkerStatusByID(ctx, module.ID)

	return &WorkerInfo{
		Network: module.Network,
		Worker:  module.Worker,
		Tags:    rss.ToTagsMap[module.Worker.(rss.Worker)],
		Status:  status,
	}
}

// determineFinalStatus determines the final status of a worker based on the statuses of its instances.
// if all instances are ready, the final status is ready,
// at least one instance is indexing or ready, the final status is indexing
// otherwise, the final status is unhealthy
func determineFinalStatus(statuses []worker.Status) worker.Status {
	hasIndexing := false

	for _, status := range statuses {
		switch status {
		case worker.StatusIndexing:
			hasIndexing = true
		case worker.StatusReady:
		default:
			return worker.StatusUnhealthy
		}
	}

	if hasIndexing {
		return worker.StatusIndexing
	}

	return worker.StatusReady
}

// getWorkerStatusByID gets worker status from Redis cache by worker id.
func (c *Component) getWorkerStatusByID(ctx context.Context, workerID string) worker.Status {
	if c.redisClient == nil {
		return worker.StatusUnknown
	}

	command := c.redisClient.B().Get().Key(c.buildWorkerIDStatusCacheKey(workerID)).Build()

	result := c.redisClient.Do(ctx, command)
	if err := result.Error(); err != nil {
		return worker.StatusUnknown
	}

	// Convert the result to worker.Status.
	statusStr, err := result.ToString()
	if err != nil {
		return worker.StatusUnknown
	}

	status, err := worker.StatusString(statusStr)
	if err != nil {
		return worker.StatusUnknown
	}

	return status
}

// buildWorkerIDStatusCacheKey builds the cache key for the worker status by id.
func (c *Component) buildWorkerIDStatusCacheKey(workerID string) string {
	return fmt.Sprintf("worker:status:id::%s", workerID)
}
