package ethereum

import (
	"math/big"

	"github.com/rss3-network/node/config"
	"github.com/samber/lo"
)

const (
	defaultRPCThreadBlocks       = uint(8)
	defaultRPCBatchBlocks        = uint(8)
	defaultRPCBatchReceipts      = uint(200)
	defaultRPCBatchBlockReceipts = uint(8)
)

type Option struct {
	BlockNumberStart  *big.Int `json:"block_number_start" mapstructure:"block_number_start"`
	BlockNumberTarget *big.Int `json:"block_number_target" mapstructure:"block_number_target"`

	// RPCThreadBlocks is the number of concurrent RPC requests associated with the blocks.
	RPCThreadBlocks       *uint `json:"rpc_thread_blocks" mapstructure:"rpc_thread_blocks"`
	RPCBatchBlocks        *uint `json:"rpc_batch_blocks" mapstructure:"rpc_batch_blocks"`
	RPCBatchReceipts      *uint `json:"rpc_batch_receipts" mapstructure:"rpc_batch_blocks"`
	RPCBatchBlockReceipts *uint `json:"rpc_batch_block_receipts" mapstructure:"rpc_batch_block_receipts"`
}

func NewOption(options *config.Options) (*Option, error) {
	var instance Option

	if options == nil {
		return &instance, nil
	}

	if err := options.Decode(&instance); err != nil {
		return nil, err
	}

	// Set default values.
	if instance.RPCThreadBlocks == nil {
		instance.RPCThreadBlocks = lo.ToPtr(defaultRPCThreadBlocks)
	}

	if instance.RPCBatchBlocks == nil {
		instance.RPCBatchBlocks = lo.ToPtr(defaultRPCBatchBlocks)
	}

	if instance.RPCBatchReceipts == nil {
		instance.RPCBatchReceipts = lo.ToPtr(defaultRPCBatchReceipts)
	}

	if instance.RPCBatchBlockReceipts == nil {
		instance.RPCBatchBlockReceipts = lo.ToPtr(defaultRPCBatchBlockReceipts)
	}

	return &instance, nil
}
