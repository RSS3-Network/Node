package cow_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/decentralized/contract/cow"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	workerx "github.com/rss3-network/node/schema/worker/decentralized"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestWorker_Cow(t *testing.T) {
	t.Parallel()

	type arguments struct {
		task   *source.Task
		config *config.Module
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      *activityx.Activity
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "Cow Single Token Swap (CHEX to ETH)",
			arguments: struct {
				task   *source.Task
				config *config.Module
			}{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:       common.HexToHash("0x6f1cad729efab3ce9a25f41f9c63204d962625b546c535353bb303c52b946ce7"),
						ParentHash: common.HexToHash("0x9212c7703b30ee7fdaf9fc944dbcef58e4cb681fa7fdddbf3526717b686e3a7b"),
						UncleHash:  common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:   common.HexToAddress("0x4838b106fce9647bdf1e7877bf73ce8b0bad5f97"),
						Number:     big.NewInt(19789140),
						GasLimit:   30000000,
						GasUsed:    12512871,
						Timestamp:  1714645623,
						BaseFee:    big.NewInt(462212328),
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x6f1cad729efab3ce9a25f41f9c63204d962625b546c535353bb303c52b946ce7"),
						From:      common.HexToAddress("0x16c473448e770ff647c69cbe19e28528877fba1b"),
						Gas:       443782,
						GasPrice:  big.NewInt(507933742),
						Hash:      common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
						Input:     hexutil.MustDecode("0x13d79a0b0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000003e000000000000000000000000000000000000000000000000000000000000000040000000000000000000000009ce84f6a69986a83d92c324df10bc8e64771030f000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee0000000000000000000000009ce84f6a69986a83d92c324df10bc8e64771030f000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000032e3bea1948eccef0ed10000000000000000000000000000000000000000000000002619e46cf01de1df000000000000000000000000000000000000000000000793e048d362d7acc3110000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000003000000000000000000000000d6bc23c282e2593f1bca3fd199148cbfa3f2f566000000000000000000000000000000000000000000000793e048d362d7acc31100000000000000000000000000000000000000000000000025efdcfe1f7048f9000000000000000000000000000000000000000000000000000000006631c69b"),
						To:        lo.ToPtr(common.HexToAddress("0x9008D19f58AAbD9eD0D60971565AA8510560ab41")),
						Value:     big.NewInt(0),
						Type:      2,
						ChainID:   big.NewInt(1),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x6f1cad729efab3ce9a25f41f9c63204d962625b546c535353bb303c52b946ce7"),
						BlockNumber:       big.NewInt(19789140),
						ContractAddress:   nil,
						CumulativeGasUsed: 2250719,
						EffectiveGasPrice: big.NewInt(507933742),
						GasUsed:           214600,
						Status:            1,
						TransactionHash:   common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
						TransactionIndex:  13,
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0x9008d19f58aabd9ed0d60971565aa8510560ab41"),
								Topics: []common.Hash{
									common.HexToHash("0xa07a543ab8a018198e99ca0184c93fe9050a79400a0a723441f84de1d972cc17"),
									common.HexToHash("0x000000000000000000000000d6bc23c282e2593f1bca3fd199148cbfa3f2f566"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000009ce84f6a69986a83d92c324df10bc8e64771030f000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee000000000000000000000000000000000000000000000793e048d362d7acc3110000000000000000000000000000000000000000000000002619e46cf01de1df000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000003833b3cf17f878352ea6b1d0e59a7a83f1c14cb66384db1da938e3b4de18337552d6bc23c282e2593f1bca3fd199148cbfa3f2f5666631c69b0000000000000000"),
								BlockNumber:     big.NewInt(19789140),
								TransactionHash: common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
								//TxIndex:     13,
								BlockHash: common.HexToHash("0x6f1cad729efab3ce9a25f41f9c63204d962625b546c535353bb303c52b946ce7"),
								Index:     75,
								Removed:   false,
							},
							{
								Address: common.HexToAddress("0x9ce84f6a69986a83d92c324df10bc8e64771030f"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x000000000000000000000000d6bc23c282e2593f1bca3fd199148cbfa3f2f566"),
									common.HexToHash("0x0000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab41"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000793e048d362d7acc311"),
								BlockNumber:     big.NewInt(19789140),
								TransactionHash: common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
								//TxIndex:     13,
								BlockHash: common.HexToHash("0x6f1cad729efab3ce9a25f41f9c63204d962625b546c535353bb303c52b946ce7"),
								Index:     76,
								Removed:   false,
							},
							{
								Address: common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x000000000000000000000000d3e9895230e8fb1460852f6cda3c4b926fbc29d8"),
									common.HexToHash("0x0000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab41"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000002619e46cf01de1de"),
								BlockNumber:     big.NewInt(19789140),
								TransactionHash: common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
								//TxIndex:     13,
								BlockHash: common.HexToHash("0x6f1cad729efab3ce9a25f41f9c63204d962625b546c535353bb303c52b946ce7"),
								Index:     77,
								Removed:   false,
							},
							{
								Address: common.HexToAddress("0x9ce84f6a69986a83d92c324df10bc8e64771030f"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x0000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab41"),
									common.HexToHash("0x000000000000000000000000d3e9895230e8fb1460852f6cda3c4b926fbc29d8"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000792f3f2176c7c41b9ee"),
								BlockNumber:     big.NewInt(19789140),
								TransactionHash: common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
								//TxIndex:     13,
								BlockHash: common.HexToHash("0x6f1cad729efab3ce9a25f41f9c63204d962625b546c535353bb303c52b946ce7"),
								Index:     78,
								Removed:   false,
							},
							{
								Address: common.HexToAddress("0xd3e9895230e8fb1460852f6cda3c4b926fbc29d8"),
								Topics: []common.Hash{
									common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"),
									common.HexToHash("0x00000000000000000000000000000000be48a3000b818e9615d85aacfed4ca97"),
									common.HexToHash("0x0000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab41"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000792f3f2176c7c41b9eeffffffffffffffffffffffffffffffffffffffffffffffffd9e61b930fe21e220000000000000000000000000000000000000000023c99f2ed4e4c321d1054d000000000000000000000000000000000000000000000043f625f52a6f3310475fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8da7"),
								BlockNumber:     big.NewInt(19789140),
								TransactionHash: common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
								//TxIndex:     13,
								BlockHash: common.HexToHash("0x6f1cad729efab3ce9a25f41f9c63204d962625b546c535353bb303c52b946ce7"),
								Index:     79,
								Removed:   false,
							},
						},
					},
				},
				config: &config.Module{
					Network: network.Ethereum,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Ethereum),
					},
				},
			},
			want: &activityx.Activity{
				ID:       "0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b",
				Network:  network.Ethereum,
				Index:    13,
				Platform: "Cow",
				From:     "0x16c473448e770ff647c69cbe19e28528877fba1b",
				To:       "0x9008D19f58AAbD9eD0D60971565AA8510560ab41",
				Type:     typex.ExchangeSwap,
				Status:   true,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x13d79a0b",
				},
				Timestamp: 1714645623,
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("109002581033200")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.ExchangeSwap,
						Platform: workerx.PlatformCow.String(),
						From:     "0xd6BC23c282e2593F1bCA3FD199148cbFA3F2F566",
						To:       "0xd6BC23c282e2593F1bCA3FD199148cbFA3F2F566",
						Metadata: metadata.ExchangeSwap{
							From: metadata.Token{
								Address:  lo.ToPtr("0x9Ce84F6A69986a83d92C324df10bC8E64771030f"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("35784398158607118287633"))),
								Name:     "Chintai Exchange Token",
								Symbol:   "CHEX",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
							To: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2745476604395119071"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
			},
			wantError: require.NoError,
		},
		{
			name: "Cow Multiple token Swaps(WETH - USDC, WETH - AAVE)",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:       common.HexToHash("0x420ddba0c2e81bd1413b53a9247d5aacd60e398e017c499eb3fd1826e222530e"),
						ParentHash: common.HexToHash("0xeb8440a3150f14d6f4c29d82cd3def31463cacffca49d18a1e0e940101b2d3a9"),
						UncleHash:  common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:   common.HexToAddress("0x4838b106fce9647bdf1e7877bf73ce8b0bad5f97"),
						Number:     big.NewInt(19789141),
						GasLimit:   30000000,
						GasUsed:    25341971,
						Timestamp:  1714645603,
						BaseFee:    big.NewInt(5590089183),
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x420ddba0c2e81bd1413b53a9247d5aacd60e398e017c499eb3fd1826e222530e"),
						From:      common.HexToAddress("0x16c473448e770ff647c69cbe19e28528877fba1b"),
						Gas:       1219506,
						GasPrice:  big.NewInt(6232011486),
						Hash:      common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
						Input:     hexutil.MustDecode("0x13d79a0b0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000002800000000000000000000000000000000000000000000000000000000000000b0000000000000000000000000000000000000000000000000000000000000000070000000000000000000000007fc66500c84a76ad7e9c93437bfc5ac33e2ddae9000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000007fc66500c84a76ad7e9c93437bfc5ac33e2ddae9000000000000000000000000000000000000000000000000000000000000000700000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000002cde9e4cbf690adf2d413694900000000000000000000000000000000000000000000002422e25377173ba43c00000000000000000000000000000000000000000000000000000005b46a15390000000000000000000000000000000000000000000000007155d41081b0aaaa000000000000000000000000000000000000000000000009ca5dcb3527fcec670000000000000000000000000000000000000000000000004563918244f40000"),
						To:        lo.ToPtr(common.HexToAddress("0x9008D19f58AAbD9eD0D60971565AA8510560ab41")),
						Value:     big.NewInt(0),
						Type:      2,
						ChainID:   big.NewInt(1),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x420ddba0c2e81bd1413b53a9247d5aacd60e398e017c499eb3fd1826e222530e"),
						BlockNumber:       lo.Must(new(big.Int).SetString("19789141", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 6874595,
						EffectiveGasPrice: big.NewInt(6232011486),
						GasUsed:           595400,
						Status:            1,
						TransactionHash:   common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
						TransactionIndex:  13,
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
								Topics: []common.Hash{
									common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"),
									common.HexToHash("0x00000000000000000000000040a50cf069e992aa4536211b23f286ef88752187"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000004563918244f40000"),
								BlockNumber:     big.NewInt(19789141),
								TransactionHash: common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
								//TxIndex:     13,
								BlockHash: common.HexToHash("0x420ddba0c2e81bd1413b53a9247d5aacd60e398e017c499eb3fd1826e222530e"),
								Index:     315,
								Removed:   false,
							},
							{
								Address: common.HexToAddress("0x9008d19f58aabd9ed0d60971565aa8510560ab41"),
								Topics: []common.Hash{
									common.HexToHash("0xed99827efb37016f2275f98c4bcf71c7551c75d59e9b450f79fa32e60be672c2"),
									common.HexToHash("0x00000000000000000000000040a50cf069e992aa4536211b23f286ef88752187"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000004c84c1c800000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     big.NewInt(19789141),
								TransactionHash: common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
								//TxIndex:     13,
								BlockHash: common.HexToHash("0x420ddba0c2e81bd1413b53a9247d5aacd60e398e017c499eb3fd1826e222530e"),
								Index:     316,
								Removed:   false,
							},
							{
								Address: common.HexToAddress("0x9008d19f58aabd9ed0d60971565aa8510560ab41"),
								Topics: []common.Hash{
									common.HexToHash("0xa07a543ab8a018198e99ca0184c93fe9050a79400a0a723441f84de1d972cc17"),
									common.HexToHash("0x000000000000000000000000087f3f5dee3fa664a57129108cd064823620f051"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000000000000000000000000000007155d41081b0aaaa00000000000000000000000000000000000000000000000000000005b46a1539000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000003815a14f307cc9338e01dcead7849de6dbac24eb0831aa4ee86026659c3defc73d087f3f5dee3fa664a57129108cd064823620f0516631da520000000000000000"),
								BlockNumber:     big.NewInt(19789141),
								TransactionHash: common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
								//TxIndex:     13,
								BlockHash: common.HexToHash("0x420ddba0c2e81bd1413b53a9247d5aacd60e398e017c499eb3fd1826e222530e"),
								Index:     317,
								Removed:   false,
							},
							{
								Address: common.HexToAddress("0x9008d19f58aabd9ed0d60971565aa8510560ab41"),
								Topics: []common.Hash{
									common.HexToHash("0xa07a543ab8a018198e99ca0184c93fe9050a79400a0a723441f84de1d972cc17"),
									common.HexToHash("0x00000000000000000000000040a50cf069e992aa4536211b23f286ef88752187"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000007fc66500c84a76ad7e9c93437bfc5ac33e2ddae90000000000000000000000000000000000000000000000004563918244f40000000000000000000000000000000000000000000000000009ca5dcb3527fcec67000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000380e46515d4d6713dc9b38e314fc6d8033826ebef22c764e053e193d3696e520f740a50cf069e992aa4536211b23f286ef88752187ffffffff0000000000000000"),
								BlockNumber:     big.NewInt(19789141),
								TransactionHash: common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
								//TxIndex:     13,
								BlockHash: common.HexToHash("0x420ddba0c2e81bd1413b53a9247d5aacd60e398e017c499eb3fd1826e222530e"),
								Index:     318,
								Removed:   false,
							},
							{
								Address: common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x000000000000000000000000087f3f5dee3fa664a57129108cd064823620f051"),
									common.HexToHash("0x0000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab41"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000007155d41081b0aaaa"),
								BlockNumber:     big.NewInt(19789141),
								TransactionHash: common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
								//TxIndex:     13,
								BlockHash: common.HexToHash("0x420ddba0c2e81bd1413b53a9247d5aacd60e398e017c499eb3fd1826e222530e"),
								Index:     319,
								Removed:   false,
							},
							{
								Address: common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x00000000000000000000000040a50cf069e992aa4536211b23f286ef88752187"),
									common.HexToHash("0x0000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab41"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000004563918244f40000"),
								BlockNumber:     big.NewInt(19789141),
								TransactionHash: common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
								//TxIndex:     13,
								BlockHash: common.HexToHash("0x420ddba0c2e81bd1413b53a9247d5aacd60e398e017c499eb3fd1826e222530e"),
								Index:     320,
								Removed:   false,
							},
							{
								Address: common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x00000000000000000000000051c72848c68a965f66fa7a88855f9f7784502a7f"),
									common.HexToHash("0x0000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab41"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000005b4c01c2a"),
								BlockNumber:     big.NewInt(19789141),
								TransactionHash: common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
								//TxIndex:     13,
								BlockHash: common.HexToHash("0x420ddba0c2e81bd1413b53a9247d5aacd60e398e017c499eb3fd1826e222530e"),
								Index:     321,
								Removed:   false,
							},
							{
								Address: common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x0000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab41"),
									common.HexToHash("0x00000000000000000000000051c72848c68a965f66fa7a88855f9f7784502a7f"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000007155d41081b0aaaa"),
								BlockNumber:     big.NewInt(19789141),
								TransactionHash: common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
								//TxIndex:     13,
								BlockHash: common.HexToHash("0x420ddba0c2e81bd1413b53a9247d5aacd60e398e017c499eb3fd1826e222530e"),
								Index:     322,
								Removed:   false,
							},
							{
								Address: common.HexToAddress("0xbeb09000fa59627dc02bb55448ac1893eaa501a5"),
								Topics: []common.Hash{
									common.HexToHash("0xc59522161f93d59c8c4520b0e7a3635fb7544133275be812a4ea970f4f14251b"),
								},
								Data:            hexutil.MustDecode("0xc07c6b6e953a51a32d28aa834bf252c77d32f986582aeefbfe14a23338595a7f"),
								BlockNumber:     big.NewInt(19789141),
								TransactionHash: common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
								//TxIndex:     13,
								BlockHash: common.HexToHash("0x420ddba0c2e81bd1413b53a9247d5aacd60e398e017c499eb3fd1826e222530e"),
								Index:     323,
								Removed:   false,
							},
							{
								Address: common.HexToAddress("0x9008d19f58aabd9ed0d60971565aa8510560ab41"),
								Topics: []common.Hash{
									common.HexToHash("0xed99827efb37016f2275f98c4bcf71c7551c75d59e9b450f79fa32e60be672c2"),
									common.HexToHash("0x000000000000000000000000beb09000fa59627dc02bb55448ac1893eaa501a5"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000000024a305f200000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     big.NewInt(19789141),
								TransactionHash: common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
								//TxIndex:     13,
								BlockHash: common.HexToHash("0x420ddba0c2e81bd1413b53a9247d5aacd60e398e017c499eb3fd1826e222530e"),
								Index:     324,
								Removed:   false,
							},
						},
					},
				},
				config: &config.Module{
					Network: network.Ethereum,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Ethereum),
					},
				},
			},
			want: &activityx.Activity{
				ID:       "0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a",
				Network:  network.Ethereum,
				Index:    13,
				Platform: "Cow",
				From:     "0x16c473448e770ff647c69cbe19e28528877fba1b",
				To:       "0x9008D19f58AAbD9eD0D60971565AA8510560ab41",
				Type:     typex.ExchangeSwap,
				Status:   true,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x13d79a0b",
				},
				Timestamp: 1714645603,
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("3710539638764400")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.ExchangeSwap,
						Platform: workerx.PlatformCow.String(),
						From:     "0x087F3F5dEE3FA664a57129108Cd064823620F051",
						To:       "0x087F3F5dEE3FA664a57129108Cd064823620F051",
						Metadata: metadata.ExchangeSwap{
							From: metadata.Token{
								Address:  lo.ToPtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("8166666666666666666"))),
								Name:     "Wrapped Ether",
								Symbol:   "WETH",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
							To: metadata.Token{
								Address:  lo.ToPtr("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("24501687609"))),
								Name:     "USD Coin",
								Symbol:   "USDC",
								Decimals: 6,
								Standard: metadata.StandardERC20,
							},
						},
					},
					{
						Type:     typex.ExchangeSwap,
						Platform: workerx.PlatformCow.String(),
						From:     "0x40A50cf069e992AA4536211B23F286eF88752187",
						To:       "0x40A50cf069e992AA4536211B23F286eF88752187",
						Metadata: metadata.ExchangeSwap{
							From: metadata.Token{
								Address:  lo.ToPtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("5000000000000000000"))),
								Name:     "Wrapped Ether",
								Symbol:   "WETH",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
							To: metadata.Token{
								Address:  lo.ToPtr("0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("180602731261046090855"))),
								Name:     "Aave Token",
								Symbol:   "AAVE",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
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

			activity, err := instance.Transform(ctx, testcase.arguments.task)
			testcase.wantError(t, err)

			require.Equal(t, testcase.want, activity)
		})
	}
}
