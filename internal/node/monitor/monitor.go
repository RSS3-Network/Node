package monitor

import (
	"context"
	"encoding/json"
	"fmt"

	workerx "github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/protocol-go/schema/network"
	"go.uber.org/zap"
)

type CheckpointState struct {
	BlockHeight      uint64 `json:"block_height"`
	BlockNumber      uint64 `json:"block_number"`
	EventID          uint64 `json:"event_id"`
	CastsBackfill    bool   `json:"casts_backfill"`
	ReactionBackfill bool   `json:"reaction_backfill"`
}

// MonitorWorkerStatus checks the worker status by comparing the current and latest block height/number.
// flags the worker as unhealthy if it's left behind the latest block height/number by more than the tolerance.
func (m *Monitor) MonitorWorkerStatus(ctx context.Context) error {
	for _, w := range m.config.Node.Decentralized {
		// get checkpoint state from database
		state, err := m.getCheckpointState(ctx, w.ID(), w.Network, w.Worker.String())
		if err != nil {
			zap.L().Error("get checkpoint state", zap.Error(err))
			return err
		}

		// get current indexing block height, number or event id and the latest block height, number, timestamp of network
		currentWorkerState, latestWorkerState, err := m.getWorkerIndexingStateByClients(ctx, w.Network, state)
		if err != nil {
			zap.L().Error("get latest block height or number", zap.Error(err))
			return err
		}

		// check worker's current status, and flag it as unhealthy if it's left behind the latest block height/number by more than the tolerance
		if err := m.flagUnhealthyWorker(ctx, w.Network.String(), w.Worker.String(), currentWorkerState, latestWorkerState, NetworkTorlerance[w.Network]); err != nil {
			return fmt.Errorf("detect unhealthy: %w", err)
		}
	}

	return nil
}

// getWorkerIndexingStateByClients gets the latest block height (arweave), block number (ethereum), event id (farcaster).
func (m *Monitor) getWorkerIndexingStateByClients(ctx context.Context, n network.Network, state CheckpointState) (uint64, uint64, error) {
	client, ok := m.clients[n]
	if !ok {
		return 0, 0, fmt.Errorf("client not ready")
	}

	latestState, err := client.LatestState(ctx)
	if err != nil {
		return 0, 0, fmt.Errorf("get latest state: %w", err)
	}

	return client.CurrentState(state), latestState, nil
}

// flagUnhealthyWorker detects by comparing the current and latest block height/number. If the difference is greater than the tolerance, the worker is flagged as unhealthy.
func (m *Monitor) flagUnhealthyWorker(ctx context.Context, network, workerName string, currentWorkerState, latestWorkerState, networkTolerance uint64) error {
	// if the worker is in Ready status and current block height/number is left behind the latest block height/number by more than the tolerance, flag the worker as unhealthy
	if m.getWorkerStatus(network, workerName) == workerx.StatusReady && latestWorkerState-currentWorkerState > networkTolerance {
		// Cache worker status to Redis.
		if err := m.updateWorkerStatus(ctx, network, workerName, workerx.StatusUnhealthy.String()); err != nil {
			return fmt.Errorf("cache token metadata: %w", err)
		}
	}

	return nil
}

// getCheckpointState gets the checkpoint state from the database.
func (m *Monitor) getCheckpointState(ctx context.Context, id string, network network.Network, worker string) (CheckpointState, error) {
	checkpoint, err := m.databaseClient.LoadCheckpoint(ctx, id, network, worker)
	if err != nil {
		return CheckpointState{}, fmt.Errorf("load checkpoint: %w", err)
	}

	var state CheckpointState
	if err := json.Unmarshal(checkpoint.State, &state); err != nil {
		zap.L().Error("unmarshal checkpoint state", zap.Error(err))
		return CheckpointState{}, err
	}

	return state, nil
}

// getWorkerStatus gets worker status from Redis cache by network and workerName.
func (m *Monitor) getWorkerStatus(network, workerName string) workerx.Status {
	if m.redisClient == nil {
		return workerx.StatusUnknown
	}

	command := m.redisClient.B().Get().Key(m.buildCacheKey(network, workerName)).Build()

	result := m.redisClient.Do(context.TODO(), command)
	if err := result.Error(); err != nil {
		return workerx.StatusUnknown
	}

	// Convert the result to worker.Status.
	statusStr, err := result.ToString()
	if err != nil {
		return workerx.StatusUnknown
	}

	status, err := workerx.StatusString(statusStr)
	if err != nil {
		return workerx.StatusUnknown
	}

	return status
}

// updateWorkerStatus caches the worker status to Redis.
func (m *Monitor) updateWorkerStatus(ctx context.Context, network, workerName string, status string) error {
	if m.redisClient == nil {
		return nil
	}

	command := m.redisClient.B().Set().Key(m.buildCacheKey(network, workerName)).Value(status).Build()

	result := m.redisClient.Do(ctx, command)
	if err := result.Error(); err != nil {
		return fmt.Errorf("redis result: %w", err)
	}

	return nil
}

// buildCacheKey builds cache key for Redis.
func (m *Monitor) buildCacheKey(network, workerName string) string {
	return fmt.Sprintf("worker:status::%s:%s", network, workerName)
}
