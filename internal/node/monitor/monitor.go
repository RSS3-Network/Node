package monitor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rss3-network/node/provider/arweave"
	"github.com/rss3-network/node/provider/ethereum"
	workerx "github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/protocol-go/schema/network"
	"go.uber.org/zap"
)

type CheckpointState struct {
	BlockHeight uint64 `json:"block_height"`
	BlockNumber uint64 `json:"block_number"`
	EventID     uint64 `json:"event_id"`
}

func (m *Monitor) CheckWorkerStatus(ctx context.Context) error {
	for _, w := range m.config.Node.Decentralized {
		// get checkpoint state from database
		state, err := m.getCheckpointState(ctx, w.ID(), w.Network, w.Worker.String())
		if err != nil {
			zap.L().Error("get checkpoint state", zap.Error(err))
			return err
		}

		// get current indexing block height, number or event id and the latest block height, number, timestamp of network
		currentBlockState, latestBlockState, err := m.getLatestBlockHeightOrNumberByClients(ctx, w.Network, state)
		if err != nil {
			zap.L().Error("get latest block height or number", zap.Error(err))
			return err
		}

		if err := m.flagUnhealthyWorker(ctx, w.Network.String(), w.Worker.String(), currentBlockState, latestBlockState, 500); err != nil {
			return fmt.Errorf("detect unhealthy: %w", err)
		}
	}

	return nil
}

// getLatestBlockHeightOrNumber gets the latest block height (arweave), block number (ethereum), event id (farcaster).
func (m *Monitor) getLatestBlockHeightOrNumberByClients(ctx context.Context, n network.Network, state CheckpointState) (uint64, uint64, error) {
	switch n {
	case network.Arweave:
		currentBlockState := state.BlockHeight

		arweaveClient := m.clients[n].(arweave.Client)

		latestBlockState, err := arweaveClient.GetBlockHeight(ctx)
		if err != nil {
			zap.L().Error("get latest block height", zap.Error(err))
			return 0, 0, err
		}

		return currentBlockState, uint64(latestBlockState), nil
	case network.Farcaster:
		return 0, 0, nil
	default:
		currentBlockState := state.BlockNumber

		evmClient := m.clients[n].(ethereum.Client)

		latestBlockState, err := evmClient.BlockNumber(ctx)
		if err != nil {
			return 0, 0, err
		}

		return currentBlockState, uint64(latestBlockState.Int64()), nil
	}
}

// flagUnhealthyWorker detects by comparing the current and latest block height/number. If the difference is greater than the tolerance, the worker is flagged as unhealthy.
func (m *Monitor) flagUnhealthyWorker(ctx context.Context, network, workerName string, currentBlockState, latestBlockState, networkTolerance uint64) error {
	if m.getWorkerStatus(network, workerName) == workerx.StatusReady && latestBlockState-currentBlockState > networkTolerance {
		// Cache worker status to Redis.
		if err := m.updateWorkerStatus(ctx, network, workerName, workerx.StatusUnhealthy.String()); err != nil {
			return fmt.Errorf("cache token metadata: %w", err)
		}
	}

	return nil
}

// getCheckpointState gets the checkpoint state from the database.
func (m *Monitor) getCheckpointState(ctx context.Context, id string, network network.Network, worker string) (CheckpointState, error) {
	// load checkpoint
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

// getWorkerStatus gets worker status from Redis cache by network and workName.
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
