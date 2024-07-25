package broadcaster

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
	"go.uber.org/zap"
)

func (b *Broadcaster) Register(ctx context.Context) error {
	request := RegisterNodeRequest{
		Address:   b.config.Discovery.Operator.EvmAddress,
		Signature: b.config.Discovery.Operator.Signature,
		Endpoint:  b.config.Discovery.Server.Endpoint,
		Stream:    b.config.Stream,
		Config:    b.config.Component,
		Type:      b.config.Type,
	}

	var response any

	err := b.sendRequest(ctx, "/nodes/register", nil, request, &response)

	zap.L().Info("registered node", zap.Any("request", request), zap.Any("response", response))

	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}

	return nil
}

func (b *Broadcaster) Heartbeat(ctx context.Context) error {
	request := NodeHeartbeatRequest{
		Address:   b.config.Discovery.Operator.EvmAddress,
		Endpoint:  b.config.Discovery.Server.Endpoint,
		Signature: b.config.Discovery.Operator.Signature,
		Timestamp: time.Now().Unix(),
	}

	var response any

	err := b.sendRequest(ctx, "/nodes/heartbeat", nil, request, &response)

	zap.L().Info("sending heartbeat", zap.Any("request", request), zap.Any("response", response))

	if err != nil {
		zap.L().Error("failed to send heartbeat", zap.Error(err))

		return fmt.Errorf("send heartbeat: %w", err)
	}

	// record heartbeat
	err = UpdateHeartbeatTime(ctx, b.redisClient, request.Timestamp)
	if err != nil {
		return fmt.Errorf("update last heartbeat time: %w", err)
	}

	return nil
}

func (b *Broadcaster) sendRequest(ctx context.Context, path string, values url.Values, body any, result any) error {
	internalURL, err := url.Parse(b.config.Discovery.Server.GlobalIndexerEndpoint)
	if err != nil {
		return err
	}

	internalURL.Path = path
	internalURL.RawQuery = values.Encode()

	requestBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, internalURL.String(), bytes.NewReader(requestBody))
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := b.httpClient.Do(req)
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

type RegisterNodeRequest struct {
	Address   common.Address    `json:"address" validate:"required"`
	Signature string            `json:"signature" validate:"required"`
	Endpoint  string            `json:"endpoint" validate:"required"`
	Stream    *config.Stream    `json:"stream,omitempty"`
	Config    *config.Component `json:"config,omitempty"`
	Type      string            `json:"type" validate:"required"`
}

type NodeHeartbeatRequest struct {
	Address   common.Address `json:"address" validate:"required"`
	Signature string         `json:"signature" validate:"required"`
	Endpoint  string         `json:"endpoint" validate:"required"`
	Timestamp int64          `json:"timestamp" validate:"required"`
}

// GetHeartbeatTime Get the last heartbeat time from redis cache
func GetHeartbeatTime(ctx context.Context, redisClient rueidis.Client) (int64, error) {
	if redisClient == nil {
		return 0, fmt.Errorf("redis client is nil")
	}

	command := redisClient.B().Get().Key(buildLastHeatbeatTimeCacheKey()).Build()

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
	lastHeatbeatTime, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse int: %w", err)
	}

	return lastHeatbeatTime, nil
}

// UpdateHeartbeatTime updates the first start time in redis cache
func UpdateHeartbeatTime(ctx context.Context, redisClient rueidis.Client, timestamp int64) error {
	if redisClient == nil {
		return fmt.Errorf("redis client is nil")
	}

	command := redisClient.B().Set().Key(buildLastHeatbeatTimeCacheKey()).Value(strconv.FormatInt(timestamp, 10)).Build()

	result := redisClient.Do(ctx, command)
	if err := result.Error(); err != nil {
		return fmt.Errorf("redis result: %w", err)
	}

	return nil
}

// buildLastHeatbeatTimeCacheKey builds the cache key for last heartbeat time
func buildLastHeatbeatTimeCacheKey() string {
	return "node:info:last_heartbeat_time"
}
