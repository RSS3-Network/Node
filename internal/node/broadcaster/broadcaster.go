package broadcaster

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"go.uber.org/zap"
)

func (b *Broadcaster) Register(ctx context.Context) error {
	request := RegisterNodeRequest{
		Address:   b.config.Discovery.Maintainer.EvmAddress,
		Signature: b.config.Discovery.Maintainer.Signature,
		Endpoint:  b.config.Discovery.Server.Endpoint,
		Stream:    b.config.Stream,
		Config:    b.config.Node,
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
		Address:   b.config.Discovery.Maintainer.EvmAddress,
		Endpoint:  b.config.Discovery.Server.Endpoint,
		Signature: b.config.Discovery.Maintainer.Signature,
		Timestamp: time.Now().Unix(),
	}

	var response any

	err := b.sendRequest(ctx, "/nodes/heartbeat", nil, request, &response)

	zap.L().Info("sending heartbeat", zap.Any("request", request), zap.Any("response", response))

	if err != nil {
		zap.L().Error("failed to send heartbeat", zap.Error(err))

		return fmt.Errorf("send heartbeat: %w", err)
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
		return fmt.Errorf("unexpected status: %s, response: %s", resp.Status, result)
	}

	return nil
}

type RegisterNodeRequest struct {
	Address   common.Address `json:"address" validate:"required"`
	Signature string         `json:"signature" validate:"required"`
	Endpoint  string         `json:"endpoint" validate:"required"`
	Stream    *config.Stream `json:"stream,omitempty"`
	Config    *config.Node   `json:"config,omitempty"`
}

type NodeHeartbeatRequest struct {
	Address   common.Address `json:"address" validate:"required"`
	Signature string         `json:"signature" validate:"required"`
	Endpoint  string         `json:"endpoint" validate:"required"`
	Timestamp int64          `json:"timestamp" validate:"required"`
}
