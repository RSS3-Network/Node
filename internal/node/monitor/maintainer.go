package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rss3-network/node/v2/config/parameter"
	"go.uber.org/zap"
)

func (m *Monitor) MaintainCoveragePeriod(ctx context.Context) error {
	// Get the latest epoch from the chain.
	epoch, err := m.settlementCaller.CurrentEpoch(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("failed to get current epoch: %w", err)
	}

	zap.L().Debug("retrieved current epoch", zap.Uint64("epoch", epoch.Uint64()))

	// Read the network parameters from the chain.
	params, err := m.networkParamsCaller.GetParams(&bind.CallOpts{}, epoch.Uint64())
	if err != nil {
		return fmt.Errorf("failed to get network parameters: %w", err)
	}

	zap.L().Debug("retrieved network parameters", zap.String("params", params))

	var paramsData parameter.NetworkParamsData

	if err := json.Unmarshal([]byte(params), &paramsData); err != nil {
		return fmt.Errorf("json unmarshal: %w", err)
	}

	// Load the coverage period from the config file.
	year, month, _ := time.Now().AddDate(0, -m.config.Database.CoveragePeriod, 0).Date()
	configTimestamp := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	zap.L().Debug("calculated coverage period timestamp",
		zap.Time("timestamp", configTimestamp),
		zap.Int("coverage_period_months", m.config.Database.CoveragePeriod))

	for network, start := range paramsData.NetworkStartBlock {
		timestamp := time.Unix(start.Timestamp, 0)

		if timestamp.After(configTimestamp) {
			zap.L().Debug("using config timestamp instead of network start time",
				zap.String("network", network.String()),
				zap.Time("network_start", timestamp),
				zap.Time("config_timestamp", configTimestamp))

			timestamp = configTimestamp
		}

		// Delete expired data.
		if err := m.databaseClient.DeleteExpiredActivities(ctx, network, timestamp); err != nil {
			zap.L().Error("delete coverage period", zap.Error(err), zap.String("network", network.String()), zap.Time("timestamp", timestamp))
		} else {
			zap.L().Info("successfully deleted expired activities",
				zap.String("network", network.String()),
				zap.Time("before_timestamp", timestamp))
		}
	}

	return nil
}
