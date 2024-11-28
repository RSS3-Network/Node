package info

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/constant"
	"github.com/rss3-network/node/internal/node/component/decentralized"
	"github.com/rss3-network/node/internal/node/component/federated"
	"github.com/rss3-network/node/internal/node/component/rss"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type VersionResponse struct {
	Data Version `json:"data"`
}

type Version struct {
	Tag    string `json:"tag"`
	Commit string `json:"commit"`
}

type NodeInfoResponse struct {
	Data NodeInfo `json:"data"`
}

type GIRewardsResponse struct {
	Data []struct {
		ID            uint64 `json:"id"`
		Distributions []struct {
			ID            uint64 `json:"id"`
			RewardedNodes []struct {
				EpochID          uint64          `json:"epoch_id"`
				OperationRewards decimal.Decimal `json:"operation_rewards"`
				StakingRewards   decimal.Decimal `json:"staking_rewards"`
				RequestCount     decimal.Decimal `json:"request_count"`
			} `json:"rewarded_nodes"`
		} `json:"distributions"`
	} `json:"data"`
}

type GINodeInfoResponse struct {
	Data struct {
		LastHeartbeat int64           `json:"last_heartbeat"`
		SlashedTokens decimal.Decimal `json:"slashed_tokens"`
	} `json:"data"`
}

type NodeInfo struct {
	Operator common.Address `json:"operator"`
	Version  Version        `json:"version"`
	Uptime   int64          `json:"uptime"`
	Coverage WorkerCoverage `json:"coverage"`
	Records  Record         `json:"records"`
}

type WorkerSupportStatus struct {
	Supported   []string `json:"supported,omitempty"`
	Unsupported []string `json:"unsupported,omitempty"`
}

type WorkerCoverage struct {
	RSS           WorkerSupportStatus `json:"rss"`
	Decentralized WorkerSupportStatus `json:"decentralized"`
	Federated     WorkerSupportStatus `json:"federated"`
}

type Record struct {
	LastHeartbeat  int64           `json:"last_heartbeat"`
	RecentRequests []string        `json:"recent_requests"`
	RecentRewards  []Reward        `json:"recent_rewards"`
	SlashedTokens  decimal.Decimal `json:"slashed_tokens"`
}

type Reward struct {
	Epoch            uint64          `json:"epoch"`
	OperationRewards decimal.Decimal `json:"operation_rewards"`
	StakingRewards   decimal.Decimal `json:"staking_rewards"`
	RequestCounts    decimal.Decimal `json:"request_counts"`
}

// GetNodeOperator returns the node information.
func (c *Component) GetNodeOperator(ctx echo.Context) error {
	go c.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, "info")
	zap.L().Debug("getting node operator info")

	// Get Operator address info
	evmAddress := common.Address{}

	if operator := c.config.Discovery.Operator; operator != nil {
		evmAddress = operator.EvmAddress
		zap.L().Debug("found operator address",
			zap.String("address", evmAddress.String()))
	} else {
		zap.L().Warn("no operator address configured")
	}

	response := fmt.Sprintf("This is an RSS3 Node operated by %s.", evmAddress)

	zap.L().Debug("successfully retrieved node operator info")

	return ctx.JSON(http.StatusOK, response)
}

// GetNodeInfo returns the node information.
func (c *Component) GetNodeInfo(ctx echo.Context) error {
	go c.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, "info")
	zap.L().Debug("getting node info")

	// Get Version info
	version := c.buildVersion()
	zap.L().Debug("retrieved version info",
		zap.String("tag", version.Tag),
		zap.String("commit", version.Commit))

	// Get Operator address info
	evmAddress := common.Address{}

	if operator := c.config.Discovery.Operator; operator != nil {
		evmAddress = operator.EvmAddress
		zap.L().Debug("found operator address",
			zap.String("address", evmAddress.String()))
	} else {
		zap.L().Warn("no operator address configured")
	}

	// Get uptime info
	uptime, err := c.getNodeUptime()
	if err != nil {
		zap.L().Error("failed to get node uptime",
			zap.Error(err))
		return err
	}

	zap.L().Debug("retrieved node uptime",
		zap.Int64("uptime", uptime))

	// Get worker coverage info
	workerCoverage := c.getNodeWorkerCoverage()
	zap.L().Debug("retrieved worker coverage",
		zap.Int("rss_supported", len(workerCoverage.RSS.Supported)),
		zap.Int("decentralized_supported", len(workerCoverage.Decentralized.Supported)),
		zap.Int("federated_supported", len(workerCoverage.Federated.Supported)))

	// get reward info
	rewards, err := c.getNodeRewards(ctx.Request().Context(), evmAddress)
	if err != nil {
		zap.L().Error("failed to get node rewards",
			zap.Error(err))
		return err
	}

	zap.L().Debug("retrieved node rewards",
		zap.Int("reward_count", len(rewards)))

	// get last heartbeat and slashed tokens
	lastHeartbeat, slashedTokens, err := c.getNodeBasicInfo(ctx.Request().Context(), evmAddress)
	if err != nil {
		zap.L().Error("failed to get node basic info",
			zap.Error(err))

		return err
	}

	zap.L().Debug("retrieved node basic info",
		zap.Int64("last_heartbeat", lastHeartbeat),
		zap.String("slashed_tokens", slashedTokens.String()))

	var recentRequests []string

	if len(c.config.Component.Decentralized) > 0 {
		decentralizedRequests := decentralized.GetRecentRequest()
		recentRequests = append(recentRequests, decentralizedRequests...)
		zap.L().Debug("retrieved decentralized requests",
			zap.Int("count", len(decentralizedRequests)))
	}

	if len(c.config.Component.Federated) > 0 {
		federatedRequests := federated.GetRecentRequest()
		recentRequests = append(recentRequests, federatedRequests...)
		zap.L().Debug("retrieved federated requests",
			zap.Int("count", len(federatedRequests)))
	}

	if c.config.Component.RSS != nil {
		rssRequests := rss.GetRecentRequest()
		recentRequests = append(recentRequests, rssRequests...)
		zap.L().Debug("retrieved RSS requests",
			zap.Int("count", len(rssRequests)))
	}

	zap.L().Debug("successfully retrieved all node info",
		zap.String("operator", evmAddress.String()),
		zap.Int("total_recent_requests", len(recentRequests)))

	return ctx.JSON(http.StatusOK, NodeInfoResponse{
		Data: NodeInfo{
			Version:  version,
			Operator: evmAddress,
			Uptime:   uptime,
			Coverage: workerCoverage,
			Records: Record{
				LastHeartbeat:  lastHeartbeat,
				RecentRequests: recentRequests,
				RecentRewards:  rewards,
				SlashedTokens:  slashedTokens,
			},
		},
	})
}

