package monitor

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/node/component/rss"
	"github.com/rss3-network/node/provider/arweave"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/farcaster"
	"github.com/rss3-network/node/provider/httpx"
)

type Client interface {
	// CurrentState returns the current block number (ethereum), height (arweave) or event id (farcaster) of the client from Checkpoints table in database.
	CurrentState(state CheckpointState) uint64
	// TargetState checks if the target block number/height is set in the parameters.
	TargetState(param *config.Parameters) uint64
	// LatestState returns the latest block number (ethereum), height (arweave) or event id (farcaster) of the client from network rpc/api.
	LatestState(ctx context.Context) (uint64, error)
}

// ethereumClient is a client implementation for ethereum.
type ethereumClient struct {
	ethereumClient ethereum.Client
}

// make sure client implements Client
var _ Client = (*ethereumClient)(nil)

func (c *ethereumClient) CurrentState(state CheckpointState) uint64 {
	return state.BlockNumber
}

func (c *ethereumClient) TargetState(param *config.Parameters) uint64 {
	return getTargetBlockFromParam(param)
}

func (c *ethereumClient) LatestState(ctx context.Context) (uint64, error) {
	latestWorkerState, err := c.ethereumClient.BlockNumber(ctx)
	if err != nil {
		return 0, err
	}

	return uint64(latestWorkerState.Int64()), nil
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

// arweaveClient is a client implementation for arweave.
type arweaveClient struct {
	arweaveClient arweave.Client
}

// make sure client implements Client
var _ Client = (*arweaveClient)(nil)

func (c *arweaveClient) CurrentState(state CheckpointState) uint64 {
	return state.BlockHeight
}

func (c *arweaveClient) TargetState(param *config.Parameters) uint64 {
	return getTargetBlockFromParam(param)
}

func (c *arweaveClient) LatestState(ctx context.Context) (uint64, error) {
	latestWorkerState, err := c.arweaveClient.GetBlockHeight(ctx)
	if err != nil {
		return 0, err
	}

	return uint64(latestWorkerState), nil
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

func (c *farcasterClient) CurrentState(state CheckpointState) uint64 {
	if state.EventID != 0 {
		return farcaster.ConvertEventIDToTimestampMilli(state.EventID)
	}

	return 0
}

func (c *farcasterClient) TargetState(_ *config.Parameters) uint64 {
	return 0
}

func (c *farcasterClient) LatestState(_ context.Context) (uint64, error) {
	return uint64(time.Now().UnixMilli()), nil
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

type rssClient struct {
	httpClient httpx.Client
	url        string
}

// make sure client implements Client
var _ Client = (*rssClient)(nil)

func (c *rssClient) CurrentState(_ CheckpointState) uint64 {
	return 0
}

func (c *rssClient) TargetState(_ *config.Parameters) uint64 {
	return 0
}

func (c *rssClient) LatestState(ctx context.Context) (uint64, error) {
	_, err := c.httpClient.Fetch(ctx, c.url)

	if err != nil {
		return 0, err
	}

	return 0, nil
}

// NewRssClient returns a new rss client.
func NewRssClient(endpoint config.Endpoint, param *config.Parameters) (Client, error) {
	base, err := url.Parse(endpoint.URL)
	if err != nil {
		return nil, fmt.Errorf("parse RSSHub endpoint: %w", err)
	}

	// used for health checks
	base.Path = path.Join(base.Path, "abc")

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
