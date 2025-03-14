package rainbow_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/v2/config"
	source "github.com/rss3-network/node/v2/internal/engine/protocol/ethereum"
	worker "github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/rainbow"
	"github.com/rss3-network/node/v2/provider/ethereum"
	"github.com/rss3-network/node/v2/provider/ethereum/contract/rainbow"
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

func TestWorker_Rainbow(t *testing.T) {
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
			name: "Swap ETH to GIGA on Ethereum",
			arguments: struct {
				task   *source.Task
				config *config.Module
			}{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x49c8f0217c18bb6c1016ab256b214cc0b7cc49188a3d62a1265981598f57547e"),
						ParentHash:   common.HexToHash("0x2d0397a462ce1b249120c4b03495c83312c0f32be9463715f9f54cb147a5b3e2"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("20951604", 0)),
						GasLimit:     30000000,
						GasUsed:      11763606,
						Timestamp:    1728761231,
						BaseFee:      lo.Must(new(big.Int).SetString("8288888059", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x53ea61bbe1656b89995690376c7c8b525bc046307546814578b7df7ec141557d"),
						From:      common.HexToAddress("0x9E7E6ee8607628362B25AdF109d5704Db3f4640F"),
						Gas:       306544,
						GasPrice:  lo.Must(new(big.Int).SetString("8848505557", 10)),
						Hash:      common.HexToHash("0x53ea61bbe1656b89995690376c7c8b525bc046307546814578b7df7ec141557d"),
						Input:     hexutil.MustDecode("0x3c2b9a7d00000000000000000000000080ed012ef8dcd8e3a347d9034487aebd991bfda5000000000000000000000000111111125421ca6dc452d289314280a0f8842a65000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000013cf57ab448000000000000000000000000000000000000000000000000000000000000000048a76dfc3b000000000000000000000000000000000000000000000000035e0c7d905655d300000000000000003b6d03406c59c360ef92ae28dcda2b60cf0ffc1c08fd3a8fd6f29312000000000000000000000000000000000000000000000000d7e44d53"),
						To:        lo.ToPtr(common.HexToAddress("0x00000000009726632680FB29d3F7A9734E3010E2")),
						Value:     lo.Must(new(big.Int).SetString("41000000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x49c8f0217c18bb6c1016ab256b214cc0b7cc49188a3d62a1265981598f57547e"),
						BlockNumber:       lo.Must(new(big.Int).SetString("20951604", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 8759489,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x20f697ad5"),
						GasUsed:           229892,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"),
								common.HexToHash("0x000000000000000000000000111111125421ca6dc452d289314280a0f8842a65"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000906c52e9163800"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20951604", 0)),
							TransactionHash: common.HexToHash("0x53ea61bbe1656b89995690376c7c8b525bc046307546814578b7df7ec141557d"),
							Index:           199,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000111111125421ca6dc452d289314280a0f8842a65"),
								common.HexToHash("0x0000000000000000000000006c59c360ef92ae28dcda2b60cf0ffc1c08fd3a8f"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000906c52e9163800"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20951604", 0)),
							TransactionHash: common.HexToHash("0x53ea61bbe1656b89995690376c7c8b525bc046307546814578b7df7ec141557d"),
							Index:           200,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x80eD012ef8dcd8e3A347D9034487AeBd991bfDA5"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000006c59c360ef92ae28dcda2b60cf0ffc1c08fd3a8f"),
								common.HexToHash("0x00000000000000000000000000000000009726632680fb29d3f7a9734e3010e2"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000036fa4416eb62355"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20951604", 0)),
							TransactionHash: common.HexToHash("0x53ea61bbe1656b89995690376c7c8b525bc046307546814578b7df7ec141557d"),
							Index:           201,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x6c59c360ef92AE28DcdA2B60CF0fFc1C08Fd3a8F"),
							Topics: []common.Hash{
								common.HexToHash("0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000004c3b6b3814f22f3ac000000000000000000000000000000000000000000000000c8388f18004f084d"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20951604", 0)),
							TransactionHash: common.HexToHash("0x53ea61bbe1656b89995690376c7c8b525bc046307546814578b7df7ec141557d"),
							Index:           202,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x6c59c360ef92AE28DcdA2B60CF0fFc1C08Fd3a8F"),
							Topics: []common.Hash{
								common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
								common.HexToHash("0x000000000000000000000000111111125421ca6dc452d289314280a0f8842a65"),
								common.HexToHash("0x00000000000000000000000000000000009726632680fb29d3f7a9734e3010e2"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000906c52e9163800000000000000000000000000000000000000000000000000036fa4416eb623550000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20951604", 0)),
							TransactionHash: common.HexToHash("0x53ea61bbe1656b89995690376c7c8b525bc046307546814578b7df7ec141557d"),
							Index:           203,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x80eD012ef8dcd8e3A347D9034487AeBd991bfDA5"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000000000000009726632680fb29d3f7a9734e3010e2"),
								common.HexToHash("0x0000000000000000000000009e7e6ee8607628362b25adf109d5704db3f4640f"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000036fa4416eb62355"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20951604", 0)),
							TransactionHash: common.HexToHash("0x53ea61bbe1656b89995690376c7c8b525bc046307546814578b7df7ec141557d"),
							Index:           204,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x53ea61bbe1656b89995690376c7c8b525bc046307546814578b7df7ec141557d"),
						TransactionIndex: 110,
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
				ID:       "0x53ea61bbe1656b89995690376c7c8b525bc046307546814578b7df7ec141557d",
				Network:  network.Ethereum,
				Index:    110,
				From:     "0x9E7E6ee8607628362B25AdF109d5704Db3f4640F",
				To:       rainbow.AddressRouter.String(),
				Type:     typex.ExchangeSwap,
				Platform: workerx.PlatformRainbow.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("2034200639509844")),
					Decimal: 18,
				},
				Calldata: &activityx.Calldata{
					FunctionHash: "0x3c2b9a7d",
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionTransfer,
						Platform: workerx.PlatformRainbow.String(),
						From:     "0x9E7E6ee8607628362B25AdF109d5704Db3f4640F",
						To:       rainbow.AddressRouter.String(),
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("348500000000000"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type:     typex.ExchangeSwap,
						Platform: workerx.PlatformRainbow.String(),
						From:     "0x9E7E6ee8607628362B25AdF109d5704Db3f4640F",
						To:       "0x9E7E6ee8607628362B25AdF109d5704Db3f4640F",
						Metadata: metadata.ExchangeSwap{
							From: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("40651500000000000"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
							To: metadata.Token{
								Address:  lo.ToPtr("0x80eD012ef8dcd8e3A347D9034487AeBd991bfDA5"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("247597105465926485"))),
								Name:     "Gigachad",
								Symbol:   "GIGA",
								Decimals: 9,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1728761231,
			},
			wantError: require.NoError,
		},
		{
			name: "Swap TITANX to ETH on Ethereum",
			arguments: struct {
				task   *source.Task
				config *config.Module
			}{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xeffb81993ffaa95a6aa1468a5ce98e6d1ca1f05a74ed0b5563b8d599b8ba49d6"),
						ParentHash:   common.HexToHash("0xccc6959e2240281b8c01ad6ea59cc164c16c2800f36216742656277812baf36c"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("20951621", 0)),
						GasLimit:     30000000,
						GasUsed:      11209801,
						Timestamp:    1728761435,
						BaseFee:      lo.Must(new(big.Int).SetString("8493926498", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xc271bf87f59dd80d59e4c8d90ba5e709304213c71831e1c9f79d9b77ff27fd3a"),
						From:      common.HexToAddress("0x5b527F39DB322e82D4faB0f6934AF3Ee7007AD1c"),
						Gas:       222940,
						GasPrice:  lo.Must(new(big.Int).SetString("8546056557", 10)),
						Hash:      common.HexToHash("0xc271bf87f59dd80d59e4c8d90ba5e709304213c71831e1c9f79d9b77ff27fd3a"),
						Input:     hexutil.MustDecode("0x999b6464000000000000000000000000f19308f923582a6f7c465e5ce7a9dc1bec6665b1000000000000000000000000111111125421ca6dc452d289314280a0f8842a6500000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000003f96ea19b9c9dbb8e000000000000000000000000000000000000000000000000000000001e32b478974000000000000000000000000000000000000000000000000000000000000000008883800a8e000000000000000000000000f19308f923582a6f7c465e5ce7a9dc1bec6665b1000000000000000000000000000000000000000003f96ea19b9c9dbb8e00000000000000000000000000000000000000000000000000000007c77ca5585f16e2380000000000000000000000c45a81bc23a64ea556ab4cdf08a86b61cdceea8bd6f29312000000000000000000000000000000000000000000000000d7e44d53"),
						To:        lo.ToPtr(common.HexToAddress("0x00000000009726632680FB29d3F7A9734E3010E2")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xeffb81993ffaa95a6aa1468a5ce98e6d1ca1f05a74ed0b5563b8d599b8ba49d6"),
						BlockNumber:       lo.Must(new(big.Int).SetString("20951621", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 9215918,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x1fd62796d"),
						GasUsed:           160057,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xF19308F923582A6f7c465e5CE7a9Dc1BEC6665B1"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000005b527f39db322e82d4fab0f6934af3ee7007ad1c"),
								common.HexToHash("0x00000000000000000000000000000000009726632680fb29d3f7a9734e3010e2"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000003f96ea19b9c9dbb8e000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20951621", 0)),
							TransactionHash: common.HexToHash("0xc271bf87f59dd80d59e4c8d90ba5e709304213c71831e1c9f79d9b77ff27fd3a"),
							Index:           263,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xF19308F923582A6f7c465e5CE7a9Dc1BEC6665B1"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x00000000000000000000000000000000009726632680fb29d3f7a9734e3010e2"),
								common.HexToHash("0x000000000000000000000000111111125421ca6dc452d289314280a0f8842a65"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000003f96ea19b9c9dbb8e000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20951621", 0)),
							TransactionHash: common.HexToHash("0xc271bf87f59dd80d59e4c8d90ba5e709304213c71831e1c9f79d9b77ff27fd3a"),
							Index:           264,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000c45a81bc23a64ea556ab4cdf08a86b61cdceea8b"),
								common.HexToHash("0x000000000000000000000000111111125421ca6dc452d289314280a0f8842a65"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000007db9a5970896c73"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20951621", 0)),
							TransactionHash: common.HexToHash("0xc271bf87f59dd80d59e4c8d90ba5e709304213c71831e1c9f79d9b77ff27fd3a"),
							Index:           265,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xF19308F923582A6f7c465e5CE7a9Dc1BEC6665B1"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x00000000000000000000000000000000009726632680fb29d3f7a9734e3010e2"),
								common.HexToHash("0x000000000000000000000000111111125421ca6dc452d289314280a0f8842a65"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20951621", 0)),
							TransactionHash: common.HexToHash("0xc271bf87f59dd80d59e4c8d90ba5e709304213c71831e1c9f79d9b77ff27fd3a"),
							Index:           266,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xF19308F923582A6f7c465e5CE7a9Dc1BEC6665B1"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000000000000009726632680fb29d3f7a9734e3010e2"),
								common.HexToHash("0x000000000000000000000000c45a81bc23a64ea556ab4cdf08a86b61cdceea8b"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000003f96ea19b9c9dbb8e000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20951621", 0)),
							TransactionHash: common.HexToHash("0xc271bf87f59dd80d59e4c8d90ba5e709304213c71831e1c9f79d9b77ff27fd3a"),
							Index:           267,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xc45A81BC23A64eA556ab4CdF08A86B61cdcEEA8b"),
							Topics: []common.Hash{
								common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"),
								common.HexToHash("0x000000000000000000000000111111125421ca6dc452d289314280a0f8842a65"),
								common.HexToHash("0x000000000000000000000000111111125421ca6dc452d289314280a0f8842a65"),
							},
							Data:            hexutil.MustDecode("0xfffffffffffffffffffffffffffffffffffffffffffffffff82465a68f76938d000000000000000000000000000000000000000003f96ea19b9c9dbb8e000000000000000000000000000000000000000000b5305a7ce383ffa91e03f1151092000000000000000000000000000000000000000000330e14b4badb6eecf8c5220000000000000000000000000000000000000000000000000000000000034779"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20951621", 0)),
							TransactionHash: common.HexToHash("0xc271bf87f59dd80d59e4c8d90ba5e709304213c71831e1c9f79d9b77ff27fd3a"),
							Index:           268,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65"),
								common.HexToHash("0x000000000000000000000000111111125421ca6dc452d289314280a0f8842a65"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000007db9a5970896c73"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20951621", 0)),
							TransactionHash: common.HexToHash("0xc271bf87f59dd80d59e4c8d90ba5e709304213c71831e1c9f79d9b77ff27fd3a"),
							Index:           269,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xc271bf87f59dd80d59e4c8d90ba5e709304213c71831e1c9f79d9b77ff27fd3a"),
						TransactionIndex: 103,
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
				ID:       "0xc271bf87f59dd80d59e4c8d90ba5e709304213c71831e1c9f79d9b77ff27fd3a",
				Network:  network.Ethereum,
				Index:    103,
				From:     "0x5b527F39DB322e82D4faB0f6934AF3Ee7007AD1c",
				To:       rainbow.AddressRouter.String(),
				Type:     typex.ExchangeSwap,
				Platform: workerx.PlatformRainbow.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("1367856174343749")),
					Decimal: 18,
				},
				Calldata: &activityx.Calldata{
					FunctionHash: "0x999b6464",
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionTransfer,
						Platform: workerx.PlatformRainbow.String(),
						From:     "0x5b527F39DB322e82D4faB0f6934AF3Ee7007AD1c",
						To:       rainbow.AddressRouter.String(),
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("4812835040315580"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type:     typex.ExchangeSwap,
						Platform: workerx.PlatformRainbow.String(),
						From:     "0x5b527F39DB322e82D4faB0f6934AF3Ee7007AD1c",
						To:       "0x5b527F39DB322e82D4faB0f6934AF3Ee7007AD1c",
						Metadata: metadata.ExchangeSwap{
							From: metadata.Token{
								Address:  lo.ToPtr("0xF19308F923582A6f7c465e5CE7a9Dc1BEC6665B1"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1230000000000000000000000000"))),
								Name:     "TITAN X",
								Symbol:   "TITANX",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
							To: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("566215887095950451"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1728761435,
			},
			wantError: require.NoError,
		},
		{
			name: "Swap PEIPEI to USDC on Ethereum",
			arguments: struct {
				task   *source.Task
				config *config.Module
			}{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x1a943b26a8490bca600c79909bf590cf9cb93177ff228ef6b9b3aeeabc180471"),
						ParentHash:   common.HexToHash("0x2454e595d850603d285aab7403edef2dd35eecd29e412bd87b84faa68b4c2701"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("20960347", 0)),
						GasLimit:     30000000,
						GasUsed:      13080242,
						Timestamp:    1728867131,
						BaseFee:      lo.Must(new(big.Int).SetString("7446526121", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xcd4e9395bb9ee875337c1e126dec94635f234c8107f5378a1f2e6eb5deb90f24"),
						From:      common.HexToAddress("0xbBb87B2E1aF93c326BA6fA10333c79A02c888ad8"),
						Gas:       344747,
						GasPrice:  lo.Must(new(big.Int).SetString("8453057815", 10)),
						Hash:      common.HexToHash("0xcd4e9395bb9ee875337c1e126dec94635f234c8107f5378a1f2e6eb5deb90f24"),
						Input:     hexutil.MustDecode("0x55e4b7be0000000000000000000000003ffeea07a27fab7ad1df5297fa75e77a43cb5790000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000111111125421ca6dc452d289314280a0f8842a6500000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000561ae8868f4c970d16d9ad38000000000000000000000000000000000000000000bb5d5e53e7f26f996904ac00000000000000000000000000000000000000000000000000000000000000a88770ba910000000000000000000000003ffeea07a27fab7ad1df5297fa75e77a43cb57900000000000000000000000000000000000000000555f8b283b64a49d7d70a88c00000000000000000000000000000000000000000000000000000000dc22049e08800000000000003b6d0340bf16540c857b4e32ce6c37d2f7725c8eec869b8b00000000000000003b6d0340b4e16d0168e52d35cacd2c6185b44281ec28c9dcd6f29312000000000000000000000000000000000000000000000000d7e44d53"),
						To:        lo.ToPtr(common.HexToAddress("0x00000000009726632680FB29d3F7A9734E3010E2")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x1a943b26a8490bca600c79909bf590cf9cb93177ff228ef6b9b3aeeabc180471"),
						BlockNumber:       lo.Must(new(big.Int).SetString("20960347", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 6717867,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x1f7d76d17"),
						GasUsed:           248433,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x3fFEea07a27Fab7ad1df5297fa75e77a43CB5790"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000bbb87b2e1af93c326ba6fa10333c79a02c888ad8"),
								common.HexToHash("0x00000000000000000000000000000000009726632680fb29d3f7a9734e3010e2"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000561ae8868f4c970d16d9ad38"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20960347", 0)),
							TransactionHash: common.HexToHash("0xcd4e9395bb9ee875337c1e126dec94635f234c8107f5378a1f2e6eb5deb90f24"),
							Index:           202,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x3fFEea07a27Fab7ad1df5297fa75e77a43CB5790"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000bbb87b2e1af93c326ba6fa10333c79a02c888ad8"),
								common.HexToHash("0x00000000000000000000000000000000009726632680fb29d3f7a9734e3010e2"),
							},
							Data:            hexutil.MustDecode("0xfffffffffffffffffffffffffffffffffffffffa6245c58bc3e3d5e07de38a18"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20960347", 0)),
							TransactionHash: common.HexToHash("0xcd4e9395bb9ee875337c1e126dec94635f234c8107f5378a1f2e6eb5deb90f24"),
							Index:           203,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x3fFEea07a27Fab7ad1df5297fa75e77a43CB5790"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x00000000000000000000000000000000009726632680fb29d3f7a9734e3010e2"),
								common.HexToHash("0x000000000000000000000000111111125421ca6dc452d289314280a0f8842a65"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000555f8b283b64a49d7d70a88c"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20960347", 0)),
							TransactionHash: common.HexToHash("0xcd4e9395bb9ee875337c1e126dec94635f234c8107f5378a1f2e6eb5deb90f24"),
							Index:           204,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x3fFEea07a27Fab7ad1df5297fa75e77a43CB5790"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000000000000009726632680fb29d3f7a9734e3010e2"),
								common.HexToHash("0x000000000000000000000000bf16540c857b4e32ce6c37d2f7725c8eec869b8b"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000555f8b283b64a49d7d70a88c"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20960347", 0)),
							TransactionHash: common.HexToHash("0xcd4e9395bb9ee875337c1e126dec94635f234c8107f5378a1f2e6eb5deb90f24"),
							Index:           205,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x3fFEea07a27Fab7ad1df5297fa75e77a43CB5790"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x00000000000000000000000000000000009726632680fb29d3f7a9734e3010e2"),
								common.HexToHash("0x000000000000000000000000111111125421ca6dc452d289314280a0f8842a65"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20960347", 0)),
							TransactionHash: common.HexToHash("0xcd4e9395bb9ee875337c1e126dec94635f234c8107f5378a1f2e6eb5deb90f24"),
							Index:           206,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000bf16540c857b4e32ce6c37d2f7725c8eec869b8b"),
								common.HexToHash("0x000000000000000000000000b4e16d0168e52d35cacd2c6185b44281ec28c9dc"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000155a6c93a446fff9"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20960347", 0)),
							TransactionHash: common.HexToHash("0xcd4e9395bb9ee875337c1e126dec94635f234c8107f5378a1f2e6eb5deb90f24"),
							Index:           207,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xbF16540c857B4e32cE6C37d2F7725C8eEC869B8b"),
							Topics: []common.Hash{
								common.HexToHash("0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000067f12992e52537d238137fb5ae000000000000000000000000000000000000000000000019fdfd59c84dc6e642"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20960347", 0)),
							TransactionHash: common.HexToHash("0xcd4e9395bb9ee875337c1e126dec94635f234c8107f5378a1f2e6eb5deb90f24"),
							Index:           208,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xbF16540c857B4e32cE6C37d2F7725C8eEC869B8b"),
							Topics: []common.Hash{
								common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
								common.HexToHash("0x000000000000000000000000111111125421ca6dc452d289314280a0f8842a65"),
								common.HexToHash("0x000000000000000000000000b4e16d0168e52d35cacd2c6185b44281ec28c9dc"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000555f8b283b64a49d7d70a88c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000155a6c93a446fff9"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20960347", 0)),
							TransactionHash: common.HexToHash("0xcd4e9395bb9ee875337c1e126dec94635f234c8107f5378a1f2e6eb5deb90f24"),
							Index:           209,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000b4e16d0168e52d35cacd2c6185b44281ec28c9dc"),
								common.HexToHash("0x00000000000000000000000000000000009726632680fb29d3f7a9734e3010e2"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000e0a0199d"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20960347", 0)),
							TransactionHash: common.HexToHash("0xcd4e9395bb9ee875337c1e126dec94635f234c8107f5378a1f2e6eb5deb90f24"),
							Index:           210,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc"),
							Topics: []common.Hash{
								common.HexToHash("0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000002785d6b5cb8c0000000000000000000000000000000000000000000003bf03b9e23d1cc952e9"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20960347", 0)),
							TransactionHash: common.HexToHash("0xcd4e9395bb9ee875337c1e126dec94635f234c8107f5378a1f2e6eb5deb90f24"),
							Index:           211,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc"),
							Topics: []common.Hash{
								common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
								common.HexToHash("0x000000000000000000000000111111125421ca6dc452d289314280a0f8842a65"),
								common.HexToHash("0x00000000000000000000000000000000009726632680fb29d3f7a9734e3010e2"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000155a6c93a446fff900000000000000000000000000000000000000000000000000000000e0a0199d0000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20960347", 0)),
							TransactionHash: common.HexToHash("0xcd4e9395bb9ee875337c1e126dec94635f234c8107f5378a1f2e6eb5deb90f24"),
							Index:           212,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000000000000009726632680fb29d3f7a9734e3010e2"),
								common.HexToHash("0x000000000000000000000000bbb87b2e1af93c326ba6fa10333c79a02c888ad8"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000e0a0199d"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20960347", 0)),
							TransactionHash: common.HexToHash("0xcd4e9395bb9ee875337c1e126dec94635f234c8107f5378a1f2e6eb5deb90f24"),
							Index:           213,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xcd4e9395bb9ee875337c1e126dec94635f234c8107f5378a1f2e6eb5deb90f24"),
						TransactionIndex: 70,
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
				ID:       "0xcd4e9395bb9ee875337c1e126dec94635f234c8107f5378a1f2e6eb5deb90f24",
				Network:  network.Ethereum,
				Index:    70,
				From:     "0xbBb87B2E1aF93c326BA6fA10333c79A02c888ad8",
				To:       rainbow.AddressRouter.String(),
				Type:     typex.ExchangeSwap,
				Platform: workerx.PlatformRainbow.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("2100018512153895")),
					Decimal: 18,
				},
				Calldata: &activityx.Calldata{
					FunctionHash: "0x55e4b7be",
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionTransfer,
						Platform: workerx.PlatformRainbow.String(),
						From:     "0xbBb87B2E1aF93c326BA6fA10333c79A02c888ad8",
						To:       rainbow.AddressRouter.String(),
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0x3fFEea07a27Fab7ad1df5297fa75e77a43CB5790"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("226510048390853047270048940"))),
							Name:     "PeiPei",
							Symbol:   "PEIPEI",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
					{
						Type:     typex.ExchangeSwap,
						Platform: workerx.PlatformRainbow.String(),
						From:     "0xbBb87B2E1aF93c326BA6fA10333c79A02c888ad8",
						To:       "0xbBb87B2E1aF93c326BA6fA10333c79A02c888ad8",
						Metadata: metadata.ExchangeSwap{
							From: metadata.Token{
								Address:  lo.ToPtr("0x3fFEea07a27Fab7ad1df5297fa75e77a43CB5790"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("26648240987159182031770463544"))),
								Name:     "PeiPei",
								Symbol:   "PEIPEI",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
							To: metadata.Token{
								Address:  lo.ToPtr("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("3768588701"))),
								Name:     "USD Coin",
								Symbol:   "USDC",
								Decimals: 6,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1728867131,
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

			// t.Log(string(lo.Must(json.MarshalIndent(activity, "", "\x20\x20"))))

			require.Equal(t, testcase.want, activity)
		})
	}
}
