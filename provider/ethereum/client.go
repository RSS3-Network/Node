package ethereum

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/samber/lo"
)

// Client provides basic RPC methods.
type Client interface {
	// ContractCaller is used for compatibility with abigen generated contract callers.
	bind.ContractCaller

	ChainID(ctx context.Context) (*big.Int, error)
	BlockNumber(ctx context.Context) (*big.Int, error)
	HeaderByHash(ctx context.Context, hash common.Hash) (*Header, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*Header, error)
	BlockByHash(ctx context.Context, hash common.Hash) (*Block, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*Block, error)
	BatchBlockByNumbers(ctx context.Context, numbers []*big.Int) ([]*Block, error)
	BlockReceipts(ctx context.Context, number *big.Int) ([]*Receipt, error)
	BatchBlockReceipts(ctx context.Context, numbers []*big.Int) ([][]*Receipt, error)
	TransactionByHash(ctx context.Context, hash common.Hash) (*Transaction, error)
	TransactionReceipt(ctx context.Context, hash common.Hash) (*Receipt, error)
	BatchTransactionReceipt(ctx context.Context, hashes []common.Hash) ([]*Receipt, error)
	StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error)
	FilterLogs(ctx context.Context, filter Filter) ([]*Log, error)
}

var _ Client = (*client)(nil)

type client struct {
	endpoint  string
	rpcClient *rpc.Client
}

// CodeAt returns the contract code of the given account.
func (c *client) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	var code hexutil.Bytes
	if err := c.rpcClient.CallContext(ctx, &code, "eth_getCode", contract, formatBlockNumber(blockNumber)); err != nil {
		return nil, err
	}

	return code, nil
}

// CallContract returns the result of the given transaction call.
func (c *client) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var value hexutil.Bytes
	if err := c.rpcClient.CallContext(ctx, &value, "eth_call", formatTransactionCall(call), formatBlockNumber(blockNumber)); err != nil {
		return nil, err
	}

	return value, nil
}

// ChainID returns the chain ID.
func (c *client) ChainID(ctx context.Context) (*big.Int, error) {
	var chainID *hexutil.Big
	if err := c.rpcClient.CallContext(ctx, &chainID, "eth_chainId"); err != nil {
		return nil, err
	}

	return chainID.ToInt(), nil
}

// BlockNumber returns the current block number.
func (c *client) BlockNumber(ctx context.Context) (*big.Int, error) {
	var number *hexutil.Big
	if err := c.rpcClient.CallContext(ctx, &number, "eth_blockNumber"); err != nil {
		return nil, err
	}

	return number.ToInt(), nil
}

// HeaderByHash returns the header with the given hash.
func (c *client) HeaderByHash(ctx context.Context, hash common.Hash) (*Header, error) {
	var header *Header
	if err := c.rpcClient.CallContext(ctx, &header, "eth_getBlockByHash", hash, false); err != nil {
		return nil, err
	}

	if header == nil {
		return nil, ethereum.NotFound
	}

	return header, nil
}

// HeaderByNumber returns the header with the given number.
func (c *client) HeaderByNumber(ctx context.Context, number *big.Int) (*Header, error) {
	var header *Header
	if err := c.rpcClient.CallContext(ctx, &header, "eth_getBlockByNumber", formatBlockNumber(number), false); err != nil {
		return nil, err
	}

	if header == nil {
		return nil, ethereum.NotFound
	}

	return header, nil
}

// BlockByHash returns the block with the given hash.
func (c *client) BlockByHash(ctx context.Context, hash common.Hash) (*Block, error) {
	var block *Block
	if err := c.rpcClient.CallContext(ctx, &block, "eth_getBlockByHash", hash, true); err != nil {
		return nil, err
	}

	if block == nil {
		return nil, ethereum.NotFound
	}

	return block, nil
}

// BlockByNumber returns the block with the given number.
func (c *client) BlockByNumber(ctx context.Context, number *big.Int) (*Block, error) {
	var block *Block
	if err := c.rpcClient.CallContext(ctx, &block, "eth_getBlockByNumber", formatBlockNumber(number), true); err != nil {
		return nil, err
	}

	if block == nil {
		return nil, ethereum.NotFound
	}

	return block, nil
}

// BatchBlockByNumbers returns the blocks with the given numbers.
func (c *client) BatchBlockByNumbers(ctx context.Context, numbers []*big.Int) ([]*Block, error) {
	batchElements := make([]rpc.BatchElem, 0, len(numbers))

	for _, number := range numbers {
		batchElement := rpc.BatchElem{
			Method: "eth_getBlockByNumber",
			Args: []any{
				formatBlockNumber(number),
				true,
			},
			Result: lo.ToPtr((*Block)(nil)),
		}

		batchElements = append(batchElements, batchElement)
	}

	if err := c.rpcClient.BatchCallContext(ctx, batchElements); err != nil {
		return nil, err
	}

	blocks, err := unwrapBatchElements[**Block](batchElements)
	if err != nil {
		return nil, err
	}

	// Check if blocks are not found for some block numbers.
	nullBlockNumbers := lo.FilterMap(blocks, func(block **Block, index int) (string, bool) {
		return numbers[index].String(), *block == nil
	})

	if len(nullBlockNumbers) > 0 {
		return nil, fmt.Errorf("blocks of block numbers %s: %w", strings.Join(nullBlockNumbers, ",\x20"), ethereum.NotFound)
	}

	return derefPtrSlice(blocks), nil
}

// BlockReceipts returns the receipts of a block by block number.
func (c *client) BlockReceipts(ctx context.Context, number *big.Int) ([]*Receipt, error) {
	var receipts *[]*Receipt
	if err := c.rpcClient.CallContext(ctx, &receipts, "eth_getBlockReceipts", formatBlockNumber(number)); err != nil {
		return nil, err
	}

	if receipts == nil {
		return nil, ethereum.NotFound
	}

	return *receipts, nil
}

