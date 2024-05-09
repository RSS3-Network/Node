package decentralized

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/schema/worker"
)

func (h *Hub) GetWorkers(c echo.Context) error {
	workers := make([]WorkerInfo, 0, len(h.config.Node.Decentralized))

	for _, w := range h.config.Node.Decentralized {
		workers = append(workers, WorkerInfo{
			Network:  w.Network,
			Worker:   w.Worker,
			Platform: worker.ToPlatformMap[w.Worker],
			Tag:      worker.ToTagsMap[w.Worker],
			Status:   h.getWorkerStatus(w.Network.String(), w.Worker.String()),
		})
	}

	return c.JSON(http.StatusOK, WorkerResponse{
		Data: workers,
	})
}

// getWorkerStatus gets worker status from Redis cache by network and workName.
func (h *Hub) getWorkerStatus(network, workerName string) worker.Status {
	if h.redisClient == nil {
		return worker.StatusUnknown
	}

	command := h.redisClient.B().Get().Key(h.buildCacheKey(network, workerName)).Build()

	result := h.redisClient.Do(context.TODO(), command)
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
