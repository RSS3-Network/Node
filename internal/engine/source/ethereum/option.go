package ethereum

import (
	"math/big"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
)

const (
	defaultConcurrentBlockRequests = uint64(8)
	defaultBlockBatchSize          = uint(8)
	defaultReceiptsBatchSize       = uint(200)
	defaultBlockReceiptsBatchSize  = uint(8)
)

type Option struct {
	// BlockStart is the block height on Arweave that the worker should start from.
	BlockStart *big.Int `json:"block_start" mapstructure:"block_start"`
	// BlockTarget is the block height on Arweave that the worker should stop at.
	BlockTarget *big.Int `json:"block_target" mapstructure:"block_target"`
	// ConcurrentBlockRequests is the number of blocks to request concurrently.
	ConcurrentBlockRequests *uint64 `json:"concurrent_block_requests" mapstructure:"concurrent_block_requests"`
	// BlockBatchSize is the number of blocks to fetch in a single batch.
	BlockBatchSize *uint `json:"block_batch_size" mapstructure:"block_batch_size"`
	// ReceiptsBatchSize is the number of receipts to fetch in a single batch.
	ReceiptsBatchSize *uint `json:"receipts_batch_size" mapstructure:"receipts_batch_size"`
	// BlockReceiptsBatchSize is the number of block receipts to fetch in a single batch.
	BlockReceiptsBatchSize *uint `json:"block_receipts_batch_size" mapstructure:"block_receipts_batch_size"`
}

func NewOption(n network.Network, parameters *config.Parameters) (*Option, error) {
	var option Option

	if parameters == nil {
		return &Option{
			BlockStart:              parameter.CurrentNetworkStartBlock[n].Block,
			ConcurrentBlockRequests: lo.ToPtr(defaultConcurrentBlockRequests),
			BlockBatchSize:          lo.ToPtr(defaultBlockBatchSize),
			ReceiptsBatchSize:       lo.ToPtr(defaultReceiptsBatchSize),
			BlockReceiptsBatchSize:  lo.ToPtr(defaultBlockReceiptsBatchSize),
		}, nil
	}

	if err := parameters.Decode(&option); err != nil {
		return nil, err
	}

	// Set default values.
	if option.ConcurrentBlockRequests == nil {
		option.ConcurrentBlockRequests = lo.ToPtr(defaultConcurrentBlockRequests)
	}

	if option.BlockBatchSize == nil {
		option.BlockBatchSize = lo.ToPtr(defaultBlockBatchSize)
	}

	if option.ReceiptsBatchSize == nil {
		option.ReceiptsBatchSize = lo.ToPtr(defaultReceiptsBatchSize)
	}

	if option.BlockReceiptsBatchSize == nil {
		option.BlockReceiptsBatchSize = lo.ToPtr(defaultBlockReceiptsBatchSize)
	}

	if option.BlockStart == nil {
		option.BlockStart = parameter.CurrentNetworkStartBlock[n].Block
	}

	return &option, nil
}
