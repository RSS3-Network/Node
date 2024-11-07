package monitor

import (
	"context"
	"fmt"
	"sync"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"go.uber.org/zap"
)

func (m *Monitor) MonitorMockWorkerStatus(ctx context.Context, currentState CheckpointState, targetWorkerState, latestState uint64) error {
	var wg sync.WaitGroup

	errChan := make(chan error, len(m.config.Component.Decentralized)+len(m.config.Component.Federated))

	for _, w := range m.config.Component.Decentralized {
		wg.Add(1)

		go func(w *config.Module) {
			defer wg.Done()

			if err := m.processMockWorker(ctx, w, currentState, targetWorkerState, latestState); err != nil {
				errChan <- err
			}
		}(w)
	}

	for _, w := range m.config.Component.Federated {
		wg.Add(1)

		go func(w *config.Module) {
			defer wg.Done()

			if err := m.processMockFederatedWorker(ctx, w); err != nil {
				errChan <- err
			}
		}(w)
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

// processWorker processes the decentralized worker status.
func (m *Monitor) processMockWorker(ctx context.Context, w *config.Module, currentState CheckpointState, targetWorkerState, latestWorkerState uint64) error {
	// get current indexing block height, number or event id and the latest block height, number, timestamp of network
	currentWorkerState, _, _, err := m.getWorkerIndexingStateByClients(ctx, w.Network, w.Worker.Name(), currentState, w.Parameters)
	if err != nil {
		zap.L().Error("get latest block height or number", zap.Error(err))
		return err
	}

	networkTolerance := parameter.CurrentNetworkTolerance[w.Network]

	if w.Worker.Name() == decentralized.Momoka.String() {
		networkTolerance = parameter.CurrentNetworkTolerance[w.Network] * 120000
	}

	// check worker's current status, and flag it as unhealthy if it's left behind the latest block height/number by more than the tolerance
	if err := m.flagWorkerStatus(ctx, w.ID, currentWorkerState, targetWorkerState, latestWorkerState, networkTolerance); err != nil {
		return fmt.Errorf("detect unhealthy: %w", err)
	}

	//	update worker progress (state)
	if err := m.UpdateWorkerProgress(ctx, w.ID, ConstructWorkerProgress(currentWorkerState, targetWorkerState, latestWorkerState, 0)); err != nil {
		return fmt.Errorf("update worker progress: %w", err)
	}

	return nil
}

// processFederatedWorker processes the federated worker status.
func (m *Monitor) processMockFederatedWorker(ctx context.Context, w *config.Module) error {
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
