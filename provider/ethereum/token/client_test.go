package token_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	"github.com/rss3-network/node/provider/ethereum/token"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestEthereumClient(t *testing.T) {
	t.Parallel()

	type arguments struct {
		network filter.Network
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
				network: filter.NetworkEthereum,
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
				network: filter.NetworkEthereum,
				address: lo.ToPtr(common.HexToAddress("0xc98D64DA73a6616c42117b582e832812e7B8D57F")),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0xc98D64DA73a6616c42117b582e832812e7B8D57F"),
				Name:     "RSS3",
				Symbol:   "RSS3",
				Decimals: 18,
				Standard: metadata.StandardERC20,
			},
		},
		{
			name: "RSS Fruits Token #61",
			arguments: arguments{
				network: filter.NetworkEthereum,
				address: lo.ToPtr(common.HexToAddress("0xAcbe98EFe2d4D103e221E04c76D7c55dB15C8E89")),
				id:      big.NewInt(61),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0xAcbe98EFe2d4D103e221E04c76D7c55dB15C8E89"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("61"))),
				Name:     "RSS Fruits Token",
				Symbol:   "RFT",
				Standard: metadata.StandardERC721,
				URI:      "https://gateway.pinata.cloud/ipfs/QmRjC25urVAke71UYAV4PoMi4mCACcqG7MizjLVnJuVQyw",
			},
		},
		{
			name: "Base, Introduced #1",
			arguments: arguments{
				network: filter.NetworkEthereum,
				address: lo.ToPtr(common.HexToAddress("0xD4307E0acD12CF46fD6cf93BC264f5D5D1598792")),
				id:      big.NewInt(1),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0xD4307E0acD12CF46fD6cf93BC264f5D5D1598792"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
				Name:     "Base, Introduced",
				Symbol:   "BASEINTRODUCED",
				Standard: metadata.StandardERC721,
				URI:      "data:application/json;base64,eyJuYW1lIjogIkJhc2UsIEludHJvZHVjZWQgMSIsICJkZXNjcmlwdGlvbiI6ICJNZWV0IEJhc2UsIGFuIEV0aGVyZXVtIEwyIHRoYXQgb2ZmZXJzIGEgc2VjdXJlLCBsb3ctY29zdCwgZGV2ZWxvcGVyLWZyaWVuZGx5IHdheSBmb3IgYW55b25lLCBhbnl3aGVyZSwgdG8gYnVpbGQgZGVjZW50cmFsaXplZCBhcHBzLlxuXG5XZSBjb2xsZWN0aXZlbHkgbWludGVkIOKAmEJhc2UsIEludHJvZHVjZWTigJkgdG8gY2VsZWJyYXRlIHRoZSB0ZXN0bmV0IGxhdW5jaCBhbmQgZ3JvdyB0aGUgYnJvYWRlciBCYXNlIGNvbW11bml0eS4gV2XigJlyZSBleGNpdGVkIHRvIGJ1aWxkIEJhc2UgdG9nZXRoZXIuIiwgImltYWdlIjogImlwZnM6Ly9iYWZ5YmVpYmh0azIzaDZzYXM0eXVhaHR5eTd2Mmt6dndvd3c3aGU0NHJoaXg3a3E0NHJmMmFmM2ZjcSIsICJwcm9wZXJ0aWVzIjogeyJudW1iZXIiOiAxLCAibmFtZSI6ICJCYXNlLCBJbnRyb2R1Y2VkIn19",
			},
		},
		{
			name: "Cheers UP #0",
			arguments: arguments{
				network: filter.NetworkEthereum,
				address: lo.ToPtr(common.HexToAddress("0x3113A3c04aEBEC2B77eB38Eabf6a2257B580c54B")),
				id:      big.NewInt(0),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0x3113A3c04aEBEC2B77eB38Eabf6a2257B580c54B"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("0"))),
				Name:     "Cheers UP",
				Symbol:   "CUP",
				Standard: metadata.StandardERC721,
				URI:      "ipfs://QmR4fuz6w9oKEo6oqwFdTmuXqWwmrsFwv659tSZr1SJiNR",
			},
		},
		{
			name: "RSS3 Whitepaper #1",
			arguments: arguments{
				network: filter.NetworkEthereum,
				address: lo.ToPtr(common.HexToAddress("0xB9619cF4F875CdF0E3CE48B28A1C725bC4f6c0fB")),
				id:      big.NewInt(1),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0xB9619cF4F875CdF0E3CE48B28A1C725bC4f6c0fB"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
				Name:     "RSS3 Whitepaper",
				Symbol:   "RWP",
				Standard: metadata.StandardERC721,
				URI:      "ipfs://QmTMD6sLA7M4iegKDhbdMPBZ4HLi5fjW27w2J16gqc5Cb7/1.json",
			},
		},
		{
			name: "OpenSea Shared Storefront #4452815761501376758038898720702591022279500679302039557361006834352129",
			arguments: arguments{
				network: filter.NetworkEthereum,
				address: lo.ToPtr(common.HexToAddress("0x495f947276749Ce646f68AC8c248420045cb7b5e")),
				id:      lo.Must(new(big.Int).SetString("4452815761501376758038898720702591022279500679302039557361006834352129", 0)),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0x495f947276749Ce646f68AC8c248420045cb7b5e"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("4452815761501376758038898720702591022279500679302039557361006834352129"))),
				Name:     "OpenSea Shared Storefront",
				Symbol:   "OPENSTORE",
				Standard: metadata.StandardERC1155,
				URI:      "https://api.opensea.io/api/v1/metadata/0x495f947276749Ce646f68AC8c248420045cb7b5e/0x{id}",
			},
		},
		{
			name: "Love, Death + Robots",
			arguments: arguments{
				network: filter.NetworkEthereum,
				address: lo.ToPtr(common.HexToAddress("0xFD43D1dA000558473822302e1d44D81dA2e4cC0d")),
				id:      big.NewInt(1),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0xFD43D1dA000558473822302e1d44D81dA2e4cC0d"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
				Name:     "Love, Death + Robots",
				Symbol:   "LDR",
				Standard: metadata.StandardERC1155,
				URI:      "ipfs://QmNjmcjL2cz1LrHmTXy3CpVifRNoy3TDh6duo7jxkFFBZH/{id}",
			},
		},
		{
			name: "RSS3 Testnet Token",
			arguments: arguments{
				network: filter.NetworkRSS3Testnet,
				address: lo.ToPtr(common.HexToAddress("0x05BF18310a20FBAeBA376282B5FC6CC0A404402B")),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0x05BF18310a20FBAeBA376282B5FC6CC0A404402B"),
				Name:     "aaa",
				Symbol:   "a",
				Decimals: 18,
				Standard: metadata.StandardERC20,
			},
		},
		{
			name: "Kiwi #979",
			arguments: arguments{
				network: filter.NetworkOptimism,
				address: lo.ToPtr(common.HexToAddress("0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45")),
				id:      big.NewInt(979),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("979"))),
				Name:     "Kiwi Pass",
				Symbol:   "$KIWI",
				Standard: metadata.StandardERC721,
				URI:      "data:application/json;base64,eyJuYW1lIjogIktpd2kgUGFzcyA5NzkiLCAiZGVzY3JpcHRpb24iOiAiS2l3aSBQYXNzIG9wZW5zIHRoZSBkb29yIHRvIHRoZSBjb21tdW5pdHkgb2YgY3VyYXRvcnMgd2hvIHNoYXBlIHRoZSBLaXdpIGZlZWQuXG5cbkdvIHRvIGtpd2luZXdzLnh5eiBmb3IgeW91ciBkYWlseSBkb3NlIG9mIGtpd2kg8J+lnVxuXG5EbyBOT1QgYnV5IEtpd2kgUGFzc2VzIG9uIHRoZSBzZWNvbmRhcnkgbWFya2V0LiBUaGV5IHdpbGwgTk9UIHdvcmsuIiwgImltYWdlIjogImlwZnM6Ly9iYWZ5YmVpYzJhcmFqNWtvanVhZ2tqZXJ5aHF5Y2Y3dnZhZmczMmJjdzNsNXUzZ3Q2YXM0MmZseWEyeSIsICJwcm9wZXJ0aWVzIjogeyJudW1iZXIiOiA5NzksICJuYW1lIjogIktpd2kgUGFzcyJ9fQ==",
			},
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
			defer cancel()

			ethereumClient, err := ethereum.Dial(ctx, endpoint.MustGet(testcase.arguments.network))
			require.NoError(t, err)

			tokenClient := token.NewClient(ethereumClient)

			chainID, err := filter.EthereumChainIDString(testcase.arguments.network.String())
			require.NoError(t, err)

			tokenMetadata, err := tokenClient.Lookup(ctx, uint64(chainID), testcase.arguments.address, testcase.arguments.id, nil)
			require.NoError(t, err)

			require.Equal(t, testcase.want, *tokenMetadata)
		})
	}
}
