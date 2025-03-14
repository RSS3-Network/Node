package base_test

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/v2/config"
	source "github.com/rss3-network/node/v2/internal/engine/protocol/ethereum"
	worker "github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/base"
	"github.com/rss3-network/node/v2/provider/ethereum"
	"github.com/rss3-network/node/v2/provider/ethereum/contract/base"
	"github.com/rss3-network/node/v2/provider/ethereum/endpoint"
	workerx "github.com/rss3-network/node/v2/schema/worker/decentralized"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestWorker_Base(t *testing.T) {
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
			name: "Deposit ETH from Ethereum Mainnet to Ethereum Base on Optimism Portal",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xfd22ef3295f2ff04af7d03aecd00251104b9c1c8e9a5ae8561ea4a758d54e876"),
						ParentHash:   common.HexToHash("0x8f85dd027a1aad0ffa917f620eff7a268e98446b50c6c9bc884ee563a1ab3073"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x690B9A9E9aa1C9dB991C7721a92d351Db4FaC990"),
						Number:       lo.Must(new(big.Int).SetString("17890185", 0)),
						GasLimit:     30000000,
						GasUsed:      11378037,
						Timestamp:    1691739287,
						BaseFee:      lo.Must(new(big.Int).SetString("21934542358", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xa01551ecd515a208bacb2f6cbfb73a993e8d49e5ad339140ca0a19c89b9fd534"),
						From:      common.HexToAddress("0x936AaDE0d4d835271C9C11B89880244330864a63"),
						Gas:       200000,
						GasPrice:  lo.Must(new(big.Int).SetString("22034542358", 10)),
						Hash:      common.HexToHash("0xa01551ecd515a208bacb2f6cbfb73a993e8d49e5ad339140ca0a19c89b9fd534"),
						Input:     hexutil.MustDecode("0xe9e05c42000000000000000000000000936aade0d4d835271c9c11b89880244330864a630000000000000000000000000000000000000000000000000de0b6b3a764000000000000000000000000000000000000000000000000000000000000000186a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000010100000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x49048044D57e1C92A77f79988d21Fa8fAF74E97e")),
						Value:     lo.Must(new(big.Int).SetString("1000000000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xfd22ef3295f2ff04af7d03aecd00251104b9c1c8e9a5ae8561ea4a758d54e876"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17890185", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 8473125,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x5215c6f16"),
						GasUsed:           52594,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x49048044D57e1C92A77f79988d21Fa8fAF74E97e"),
							Topics: []common.Hash{
								common.HexToHash("0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32"),
								common.HexToHash("0x000000000000000000000000936aade0d4d835271c9c11b89880244330864a63"),
								common.HexToHash("0x000000000000000000000000936aade0d4d835271c9c11b89880244330864a63"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000004a0000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000de0b6b3a764000000000000000186a0000100000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17890185", 0)),
							TransactionHash: common.HexToHash("0xa01551ecd515a208bacb2f6cbfb73a993e8d49e5ad339140ca0a19c89b9fd534"),
							Index:           156,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xa01551ecd515a208bacb2f6cbfb73a993e8d49e5ad339140ca0a19c89b9fd534"),
						TransactionIndex: 81,
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
				ID:      "0xa01551ecd515a208bacb2f6cbfb73a993e8d49e5ad339140ca0a19c89b9fd534",
				Network: network.Ethereum,
				Index:   81,
				From:    "0x936AaDE0d4d835271C9C11B89880244330864a63",
				To:      base.AddressOptimismPortal.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xe9e05c42",
				},
				Tag:      tag.Transaction,
				Platform: workerx.PlatformBase.String(),

				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("1158884720776652")),
					Decimal: 18,
				},
				TotalActions: 1,
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Tag:      tag.Transaction,
						Platform: workerx.PlatformBase.String(),
						From:     "0x936AaDE0d4d835271C9C11B89880244330864a63",
						To:       "0x936AaDE0d4d835271C9C11B89880244330864a63",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeDeposit,
							SourceNetwork: network.Ethereum,
							TargetNetwork: network.Base,
							Token: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1000000000000000000"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1691739287,
			},
			wantError: require.NoError,
		},
		{
			name: "Deposit ETH from Ethereum Mainnet to Ethereum Base on L1 Standard Bridge",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x18ce6cdef97122caa0320de32dfd7bbc2187574ad8540454f3ee5e45bd468226"),
						ParentHash:   common.HexToHash("0x8efc890a8f9be6d5eba2df2622540cc317ad3d6ff651d27a62c1998f4443d3c9"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xDAFEA492D9c6733ae3d56b7Ed1ADB60692c98Bc5"),
						Number:       lo.Must(new(big.Int).SetString("17890108", 0)),
						GasLimit:     30000000,
						GasUsed:      14895124,
						Timestamp:    1691738363,
						BaseFee:      lo.Must(new(big.Int).SetString("15576551442", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x185d83cec71d01c7ddf3f409ea57ae76c53f835b51f61d609dfaf603644ead02"),
						From:      common.HexToAddress("0xFBFd6Fa9F73Ac6A058E01259034C28001BEf8247"),
						Gas:       199501,
						GasPrice:  lo.Must(new(big.Int).SetString("15676551442", 10)),
						Hash:      common.HexToHash("0x185d83cec71d01c7ddf3f409ea57ae76c53f835b51f61d609dfaf603644ead02"),
						Input:     hexutil.MustDecode("0xb1a1a8820000000000000000000000000000000000000000000000000000000000030d4000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x3154Cf16ccdb4C6d922629664174b904d80F2C35")),
						Value:     lo.Must(new(big.Int).SetString("400000000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x18ce6cdef97122caa0320de32dfd7bbc2187574ad8540454f3ee5e45bd468226"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17890108", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 10059062,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3a6653112"),
						GasUsed:           127028,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x3154Cf16ccdb4C6d922629664174b904d80F2C35"),
							Topics: []common.Hash{
								common.HexToHash("0x35d79ab81f2b2017e19afb5c5571778877782d7a8786f5907f93b0f4702f4f23"),
								common.HexToHash("0x000000000000000000000000fbfd6fa9f73ac6a058e01259034c28001bef8247"),
								common.HexToHash("0x000000000000000000000000fbfd6fa9f73ac6a058e01259034c28001bef8247"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000058d15e17628000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17890108", 0)),
							TransactionHash: common.HexToHash("0x185d83cec71d01c7ddf3f409ea57ae76c53f835b51f61d609dfaf603644ead02"),
							Index:           222,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x3154Cf16ccdb4C6d922629664174b904d80F2C35"),
							Topics: []common.Hash{
								common.HexToHash("0x2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5"),
								common.HexToHash("0x000000000000000000000000fbfd6fa9f73ac6a058e01259034c28001bef8247"),
								common.HexToHash("0x000000000000000000000000fbfd6fa9f73ac6a058e01259034c28001bef8247"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000058d15e17628000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17890108", 0)),
							TransactionHash: common.HexToHash("0x185d83cec71d01c7ddf3f409ea57ae76c53f835b51f61d609dfaf603644ead02"),
							Index:           223,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x49048044D57e1C92A77f79988d21Fa8fAF74E97e"),
							Topics: []common.Hash{
								common.HexToHash("0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32"),
								common.HexToHash("0x000000000000000000000000977f82a600a1414e583f7f13623f1ac5d58b1c0b"),
								common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000007"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000001ed000000000000000000000000000000000000000000000000058d15e176280000000000000000000000000000000000000000000000000000058d15e1762800000000000000077d2e00d764ad0b0001000000000000000000000000000000000000000000000000000000004cbb0000000000000000000000003154cf16ccdb4c6d922629664174b904d80f2c350000000000000000000000004200000000000000000000000000000000000010000000000000000000000000000000000000000000000000058d15e1762800000000000000000000000000000000000000000000000000000000000000030d4000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000a41635f5fd000000000000000000000000fbfd6fa9f73ac6a058e01259034c28001bef8247000000000000000000000000fbfd6fa9f73ac6a058e01259034c28001bef8247000000000000000000000000000000000000000000000000058d15e176280000000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17890108", 0)),
							TransactionHash: common.HexToHash("0x185d83cec71d01c7ddf3f409ea57ae76c53f835b51f61d609dfaf603644ead02"),
							Index:           224,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x866E82a600A1414e583f7F13623F1aC5d58b0Afa"),
							Topics: []common.Hash{
								common.HexToHash("0xcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a"),
								common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000010"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000003154cf16ccdb4c6d922629664174b904d80f2c3500000000000000000000000000000000000000000000000000000000000000800001000000000000000000000000000000000000000000000000000000004cbb0000000000000000000000000000000000000000000000000000000000030d4000000000000000000000000000000000000000000000000000000000000000a41635f5fd000000000000000000000000fbfd6fa9f73ac6a058e01259034c28001bef8247000000000000000000000000fbfd6fa9f73ac6a058e01259034c28001bef8247000000000000000000000000000000000000000000000000058d15e1762800000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17890108", 0)),
							TransactionHash: common.HexToHash("0x185d83cec71d01c7ddf3f409ea57ae76c53f835b51f61d609dfaf603644ead02"),
							Index:           225,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x866E82a600A1414e583f7F13623F1aC5d58b0Afa"),
							Topics: []common.Hash{
								common.HexToHash("0x8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d546"),
								common.HexToHash("0x0000000000000000000000003154cf16ccdb4c6d922629664174b904d80f2c35"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000058d15e176280000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17890108", 0)),
							TransactionHash: common.HexToHash("0x185d83cec71d01c7ddf3f409ea57ae76c53f835b51f61d609dfaf603644ead02"),
							Index:           226,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x185d83cec71d01c7ddf3f409ea57ae76c53f835b51f61d609dfaf603644ead02"),
						TransactionIndex: 92,
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
				ID:      "0x185d83cec71d01c7ddf3f409ea57ae76c53f835b51f61d609dfaf603644ead02",
				Network: network.Ethereum,
				Index:   92,
				From:    "0xFBFd6Fa9F73Ac6A058E01259034C28001BEf8247",
				To:      base.AddressL1StandardBridge.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xb1a1a882",
				},
				Tag:      tag.Transaction,
				Platform: workerx.PlatformBase.String(),

				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("1991360976574376")),
					Decimal: 18,
				},
				TotalActions: 1,
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Tag:      tag.Transaction,
						Platform: workerx.PlatformBase.String(),
						From:     "0xFBFd6Fa9F73Ac6A058E01259034C28001BEf8247",
						To:       "0xFBFd6Fa9F73Ac6A058E01259034C28001BEf8247",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeDeposit,
							SourceNetwork: network.Ethereum,
							TargetNetwork: network.Base,
							Token: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("400000000000000000"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1691738363,
			},
			wantError: require.NoError,
		},
		{
			name: "Deposit ERC20 from Ethereum Mainnet to Ethereum Base on L1 Standard Bridge",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x77d68f0268295c5ed7ceeccf899998973de591aa26c753633e18985f1023ff11"),
						ParentHash:   common.HexToHash("0x7930e89737c12a4604f88b15535b97ee8d09821c4d55b6598ebd604b3ff5101e"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xDAFEA492D9c6733ae3d56b7Ed1ADB60692c98Bc5"),
						Number:       lo.Must(new(big.Int).SetString("17890089", 0)),
						GasLimit:     30000000,
						GasUsed:      14072715,
						Timestamp:    1691738135,
						BaseFee:      lo.Must(new(big.Int).SetString("15302876294", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xd5eabb9725773af1e243f70936a3d00834cf015e262f5a60d381486e9c18244d"),
						From:      common.HexToAddress("0x00005fF7b4E62B6b868D36809F8a30c85A6D0000"),
						Gas:       355374,
						GasPrice:  lo.Must(new(big.Int).SetString("15500000000", 10)),
						Hash:      common.HexToHash("0xd5eabb9725773af1e243f70936a3d00834cf015e262f5a60d381486e9c18244d"),
						Input:     hexutil.MustDecode("0x58a997f6000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000d9aaec86b65d86f6a7b5b1b0c42ffa531710b6ca0000000000000000000000000000000000000000000000000000000df82b864400000000000000000000000000000000000000000000000000000000000186a000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000010100000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x3154Cf16ccdb4C6d922629664174b904d80F2C35")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x77d68f0268295c5ed7ceeccf899998973de591aa26c753633e18985f1023ff11"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17890089", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 6771576,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x39bdf3b00"),
						GasUsed:           165923,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000000005ff7b4e62b6b868d36809f8a30c85a6d0000"),
								common.HexToHash("0x0000000000000000000000003154cf16ccdb4c6d922629664174b904d80f2c35"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000df82b8644"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17890089", 0)),
							TransactionHash: common.HexToHash("0xd5eabb9725773af1e243f70936a3d00834cf015e262f5a60d381486e9c18244d"),
							Index:           167,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x3154Cf16ccdb4C6d922629664174b904d80F2C35"),
							Topics: []common.Hash{
								common.HexToHash("0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396"),
								common.HexToHash("0x000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
								common.HexToHash("0x000000000000000000000000d9aaec86b65d86f6a7b5b1b0c42ffa531710b6ca"),
								common.HexToHash("0x00000000000000000000000000005ff7b4e62b6b868d36809f8a30c85a6d0000"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000005ff7b4e62b6b868d36809f8a30c85a6d00000000000000000000000000000000000000000000000000000000000df82b8644000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000010100000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17890089", 0)),
							TransactionHash: common.HexToHash("0xd5eabb9725773af1e243f70936a3d00834cf015e262f5a60d381486e9c18244d"),
							Index:           168,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x3154Cf16ccdb4C6d922629664174b904d80F2C35"),
							Topics: []common.Hash{
								common.HexToHash("0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf"),
								common.HexToHash("0x000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
								common.HexToHash("0x000000000000000000000000d9aaec86b65d86f6a7b5b1b0c42ffa531710b6ca"),
								common.HexToHash("0x00000000000000000000000000005ff7b4e62b6b868d36809f8a30c85a6d0000"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000005ff7b4e62b6b868d36809f8a30c85a6d00000000000000000000000000000000000000000000000000000000000df82b8644000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000010100000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17890089", 0)),
							TransactionHash: common.HexToHash("0xd5eabb9725773af1e243f70936a3d00834cf015e262f5a60d381486e9c18244d"),
							Index:           169,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x49048044D57e1C92A77f79988d21Fa8fAF74E97e"),
							Topics: []common.Hash{
								common.HexToHash("0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32"),
								common.HexToHash("0x000000000000000000000000977f82a600a1414e583f7f13623f1ac5d58b1c0b"),
								common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000007"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000024d00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005f65b00d764ad0b0001000000000000000000000000000000000000000000000000000000004c9d0000000000000000000000003154cf16ccdb4c6d922629664174b904d80f2c350000000000000000000000004200000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000186a000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001040166a07a000000000000000000000000d9aaec86b65d86f6a7b5b1b0c42ffa531710b6ca000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb4800000000000000000000000000005ff7b4e62b6b868d36809f8a30c85a6d000000000000000000000000000000005ff7b4e62b6b868d36809f8a30c85a6d00000000000000000000000000000000000000000000000000000000000df82b864400000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000101000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17890089", 0)),
							TransactionHash: common.HexToHash("0xd5eabb9725773af1e243f70936a3d00834cf015e262f5a60d381486e9c18244d"),
							Index:           170,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x866E82a600A1414e583f7F13623F1aC5d58b0Afa"),
							Topics: []common.Hash{
								common.HexToHash("0xcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a"),
								common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000010"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000003154cf16ccdb4c6d922629664174b904d80f2c3500000000000000000000000000000000000000000000000000000000000000800001000000000000000000000000000000000000000000000000000000004c9d00000000000000000000000000000000000000000000000000000000000186a000000000000000000000000000000000000000000000000000000000000001040166a07a000000000000000000000000d9aaec86b65d86f6a7b5b1b0c42ffa531710b6ca000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb4800000000000000000000000000005ff7b4e62b6b868d36809f8a30c85a6d000000000000000000000000000000005ff7b4e62b6b868d36809f8a30c85a6d00000000000000000000000000000000000000000000000000000000000df82b864400000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000001010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17890089", 0)),
							TransactionHash: common.HexToHash("0xd5eabb9725773af1e243f70936a3d00834cf015e262f5a60d381486e9c18244d"),
							Index:           171,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x866E82a600A1414e583f7F13623F1aC5d58b0Afa"),
							Topics: []common.Hash{
								common.HexToHash("0x8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d546"),
								common.HexToHash("0x0000000000000000000000003154cf16ccdb4c6d922629664174b904d80f2c35"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17890089", 0)),
							TransactionHash: common.HexToHash("0xd5eabb9725773af1e243f70936a3d00834cf015e262f5a60d381486e9c18244d"),
							Index:           172,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xd5eabb9725773af1e243f70936a3d00834cf015e262f5a60d381486e9c18244d"),
						TransactionIndex: 88,
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
				ID:      "0xd5eabb9725773af1e243f70936a3d00834cf015e262f5a60d381486e9c18244d",
				Network: network.Ethereum,
				Index:   88,
				From:    "0x00005fF7b4E62B6b868D36809F8a30c85A6D0000",
				To:      base.AddressL1StandardBridge.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x58a997f6",
				},
				Tag:      tag.Transaction,
				Platform: workerx.PlatformBase.String(),

				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("2571806500000000")),
					Decimal: 18,
				},
				TotalActions: 1,
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Tag:      tag.Transaction,
						Platform: workerx.PlatformBase.String(),
						From:     "0x00005fF7b4E62B6b868D36809F8a30c85A6D0000",
						To:       "0x00005fF7b4E62B6b868D36809F8a30c85A6D0000",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeDeposit,
							SourceNetwork: network.Ethereum,
							TargetNetwork: network.Base,
							Token: metadata.Token{
								Address:  lo.ToPtr("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("59998176836"))),
								Name:     "USD Coin",
								Symbol:   "USDC",
								Decimals: 6,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1691738135,
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

			data, err := json.MarshalIndent(activity, "", "\x20\x20")
			require.NoError(t, err)

			t.Log(string(data))
			require.Equal(t, testcase.want, activity)
		})
	}
}
