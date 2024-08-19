package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/config/parameter"
	workerx "github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CheckpointState struct {
	BlockHeight      uint64 `json:"block_height"`
	BlockTimestamp   uint64 `json:"block_timestamp"`
	BlockNumber      uint64 `json:"block_number"`
	EventID          uint64 `json:"event_id"`
	CastsBackfill    bool   `json:"casts_backfill"`
	ReactionBackfill bool   `json:"reaction_backfill"`
}

type WorkerProgress struct {
	RemoteState  uint64 `json:"remote_state"`
	IndexedState uint64 `json:"indexed_state"`
	IndexCount   int64  `json:"index_count"`
}

// MonitorWorkerStatus checks the worker status by comparing the current and latest block height/number.
// flags the worker as unhealthy if it's left behind the latest block height/number by more than the tolerance.
func (m *Monitor) MonitorWorkerStatus(ctx context.Context) error {
	var wg sync.WaitGroup

	errChan := make(chan error, len(m.config.Component.Decentralized)+lo.Ternary(m.config.Component.RSS != nil, 1, 0)+len(m.config.Component.Federated))

	processWorker := func(w *config.Module, processFunc func(context.Context, *config.Module) error) {
		wg.Add(1)

		go func(w *config.Module) {
			defer wg.Done()

			if err := processFunc(ctx, w); err != nil {
				errChan <- err
			}
		}(w)
	}

	for _, w := range m.config.Component.Decentralized {
		processWorker(w, m.processDecentralizedWorker)
	}

	if m.config.Component.RSS != nil {
		processWorker(m.config.Component.RSS, m.processRSSWorker)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

// processDecentralizedWorker processes the decentralized worker status.
func (m *Monitor) processDecentralizedWorker(ctx context.Context, w *config.Module) error {
	// get checkpoint info from database
	indexCount, state, err := m.getCheckpointState(ctx, w.ID, w.Network, w.Worker.Name())
	if err != nil {
		zap.L().Error("get checkpoint info", zap.Error(err))
		return err
	}

	// get current indexing block height, number or event id and the latest block height, number, timestamp of network
	currentWorkerState, targetWorkerState, latestWorkerState, err := m.getWorkerIndexingStateByClients(ctx, w.Network, w.Worker.Name(), state, w.Parameters)
	if err != nil {
		zap.L().Error("get latest block height or number", zap.Error(err))
		return err
	}

	networkTolerance := parameter.CurrentNetworkTolerance[w.Network]

	if w.Worker.Name() == decentralized.Momoka.String() {
		networkTolerance = parameter.CurrentNetworkTolerance[w.Network] * 120000
	}

	// check worker's current status, and flag it as unhealthy if it's behind the latest block height/number by more than the tolerated amount
	if err := m.flagWorkerStatus(ctx, w.ID, currentWorkerState, targetWorkerState, latestWorkerState, networkTolerance); err != nil {
		return fmt.Errorf("detect unhealthy: %w", err)
	}

	if err := m.UpdateWorkerProgress(ctx, w.ID, ConstructWorkerProgress(currentWorkerState, targetWorkerState, latestWorkerState, indexCount)); err != nil {
		return fmt.Errorf("update worker progress: %w", err)
	}

	return nil
}

// processRSSWorker processes the rss worker status.
func (m *Monitor) processRSSWorker(ctx context.Context, w *config.Module) error {
	client, ok := m.clients[w.Network]
	if !ok {
		return fmt.Errorf("client not exist")
	}

	targetStatus := workerx.StatusReady
	if _, _, err := client.LatestState(ctx); err != nil {
		targetStatus = workerx.StatusUnhealthy
	}

	return m.UpdateWorkerStatusByID(ctx, w.ID, targetStatus.String())
}

// getWorkerIndexingStateByClients gets the latest block height (arweave), block number (ethereum), event id (farcaster).
func (m *Monitor) getWorkerIndexingStateByClients(ctx context.Context, n network.Network, w string, state CheckpointState, param *config.Parameters) (uint64, uint64, uint64, error) {
	client, ok := m.clients[n]
	if !ok {
		return 0, 0, 0, fmt.Errorf("client not ready")
	}

	var current, target, latest uint64

	currentBlock, currentBlockTimestamp := client.CurrentState(state)
	targetBlock, targetBlockTimestamp := client.TargetState(param)
	latestBlock, latestBlockTimestamp, err := client.LatestState(ctx)

	if err != nil {
		return 0, 0, 0, fmt.Errorf("get latest state: %w", err)
	}

	if w == decentralized.Momoka.String() {
		current = currentBlockTimestamp
		target = targetBlockTimestamp
		latest = latestBlockTimestamp
	} else {
		current = currentBlock
		target = targetBlock
		latest = latestBlock
	}

	return current, target, latest, nil
}

// flagWorkerStatus detects by comparing the current and latest block height/number. Could be converted to ready, indexing, unhealthy.
func (m *Monitor) flagWorkerStatus(ctx context.Context, workerID string, currentWorkerState, targetWorkerState, latestWorkerState, networkTolerance uint64) error {
	currentStatus := m.GetWorkerStatusByID(ctx, workerID)
	targetStatus := currentStatus

	if latestWorkerState < currentWorkerState {
		return nil
	}

	switch currentStatus {
	case workerx.StatusUnknown:
		// if the worker status is unknown and the current worker state is greater than 0, flag it as indexing
		if currentWorkerState > 0 {
			targetStatus = workerx.StatusIndexing
		} else {
			targetStatus = workerx.StatusUnhealthy
		}
	case workerx.StatusIndexing:
		if currentWorkerState <= m.getWorkerProgress(ctx, workerID).IndexedState {
			// if the worker is indexing but didn't make any progress in the last cycle, flag it as unhealthy
			targetStatus = workerx.StatusUnhealthy
			break
		}

		if targetWorkerState == 0 && latestWorkerState-currentWorkerState < networkTolerance || targetWorkerState > 0 && currentWorkerState >= targetWorkerState {
			// if the worker is indexing without a target param and the difference between the current and latest block height/number is less than the tolerance, flag it as ready
			// or if the worker is indexing with a target param and the current state reaches the target state, flag it as ready
			targetStatus = workerx.StatusReady
		}
	case workerx.StatusReady:
		// if the worker is ready and the target block height/number is set, it will remain ready because it's finished.
		// otherwise the difference between the current and latest block height/number is greater than the tolerance, flag it as indexing
		if targetWorkerState == 0 && latestWorkerState-currentWorkerState > networkTolerance {
			targetStatus = workerx.StatusIndexing
		}
	case workerx.StatusUnhealthy:
		if currentWorkerState > m.getWorkerProgress(ctx, workerID).IndexedState {
			// if the worker is unhealthy and made progress in the last cycle, flag it as indexing
			targetStatus = workerx.StatusIndexing
		}
	default:
	}

	if targetStatus != currentStatus {
		if err := m.UpdateWorkerStatusByID(ctx, workerID, targetStatus.String()); err != nil {
			return fmt.Errorf("update worker status: %w", err)
		}
	}

	return nil
}

// getCheckpointState gets the checkpoint state from the database.
func (m *Monitor) getCheckpointState(ctx context.Context, id string, network network.Network, worker string) (int64, CheckpointState, error) {
	checkpoint, err := m.databaseClient.LoadCheckpoint(ctx, id, network, worker)
	if err != nil {
		return 0, CheckpointState{}, fmt.Errorf("load checkpoint: %w", err)
	}

	var state CheckpointState
	if err := json.Unmarshal(checkpoint.State, &state); err != nil {
		zap.L().Error("unmarshal checkpoint state", zap.Error(err))
		return 0, CheckpointState{}, err
	}

	return checkpoint.IndexCount, state, nil
}

// GetWorkerStatusByID gets worker status from Redis cache by network and workerName.
func (m *Monitor) GetWorkerStatusByID(ctx context.Context, workerID string) workerx.Status {
	if m.redisClient == nil {
		return workerx.StatusUnknown
	}

	command := m.redisClient.B().Get().Key(m.buildWorkerIDStatusCacheKey(workerID)).Build()

	result := m.redisClient.Do(ctx, command)
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

// UpdateWorkerStatusByID updates the worker status in the Redis cache by id.
func (m *Monitor) UpdateWorkerStatusByID(ctx context.Context, workerID string, status string) error {
	if m.redisClient == nil {
		return nil
	}

	command := m.redisClient.B().Set().Key(m.buildWorkerIDStatusCacheKey(workerID)).Value(status).Build()

	result := m.redisClient.Do(ctx, command)
	if err := result.Error(); err != nil {
		return fmt.Errorf("redis result: %w", err)
	}

	return nil
}

// getWorkerProgress gets worker progress from Redis cache by worker ID.
func (m *Monitor) getWorkerProgress(ctx context.Context, workerID string) WorkerProgress {
	var progress WorkerProgress
	if m.redisClient == nil {
		return progress
	}

	command := m.redisClient.B().Get().Key(m.buildWorkerProgressCacheKey(workerID)).Build()

	result := m.redisClient.Do(ctx, command)
	if err := result.Error(); err != nil {
		return progress
	}

	progressStr, err := result.ToString()
	if err != nil {
		return progress
	}

	err = json.Unmarshal([]byte(progressStr), &progress)
	if err != nil {
		return WorkerProgress{}
	}

	return progress
}

// UpdateWorkerProgress updates worker progress (state) in each monitoring cycle to Redis Cache.
func (m *Monitor) UpdateWorkerProgress(ctx context.Context, workerID string, progress WorkerProgress) error {
	if m.redisClient == nil {
		return nil
	}

	progressJSON, err := json.Marshal(progress)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	command := m.redisClient.B().Set().Key(m.buildWorkerProgressCacheKey(workerID)).Value(string(progressJSON)).Build()

	result := m.redisClient.Do(ctx, command)
	if err := result.Error(); err != nil {
		return fmt.Errorf("redis result: %w", err)
	}

	return nil
}

// buildWorkerIDStatusCacheKey builds the cache key for the worker status by id.
func (m *Monitor) buildWorkerIDStatusCacheKey(workerID string) string {
	return fmt.Sprintf("worker:status:id:%s", workerID)
}

// buildWorkerProgressCacheKey builds cache key for worker current state in each monitoring cycle.
func (m *Monitor) buildWorkerProgressCacheKey(workerID string) string {
	return fmt.Sprintf("worker:progress:%s", workerID)
}

// ConstructWorkerProgress constructs the worker progress from current, target and latest block height/number.
func ConstructWorkerProgress(currentWorkerState, targetWorkerState, latestWorkerState uint64, indexCount int64) WorkerProgress {
	workerProgress := WorkerProgress{
		RemoteState:  latestWorkerState,
		IndexedState: currentWorkerState,
		IndexCount:   indexCount,
	}

	// set remote to target if it's set
	if targetWorkerState > 0 {
		workerProgress.RemoteState = targetWorkerState
	}

	return workerProgress
}
