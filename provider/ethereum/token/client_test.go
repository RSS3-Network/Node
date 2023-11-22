package token_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/endpoint"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/token"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestEthereumClient(t *testing.T) {
	t.Parallel()

	type arguments struct {
		chain   filter.ChainEthereum
		address *common.Address
		id      *big.Int
	}

	var testcases = []struct {
		name      string
		arguments arguments
		want      metadata.Token
	}{
		{
			name: "ETH",
			arguments: arguments{
				chain: filter.ChainEthereumMainnet,
			},
			want: metadata.Token{
				Name:     "Ethereum",
				Symbol:   "ETH",
				Decimals: 18,
			},
		},
		{
			name: "RSS3",
			arguments: arguments{
				chain:   filter.ChainEthereumMainnet,
				address: lo.ToPtr(common.HexToAddress("0xc98D64DA73a6616c42117b582e832812e7B8D57F")),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0xc98D64DA73a6616c42117b582e832812e7B8D57F"),
				Name:     "RSS3",
				Symbol:   "RSS3",
				Decimals: 18,
				Standard: contract.StandardERC20,
			},
		},
		{
			name: "RSS Fruits Token #61",
			arguments: arguments{
				chain:   filter.ChainEthereumMainnet,
				address: lo.ToPtr(common.HexToAddress("0xAcbe98EFe2d4D103e221E04c76D7c55dB15C8E89")),
				id:      big.NewInt(61),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0xAcbe98EFe2d4D103e221E04c76D7c55dB15C8E89"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("61"))),
				Name:     "RSS Fruits Token",
				Symbol:   "RFT",
				Standard: contract.StandardERC721,
				URI:      "https://gateway.pinata.cloud/ipfs/QmRjC25urVAke71UYAV4PoMi4mCACcqG7MizjLVnJuVQyw",
			},
		},
		{
			name: "Base, Introduced #1",
			arguments: arguments{
				chain:   filter.ChainEthereumMainnet,
				address: lo.ToPtr(common.HexToAddress("0xD4307E0acD12CF46fD6cf93BC264f5D5D1598792")),
				id:      big.NewInt(1),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0xD4307E0acD12CF46fD6cf93BC264f5D5D1598792"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
				Name:     "Base, Introduced",
				Symbol:   "BASEINTRODUCED",
				Standard: contract.StandardERC721,
				URI:      "data:application/json;base64,eyJuYW1lIjogIkJhc2UsIEludHJvZHVjZWQgMSIsICJkZXNjcmlwdGlvbiI6ICJNZWV0IEJhc2UsIGFuIEV0aGVyZXVtIEwyIHRoYXQgb2ZmZXJzIGEgc2VjdXJlLCBsb3ctY29zdCwgZGV2ZWxvcGVyLWZyaWVuZGx5IHdheSBmb3IgYW55b25lLCBhbnl3aGVyZSwgdG8gYnVpbGQgZGVjZW50cmFsaXplZCBhcHBzLlxuXG5XZSBjb2xsZWN0aXZlbHkgbWludGVkIOKAmEJhc2UsIEludHJvZHVjZWTigJkgdG8gY2VsZWJyYXRlIHRoZSB0ZXN0bmV0IGxhdW5jaCBhbmQgZ3JvdyB0aGUgYnJvYWRlciBCYXNlIGNvbW11bml0eS4gV2XigJlyZSBleGNpdGVkIHRvIGJ1aWxkIEJhc2UgdG9nZXRoZXIuIiwgImltYWdlIjogImlwZnM6Ly9iYWZ5YmVpYmh0azIzaDZzYXM0eXVhaHR5eTd2Mmt6dndvd3c3aGU0NHJoaXg3a3E0NHJmMmFmM2ZjcSIsICJwcm9wZXJ0aWVzIjogeyJudW1iZXIiOiAxLCAibmFtZSI6ICJCYXNlLCBJbnRyb2R1Y2VkIn19",
			},
		},
		{
			name: "Cheers UP #0",
			arguments: arguments{
				chain:   filter.ChainEthereumMainnet,
				address: lo.ToPtr(common.HexToAddress("0x3113A3c04aEBEC2B77eB38Eabf6a2257B580c54B")),
				id:      big.NewInt(0),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0x3113A3c04aEBEC2B77eB38Eabf6a2257B580c54B"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("0"))),
				Name:     "Cheers UP",
				Symbol:   "CUP",
				Standard: contract.StandardERC721,
				URI:      "ipfs://QmR4fuz6w9oKEo6oqwFdTmuXqWwmrsFwv659tSZr1SJiNR",
			},
		},
		{
			name: "RSS3 Whitepaper #1",
			arguments: arguments{
				chain:   filter.ChainEthereumMainnet,
				address: lo.ToPtr(common.HexToAddress("0xB9619cF4F875CdF0E3CE48B28A1C725bC4f6c0fB")),
				id:      big.NewInt(1),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0xB9619cF4F875CdF0E3CE48B28A1C725bC4f6c0fB"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
				Name:     "RSS3 Whitepaper",
				Symbol:   "RWP",
				Standard: contract.StandardERC721,
				URI:      "ipfs://QmTMD6sLA7M4iegKDhbdMPBZ4HLi5fjW27w2J16gqc5Cb7/1.json",
			},
		},
		{
			name: "OpenSea Shared Storefront #4452815761501376758038898720702591022279500679302039557361006834352129",
			arguments: arguments{
				chain:   filter.ChainEthereumMainnet,
				address: lo.ToPtr(common.HexToAddress("0x495f947276749Ce646f68AC8c248420045cb7b5e")),
				id:      lo.Must(new(big.Int).SetString("4452815761501376758038898720702591022279500679302039557361006834352129", 0)),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0x495f947276749Ce646f68AC8c248420045cb7b5e"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("4452815761501376758038898720702591022279500679302039557361006834352129"))),
				Name:     "OpenSea Shared Storefront",
				Symbol:   "OPENSTORE",
				Standard: contract.StandardERC1155,
				URI:      "https://api.opensea.io/api/v1/metadata/0x495f947276749Ce646f68AC8c248420045cb7b5e/0x{id}",
			},
		},
		{
			name: "Love, Death + Robots",
			arguments: arguments{
				chain:   filter.ChainEthereumMainnet,
				address: lo.ToPtr(common.HexToAddress("0xFD43D1dA000558473822302e1d44D81dA2e4cC0d")),
				id:      big.NewInt(1),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0xFD43D1dA000558473822302e1d44D81dA2e4cC0d"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
				Name:     "Love, Death + Robots",
				Symbol:   "LDR",
				Standard: contract.StandardERC1155,
				URI:      "ipfs://QmNjmcjL2cz1LrHmTXy3CpVifRNoy3TDh6duo7jxkFFBZH/{id}",
			},
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
			defer cancel()

			ethereumClient, err := ethereum.Dial(ctx, endpoint.MustGet(testcase.arguments.chain))
			require.NoError(t, err)

			tokenClient := token.NewClient(ethereumClient)

			tokenMetadata, err := tokenClient.Lookup(ctx, testcase.arguments.chain, testcase.arguments.address, testcase.arguments.id, nil)
			require.NoError(t, err)

			require.Equal(t, testcase.want, *tokenMetadata)
		})
	}
}
