package lido_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/contract/lido"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/lido"
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
		{
			name: "Add ETH liquidity and mint stETH",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xa6041b2aa1b9099f4ef1b24d50c53f3734cffffb7f117dd9cceb3ad8be9d15e6"),
						ParentHash:   common.HexToHash("0x14823424e41a939d92b4a83e512dd83383644b4a8769522a8edebaed8cf1470c"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("17487526", 0)),
						GasLimit:     30000000,
						GasUsed:      18143485,
						Timestamp:    1686859415,
						BaseFee:      lo.Must(new(big.Int).SetString("17502326215", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xd632e7c240f2636d6f757a23480e27c968ba28ed7240c0f2c7b8206903252b22"),
						From:      common.HexToAddress("0xcF7c2e2283e31984cA3e809e01ee35E80bCdAEea"),
						Gas:       112608,
						GasPrice:  lo.Must(new(big.Int).SetString("19002326215", 10)),
						Hash:      common.HexToHash("0xd632e7c240f2636d6f757a23480e27c968ba28ed7240c0f2c7b8206903252b22"),
						Input:     hexutil.MustDecode("0xa1903eab0000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84")),
						Value:     lo.Must(new(big.Int).SetString("967207443783628052990", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xa6041b2aa1b9099f4ef1b24d50c53f3734cffffb7f117dd9cceb3ad8be9d15e6"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17487526", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 5425922,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x46ca07cc7"),
						GasUsed:           97402,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
							Topics: []common.Hash{
								common.HexToHash("0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a"),
								common.HexToHash("0x000000000000000000000000cf7c2e2283e31984ca3e809e01ee35e80bcdaeea"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000346eb31da2f4bf59fe0000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17487526", 0)),
							TransactionHash: common.HexToHash("0xd632e7c240f2636d6f757a23480e27c968ba28ed7240c0f2c7b8206903252b22"),
							Index:           125,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000cf7c2e2283e31984ca3e809e01ee35e80bcdaeea"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000346eb31da2f4bf59fd"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17487526", 0)),
							TransactionHash: common.HexToHash("0xd632e7c240f2636d6f757a23480e27c968ba28ed7240c0f2c7b8206903252b22"),
							Index:           126,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
							Topics: []common.Hash{
								common.HexToHash("0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000cf7c2e2283e31984ca3e809e01ee35e80bcdaeea"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000002e7638363eb2832c93"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17487526", 0)),
							TransactionHash: common.HexToHash("0xd632e7c240f2636d6f757a23480e27c968ba28ed7240c0f2c7b8206903252b22"),
							Index:           127,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xd632e7c240f2636d6f757a23480e27c968ba28ed7240c0f2c7b8206903252b22"),
						TransactionIndex: 56,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0xd632e7c240f2636d6f757a23480e27c968ba28ed7240c0f2c7b8206903252b22",
				Network: filter.NetworkEthereum,
				Index:   56,
				From:    "0xcF7c2e2283e31984cA3e809e01ee35E80bCdAEea",
				To:      lido.AddressStakedETH.String(),
				Type:    filter.TypeExchangeLiquidity,
				Calldata: &schema.Calldata{
					FunctionHash: "0xa1903eab",
				},
				Platform: lo.ToPtr(filter.PlatformLido),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("1850864577993430")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeLiquidity,
						Platform: filter.PlatformLido.String(),
						From:     "0xcF7c2e2283e31984cA3e809e01ee35E80bCdAEea",
						To:       lido.AddressStakedETH.String(),
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquidityAdd,
							Tokens: []metadata.Token{
								{
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("967207443783628052990"))),
									Name:     "Ethereum",
									Symbol:   "ETH",
									Decimals: 18,
								},
							},
						},
					},
					{
						Type:     filter.TypeTransactionMint,
						Platform: filter.PlatformLido.String(),
						From:     ethereum.AddressGenesis.String(),
						To:       "0xcF7c2e2283e31984cA3e809e01ee35E80bCdAEea",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("967207443783628052990"))),
							Name:     "Liquid staked Ether 2.0",
							Symbol:   "stETH",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
				},
				Status:    true,
				Timestamp: 1686859415,
			},
			wantError: require.NoError,
		},
		{
			name: "Remove stETH liquidity and mint unstETH #2427",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x96a4e104794025dfb9bcd92805c24efccf6811a9aea17a6f7d10adef44d02f47"),
						ParentHash:   common.HexToHash("0xbd6092d609c0ee98dba12c2cd794d2cfd4e1efa7079b83ab0935cebfb730fe44"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x690B9A9E9aa1C9dB991C7721a92d351Db4FaC990"),
						Number:       lo.Must(new(big.Int).SetString("17489467", 0)),
						GasLimit:     30000000,
						GasUsed:      10669796,
						Timestamp:    1686882899,
						BaseFee:      lo.Must(new(big.Int).SetString("14935721596", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xe1bc6380c76131ed87a875bd0740f187548c22c1e79eec10d9a7f00c51ad914c"),
						From:      common.HexToAddress("0x64833AEbDB5a39EAD24b7F053DbDA0FE38ad632e"),
						Gas:       289684,
						GasPrice:  lo.Must(new(big.Int).SetString("16435721596", 10)),
						Hash:      common.HexToHash("0xe1bc6380c76131ed87a875bd0740f187548c22c1e79eec10d9a7f00c51ad914c"),
						Input:     hexutil.MustDecode("0xacf41e4d00000000000000000000000000000000000000000000000000000000000000e000000000000000000000000064833aebdb5a39ead24b7f053dbda0fe38ad632e00000000000000000000000000000000000000000000000029e217028c210000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000000000000000000000000000001c707dd280cdbaf815c602345c54af5a2eeb1e811033adefbf10d82fc00ee296b255ca29d505d68a7bf1b8384fd69370445cf684b7922001759988df3b48672107000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000029e217028c210000"),
						To:        lo.ToPtr(common.HexToAddress("0x889edC2eDab5f40e902b864aD4d7AdE8E412F9B1")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x96a4e104794025dfb9bcd92805c24efccf6811a9aea17a6f7d10adef44d02f47"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17489467", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 4643762,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3d3a5357c"),
						GasUsed:           265756,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x00000000000000000000000064833aebdb5a39ead24b7f053dbda0fe38ad632e"),
								common.HexToHash("0x000000000000000000000000889edc2edab5f40e902b864ad4d7ade8e412f9b1"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000029e217028c210000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17489467", 0)),
							TransactionHash: common.HexToHash("0xe1bc6380c76131ed87a875bd0740f187548c22c1e79eec10d9a7f00c51ad914c"),
							Index:           143,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x00000000000000000000000064833aebdb5a39ead24b7f053dbda0fe38ad632e"),
								common.HexToHash("0x000000000000000000000000889edc2edab5f40e902b864ad4d7ade8e412f9b1"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17489467", 0)),
							TransactionHash: common.HexToHash("0xe1bc6380c76131ed87a875bd0740f187548c22c1e79eec10d9a7f00c51ad914c"),
							Index:           144,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000064833aebdb5a39ead24b7f053dbda0fe38ad632e"),
								common.HexToHash("0x000000000000000000000000889edc2edab5f40e902b864ad4d7ade8e412f9b1"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000029e217028c210000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17489467", 0)),
							TransactionHash: common.HexToHash("0xe1bc6380c76131ed87a875bd0740f187548c22c1e79eec10d9a7f00c51ad914c"),
							Index:           145,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
							Topics: []common.Hash{
								common.HexToHash("0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb"),
								common.HexToHash("0x00000000000000000000000064833aebdb5a39ead24b7f053dbda0fe38ad632e"),
								common.HexToHash("0x000000000000000000000000889edc2edab5f40e902b864ad4d7ade8e412f9b1"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000251d22bf003814ee"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17489467", 0)),
							TransactionHash: common.HexToHash("0xe1bc6380c76131ed87a875bd0740f187548c22c1e79eec10d9a7f00c51ad914c"),
							Index:           146,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x889edC2eDab5f40e902b864aD4d7AdE8E412F9B1"),
							Topics: []common.Hash{
								common.HexToHash("0xf0cb471f23fb74ea44b8252eb1881a2dca546288d9f6e90d1a0e82fe0ed342ab"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000097b"),
								common.HexToHash("0x00000000000000000000000064833aebdb5a39ead24b7f053dbda0fe38ad632e"),
								common.HexToHash("0x00000000000000000000000064833aebdb5a39ead24b7f053dbda0fe38ad632e"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000029e217028c210000000000000000000000000000000000000000000000000000251d22bf003814ee"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17489467", 0)),
							TransactionHash: common.HexToHash("0xe1bc6380c76131ed87a875bd0740f187548c22c1e79eec10d9a7f00c51ad914c"),
							Index:           147,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x889edC2eDab5f40e902b864aD4d7AdE8E412F9B1"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x00000000000000000000000064833aebdb5a39ead24b7f053dbda0fe38ad632e"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000097b"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("17489467", 0)),
							TransactionHash: common.HexToHash("0xe1bc6380c76131ed87a875bd0740f187548c22c1e79eec10d9a7f00c51ad914c"),
							Index:           148,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xe1bc6380c76131ed87a875bd0740f187548c22c1e79eec10d9a7f00c51ad914c"),
						TransactionIndex: 47,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0xe1bc6380c76131ed87a875bd0740f187548c22c1e79eec10d9a7f00c51ad914c",
				Network: filter.NetworkEthereum,
				Index:   47,
				From:    "0x64833AEbDB5a39EAD24b7F053DbDA0FE38ad632e",
				To:      lido.AddressStakedETHWithdrawalNFT.String(),
				Type:    filter.TypeExchangeLiquidity,
				Calldata: &schema.Calldata{
					FunctionHash: "0xacf41e4d",
				},
				Platform: lo.ToPtr(filter.PlatformLido),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("4367891628466576")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeLiquidity,
						Platform: filter.PlatformLido.String(),
						From:     "0x64833AEbDB5a39EAD24b7F053DbDA0FE38ad632e",
						To:       lido.AddressStakedETHWithdrawalNFT.String(),
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquidityRemove,
							Tokens: []metadata.Token{
								{
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("3018000000000000000"))),
									Name:     "Ethereum",
									Symbol:   "ETH",
									Decimals: 18,
								},
							},
						},
					},
					{
						Type:     filter.TypeCollectibleMint,
						Platform: filter.PlatformLido.String(),
						From:     ethereum.AddressGenesis.String(),
						To:       "0x64833AEbDB5a39EAD24b7F053DbDA0FE38ad632e",
						Metadata: metadata.CollectibleTransfer{
							Address:  lo.ToPtr("0x889edC2eDab5f40e902b864aD4d7AdE8E412F9B1"),
							ID:       lo.ToPtr(lo.Must(decimal.NewFromString("2427"))),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
							Name:     "Lido: stETH Withdrawal NFT",
							Symbol:   "unstETH",
							Standard: metadata.StandardERC721,
							URI:      "https://wq-api.lido.fi/v1/nft/2427?requested=3018000000000000000&created_at=1686882899",
						},
					},
				},
				Status:    true,
				Timestamp: 1686882899,
			},
			wantError: require.NoError,
		},
		{
			name: "Burn unstETH #1685 and receive ETH",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x5e39c72459101c14e25455f4b8cc92d015c9ef4d8a604710b9c5e885ce26baa4"),
						ParentHash:   common.HexToHash("0xa21c66b0ab680bcff9da62d48807408bf7195f18c4fc521ed18c409d656760fb"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xf272E9AAdC2a082b7AF45ec89faB88f4ae899394"),
						Number:       lo.Must(new(big.Int).SetString("17489512", 0)),
						GasLimit:     30000000,
						GasUsed:      9791829,
						Timestamp:    1686883439,
						BaseFee:      lo.Must(new(big.Int).SetString("16816544309", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xcc70adde277e1f7faff67d5ebda2b8fdae071933f4b751790b5aba4b65b012c8"),
						From:      common.HexToAddress("0x0C71f889D32002bCAC135434cac861Baf17A5f67"),
						Gas:       89842,
						GasPrice:  lo.Must(new(big.Int).SetString("18316544309", 10)),
						Hash:      common.HexToHash("0xcc70adde277e1f7faff67d5ebda2b8fdae071933f4b751790b5aba4b65b012c8"),
						Input:     hexutil.MustDecode("0xe3afe0a3000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000069500000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000017"),
						To:        lo.ToPtr(common.HexToAddress("0x889edC2eDab5f40e902b864aD4d7AdE8E412F9B1")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x5e39c72459101c14e25455f4b8cc92d015c9ef4d8a604710b9c5e885ce26baa4"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17489512", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 3560347,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x443c04935"),
						GasUsed:           74498,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x889edC2eDab5f40e902b864aD4d7AdE8E412F9B1"),
							Topics: []common.Hash{
								common.HexToHash("0x6ad26c5e238e7d002799f9a5db07e81ef14e37386ae03496d7a7ef04713e145b"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000695"),
								common.HexToHash("0x0000000000000000000000000c71f889d32002bcac135434cac861baf17a5f67"),
								common.HexToHash("0x0000000000000000000000000c71f889d32002bcac135434cac861baf17a5f67"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000001262c1d8a9cd9669"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17489512", 0)),
							TransactionHash: common.HexToHash("0xcc70adde277e1f7faff67d5ebda2b8fdae071933f4b751790b5aba4b65b012c8"),
							Index:           100,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x889edC2eDab5f40e902b864aD4d7AdE8E412F9B1"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000c71f889d32002bcac135434cac861baf17a5f67"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000695"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("17489512", 0)),
							TransactionHash: common.HexToHash("0xcc70adde277e1f7faff67d5ebda2b8fdae071933f4b751790b5aba4b65b012c8"),
							Index:           101,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xcc70adde277e1f7faff67d5ebda2b8fdae071933f4b751790b5aba4b65b012c8"),
						TransactionIndex: 44,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0xcc70adde277e1f7faff67d5ebda2b8fdae071933f4b751790b5aba4b65b012c8",
				Network: filter.NetworkEthereum,
				Index:   44,
				From:    "0x0C71f889D32002bCAC135434cac861Baf17A5f67",
				To:      lido.AddressStakedETHWithdrawalNFT.String(),
				Type:    filter.TypeCollectibleBurn,
				Calldata: &schema.Calldata{
					FunctionHash: "0xe3afe0a3",
				},
				Platform: lo.ToPtr(filter.PlatformLido),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("1364545917931882")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeCollectibleBurn,
						Platform: filter.PlatformLido.String(),
						From:     "0x0C71f889D32002bCAC135434cac861Baf17A5f67",
						To:       ethereum.AddressGenesis.String(),
						Metadata: metadata.CollectibleTransfer{
							Address:  lo.ToPtr("0x889edC2eDab5f40e902b864aD4d7AdE8E412F9B1"),
							ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1685"))),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
							Name:     "Lido: stETH Withdrawal NFT",
							Symbol:   "unstETH",
							Standard: metadata.StandardERC721,
						},
					},
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformLido.String(),
						From:     lido.AddressStakedETHWithdrawalNFT.String(),
						To:       "0x0C71f889D32002bCAC135434cac861Baf17A5f67",
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1324834376706266729"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
				},
				Status:    true,
				Timestamp: 1686883439,
			},
			wantError: require.NoError,
		},
		{
			name: "Add MATIC liquidity and mint stMATIC",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x5441b85042d561a65fa0f3fb2c030de8cbe6bed4460a079b36ada369015d9511"),
						ParentHash:   common.HexToHash("0xfc5f3efbe3ac4311c07a3d734b7ceb94115838207614352abf853cd5a2555538"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xFeebabE6b0418eC13b30aAdF129F5DcDd4f70CeA"),
						Number:       lo.Must(new(big.Int).SetString("17489412", 0)),
						GasLimit:     30000000,
						GasUsed:      29776079,
						Timestamp:    1686882227,
						BaseFee:      lo.Must(new(big.Int).SetString("13448442191", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xa8405c2dfb85565b406d5b6b60f0b460f61c211977d967ee546967ae43be2c2f"),
						From:      common.HexToAddress("0xc9fE2F54234c513dbeadbF12Af72a7c1eDB0058f"),
						Gas:       782183,
						GasPrice:  lo.Must(new(big.Int).SetString("15448442191", 10)),
						Hash:      common.HexToHash("0xa8405c2dfb85565b406d5b6b60f0b460f61c211977d967ee546967ae43be2c2f"),
						Input:     hexutil.MustDecode("0xf532e86a00000000000000000000000000000000000000000000003c19ad6d0b59267800000000000000000000000000c9fe2f54234c513dbeadbf12af72a7c1edb0058f"),
						To:        lo.ToPtr(common.HexToAddress("0x9ee91F9f426fA633d227f7a9b000E28b9dfd8599")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x5441b85042d561a65fa0f3fb2c030de8cbe6bed4460a079b36ada369015d9511"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17489412", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 2836245,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x398cc854f"),
						GasUsed:           752101,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000c9fe2f54234c513dbeadbf12af72a7c1edb0058f"),
								common.HexToHash("0x0000000000000000000000009ee91f9f426fa633d227f7a9b000e28b9dfd8599"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000003c19ad6d0b59267800"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17489412", 0)),
							TransactionHash: common.HexToHash("0xa8405c2dfb85565b406d5b6b60f0b460f61c211977d967ee546967ae43be2c2f"),
							Index:           47,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000c9fe2f54234c513dbeadbf12af72a7c1edb0058f"),
								common.HexToHash("0x0000000000000000000000009ee91f9f426fa633d227f7a9b000e28b9dfd8599"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17489412", 0)),
							TransactionHash: common.HexToHash("0xa8405c2dfb85565b406d5b6b60f0b460f61c211977d967ee546967ae43be2c2f"),
							Index:           48,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x9ee91F9f426fA633d227f7a9b000E28b9dfd8599"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000c9fe2f54234c513dbeadbf12af72a7c1edb0058f"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000037edd1d2d7404cdde2"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17489412", 0)),
							TransactionHash: common.HexToHash("0xa8405c2dfb85565b406d5b6b60f0b460f61c211977d967ee546967ae43be2c2f"),
							Index:           49,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x28e4F3a7f651294B9564800b2D01f35189A5bFbE"),
							Topics: []common.Hash{
								common.HexToHash("0x103fed9db65eac19c4d870f49ab7520fe03b99f1838e5996caf47e9e43308392"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000292cf3"),
								common.HexToHash("0x0000000000000000000000008397259c983751daf40400790063935a11afa28a"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000c7dd5c30dca04f487c9ede0c5ac580c91587fc660000000000000000000000000833f5bd45803e05ef54e119a77e463ce6b1a963000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000006899f2330f95dc2bdd72f800000000000000000000000000000000000000000070671c471992f0603e0a92"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17489412", 0)),
							TransactionHash: common.HexToHash("0xa8405c2dfb85565b406d5b6b60f0b460f61c211977d967ee546967ae43be2c2f"),
							Index:           50,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x9ee91F9f426fA633d227f7a9b000E28b9dfd8599"),
							Topics: []common.Hash{
								common.HexToHash("0x98d2bc018caf34c71a8f920d9d93d4ed62e9789506b74087b48570c17b28ed99"),
								common.HexToHash("0x000000000000000000000000c9fe2f54234c513dbeadbf12af72a7c1edb0058f"),
								common.HexToHash("0x000000000000000000000000c9fe2f54234c513dbeadbf12af72a7c1edb0058f"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000003c19ad6d0b59267800"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17489412", 0)),
							TransactionHash: common.HexToHash("0xa8405c2dfb85565b406d5b6b60f0b460f61c211977d967ee546967ae43be2c2f"),
							Index:           51,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xa8405c2dfb85565b406d5b6b60f0b460f61c211977d967ee546967ae43be2c2f"),
						TransactionIndex: 29,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0xa8405c2dfb85565b406d5b6b60f0b460f61c211977d967ee546967ae43be2c2f",
				Network: filter.NetworkEthereum,
				Index:   29,
				From:    "0xc9fE2F54234c513dbeadbF12Af72a7c1eDB0058f",
				To:      lido.AddressStakedMATIC.String(),
				Type:    filter.TypeExchangeLiquidity,
				Calldata: &schema.Calldata{
					FunctionHash: "0xf532e86a",
				},
				Platform: lo.ToPtr(filter.PlatformLido),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("11618788820293291")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeLiquidity,
						Platform: filter.PlatformLido.String(),
						From:     "0xc9fE2F54234c513dbeadbF12Af72a7c1eDB0058f",
						To:       lido.AddressStakedMATIC.String(),
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquidityAdd,
							Tokens: []metadata.Token{
								{
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1108654899340000000000"))),
									Name:     "Ethereum",
									Symbol:   "ETH",
									Decimals: 18,
								},
							},
						},
					},
					{
						Type:     filter.TypeTransactionMint,
						Platform: filter.PlatformLido.String(),
						From:     ethereum.AddressGenesis.String(),
						To:       "0xc9fE2F54234c513dbeadbF12Af72a7c1eDB0058f",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0x9ee91F9f426fA633d227f7a9b000E28b9dfd8599"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1108654899340000000000"))),
							Name:     "Staked MATIC",
							Symbol:   "stMATIC",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
				},
				Status:    true,
				Timestamp: 1686882227,
			},
			wantError: require.NoError,
		},
		{
			name: "Remove stMATIC liquidity and mint PLO #1376",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x078fa6c6c9d8f68c36647575b16ce012da583aa32579b63e4ff5309e4cd0fc8a"),
						ParentHash:   common.HexToHash("0x40e46e912b49087b499f88e2d8b0ed6e18778110041df7e4114026e1513a04d6"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"),
						Number:       lo.Must(new(big.Int).SetString("17477988", 0)),
						GasLimit:     30000000,
						GasUsed:      20212491,
						Timestamp:    1686743375,
						BaseFee:      lo.Must(new(big.Int).SetString("14286332938", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x9247508d5690e97814a3f32b1f8ffda14c7b3404a7ae280eea6e93df2423dfd2"),
						From:      common.HexToAddress("0x36cc7B13029B5DEe4034745FB4F24034f3F2ffc6"),
						Gas:       1332415,
						GasPrice:  lo.Must(new(big.Int).SetString("14386332938", 10)),
						Hash:      common.HexToHash("0x9247508d5690e97814a3f32b1f8ffda14c7b3404a7ae280eea6e93df2423dfd2"),
						Input:     hexutil.MustDecode("0xccc143b800000000000000000000000000000000000000000000197b04d6de02429b2be60000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x9ee91F9f426fA633d227f7a9b000E28b9dfd8599")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x078fa6c6c9d8f68c36647575b16ce012da583aa32579b63e4ff5309e4cd0fc8a"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17477988", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 13476503,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3597e050a"),
						GasUsed:           1299557,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x60a91E2B7A1568f0848f3D43353C453730082E46"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x00000000000000000000000036cc7b13029b5dee4034745fb4f24034f3f2ffc6"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000560"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("17477988", 0)),
							TransactionHash: common.HexToHash("0x9247508d5690e97814a3f32b1f8ffda14c7b3404a7ae280eea6e93df2423dfd2"),
							Index:           186,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000005e3ef299fddf15eaa0432e6e66473ace8c13d908"),
								common.HexToHash("0x0000000000000000000000009ee91f9f426fa633d227f7a9b000e28b9dfd8599"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000005dbeb0ca0ac842d867"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17477988", 0)),
							TransactionHash: common.HexToHash("0x9247508d5690e97814a3f32b1f8ffda14c7b3404a7ae280eea6e93df2423dfd2"),
							Index:           187,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xa59C847Bd5aC0172Ff4FE912C5d29E5A71A7512B"),
							Topics: []common.Hash{
								common.HexToHash("0x31d1715032654fde9867c0f095aecce1113049e30b9f4ecbaa6954ed6c63b8df"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000075"),
								common.HexToHash("0x0000000000000000000000009ee91f9f426fa633d227f7a9b000e28b9dfd8599"),
								common.HexToHash("0x00000000000000000000000000000000000000000000005dbeb0ca0ac842d867"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("17477988", 0)),
							TransactionHash: common.HexToHash("0x9247508d5690e97814a3f32b1f8ffda14c7b3404a7ae280eea6e93df2423dfd2"),
							Index:           188,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC55D28Ac155C1a43eF2309869b704dF677788E81"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000009ee91f9f426fa633d227f7a9b000e28b9dfd8599"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000001b5fefd3bf56d174cee7"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17477988", 0)),
							TransactionHash: common.HexToHash("0x9247508d5690e97814a3f32b1f8ffda14c7b3404a7ae280eea6e93df2423dfd2"),
							Index:           189,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x6dF5CB08d3f0193C768C8A01f42ac4424DC5086b"),
							Topics: []common.Hash{
								common.HexToHash("0x502f5a6697a92e602c626b931a4a771fef4360eb9cece7bf906f10c5cec96aaa"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000075"),
								common.HexToHash("0x0000000000000000000000009ee91f9f426fa633d227f7a9b000e28b9dfd8599"),
								common.HexToHash("0x000000000000000000000000000000000000000000001b5fefd3bf56d174cee7"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000001b5fefd3bf56d174cee700000000000000000000000000000000000000000000000000000000000000c7"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17477988", 0)),
							TransactionHash: common.HexToHash("0x9247508d5690e97814a3f32b1f8ffda14c7b3404a7ae280eea6e93df2423dfd2"),
							Index:           190,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xa59C847Bd5aC0172Ff4FE912C5d29E5A71A7512B"),
							Topics: []common.Hash{
								common.HexToHash("0x35af9eea1f0e7b300b0a14fae90139a072470e44daa3f14b5069bebbc1265bda"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000075"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000229"),
								common.HexToHash("0x0000000000000000000000000000000000000000001a29d459884950f51ae4a3"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("17477988", 0)),
							TransactionHash: common.HexToHash("0x9247508d5690e97814a3f32b1f8ffda14c7b3404a7ae280eea6e93df2423dfd2"),
							Index:           191,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x9ee91F9f426fA633d227f7a9b000E28b9dfd8599"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000036cc7b13029b5dee4034745fb4f24034f3f2ffc6"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000197b04d6de02429b2be6"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17477988", 0)),
							TransactionHash: common.HexToHash("0x9247508d5690e97814a3f32b1f8ffda14c7b3404a7ae280eea6e93df2423dfd2"),
							Index:           192,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x28e4F3a7f651294B9564800b2D01f35189A5bFbE"),
							Topics: []common.Hash{
								common.HexToHash("0x103fed9db65eac19c4d870f49ab7520fe03b99f1838e5996caf47e9e43308392"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000292820"),
								common.HexToHash("0x0000000000000000000000008397259c983751daf40400790063935a11afa28a"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000c7dd5c30dca04f487c9ede0c5ac580c91587fc660000000000000000000000000833f5bd45803e05ef54e119a77e463ce6b1a9630000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000068b3197bd5c57243de395a000000000000000000000000000000000000000000707ba0de1f12ad6ba79d3d"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17477988", 0)),
							TransactionHash: common.HexToHash("0x9247508d5690e97814a3f32b1f8ffda14c7b3404a7ae280eea6e93df2423dfd2"),
							Index:           193,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x9ee91F9f426fA633d227f7a9b000E28b9dfd8599"),
							Topics: []common.Hash{
								common.HexToHash("0x4318b22a7b774533f1c9cd7102530d96faffc18ef44a1ecb56abc9a55d49fd8b"),
								common.HexToHash("0x00000000000000000000000036cc7b13029b5dee4034745fb4f24034f3f2ffc6"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000197b04d6de02429b2be6"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17477988", 0)),
							TransactionHash: common.HexToHash("0x9247508d5690e97814a3f32b1f8ffda14c7b3404a7ae280eea6e93df2423dfd2"),
							Index:           194,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x9247508d5690e97814a3f32b1f8ffda14c7b3404a7ae280eea6e93df2423dfd2"),
						TransactionIndex: 121,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x9247508d5690e97814a3f32b1f8ffda14c7b3404a7ae280eea6e93df2423dfd2",
				Network: filter.NetworkEthereum,
				Index:   121,
				From:    "0x36cc7B13029B5DEe4034745FB4F24034f3F2ffc6",
				To:      lido.AddressStakedMATIC.String(),
				Type:    filter.TypeCollectibleMint,
				Calldata: &schema.Calldata{
					FunctionHash: "0xccc143b8",
				},
				Platform: lo.ToPtr(filter.PlatformLido),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("18695859673908466")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeLiquidity,
						Platform: filter.PlatformLido.String(),
						From:     "0x36cc7B13029B5DEe4034745FB4F24034f3F2ffc6",
						To:       lido.AddressStakedMATICWithdrawalNFT.String(),
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquidityRemove,
							Tokens: []metadata.Token{
								{
									Address:  lo.ToPtr("0x9ee91F9f426fA633d227f7a9b000E28b9dfd8599"),
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("120328460302929861749734"))),
									Name:     "Staked MATIC",
									Symbol:   "stMATIC",
									Decimals: 18,
									Standard: metadata.StandardERC20,
								},
							},
						},
					},
					{
						Type:     filter.TypeCollectibleMint,
						Platform: filter.PlatformLido.String(),
						From:     ethereum.AddressGenesis.String(),
						To:       "0x36cc7B13029B5DEe4034745FB4F24034f3F2ffc6",
						Metadata: metadata.CollectibleTransfer{
							Address:  lo.ToPtr("0x60a91E2B7A1568f0848f3D43353C453730082E46"),
							ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1376"))),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
							Name:     "PoLido",
							Symbol:   "PLO",
							Standard: metadata.StandardERC721,
						},
					},
				},
				Status:    true,
				Timestamp: 1686743375,
			},
			wantError: require.NoError,
		},
		{
			name: "Burn PLO #1371 and receive MATIC",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x16f1f21c611d964651e7d4ca54d719dc028727fd1d85454740f1e3a3b854f95d"),
						ParentHash:   common.HexToHash("0x0f11cc07e71802b51d831286ea354ba45a2ac352a7c2bb6a366a3da84cd5a051"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x690B9A9E9aa1C9dB991C7721a92d351Db4FaC990"),
						Number:       lo.Must(new(big.Int).SetString("17484511", 0)),
						GasLimit:     30000000,
						GasUsed:      13435168,
						Timestamp:    1686822695,
						BaseFee:      lo.Must(new(big.Int).SetString("16870254903", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xfb4fc03b79772dc77bf6af4f35eaf1e32746eda8f2fdf89d40ca2d8a06e4225f"),
						From:      common.HexToAddress("0x22510fe99F63aE03BA792c21A29Ec10Fd87cae08"),
						Gas:       271292,
						GasPrice:  lo.Must(new(big.Int).SetString("17000000000", 10)),
						Hash:      common.HexToHash("0xfb4fc03b79772dc77bf6af4f35eaf1e32746eda8f2fdf89d40ca2d8a06e4225f"),
						Input:     hexutil.MustDecode("0x46e04a2f000000000000000000000000000000000000000000000000000000000000055b"),
						To:        lo.ToPtr(common.HexToAddress("0x9ee91F9f426fA633d227f7a9b000E28b9dfd8599")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x16f1f21c611d964651e7d4ca54d719dc028727fd1d85454740f1e3a3b854f95d"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17484511", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 6625603,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3f5476a00"),
						GasUsed:           216737,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x60a91E2B7A1568f0848f3D43353C453730082E46"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x00000000000000000000000022510fe99f63ae03ba792c21a29ec10fd87cae08"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000055b"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("17484511", 0)),
							TransactionHash: common.HexToHash("0xfb4fc03b79772dc77bf6af4f35eaf1e32746eda8f2fdf89d40ca2d8a06e4225f"),
							Index:           171,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x60a91E2B7A1568f0848f3D43353C453730082E46"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000022510fe99f63ae03ba792c21a29ec10fd87cae08"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000055b"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("17484511", 0)),
							TransactionHash: common.HexToHash("0xfb4fc03b79772dc77bf6af4f35eaf1e32746eda8f2fdf89d40ca2d8a06e4225f"),
							Index:           172,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000005e3ef299fddf15eaa0432e6e66473ace8c13d908"),
								common.HexToHash("0x0000000000000000000000009ee91f9f426fa633d227f7a9b000e28b9dfd8599"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000009e8159b8ab3a32ff93"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17484511", 0)),
							TransactionHash: common.HexToHash("0xfb4fc03b79772dc77bf6af4f35eaf1e32746eda8f2fdf89d40ca2d8a06e4225f"),
							Index:           173,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x6dF5CB08d3f0193C768C8A01f42ac4424DC5086b"),
							Topics: []common.Hash{
								common.HexToHash("0x7beb9bef91040fcf7607c78d4fc668370a9036788d7e6f202a4a2db5b0c28c80"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000075"),
								common.HexToHash("0x0000000000000000000000009ee91f9f426fa633d227f7a9b000e28b9dfd8599"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000009e8159b8ab3a32ff9300000000000000000000000000000000000000000000000000000000000000c2"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17484511", 0)),
							TransactionHash: common.HexToHash("0xfb4fc03b79772dc77bf6af4f35eaf1e32746eda8f2fdf89d40ca2d8a06e4225f"),
							Index:           174,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000009ee91f9f426fa633d227f7a9b000e28b9dfd8599"),
								common.HexToHash("0x00000000000000000000000022510fe99f63ae03ba792c21a29ec10fd87cae08"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000009e8159b8ab3a32ff93"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17484511", 0)),
							TransactionHash: common.HexToHash("0xfb4fc03b79772dc77bf6af4f35eaf1e32746eda8f2fdf89d40ca2d8a06e4225f"),
							Index:           175,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x9ee91F9f426fA633d227f7a9b000E28b9dfd8599"),
							Topics: []common.Hash{
								common.HexToHash("0xaca94a3466fab333b79851ab29b0715612740e4ae0d891ef8e9bd2a1bf5e24dd"),
								common.HexToHash("0x00000000000000000000000022510fe99f63ae03ba792c21a29ec10fd87cae08"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000055b"),
								common.HexToHash("0x00000000000000000000000000000000000000000000009e8159b8ab3a32ff93"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17484511", 0)),
							TransactionHash: common.HexToHash("0xfb4fc03b79772dc77bf6af4f35eaf1e32746eda8f2fdf89d40ca2d8a06e4225f"),
							Index:           176,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xfb4fc03b79772dc77bf6af4f35eaf1e32746eda8f2fdf89d40ca2d8a06e4225f"),
						TransactionIndex: 83,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0xfb4fc03b79772dc77bf6af4f35eaf1e32746eda8f2fdf89d40ca2d8a06e4225f",
				Network: filter.NetworkEthereum,
				Index:   83,
				From:    "0x22510fe99F63aE03BA792c21A29Ec10Fd87cae08",
				To:      lido.AddressStakedMATIC.String(),
				Type:    filter.TypeCollectibleBurn,
				Calldata: &schema.Calldata{
					FunctionHash: "0x46e04a2f",
				},
				Platform: lo.ToPtr(filter.PlatformLido),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("3684529000000000")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeCollectibleBurn,
						Platform: filter.PlatformLido.String(),
						From:     "0x22510fe99F63aE03BA792c21A29Ec10Fd87cae08",
						To:       ethereum.AddressGenesis.String(),
						Metadata: metadata.CollectibleTransfer{
							Address:  lo.ToPtr(lido.AddressStakedMATICWithdrawalNFT.String()),
							ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1371"))),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
							Name:     "PoLido",
							Symbol:   "PLO",
							Standard: metadata.StandardERC721,
						},
					},
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformLido.String(),
						From:     lido.AddressStakedMATIC.String(),
						To:       "0x22510fe99F63aE03BA792c21A29Ec10Fd87cae08",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr(lido.AddressMATIC.String()),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2923906247595484446611"))),
							Name:     "Matic Token",
							Symbol:   "MATIC",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
				},
				Status:    true,
				Timestamp: 1686822695,
			},
			wantError: require.NoError,
		},
		{
			name: "Wrap stETH for wstETH",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xbd8b0c681de0ca40260dd063fff4ccc3e4f60f27980aa3fcb0a3b507dbaa289b"),
						ParentHash:   common.HexToHash("0x578def7b942d544ba271d88afed21f0bc58551329e0da244846f5c2eb7264fba"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xDAFEA492D9c6733ae3d56b7Ed1ADB60692c98Bc5"),
						Number:       lo.Must(new(big.Int).SetString("17488034", 0)),
						GasLimit:     30000000,
						GasUsed:      12672298,
						Timestamp:    1686865559,
						BaseFee:      lo.Must(new(big.Int).SetString("16169946604", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x89b84ef9d6cd376cffd554d7c63ff3c301eea7ffe19371aae59e7b0e7aeee2b5"),
						From:      common.HexToAddress("0x8F1853b7891955791513A5E1262258D42468005a"),
						Gas:       128930,
						GasPrice:  lo.Must(new(big.Int).SetString("17669946604", 10)),
						Hash:      common.HexToHash("0x89b84ef9d6cd376cffd554d7c63ff3c301eea7ffe19371aae59e7b0e7aeee2b5"),
						Input:     hexutil.MustDecode("0xea598cb000000000000000000000000000000000000000000000000821ab0d441497ffff"),
						To:        lo.ToPtr(common.HexToAddress("0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xbd8b0c681de0ca40260dd063fff4ccc3e4f60f27980aa3fcb0a3b507dbaa289b"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17488034", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 4138993,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x41d35fcec"),
						GasUsed:           118575,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x0000000000000000000000008f1853b7891955791513a5e1262258d42468005a"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000007349f7217a106fd71"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17488034", 0)),
							TransactionHash: common.HexToHash("0x89b84ef9d6cd376cffd554d7c63ff3c301eea7ffe19371aae59e7b0e7aeee2b5"),
							Index:           95,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x0000000000000000000000008f1853b7891955791513a5e1262258d42468005a"),
								common.HexToHash("0x0000000000000000000000007f39c581f595b53c5cb19bd0b3f8da6c935e2ca0"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000001"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17488034", 0)),
							TransactionHash: common.HexToHash("0x89b84ef9d6cd376cffd554d7c63ff3c301eea7ffe19371aae59e7b0e7aeee2b5"),
							Index:           96,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000008f1853b7891955791513a5e1262258d42468005a"),
								common.HexToHash("0x0000000000000000000000007f39c581f595b53c5cb19bd0b3f8da6c935e2ca0"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000821ab0d441497ffff"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17488034", 0)),
							TransactionHash: common.HexToHash("0x89b84ef9d6cd376cffd554d7c63ff3c301eea7ffe19371aae59e7b0e7aeee2b5"),
							Index:           97,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
							Topics: []common.Hash{
								common.HexToHash("0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb"),
								common.HexToHash("0x0000000000000000000000008f1853b7891955791513a5e1262258d42468005a"),
								common.HexToHash("0x0000000000000000000000007f39c581f595b53c5cb19bd0b3f8da6c935e2ca0"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000007349f7217a106fd71"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17488034", 0)),
							TransactionHash: common.HexToHash("0x89b84ef9d6cd376cffd554d7c63ff3c301eea7ffe19371aae59e7b0e7aeee2b5"),
							Index:           98,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x89b84ef9d6cd376cffd554d7c63ff3c301eea7ffe19371aae59e7b0e7aeee2b5"),
						TransactionIndex: 47,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x89b84ef9d6cd376cffd554d7c63ff3c301eea7ffe19371aae59e7b0e7aeee2b5",
				Network: filter.NetworkEthereum,
				Index:   47,
				From:    "0x8F1853b7891955791513A5E1262258D42468005a",
				To:      lido.AddressWrappedStakedETH.String(),
				Type:    filter.TypeExchangeSwap,
				Calldata: &schema.Calldata{
					FunctionHash: "0xea598cb0",
				},
				Platform: lo.ToPtr(filter.PlatformLido),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("2095213918569300")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeSwap,
						Platform: filter.PlatformLido.String(),
						From:     "0x8F1853b7891955791513A5E1262258D42468005a",
						To:       "0x8F1853b7891955791513A5E1262258D42468005a",
						Metadata: metadata.ExchangeSwap{
							From: metadata.Token{
								Address:  lo.ToPtr("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("149999999999999999999"))),
								Name:     "Liquid staked Ether 2.0",
								Symbol:   "stETH",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
							To: metadata.Token{
								Address:  lo.ToPtr("0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("132919083373047512433"))),
								Name:     "Wrapped liquid staked Ether 2.0",
								Symbol:   "wstETH",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1686865559,
			},
			wantError: require.NoError,
		},
		{
			name: "Unwrap wstETH for stETH",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x1ea3a25928a89fbd34b60ba327fa4047886cff127791f1fe25533fbcdcd775d5"),
						ParentHash:   common.HexToHash("0x11773d8a3b20e2069fa3664d0075b277d601d67aeb07405f1974bc74c3f07f13"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("17487958", 0)),
						GasLimit:     30000000,
						GasUsed:      14640645,
						Timestamp:    1686864647,
						BaseFee:      lo.Must(new(big.Int).SetString("18015954095", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xdef8ed7f55480f66ed52be86b9b66c86154e6a5df289e894da9ecdada5b603cb"),
						From:      common.HexToAddress("0x8375eC8bB9902512a63d488C09548e40e13FB81F"),
						Gas:       122985,
						GasPrice:  lo.Must(new(big.Int).SetString("19515954095", 10)),
						Hash:      common.HexToHash("0xdef8ed7f55480f66ed52be86b9b66c86154e6a5df289e894da9ecdada5b603cb"),
						Input:     hexutil.MustDecode("0xde0e9a3e000000000000000000000000000000000000000000000075539da0d08d66ff59"),
						To:        lo.ToPtr(common.HexToAddress("0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x1ea3a25928a89fbd34b60ba327fa4047886cff127791f1fe25533fbcdcd775d5"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17487958", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 5302339,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x48b3dd3af"),
						GasUsed:           107666,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000008375ec8bb9902512a63d488c09548e40e13fb81f"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000075539da0d08d66ff59"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17487958", 0)),
							TransactionHash: common.HexToHash("0xdef8ed7f55480f66ed52be86b9b66c86154e6a5df289e894da9ecdada5b603cb"),
							Index:           143,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000007f39c581f595b53c5cb19bd0b3f8da6c935e2ca0"),
								common.HexToHash("0x0000000000000000000000008375ec8bb9902512a63d488c09548e40e13fb81f"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000008467605b0e35376fa2"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17487958", 0)),
							TransactionHash: common.HexToHash("0xdef8ed7f55480f66ed52be86b9b66c86154e6a5df289e894da9ecdada5b603cb"),
							Index:           144,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
							Topics: []common.Hash{
								common.HexToHash("0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb"),
								common.HexToHash("0x0000000000000000000000007f39c581f595b53c5cb19bd0b3f8da6c935e2ca0"),
								common.HexToHash("0x0000000000000000000000008375ec8bb9902512a63d488c09548e40e13fb81f"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000075539da0d08d66ff58"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17487958", 0)),
							TransactionHash: common.HexToHash("0xdef8ed7f55480f66ed52be86b9b66c86154e6a5df289e894da9ecdada5b603cb"),
							Index:           145,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xdef8ed7f55480f66ed52be86b9b66c86154e6a5df289e894da9ecdada5b603cb"),
						TransactionIndex: 71,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0xdef8ed7f55480f66ed52be86b9b66c86154e6a5df289e894da9ecdada5b603cb",
				Network: filter.NetworkEthereum,
				Index:   71,
				From:    "0x8375eC8bB9902512a63d488C09548e40e13FB81F",
				To:      lido.AddressWrappedStakedETH.String(),
				Type:    filter.TypeExchangeSwap,
				Calldata: &schema.Calldata{
					FunctionHash: "0xde0e9a3e",
				},
				Platform: lo.ToPtr(filter.PlatformLido),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("2101204713592270")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeSwap,
						Platform: filter.PlatformLido.String(),
						From:     "0x8375eC8bB9902512a63d488C09548e40e13FB81F",
						To:       "0x8375eC8bB9902512a63d488C09548e40e13FB81F",
						Metadata: metadata.ExchangeSwap{
							From: metadata.Token{
								Address:  lo.ToPtr("0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2164294205318095109977"))),
								Name:     "Wrapped liquid staked Ether 2.0",
								Symbol:   "wstETH",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
							To: metadata.Token{
								Address:  lo.ToPtr("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2442419271629912108962"))),
								Name:     "Liquid staked Ether 2.0",
								Symbol:   "stETH",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1686864647,
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
			require.Equal(t, testcase.want, feed)

			t.Log(feed)
		})
	}
}