func (c *Component) buildVersion() Version {
	tag, commit := constant.BuildVersionDetail()

	return Version{
		Tag:    tag,
		Commit: commit,
	}
}

// getNodeWorkerCoverage returns the worker coverage in network-worker format.
func (c *Component) getNodeWorkerCoverage() WorkerCoverage {
	workerCoverage := WorkerCoverage{}

	addCoverage := func(network, workerName string, supported, unsupported *[]string) {
		coverage := fmt.Sprintf("%s_%s", network, workerName)
		target := unsupported
		logMessage := "added unsupported worker coverage"

		if workerName != "" {
			target = supported
			logMessage = "added worker coverage"
		}

		*target = append(*target, coverage)
		zap.L().Debug(logMessage, zap.String("coverage", coverage))
	}

	processWorkers := func(workers []*config.Module, supported, unsupported *[]string) {
		for _, worker := range workers {
			addCoverage(worker.Network.String(), worker.Worker.Name(), supported, unsupported)
		}
	}

	processWorkers(c.config.Component.Decentralized, &workerCoverage.Decentralized.Supported, &workerCoverage.Decentralized.Unsupported)

	if len(c.config.Component.Federated) > 0 {
		processWorkers(c.config.Component.Federated, &workerCoverage.Federated.Supported, &workerCoverage.Federated.Unsupported)
	} else {
		addCoverage(network.Mastodon.String(), "core", &workerCoverage.Federated.Unsupported, &workerCoverage.Federated.Unsupported)
	}

	if c.config.Component.RSS != nil {
		addCoverage(network.RSSHub.String(), c.config.Component.RSS.Worker.Name(), &workerCoverage.RSS.Supported, &workerCoverage.RSS.Unsupported)
	} else {
		addCoverage(network.RSSHub.String(), "core", &workerCoverage.RSS.Unsupported, &workerCoverage.RSS.Unsupported)
	}

	for net, workers := range NetworkToWorkersMap {
		if net == network.RSSHub || net == network.SatoshiVM || net == network.Mastodon {
			continue
		}

		for _, worker := range workers {
			coverage := fmt.Sprintf("%s_%s", net, worker.Name())
			if !lo.Contains(workerCoverage.Decentralized.Supported, coverage) {
				workerCoverage.Decentralized.Unsupported = append(workerCoverage.Decentralized.Unsupported, coverage)
				zap.L().Debug("added decentralized unsupported worker coverage", zap.String("coverage", coverage))
			}
		}
	}

	uniqueAndSort := func(supported, unsupported *[]string) {
		*supported = lo.Uniq(*supported)

		*unsupported = lo.Uniq(*unsupported)

		sort.Strings(*supported)

		sort.Strings(*unsupported)
	}

	uniqueAndSort(&workerCoverage.RSS.Supported, &workerCoverage.RSS.Unsupported)
	uniqueAndSort(&workerCoverage.Decentralized.Supported, &workerCoverage.Decentralized.Unsupported)
	uniqueAndSort(&workerCoverage.Federated.Supported, &workerCoverage.Federated.Unsupported)

	zap.L().Debug("completed getting worker coverage",
		zap.Int("total_coverage", len(workerCoverage.RSS.Supported)+len(workerCoverage.Decentralized.Supported)+len(workerCoverage.Federated.Supported)))

	return workerCoverage
}

