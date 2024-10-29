package monitor

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine/protocol/activitypub"
	"github.com/rss3-network/node/internal/node/component/rss"
	"github.com/rss3-network/node/provider/activitypub/mastodon"
	"github.com/rss3-network/node/provider/arweave"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/farcaster"
	"github.com/rss3-network/node/provider/httpx"
	"github.com/rss3-network/node/provider/near"
	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/node/schema/worker/federated"
	"github.com/rss3-network/protocol-go/schema/network"
	"go.uber.org/zap"
)

type Client interface {
	// CurrentState returns the current block number (ethereum), height (arweave) or event id (farcaster) of the client from Checkpoints table in database.
	CurrentState(state CheckpointState) (uint64, uint64)
	// TargetState checks if the target block number/height is set in the parameters.
	TargetState(param *config.Parameters) (uint64, uint64)
	// LatestState returns the latest block number (ethereum), height (arweave) or event id (farcaster) or err (rss) of the client from network rpc/api.
	LatestState(ctx context.Context) (uint64, uint64, error)
}

// activitypubClient is a client implementation for ActivityPub.
type activitypubClient struct {
	activitypubClient mastodon.Client // ToDo: (Note) Currently use Matodon Client for all ActivityPub Source.
}

// set a default client
var _ Client = (*activitypubClient)(nil)

// ethereumClient is a client implementation for ethereum.
type ethereumClient struct {
	ethereumClient ethereum.Client
}

// make sure client implements Client
var _ Client = (*ethereumClient)(nil)

func (c *ethereumClient) CurrentState(state CheckpointState) (uint64, uint64) {
	return state.BlockNumber, 0
}

func (c *ethereumClient) TargetState(param *config.Parameters) (uint64, uint64) {
	return getTargetBlockFromParam(param), 0
}

func (c *ethereumClient) LatestState(ctx context.Context) (uint64, uint64, error) {
	latestWorkerState, err := c.ethereumClient.BlockNumber(ctx)
	if err != nil {
		return 0, 0, err
	}

	return uint64(latestWorkerState.Int64()), 0, nil
}

// NewEthereumClient returns a new ethereum client.
func NewEthereumClient(endpoint config.Endpoint) (Client, error) {
	evmClient, err := ethereum.Dial(context.Background(), endpoint.URL, endpoint.BuildEthereumOptions()...)
	if err != nil {
		return nil, err
	}

	return &ethereumClient{
		ethereumClient: evmClient,
	}, nil
}

type nearClient struct {
	client near.Client
}

// make sure client implements Client
var _ Client = (*nearClient)(nil)

func (c *nearClient) CurrentState(state CheckpointState) (uint64, uint64) {
	return state.BlockHeight, 0
}

func (c *nearClient) TargetState(param *config.Parameters) (uint64, uint64) {
	return getTargetBlockFromParam(param), 0
}

func (c *nearClient) LatestState(ctx context.Context) (uint64, uint64, error) {
	latestWorkerState, err := c.client.GetBlockHeight(ctx)
	if err != nil {
		return 0, 0, err
	}

	return uint64(latestWorkerState), 0, nil
}

func NewNearClient(endpoint config.Endpoint) (Client, error) {
	client, err := near.Dial(context.Background(), endpoint.URL)
	if err != nil {
		return nil, err
	}

	return &nearClient{
		client: client,
	}, nil
}

// arweaveClient is a client implementation for arweave.
type arweaveClient struct {
	arweaveClient arweave.Client
}

// make sure client implements Client
var _ Client = (*arweaveClient)(nil)

func (c *arweaveClient) CurrentState(state CheckpointState) (uint64, uint64) {
	return state.BlockHeight, state.BlockTimestamp
}

func (c *arweaveClient) TargetState(param *config.Parameters) (uint64, uint64) {
	return getTargetBlockFromParam(param), 0
}

func (c *arweaveClient) LatestState(ctx context.Context) (uint64, uint64, error) {
	latestWorkerState, err := c.arweaveClient.GetBlockHeight(ctx)
	if err != nil {
		return 0, 0, err
	}

	return uint64(latestWorkerState), uint64(time.Now().UnixMilli()), nil
}

// NewArweaveClient returns a new arweave client.
func NewArweaveClient() (Client, error) {
	arClient, err := arweave.NewClient()
	if err != nil {
		return nil, err
	}

	return &arweaveClient{
		arweaveClient: arClient,
	}, nil
}

// farcasterClient is a client implementation for farcaster.
type farcasterClient struct{}

// make sure client implements Client
var _ Client = (*farcasterClient)(nil)

func (c *farcasterClient) CurrentState(state CheckpointState) (uint64, uint64) {
	if state.EventID != 0 {
		return farcaster.ConvertEventIDToTimestampMilli(state.EventID), 0
	}

	return 0, 0
}

func (c *farcasterClient) TargetState(_ *config.Parameters) (uint64, uint64) {
	return 0, 0
}

func (c *farcasterClient) LatestState(_ context.Context) (uint64, uint64, error) {
	return uint64(time.Now().UnixMilli()), 0, nil
}

// getTargetBlockFromParam returns the target block number/height from the parameters.
func getTargetBlockFromParam(param *config.Parameters) uint64 {
	if param == nil {
		return 0
	}

	targetBlock, exists := (*param)["block_target"]
	if !exists || targetBlock == nil {
		return 0
	}

	targetBlockUint, err := convertToUint64(targetBlock)
	if err != nil {
		return 0
	}

	return targetBlockUint
}

