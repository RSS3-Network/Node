package ethereum

import (
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/samber/lo"
)

const (
	// Optimism
	// https://github.com/Inphi/optimism-rosetta/blob/d3c37e359fa4262a2a53014c424a8ae53efaf3d0/optimism/client_bedrock.go#L138
	TransactionTypeOptimismDeposit = 0x7E
)

//go:generate go run --mod=mod github.com/fjl/gencodec --type Header --field-override headerMarshal --out type_header.go
type Header struct {
	Hash         common.Hash    `json:"hash"`
	ParentHash   common.Hash    `json:"parentHash"`
	UncleHash    common.Hash    `json:"sha3Uncles"`
	Coinbase     common.Address `json:"miner"`
	Number       *big.Int       `json:"number"`
	GasLimit     uint64         `json:"gasLimit"`
	GasUsed      uint64         `json:"gasUsed"`
	Timestamp    uint64         `json:"timestamp"`
	BaseFee      *big.Int       `json:"baseFeePerGas"`
	Transactions []common.Hash  `json:"transactions"`
}

type headerMarshal struct {
	Number    *hexutil.Big
	GasLimit  hexutil.Uint64
	GasUsed   hexutil.Uint64
	Timestamp hexutil.Uint64
	BaseFee   *hexutil.Big
}

//go:generate go run --mod=mod github.com/fjl/gencodec --type Block --field-override blockMarshal --out type_block.go
type Block struct {
	Hash         common.Hash    `json:"hash"`
	ParentHash   common.Hash    `json:"parentHash"`
	UncleHash    common.Hash    `json:"sha3Uncles"`
	Coinbase     common.Address `json:"miner"`
	Number       *big.Int       `json:"number"`
	GasLimit     uint64         `json:"gasLimit"`
	GasUsed      uint64         `json:"gasUsed"`
	Timestamp    uint64         `json:"timestamp"`
	BaseFee      *big.Int       `json:"baseFeePerGas"`
	Transactions []*Transaction `json:"transactions"`
}

type blockMarshal struct {
	Number    *hexutil.Big
	GasLimit  hexutil.Uint64
	GasUsed   hexutil.Uint64
	Timestamp hexutil.Uint64
	BaseFee   *hexutil.Big
}

func (b Block) Header() *Header {
	header := Header{
		Hash:       b.Hash,
		ParentHash: b.ParentHash,
		UncleHash:  b.UncleHash,
		Coinbase:   b.Coinbase,
		Number:     b.Number,
		GasLimit:   b.GasLimit,
		GasUsed:    b.GasUsed,
		Timestamp:  b.Timestamp,
		BaseFee:    b.BaseFee,
		Transactions: lo.Map(b.Transactions, func(transaction *Transaction, _ int) common.Hash {
			return transaction.Hash
		}),
	}

	return &header
}

//go:generate go run --mod=mod github.com/fjl/gencodec --type Transaction --field-override transactionMarshal --out type_transaction.go
type Transaction struct {
	BlockHash   common.Hash     `json:"blockHash"`
	BlockNumber *big.Int        `json:"blockNumber"`
	From        common.Address  `json:"from"`
	Gas         uint64          `json:"gas"`
	GasPrice    *big.Int        `json:"gasPrice"`
	GasTipCap   *big.Int        `json:"maxPriorityFeePerGas"`
	Hash        common.Hash     `json:"hash"`
	Input       []byte          `json:"input"`
	To          *common.Address `json:"to"`
	Index       uint            `json:"index"`
	Value       *big.Int        `json:"value"`
	Type        uint64          `json:"type"`
	ChainID     *big.Int        `json:"chainId"`
}

type transactionMarshal struct {
	BlockNumber *hexutil.Big
	Gas         hexutil.Uint64
	GasPrice    *hexutil.Big
	GasTipCap   *hexutil.Big
	Input       hexutil.Bytes
	Index       hexutil.Uint
	Value       *hexutil.Big
	Type        hexutil.Uint64
	ChainID     *hexutil.Big
}

