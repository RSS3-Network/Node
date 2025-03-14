package cow_test

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/v2/config"
	source "github.com/rss3-network/node/v2/internal/engine/protocol/ethereum"
	worker "github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/cow"
	"github.com/rss3-network/node/v2/provider/ethereum"
	"github.com/rss3-network/node/v2/provider/ethereum/endpoint"
	workerx "github.com/rss3-network/node/v2/schema/worker/decentralized"
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
						Hash:         common.HexToHash("0x6f1cad729efab3ce9a25f41f9c63204d962625b546c535353bb303c52b946ce7"),
						ParentHash:   common.HexToHash("0x9212c7703b30ee7fdaf9fc944dbcef58e4cb681fa7fdddbf3526717b686e3a7b"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"),
						Number:       lo.Must(new(big.Int).SetString("19772788", 0)),
						GasLimit:     30000000,
						GasUsed:      12566119,
						Timestamp:    1714536407,
						BaseFee:      lo.Must(new(big.Int).SetString("7390004874", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
						From:      common.HexToAddress("0x16C473448E770Ff647c69CBe19e28528877fba1B"),
						Gas:       442758,
						GasPrice:  lo.Must(new(big.Int).SetString("7975644964", 10)),
						Hash:      common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
						Input:     hexutil.MustDecode("0x13d79a0b0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000003e000000000000000000000000000000000000000000000000000000000000000040000000000000000000000009ce84f6a69986a83d92c324df10bc8e64771030f000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee0000000000000000000000009ce84f6a69986a83d92c324df10bc8e64771030f000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000032e3bea1948eccef0ed10000000000000000000000000000000000000000000000002619e46cf01de1df000000000000000000000000000000000000000000000793e048d362d7acc3110000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000003000000000000000000000000d6bc23c282e2593f1bca3fd199148cbfa3f2f566000000000000000000000000000000000000000000000793e048d362d7acc31100000000000000000000000000000000000000000000000025efdcfe1f7048f9000000000000000000000000000000000000000000000000000000006631c69ba65aca2aeb89ba5841b0891a0e84a5b3b3db451fbb66f5210b63231845ae953e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000793e048d362d7acc311000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000415c7d03ee684f73dc159763ae5affd6f318fd84a3a7a5f5c78cd88251456e8c98728f27036ad04dba125410c8b0a76bc94f15fe1c6929e666cf8f60b2621cc05b1b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000280000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000be48a3000b818e9615d85aacfed4ca9700000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000004f0000000101010000000000000792f3f2176c7c41b9eed3e9895230e8fb1460852f6cda3c4b926fbc29d8000000000000000025f0000000000000019ce84f6a69986a83d92c324df10bc8e64771030f0000000000000000000000000000000000000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000242e1a7d4d0000000000000000000000000000000000000000000000002619e46cf01de1df0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000868698"),
						To:        lo.ToPtr(common.HexToAddress("0x9008D19f58AAbD9eD0D60971565AA8510560ab41")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x6f1cad729efab3ce9a25f41f9c63204d962625b546c535353bb303c52b946ce7"),
						BlockNumber:       lo.Must(new(big.Int).SetString("19772788", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 2255839,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x1db62af24"),
						GasUsed:           215120,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x9008D19f58AAbD9eD0D60971565AA8510560ab41"),
							Topics: []common.Hash{
								common.HexToHash("0xa07a543ab8a018198e99ca0184c93fe9050a79400a0a723441f84de1d972cc17"),
								common.HexToHash("0x000000000000000000000000d6bc23c282e2593f1bca3fd199148cbfa3f2f566"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000009ce84f6a69986a83d92c324df10bc8e64771030f000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee000000000000000000000000000000000000000000000793e048d362d7acc3110000000000000000000000000000000000000000000000002619e46cf01de1df000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000003833b3cf17f878352ea6b1d0e59a7a83f1c14cb66384db1da938e3b4de18337552d6bc23c282e2593f1bca3fd199148cbfa3f2f5666631c69b0000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19772788", 0)),
							TransactionHash: common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
							Index:           75,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x9Ce84F6A69986a83d92C324df10bC8E64771030f"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000d6bc23c282e2593f1bca3fd199148cbfa3f2f566"),
								common.HexToHash("0x0000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab41"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000793e048d362d7acc311"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19772788", 0)),
							TransactionHash: common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
							Index:           76,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000d3e9895230e8fb1460852f6cda3c4b926fbc29d8"),
								common.HexToHash("0x0000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab41"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000002619e46cf01de1de"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19772788", 0)),
							TransactionHash: common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
							Index:           77,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x9Ce84F6A69986a83d92C324df10bC8E64771030f"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab41"),
								common.HexToHash("0x000000000000000000000000d3e9895230e8fb1460852f6cda3c4b926fbc29d8"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000792f3f2176c7c41b9ee"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19772788", 0)),
							TransactionHash: common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
							Index:           78,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xD3e9895230E8FB1460852F6cDA3C4B926FbC29D8"),
							Topics: []common.Hash{
								common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"),
								common.HexToHash("0x00000000000000000000000000000000be48a3000b818e9615d85aacfed4ca97"),
								common.HexToHash("0x0000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab41"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000792f3f2176c7c41b9eeffffffffffffffffffffffffffffffffffffffffffffffffd9e61b930fe21e220000000000000000000000000000000000000000023c99f2ed4e4c321d1054d000000000000000000000000000000000000000000000043f625f52a6f3310475fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8da7"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19772788", 0)),
							TransactionHash: common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
							Index:           79,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x9008D19f58AAbD9eD0D60971565AA8510560ab41"),
							Topics: []common.Hash{
								common.HexToHash("0xed99827efb37016f2275f98c4bcf71c7551c75d59e9b450f79fa32e60be672c2"),
								common.HexToHash("0x00000000000000000000000000000000be48a3000b818e9615d85aacfed4ca97"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19772788", 0)),
							TransactionHash: common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
							Index:           80,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65"),
								common.HexToHash("0x0000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab41"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000002619e46cf01de1df"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19772788", 0)),
							TransactionHash: common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
							Index:           81,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x9008D19f58AAbD9eD0D60971565AA8510560ab41"),
							Topics: []common.Hash{
								common.HexToHash("0xed99827efb37016f2275f98c4bcf71c7551c75d59e9b450f79fa32e60be672c2"),
								common.HexToHash("0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000002e1a7d4d00000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19772788", 0)),
							TransactionHash: common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
							Index:           82,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x9008D19f58AAbD9eD0D60971565AA8510560ab41"),
							Topics: []common.Hash{
								common.HexToHash("0x40338ce1a7c49204f0099533b1e9a7ee0a3d261f84974ab7af36105b8c4e9db4"),
								common.HexToHash("0x00000000000000000000000016c473448e770ff647c69cbe19e28528877fba1b"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("19772788", 0)),
							TransactionHash: common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
							Index:           83,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xa5dfcff7dcc02511d02922f9bcf5f324147fc9e9f484d16b9ba8ddb7a2001d8b"),
						TransactionIndex: 13,
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
				From:     "0x16C473448E770Ff647c69CBe19e28528877fba1B",
				To:       "0x9008D19f58AAbD9eD0D60971565AA8510560ab41",
				Type:     typex.ExchangeSwap,
				Platform: workerx.PlatformCow.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("1715720744655680")),
					Decimal: 18,
				},
				Calldata: &activityx.Calldata{
					FunctionHash: "0x13d79a0b",
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
				Status:    true,
				Timestamp: 1714536407,
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
						Hash:         common.HexToHash("0x420ddba0c2e81bd1413b53a9247d5aacd60e398e017c499eb3fd1826e222530e"),
						ParentHash:   common.HexToHash("0xeb8440a3150f14d6f4c29d82cd3def31463cacffca49d18a1e0e940101b2d3a9"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"),
						Number:       lo.Must(new(big.Int).SetString("19772757", 0)),
						GasLimit:     30000000,
						GasUsed:      25442323,
						Timestamp:    1714536035,
						BaseFee:      lo.Must(new(big.Int).SetString("5680726495", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
						From:      common.HexToAddress("0x16C473448E770Ff647c69CBe19e28528877fba1B"),
						Gas:       1236914,
						GasPrice:  lo.Must(new(big.Int).SetString("6265321182", 10)),
						Hash:      common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
						Input:     hexutil.MustDecode("0x13d79a0b0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000002800000000000000000000000000000000000000000000000000000000000000b0000000000000000000000000000000000000000000000000000000000000000070000000000000000000000007fc66500c84a76ad7e9c93437bfc5ac33e2ddae9000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000007fc66500c84a76ad7e9c93437bfc5ac33e2ddae9000000000000000000000000000000000000000000000000000000000000000700000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000002cde9e4cbf690adf2d413694900000000000000000000000000000000000000000000002422e25377173ba43c00000000000000000000000000000000000000000000000000000005b46a15390000000000000000000000000000000000000000000000007155d41081b0aaaa000000000000000000000000000000000000000000000009ca5dcb3527fcec670000000000000000000000000000000000000000000000004563918244f400000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000006a000000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000004000000000000000000000000087f3f5dee3fa664a57129108cd064823620f0510000000000000000000000000000000000000000000000007155d41081b0aaaa00000000000000000000000000000000000000000000000000000005008ecffa000000000000000000000000000000000000000000000000000000006631da521c3ebdf32bf66763f2e67af49eabcfd20dcdee2c1770c859c53cf2bfb8ae0652000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000007155d41081b0aaaa000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000004d8087f3f5dee3fa664a57129108cd064823620f0515fd7e97dc078f884a2676e1345748b1feace7b0abee5d00ecadb6e574dcdd109a63e8943d5a25ba2e97094ad7d83dc28a6572da797d6b3e7fc6663bd93efb789fc17e489000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000002200000000000000000000000000000000000000000000000000000000000000180000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000087f3f5dee3fa664a57129108cd064823620f0510000000000000000000000000000000000000000000000007155d41081b0aaaa00000000000000000000000000000000000000000000000000000005008ecffa000000000000000000000000000000000000000000000000000000006631da521c3ebdf32bf66763f2e67af49eabcfd20dcdee2c1770c859c53cf2bfb8ae06520000000000000000000000000000000000000000000000000000000000000000f3b277728b3fee749481eb3e0b3b48980dbbab78658fc419025cb16eee34677500000000000000000000000000000000000000000000000000000000000000005a28e9363bb942b639270062aa6bb295f434bcdfc42c97267bf003f272060dc95a28e9363bb942b639270062aa6bb295f434bcdfc42c97267bf003f272060dc90000000000000000000000000000000000000000000000000000000000000280000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000024000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006cf1e9ca41f7611def408122793c358a3d11e5a500000000000000000000000000000000000000000000000000000018f309603c00000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000140000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000087f3f5dee3fa664a57129108cd064823620f0510000000000000000000000000000000000000000000000007155d41081b0aaaa00000000000000000000000000000000000000000000000000000005008ecffa000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000001c2000000000000000000000000000000000000000000000000000000000000000001c3ebdf32bf66763f2e67af49eabcfd20dcdee2c1770c859c53cf2bfb8ae065200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000060000000000000000000000007338afb07db145220849b04a45243956f20b14d90000000000000000000000000000000000000000000000004563918244f400000000000000000000000000000000000000000000000000099849302b37b3d18600000000000000000000000000000000000000000000000000000000ffffffffc5e8ee8dda38627f3c743f0e367881b1b98c033bf295e0636b414c241e541953000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000004563918244f400000000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000002840a50cf069e992aa4536211b23f286ef8875218740a50cf069e992aa4536211b23f286ef887521870000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000010600000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000040a50cf069e992aa4536211b23f286ef887521870000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000044c84c1c800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000760000000000000000000000000000000000000000000000000000000000000082000000000000000000000000000000000000000000000000000000000000008e00000000000000000000000000000000000000000000000000000000000000d800000000000000000000000000000000000000000000000000000000000000e40000000000000000000000000beb09000fa59627dc02bb55448ac1893eaa501a500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000060424a305f20000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000046000000000000000000000000000000000000000000000000000000000000004c0000000000000000000000000000000000000000000000000000000006631beaf0000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab410000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000001c0000000000000000000000000000000000000000000000000000000000000024000000000000000000000000000000000000000000000000000000000000002c000000000000000000000000000000000000000000000000000000000000003400000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab4100000000000000000000000000000000000000000000000000000000000003c0000000000000000000000000000000000000000000000000000000000000000100000000000000000000000051c72848c68a965f66fa7a88855f9f7784502a7f00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000031dd9e51133000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000007155d41081b0aaaa00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000005b4c01c2a000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000041b5d7c445a42cc3eca0327e249a8c1cedf3a51b98148edce88612e81bf32fdb9d625bd4eb062afba6a84cc474dc6da53f2691f326d393fbefc91ed4e4ed305f481b000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000093d199263632a4ef4bb438f1feb99e57b4b5f0bd00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000002438e9922e000000000000000000000000000000000000000000000000000012309ce54000000000000000000000000000000000000000000000000000000000000000000000000000000000003de27efa2f1aa663ae5d458857e731c129069f2900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000002438e9922e00000000000000000000000000000000000000000000000000038d7ea4c6800000000000000000000000000000000000000000000000000000000000000000000000000000000000ba12222222228d8ba445958a75a0704d566bf2c8000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000404945bcec90000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000003000000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab4100000000000000000000000000000000000000000000000000000000000000000000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab410000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000038000000000000000000000000000000000000000000000000000000f000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000010093d199263632a4ef4bb438f1feb99e57b4b5f0bd0000000000000000000005c200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000455bd74da8ba83bc00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000003de27efa2f1aa663ae5d458857e731c129069f2900020000000000000000058800000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000007f39c581f595b53c5cb19bd0b3f8da6c935e2ca00000000000000000000000007fc66500c84a76ad7e9c93437bfc5ac33e2ddae90000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000455bd74da8ba83bc0000000000000000000000000000000000000000000000000000000000000000fffffffffffffffffffffffffffffffffffffffffffffff63b000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000093d199263632a4ef4bb438f1feb99e57b4b5f0bd00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000002438e9922e00000000000000000000000000000000000000000000000000005af3107a4000000000000000000000000000000000000000000000000000000000000000000000000000000000003de27efa2f1aa663ae5d458857e731c129069f2900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000002438e9922e0000000000000000000000000000000000000000000000000011c37937e080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000868682"),
						To:        lo.ToPtr(common.HexToAddress("0x9008D19f58AAbD9eD0D60971565AA8510560ab41")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x420ddba0c2e81bd1413b53a9247d5aacd60e398e017c499eb3fd1826e222530e"),
						BlockNumber:       lo.Must(new(big.Int).SetString("19772757", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 6876131,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x1757136de"),
						GasUsed:           596424,

						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
								Topics: []common.Hash{
									common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"),
									common.HexToHash("0x00000000000000000000000040a50cf069e992aa4536211b23f286ef88752187"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000004563918244f40000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("19772757", 0)),
								TransactionHash: common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
								Index:           315,
								Removed:         false,
							}, {
								Address: common.HexToAddress("0x9008D19f58AAbD9eD0D60971565AA8510560ab41"),
								Topics: []common.Hash{
									common.HexToHash("0xed99827efb37016f2275f98c4bcf71c7551c75d59e9b450f79fa32e60be672c2"),
									common.HexToHash("0x00000000000000000000000040a50cf069e992aa4536211b23f286ef88752187"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000004c84c1c800000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("19772757", 0)),
								TransactionHash: common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
								Index:           316,
								Removed:         false,
							}, {
								Address: common.HexToAddress("0x9008D19f58AAbD9eD0D60971565AA8510560ab41"),
								Topics: []common.Hash{
									common.HexToHash("0xa07a543ab8a018198e99ca0184c93fe9050a79400a0a723441f84de1d972cc17"),
									common.HexToHash("0x000000000000000000000000087f3f5dee3fa664a57129108cd064823620f051"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000000000000000000000000000007155d41081b0aaaa00000000000000000000000000000000000000000000000000000005b46a1539000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000003815a14f307cc9338e01dcead7849de6dbac24eb0831aa4ee86026659c3defc73d087f3f5dee3fa664a57129108cd064823620f0516631da520000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("19772757", 0)),
								TransactionHash: common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
								Index:           317,
								Removed:         false,
							}, {
								Address: common.HexToAddress("0x9008D19f58AAbD9eD0D60971565AA8510560ab41"),
								Topics: []common.Hash{
									common.HexToHash("0xa07a543ab8a018198e99ca0184c93fe9050a79400a0a723441f84de1d972cc17"),
									common.HexToHash("0x00000000000000000000000040a50cf069e992aa4536211b23f286ef88752187"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000007fc66500c84a76ad7e9c93437bfc5ac33e2ddae90000000000000000000000000000000000000000000000004563918244f40000000000000000000000000000000000000000000000000009ca5dcb3527fcec67000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000380e46515d4d6713dc9b38e314fc6d8033826ebef22c764e053e193d3696e520f740a50cf069e992aa4536211b23f286ef88752187ffffffff0000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("19772757", 0)),
								TransactionHash: common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
								Index:           318,
								Removed:         false,
							}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x16179754d4467ec1476622f9d5507e74af8f695324a821f6ad2c5a73fae4595a"),
						TransactionIndex: 13,
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
				From:     "0x16C473448E770Ff647c69CBe19e28528877fba1B",
				To:       "0x9008D19f58AAbD9eD0D60971565AA8510560ab41",
				Type:     typex.ExchangeSwap,
				Platform: workerx.PlatformCow.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("3736787920653168")),
					Decimal: 18,
				},
				Calldata: &activityx.Calldata{
					FunctionHash: "0x13d79a0b",
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
				Status:    true,
				Timestamp: 1714536035,
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

			t.Log(string(lo.Must(json.MarshalIndent(activity, "", "\x20\x20"))))

			require.Equal(t, testcase.want, activity)
		})
	}
}
