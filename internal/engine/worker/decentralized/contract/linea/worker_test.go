package linea_test

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/decentralized/contract/linea"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/linea"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	workerx "github.com/rss3-network/node/schema/worker/decentralized"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestWorker_Linea(t *testing.T) {
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
			name: "Deposit ETH from Ethereum Mainnet to Ethereum Linea on Ethereum Mainnet",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xcd51f362f988b3f8dc45f7e579408cda2ef2398faed35432d50409337fa6e48a"),
						ParentHash:   common.HexToHash("0x8bf3c454d359f0050cb3a0cc472d47e5339182b7d9e35c111bc09449b04addad"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"),
						Number:       lo.Must(new(big.Int).SetString("18246708", 0)),
						GasLimit:     30000000,
						GasUsed:      26645446,
						Timestamp:    1696054991,
						BaseFee:      lo.Must(new(big.Int).SetString("7560221523", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x5a63ba05277c089ca86439c8e57c17a17fe3578eb26268e2f19bd94482e236b9"),
						From:      common.HexToAddress("0xfF735406f897a32E3e815fb3C53FB8BA16214d72"),
						Gas:       99301,
						GasPrice:  lo.Must(new(big.Int).SetString("7660221523", 10)),
						Hash:      common.HexToHash("0x5a63ba05277c089ca86439c8e57c17a17fe3578eb26268e2f19bd94482e236b9"),
						Input:     hexutil.MustDecode("0x9f3ce55a000000000000000000000000ff735406f897a32e3e815fb3c53fb8ba16214d72000000000000000000000000000000000000000000000000000063ab79d1b54000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xd19d4B5d358258f05D7B411E21A1460D11B0876F")),
						Value:     lo.Must(new(big.Int).SetString("1617283463770166030", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xcd51f362f988b3f8dc45f7e579408cda2ef2398faed35432d50409337fa6e48a"),
						BlockNumber:       lo.Must(new(big.Int).SetString("18246708", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 5970245,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x1c895b453"),
						GasUsed:           65659,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xd19d4B5d358258f05D7B411E21A1460D11B0876F"),
							Topics: []common.Hash{
								common.HexToHash("0xe856c2b8bd4eb0027ce32eeaf595c21b0b6b4644b326e5b7bd80a1cf8db72e6c"),
								common.HexToHash("0x000000000000000000000000ff735406f897a32e3e815fb3c53fb8ba16214d72"),
								common.HexToHash("0x000000000000000000000000ff735406f897a32e3e815fb3c53fb8ba16214d72"),
								common.HexToHash("0x48a4bcb472d72cc281e53e5cb150cb004b50fa7b6380c101debecef14641312a"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000063ab79d1b54000000000000000000000000000000000000000000000000016715b125d509dce000000000000000000000000000000000000000000000000000000000004f1df00000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("18246708", 0)),
							TransactionHash: common.HexToHash("0x5a63ba05277c089ca86439c8e57c17a17fe3578eb26268e2f19bd94482e236b9"),
							Index:           118,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x5a63ba05277c089ca86439c8e57c17a17fe3578eb26268e2f19bd94482e236b9"),
						TransactionIndex: 53,
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
				ID:      "0x5a63ba05277c089ca86439c8e57c17a17fe3578eb26268e2f19bd94482e236b9",
				Network: network.Ethereum,
				Index:   53,
				From:    "0xfF735406f897a32E3e815fb3C53FB8BA16214d72",
				To:      linea.AddressZKEVMV2.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x9f3ce55a",
				},
				Tag:      tag.Transaction,
				Platform: workerx.PlatformLinea.String(),

				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("502962484978657")),
					Decimal: 18,
				},
				TotalActions: 1,
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Tag:      tag.Transaction,
						Platform: workerx.PlatformLinea.String(),
						From:     "0xfF735406f897a32E3e815fb3C53FB8BA16214d72",
						To:       "0xfF735406f897a32E3e815fb3C53FB8BA16214d72",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeDeposit,
							SourceNetwork: network.Ethereum,
							TargetNetwork: network.Linea,
							Token: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1617173875635822030"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1696054991,
			},
			wantError: require.NoError,
		},
		{
			name: "Withdraw ETH from Ethereum Linea to Ethereum Mainnet on Ethereum Mainnet",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x2bcf73dc999ec5dd6d43130d706ee4cf63354f68951e7b3efd11e51b9d534dc1"),
						ParentHash:   common.HexToHash("0x63c60ddd36871e5cd6caa9b610443a2c7eb2dab8b54905bfd7348f5b57f10f4a"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("18246566", 0)),
						GasLimit:     30000000,
						GasUsed:      18130698,
						Timestamp:    1696053275,
						BaseFee:      lo.Must(new(big.Int).SetString("6851097859", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xef8212b2b830d5be9a1685253a6ee6c818ac12bd433dbdb27dc07942dd0b01a7"),
						From:      common.HexToAddress("0x877180d5a580696daddDd8b7410889552E3C9295"),
						Gas:       107685,
						GasPrice:  lo.Must(new(big.Int).SetString("6951097859", 10)),
						Hash:      common.HexToHash("0xef8212b2b830d5be9a1685253a6ee6c818ac12bd433dbdb27dc07942dd0b01a7"),
						Input:     hexutil.MustDecode("0x491e0936000000000000000000000000877180d5a580696dadddd8b7410889552e3c9295000000000000000000000000877180d5a580696dadddd8b7410889552e3c929500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001bc16d674ec80000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000066470000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xd19d4B5d358258f05D7B411E21A1460D11B0876F")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x2bcf73dc999ec5dd6d43130d706ee4cf63354f68951e7b3efd11e51b9d534dc1"),
						BlockNumber:       lo.Must(new(big.Int).SetString("18246566", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 12521471,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x19e515603"),
						GasUsed:           58606,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xd19d4B5d358258f05D7B411E21A1460D11B0876F"),
							Topics: []common.Hash{
								common.HexToHash("0xa4c827e719e911e8f19393ccdb85b5102f08f0910604d340ba38390b7ff2ab0e"),
								common.HexToHash("0x3a393cec3d38c33e8638cf1235fd1e7b73e8c711cef6edfa341e1dcd65b204c6"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("18246566", 0)),
							TransactionHash: common.HexToHash("0xef8212b2b830d5be9a1685253a6ee6c818ac12bd433dbdb27dc07942dd0b01a7"),
							Index:           204,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xef8212b2b830d5be9a1685253a6ee6c818ac12bd433dbdb27dc07942dd0b01a7"),
						TransactionIndex: 80,
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
				ID:      "0xef8212b2b830d5be9a1685253a6ee6c818ac12bd433dbdb27dc07942dd0b01a7",
				Network: network.Ethereum,
				Index:   80,
				From:    "0x877180d5a580696daddDd8b7410889552E3C9295",
				To:      linea.AddressZKEVMV2.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x491e0936",
				},
				Tag:      tag.Transaction,
				Platform: workerx.PlatformLinea.String(),

				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("407376041124554")),
					Decimal: 18,
				},
				TotalActions: 1,
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Tag:      tag.Transaction,
						Platform: workerx.PlatformLinea.String(),
						From:     "0x877180d5a580696daddDd8b7410889552E3C9295",
						To:       "0x877180d5a580696daddDd8b7410889552E3C9295",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeWithdraw,
							SourceNetwork: network.Linea,
							TargetNetwork: network.Ethereum,
							Token: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2000000000000000000"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1696053275,
			},
			wantError: require.NoError,
		},
		{
			name: "Withdraw Uni from Ethereum Linea to Ethereum Mainnet on Ethereum Mainnet",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xac0210f0d3527734c81979918fe52cb0b7969ef0a28f62542dcf471e6b3c6a5d"),
						ParentHash:   common.HexToHash("0x3e2b9503f0c35660edbb15a97cc709713f5fa042d6196fa6a680f45da682219d"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"),
						Number:       lo.Must(new(big.Int).SetString("20667438", 0)),
						GasLimit:     30000000,
						GasUsed:      13861506,
						Timestamp:    1725334991,
						BaseFee:      lo.Must(new(big.Int).SetString("779389969", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x1f2c1706277d879e8f9c95cdd4b0cd5694e0e22511ea7fce90081e152a7628f7"),
						From:      common.HexToAddress("0x9bacb15D8935D7fc93E119Cc99C711Dc59aea660"),
						Gas:       168324,
						GasPrice:  lo.Must(new(big.Int).SetString("1279389969", 10)),
						Hash:      common.HexToHash("0x1f2c1706277d879e8f9c95cdd4b0cd5694e0e22511ea7fce90081e152a7628f7"),
						Input:     hexutil.MustDecode("0x6463fb2a0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000129630000000000000000000000000000000000000000000000000000000000000002000000000000000000000000353012dc4a9a6cf55c941badc267f82004a8ceb9000000000000000000000000051f1d88f0af5763fb888ec4378b4d8b29ea331900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000069c4f790c6130f7951d6acfd5ef61c50a48aaa260498eeed38fd2606b3857b0f000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000057143fc0f1736c515aa2d0f8e2c7e1eaf5ec20c1986e98ca8123814b721d9eb1a8b84e58356b6b93f58fc182ba1ee3c197cb182369f89abc73f7b9036b49322759f5761d76dd07f1128ffab5c6d3594b3e5057bf6dcd29002e7dc162f740fb33a21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba85e58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a1934400000000000000000000000000000000000000000000000000000000000000c4e4d274510000000000000000000000001f9840a85d5af5bf1d1762f925bdaddc4201f984000000000000000000000000000000000000000000000001fc519068fc45a0000000000000000000000000009bacb15d8935d7fc93e119cc99c711dc59aea660000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xd19d4B5d358258f05D7B411E21A1460D11B0876F")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xac0210f0d3527734c81979918fe52cb0b7969ef0a28f62542dcf471e6b3c6a5d"),
						BlockNumber:       lo.Must(new(big.Int).SetString("20667438", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 12181911,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x4c41f111"),
						GasUsed:           106209,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000051f1d88f0af5763fb888ec4378b4d8b29ea3319"),
								common.HexToHash("0x0000000000000000000000009bacb15d8935d7fc93e119cc99c711dc59aea660"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000001fc519068fc45a000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20667438", 0)),
							TransactionHash: common.HexToHash("0x1f2c1706277d879e8f9c95cdd4b0cd5694e0e22511ea7fce90081e152a7628f7"),
							Index:           307,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x051F1D88f0aF5763fB888eC4378b4D8B29ea3319"),
							Topics: []common.Hash{
								common.HexToHash("0x6ed06519caca659cdefa71015c79a561928d3cf8cc4a3e9739fde9fb5fb38d64"),
								common.HexToHash("0x0000000000000000000000001f9840a85d5af5bf1d1762f925bdaddc4201f984"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x0000000000000000000000009bacb15d8935d7fc93e119cc99c711dc59aea660"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000001fc519068fc45a000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20667438", 0)),
							TransactionHash: common.HexToHash("0x1f2c1706277d879e8f9c95cdd4b0cd5694e0e22511ea7fce90081e152a7628f7"),
							Index:           308,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xd19d4B5d358258f05D7B411E21A1460D11B0876F"),
							Topics: []common.Hash{
								common.HexToHash("0xa4c827e719e911e8f19393ccdb85b5102f08f0910604d340ba38390b7ff2ab0e"),
								common.HexToHash("0xa0bf9527e4f1fbae8339d5f0a9d3f95062b1ac315937d153bd3d7893dda90d3b"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("20667438", 0)),
							TransactionHash: common.HexToHash("0x1f2c1706277d879e8f9c95cdd4b0cd5694e0e22511ea7fce90081e152a7628f7"),
							Index:           309,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x1f2c1706277d879e8f9c95cdd4b0cd5694e0e22511ea7fce90081e152a7628f7"),
						TransactionIndex: 124,
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
				ID:      "0x1f2c1706277d879e8f9c95cdd4b0cd5694e0e22511ea7fce90081e152a7628f7",
				Network: network.Ethereum,
				Index:   124,
				From:    "0x9bacb15D8935D7fc93E119Cc99C711Dc59aea660",
				To:      linea.AddressZKEVMV2.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x6463fb2a",
				},
				Tag:      tag.Transaction,
				Platform: workerx.PlatformLinea.String(),

				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("135882729217521")),
					Decimal: 18,
				},
				TotalActions: 1,
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Tag:      tag.Transaction,
						Platform: workerx.PlatformLinea.String(),
						From:     "0x9bacb15D8935D7fc93E119Cc99C711Dc59aea660",
						To:       "0x9bacb15D8935D7fc93E119Cc99C711Dc59aea660",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeWithdraw,
							SourceNetwork: network.Ethereum,
							TargetNetwork: network.Linea,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("36628216024964374528"))),
								Name:     "Uniswap",
								Symbol:   "UNI",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1725334991,
			},
			wantError: require.NoError,
		},
		{
			name: "Deposit Wrapped BTC from Ethereum Mainnet to Ethereum Linea on Ethereum Mainnet",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x593427df897c60187efa338d18c62c8134b1eaa06d7554e3fb77d25b2f293a8f"),
						ParentHash:   common.HexToHash("0x4678738a59ff27bd0d03c7f17ba17b5da21eea61ccdb0f8cb7035b2cd7c9acc8"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("20667178", 0)),
						GasLimit:     30000000,
						GasUsed:      9080611,
						Timestamp:    1725331847,
						BaseFee:      lo.Must(new(big.Int).SetString("870634089", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x36aa2360ca9a4eb79c742d1d8bee631336facbd876f0b9b1ecc10203e4c8ea7a"),
						From:      common.HexToAddress("0xA2D37a9DdB9bC09eeE1A6E0FE66ac6c1662D0954"),
						Gas:       200733,
						GasPrice:  lo.Must(new(big.Int).SetString("1370634089", 10)),
						Hash:      common.HexToHash("0x36aa2360ca9a4eb79c742d1d8bee631336facbd876f0b9b1ecc10203e4c8ea7a"),
						Input:     hexutil.MustDecode("0x522ea81a0000000000000000000000002260fac5e5542a773aa44fbcfedf7c193bc2c59900000000000000000000000000000000000000000000000000000000000d4670000000000000000000000000a2d37a9ddb9bc09eee1a6e0fe66ac6c1662d0954"),
						To:        lo.ToPtr(common.HexToAddress("0x051F1D88f0aF5763fB888eC4378b4D8B29ea3319")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x593427df897c60187efa338d18c62c8134b1eaa06d7554e3fb77d25b2f293a8f"),
						BlockNumber:       lo.Must(new(big.Int).SetString("20667178", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 6202413,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x51b23769"),
						GasUsed:           117032,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000a2d37a9ddb9bc09eee1a6e0fe66ac6c1662d0954"),
								common.HexToHash("0x000000000000000000000000051f1d88f0af5763fb888ec4378b4d8b29ea3319"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000d4670"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20667178", 0)),
							TransactionHash: common.HexToHash("0x36aa2360ca9a4eb79c742d1d8bee631336facbd876f0b9b1ecc10203e4c8ea7a"),
							Index:           164,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xd19d4B5d358258f05D7B411E21A1460D11B0876F"),
							Topics: []common.Hash{
								common.HexToHash("0xea3b023b4c8680d4b4824f0143132c95476359a2bb70a81d6c5a36f6918f6339"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000ab0db"),
								common.HexToHash("0x8578562997ba8320b241e87289f52e4957e9f5133273412ea3f6a45a1d5b7a1b"),
								common.HexToHash("0x34139f2f3e2a091b6477508b0abaa25eaa4bd386da7d223b452ff2a433b9a8da"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("20667178", 0)),
							TransactionHash: common.HexToHash("0x36aa2360ca9a4eb79c742d1d8bee631336facbd876f0b9b1ecc10203e4c8ea7a"),
							Index:           165,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xd19d4B5d358258f05D7B411E21A1460D11B0876F"),
							Topics: []common.Hash{
								common.HexToHash("0xe856c2b8bd4eb0027ce32eeaf595c21b0b6b4644b326e5b7bd80a1cf8db72e6c"),
								common.HexToHash("0x000000000000000000000000051f1d88f0af5763fb888ec4378b4d8b29ea3319"),
								common.HexToHash("0x000000000000000000000000353012dc4a9a6cf55c941badc267f82004a8ceb9"),
								common.HexToHash("0x34139f2f3e2a091b6477508b0abaa25eaa4bd386da7d223b452ff2a433b9a8da"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ab0db000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c4e4d274510000000000000000000000002260fac5e5542a773aa44fbcfedf7c193bc2c59900000000000000000000000000000000000000000000000000000000000d4670000000000000000000000000a2d37a9ddb9bc09eee1a6e0fe66ac6c1662d0954000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20667178", 0)),
							TransactionHash: common.HexToHash("0x36aa2360ca9a4eb79c742d1d8bee631336facbd876f0b9b1ecc10203e4c8ea7a"),
							Index:           166,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x051F1D88f0aF5763fB888eC4378b4D8B29ea3319"),
							Topics: []common.Hash{
								common.HexToHash("0x8780a94875b70464f8ac6c28851501d32e7fd4ee574e4b94beb28923a3c42d9c"),
								common.HexToHash("0x000000000000000000000000a2d37a9ddb9bc09eee1a6e0fe66ac6c1662d0954"),
								common.HexToHash("0x000000000000000000000000a2d37a9ddb9bc09eee1a6e0fe66ac6c1662d0954"),
								common.HexToHash("0x0000000000000000000000002260fac5e5542a773aa44fbcfedf7c193bc2c599"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000d4670"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20667178", 0)),
							TransactionHash: common.HexToHash("0x36aa2360ca9a4eb79c742d1d8bee631336facbd876f0b9b1ecc10203e4c8ea7a"),
							Index:           167,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x36aa2360ca9a4eb79c742d1d8bee631336facbd876f0b9b1ecc10203e4c8ea7a"),
						TransactionIndex: 75,
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
				ID:      "0x36aa2360ca9a4eb79c742d1d8bee631336facbd876f0b9b1ecc10203e4c8ea7a",
				Network: network.Ethereum,
				Index:   75,
				From:    "0xA2D37a9DdB9bC09eeE1A6E0FE66ac6c1662D0954",
				To:      linea.AddressTokenBridge.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x522ea81a",
				},
				Tag:      tag.Transaction,
				Platform: workerx.PlatformLinea.String(),

				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("160408048703848")),
					Decimal: 18,
				},
				TotalActions: 1,
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Tag:      tag.Transaction,
						Platform: workerx.PlatformLinea.String(),
						From:     "0xA2D37a9DdB9bC09eeE1A6E0FE66ac6c1662D0954",
						To:       "0xA2D37a9DdB9bC09eeE1A6E0FE66ac6c1662D0954",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeDeposit,
							SourceNetwork: network.Ethereum,
							TargetNetwork: network.Linea,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("870000"))),
								Name:     "Wrapped BTC",
								Symbol:   "WBTC",
								Decimals: 8,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1725331847,
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
