package decentralized

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/schema/worker"
)

func (h *Hub) GetWorkers(c echo.Context) error {
	var wg sync.WaitGroup

	workerInfoChan := make(chan WorkerInfo, len(h.config.Node.Decentralized))

	workers := make([]WorkerInfo, 0, len(h.config.Node.Decentralized))

	for _, w := range h.config.Node.Decentralized {
		wg.Add(1)

		go func(w *config.Module) {
			defer wg.Done()

			status := h.getWorkerStatus(c.Request().Context(), w.Network.String(), w.Worker.String())

			workerInfo := WorkerInfo{
				Network:  w.Network,
				Worker:   w.Worker,
				Platform: worker.ToPlatformMap[w.Worker],
				Tags:     worker.ToTagsMap[w.Worker],
				Status:   status,
			}

			workerInfoChan <- workerInfo
		}(w)
	}

	go func() {
		wg.Wait()

		close(workerInfoChan)
	}()

	for workerInfo := range workerInfoChan {
		workers = append(workers, workerInfo)
	}

	return c.JSON(http.StatusOK, WorkerResponse{
		Data: workers,
	})
}

// getWorkerStatus gets worker status from Redis cache by network and workName.
func (h *Hub) getWorkerStatus(ctx context.Context, network, workerName string) worker.Status {
	if h.redisClient == nil {
		return worker.StatusUnknown
	}

	command := h.redisClient.B().Get().Key(h.buildCacheKey(network, workerName)).Build()

	result := h.redisClient.Do(ctx, command)
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

// buildCacheKey builds cache key for Redis.
func (h *Hub) buildCacheKey(network, workerName string) string {
	return fmt.Sprintf("worker:status::%s:%s", network, workerName)
}
