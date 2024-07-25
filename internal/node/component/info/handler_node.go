package info

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
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
	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
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
		ID                    uint64 `json:"id"`
		TotalOperationRewards string `json:"total_operation_rewards"`
		TotalStakingRewards   string `json:"total_staking_rewards"`
		TotalRequestCounts    string `json:"total_request_counts"`
	} `json:"data"`
}

type GINodeInfoResponse struct {
	Data struct {
		LastHeartbeat int64  `json:"last_heartbeat"`
		SlashedTokens string `json:"slashed_tokens"`
	} `json:"data"`
}

type NodeInfo struct {
	Operator   common.Address `json:"operator"`
	Version    Version        `json:"version"`
	Uptime     int64          `json:"uptime"`
	Parameters string         `json:"parameters"`
	Coverage   float64        `json:"coverage"`
	Records    Record         `json:"records"`
}

type Record struct {
	LastHeartbeat  int64    `json:"last_heartbeat"`
	RecentRequests []string `json:"recent_requests"`
	RecentRewards  []Reward `json:"recent_rewards"`
	SlashedTokens  string   `json:"slashed_tokens"`
}

type Reward struct {
	Epoch            uint64 `json:"epoch"`
	OperationRewards string `json:"operation_rewards"`
	StakingRewards   string `json:"staking_rewards"`
	RequestCounts    string `json:"request_counts"`
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
				RecentRequests: decentralized.RecentRequests,
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
func (c *Component) getNodeWorkerCoverage() float64 {
	// should be divided by the total workers to get worker coverage
	currentWorkerCount := len(c.config.Component.Decentralized) + len(c.config.Component.RSS) + len(c.config.Component.Federated)
	totalWorkerCount := calculateTotalWorkers(NetworkToWorkersMap)

	// Calculate worker coverage
	workerCoverage := 0.0
	if totalWorkerCount > 0 {
		workerCoverage = float64(currentWorkerCount) / float64(totalWorkerCount)
		workerCoverage = math.Round(workerCoverage*10000) / 10000
	}

	return workerCoverage
}

// calculateTotalWorkers calculates the total number of workers in the network.
func calculateTotalWorkers(networkToWorkersMap map[network.Network][]worker.Worker) int {
	totalWorkerCount := 0
	for _, workers := range networkToWorkersMap {
		totalWorkerCount += len(workers)
	}

	return totalWorkerCount
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
	var Resp GIRewardsResponse

	err := c.sendRequest(ctx, fmt.Sprintf("/nta/epochs/%s/rewards", address.Hex()), &Resp)
	if err != nil {
		return nil, fmt.Errorf("failed to get node rewards: %w", err)
	}

	rewards := make([]Reward, 0, len(Resp.Data))

	for _, data := range Resp.Data {
		reward := Reward{
			Epoch:            data.ID,
			OperationRewards: data.TotalOperationRewards,
			StakingRewards:   data.TotalStakingRewards,
			RequestCounts:    data.TotalRequestCounts,
		}
		rewards = append(rewards, reward)
	}

	return rewards, nil
}

// getNodeInfo returns the node slashed token.
func (c *Component) getNodeBasicInfo(ctx context.Context, address common.Address) (int64, string, error) {
	var Resp GINodeInfoResponse

	err := c.sendRequest(ctx, fmt.Sprintf("/nta/nodes/%s", address.Hex()), &Resp)
	if err != nil {
		return 0, "", fmt.Errorf("failed to get node slashed tokens: %w", err)
	}

	return Resp.Data.LastHeartbeat, Resp.Data.SlashedTokens, nil
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

	// Retrieve the result as a string
	timestampStr, err := result.ToString()
	if err != nil {
		return 0, fmt.Errorf("redis result to string: %w", err)
	}

	// Convert the string to an uint64
	firstStartTime, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse int: %w", err)
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
