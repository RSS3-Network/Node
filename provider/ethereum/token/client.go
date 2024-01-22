package token

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/erc1155"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/multicall3"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

// LookupFunc is a function to look up token metadata.
type LookupFunc = func(ctx context.Context, chainID uint64, address *common.Address, id, blockNumber *big.Int) (*metadata.Token, error)

// Client is a client to look up token metadata.
type Client interface {
	Lookup(ctx context.Context, chainID uint64, address *common.Address, id, blockNumber *big.Int) (*metadata.Token, error)
}

// nativeTokenMap is a map of native token metadata.
var nativeTokenMap = map[uint64]metadata.Token{
	uint64(filter.EthereumChainIDMainnet): {
		Name:     "Ethereum",
		Symbol:   "ETH",
		Decimals: 18,
	},
	uint64(filter.EthereumChainIDOptimism): {
		Name:     "Ethereum",
		Symbol:   "ETH",
		Decimals: 18,
	},
	uint64(filter.EthereumChainIDPolygon): {
		Name:     "Ethereum",
		Symbol:   "ETH",
		Decimals: 18,
	},
	uint64(filter.EthereumChainIDBase): {
		Name:     "Ethereum",
		Symbol:   "ETH",
		Decimals: 18,
	},
	uint64(filter.EthereumChainIDArbitrum): {
		Name:     "Ethereum",
		Symbol:   "ETH",
		Decimals: 18,
	},
	uint64(filter.EthereumChainIDFantom): {
		Name:     "Ethereum",
		Symbol:   "ETH",
		Decimals: 18,
	},
}

var _ Client = (*client)(nil)

// client is a client to look up token metadata.
type client struct {
	ethereumClient     ethereum.Client
	unexpectedTokenMap map[uint64]map[common.Address]LookupFunc
}

// Lookup looks up token metadata, it supports ERC-20, ERC-721, ERC-1155 and native token.
func (c *client) Lookup(ctx context.Context, chainID uint64, address *common.Address, id, blockNumber *big.Int) (*metadata.Token, error) {
	// Lookup unexpected token
	if address != nil {
		if lookupMap, exists := c.unexpectedTokenMap[chainID]; exists {
			if lookup, exists := lookupMap[*address]; exists {
				return lookup(ctx, chainID, address, id, blockNumber)
			}
		}
	}

	switch {
	case address != nil && id == nil: // ERC-20 token
		return c.lookupERC20(ctx, chainID, *address, blockNumber)
	case address != nil && id != nil: // ERC-721 and ERC-1155 token
		return c.lookupNFT(ctx, chainID, *address, id, blockNumber)
	default: // Native token
		return c.lookupNative(ctx, chainID, address, id, blockNumber)
	}
}

// lookupERC20 looks up ERC-20 token metadata.
func (c *client) lookupERC20(ctx context.Context, chainID uint64, address common.Address, blockNumber *big.Int) (*metadata.Token, error) {
	tokenMetadata, err := c.lookupERC20ByRPC(ctx, chainID, address, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("lookup ERC20 by rpc: %w", err)
	}

	// Fallback NFT approval transactions
	standard, err := contract.DetectNFTStandard(ctx, chainID, address, blockNumber, c.ethereumClient)
	if err == nil && standard != contract.StandardUnknown {
		tokenMetadata.Standard = standard
	}

	return tokenMetadata, nil
}

// lookupERC20ByRPC looks up ERC-20 token metadata by RPC.
func (c *client) lookupERC20ByRPC(ctx context.Context, chainID uint64, address common.Address, blockNumber *big.Int) (*metadata.Token, error) {
	tokenMetadata := metadata.Token{
		Address:  lo.ToPtr(address.String()),
		Standard: contract.StandardERC20,
	}

	abi, err := erc20.ERC20MetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("load abi: %w", err)
	}

	calls := []multicall3.Multicall3Call3{
		{
			Target:       address,
			AllowFailure: true,
			CallData:     lo.Must(abi.Pack("name")),
		},
		{
			Target:       address,
			AllowFailure: true,
			CallData:     lo.Must(abi.Pack("symbol")),
		},
		{
			Target:       address,
			AllowFailure: true,
			CallData:     lo.Must(abi.Pack("decimals")),
		},
	}

	results, err := multicall3.Aggregate3(ctx, chainID, calls, blockNumber, c.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("call aggregate3 method: %w", err)
	}

	if results[0].Success {
		if err := abi.UnpackIntoInterface(&tokenMetadata.Name, "name", results[0].ReturnData); err != nil {
			return nil, fmt.Errorf("unpack name: %w", err)
		}
	}

	if results[1].Success {
		if err := abi.UnpackIntoInterface(&tokenMetadata.Symbol, "symbol", results[1].ReturnData); err != nil {
			return nil, fmt.Errorf("unpack symbol: %w", err)
		}
	}

	if results[2].Success {
		if err := abi.UnpackIntoInterface(&tokenMetadata.Decimals, "decimals", results[2].ReturnData); err != nil {
			return nil, fmt.Errorf("unpack decimals: %w", err)
		}
	}

	return &tokenMetadata, nil
}

// lookupNFT looks up NFT token metadata, it supports ERC-721 and ERC-1155.
func (c *client) lookupNFT(ctx context.Context, chain uint64, address common.Address, id *big.Int, blockNumber *big.Int) (*metadata.Token, error) {
	// Detect NFT standard by ERC-165.
	standard, err := contract.DetectNFTStandard(ctx, chain, address, blockNumber, c.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("detect NFT standard: %w", err)
	}

	var tokenMetadata *metadata.Token

	switch standard {
	case contract.StandardERC721:
		tokenMetadata, err = c.lookupERC721(ctx, chain, address, id, blockNumber)
	case contract.StandardERC1155:
		tokenMetadata, err = c.lookupERC1155(ctx, chain, address, id, blockNumber)
	default:
		err = fmt.Errorf("unsupported NFT standard %s", standard)
	}

	if err != nil {
		return nil, err
	}

	return tokenMetadata, nil
}

