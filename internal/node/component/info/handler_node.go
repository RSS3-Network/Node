package info

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/node/internal/constant"
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

type NodeInfo struct {
	Operator   common.Address `json:"operator"`
	Version    Version        `json:"version"`
	Uptime     int64          `json:"uptime"`
	Parameters string         `json:"parameters"`
}

// GetNodeInfo returns the node information.
func (c *Component) GetNodeInfo(ctx echo.Context) error {
	go c.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, "info")

	zap.L().Debug("get node info", zap.String("request.ip", ctx.Request().RemoteAddr))

	var uptime int64

	version := c.buildVersion()

	evmAddress := common.Address{}

	if !lo.IsEmpty(c.config.Discovery.Operator.EvmAddress) {
		evmAddress = c.config.Discovery.Operator.EvmAddress
	}

	currentEpoch, err := parameter.GetCurrentEpoch(ctx.Request().Context(), c.redisClient)
	if err != nil {
		return fmt.Errorf("failed to get current epoch from cache: %w", err)
	}

	// Get parameters from the network
	params, err := parameter.PullNetworkParamsFromVSL(c.networkParamsCaller, uint64(currentEpoch))
	if err != nil {
		return err
	}

	// get first start time from redis cache and calculate uptime
	firstStartTime, err := GetFirstStartTime(ctx.Request().Context(), c.redisClient)
	if err != nil {
		return fmt.Errorf("failed to get first start time from cache: %w", err)
	}

	if firstStartTime == 0 {
		uptime = 0

		err := UpdateFirstStartTime(ctx.Request().Context(), c.redisClient, time.Now().Unix())
		if err != nil {
			return fmt.Errorf("failed to update first start time: %w", err)
		}
	} else {
		uptime = time.Now().Unix() - firstStartTime
	}

	return ctx.JSON(http.StatusOK, NodeInfoResponse{
		Data: NodeInfo{
			Version:    version,
			Operator:   evmAddress,
			Parameters: params,
			Uptime:     uptime,
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
