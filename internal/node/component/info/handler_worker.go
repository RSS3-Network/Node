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
	monitor.WorkerProgress
}

// WorkerKey is the key for the worker status aggregator.
type WorkerKey struct {
	Network network.Network
	Worker  worker.Worker
}

// WorkerStatusAggregator aggregates the statuses of workers with the same Network+Worker.
type WorkerStatusAggregator struct {
	Statuses   []worker.Status
	Progresses []monitor.WorkerProgress
}

// GetWorkersStatus returns the status of all workers.
func (c *Component) GetWorkersStatus(ctx echo.Context) error {
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
			aggregatedWorkers[key] = &WorkerStatusAggregator{Statuses: []worker.Status{}, Progresses: []monitor.WorkerProgress{}}
		}

		aggregatedWorkers[key].Statuses = append(aggregatedWorkers[key].Statuses, workerInfo.Status)
		aggregatedWorkers[key].Progresses = append(aggregatedWorkers[key].Progresses, workerInfo.WorkerProgress)
	}

	return aggregatedWorkers
}

// buildWorkerResponse aggregates statuses to determine the final status of a worker
func (c *Component) buildWorkerResponse(aggregatedWorkers map[WorkerKey]*WorkerStatusAggregator) []*WorkerInfo {
	workers := make([]*WorkerInfo, 0, len(aggregatedWorkers))

	for key, aggregator := range aggregatedWorkers {
		finalStatus := determineFinalStatus(aggregator.Statuses)
		finalProgress := determineFinalProgress(aggregator.Progresses)

		workers = append(workers, &WorkerInfo{
			Network:        key.Network,
			Worker:         key.Worker,
			Status:         finalStatus,
			Tags:           worker.ToTagsMap[key.Worker],
			Platform:       worker.ToPlatformMap[key.Worker],
			WorkerProgress: finalProgress,
		})
	}

	return workers
}

func (c *Component) fetchWorkerInfo(ctx context.Context, module *config.Module) *WorkerInfo {
	// Fetch status and progress from a specific worker by id.
	status, workerProgress := c.getWorkerStatusAndProgressByID(ctx, module.ID)

	return &WorkerInfo{
		Network:  module.Network,
		Worker:   module.Worker,
		Platform: worker.ToPlatformMap[module.Worker],
		Tags:     worker.ToTagsMap[module.Worker],
		Status:   status,
		WorkerProgress: monitor.WorkerProgress{
			LatestRemoteBlock:  workerProgress.LatestRemoteBlock,
			LatestIndexedBlock: workerProgress.LatestIndexedBlock,
		},
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

// determineFinalProgress determines the final progress of a worker based on the progresses of its instances.
// if user runs more than one worker instance, we can determine the final progress by getting the larger value of each progress field
func determineFinalProgress(progresses []monitor.WorkerProgress) monitor.WorkerProgress {
	finalProgress := monitor.WorkerProgress{}

	for _, progress := range progresses {
		if progress.LatestRemoteBlock > finalProgress.LatestRemoteBlock {
			finalProgress.LatestRemoteBlock = progress.LatestRemoteBlock
		}

		if progress.LatestIndexedBlock > finalProgress.LatestIndexedBlock {
			finalProgress.LatestIndexedBlock = progress.LatestIndexedBlock
		}
	}

	return finalProgress
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
	return fmt.Sprintf("worker:status:id::%s", workerID)
}

// buildWorkerProgressCacheKey builds the cache key for the worker progress by id.
func (c *Component) buildWorkerProgressCacheKey(workerID string) string {
	return fmt.Sprintf("worker:progress::%s", workerID)
}
