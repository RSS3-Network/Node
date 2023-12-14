package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
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
	BlockReceipts(ctx context.Context, number *big.Int) ([]*Receipt, error)
	TransactionByHash(ctx context.Context, hash common.Hash) (*Transaction, error)
	TransactionReceipt(ctx context.Context, hash common.Hash) (*Receipt, error)
	StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error)
	FilterLogs(ctx context.Context, filter Filter) ([]*Log, error)
}

var _ Client = (*client)(nil)

type client struct {
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
	var header Header
	if err := c.rpcClient.CallContext(ctx, &header, "eth_getBlockByHash", hash, false); err != nil {
		return nil, err
	}

	return &header, nil
}

// HeaderByNumber returns the header with the given number.
func (c *client) HeaderByNumber(ctx context.Context, number *big.Int) (*Header, error) {
	var header Header
	if err := c.rpcClient.CallContext(ctx, &header, "eth_getBlockByNumber", formatBlockNumber(number), false); err != nil {
		return nil, err
	}

	return &header, nil
}

// BlockByHash returns the block with the given hash.
func (c *client) BlockByHash(ctx context.Context, hash common.Hash) (*Block, error) {
	var block Block
	if err := c.rpcClient.CallContext(ctx, &block, "eth_getBlockByHash", hash, true); err != nil {
		return nil, err
	}

	return &block, nil
}

// BlockByNumber returns the block with the given number.
func (c *client) BlockByNumber(ctx context.Context, number *big.Int) (*Block, error) {
	var block Block
	if err := c.rpcClient.CallContext(ctx, &block, "eth_getBlockByNumber", formatBlockNumber(number), true); err != nil {
		return nil, err
	}

	return &block, nil
}

// BlockReceipts returns the receipts of a block by block number.
func (c *client) BlockReceipts(ctx context.Context, number *big.Int) ([]*Receipt, error) {
	var receipts []*Receipt
	if err := c.rpcClient.CallContext(ctx, &receipts, "eth_getBlockReceipts", formatBlockNumber(number)); err != nil {
		return nil, err
	}

	return receipts, nil
}

// TransactionByHash returns the transaction with the given hash.
func (c *client) TransactionByHash(ctx context.Context, hash common.Hash) (*Transaction, error) {
	var transaction Transaction
	if err := c.rpcClient.CallContext(ctx, &transaction, "eth_getTransactionByHash", hash); err != nil {
		return nil, err
	}

	return &transaction, nil
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
func (c *client) TransactionReceipt(ctx context.Context, hash common.Hash) (*Receipt, error) {
	var receipt Receipt
	if err := c.rpcClient.CallContext(ctx, &receipt, "eth_getTransactionReceipt", hash); err != nil {
		return nil, err
	}

	return &receipt, nil
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
func Dial(ctx context.Context, endpoint string) (Client, error) {
	rpcClient, err := rpc.DialContext(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	instance := client{
		rpcClient: rpcClient,
	}

	return &instance, nil
}
