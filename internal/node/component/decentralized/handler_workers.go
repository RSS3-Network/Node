package decentralized

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
)

type WorkerResponse struct {
	Data []*WorkerInfo `json:"data"`
}

type WorkerInfo struct {
	Network  network.Network `json:"network"`
	Worker   worker.Worker   `json:"worker"`
	Tags     []tag.Tag       `json:"tags"`
	Platform worker.Platform `json:"platform"`
	Status   worker.Status   `json:"status"`
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

// GetWorkers returns the status of all workers.
func (c *Component) GetWorkers(ctx echo.Context) error {
	workerCount := len(c.config.Component.Decentralized)
	workerInfoChan := make(chan *WorkerInfo, workerCount)

	// Fetch all worker info concurrently.
	c.fetchAllWorkerInfo(ctx, workerInfoChan)

	// Aggregate the statuses of workers.
	aggregatedWorkers := c.aggregateWorkers(workerInfoChan)

	// Build the worker response.
	workers := c.buildWorkerResponse(aggregatedWorkers)

	return ctx.JSON(http.StatusOK, WorkerResponse{
		Data: workers,
	})
}

// fetchAllWorkerInfo fetches the status of all workers concurrently.
func (c *Component) fetchAllWorkerInfo(ctx echo.Context, workerInfoChan chan<- *WorkerInfo) {
	var wg sync.WaitGroup

	for _, w := range c.config.Component.Decentralized {
		wg.Add(1)

		go func(module *config.Module) {
			defer wg.Done()

			workerInfoChan <- c.fetchWorkerInfo(ctx.Request().Context(), module)
		}(w)
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
func (c *Component) buildWorkerResponse(aggregatedWorkers map[WorkerKey]*WorkerStatusAggregator) []*WorkerInfo {
	workers := make([]*WorkerInfo, 0, len(aggregatedWorkers))

	for key, aggregator := range aggregatedWorkers {
		finalStatus := determineFinalStatus(aggregator.Statuses)

		workers = append(workers, &WorkerInfo{
			Network:  key.Network,
			Worker:   key.Worker,
			Status:   finalStatus,
			Tags:     worker.ToTagsMap[key.Worker],
			Platform: worker.ToPlatformMap[key.Worker],
		})
	}

	return workers
}

func (c *Component) fetchWorkerInfo(ctx context.Context, module *config.Module) *WorkerInfo {
	// Fetch status from a specific worker by id.
	status := c.getWorkerStatusByID(ctx, module.ID)

	return &WorkerInfo{
		Network:  module.Network,
		Worker:   module.Worker,
		Platform: worker.ToPlatformMap[module.Worker],
		Tags:     worker.ToTagsMap[module.Worker],
		Status:   status,
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