// convertToUint64 a helper func which converts the value to uint64.
func convertToUint64(value interface{}) (uint64, error) {
	switch v := value.(type) {
	case string:
		return strconv.ParseUint(v, 10, 64)
	case int:
		return uint64(v), nil
	case int64:
		return uint64(v), nil
	case uint:
		return uint64(v), nil
	case uint64:
		return v, nil
	default:
		return 0, fmt.Errorf("unsupported type: %T", v)
	}
}

// NewFarcasterClient returns a new farcaster client.
func NewFarcasterClient() (Client, error) {
	return &farcasterClient{}, nil
}

// rssClient is a client implementation for rss.
type rssClient struct {
	httpClient httpx.Client
	url        string
}

// make sure client implements Client
var _ Client = (*rssClient)(nil)

func (c *rssClient) CurrentState(_ CheckpointState) (uint64, uint64) {
	return 0, 0
}

func (c *rssClient) TargetState(_ *config.Parameters) (uint64, uint64) {
	return 0, 0
}

// LatestState requests a URL to check if it can be accessed correctly.
func (c *rssClient) LatestState(ctx context.Context) (uint64, uint64, error) {
	_, err := c.httpClient.Fetch(ctx, c.url)

	if err != nil {
		return 0, 0, err
	}

	return 0, 0, nil
}

// NewRssClient returns a new rss client.
func NewRssClient(endpoint string, param *config.Parameters) (Client, error) {
	base, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("parse RSSHub endpoint: %w", err)
	}

	// used for health checks
	base.Path = path.Join(base.Path, "healthz")

	option, err := rss.NewOption(param)
	if err != nil {
		return nil, fmt.Errorf("parse config parameters: %w", err)
	}

	if option.Authentication.AccessKey != "" {
		query := base.Query()
		query.Set("key", option.Authentication.AccessKey)
		base.RawQuery = query.Encode()
	}

	httpClient, err := httpx.NewHTTPClient()
	if err != nil {
		return nil, err
	}

	return &rssClient{
		httpClient: httpClient,
		url:        base.String(),
	}, nil
}

func (c *activitypubClient) CurrentState(_ CheckpointState) (uint64, uint64) {
	return 0, 0
}

func (c *activitypubClient) TargetState(_ *config.Parameters) (uint64, uint64) {
	return 0, 0
}

// LatestState checks the health of the ActivityPub connection.
// Returns current timestamp if healthy, error otherwise.
func (c *activitypubClient) LatestState(ctx context.Context) (uint64, uint64, error) {
	if c.activitypubClient == nil {
		return 0, 0, fmt.Errorf("client not initialized")
	}

	msgChan, err := c.activitypubClient.GetMessageChan()
	if err != nil {
		return 0, 0, fmt.Errorf("get message channel: %w", err)
	}

	// Create a timeout context
	healthCheckTimeout := 10 * time.Second
	checkCtx, cancel := context.WithTimeout(ctx, healthCheckTimeout)

	defer cancel()

	// Check server readiness first
	select {
	case <-c.activitypubClient.GetReadyChan():
		zap.L().Debug("server is ready")
	case <-checkCtx.Done():
		return 0, 0, fmt.Errorf("timeout waiting for server readiness")
	}

	// Check connection health and message flow
	select {
	case <-checkCtx.Done():
		if ctx.Err() != nil {
			return 0, 0, fmt.Errorf("context cancelled: %w", ctx.Err())
		}

		return 0, uint64(time.Now().Unix()), nil

	case msg := <-msgChan:
		currentTime := uint64(time.Now().Unix())

		// Re-send the message back to the channel for other listeners
		c.activitypubClient.SendMessage(msg)

		zap.L().Info("received message during health check",
			zap.String("message", msg),
			zap.Uint64("timestamp", currentTime))

		return 0, currentTime, nil
	}
}

// NewActivityPubClient returns a new ActivityPub client.
func NewActivityPubClient(endpoint config.Endpoint, worker worker.Worker) (Client, error) {
	if endpoint.URL == "" {
		return nil, fmt.Errorf("endpoint URL is required")
	}

	// Get relay URLs directly from parameters using NewOption
	option, err := activitypub.NewOption(network.Mastodon, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create option: %w", err)
	}

	relayURLList := option.RelayURLList
	port := mastodon.DefaultMonitorServerPort // Default port for Monitor client

	zap.L().Info("Using relay URL list and port in NewActivityPubClient", zap.Strings("relayURLList", relayURLList), zap.Int64("port", int64(port)))

	// Create a context with cancellation
	ctx := context.Background()
	errorChan := make(chan error, 1)

	workerType := worker.Name()
	switch workerType {
	case federated.Core.String():
		mastodonClient, err := mastodon.NewClient(ctx, endpoint.URL, relayURLList, int64(port), errorChan)
		if err != nil {
			return nil, fmt.Errorf("create mastodon client: %w", err)
		}

		// Monitor error channel in background
		go func() {
			for err := range errorChan {
				if err != nil {
					zap.L().Error("mastodon client error",
						zap.Error(err))
				}
			}
		}()

		return &activitypubClient{
			activitypubClient: mastodonClient,
		}, nil

	default:
		return nil, fmt.Errorf("unsupported worker type: %s", workerType)
	}
}
