package token

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/erc1155"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

type LookupFunc = func(ctx context.Context, chain filter.ChainEthereum, address *common.Address, id, blockNumber *big.Int) (*metadata.Token, error)

type Client interface {
	Lookup(ctx context.Context, chain filter.ChainEthereum, address *common.Address, id, blockNumber *big.Int) (*metadata.Token, error)
}

var nativeTokenMap = map[filter.Chain]metadata.Token{
	filter.ChainEthereumMainnet: {
		Name:     "Ethereum",
		Symbol:   "ETH",
		Decimals: 18,
	},
}

var _ Client = (*client)(nil)

type client struct {
	ethereumClient ethereum.Client
}

func (c *client) Lookup(ctx context.Context, chain filter.ChainEthereum, address *common.Address, id, blockNumber *big.Int) (*metadata.Token, error) {
	switch {
	case address != nil && id == nil: // ERC-20 token
		return c.lookupERC20(ctx, chain, *address, blockNumber)
	case address != nil && id != nil: // ERC-721 and ERC-1155 token
		return c.lookupNFT(ctx, chain, *address, id, blockNumber)
	default: // Native token
		return c.lookupNative(ctx, chain, address, id, blockNumber)
	}
}

func (c *client) lookupERC20(ctx context.Context, chain filter.ChainEthereum, address common.Address, blockNumber *big.Int) (*metadata.Token, error) {
	tokenMetadata, err := c.lookupERC20ByRPC(ctx, chain, address, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("lookup ERC20 by rpc: %w", err)
	}

	// Fallback NFT approval transactions
	standard, err := contract.DetectNFTStandard(ctx, new(big.Int).SetUint64(chain.ID()), address, blockNumber, c.ethereumClient)
	if err == nil && standard != contract.StandardUnknown {
		tokenMetadata.Standard = standard
	}

	return tokenMetadata, nil
}

func (c *client) lookupERC20ByRPC(ctx context.Context, _ filter.ChainEthereum, address common.Address, blockNumber *big.Int) (*metadata.Token, error) {
	tokenMetadata := metadata.Token{
		Address:  lo.ToPtr(address.String()),
		Standard: contract.StandardERC20,
	}

	caller, err := erc20.NewERC20Caller(address, c.ethereumClient)
	if err != nil {
		return nil, err
	}

	callOptions := bind.CallOpts{
		BlockNumber: blockNumber,
		Context:     ctx,
	}

	if tokenMetadata.Name, err = caller.Name(&callOptions); err != nil {
		return nil, err
	}

	if tokenMetadata.Symbol, err = caller.Symbol(&callOptions); err != nil {
		return nil, err
	}

	if tokenMetadata.Decimals, err = caller.Decimals(&callOptions); err != nil {
		return nil, err
	}

	return &tokenMetadata, nil
}

func (c *client) lookupNFT(ctx context.Context, chain filter.ChainEthereum, address common.Address, id *big.Int, blockNumber *big.Int) (*metadata.Token, error) {
	chainID := new(big.Int).SetUint64(chain.ID())

	standard, err := contract.DetectNFTStandard(ctx, chainID, address, blockNumber, c.ethereumClient)
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

func (c *client) lookupERC721(ctx context.Context, _ filter.ChainEthereum, address common.Address, id *big.Int, blockNumber *big.Int) (*metadata.Token, error) {
	tokenMetadata := metadata.Token{
		Address:  lo.ToPtr(address.String()),
		ID:       lo.ToPtr(decimal.NewFromBigInt(id, 0)),
		Standard: contract.StandardERC721,
	}

	caller, err := erc721.NewERC721Caller(address, c.ethereumClient)
	if err != nil {
		return nil, err
	}

	callOptions := bind.CallOpts{
		BlockNumber: blockNumber,
		Context:     ctx,
	}

	if tokenMetadata.Name, err = caller.Name(&callOptions); err != nil {
		return nil, err
	}

	if tokenMetadata.Symbol, err = caller.Symbol(&callOptions); err != nil {
		return nil, err
	}

	if tokenMetadata.URI, err = caller.TokenURI(&callOptions, id); err != nil {
		return nil, err
	}

	return &tokenMetadata, nil
}

func (c *client) lookupERC1155(ctx context.Context, _ filter.ChainEthereum, address common.Address, id *big.Int, blockNumber *big.Int) (*metadata.Token, error) {
	tokenMetadata := metadata.Token{
		Address:  lo.ToPtr(address.String()),
		ID:       lo.ToPtr(decimal.NewFromBigInt(id, 0)),
		Standard: contract.StandardERC1155,
	}

	caller, err := erc1155.NewERC1155Caller(address, c.ethereumClient)
	if err != nil {
		return nil, err
	}

	callOptions := bind.CallOpts{
		BlockNumber: blockNumber,
		Context:     ctx,
	}

	if tokenMetadata.Name, err = caller.Name(&callOptions); err != nil {
		return nil, err
	}

	if tokenMetadata.Symbol, err = caller.Symbol(&callOptions); err != nil {
		return nil, err
	}

	if tokenMetadata.URI, err = caller.Uri(&callOptions, id); err != nil {
		return nil, err
	}

	return &tokenMetadata, nil
}

func (c *client) lookupNative(_ context.Context, chain filter.ChainEthereum, _ *common.Address, _, _ *big.Int) (*metadata.Token, error) {
	tokenMetadata, exists := nativeTokenMap[chain]
	if !exists {
		return nil, fmt.Errorf("chain %s does not have a native token", chain)
	}

	return &tokenMetadata, nil
}

func NewClient(ethereumClient ethereum.Client) Client {
	instance := client{
		ethereumClient: ethereumClient,
	}

	return &instance
}
