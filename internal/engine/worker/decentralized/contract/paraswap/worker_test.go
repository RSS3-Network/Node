package paraswap_test

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/v2/config"
	source "github.com/rss3-network/node/v2/internal/engine/protocol/ethereum"
	worker "github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/paraswap"
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

func TestWorker_Paraswap(t *testing.T) {
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
			name: "Paraswap Token Swap (REZ to ETH)",
			arguments: struct {
				task   *source.Task
				config *config.Module
			}{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xa20374ef4f54e40ea0b12bea342d038217c1c52e9579f8784868422b8757619e"),
						ParentHash:   common.HexToHash("0xbc369c24220496d87af0ebddb713ad7fa96a6756088efd3289d99a0e03162cd3"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"),
						Number:       lo.Must(new(big.Int).SetString("19771930", 0)),
						GasLimit:     30000000,
						GasUsed:      11600032,
						Timestamp:    1714526063,
						BaseFee:      lo.Must(new(big.Int).SetString("6866994584", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x2c668402f39a485ef57cbfe31e24d7e0f9342a4584de7f94889fd7b02f90d77b"),
						From:      common.HexToAddress("0x7c614681E6F6ECCa2E84D311D3f6dda6c460E55B"),
						Gas:       306971,
						GasPrice:  lo.Must(new(big.Int).SetString("6956994584", 10)),
						Hash:      common.HexToHash("0x2c668402f39a485ef57cbfe31e24d7e0f9342a4584de7f94889fd7b02f90d77b"),
						Input:     hexutil.MustDecode("0xa6886da900000000000000000000000000000000000000000000000000000000000000200000000000000000000000003b50805453023a91a8bf641e279401a0b23fa6f9000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee000000000000000000000000e592427a0aece92de3edee1f18e0157c05861564000000000000000000000000000000000000000000000052f103edb66ba80000000000000000000000000000000000000000000000000000010338174146251a0000000000000000000000000000000000000000000000000104858f02911c490100000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000006631ebc900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002201704308b9f1748d39431ba15f2bedcc500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002b3b50805453023a91a8bf641e279401a0b23fa6f9000bb8c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000007c614681e6f6ecca2e84d311d3f6dda6c460e55b000000000000000000000000216b4b4ba9f3e719726886d34a177484278bfcae000000000000000000000000000000000000000000000052f103edb66ba800000000000000000000000000000000000000000000000000000000000066358bde000000000000000000000000000000000000000000000000000000000000001bdcd2799ac8a0174f1f44a3be56a31351b0fbfefb7620766293cfbec2f2343ff0653c54e08e4c433cd970a875c8eba3480fc58299d32e784ecbe541af4c451007"),
						To:        lo.ToPtr(common.HexToAddress("0xDEF171Fe48CF0115B1d80b88dc8eAB59176FEe57")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xa20374ef4f54e40ea0b12bea342d038217c1c52e9579f8784868422b8757619e"),
						BlockNumber:       lo.Must(new(big.Int).SetString("19771930", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 5998846,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x19eab5018"),
						GasUsed:           200757,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x3B50805453023a91a8bf641e279401a0b23FA6F9"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x0000000000000000000000007c614681e6f6ecca2e84d311d3f6dda6c460e55b"),
								common.HexToHash("0x000000000000000000000000216b4b4ba9f3e719726886d34a177484278bfcae"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000052f103edb66ba80000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19771930", 0)),
							TransactionHash: common.HexToHash("0x2c668402f39a485ef57cbfe31e24d7e0f9342a4584de7f94889fd7b02f90d77b"),
							Index:           160,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x3B50805453023a91a8bf641e279401a0b23FA6F9"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000007c614681e6f6ecca2e84d311d3f6dda6c460e55b"),
								common.HexToHash("0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000052f103edb66ba80000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19771930", 0)),
							TransactionHash: common.HexToHash("0x2c668402f39a485ef57cbfe31e24d7e0f9342a4584de7f94889fd7b02f90d77b"),
							Index:           161,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000059c68e0eca9f0844d9d4cee0bedf23c8dee462c5"),
								common.HexToHash("0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000104579503cdcadc"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19771930", 0)),
							TransactionHash: common.HexToHash("0x2c668402f39a485ef57cbfe31e24d7e0f9342a4584de7f94889fd7b02f90d77b"),
							Index:           162,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x3B50805453023a91a8bf641e279401a0b23FA6F9"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57"),
								common.HexToHash("0x00000000000000000000000059c68e0eca9f0844d9d4cee0bedf23c8dee462c5"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000052f103edb66ba80000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19771930", 0)),
							TransactionHash: common.HexToHash("0x2c668402f39a485ef57cbfe31e24d7e0f9342a4584de7f94889fd7b02f90d77b"),
							Index:           163,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x59C68e0ECa9F0844d9D4CeE0BEDf23c8dEE462c5"),
							Topics: []common.Hash{
								common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"),
								common.HexToHash("0x000000000000000000000000e592427a0aece92de3edee1f18e0157c05861564"),
								common.HexToHash("0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000052f103edb66ba80000fffffffffffffffffffffffffffffffffffffffffffffffffefba86afc323524000000000000000000000000000000000000000001c632f4c47ef9821f94f0b1000000000000000000000000000000000000000000000ef1fb85fa4e8dd48af1fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe7b8e"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19771930", 0)),
							TransactionHash: common.HexToHash("0x2c668402f39a485ef57cbfe31e24d7e0f9342a4584de7f94889fd7b02f90d77b"),
							Index:           164,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65"),
								common.HexToHash("0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000104579503cdcadc"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19771930", 0)),
							TransactionHash: common.HexToHash("0x2c668402f39a485ef57cbfe31e24d7e0f9342a4584de7f94889fd7b02f90d77b"),
							Index:           165,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xDEF171Fe48CF0115B1d80b88dc8eAB59176FEe57"),
							Topics: []common.Hash{
								common.HexToHash("0xd2d73da2b5fd52cd654d8fd1b514ad57355bad741de639e3a1c3a20dd9f17347"),
								common.HexToHash("0x0000000000000000000000007c614681e6f6ecca2e84d311d3f6dda6c460e55b"),
								common.HexToHash("0x0000000000000000000000003b50805453023a91a8bf641e279401a0b23fa6f9"),
								common.HexToHash("0x000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"),
							},
							Data:            hexutil.MustDecode("0x1704308b9f1748d39431ba15f2bedcc500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000040000000000000000000000000007c614681e6f6ecca2e84d311d3f6dda6c460e55b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000052f103edb66ba800000000000000000000000000000000000000000000000000000104579503cdcadc0000000000000000000000000000000000000000000000000104858f02911c49"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19771930", 0)),
							TransactionHash: common.HexToHash("0x2c668402f39a485ef57cbfe31e24d7e0f9342a4584de7f94889fd7b02f90d77b"),
							Index:           166,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x2c668402f39a485ef57cbfe31e24d7e0f9342a4584de7f94889fd7b02f90d77b"),
						TransactionIndex: 71,
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
				ID:       "0x2c668402f39a485ef57cbfe31e24d7e0f9342a4584de7f94889fd7b02f90d77b",
				Network:  network.Ethereum,
				Index:    71,
				From:     "0x7c614681E6F6ECCa2E84D311D3f6dda6c460E55B",
				To:       "0xDEF171Fe48CF0115B1d80b88dc8eAB59176FEe57",
				Type:     typex.ExchangeSwap,
				Platform: workerx.PlatformParaswap.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("1396665361700088")),
					Decimal: 18,
				},
				Calldata: &activityx.Calldata{
					FunctionHash: "0xa6886da9",
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.ExchangeSwap,
						Platform: workerx.PlatformParaswap.String(),
						From:     "0x7c614681E6F6ECCa2E84D311D3f6dda6c460E55B",
						To:       "0x7c614681E6F6ECCa2E84D311D3f6dda6c460E55B",
						Metadata: metadata.ExchangeSwap{
							From: metadata.Token{
								Address:  lo.ToPtr("0x3B50805453023a91a8bf641e279401a0b23FA6F9"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1530000000000000000000"))),
								Name:     "Renzo",
								Symbol:   "REZ",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
							To: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("73279791470332636"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1714526063,
			},
			wantError: require.NoError,
		},
		{
			name: "Paraswap Token Swap (WETH to USDC)",
			arguments: struct {
				task   *source.Task
				config *config.Module
			}{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x0e8ad5c0a8197420174e4622129960d4d9e551839eb0defab002c48b4f00b800"),
						ParentHash:   common.HexToHash("0x97206ab738d3f885446763cf64dc4b354e3781a3d044659cb9e8511885acb17b"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("19771969", 0)),
						GasLimit:     30000000,
						GasUsed:      12027224,
						Timestamp:    1714526531,
						BaseFee:      lo.Must(new(big.Int).SetString("6640967522", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x647baef594b4acaab7e8e4e18a56d442b679a7e70d7b552da8250d675004176b"),
						From:      common.HexToAddress("0x47ed2b8a79f6D67dFD0C50313e260d07F5AC7E95"),
						Gas:       2026878,
						GasPrice:  lo.Must(new(big.Int).SetString("6690967522", 10)),
						Hash:      common.HexToHash("0x647baef594b4acaab7e8e4e18a56d442b679a7e70d7b552da8250d675004176b"),
						Input:     hexutil.MustDecode("0x6092eae500000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000000000000000000000000000026d6a0823aceefc0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000380000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb4800000000000000000000000000000000000000000000000026d6a0823aceefc000000000000000000000000000000000000000000000000000000001f48502e5000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57000000000000000000000000216b4b4ba9f3e719726886d34a177484278bfcae00000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000264a6886da90000000000000000000000000000000000000000000000000000000000000020000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000e592427a0aece92de3edee1f18e0157c0586156400000000000000000000000000000000000000000000000026d6a0823aceefc000000000000000000000000000000000000000000000000000000001ef83ae6200000000000000000000000000000000000000000000000000000001f48502e50100000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000006631ed9c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002201d0deb7b7778495c981e4402121ab77b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002bc02aaa39b223fe8d0a0e5c4f27ead9083c756cc20001f4a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xe6FC6A67607EFE4C44224c16be190980779F2dC1")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x0e8ad5c0a8197420174e4622129960d4d9e551839eb0defab002c48b4f00b800"),
						BlockNumber:       lo.Must(new(big.Int).SetString("19771969", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 11301628,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x18ed00fe2"),
						GasUsed:           999677,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000004d5f47fa6a74757f35c14fd3a6ef8e3c9bc514e8"),
								common.HexToHash("0x0000000000000000000000006e5743e228614b000104b7d9b95d8981c92c26f0"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000026d6a0823aceefc0"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19771969", 0)),
							TransactionHash: common.HexToHash("0x647baef594b4acaab7e8e4e18a56d442b679a7e70d7b552da8250d675004176b"),
							Index:           262,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x0000000000000000000000006e5743e228614b000104b7d9b95d8981c92c26f0"),
								common.HexToHash("0x000000000000000000000000216b4b4ba9f3e719726886d34a177484278bfcae"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000026d6a0823aceefc0"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19771969", 0)),
							TransactionHash: common.HexToHash("0x647baef594b4acaab7e8e4e18a56d442b679a7e70d7b552da8250d675004176b"),
							Index:           263,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000006e5743e228614b000104b7d9b95d8981c92c26f0"),
								common.HexToHash("0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000026d6a0823aceefc0"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19771969", 0)),
							TransactionHash: common.HexToHash("0x647baef594b4acaab7e8e4e18a56d442b679a7e70d7b552da8250d675004176b"),
							Index:           264,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000088e6a0c2ddd26feeb64f039a2c41296fcb3f5640"),
								common.HexToHash("0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000001f48502e5"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19771969", 0)),
							TransactionHash: common.HexToHash("0x647baef594b4acaab7e8e4e18a56d442b679a7e70d7b552da8250d675004176b"),
							Index:           265,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57"),
								common.HexToHash("0x00000000000000000000000088e6a0c2ddd26feeb64f039a2c41296fcb3f5640"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000026d6a0823aceefc0"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19771969", 0)),
							TransactionHash: common.HexToHash("0x647baef594b4acaab7e8e4e18a56d442b679a7e70d7b552da8250d675004176b"),
							Index:           266,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640"),
							Topics: []common.Hash{
								common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"),
								common.HexToHash("0x000000000000000000000000e592427a0aece92de3edee1f18e0157c05861564"),
								common.HexToHash("0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57"),
							},
							Data:            hexutil.MustDecode("0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffe0b7afd1b00000000000000000000000000000000000000000000000026d6a0823aceefc0000000000000000000000000000000000000474b5fc02aca0c9606d9228e78f10000000000000000000000000000000000000000000000006e880344f151bb84000000000000000000000000000000000000000000000000000000000002fe99"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19771969", 0)),
							TransactionHash: common.HexToHash("0x647baef594b4acaab7e8e4e18a56d442b679a7e70d7b552da8250d675004176b"),
							Index:           267,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57"),
								common.HexToHash("0x0000000000000000000000006e5743e228614b000104b7d9b95d8981c92c26f0"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000001f48502e5"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19771969", 0)),
							TransactionHash: common.HexToHash("0x647baef594b4acaab7e8e4e18a56d442b679a7e70d7b552da8250d675004176b"),
							Index:           268,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xDEF171Fe48CF0115B1d80b88dc8eAB59176FEe57"),
							Topics: []common.Hash{
								common.HexToHash("0xd2d73da2b5fd52cd654d8fd1b514ad57355bad741de639e3a1c3a20dd9f17347"),
								common.HexToHash("0x0000000000000000000000006e5743e228614b000104b7d9b95d8981c92c26f0"),
								common.HexToHash("0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
								common.HexToHash("0x000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
							},
							Data:            hexutil.MustDecode("0x1d0deb7b7778495c981e4402121ab77b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000040000000000000000000000000006e5743e228614b000104b7d9b95d8981c92c26f0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000026d6a0823aceefc000000000000000000000000000000000000000000000000000000001f48502e500000000000000000000000000000000000000000000000000000001f48502e5"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19771969", 0)),
							TransactionHash: common.HexToHash("0x647baef594b4acaab7e8e4e18a56d442b679a7e70d7b552da8250d675004176b"),
							Index:           269,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x647baef594b4acaab7e8e4e18a56d442b679a7e70d7b552da8250d675004176b"),
						TransactionIndex: 135,
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
				ID:       "0x647baef594b4acaab7e8e4e18a56d442b679a7e70d7b552da8250d675004176b",
				Network:  network.Ethereum,
				Index:    135,
				From:     "0x47ed2b8a79f6D67dFD0C50313e260d07F5AC7E95",
				To:       "0xe6FC6A67607EFE4C44224c16be190980779F2dC1",
				Type:     typex.ExchangeSwap,
				Platform: workerx.PlatformParaswap.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("6688806339490394")),
					Decimal: 18,
				},
				Calldata: &activityx.Calldata{
					FunctionHash: "0x6092eae5",
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.ExchangeSwap,
						Platform: workerx.PlatformParaswap.String(),
						From:     "0x6E5743E228614B000104b7D9B95d8981c92c26F0",
						To:       "0x6E5743E228614B000104b7D9B95d8981c92c26F0",
						Metadata: metadata.ExchangeSwap{
							From: metadata.Token{
								Address:  lo.ToPtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2798600699650174912"))),
								Name:     "Wrapped Ether",
								Symbol:   "WETH",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
							To: metadata.Token{
								Address:  lo.ToPtr("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("8397325029"))),
								Name:     "USD Coin",
								Symbol:   "USDC",
								Decimals: 6,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1714526531,
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