// lookupERC721 looks up ERC-721 token metadata.
func (c *client) lookupERC721(ctx context.Context, chainID uint64, address common.Address, id *big.Int, blockNumber *big.Int) (*metadata.Token, error) {
	tokenMetadata := metadata.Token{
		Address:  lo.ToPtr(address.String()),
		ID:       lo.ToPtr(decimal.NewFromBigInt(id, 0)),
		Standard: contract.StandardERC721,
	}

	abi, err := erc721.ERC721MetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("load abi: %w", err)
	}

	calls := []multicall3.Multicall3Call3{
		{
			Target:       address,
			AllowFailure: true,
			CallData:     lo.Must(abi.Pack("name")),
		},
		{
			Target:       address,
			AllowFailure: true,
			CallData:     lo.Must(abi.Pack("symbol")),
		},
		{
			Target:       address,
			AllowFailure: true,
			CallData:     lo.Must(abi.Pack("tokenURI", id)),
		},
	}

	results, err := multicall3.Aggregate3(ctx, chainID, calls, blockNumber, c.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("aggregate calls: %w", err)
	}

	if results[0].Success {
		if err := abi.UnpackIntoInterface(&tokenMetadata.Name, "name", results[0].ReturnData); err != nil {
			return nil, fmt.Errorf("unpack name: %w", err)
		}
	}

	if results[1].Success {
		if err := abi.UnpackIntoInterface(&tokenMetadata.Symbol, "symbol", results[1].ReturnData); err != nil {
			return nil, fmt.Errorf("unpack symbol: %w", err)
		}
	}

	if results[2].Success {
		if err := abi.UnpackIntoInterface(&tokenMetadata.URI, "tokenURI", results[2].ReturnData); err != nil {
			return nil, fmt.Errorf("unpack token uri: %w", err)
		}
	}

	return &tokenMetadata, nil
}

// lookupERC1155 looks up ERC-1155 token metadata.
func (c *client) lookupERC1155(ctx context.Context, chainID uint64, address common.Address, id *big.Int, blockNumber *big.Int) (*metadata.Token, error) {
	tokenMetadata := metadata.Token{
		Address:  lo.ToPtr(address.String()),
		ID:       lo.ToPtr(decimal.NewFromBigInt(id, 0)),
		Standard: contract.StandardERC1155,
	}

	abi, err := erc1155.ERC1155MetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("load abi: %w", err)
	}

	calls := []multicall3.Multicall3Call3{
		{
			Target:       address,
			AllowFailure: true,
			CallData:     lo.Must(abi.Pack("name")),
		},
		{
			Target:       address,
			AllowFailure: true,
			CallData:     lo.Must(abi.Pack("symbol")),
		},
		{
			Target:   address,
			CallData: lo.Must(abi.Pack("uri", id)),
		},
	}

	results, err := multicall3.Aggregate3(ctx, chainID, calls, blockNumber, c.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("aggregate calls: %w", err)
	}

	if results[0].Success {
		if err := abi.UnpackIntoInterface(&tokenMetadata.Name, "name", results[0].ReturnData); err != nil {
			return nil, fmt.Errorf("unpack name: %w", err)
		}
	}

	if results[1].Success {
		if err := abi.UnpackIntoInterface(&tokenMetadata.Symbol, "symbol", results[1].ReturnData); err != nil {
			return nil, fmt.Errorf("unpack symbol: %w", err)
		}
	}

	if results[2].Success {
		if err := abi.UnpackIntoInterface(&tokenMetadata.URI, "uri", results[2].ReturnData); err != nil {
			return nil, fmt.Errorf("unpack tokenURI: %w", err)
		}
	}

	return &tokenMetadata, nil
}

// lookupNative looks up native token metadata.
func (c *client) lookupNative(_ context.Context, chainID uint64, _ *common.Address, _, _ *big.Int) (*metadata.Token, error) {
	tokenMetadata, exists := nativeTokenMap[chainID]
	if !exists {
		return nil, fmt.Errorf("chain id %d does not have a native token", chainID)
	}

	return &tokenMetadata, nil
}

// lookupENS looks up ENS token metadata.
func (c *client) lookupENS(_ context.Context, _ uint64, address *common.Address, id, _ *big.Int) (*metadata.Token, error) {
	tokenMetadata := metadata.Token{
		Address:  lo.ToPtr(address.String()),
		Symbol:   "ENS",
		URI:      fmt.Sprintf("https://metadata.ens.domains/mainnet/%v/%v", address, id),
		Standard: contract.StandardERC721,
	}

	if id == nil {
		return &tokenMetadata, nil
	}

	tokenMetadata.ID = lo.ToPtr(decimal.NewFromBigInt(id, 0))

	return &tokenMetadata, nil
}

// NewClient returns a new client.
func NewClient(ethereumClient ethereum.Client) Client {
	instance := client{
		ethereumClient: ethereumClient,
	}

	instance.unexpectedTokenMap = map[uint64]map[common.Address]LookupFunc{
		uint64(filter.EthereumChainIDMainnet): {
			// ENS
			common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85"): instance.lookupENS,
		},
	}

	return &instance
}
