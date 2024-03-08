package token

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"strings"
	"unicode/utf8"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract"
	"github.com/rss3-network/node/provider/ethereum/contract/erc1155"
	"github.com/rss3-network/node/provider/ethereum/contract/erc20"
	"github.com/rss3-network/node/provider/ethereum/contract/erc721"
	"github.com/rss3-network/node/provider/ethereum/contract/multicall3"
	"github.com/rss3-network/node/provider/httpx"
	"github.com/rss3-network/node/provider/ipfs"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/tdewolff/minify/v2/minify"
	"github.com/tidwall/gjson"
	"github.com/vincent-petithory/dataurl"
)

// LookupFunc is a function to look up token metadata.
type LookupFunc = func(ctx context.Context, chainID uint64, address *common.Address, id, blockNumber *big.Int) (*metadata.Token, error)

// Client is a client to look up token metadata.
type Client interface {
	Lookup(ctx context.Context, chainID uint64, address *common.Address, id, blockNumber *big.Int, options ...Option) (*metadata.Token, error)
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
		Name:     "Polygon",
		Symbol:   "MATIC",
		Decimals: 18,
	},
	uint64(filter.EthereumChainIDBase): {
		Name:     "Ethereum",
		Symbol:   "ETH",
		Decimals: 18,
	},
	uint64(filter.EthereumChainIDFantom): {
		Name:     "Fantom",
		Symbol:   "FTM",
		Decimals: 18,
	},
	uint64(filter.EthereumChainIDCrossbell): {
		Name:     "CSB",
		Symbol:   "CSB",
		Decimals: 18,
	},
	uint64(filter.EthereumChainIDArbitrum): {
		Name:     "Ethereum",
		Symbol:   "ETH",
		Decimals: 18,
	},
	uint64(filter.EthereumChainIDAvalanche): {
		Name:     "Avalanche",
		Symbol:   "AVAX",
		Decimals: 18,
	},
	uint64(filter.EthereumChainIDRSS3Testnet): {
		Name:     "RSS3",
		Symbol:   "RSS3",
		Decimals: 18,
	},
}

var _ Client = (*client)(nil)

// client is a client to look up token metadata.
type client struct {
	ethereumClient     ethereum.Client
	ipfsClient         ipfs.HTTPClient
	httpClient         httpx.Client
	unexpectedTokenMap map[uint64]map[common.Address]LookupFunc
	parseTokenMetadata bool
}

type Option func(*client) error

func WithParseTokenMetadata(value bool) Option {
	return func(c *client) error {
		c.parseTokenMetadata = value
		return nil
	}
}