// getNodeUptime returns the node uptime.
func (c *Component) getNodeUptime() (int64, error) {
	zap.L().Debug("getting node uptime")

	var uptime int64

	// get first start time from redis cache and calculate uptime
	firstStartTime, err := GetFirstStartTime()
	if err != nil {
		zap.L().Error("failed to get first start time from cache",
			zap.Error(err))
		return 0, err
	}

	if firstStartTime == 0 {
		uptime = 0

		zap.L().Info("first time node startup detected")

		err := UpdateFirstStartTime(time.Now().Unix())
		if err != nil {
			zap.L().Error("failed to update first start time",
				zap.Error(err))

			return 0, err
		}

		zap.L().Debug("successfully updated first start time")
	} else {
		uptime = time.Now().Unix() - firstStartTime
		zap.L().Debug("calculated uptime",
			zap.Int64("uptime_seconds", uptime))
	}

	return uptime, nil
}

// getNodeRewards returns the node rewards.
func (c *Component) getNodeRewards(ctx context.Context, address common.Address) ([]Reward, error) {
	zap.L().Debug("getting node rewards",
		zap.String("address", address.String()))

	var resp GIRewardsResponse

	err := c.sendRequest(ctx, fmt.Sprintf("/nta/epochs/%s/rewards", address), &resp)
	if err != nil {
		zap.L().Error("failed to get node rewards", zap.Error(err))

		return nil, err
	}

	rewards := make([]Reward, 0, len(resp.Data))

	for _, data := range resp.Data {
		if len(data.Distributions) > 0 && len(data.Distributions[0].RewardedNodes) > 0 {
			rewardedNode := data.Distributions[0].RewardedNodes[0]

			var reward Reward

			if rewardedNode.EpochID == data.ID {
				reward = Reward{
					Epoch:            data.ID,
					OperationRewards: rewardedNode.OperationRewards,
					StakingRewards:   rewardedNode.StakingRewards,
					RequestCounts:    rewardedNode.RequestCount,
				}

				zap.L().Debug("processed reward for epoch",
					zap.Uint64("epoch", data.ID),
					zap.String("operation_rewards", rewardedNode.OperationRewards.String()),
					zap.String("staking_rewards", rewardedNode.StakingRewards.String()))
			}

			rewards = append(rewards, reward)
		}
	}

	zap.L().Debug("successfully retrieved node rewards",
		zap.Int("total_rewards", len(rewards)))

	return rewards, nil
}

// getNodeInfo returns the node slashed token.
func (c *Component) getNodeBasicInfo(ctx context.Context, address common.Address) (int64, decimal.Decimal, error) {
	zap.L().Debug("getting node basic info",
		zap.String("address", address.String()))

	var resp GINodeInfoResponse

	err := c.sendRequest(ctx, fmt.Sprintf("/nta/nodes/%s", address), &resp)
	if err != nil {
		zap.L().Error("failed to get node basic info",
			zap.String("address", address.String()),
			zap.Error(err))

		return 0, decimal.Decimal{}, err
	}

	zap.L().Debug("successfully retrieved node basic info",
		zap.Int64("last_heartbeat", resp.Data.LastHeartbeat),
		zap.String("slashed_tokens", resp.Data.SlashedTokens.String()))

	return resp.Data.LastHeartbeat, resp.Data.SlashedTokens, nil
}

// sendRequest sends a request to the global indexer.
func (c *Component) sendRequest(ctx context.Context, path string, result any) error {
	internalURL, err := url.Parse(c.config.Discovery.Server.GlobalIndexerEndpoint)
	if err != nil {
		zap.L().Error("failed to parse global indexer endpoint",
			zap.Error(err))
		return err
	}

	internalURL.Path = path

	zap.L().Debug("sending request to global indexer",
		zap.String("url", internalURL.String()))

	body, err := c.httpClient.Fetch(ctx, internalURL.String())
	if err != nil {
		return fmt.Errorf("fetch request: %w", err)
	}
	defer body.Close()

	if err = json.NewDecoder(body).Decode(&result); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	zap.L().Debug("successfully completed request to global indexer")

	return nil
}

// GetFirstStartTime Get the first start time from local file
func GetFirstStartTime() (int64, error) {
	filePath := "first_start_time.txt"

	zap.L().Debug("getting first start time from file",
		zap.String("path", filePath))

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		zap.L().Debug("first start time file does not exist")
		return 0, nil
	}

	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return 0, fmt.Errorf("read file: %w", err)
	}

	// Convert content to int64
	firstStartTime, err := strconv.ParseInt(strings.TrimSpace(string(content)), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse int64: %w", err)
	}

	zap.L().Debug("successfully retrieved first start time",
		zap.Int64("timestamp", firstStartTime))

	return firstStartTime, nil
}

// UpdateFirstStartTime updates the first start time in local file
func UpdateFirstStartTime(timestamp int64) error {
	filePath := "first_start_time.txt"

	zap.L().Debug("updating first start time",
		zap.String("path", filePath),
		zap.Int64("timestamp", timestamp))

	// Convert timestamp to string
	content := strconv.FormatInt(timestamp, 10)

	// Write to file
	err := os.WriteFile(filePath, []byte(content), 0600)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	zap.L().Debug("successfully updated first start time")

	return nil
}
