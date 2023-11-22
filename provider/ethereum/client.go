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
}

var _ Client = (*client)(nil)

type client struct {
	rpcClient *rpc.Client
}

func (c *client) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	var code hexutil.Bytes
	if err := c.rpcClient.CallContext(ctx, &code, "eth_getCode", contract, formatBlockNumber(blockNumber)); err != nil {
		return nil, err
	}

	return code, nil
}

func (c *client) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var value hexutil.Bytes
	if err := c.rpcClient.CallContext(ctx, &value, "eth_call", formatTransactionCall(call), formatBlockNumber(blockNumber)); err != nil {
		return nil, err
	}

	return value, nil
}

func (c *client) ChainID(ctx context.Context) (*big.Int, error) {
	var chainID *hexutil.Big
	if err := c.rpcClient.CallContext(ctx, &chainID, "eth_chainId"); err != nil {
		return nil, err
	}

	return chainID.ToInt(), nil
}

func (c *client) BlockNumber(ctx context.Context) (*big.Int, error) {
	var number *hexutil.Big
	if err := c.rpcClient.CallContext(ctx, &number, "eth_blockNumber"); err != nil {
		return nil, err
	}

	return number.ToInt(), nil
}

func (c *client) HeaderByHash(ctx context.Context, hash common.Hash) (*Header, error) {
	var header Header
	if err := c.rpcClient.CallContext(ctx, &header, "eth_getBlockByHash", hash, false); err != nil {
		return nil, err
	}

	return &header, nil
}

func (c *client) HeaderByNumber(ctx context.Context, number *big.Int) (*Header, error) {
	var header Header
	if err := c.rpcClient.CallContext(ctx, &header, "eth_getBlockByNumber", formatBlockNumber(number), false); err != nil {
		return nil, err
	}

	return &header, nil
}

func (c *client) BlockByHash(ctx context.Context, hash common.Hash) (*Block, error) {
	var block Block
	if err := c.rpcClient.CallContext(ctx, &block, "eth_getBlockByHash", hash, true); err != nil {
		return nil, err
	}

	return &block, nil
}

func (c *client) BlockByNumber(ctx context.Context, number *big.Int) (*Block, error) {
	var block Block
	if err := c.rpcClient.CallContext(ctx, &block, "eth_getBlockByNumber", formatBlockNumber(number), true); err != nil {
		return nil, err
	}

	return &block, nil
}

func (c *client) BlockReceipts(ctx context.Context, number *big.Int) ([]*Receipt, error) {
	var receipts []*Receipt
	if err := c.rpcClient.CallContext(ctx, &receipts, "eth_getBlockReceipts", formatBlockNumber(number)); err != nil {
		return nil, err
	}

	return receipts, nil
}

func (c *client) TransactionByHash(ctx context.Context, hash common.Hash) (*Transaction, error) {
	var transaction Transaction
	if err := c.rpcClient.CallContext(ctx, &transaction, "eth_getTransactionByHash", hash); err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (c *client) TransactionReceipt(ctx context.Context, hash common.Hash) (*Receipt, error) {
	var receipt Receipt
	if err := c.rpcClient.CallContext(ctx, &receipt, "eth_getTransactionReceipt", hash); err != nil {
		return nil, err
	}

	return &receipt, nil
}

func (c *client) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	var value hexutil.Bytes
	if err := c.rpcClient.CallContext(ctx, &value, "eth_getStorageAt", account, key, formatBlockNumber(blockNumber)); err != nil {
		return nil, err
	}

	return value, nil
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
