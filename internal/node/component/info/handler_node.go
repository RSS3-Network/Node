package info

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/node/internal/constant"
	"github.com/rss3-network/node/internal/node/component/decentralized"
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
	Operator   common.Address `json:"operator"`
	Version    Version        `json:"version"`
	Uptime     int64          `json:"uptime"`
	Parameters string         `json:"parameters"`
	Coverage   []string       `json:"coverage"`
	Records    Record         `json:"records"`
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

	zap.L().Debug("get node info", zap.String("request.ip", ctx.Request().RemoteAddr))

	// Get Operator address info
	evmAddress := "0x0000000000000000000000000000000000000000"

	if operator := c.config.Discovery.Operator; operator != nil {
		evmAddress = operator.EvmAddress.String()
	}

	response := fmt.Sprintf("This is an RSS3 Node operated by %s.", evmAddress)

	return ctx.JSON(http.StatusOK, response)
}

// GetNodeInfo returns the node information.
func (c *Component) GetNodeInfo(ctx echo.Context) error {
	go c.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, "info")

	zap.L().Debug("get node info", zap.String("request.ip", ctx.Request().RemoteAddr))

	// Get Version info
	version := c.buildVersion()

	// Get Operator address info
	evmAddress := common.Address{}

	if !lo.IsEmpty(c.config.Discovery.Operator.EvmAddress) {
		evmAddress = c.config.Discovery.Operator.EvmAddress
	}

	// Get network params info
	params, err := c.getNetworkParams(ctx.Request().Context())
	if err != nil {
		return fmt.Errorf("failed to get network params: %w", err)
	}

	// Get uptime info
	uptime, err := c.getNodeUptime(ctx.Request().Context())
	if err != nil {
		return fmt.Errorf("failed to get node uptime: %w", err)
	}

	// Get worker coverage info
	workerCoverage := c.getNodeWorkerCoverage()

	// get reward info
	rewards, err := c.getNodeRewards(ctx.Request().Context(), evmAddress)
	if err != nil {
		return fmt.Errorf("failed to get node rewards: %w", err)
	}

	// get last heartbeat and slashed tokens
	lastHeartbeat, slashedTokens, err := c.getNodeBasicInfo(ctx.Request().Context(), evmAddress)
	if err != nil {
		return fmt.Errorf("failed to get node slashed tokens: %w", err)
	}

	return ctx.JSON(http.StatusOK, NodeInfoResponse{
		Data: NodeInfo{
			Version:    version,
			Operator:   evmAddress,
			Parameters: params,
			Uptime:     uptime,
			Coverage:   workerCoverage,
			Records: Record{
				LastHeartbeat:  lastHeartbeat,
				RecentRequests: decentralized.GetRecentRequest(),
				RecentRewards:  rewards,
				SlashedTokens:  slashedTokens,
			},
		},
	})
}

// GetVersion returns the version of the network component.
func (c *Component) GetVersion(ctx echo.Context) error {
	version := c.buildVersion()

	return ctx.JSON(http.StatusOK, VersionResponse{
		Data: version,
	})
}

func (c *Component) buildVersion() Version {
	tag, commit := constant.BuildVersionDetail()

	return Version{
		Tag:    tag,
		Commit: commit,
	}
}

// getNodeWorkerCoverage returns the worker coverage.
func (c *Component) getNodeWorkerCoverage() []string {
	workerCoverage := make([]string, 0, len(c.config.Component.Decentralized)+lo.Ternary(c.config.Component.RSS != nil, 1, 0)+len(c.config.Component.Federated))

	// append all workers
	for _, worker := range c.config.Component.Decentralized {
		workerCoverage = append(workerCoverage, worker.Worker.Name())
	}

	if c.config.Component.RSS != nil {
		workerCoverage = append(workerCoverage, c.config.Component.RSS.Worker.Name())
	}

	for _, worker := range c.config.Component.Federated {
		workerCoverage = append(workerCoverage, worker.Worker.Name())
	}

	return lo.Uniq(workerCoverage)
}