// Lookup looks up token metadata, it supports ERC-20, ERC-721, ERC-1155 and native token.
func (c *client) Lookup(ctx context.Context, chainID uint64, address *common.Address, id, blockNumber *big.Int, options ...Option) (*metadata.Token, error) {
	// Lookup unexpected token
	if address != nil {
		if lookupMap, exists := c.unexpectedTokenMap[chainID]; exists {
			if lookup, exists := lookupMap[*address]; exists {
				return lookup(ctx, chainID, address, id, blockNumber)
			}
		}
	}

	for _, option := range options {
		if err := option(c); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
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
	if err == nil && standard != metadata.StandardUnknown {
		tokenMetadata.Standard = standard
	}

	return tokenMetadata, nil
}

// lookupERC20ByRPC looks up ERC-20 token metadata by RPC.
func (c *client) lookupERC20ByRPC(ctx context.Context, chainID uint64, address common.Address, blockNumber *big.Int) (*metadata.Token, error) {
	tokenMetadata := metadata.Token{
		Address:  lo.ToPtr(address.String()),
		Standard: metadata.StandardERC20,
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

	if c.parseTokenMetadata {
		// When handling nft, we should get image uri from tokenURI
		// Initialize ipfs client.
		if c.ipfsClient, err = ipfs.NewHTTPClient(); err != nil {
			return nil, fmt.Errorf("new ipfs client: %w", err)
		}

		// Initialize http client.
		if c.httpClient, err = httpx.NewHTTPClient(); err != nil {
			return nil, fmt.Errorf("new http client: %w", err)
		}
	}

	switch standard {
	case metadata.StandardERC721:
		tokenMetadata, err = c.lookupERC721(ctx, chain, address, id, blockNumber)
	case metadata.StandardERC1155:
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
		Standard: metadata.StandardERC721,
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

	if c.parseTokenMetadata {
		// Ignore invalid URI
		nonFungibleTokenMetadata, err := c.buildNonFungibleTokenMetadata(ctx, tokenMetadata.URI, id)
		if err == nil {
			tokenMetadata.ParsedImageURL = nonFungibleTokenMetadata.ImageURL
		}

		if isDataURIErr(err) {
			return nil, fmt.Errorf("unsupport data uri %s %w", tokenMetadata.URI, err)
		}
	}

	return &tokenMetadata, nil
}

// lookupERC1155 looks up ERC-1155 token metadata.
func (c *client) lookupERC1155(ctx context.Context, chainID uint64, address common.Address, id *big.Int, blockNumber *big.Int) (*metadata.Token, error) {
	tokenMetadata := metadata.Token{
		Address:  lo.ToPtr(address.String()),
		ID:       lo.ToPtr(decimal.NewFromBigInt(id, 0)),
		Standard: metadata.StandardERC1155,
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

	if c.parseTokenMetadata {
		// Ignore invalid URI
		nonFungibleTokenMetadata, err := c.buildNonFungibleTokenMetadata(ctx, tokenMetadata.URI, id)
		if err == nil {
			tokenMetadata.ParsedImageURL = nonFungibleTokenMetadata.ImageURL
		}

		if isDataURIErr(err) {
			return nil, fmt.Errorf("unsupport data uri %s %w", tokenMetadata.URI, err)
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
		Standard: metadata.StandardERC721,
	}

	if id == nil {
		return &tokenMetadata, nil
	}

	tokenMetadata.ID = lo.ToPtr(decimal.NewFromBigInt(id, 0))

	return &tokenMetadata, nil
}

func (c *client) lookupMaker(_ context.Context, _ uint64, address *common.Address, _, _ *big.Int) (*metadata.Token, error) {
	tokenMetadata := metadata.Token{
		Address:  lo.ToPtr(address.String()),
		Name:     "Maker",
		Symbol:   "MKR",
		Decimals: 18,
		Standard: metadata.StandardERC20,
	}

	return &tokenMetadata, nil
}

// buildNonFungibleTokenMetadata builds non-fungible token metadata.
func (c *client) buildNonFungibleTokenMetadata(ctx context.Context, uri string, id *big.Int) (*metadata.NonFungibleTokenMetadata, error) {
	var (
		nonFungibleTokenMetadata metadata.NonFungibleTokenMetadata
		metadata                 json.RawMessage
	)

	if uri == "" {
		return &nonFungibleTokenMetadata, nil
	}

	var isDataURI bool

	dataURI, err := dataurl.DecodeString(uri)

	if err == nil {
		metadata, isDataURI = dataURI.Data, true
	}

	if isDataURIErr(err) {
		return nil, err
	}

	// Format URI with ID
	uri = strings.ReplaceAll(uri, "0x{id}", hexutil.EncodeBig(id))
	uri = strings.ReplaceAll(uri, "{id}", id.String())

	_, path, err := ipfs.ParseURL(uri)

	switch {
	case isDataURI: // Data URI
	case path != "": // IPFS
		if err != nil {
			return nil, fmt.Errorf("parse IPFS URL: %w", err)
		}

		readCloser, err := c.ipfsClient.Fetch(ctx, path, ipfs.FetchModeQuick)
		if err != nil {
			return nil, fmt.Errorf("quick fetch %s: %w", path, err)
		}

		if metadata, err = io.ReadAll(readCloser); err != nil {
			return nil, fmt.Errorf("read metadata from IPFS: %w", err)
		}
	case strings.HasPrefix(uri, "ar://"): // Arweave
		uri = strings.Replace(uri, "ar://", "https://arweave.net/", 1)

		fallthrough
	default:
		response, err := c.httpClient.Fetch(ctx, uri)
		if err != nil {
			return nil, fmt.Errorf("fetch metadata from HTTP %s: %w", uri, err)
		}

		if metadata, err = io.ReadAll(response); err != nil {
			return nil, fmt.Errorf("read metadata from HTTP %s: %w", uri, err)
		}
	}

	if err := c.standardizeNonFungibleTokenMetadata(ctx, &nonFungibleTokenMetadata, metadata); err != nil {
		return nil, fmt.Errorf("standardize NFT metadata: %w", err)
	}

	return &nonFungibleTokenMetadata, nil
}

// TODO: Need better Data URI package.
func isDataURIErr(err error) bool {
	return err != nil && strings.HasPrefix(err.Error(), "expected base64, got")
}

// removeInvalidCharacter removes invalid character.
func removeInvalidCharacter(character string) string {
	if !utf8.ValidString(character) {
		return ""
	}

	return character
}

// standardizeNonFungibleTokenMetadata standardizes non-fungible token metadata.
func (c *client) standardizeNonFungibleTokenMetadata(_ context.Context, nonFungibleTokenMetadata *metadata.NonFungibleTokenMetadata, data []byte) error {
	result := gjson.ParseBytes(data)

	if name := result.Get("name"); name.Exists() {
		nonFungibleTokenMetadata.Title = removeInvalidCharacter(name.String())
	} else {
		nonFungibleTokenMetadata.Title = removeInvalidCharacter(result.Get("title").String())
	}

	nonFungibleTokenMetadata.Description = removeInvalidCharacter(result.Get("description").String())

	if imageURL := result.Get("image"); imageURL.Exists() {
		nonFungibleTokenMetadata.ImageURL = imageURL.String()
	} else {
		nonFungibleTokenMetadata.ImageURL = result.Get("image_url").String()
	}

	nonFungibleTokenMetadata.ExternalURL = result.Get("external_url").String()
	nonFungibleTokenMetadata.MediaURL = result.Get("media_url").String()
	nonFungibleTokenMetadata.AnimationURL = result.Get("animation_url").String()

	var properties string

	switch {
	case result.Get("properties").Exists():
		properties = removeInvalidCharacter(result.Get("properties").Raw)
	case result.Get("attributes").Exists():
		properties = removeInvalidCharacter(result.Get("attributes").Raw)
	case result.Get("traits").Exists():
		properties = removeInvalidCharacter(result.Get("traits").Raw)
	}

	properties, err := minify.JSON(properties)
	if err != nil {
		return fmt.Errorf("minify JSON: %w", err)
	}

	nonFungibleTokenMetadata.Properties = json.RawMessage(properties)

	// Fallback to default value if properties is invalid JSON
	if !json.Valid(nonFungibleTokenMetadata.Properties) {
		nonFungibleTokenMetadata.Properties = json.RawMessage{}
	}

	return nil
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
			common.HexToAddress("0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2"): instance.lookupMaker,
		},
	}

	return &instance
}
