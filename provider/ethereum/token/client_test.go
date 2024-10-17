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
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestEthereumClient(t *testing.T) {
	t.Parallel()

	type arguments struct {
		network network.Network
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
				network: network.Ethereum,
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
				network: network.Ethereum,
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
				network: network.Ethereum,
				address: lo.ToPtr(common.HexToAddress("0xAcbe98EFe2d4D103e221E04c76D7c55dB15C8E89")),
				id:      big.NewInt(61),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0xAcbe98EFe2d4D103e221E04c76D7c55dB15C8E89"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("61"))),
				Name:     "RSS Fruits Token",
				Symbol:   "RFT",
				Standard: metadata.StandardERC721,
			},
		},
		{
			name: "Base, Introduced #1",
			arguments: arguments{
				network: network.Ethereum,
				address: lo.ToPtr(common.HexToAddress("0xD4307E0acD12CF46fD6cf93BC264f5D5D1598792")),
				id:      big.NewInt(1),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0xD4307E0acD12CF46fD6cf93BC264f5D5D1598792"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
				Name:     "Base, Introduced",
				Symbol:   "BASEINTRODUCED",
				Standard: metadata.StandardERC721,
			},
		},
		{
			name: "Cheers UP #0",
			arguments: arguments{
				network: network.Ethereum,
				address: lo.ToPtr(common.HexToAddress("0x3113A3c04aEBEC2B77eB38Eabf6a2257B580c54B")),
				id:      big.NewInt(0),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0x3113A3c04aEBEC2B77eB38Eabf6a2257B580c54B"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("0"))),
				Name:     "Cheers UP",
				Symbol:   "CUP",
				Standard: metadata.StandardERC721,
			},
		},
		{
			name: "RSS3 Whitepaper #1",
			arguments: arguments{
				network: network.Ethereum,
				address: lo.ToPtr(common.HexToAddress("0xB9619cF4F875CdF0E3CE48B28A1C725bC4f6c0fB")),
				id:      big.NewInt(1),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0xB9619cF4F875CdF0E3CE48B28A1C725bC4f6c0fB"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
				Name:     "RSS3 Whitepaper",
				Symbol:   "RWP",
				Standard: metadata.StandardERC721,
			},
		},
		{
			name: "OpenSea Shared Storefront #4452815761501376758038898720702591022279500679302039557361006834352129",
			arguments: arguments{
				network: network.Ethereum,
				address: lo.ToPtr(common.HexToAddress("0x495f947276749Ce646f68AC8c248420045cb7b5e")),
				id:      lo.Must(new(big.Int).SetString("4452815761501376758038898720702591022279500679302039557361006834352129", 0)),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0x495f947276749Ce646f68AC8c248420045cb7b5e"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("4452815761501376758038898720702591022279500679302039557361006834352129"))),
				Name:     "OpenSea Shared Storefront",
				Symbol:   "OPENSTORE",
				Standard: metadata.StandardERC1155,
			},
		},
		{
			name: "Love, Death + Robots",
			arguments: arguments{
				network: network.Ethereum,
				address: lo.ToPtr(common.HexToAddress("0xFD43D1dA000558473822302e1d44D81dA2e4cC0d")),
				id:      big.NewInt(1),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0xFD43D1dA000558473822302e1d44D81dA2e4cC0d"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
				Name:     "Love, Death + Robots",
				Symbol:   "LDR",
				Standard: metadata.StandardERC1155,
			},
		},
		{
			name: "Kiwi #979",
			arguments: arguments{
				network: network.Optimism,
				address: lo.ToPtr(common.HexToAddress("0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45")),
				id:      big.NewInt(979),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45"),
				ID:       lo.ToPtr(lo.Must(decimal.NewFromString("979"))),
				Name:     "Kiwi Pass",
				Symbol:   "$KIWI",
				Standard: metadata.StandardERC721,
			},
		},
		{
			name: "RSS3 Token",
			arguments: arguments{
				network: network.VSL,
				address: lo.ToPtr(common.HexToAddress("0x4200000000000000000000000000000000000042")),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0x4200000000000000000000000000000000000042"),
				Name:     "RSS3",
				Symbol:   "RSS3",
				Decimals: 18,
				Standard: metadata.StandardERC20,
			},
		},
		{
			name: "RSS3 VSL Open Chips",
			arguments: arguments{
				network: network.VSL,
				address: lo.ToPtr(common.HexToAddress("0x849f8F55078dCc69dD857b58Cc04631EBA54E4DE")),
			},
			want: metadata.Token{
				Address:  lo.ToPtr("0x849f8F55078dCc69dD857b58Cc04631EBA54E4DE"),
				Name:     "Open Chips",
				Symbol:   "Chips",
				Standard: metadata.StandardERC721,
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

			chainID, err := network.EthereumChainIDString(testcase.arguments.network.String())
			require.NoError(t, err)

			tokenMetadata, err := tokenClient.Lookup(ctx, uint64(chainID), testcase.arguments.address, testcase.arguments.id, nil)
			require.NoError(t, err)

			require.Equal(t, testcase.want, *tokenMetadata)
		})
	}
}
