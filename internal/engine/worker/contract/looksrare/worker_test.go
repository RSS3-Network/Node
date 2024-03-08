package looksrare_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/contract/looksrare"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestWorker_Ethereum(t *testing.T) {
	t.Parallel()

	type arguments struct {
		task   *source.Task
		config *config.Module
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      *schema.Feed
		wantError require.ErrorAssertionFunc
	}{
		// V1
		{
			name: "Looksrare V1 Buy",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xbb7ed57364331776fe5b30600e7e83fe2aa17721cdf841e77223f59f261014e8"),
						ParentHash:   common.HexToHash("0x5a234165e36cc8a0838bee87be9ed0190f46eb8ff87a4b48f146504bde918489"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x51a1449b3B6D635EddeC781cD47a99221712De97"),
						Number:       lo.Must(new(big.Int).SetString("16529836", 0)),
						GasLimit:     30000000,
						GasUsed:      14356372,
						Timestamp:    1675204655,
						BaseFee:      lo.Must(new(big.Int).SetString("42724603776", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xceac035ba1023ba84b1c3a457170bf6dd3fe892e9718699eb88c4624efd4625c"),
						From:      common.HexToAddress("0x71D1988C74A2321a4e71B99490a4D61c72Fe2C96"),
						Gas:       340726,
						GasPrice:  lo.Must(new(big.Int).SetString("44224603776", 10)),
						Hash:      common.HexToHash("0xceac035ba1023ba84b1c3a457170bf6dd3fe892e9718699eb88c4624efd4625c"),
						Input:     hexutil.MustDecode("0xb4e4b29600000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000000000000000000000000000000071d1988c74a2321a4e71b99490a4d61c72fe2c96000000000000000000000000000000000000000000000000075762b27e2e80000000000000000000000000000000000000000000000000000000000000001677000000000000000000000000000000000000000000000000000000000000264800000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000df2c823e219a0c50b8fda1338936334003676811000000000000000000000000960b7a6bcd451c9968473f7bbfd9be826efd549a000000000000000000000000000000000000000000000000075762b27e2e800000000000000000000000000000000000000000000000000000000000000016770000000000000000000000000000000000000000000000000000000000000001000000000000000000000000579af6fd30bf83a5ac0d636bc619f98dbdeb930c000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000000000000000000000000000000000000000000b70000000000000000000000000000000000000000000000000000000063d9787b0000000000000000000000000000000000000000000000000000000063ddb6f100000000000000000000000000000000000000000000000000000000000026480000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001b342dc0cc2dc61ecb7eae231949f31736a9dc48974b2ddc66fda85099a2e20fb847320a93db917d8a91bb6192356b6c74eea18ae6184ab8fe2770662e14a8ed650000000000000000000000000000000000000000000000000000000000000000332d1229"),
						To:        lo.ToPtr(common.HexToAddress("0x59728544B08AB483533076417FbBB2fD0B17CE3a")),
						Value:     lo.Must(new(big.Int).SetString("529000000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xbb7ed57364331776fe5b30600e7e83fe2aa17721cdf841e77223f59f261014e8"),
						BlockNumber:       lo.Must(new(big.Int).SetString("16529836", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 4798076,
						EffectiveGasPrice: hexutil.MustDecodeBig("0xa4bfde680"),
						GasUsed:           238865,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"),
								common.HexToHash("0x00000000000000000000000059728544b08ab483533076417fbbb2fd0b17ce3a"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000075762b27e2e8000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16529836", 0)),
							TransactionHash: common.HexToHash("0xceac035ba1023ba84b1c3a457170bf6dd3fe892e9718699eb88c4624efd4625c"),
							Index:           105,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000059728544b08ab483533076417fbbb2fd0b17ce3a"),
								common.HexToHash("0x0000000000000000000000005924a28caaf1cc016617874a2f0c3710d881f3c1"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000001c30d7284af000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16529836", 0)),
							TransactionHash: common.HexToHash("0xceac035ba1023ba84b1c3a457170bf6dd3fe892e9718699eb88c4624efd4625c"),
							Index:           106,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000059728544b08ab483533076417fbbb2fd0b17ce3a"),
								common.HexToHash("0x000000000000000000000000798a07190b529a7beaa0b64f86865dede0f33500"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000009659d0d6e5000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16529836", 0)),
							TransactionHash: common.HexToHash("0xceac035ba1023ba84b1c3a457170bf6dd3fe892e9718699eb88c4624efd4625c"),
							Index:           107,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x59728544B08AB483533076417FbBB2fD0B17CE3a"),
							Topics: []common.Hash{
								common.HexToHash("0x27c4f0403323142b599832f26acd21c74a9e5b809f2215726e244a4ac588cd7d"),
								common.HexToHash("0x000000000000000000000000960b7a6bcd451c9968473f7bbfd9be826efd549a"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000001677"),
								common.HexToHash("0x000000000000000000000000798a07190b529a7beaa0b64f86865dede0f33500"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000009659d0d6e5000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16529836", 0)),
							TransactionHash: common.HexToHash("0xceac035ba1023ba84b1c3a457170bf6dd3fe892e9718699eb88c4624efd4625c"),
							Index:           108,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000059728544b08ab483533076417fbbb2fd0b17ce3a"),
								common.HexToHash("0x000000000000000000000000df2c823e219a0c50b8fda1338936334003676811"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000731cc3e48754000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16529836", 0)),
							TransactionHash: common.HexToHash("0xceac035ba1023ba84b1c3a457170bf6dd3fe892e9718699eb88c4624efd4625c"),
							Index:           109,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x960b7a6BCD451c9968473f7bbFd9Be826EFd549A"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000df2c823e219a0c50b8fda1338936334003676811"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000001677"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("16529836", 0)),
							TransactionHash: common.HexToHash("0xceac035ba1023ba84b1c3a457170bf6dd3fe892e9718699eb88c4624efd4625c"),
							Index:           110,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x960b7a6BCD451c9968473f7bbFd9Be826EFd549A"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000df2c823e219a0c50b8fda1338936334003676811"),
								common.HexToHash("0x00000000000000000000000071d1988c74a2321a4e71b99490a4d61c72fe2c96"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000001677"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("16529836", 0)),
							TransactionHash: common.HexToHash("0xceac035ba1023ba84b1c3a457170bf6dd3fe892e9718699eb88c4624efd4625c"),
							Index:           111,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x59728544B08AB483533076417FbBB2fD0B17CE3a"),
							Topics: []common.Hash{
								common.HexToHash("0x95fb6205e23ff6bda16a2d1dba56b9ad7c783f67c96fa149785052f47696f2be"),
								common.HexToHash("0x00000000000000000000000071d1988c74a2321a4e71b99490a4d61c72fe2c96"),
								common.HexToHash("0x000000000000000000000000df2c823e219a0c50b8fda1338936334003676811"),
								common.HexToHash("0x000000000000000000000000579af6fd30bf83a5ac0d636bc619f98dbdeb930c"),
							},
							Data:            hexutil.MustDecode("0x55e0053974972ffd24f690427a5ba27b23b5ebf9ad20cdf53347a9e14954306d00000000000000000000000000000000000000000000000000000000000000b7000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000960b7a6bcd451c9968473f7bbfd9be826efd549a00000000000000000000000000000000000000000000000000000000000016770000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000075762b27e2e8000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16529836", 0)),
							TransactionHash: common.HexToHash("0xceac035ba1023ba84b1c3a457170bf6dd3fe892e9718699eb88c4624efd4625c"),
							Index:           112,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xceac035ba1023ba84b1c3a457170bf6dd3fe892e9718699eb88c4624efd4625c"),
						TransactionIndex: 47,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0xceac035ba1023ba84b1c3a457170bf6dd3fe892e9718699eb88c4624efd4625c",
				Network: filter.NetworkEthereum,
				Index:   47,
				From:    "0x71D1988C74A2321a4e71B99490a4D61c72Fe2C96",
				To:      "0x59728544B08AB483533076417FbBB2fD0B17CE3a",
				Type:    filter.TypeCollectibleTrade,
				Calldata: &schema.Calldata{
					FunctionHash: "0xb4e4b296",
				},
				Platform: lo.ToPtr(filter.PlatformLooksRare),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("10563709980954240")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0xdf2c823E219a0C50b8fdA1338936334003676811",
						To:       "0x5924A28caAF1cc016617874a2f0C3710d881f3c1",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("7935000000000000"))),
							Name:     "Wrapped Ether",
							Symbol:   "WETH",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0xdf2c823E219a0C50b8fdA1338936334003676811",
						To:       "0x798a07190B529a7bEAA0b64F86865dedE0F33500",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2645000000000000"))),
							Name:     "Wrapped Ether",
							Symbol:   "WETH",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
					{
						Type:     filter.TypeCollectibleTrade,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0xdf2c823E219a0C50b8fdA1338936334003676811",
						To:       "0x71D1988C74A2321a4e71B99490a4D61c72Fe2C96",
						Metadata: metadata.CollectibleTrade{
							Action: metadata.ActionCollectibleTradeBuy,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x960b7a6BCD451c9968473f7bbFd9Be826EFd549A"),
								ID:       lo.ToPtr(lo.Must(decimal.NewFromString("5751"))),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
								Name:     "OnChainMonkey",
								Symbol:   "OCMONK",
								Standard: metadata.StandardERC721,
								URI:      "data:application/json;base64,eyJuYW1lIjogIk9uQ2hhaW4gTW9ua2V5ICM1NzUxIiwgImF0dHJpYnV0ZXMiOiBbeyJ0cmFpdF90eXBlIjogIkJhY2tncm91bmQiLCJ2YWx1ZSI6ICIyIn0seyJ0cmFpdF90eXBlIjogIkZ1ciIsInZhbHVlIjogIjMifSx7InRyYWl0X3R5cGUiOiAiRWFycmluZyIsInZhbHVlIjogIjAifSx7InRyYWl0X3R5cGUiOiAiSGF0IiwidmFsdWUiOiAiMjgifSx7InRyYWl0X3R5cGUiOiAiRXllcyIsInZhbHVlIjogIjgifSx7InRyYWl0X3R5cGUiOiAiQ2xvdGhlcyIsInZhbHVlIjogIjgifSx7InRyYWl0X3R5cGUiOiAiTW91dGgiLCJ2YWx1ZSI6ICIxMCJ9XSwiaW1hZ2UiOiAiZGF0YTppbWFnZS9zdmcreG1sO2Jhc2U2NCxQSE4yWnlCNGJXeHVjejBpYUhSMGNEb3ZMM2QzZHk1M015NXZjbWN2TWpBd01DOXpkbWNpSUhCeVpYTmxjblpsUVhOd1pXTjBVbUYwYVc4OUluaE5hVzVaVFdsdUlHMWxaWFFpSUhacFpYZENiM2c5SWpBZ01DQTFNREFnTlRBd0lqNDhjbVZqZENCNFBTSXdJaUI1UFNJd0lpQjNhV1IwYUQwaU5UQXdJaUJvWldsbmFIUTlJalV3TUNJZ2MzUjViR1U5SW1acGJHdzZJMlU1TWlJdlBqeHlaV04wSUhkcFpIUm9QU0l6TURBaUlHaGxhV2RvZEQwaU1USXdJaUI0UFNJNU9TSWdlVDBpTkRBd0lpQnpkSGxzWlQwaVptbHNiRG9qTmpVeklpOCtQR05wY21Oc1pTQmplRDBpTVRrd0lpQmplVDBpTkRjd0lpQnlQU0kxSWlCemRIbHNaVDBpWm1sc2JEb2pZVGN4SWk4K1BHTnBjbU5zWlNCamVEMGlNekV3SWlCamVUMGlORGN3SWlCeVBTSTFJaUJ6ZEhsc1pUMGlabWxzYkRvallUY3hJaTgrUEdOcGNtTnNaU0JqZUQwaU1UQXdJaUJqZVQwaU1qVXdJaUJ5UFNJMU1DSWdjM1I1YkdVOUltWnBiR3c2SXpZMU15SXZQanhqYVhKamJHVWdZM2c5SWpFd01DSWdZM2s5SWpJMU1DSWdjajBpTWpBaUlITjBlV3hsUFNKbWFXeHNPaU5oTnpFaUx6NDhZMmx5WTJ4bElHTjRQU0kwTURBaUlHTjVQU0l5TlRBaUlISTlJalV3SWlCemRIbHNaVDBpWm1sc2JEb2pOalV6SWk4K1BHTnBjbU5zWlNCamVEMGlOREF3SWlCamVUMGlNalV3SWlCeVBTSXlNQ0lnYzNSNWJHVTlJbVpwYkd3NkkyRTNNU0l2UGp4amFYSmpiR1VnWTNnOUlqSTFNQ0lnWTNrOUlqSTFNQ0lnY2owaU1UVXdJaUJ6ZEhsc1pUMGlabWxzYkRvak5qVXpJaTgrUEdOcGNtTnNaU0JqZUQwaU1qVXdJaUJqZVQwaU1qVXdJaUJ5UFNJeE1qQWlJSE4wZVd4bFBTSm1hV3hzT2lOaE56RWlMejQ4WTJseVkyeGxJR040UFNJeU1EQWlJR041UFNJeU1UVWlJSEk5SWpNMUlpQnpkSGxzWlQwaVptbHNiRG9qWm1abUlpOCtQR05wY21Oc1pTQmplRDBpTXpBMUlpQmplVDBpTWpJeUlpQnlQU0l6TVNJZ2MzUjViR1U5SW1acGJHdzZJMlptWmlJdlBqeGphWEpqYkdVZ1kzZzlJakl3TUNJZ1kzazlJakl5TUNJZ2NqMGlNakFpSUhOMGVXeGxQU0ptYVd4c09pTTRPRGdpTHo0OFkybHlZMnhsSUdONFBTSXpNREFpSUdONVBTSXlNakFpSUhJOUlqSXdJaUJ6ZEhsc1pUMGlabWxzYkRvak9EZzRJaTgrUEdOcGNtTnNaU0JqZUQwaU1qQXdJaUJqZVQwaU1qSXdJaUJ5UFNJM0lpQnpkSGxzWlQwaVptbHNiRG9qTURBd0lpOCtQR05wY21Oc1pTQmplRDBpTXpBd0lpQmplVDBpTWpJd0lpQnlQU0kzSWlCemRIbHNaVDBpWm1sc2JEb2pNREF3SWk4K1BISmxZM1FnZUQwaU1UVXdJaUI1UFNJeE9UQWlJSGRwWkhSb1BTSXlNREFpSUdobGFXZG9kRDBpTXpBaUlITjBlV3hsUFNKbWFXeHNPaU5oTnpFaUx6NDhjbVZqZENCNFBTSXhOakFpSUhrOUlqRTNNQ0lnZDJsa2RHZzlJakU0TUNJZ2FHVnBaMmgwUFNJMU1DSWdjM1I1YkdVOUltWnBiR3c2STJFM01TSXZQanhsYkd4cGNITmxJR040UFNJeU5UQWlJR041UFNJek1UVWlJSEo0UFNJNE5DSWdjbms5SWpNMElpQnpkSGxzWlQwaVptbHNiRG9qTmpVeklpOCtQSEpsWTNRZ2VEMGlNVGsxSWlCNVBTSXpNekFpSUhkcFpIUm9QU0l4TVRBaUlHaGxhV2RvZEQwaU15SWdjM1I1YkdVOUltWnBiR3c2SXpBd01DSXZQanhqYVhKamJHVWdZM2c5SWpJMk9DSWdZM2s5SWpJNU5TSWdjajBpTlNJZ2MzUjViR1U5SW1acGJHdzZJekF3TUNJdlBqeGphWEpqYkdVZ1kzZzlJakl6TWlJZ1kzazlJakk1TlNJZ2NqMGlOU0lnYzNSNWJHVTlJbVpwYkd3Nkl6QXdNQ0l2UGp4c2FXNWxJSGd4UFNJeE56VWlJSGt4UFNJek1EY2lJSGd5UFNJeE56VWlJSGt5UFNJek1USWlJSE4wZVd4bFBTSnpkSEp2YTJVNkl6QXdNRHR6ZEhKdmEyVXRkMmxrZEdnNk1pSXZQanhzYVc1bElIZ3hQU0l5TURBaUlIa3hQU0l6TURjaUlIZ3lQU0l5TURBaUlIa3lQU0l6TVRJaUlITjBlV3hsUFNKemRISnZhMlU2SXpBd01EdHpkSEp2YTJVdGQybGtkR2c2TWlJdlBqeHNhVzVsSUhneFBTSXlNalVpSUhreFBTSXpNRGNpSUhneVBTSXlNalVpSUhreVBTSXpNVElpSUhOMGVXeGxQU0p6ZEhKdmEyVTZJekF3TUR0emRISnZhMlV0ZDJsa2RHZzZNaUl2UGp4c2FXNWxJSGd4UFNJeU5UQWlJSGt4UFNJek1EY2lJSGd5UFNJeU5UQWlJSGt5UFNJek1USWlJSE4wZVd4bFBTSnpkSEp2YTJVNkl6QXdNRHR6ZEhKdmEyVXRkMmxrZEdnNk1pSXZQanhzYVc1bElIZ3hQU0l5TnpVaUlIa3hQU0l6TURjaUlIZ3lQU0l5TnpVaUlIa3lQU0l6TVRJaUlITjBlV3hsUFNKemRISnZhMlU2SXpBd01EdHpkSEp2YTJVdGQybGtkR2c2TWlJdlBqeHNhVzVsSUhneFBTSXpNREFpSUhreFBTSXpNRGNpSUhneVBTSXpNREFpSUhreVBTSXpNVElpSUhOMGVXeGxQU0p6ZEhKdmEyVTZJekF3TUR0emRISnZhMlV0ZDJsa2RHZzZNaUl2UGp4c2FXNWxJSGd4UFNJek1qVWlJSGt4UFNJek1EY2lJSGd5UFNJek1qVWlJSGt5UFNJek1USWlJSE4wZVd4bFBTSnpkSEp2YTJVNkl6QXdNRHR6ZEhKdmEyVXRkMmxrZEdnNk1pSXZQanhzYVc1bElIZ3hQU0l4T0RjaUlIa3hQU0l6TVRjaUlIZ3lQU0l4T0RjaUlIa3lQU0l6TWpJaUlITjBlV3hsUFNKemRISnZhMlU2SXpBd01EdHpkSEp2YTJVdGQybGtkR2c2TWlJdlBqeHNhVzVsSUhneFBTSXlNVElpSUhreFBTSXpNVGNpSUhneVBTSXlNVElpSUhreVBTSXpNaklpSUhOMGVXeGxQU0p6ZEhKdmEyVTZJekF3TUR0emRISnZhMlV0ZDJsa2RHZzZNaUl2UGp4c2FXNWxJSGd4UFNJeU16Y2lJSGt4UFNJek1UY2lJSGd5UFNJeU16Y2lJSGt5UFNJek1qSWlJSE4wZVd4bFBTSnpkSEp2YTJVNkl6QXdNRHR6ZEhKdmEyVXRkMmxrZEdnNk1pSXZQanhzYVc1bElIZ3hQU0l5TmpJaUlIa3hQU0l6TVRjaUlIZ3lQU0l5TmpJaUlIa3lQU0l6TWpJaUlITjBlV3hsUFNKemRISnZhMlU2SXpBd01EdHpkSEp2YTJVdGQybGtkR2c2TWlJdlBqeHNhVzVsSUhneFBTSXlPRGNpSUhreFBTSXpNVGNpSUhneVBTSXlPRGNpSUhreVBTSXpNaklpSUhOMGVXeGxQU0p6ZEhKdmEyVTZJekF3TUR0emRISnZhMlV0ZDJsa2RHZzZNaUl2UGp4c2FXNWxJSGd4UFNJek1USWlJSGt4UFNJek1UY2lJSGd5UFNJek1USWlJSGt5UFNJek1qSWlJSE4wZVd4bFBTSnpkSEp2YTJVNkl6QXdNRHR6ZEhKdmEyVXRkMmxrZEdnNk1pSXZQanh5WldOMElIZHBaSFJvUFNJeU1EQWlJR2hsYVdkb2REMGlPVGtpSUhnOUlqRTFNQ0lnZVQwaU5EQWlJSE4wZVd4bFBTSm1hV3hzT2lObU9EQWlMejQ4Y21WamRDQjNhV1IwYUQwaU1qQXdJaUJvWldsbmFIUTlJak16SWlCNFBTSXhOVEFpSUhrOUlqRXdOaUlnYzNSNWJHVTlJbVpwYkd3Nkl6a3daaUl2UGp4eVpXTjBJSGRwWkhSb1BTSXpNREFpSUdobGFXZG9kRDBpTVRJd0lpQjRQU0k1T1NJZ2VUMGlOREF3SWlCemRIbHNaVDBpWm1sc2JEb2paakF3SWk4K1BISmxZM1FnZDJsa2RHZzlJalV3SWlCb1pXbG5hSFE5SWpVMUlpQjRQU0l5T0RBaUlIazlJalF6TUNJZ2MzUjViR1U5SW1acGJHdzZJekJtTUNJdlBqd3ZjM1puUGc9PSJ9",
							},

							Cost: &metadata.Token{
								Address:  lo.ToPtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("529000000000000000"))),
								Name:     "Wrapped Ether",
								Symbol:   "WETH",
								Standard: metadata.StandardERC20,
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1675204655,
			},
			wantError: require.NoError,
		},
		{
			name: "Looksrare V1 Sell",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x12489b5c7be4e12caf2390389a4d6353dcfec647737156c806556468dfc21b4d"),
						ParentHash:   common.HexToHash("0xecff1900f2a824898262193cc7931dcbd42541d5086d3614dde99728f9432b41"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xDAFEA492D9c6733ae3d56b7Ed1ADB60692c98Bc5"),
						Number:       lo.Must(new(big.Int).SetString("16095612", 0)),
						GasLimit:     30000000,
						GasUsed:      16844545,
						Timestamp:    1669965899,
						BaseFee:      lo.Must(new(big.Int).SetString("11731150072", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xdfa1cdcbaf519757c74bd7a01ce3eed846c8866fd944207e277422067e2cdf01"),
						From:      common.HexToAddress("0x53b620A822984dE20Bd147319747b314e2A901bA"),
						Gas:       326284,
						GasPrice:  lo.Must(new(big.Int).SetString("13231150072", 10)),
						Hash:      common.HexToHash("0xdfa1cdcbaf519757c74bd7a01ce3eed846c8866fd944207e277422067e2cdf01"),
						Input:     hexutil.MustDecode("0x3b6d032e00000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000000100000000000000000000000053b620a822984de20bd147319747b314e2a901ba00000000000000000000000000000000000000000000000011c761a995c0c000000000000000000000000000000000000000000000000000000000000000f077000000000000000000000000000000000000000000000000000000000000264800000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000009f28152ae399a6f1b9760c04972de4f9303c0c9f00000000000000000000000034d85c9cdeb23fa97cb08333b511ac86e1c4e25800000000000000000000000000000000000000000000000011c761a995c0c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000009f93623019049c76209c26517acc2af9d49c69b000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000000000000000000000000000000000000000047c0000000000000000000000000000000000000000000000000000000063897fda000000000000000000000000000000000000000000000000000000006389f47800000000000000000000000000000000000000000000000000000000000026480000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001c982af7971d2a160a6ac0fefbb92b418d398d9de370371e53b9f93bffa8d2b61a0cb1e80bf6298e75d9916fe7cbbeab0ae33ad73fdc756dd0b8d0ded49d3755c10000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x59728544B08AB483533076417FbBB2fD0B17CE3a")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x12489b5c7be4e12caf2390389a4d6353dcfec647737156c806556468dfc21b4d"),
						BlockNumber:       lo.Must(new(big.Int).SetString("16095612", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 6816372,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x314a353f8"),
						GasUsed:           238404,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x34d85c9CDeB23FA97cb08333b511ac86E1C4E258"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x00000000000000000000000053b620a822984de20bd147319747b314e2a901ba"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000f077"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("16095612", 0)),
							TransactionHash: common.HexToHash("0xdfa1cdcbaf519757c74bd7a01ce3eed846c8866fd944207e277422067e2cdf01"),
							Index:           123,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x34d85c9CDeB23FA97cb08333b511ac86E1C4E258"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000053b620a822984de20bd147319747b314e2a901ba"),
								common.HexToHash("0x0000000000000000000000009f28152ae399a6f1b9760c04972de4f9303c0c9f"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000f077"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("16095612", 0)),
							TransactionHash: common.HexToHash("0xdfa1cdcbaf519757c74bd7a01ce3eed846c8866fd944207e277422067e2cdf01"),
							Index:           124,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000009f28152ae399a6f1b9760c04972de4f9303c0c9f"),
								common.HexToHash("0x0000000000000000000000005924a28caaf1cc016617874a2f0c3710d881f3c1"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000044454e10538800"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16095612", 0)),
							TransactionHash: common.HexToHash("0xdfa1cdcbaf519757c74bd7a01ce3eed846c8866fd944207e277422067e2cdf01"),
							Index:           125,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000009f28152ae399a6f1b9760c04972de4f9303c0c9f"),
								common.HexToHash("0x00000000000000000000000037ceb4ba093d40234c6fb312d9791b67c04ef49a"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000016c1c4b01bd800"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16095612", 0)),
							TransactionHash: common.HexToHash("0xdfa1cdcbaf519757c74bd7a01ce3eed846c8866fd944207e277422067e2cdf01"),
							Index:           126,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x59728544B08AB483533076417FbBB2fD0B17CE3a"),
							Topics: []common.Hash{
								common.HexToHash("0x27c4f0403323142b599832f26acd21c74a9e5b809f2215726e244a4ac588cd7d"),
								common.HexToHash("0x00000000000000000000000034d85c9cdeb23fa97cb08333b511ac86e1c4e258"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000f077"),
								common.HexToHash("0x00000000000000000000000037ceb4ba093d40234c6fb312d9791b67c04ef49a"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000016c1c4b01bd800"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16095612", 0)),
							TransactionHash: common.HexToHash("0xdfa1cdcbaf519757c74bd7a01ce3eed846c8866fd944207e277422067e2cdf01"),
							Index:           127,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000009f28152ae399a6f1b9760c04972de4f9303c0c9f"),
								common.HexToHash("0x00000000000000000000000053b620a822984de20bd147319747b314e2a901ba"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000116c5a96d5516000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16095612", 0)),
							TransactionHash: common.HexToHash("0xdfa1cdcbaf519757c74bd7a01ce3eed846c8866fd944207e277422067e2cdf01"),
							Index:           128,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x59728544B08AB483533076417FbBB2fD0B17CE3a"),
							Topics: []common.Hash{
								common.HexToHash("0x68cd251d4d267c6e2034ff0088b990352b97b2002c0476587d0c4da889c11330"),
								common.HexToHash("0x00000000000000000000000053b620a822984de20bd147319747b314e2a901ba"),
								common.HexToHash("0x0000000000000000000000009f28152ae399a6f1b9760c04972de4f9303c0c9f"),
								common.HexToHash("0x00000000000000000000000009f93623019049c76209c26517acc2af9d49c69b"),
							},
							Data:            hexutil.MustDecode("0xf3545ac88958479d46b913dde2773410400a4193e70a75b1325ae85fbd104bc4000000000000000000000000000000000000000000000000000000000000047c000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000034d85c9cdeb23fa97cb08333b511ac86e1c4e258000000000000000000000000000000000000000000000000000000000000f077000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000011c761a995c0c000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16095612", 0)),
							TransactionHash: common.HexToHash("0xdfa1cdcbaf519757c74bd7a01ce3eed846c8866fd944207e277422067e2cdf01"),
							Index:           129,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xdfa1cdcbaf519757c74bd7a01ce3eed846c8866fd944207e277422067e2cdf01"),
						TransactionIndex: 71,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0xdfa1cdcbaf519757c74bd7a01ce3eed846c8866fd944207e277422067e2cdf01",
				Network: filter.NetworkEthereum,
				Index:   71,
				From:    "0x53b620A822984dE20Bd147319747b314e2A901bA",
				To:      "0x59728544B08AB483533076417FbBB2fD0B17CE3a",
				Type:    filter.TypeCollectibleTrade,
				Calldata: &schema.Calldata{
					FunctionHash: "0x3b6d032e",
				},
				Platform: lo.ToPtr(filter.PlatformLooksRare),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("3154359101765088")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0x9f28152AE399a6f1B9760c04972DE4f9303C0C9f",
						To:       "0x5924A28caAF1cc016617874a2f0C3710d881f3c1",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("19216500000000000"))),
							Name:     "Wrapped Ether",
							Symbol:   "WETH",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0x9f28152AE399a6f1B9760c04972DE4f9303C0C9f",
						To:       "0x37ceB4bA093D40234c6fB312d9791B67c04eF49A",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("6405500000000000"))),
							Name:     "Wrapped Ether",
							Symbol:   "WETH",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
					{
						Type:     filter.TypeCollectibleTrade,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0x53b620A822984dE20Bd147319747b314e2A901bA",
						To:       "0x9f28152AE399a6f1B9760c04972DE4f9303C0C9f",
						Metadata: metadata.CollectibleTrade{
							Action: metadata.ActionCollectibleTradeSell,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x34d85c9CDeB23FA97cb08333b511ac86E1C4E258"),
								ID:       lo.ToPtr(lo.Must(decimal.NewFromString("61559"))),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
								Name:     "Otherdeed",
								Symbol:   "OTHR",
								Standard: metadata.StandardERC721,
								URI:      "https://api.otherside.xyz/lands/61559",
							},
							Cost: &metadata.Token{
								Address:  lo.ToPtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1281100000000000000"))),
								Name:     "Wrapped Ether",
								Symbol:   "WETH",
								Standard: metadata.StandardERC20,
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1669965899,
			},
			wantError: require.NoError,
		},
		{
			name: "Looksrare Aggregator Execute V1",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x1aab162cd63a1fcde8730d47f17973f68b00acd58002b035aee6db68707a2a8f"),
						ParentHash:   common.HexToHash("0x7fecd934c6faf6d0d46056e96d8dca5a0cee4dbdd1411ab69ae6769d13bbc5f4"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x3b64216AD1a58f61538b4fA1B27327675Ab7ED67"),
						Number:       lo.Must(new(big.Int).SetString("16912582", 0)),
						GasLimit:     30000000,
						GasUsed:      12927872,
						Timestamp:    1679845571,
						BaseFee:      lo.Must(new(big.Int).SetString("21892971416", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x6d18b33acdbfba717ffe8c11831cbe7fda5a1dd2b7e90f6d2ae3af1b503fac1d"),
						From:      common.HexToAddress("0x6d0267156f1c6CE44Caa4BF129B76009d3d41830"),
						Gas:       400000,
						GasPrice:  lo.Must(new(big.Int).SetString("21992971416", 10)),
						Hash:      common.HexToHash("0x6d18b33acdbfba717ffe8c11831cbe7fda5a1dd2b7e90f6d2ae3af1b503fac1d"),
						Input:     hexutil.MustDecode("0x12abee1200000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000006d0267156f1c6ce44caa4bf129b76009d3d418300000000000000000000000006d0267156f1c6ce44caa4bf129b76009d3d4183000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000da151039ed034d1c5bacb47c284ed1a809ce350000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000032000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000017a511c56473690f7365097726874f8751f5e70400000000000000000000000034d85c9cdeb23fa97cb08333b511ac86e1c4e25800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000016345785d8a00000000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000000000064206870000000000000000000000000000000000000000000000000000000006421b9e700000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000013d760000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000419f86054b8ec333376e9401982950f92b9b25f72daa3cf4ea3deff3213aa9e1510538840b98db4191f879e1bc9a4da85315f9a2697decf5e1297d71adcf040bfa1c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000016345785d8a000000000000000000000000000000000000000000000000000000000000000002648000000000000000000000000000000000000000000000000000000000000004d000000000000000000000000579af6fd30bf83a5ac0d636bc619f98dbdeb930c00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x00000000005228B791a99a61f36A130d50600106")),
						Value:     lo.Must(new(big.Int).SetString("1600000000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x1aab162cd63a1fcde8730d47f17973f68b00acd58002b035aee6db68707a2a8f"),
						BlockNumber:       lo.Must(new(big.Int).SetString("16912582", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 2515689,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x51ee21c98"),
						GasUsed:           285013,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"),
								common.HexToHash("0x00000000000000000000000059728544b08ab483533076417fbbb2fd0b17ce3a"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000016345785d8a00000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16912582", 0)),
							TransactionHash: common.HexToHash("0x6d18b33acdbfba717ffe8c11831cbe7fda5a1dd2b7e90f6d2ae3af1b503fac1d"),
							Index:           48,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000059728544b08ab483533076417fbbb2fd0b17ce3a"),
								common.HexToHash("0x0000000000000000000000005924a28caaf1cc016617874a2f0c3710d881f3c1"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000005543df729c0000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16912582", 0)),
							TransactionHash: common.HexToHash("0x6d18b33acdbfba717ffe8c11831cbe7fda5a1dd2b7e90f6d2ae3af1b503fac1d"),
							Index:           49,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000059728544b08ab483533076417fbbb2fd0b17ce3a"),
								common.HexToHash("0x00000000000000000000000037ceb4ba093d40234c6fb312d9791b67c04ef49a"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000001c6bf526340000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16912582", 0)),
							TransactionHash: common.HexToHash("0x6d18b33acdbfba717ffe8c11831cbe7fda5a1dd2b7e90f6d2ae3af1b503fac1d"),
							Index:           50,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x59728544B08AB483533076417FbBB2fD0B17CE3a"),
							Topics: []common.Hash{
								common.HexToHash("0x27c4f0403323142b599832f26acd21c74a9e5b809f2215726e244a4ac588cd7d"),
								common.HexToHash("0x00000000000000000000000034d85c9cdeb23fa97cb08333b511ac86e1c4e258"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000013d76"),
								common.HexToHash("0x00000000000000000000000037ceb4ba093d40234c6fb312d9791b67c04ef49a"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000000000000000000000000000001c6bf526340000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16912582", 0)),
							TransactionHash: common.HexToHash("0x6d18b33acdbfba717ffe8c11831cbe7fda5a1dd2b7e90f6d2ae3af1b503fac1d"),
							Index:           51,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000059728544b08ab483533076417fbbb2fd0b17ce3a"),
								common.HexToHash("0x00000000000000000000000017a511c56473690f7365097726874f8751f5e704"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000015c2a7b13fd00000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16912582", 0)),
							TransactionHash: common.HexToHash("0x6d18b33acdbfba717ffe8c11831cbe7fda5a1dd2b7e90f6d2ae3af1b503fac1d"),
							Index:           52,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x34d85c9CDeB23FA97cb08333b511ac86E1C4E258"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x00000000000000000000000017a511c56473690f7365097726874f8751f5e704"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000013d76"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("16912582", 0)),
							TransactionHash: common.HexToHash("0x6d18b33acdbfba717ffe8c11831cbe7fda5a1dd2b7e90f6d2ae3af1b503fac1d"),
							Index:           53,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x34d85c9CDeB23FA97cb08333b511ac86E1C4E258"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000017a511c56473690f7365097726874f8751f5e704"),
								common.HexToHash("0x00000000000000000000000000000000005228b791a99a61f36a130d50600106"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000013d76"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("16912582", 0)),
							TransactionHash: common.HexToHash("0x6d18b33acdbfba717ffe8c11831cbe7fda5a1dd2b7e90f6d2ae3af1b503fac1d"),
							Index:           54,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x59728544B08AB483533076417FbBB2fD0B17CE3a"),
							Topics: []common.Hash{
								common.HexToHash("0x95fb6205e23ff6bda16a2d1dba56b9ad7c783f67c96fa149785052f47696f2be"),
								common.HexToHash("0x00000000000000000000000000000000005228b791a99a61f36a130d50600106"),
								common.HexToHash("0x00000000000000000000000017a511c56473690f7365097726874f8751f5e704"),
								common.HexToHash("0x000000000000000000000000579af6fd30bf83a5ac0d636bc619f98dbdeb930c"),
							},
							Data:            hexutil.MustDecode("0x52b4dde7e9ba363e873743b98d873f6d8c837e91570f02824a1ab4eed0ebe257000000000000000000000000000000000000000000000000000000000000004d000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000034d85c9cdeb23fa97cb08333b511ac86e1c4e2580000000000000000000000000000000000000000000000000000000000013d76000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000016345785d8a00000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16912582", 0)),
							TransactionHash: common.HexToHash("0x6d18b33acdbfba717ffe8c11831cbe7fda5a1dd2b7e90f6d2ae3af1b503fac1d"),
							Index:           55,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x34d85c9CDeB23FA97cb08333b511ac86E1C4E258"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x00000000000000000000000000000000005228b791a99a61f36a130d50600106"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000013d76"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("16912582", 0)),
							TransactionHash: common.HexToHash("0x6d18b33acdbfba717ffe8c11831cbe7fda5a1dd2b7e90f6d2ae3af1b503fac1d"),
							Index:           56,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x34d85c9CDeB23FA97cb08333b511ac86E1C4E258"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000000000000005228b791a99a61f36a130d50600106"),
								common.HexToHash("0x0000000000000000000000006d0267156f1c6ce44caa4bf129b76009d3d41830"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000013d76"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("16912582", 0)),
							TransactionHash: common.HexToHash("0x6d18b33acdbfba717ffe8c11831cbe7fda5a1dd2b7e90f6d2ae3af1b503fac1d"),
							Index:           57,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x00000000005228B791a99a61f36A130d50600106"),
							Topics: []common.Hash{
								common.HexToHash("0x807273efecfbeb7ae7d3a2189d1ed5a7db80074eed86e7d80b10bb925cd1db73"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000006d0267156f1c6ce44caa4bf129b76009d3d41830"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16912582", 0)),
							TransactionHash: common.HexToHash("0x6d18b33acdbfba717ffe8c11831cbe7fda5a1dd2b7e90f6d2ae3af1b503fac1d"),
							Index:           58,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x6d18b33acdbfba717ffe8c11831cbe7fda5a1dd2b7e90f6d2ae3af1b503fac1d"),
						TransactionIndex: 31,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x6d18b33acdbfba717ffe8c11831cbe7fda5a1dd2b7e90f6d2ae3af1b503fac1d",
				Network: filter.NetworkEthereum,
				Index:   31,
				From:    "0x6d0267156f1c6CE44Caa4BF129B76009d3d41830",
				To:      "0x00000000005228B791a99a61f36A130d50600106",
				Type:    filter.TypeCollectibleTrade,
				Calldata: &schema.Calldata{
					FunctionHash: "0x12abee12",
				},
				Platform: lo.ToPtr(filter.PlatformLooksRare),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("6268282762188408")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0x17a511C56473690f7365097726874F8751F5E704",
						To:       "0x5924A28caAF1cc016617874a2f0C3710d881f3c1",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("24000000000000000"))),
							Name:     "Wrapped Ether",
							Symbol:   "WETH",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0x17a511C56473690f7365097726874F8751F5E704",
						To:       "0x37ceB4bA093D40234c6fB312d9791B67c04eF49A",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("8000000000000000"))),
							Name:     "Wrapped Ether",
							Symbol:   "WETH",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
					{
						Type:     filter.TypeCollectibleTrade,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0x17a511C56473690f7365097726874F8751F5E704",
						To:       "0x6d0267156f1c6CE44Caa4BF129B76009d3d41830",
						Metadata: metadata.CollectibleTrade{
							Action: metadata.ActionCollectibleTradeBuy,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x34d85c9CDeB23FA97cb08333b511ac86E1C4E258"),
								ID:       lo.ToPtr(lo.Must(decimal.NewFromString("81270"))),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
								Name:     "Otherdeed",
								Symbol:   "OTHR",
								Standard: metadata.StandardERC721,
								URI:      "https://api.otherside.xyz/lands/81270",
							},
							Cost: &metadata.Token{
								Address:  lo.ToPtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1600000000000000000"))),
								Name:     "Wrapped Ether",
								Symbol:   "WETH",
								Standard: metadata.StandardERC20,
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1679845571,
			},
			wantError: require.NoError,
		},

		//	V2
		{
			name: "Looksrare V2 Buy",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xe73fdf68e9a0ee1c86440a52980f3139cb03af51d3d75fddf88f1e07d894086e"),
						ParentHash:   common.HexToHash("0x6e2ccbf0cac76c31c4b89ac021af382e97b0059b50d70dfab06bb1874ca021c0"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"),
						Number:       lo.Must(new(big.Int).SetString("18282662", 0)),
						GasLimit:     30000000,
						GasUsed:      11775095,
						Timestamp:    1696488995,
						BaseFee:      lo.Must(new(big.Int).SetString("6045650961", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x15713058474370748513e4499f5707f1477f3fbd1a8b59e569877a6445c3ac74"),
						From:      common.HexToAddress("0x60f0A4747900A24618d0864D06b10F11d89b1d8e"),
						Gas:       224041,
						GasPrice:  lo.Must(new(big.Int).SetString("6145650961", 10)),
						Hash:      common.HexToHash("0x15713058474370748513e4499f5707f1477f3fbd1a8b59e569877a6445c3ac74"),
						Input:     hexutil.MustDecode("0x8585ae0300000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000003800000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000060f0a4747900a24618d0864d06b10f11d89b1d8e00000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d4296688906664c2c2329d2198a8716a2b8e6661000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006b6422ca906a38159cd741f965f6f4e5d05861800000000000000000000000000000000000000000000000000000000651bbf5c0000000000000000000000000000000000000000000000000000000065435a5d000000000000000000000000000000000000000000000000000690e1f43ddc0000000000000000000000000000000000000000000000000000000000000001e0000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000002600000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000067e00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000411d910bee43117ab1b6f48f4fc48606b25a97d44fb14c66fd3c9177b2e6110ccc7feb7a7932fe5e925d123fe590c31fb0d01ddeff52fbc328a91475be52ba4e081c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000072db8c0b"),
						To:        lo.ToPtr(common.HexToAddress("0x0000000000E655fAe4d56241588680F86E3b2377")),
						Value:     lo.Must(new(big.Int).SetString("1848150000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xe73fdf68e9a0ee1c86440a52980f3139cb03af51d3d75fddf88f1e07d894086e"),
						BlockNumber:       lo.Must(new(big.Int).SetString("18282662", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 6255041,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x16e4f3111"),
						GasUsed:           197237,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xd4296688906664c2C2329d2198A8716A2b8e6661"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000006b6422ca906a38159cd741f965f6f4e5d058618"),
								common.HexToHash("0x00000000000000000000000060f0a4747900a24618d0864d06b10f11d89b1d8e"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000067e"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("18282662", 0)),
							TransactionHash: common.HexToHash("0x15713058474370748513e4499f5707f1477f3fbd1a8b59e569877a6445c3ac74"),
							Index:           144,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x0000000000E655fAe4d56241588680F86E3b2377"),
							Topics: []common.Hash{
								common.HexToHash("0x3ee3de4684413690dee6fff1a0a4f92916a1b97d1c5a83cdf24671844306b2e3"),
							},
							Data:            hexutil.MustDecode("0x7ee6b938a1e3af6b27ce50166de238ce3162b4d583fe6ad73f97d20612ff664b000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000000100000000000000000000000060f0a4747900a24618d0864d06b10f11d89b1d8e00000000000000000000000060f0a4747900a24618d0864d06b10f11d89b1d8e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d4296688906664c2c2329d2198a8716a2b8e666100000000000000000000000000000000000000000000000000000000000001e0000000000000000000000000000000000000000000000000000000000000022000000000000000000000000006b6422ca906a38159cd741f965f6f4e5d05861800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006887a6c9ed480000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000867879f07800000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000067e00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001"),
							BlockNumber:     lo.Must(new(big.Int).SetString("18282662", 0)),
							TransactionHash: common.HexToHash("0x15713058474370748513e4499f5707f1477f3fbd1a8b59e569877a6445c3ac74"),
							Index:           145,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x15713058474370748513e4499f5707f1477f3fbd1a8b59e569877a6445c3ac74"),
						TransactionIndex: 42,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x15713058474370748513e4499f5707f1477f3fbd1a8b59e569877a6445c3ac74",
				Network: filter.NetworkEthereum,
				Index:   42,
				From:    "0x60f0A4747900A24618d0864D06b10F11d89b1d8e",
				To:      "0x0000000000E655fAe4d56241588680F86E3b2377",
				Type:    filter.TypeCollectibleTrade,
				Calldata: &schema.Calldata{
					FunctionHash: "0x8585ae03",
				},
				Platform: lo.ToPtr(filter.PlatformLooksRare),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("1212149758594757")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeCollectibleTrade,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0x06B6422CA906A38159cD741F965f6f4E5d058618",
						To:       "0x60f0A4747900A24618d0864D06b10F11d89b1d8e",
						Metadata: metadata.CollectibleTrade{
							Action: metadata.ActionCollectibleTradeBuy,
							Token: metadata.Token{
								Address:  lo.ToPtr("0xd4296688906664c2C2329d2198A8716A2b8e6661"),
								ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1662"))),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
								Name:     "Redacted Remilio Party",
								Symbol:   "RRP",
								Standard: metadata.StandardERC721,
								URI:      "ipfs://bafybeicbgyhvddvn2jdqwvlavszpbk2rek62g6ah4zoxmbsijlr47imcfi/1662",
							},

							Cost: &metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1848150000000000"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0x60f0A4747900A24618d0864D06b10F11d89b1d8e",
						To:       "0x1838De7d4e4e42c8eB7b204A91e28E9fad14F536",
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("9240750000000"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
				},
				Status:    true,
				Timestamp: 1696488995,
			},
			wantError: require.NoError,
		},
		{
			name: "Looksrare V2 Sell",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xa871a0464d663d8fc4176ad8743532372380a31f38f05c42019e183f8b9c38e3"),
						ParentHash:   common.HexToHash("0xf9e9a28e8d9b04038f99bd727e7fae369d1486b85ef551b8b4a29415dec6e0cf"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("18274734", 0)),
						GasLimit:     30000000,
						GasUsed:      21253229,
						Timestamp:    1696393271,
						BaseFee:      lo.Must(new(big.Int).SetString("5984154729", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x1619e96a36880cbd96b6876f44d1a6d4f39a74608c649fbb3cb290919b84575f"),
						From:      common.HexToAddress("0xE105106F401d5caA68936d595F5D96f7f07D67aC"),
						Gas:       169017,
						GasPrice:  lo.Must(new(big.Int).SetString("6084154729", 10)),
						Hash:      common.HexToHash("0x1619e96a36880cbd96b6876f44d1a6d4f39a74608c649fbb3cb290919b84575f"),
						Input:     hexutil.MustDecode("0xe72853e100000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000038000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002a2e8716de67d40ddd28e15c765d58d2700000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f5e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000306b1ea3ecdf94ab739f1910bbda052ed4a9f949000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000094e9d73a2aeec807d85756b84759965b4186645000000000000000000000000000000000000000000000000000000000651c55b100000000000000000000000000000000000000000000000000000000652ecaa900000000000000000000000000000000000000000000000000254db1c224400000000000000000000000000000000000000000000000000000000000000001e0000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000002600000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000009f00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000419a06e78335ea5a55dd72c524f046870503a7bc2d61223a584ecc410c2175290f3ef81a86fe046f3dec5e41e07d1ca580e02fb69f77cf309e08bb4f1a61275d751b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x0000000000E655fAe4d56241588680F86E3b2377")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xa871a0464d663d8fc4176ad8743532372380a31f38f05c42019e183f8b9c38e3"),
						BlockNumber:       lo.Must(new(big.Int).SetString("18274734", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 16989952,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x16aa4d569"),
						GasUsed:           159241,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x306b1ea3ecdf94aB739F1910bbda052Ed4A9f949"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000e105106f401d5caa68936d595f5d96f7f07d67ac"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000009f"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("18274734", 0)),
							TransactionHash: common.HexToHash("0x1619e96a36880cbd96b6876f44d1a6d4f39a74608c649fbb3cb290919b84575f"),
							Index:           712,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x306b1ea3ecdf94aB739F1910bbda052Ed4A9f949"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000e105106f401d5caa68936d595f5d96f7f07d67ac"),
								common.HexToHash("0x00000000000000000000000094e9d73a2aeec807d85756b84759965b41866450"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000009f"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("18274734", 0)),
							TransactionHash: common.HexToHash("0x1619e96a36880cbd96b6876f44d1a6d4f39a74608c649fbb3cb290919b84575f"),
							Index:           713,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000094e9d73a2aeec807d85756b84759965b41866450"),
								common.HexToHash("0x000000000000000000000000e105106f401d5caa68936d595f5d96f7f07d67ac"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000251df2264a7800"),
							BlockNumber:     lo.Must(new(big.Int).SetString("18274734", 0)),
							TransactionHash: common.HexToHash("0x1619e96a36880cbd96b6876f44d1a6d4f39a74608c649fbb3cb290919b84575f"),
							Index:           714,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x0000000000E655fAe4d56241588680F86E3b2377"),
							Topics: []common.Hash{
								common.HexToHash("0x9aaa45d6db2ef74ead0751ea9113263d1dec1b50cea05f0ca2002cb8063564a4"),
							},
							Data:            hexutil.MustDecode("0x8acafe2c10bad725b3d7532de4e5d40e852293bc6362068bfc4f0adf2b55c97e0000000000000000000000000000000000000000000000000000000000000f5e0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000e105106f401d5caa68936d595f5d96f7f07d67ac00000000000000000000000094e9d73a2aeec807d85756b84759965b418664500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000306b1ea3ecdf94ab739f1910bbda052ed4a9f94900000000000000000000000000000000000000000000000000000000000001e00000000000000000000000000000000000000000000000000000000000000220000000000000000000000000e105106f401d5caa68936d595f5d96f7f07d67ac000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000251df2264a7800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002fbf9bd9c8000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000009f00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001"),
							BlockNumber:     lo.Must(new(big.Int).SetString("18274734", 0)),
							TransactionHash: common.HexToHash("0x1619e96a36880cbd96b6876f44d1a6d4f39a74608c649fbb3cb290919b84575f"),
							Index:           715,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000094e9d73a2aeec807d85756b84759965b41866450"),
								common.HexToHash("0x0000000000000000000000001838de7d4e4e42c8eb7b204a91e28e9fad14f536"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000002fbf9bd9c800"),
							BlockNumber:     lo.Must(new(big.Int).SetString("18274734", 0)),
							TransactionHash: common.HexToHash("0x1619e96a36880cbd96b6876f44d1a6d4f39a74608c649fbb3cb290919b84575f"),
							Index:           716,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x1619e96a36880cbd96b6876f44d1a6d4f39a74608c649fbb3cb290919b84575f"),
						TransactionIndex: 114,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x1619e96a36880cbd96b6876f44d1a6d4f39a74608c649fbb3cb290919b84575f",
				Network: filter.NetworkEthereum,
				Index:   114,
				From:    "0xE105106F401d5caA68936d595F5D96f7f07D67aC",
				To:      "0x0000000000E655fAe4d56241588680F86E3b2377",
				Type:    filter.TypeCollectibleTrade,
				Calldata: &schema.Calldata{
					FunctionHash: "0xe72853e1",
				},
				Platform: lo.ToPtr(filter.PlatformLooksRare),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("968846883200689")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeCollectibleTrade,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0xE105106F401d5caA68936d595F5D96f7f07D67aC",
						To:       "0x94e9D73a2aEEC807d85756B84759965b41866450",
						Metadata: metadata.CollectibleTrade{
							Action: metadata.ActionCollectibleTradeSell,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x306b1ea3ecdf94aB739F1910bbda052Ed4A9f949"),
								ID:       lo.ToPtr(lo.Must(decimal.NewFromString("159"))),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
								Name:     "Beanz",
								Symbol:   "BEANZ",
								Standard: metadata.StandardERC721,
								URI:      "ipfs://QmdYeDpkVZedk1mkGodjNmF35UNxwafhFLVvsHrWgJoz6A/beanz_metadata/159",
							},

							Cost: &metadata.Token{
								Address:  lo.ToPtr(common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2").String()),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("10500000000000000"))),
								Name:     "Wrapped Ether",
								Symbol:   "WETH",
								Decimals: 18,
								Standard: 20,
							},
						},
					},
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0x94e9D73a2aEEC807d85756B84759965b41866450",
						To:       "0x1838De7d4e4e42c8eB7b204A91e28E9fad14F536",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("52500000000000"))),
							Name:     "Wrapped Ether",
							Symbol:   "WETH",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
				},
				Status:    true,
				Timestamp: 1696393271,
			},
			wantError: require.NoError,
		},
		{
			name: "LooksRare V2 Aggregator Execute",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xf74ad15c3bb59a96679a5f91d10ce550d8b81c908d78b412ff9d1bc685bc1231"),
						ParentHash:   common.HexToHash("0xcf1a87b193ede34ba901dfa56f19150e879257d151ca478a763adef3fed98b95"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"),
						Number:       lo.Must(new(big.Int).SetString("18283530", 0)),
						GasLimit:     30000000,
						GasUsed:      14018077,
						Timestamp:    1696499483,
						BaseFee:      lo.Must(new(big.Int).SetString("6730172615", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x3d02d4b9d163b0bbe9b2212a699eeeed2c3fc8f7bce870d0952a5f5e351407c7"),
						From:      common.HexToAddress("0x5A46Cd7ab0CEEbe787267bEd3Cf368Ac4C240b16"),
						Gas:       353257,
						GasPrice:  lo.Must(new(big.Int).SetString("6830172615", 10)),
						Hash:      common.HexToHash("0x3d02d4b9d163b0bbe9b2212a699eeeed2c3fc8f7bce870d0952a5f5e351407c7"),
						Input:     hexutil.MustDecode("0x12abee1200000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005a46cd7ab0ceebe787267bed3cf368ac4c240b160000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000055d65008f1dff7167f24e70db431f6a809ce350000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000058000000000000000000000000000000000000000000000000000000000000009e0000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000260000000000000000000000000ce49a5b02943335cdb1512eea1073c28a5d16f5500000000000000000000000029e25ecaa1bb58efb51433dfb65bd18ad38e4c2500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000038d7ea4c68000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000651e7f3c00000000000000000000000000000000000000000000000000000000651e87b800000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000003b000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000040f7e0dea03f7d51db0cef1647a8b52e5945d00e2694bebc089e3b4ab36b690b839f9c20d8401b9668d06873eb532a0b243d76f4affa40b1301036172378fb7f3d000000000000000000000000a86f5324129c34312187cde5b42fe283fc493fd800000000000000000000000029e25ecaa1bb58efb51433dfb65bd18ad38e4c250000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000001800000000000000000000000000000000000000000000000000006c00a3912c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000651e7e5100000000000000000000000000000000000000000000000000000000660fc0d100000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000003a00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000006311f92c3905d29eb6f2e0c476606c0584dd5f4b0766984cb44f746fc1b4f977fa7a550fb71a3d38fe82d4c22c8b5f3e20fc34dbca58c38e938a012930441ed7570000019cfb8c98bb197beb03d27a58de0cf7f53b4c1153984c14ee1c2a3cb33e04b66e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000026000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004c00500000ad104d7dbd00e3ae0a5c00560c000000000000000000000000000000000000000000000000000000000000000000360c6ebe0000000000000000000000000000000000000000611aea950096ef9b0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000376aa98310800000000000000000000000000ce49a5b02943335cdb1512eea1073c28a5d16f55000000000000000000000000000000000000000000000000000016bcc41e90000000000000000000000000000000a26b00c1f0df003000390027140000faa719000000000000000000000000000000000000000000000000000000174876e8000000000000000000000000005a46cd7ab0ceebe787267bed3cf368ac4c240b1600000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004c00500000ad104d7dbd00e3ae0a5c00560c000000000000000000000000000000000000000000000000000000000000000000360c6ebe00000000000000000000000000000000000000008ace854bda2c5c040000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000694d6c4724800000000000000000000000000a86f5324129c34312187cde5b42fe283fc493fd800000000000000000000000000000000000000000000000000002b3374a078000000000000000000000000000000a26b00c1f0df003000390027140000faa71900000000000000000000000000000000000000000000000000000000000003e00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000001e000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x00000000005228B791a99a61f36A130d50600106")),
						Value:     lo.Must(new(big.Int).SetString("2900000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xf74ad15c3bb59a96679a5f91d10ce550d8b81c908d78b412ff9d1bc685bc1231"),
						BlockNumber:       lo.Must(new(big.Int).SetString("18283530", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 3373349,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x1971c29c7"),
						GasUsed:           307180,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x29e25EcAa1bb58Efb51433dFb65Bd18ad38e4c25"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000ce49a5b02943335cdb1512eea1073c28a5d16f55"),
								common.HexToHash("0x0000000000000000000000005a46cd7ab0ceebe787267bed3cf368ac4c240b16"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000003b"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("18283530", 0)),
							TransactionHash: common.HexToHash("0x3d02d4b9d163b0bbe9b2212a699eeeed2c3fc8f7bce870d0952a5f5e351407c7"),
							Index:           57,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x00000000000000ADc04C56Bf30aC9d3c0aAF14dC"),
							Topics: []common.Hash{
								common.HexToHash("0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31"),
								common.HexToHash("0x000000000000000000000000ce49a5b02943335cdb1512eea1073c28a5d16f55"),
								common.HexToHash("0x000000000000000000000000004c00500000ad104d7dbd00e3ae0a5c00560c00"),
							},
							Data:            hexutil.MustDecode("0x9ce962d6c5a262152990c95e4388542ebd631cd1388e984c160102edc4280dec0000000000000000000000005a46cd7ab0ceebe787267bed3cf368ac4c240b16000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000029e25ecaa1bb58efb51433dfb65bd18ad38e4c25000000000000000000000000000000000000000000000000000000000000003b00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000376aa98310800000000000000000000000000ce49a5b02943335cdb1512eea1073c28a5d16f55000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000016bcc41e90000000000000000000000000000000a26b00c1f0df003000390027140000faa719000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000174876e8000000000000000000000000005a46cd7ab0ceebe787267bed3cf368ac4c240b16"),
							BlockNumber:     lo.Must(new(big.Int).SetString("18283530", 0)),
							TransactionHash: common.HexToHash("0x3d02d4b9d163b0bbe9b2212a699eeeed2c3fc8f7bce870d0952a5f5e351407c7"),
							Index:           58,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x29e25EcAa1bb58Efb51433dFb65Bd18ad38e4c25"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000a86f5324129c34312187cde5b42fe283fc493fd8"),
								common.HexToHash("0x0000000000000000000000005a46cd7ab0ceebe787267bed3cf368ac4c240b16"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000003a"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("18283530", 0)),
							TransactionHash: common.HexToHash("0x3d02d4b9d163b0bbe9b2212a699eeeed2c3fc8f7bce870d0952a5f5e351407c7"),
							Index:           59,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x00000000000000ADc04C56Bf30aC9d3c0aAF14dC"),
							Topics: []common.Hash{
								common.HexToHash("0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31"),
								common.HexToHash("0x000000000000000000000000a86f5324129c34312187cde5b42fe283fc493fd8"),
								common.HexToHash("0x000000000000000000000000004c00500000ad104d7dbd00e3ae0a5c00560c00"),
							},
							Data:            hexutil.MustDecode("0x0212b5cd5c026ece52ddd87aafc11f19fab03c64e89d4de619c9c21b9e393df50000000000000000000000005a46cd7ab0ceebe787267bed3cf368ac4c240b16000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000029e25ecaa1bb58efb51433dfb65bd18ad38e4c25000000000000000000000000000000000000000000000000000000000000003a00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000694d6c4724800000000000000000000000000a86f5324129c34312187cde5b42fe283fc493fd800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002b3374a078000000000000000000000000000000a26b00c1f0df003000390027140000faa719"),
							BlockNumber:     lo.Must(new(big.Int).SetString("18283530", 0)),
							TransactionHash: common.HexToHash("0x3d02d4b9d163b0bbe9b2212a699eeeed2c3fc8f7bce870d0952a5f5e351407c7"),
							Index:           60,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x00000000005228B791a99a61f36A130d50600106"),
							Topics: []common.Hash{
								common.HexToHash("0x807273efecfbeb7ae7d3a2189d1ed5a7db80074eed86e7d80b10bb925cd1db73"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000005a46cd7ab0ceebe787267bed3cf368ac4c240b16"),
							BlockNumber:     lo.Must(new(big.Int).SetString("18283530", 0)),
							TransactionHash: common.HexToHash("0x3d02d4b9d163b0bbe9b2212a699eeeed2c3fc8f7bce870d0952a5f5e351407c7"),
							Index:           61,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x3d02d4b9d163b0bbe9b2212a699eeeed2c3fc8f7bce870d0952a5f5e351407c7"),
						TransactionIndex: 25,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x3d02d4b9d163b0bbe9b2212a699eeeed2c3fc8f7bce870d0952a5f5e351407c7",
				Network: filter.NetworkEthereum,
				Index:   25,
				From:    "0x5A46Cd7ab0CEEbe787267bEd3Cf368Ac4C240b16",
				To:      "0x00000000005228B791a99a61f36A130d50600106",
				Type:    filter.TypeCollectibleTrade,
				Calldata: &schema.Calldata{
					FunctionHash: "0x12abee12",
				},
				Platform: lo.ToPtr(filter.PlatformLooksRare),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("2098092423875700")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0x5A46Cd7ab0CEEbe787267bEd3Cf368Ac4C240b16",
						To:       "0xCE49a5B02943335CdB1512eEA1073C28a5d16F55",
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("974900000000000"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0x5A46Cd7ab0CEEbe787267bEd3Cf368Ac4C240b16",
						To:       "0x0000a26b00c1F0DF003000390027140000fAa719",
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("25000000000000"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type:     filter.TypeCollectibleTrade,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0xCE49a5B02943335CdB1512eEA1073C28a5d16F55",
						To:       "0x5A46Cd7ab0CEEbe787267bEd3Cf368Ac4C240b16",
						Metadata: metadata.CollectibleTrade{
							Action: metadata.ActionCollectibleTradeBuy,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x29e25EcAa1bb58Efb51433dFb65Bd18ad38e4c25"),
								ID:       lo.ToPtr(lo.Must(decimal.NewFromString("59"))),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
								Name:     "Ghostd Ghosties",
								Symbol:   "GST",
								Standard: metadata.StandardERC721,
								URI:      "ipfs://bafybeigb5nf5o4rfdyuas3kdaxugr5uv4tboeprsqa5vgn5lhndvbx3d2a/59",
							},
							Cost: &metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("999900000000000"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0x5A46Cd7ab0CEEbe787267bEd3Cf368Ac4C240b16",
						To:       "0xA86f5324129c34312187CdE5B42Fe283FC493fD8",
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1852500000000000"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0x5A46Cd7ab0CEEbe787267bEd3Cf368Ac4C240b16",
						To:       "0x0000a26b00c1F0DF003000390027140000fAa719",
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("47500000000000"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type:     filter.TypeCollectibleTrade,
						Platform: filter.PlatformLooksRare.String(),
						From:     "0xA86f5324129c34312187CdE5B42Fe283FC493fD8",
						To:       "0x5A46Cd7ab0CEEbe787267bEd3Cf368Ac4C240b16",
						Metadata: metadata.CollectibleTrade{
							Action: metadata.ActionCollectibleTradeBuy,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x29e25EcAa1bb58Efb51433dFb65Bd18ad38e4c25"),
								ID:       lo.ToPtr(lo.Must(decimal.NewFromString("58"))),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
								Name:     "Ghostd Ghosties",
								Symbol:   "GST",
								Standard: metadata.StandardERC721,
								URI:      "ipfs://bafybeigb5nf5o4rfdyuas3kdaxugr5uv4tboeprsqa5vgn5lhndvbx3d2a/58",
							},
							Cost: &metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1900000000000000"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1696499483,
			},
			wantError: require.NoError,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			instance, err := worker.NewWorker(testcase.arguments.config)
			require.NoError(t, err)

			matched, err := instance.Match(ctx, testcase.arguments.task)
			testcase.wantError(t, err)
			require.True(t, matched)

			feed, err := instance.Transform(ctx, testcase.arguments.task)
			testcase.wantError(t, err)

			//t.Log(string(lo.Must(json.MarshalIndent(feed, "", "\x20\x20"))))

			require.Equal(t, testcase.want, feed)
		})
	}
}
