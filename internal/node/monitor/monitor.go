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
		switch w.Network {
		case network.Arweave:
			// get arweave client
			arweaveClient := m.clients[w.Network].(arweave.Client)

			// get checkpoint state
			state, err := m.getCheckpointState(ctx, w.ID(), w.Network, w.Worker.String())
			if err != nil {
				zap.L().Error("get checkpoint state", zap.Error(err))
				return err
			}

			currentHeight := state.BlockHeight

			latestHeight, err := arweaveClient.GetBlockHeight(ctx)
			if err != nil {
				zap.L().Error("get latest block height", zap.Error(err))
				return err
			}

			if err := m.DetectUnhealthy(ctx, w.Network.String(), w.Worker.String(), currentHeight, uint64(latestHeight), 500); err != nil {
				return fmt.Errorf("detect unhealthy: %w", err)
			}
		case network.Farcaster:
			break
		default:
			evmClient := m.clients[w.Network].(ethereum.Client)

			// get checkpoint state
			state, err := m.getCheckpointState(ctx, w.ID(), w.Network, w.Worker.String())
			if err != nil {
				zap.L().Error("get checkpoint state", zap.Error(err))
				return err
			}

			currentNumber := state.BlockNumber

			latestBlock, err := evmClient.BlockNumber(ctx)
			if err != nil {
				zap.L().Error("get latest block number", zap.Error(err))
				return err
			}

			latestBlockNumber := latestBlock.Int64()

			if err := m.DetectUnhealthy(ctx, w.Network.String(), w.Worker.String(), currentNumber, uint64(latestBlockNumber), 500); err != nil {
				return fmt.Errorf("detect unhealthy: %w", err)
			}
		}
	}

	return nil
}

// DetectUnhealthy detects unhealthy workers.
func (m *Monitor) DetectUnhealthy(ctx context.Context, network, worker string, current, latest, tolerance uint64) error {
	if m.getWorkerStatus(network, worker) == workerx.StatusReady && latest-current > tolerance {
		// Cache worker status to Redis.
		if err := m.updateWorkerStatus(ctx, network, worker, workerx.StatusUnhealthy.String()); err != nil {
			return fmt.Errorf("cache token metadata: %w", err)
		}
	}

	return nil
}

// getCheckpointState gets the checkpoint state.
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

// getWorkerStatus gets the status of the worker.
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
