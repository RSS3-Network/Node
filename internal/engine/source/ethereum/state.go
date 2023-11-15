package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
)

type State struct {
	ChainID     uint64      `json:"chain_id"`
	BlockHash   common.Hash `json:"block_hash"`
	BlockNumber uint64      `json:"block_number"`
}
