package ethereum

import (
	"math/big"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
)

const (
	defaultConcurrentBlockRequests = uint(8)
	defaultBlockBatchSize          = uint(8)
	defaultReceiptsBatchSize       = uint(200)
	defaultBlockReceiptsBatchSize  = uint(8)
)

type Option struct {
	// BlockStart is the starting block number on evm chain.
	BlockStart *big.Int `json:"block_start" mapstructure:"block_start"`
	// BlockTarget is the target block number on evm chain.
	BlockTarget *big.Int `json:"block_target" mapstructure:"block_target"`
	// ConcurrentBlockRequests is the number of concurrent RPC requests associated with the blocks.
	ConcurrentBlockRequests *uint `json:"concurrent_block_requests" mapstructure:"concurrent_block_requests"`
	// BlockBatchSize is the number of blocks to fetch in a single batch.
	BlockBatchSize *uint `json:"block_batch_size" mapstructure:"block_batch_size"`
	// ReceiptsBatchSize is the number of receipts to fetch in a single batch.
	ReceiptsBatchSize *uint `json:"receipts_batch_size" mapstructure:"receipts_batch_size"`
	// BlockReceiptsBatchSize is the number of block receipts to fetch in a single batch.
	BlockReceiptsBatchSize *uint `json:"block_receipts_batch_size" mapstructure:"block_receipts_batch_size"`
}

func NewOption(n network.Network, options *config.Parameters) (*Option, error) {
	var instance Option

	if options == nil {
		return &Option{
			BlockStart:              parameter.NetworkStartBlock[n],
			ConcurrentBlockRequests: lo.ToPtr(defaultConcurrentBlockRequests),
			BlockBatchSize:          lo.ToPtr(defaultBlockBatchSize),
			ReceiptsBatchSize:       lo.ToPtr(defaultReceiptsBatchSize),
			BlockReceiptsBatchSize:  lo.ToPtr(defaultBlockReceiptsBatchSize),
		}, nil
	}

	if err := options.Decode(&instance); err != nil {
		return nil, err
	}

	// Set default values.
	if instance.ConcurrentBlockRequests == nil {
		instance.ConcurrentBlockRequests = lo.ToPtr(defaultConcurrentBlockRequests)
	}

	if instance.BlockBatchSize == nil {
		instance.BlockBatchSize = lo.ToPtr(defaultBlockBatchSize)
	}

	if instance.ReceiptsBatchSize == nil {
		instance.ReceiptsBatchSize = lo.ToPtr(defaultReceiptsBatchSize)
	}

	if instance.BlockReceiptsBatchSize == nil {
		instance.BlockReceiptsBatchSize = lo.ToPtr(defaultBlockReceiptsBatchSize)
	}

	if instance.BlockStart == nil {
		instance.BlockStart = parameter.NetworkStartBlock[n]
	}

	return &instance, nil
}