//go:generate go run --mod=mod github.com/fjl/gencodec --type Receipt --field-override receiptMarshal --out type_receipt.go
type Receipt struct {
	BlockHash         common.Hash     `json:"blockHash"`
	BlockNumber       *big.Int        `json:"blockNumber"`
	ContractAddress   *common.Address `json:"contractAddress"`
	CumulativeGasUsed uint64          `json:"cumulativeGasUsed"`
	EffectiveGasPrice *big.Int        `json:"effectiveGasPrice"`
	GasUsed           uint64          `json:"gasUsed"`

	// Optimism
	L1GasPrice *big.Int   `json:"l1GasPrice,omitempty"`
	L1GasUsed  *big.Int   `json:"l1GasUsed,omitempty"`
	L1Fee      *big.Int   `json:"l1Fee,omitempty"`
	FeeScalar  *big.Float `json:"l1FeeScalar,omitempty"`

	Logs             []*Log      `json:"logs"`
	Status           uint64      `json:"status"`
	TransactionHash  common.Hash `json:"transactionHash"`
	TransactionIndex uint        `json:"transactionIndex"`
}

type receiptMarshal struct {
	BlockNumber       *hexutil.Big
	CumulativeGasUsed hexutil.Uint64
	EffectiveGasPrice *hexutil.Big
	GasUsed           hexutil.Uint64
	Status            hexutil.Uint64
	TransactionIndex  hexutil.Uint

	// Optimism
	L1GasPrice *hexutil.Big
	L1GasUsed  *hexutil.Big
	L1Fee      *hexutil.Big
	FeeScalar  *big.Float
}

//go:generate go run --mod=mod github.com/fjl/gencodec --type Log --field-override logMarshal --out type_log.go
type Log struct {
	Address          common.Address `json:"address"`
	Topics           []common.Hash  `json:"topics"`
	Data             []byte         `json:"data"`
	BlockNumber      *big.Int       `json:"blockNumber"`
	TransactionHash  common.Hash    `json:"transactionHash"`
	TransactionIndex uint           `json:"transactionIndex"`
	BlockHash        common.Hash    `json:"blockHash"`
	Index            uint           `json:"logIndex"`
	Removed          bool           `json:"removed"`
}

func (l Log) Export() types.Log {
	return types.Log{
		Address:     l.Address,
		Topics:      l.Topics,
		Data:        l.Data,
		BlockNumber: l.BlockNumber.Uint64(),
		TxHash:      l.TransactionHash,
		TxIndex:     l.TransactionIndex,
		BlockHash:   l.BlockHash,
		Index:       l.Index,
		Removed:     l.Removed,
	}
}

type logMarshal struct {
	Data             hexutil.Bytes
	BlockNumber      *hexutil.Big
	TransactionIndex hexutil.Uint
	Index            hexutil.Uint
}

//go:generate go run --mod=mod github.com/fjl/gencodec --type TransactionCall --field-override transactionCallMarshal --out type_transaction_call.go
type TransactionCall struct {
	From     common.Address  `json:"from"`
	To       *common.Address `json:"to,omitempty"`
	Gas      uint64          `json:"gas,omitempty"`
	GasPrice *big.Int        `json:"gasPrice,omitempty"`
	Value    *big.Int        `json:"value,omitempty"`
	Data     []byte          `json:"data,omitempty"`
}

type transactionCallMarshal struct {
	Gas      hexutil.Uint64
	GasPrice *hexutil.Big
	Value    *hexutil.Big
	Data     hexutil.Bytes
}

func formatBlockNumber(blockNumber *big.Int) string {
	switch {
	case blockNumber == nil:
		return "latest"
	case blockNumber.Int64() == -1:
		return "pending"
	default:
		return hexutil.EncodeBig(blockNumber)
	}
}

func formatTransactionCall(callMessage ethereum.CallMsg) TransactionCall {
	return TransactionCall{
		From:     callMessage.From,
		To:       callMessage.To,
		Gas:      callMessage.Gas,
		GasPrice: callMessage.GasPrice,
		Value:    callMessage.Value,
		Data:     callMessage.Data,
	}
}

type Filter struct {
	BlockHash *common.Hash     `json:"blockHash"`
	FromBlock *big.Int         `json:"fromBlock"`
	ToBlock   *big.Int         `json:"toBlock"`
	Addresses []common.Address `json:"addresses"`
	Topics    [][]common.Hash  `json:"topics"`
}

func formatFilter(filter Filter) map[string]any {
	result := map[string]interface{}{
		"address": filter.Addresses,
		"topics":  filter.Topics,
	}

	if filter.BlockHash == nil {
		if filter.FromBlock == nil {
			result["fromBlock"] = "0x0"
		} else {
			result["fromBlock"] = formatBlockNumber(filter.FromBlock)
		}

		result["toBlock"] = formatBlockNumber(filter.ToBlock)
	} else {
		result["blockHash"] = *filter.BlockHash
	}

	return result
}
