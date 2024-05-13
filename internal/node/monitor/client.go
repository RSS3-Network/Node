package monitor

import (
	"context"
	"math/big"
	"time"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/provider/arweave"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/farcaster"
)

type Client interface {
	// CurrentState returns the current block number (ethereum), height (arweave) or event id (farcaster) of the client from Checkpoints table in database.
	CurrentState(state CheckpointState) uint64
	IsTargetParamSet(param *config.Parameters) bool
	// LatestState returns the latest block number (ethereum), height (arweave) or event id (farcaster) of the client from network rpc/api.
	LatestState(ctx context.Context) (uint64, error)
}

type TargetParam struct {
	BlockHeightTarget *big.Int `json:"block_height_target"`
	BlockNumberTarget *big.Int `json:"block_number_target"`
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

func (c *ethereumClient) IsTargetParamSet(param *config.Parameters) bool {
	// Check if the target block number is set.
	if param == nil {
		return false
	}

	targetBlock, exists := (*param)["block_number_target"]
	if exists && targetBlock != nil {
		return true
	}

	return false
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

func (c *arweaveClient) IsTargetParamSet(param *config.Parameters) bool {
	// Check if the target block number is set.
	if param == nil {
		return false
	}

	targetBlock, exists := (*param)["block_height_target"]
	if exists && targetBlock != nil {
		return true
	}

	return false
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

func (c *farcasterClient) IsTargetParamSet(_ *config.Parameters) bool {
	return false
}

func (c *farcasterClient) LatestState(_ context.Context) (uint64, error) {
	return uint64(time.Now().UnixMilli()), nil
}

// NewFarcasterClient returns a new farcaster client.
func NewFarcasterClient() (Client, error) {
	return &farcasterClient{}, nil
}
