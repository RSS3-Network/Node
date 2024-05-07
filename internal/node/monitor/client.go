package monitor

import (
	"context"

	"github.com/rss3-network/node/provider/arweave"
	"github.com/rss3-network/node/provider/ethereum"
)

type Client interface {
	CurrentState(state CheckpointState) uint64
	LatestState(ctx context.Context) (uint64, error)
}

type ethereumClient struct {
	ethereumClient ethereum.Client
}

// make sure client implements monitor.Client
var _ Client = (*ethereumClient)(nil)

func (c *ethereumClient) CurrentState(state CheckpointState) uint64 {
	return state.BlockNumber
}

func (c *ethereumClient) LatestState(ctx context.Context) (uint64, error) {
	latestWorkerState, err := c.ethereumClient.BlockNumber(ctx)
	if err != nil {
		return 0, err
	}

	return uint64(latestWorkerState.Int64()), nil
}

func NewEthereumClient(endpoint string) (Client, error) {
	evmClient, err := ethereum.Dial(context.Background(), endpoint)
	if err != nil {
		return nil, err
	}

	return &ethereumClient{
		ethereumClient: evmClient,
	}, nil
}

type arweaveClient struct {
	arweaveClient arweave.Client
}

// make sure client implements Client
var _ Client = (*arweaveClient)(nil)

func (c *arweaveClient) CurrentState(state CheckpointState) uint64 {
	return state.BlockHeight
}

func (c *arweaveClient) LatestState(ctx context.Context) (uint64, error) {
	latestWorkerState, err := c.arweaveClient.GetBlockHeight(ctx)
	if err != nil {
		return 0, err
	}

	return uint64(latestWorkerState), nil
}

func NewArweaveClient() (Client, error) {
	arClient, err := arweave.NewClient()
	if err != nil {
		return nil, err
	}

	return &arweaveClient{
		arweaveClient: arClient,
	}, nil
}
