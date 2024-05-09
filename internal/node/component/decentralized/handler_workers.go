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

func (c *Component) GetWorkers(ctx echo.Context) error {
	var wg sync.WaitGroup

	workerCount := len(c.config.Component.Decentralized)

	workerInfoChan := make(chan *WorkerInfo, workerCount)

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

	workers := make([]*WorkerInfo, 0, workerCount)
	for workerInfo := range workerInfoChan {
		workers = append(workers, workerInfo)
	}

	return ctx.JSON(http.StatusOK, WorkerResponse{
		Data: workers,
	})
}

// getWorkerStatus gets worker status from Redis cache by network and workerName.
func (c *Component) getWorkerStatus(ctx context.Context, network, workerName string) worker.Status {
	if c.redisClient == nil {
		return worker.StatusUnknown
	}

	command := c.redisClient.B().Get().Key(c.buildCacheKey(network, workerName)).Build()

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

func (c *Component) fetchWorkerInfo(ctx context.Context, module *config.Module) *WorkerInfo {
	// Fetch status from a specific worker network.
	status := c.getWorkerStatus(ctx, module.Network.String(), module.Worker.String())

	return &WorkerInfo{
		Network:  module.Network,
		Worker:   module.Worker,
		Platform: worker.ToPlatformMap[module.Worker],
		Tags:     worker.ToTagsMap[module.Worker],
		Status:   status,
	}
}

// buildCacheKey builds cache key for Redis.
func (c *Component) buildCacheKey(network, workerName string) string {
	return fmt.Sprintf("worker:status::%s:%s", network, workerName)
}