// BatchBlockReceipts returns the receipts of multiple blocks by block numbers.
func (c *client) BatchBlockReceipts(ctx context.Context, numbers []*big.Int) ([][]*Receipt, error) {
	batchElements := make([]rpc.BatchElem, 0, len(numbers))

	for _, number := range numbers {
		batchElement := rpc.BatchElem{
			Method: "eth_getBlockReceipts",
			Args: []any{
				formatBlockNumber(number),
			},
			Result: new([]*Receipt),
		}

		batchElements = append(batchElements, batchElement)
	}

	if err := c.rpcClient.BatchCallContext(ctx, batchElements); err != nil {
		return nil, err
	}

	batchReceipts, err := unwrapBatchElements[*[]*Receipt](batchElements)
	if err != nil {
		return nil, err
	}

	// Check if receipts are not found for some blocks.
	nullBlockNumbers := lo.FilterMap(batchReceipts, func(receipts *[]*Receipt, index int) (string, bool) {
		return numbers[index].String(), receipts == nil
	})

	if len(nullBlockNumbers) > 0 {
		return nil, fmt.Errorf("receipts of block numbers %s: %w", strings.Join(nullBlockNumbers, ",\x20"), ethereum.NotFound)
	}

	return lo.Map(batchReceipts, func(receipts *[]*Receipt, _ int) []*Receipt { return *receipts }), nil
}

// TransactionByHash returns the transaction with the given hash.
func (c *client) TransactionByHash(ctx context.Context, hash common.Hash) (*Transaction, error) {
	var transaction *Transaction
	if err := c.rpcClient.CallContext(ctx, &transaction, "eth_getTransactionByHash", hash); err != nil {
		return nil, err
	}

	if transaction == nil {
		return nil, ethereum.NotFound
	}

	return transaction, nil
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
func (c *client) TransactionReceipt(ctx context.Context, hash common.Hash) (*Receipt, error) {
	var receipt *Receipt
	if err := c.rpcClient.CallContext(ctx, &receipt, "eth_getTransactionReceipt", hash); err != nil {
		return nil, err
	}

	if receipt == nil {
		return nil, ethereum.NotFound
	}

	return receipt, nil
}

// BatchTransactionReceipt returns the receipts of multiple transactions by transaction hashes.
func (c *client) BatchTransactionReceipt(ctx context.Context, hashes []common.Hash) ([]*Receipt, error) {
	batchElements := make([]rpc.BatchElem, 0, len(hashes))

	for _, hash := range hashes {
		batchElement := rpc.BatchElem{
			Method: "eth_getTransactionReceipt",
			Args: []any{
				hash,
			},
			Result: lo.ToPtr((*Receipt)(nil)),
		}

		batchElements = append(batchElements, batchElement)
	}

	if err := c.rpcClient.BatchCallContext(ctx, batchElements); err != nil {
		return nil, err
	}

	receipts, err := unwrapBatchElements[**Receipt](batchElements)
	if err != nil {
		return nil, err
	}

	// Check if receipts are not found for some transaction hashes.
	nullTransactionHashes := lo.FilterMap(receipts, func(receipt **Receipt, index int) (string, bool) {
		return hashes[index].String(), *receipt == nil
	})

	if len(nullTransactionHashes) > 0 {
		return nil, fmt.Errorf("receipts of transaction hashes %s: %w", strings.Join(nullTransactionHashes, ",\x20"), ethereum.NotFound)
	}

	return derefPtrSlice(receipts), nil
}

// StorageAt returns the contract storage of the given account.
func (c *client) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	var value hexutil.Bytes
	if err := c.rpcClient.CallContext(ctx, &value, "eth_getStorageAt", account, key, formatBlockNumber(blockNumber)); err != nil {
		return nil, err
	}

	return value, nil
}

// FilterLogs returns the logs that satisfy the filter conditions.
func (c *client) FilterLogs(ctx context.Context, filter Filter) ([]*Log, error) {
	var logs []*Log
	if err := c.rpcClient.CallContext(ctx, &logs, "eth_getLogs", formatFilter(filter)); err != nil {
		return nil, err
	}

	return logs, nil
}

// Dial creates a new client for the given endpoint.
func Dial(ctx context.Context, endpoint string, options ...Option) (Client, error) {
	rpcClient, err := rpc.DialContext(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	instance := client{
		endpoint:  endpoint,
		rpcClient: rpcClient,
	}

	for _, option := range options {
		if err := option(ctx, &instance); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return &instance, nil
}

// unwrapBatchElement unwraps the result of a batch call.
func unwrapBatchElement[T any](batchElement rpc.BatchElem) (T, error) {
	var empty T // The default value is nil

	if batchElement.Error != nil {
		return empty, batchElement.Error
	}

	result, ok := batchElement.Result.(T)
	if !ok {
		return empty, fmt.Errorf("invalid type %T", batchElement.Result)
	}

	return result, nil
}

// unwrapBatchElements unwraps the results of a batch call.
func unwrapBatchElements[T any](batchElements []rpc.BatchElem) ([]T, error) {
	var (
		empty   []T // The default value is nil
		results = make([]T, 0, len(batchElements))
	)

	for _, batchElement := range batchElements {
		result, err := unwrapBatchElement[T](batchElement)
		if err != nil {
			return empty, err
		}

		results = append(results, result)
	}

	return results, nil
}

// derefPtrSlice converts a slice of pointers to a slice of values.
func derefPtrSlice[T any](values []*T) []T {
	return lo.Map(values, func(value *T, _ int) T {
		return *value
	})
}
