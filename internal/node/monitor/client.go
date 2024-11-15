package monitor

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine/protocol/activitypub"
	"github.com/rss3-network/node/provider/arweave"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/farcaster"
	"github.com/rss3-network/node/provider/httpx"
	"github.com/rss3-network/node/provider/near"
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
	httpClient httpx.Client
	relayURLs  []string
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

func (c *activitypubClient) CurrentState(_ CheckpointState) (uint64, uint64) {
	return 0, 0
}

func (c *activitypubClient) TargetState(_ *config.Parameters) (uint64, uint64) {
	return 0, 0
}

// LatestState checks the health of the ActivityPub connection.
// Returns current timestamp if healthy, error otherwise.
func (c *activitypubClient) LatestState(ctx context.Context) (uint64, uint64, error) {
	currentTime := uint64(time.Now().Unix())

	// Check each relay URL
	for _, relayURL := range c.relayURLs {
		resp, err := c.httpClient.Fetch(ctx, relayURL)
		if err != nil {
			zap.L().Error("relay health check failed",
				zap.String("relay_url", relayURL),
				zap.Error(err))

			continue
		}

		// Close response body
		if resp != nil {
			resp.Close()
		}

		// If we got a successful response from any relay, consider the service healthy
		zap.L().Info("relay check succeeded",
			zap.String("relay_url", relayURL))

		return 0, currentTime, nil
	}

	// No relays were accessible
	return 0, 0, fmt.Errorf("no accessible relays found")
}

// NewActivityPubClient returns a new ActivityPub client.
func NewActivityPubClient(network network.Network, param *config.Parameters) (Client, error) {
	// Get relay URLs directly from parameters using NewOption
	option, err := activitypub.NewOption(network, param, true)
	if err != nil {
		return nil, fmt.Errorf("failed to create option: %w", err)
	}

	relayURLList := option.RelayURLList
	httpClient, err := httpx.NewHTTPClient()

	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP client: %w", err)
	}

	zap.L().Info("initializing ActivityPub client",
		zap.Strings("relay_urls", relayURLList))

	return &activitypubClient{
		httpClient: httpClient,
		relayURLs:  relayURLList,
	}, nil
}