// getNodeUptime returns the node uptime.
func (c *Component) getNodeUptime(ctx context.Context) (int64, error) {
	var uptime int64

	// get first start time from redis cache and calculate uptime
	firstStartTime, err := GetFirstStartTime(ctx, c.redisClient)
	if err != nil {
		return 0, fmt.Errorf("failed to get first start time from cache: %w", err)
	}

	if firstStartTime == 0 {
		uptime = 0

		err := UpdateFirstStartTime(ctx, c.redisClient, time.Now().Unix())
		if err != nil {
			return 0, fmt.Errorf("failed to update first start time: %w", err)
		}
	} else {
		uptime = time.Now().Unix() - firstStartTime
	}

	return uptime, nil
}

// getNetworkParams returns the network parameters.
func (c *Component) getNetworkParams(ctx context.Context) (string, error) {
	currentEpoch, err := parameter.GetCurrentEpoch(ctx, c.redisClient)
	if err != nil {
		return "", fmt.Errorf("failed to get current epoch from cache: %w", err)
	}

	// Get parameters from the network
	params, err := parameter.PullNetworkParamsFromVSL(c.networkParamsCaller, uint64(currentEpoch))
	if err != nil {
		return "", fmt.Errorf("failed to get network parameters: %w", err)
	}

	return params, nil
}

// getNodeRewards returns the node rewards.
func (c *Component) getNodeRewards(ctx context.Context, address common.Address) ([]Reward, error) {
	var resp GIRewardsResponse

	err := c.sendRequest(ctx, fmt.Sprintf("/nta/epochs/%s/rewards", address), &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to get node rewards: %w", err)
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
			}

			rewards = append(rewards, reward)
		}
	}

	return rewards, nil
}

// getNodeInfo returns the node slashed token.
func (c *Component) getNodeBasicInfo(ctx context.Context, address common.Address) (int64, decimal.Decimal, error) {
	var resp GINodeInfoResponse

	err := c.sendRequest(ctx, fmt.Sprintf("/nta/nodes/%s", address), &resp)
	if err != nil {
		return 0, decimal.Decimal{}, fmt.Errorf("failed to get node slashed tokens: %w", err)
	}

	return resp.Data.LastHeartbeat, resp.Data.SlashedTokens, nil
}

// sendRequest sends a request to the global indexer.
func (c *Component) sendRequest(ctx context.Context, path string, result any) error {
	internalURL, err := url.Parse(c.config.Discovery.Server.GlobalIndexerEndpoint)
	if err != nil {
		return err
	}

	internalURL.Path = path

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, internalURL.String(), nil)
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		marshal, _ := json.Marshal(result)

		return fmt.Errorf("unexpected status: %s, response: %s", resp.Status, string(marshal))
	}

	return nil
}

// GetFirstStartTime Get the first start time from redis cache
func GetFirstStartTime(ctx context.Context, redisClient rueidis.Client) (int64, error) {
	if redisClient == nil {
		return 0, fmt.Errorf("redis client is nil")
	}

	command := redisClient.B().Get().Key(buildFirstStartTimeCacheKey()).Build()

	result := redisClient.Do(ctx, command)
	if err := result.Error(); err != nil {
		if errors.Is(err, rueidis.Nil) {
			// Key doesn't exist, return 0 or a default value
			return 0, nil
		}

		return 0, fmt.Errorf("redis result: %w", err)
	}

	firstStartTime, err := result.AsInt64()
	if err != nil {
		return 0, fmt.Errorf("redis result to int64: %w", err)
	}

	return firstStartTime, nil
}

// UpdateFirstStartTime updates the first start time in redis cache
func UpdateFirstStartTime(ctx context.Context, redisClient rueidis.Client, timestamp int64) error {
	if redisClient == nil {
		return fmt.Errorf("redis client is nil")
	}

	command := redisClient.B().Set().Key(buildFirstStartTimeCacheKey()).Value(strconv.FormatInt(timestamp, 10)).Build()

	result := redisClient.Do(ctx, command)
	if err := result.Error(); err != nil {
		return fmt.Errorf("redis result: %w", err)
	}

	return nil
}

// buildFirstStartTimeCacheKey builds the cache key for the first start time
func buildFirstStartTimeCacheKey() string {
	return "node:info:first_start_time"
}
