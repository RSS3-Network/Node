package nouns_test

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/decentralized/contract/nouns"
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

func TestWorker_Ethereum(t *testing.T) {
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
			name: "Nouns Auction Settlement and New Noun Minting",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xeb56b4abb8cca128eddfddee5730f2fdf4a95b7b1f1feaccbc46ba6290b5a5bf"),
						ParentHash:   common.HexToHash("0x7c36940a2f532152df792f4f8143a2f49dab3a0ab98dcb87ef9de7bf286d3a32"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"),
						Number:       lo.Must(new(big.Int).SetString("20000137", 0)),
						GasLimit:     30000000,
						GasUsed:      27266432,
						Timestamp:    1717283063,
						BaseFee:      lo.Must(new(big.Int).SetString("5035867115", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xe745d7f27c6908cf94e087d3c93a0aab32926556597962639f073e11d7ebb627"),
						From:      common.HexToAddress("0x85906cF629ae1DA297548769ecE3e3E6a4f3288f"),
						Gas:       650000,
						GasPrice:  lo.Must(new(big.Int).SetString("15035867115", 10)),
						Hash:      common.HexToHash("0xe745d7f27c6908cf94e087d3c93a0aab32926556597962639f073e11d7ebb627"),
						Input:     hexutil.MustDecode("0xcbd5796f7c36940a2f532152df792f4f8143a2f49dab3a0ab98dcb87ef9de7bf286d3a32"),
						To:        lo.ToPtr(common.HexToAddress("0xb2341612271e122ff20905c9e389c3d7f0F222a1")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xeb56b4abb8cca128eddfddee5730f2fdf4a95b7b1f1feaccbc46ba6290b5a5bf"),
						BlockNumber:       lo.Must(new(big.Int).SetString("20000137", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 22496298,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x380351feb"),
						GasUsed:           396013,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03"),
							Topics: []common.Hash{
								common.HexToHash("0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724"),
								common.HexToHash("0x000000000000000000000000830bd73e4184cef73443c15111a1df14e495c706"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20000137", 0)),
							TransactionHash: common.HexToHash("0xe745d7f27c6908cf94e087d3c93a0aab32926556597962639f073e11d7ebb627"),
							Index:           412,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03"),
							Topics: []common.Hash{
								common.HexToHash("0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724"),
								common.HexToHash("0x000000000000000000000000560ddbb5ccaf91d27e91f0e7c0fa099e7ee179e6"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20000137", 0)),
							TransactionHash: common.HexToHash("0xe745d7f27c6908cf94e087d3c93a0aab32926556597962639f073e11d7ebb627"),
							Index:           413,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000830bd73e4184cef73443c15111a1df14e495c706"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000046c"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("20000137", 0)),
							TransactionHash: common.HexToHash("0xe745d7f27c6908cf94e087d3c93a0aab32926556597962639f073e11d7ebb627"),
							Index:           414,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000830bd73e4184cef73443c15111a1df14e495c706"),
								common.HexToHash("0x000000000000000000000000560ddbb5ccaf91d27e91f0e7c0fa099e7ee179e6"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000046c"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("20000137", 0)),
							TransactionHash: common.HexToHash("0xe745d7f27c6908cf94e087d3c93a0aab32926556597962639f073e11d7ebb627"),
							Index:           415,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x830BD73E4184ceF73443C15111a1DF14e495C706"),
							Topics: []common.Hash{
								common.HexToHash("0xc9f72b276a388619c6d185d146697036241880c36654b1a3ffdad07c24038d99"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000046c"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000560ddbb5ccaf91d27e91f0e7c0fa099e7ee179e600000000000000000000000000000000000000000000000053d263ffab5c0000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20000137", 0)),
							TransactionHash: common.HexToHash("0xe745d7f27c6908cf94e087d3c93a0aab32926556597962639f073e11d7ebb627"),
							Index:           416,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03"),
							Topics: []common.Hash{
								common.HexToHash("0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724"),
								common.HexToHash("0x000000000000000000000000830bd73e4184cef73443c15111a1df14e495c706"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20000137", 0)),
							TransactionHash: common.HexToHash("0xe745d7f27c6908cf94e087d3c93a0aab32926556597962639f073e11d7ebb627"),
							Index:           417,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000b1a32fc9f9d8b2cf86c068cae13108809547ef71"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000046d"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("20000137", 0)),
							TransactionHash: common.HexToHash("0xe745d7f27c6908cf94e087d3c93a0aab32926556597962639f073e11d7ebb627"),
							Index:           418,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000b1a32fc9f9d8b2cf86c068cae13108809547ef71"),
								common.HexToHash("0x000000000000000000000000830bd73e4184cef73443c15111a1df14e495c706"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000046d"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("20000137", 0)),
							TransactionHash: common.HexToHash("0xe745d7f27c6908cf94e087d3c93a0aab32926556597962639f073e11d7ebb627"),
							Index:           419,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03"),
							Topics: []common.Hash{
								common.HexToHash("0x1106ee9d020bfbb5ee34cf5535a5fbf024a011bd130078088cbf124ab3092478"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000046d"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000032000000000000000000000000000000000000000000000000000000000000003a0000000000000000000000000000000000000000000000000000000000000010"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20000137", 0)),
							TransactionHash: common.HexToHash("0xe745d7f27c6908cf94e087d3c93a0aab32926556597962639f073e11d7ebb627"),
							Index:           420,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x830BD73E4184ceF73443C15111a1DF14e495C706"),
							Topics: []common.Hash{
								common.HexToHash("0xd6eddd1118d71820909c1197aa966dbc15ed6f508554252169cc3d5ccac756ca"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000046d"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000665ba8f700000000000000000000000000000000000000000000000000000000665cfa77"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20000137", 0)),
							TransactionHash: common.HexToHash("0xe745d7f27c6908cf94e087d3c93a0aab32926556597962639f073e11d7ebb627"),
							Index:           421,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xe745d7f27c6908cf94e087d3c93a0aab32926556597962639f073e11d7ebb627"),
						TransactionIndex: 206,
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
				ID:      "0xe745d7f27c6908cf94e087d3c93a0aab32926556597962639f073e11d7ebb627",
				Network: network.Ethereum,
				Index:   206,
				From:    "0x85906cF629ae1DA297548769ecE3e3E6a4f3288f",
				To:      "0xb2341612271e122ff20905c9e389c3d7f0F222a1",
				Type:    typex.CollectibleAuction,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xcbd5796f",
				},
				Platform: workerx.PlatformNouns.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("5954398843812495")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.CollectibleAuction,
						Platform: workerx.PlatformNouns.String(),
						From:     "0x830BD73E4184ceF73443C15111a1DF14e495C706",
						To:       "0x560DDBB5cCaf91d27e91f0e7c0fA099e7Ee179E6",
						Metadata: metadata.CollectibleAuction{
							Action: metadata.ActionCollectibleAuctionFinalize,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03"),
								ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1132"))),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
								Name:     "Nouns",
								Symbol:   "NOUN",
								URI:      "",
								Decimals: 0,
								Standard: metadata.StandardERC721,
							},
							Cost: &metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("6040000000000000000"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
					{
						Type:     typex.CollectibleMint,
						Platform: workerx.PlatformNouns.String(),
						From:     "0x0000000000000000000000000000000000000000",
						To:       "0x830BD73E4184ceF73443C15111a1DF14e495C706",
						Metadata: metadata.CollectibleTransfer{
							Address:  lo.ToPtr("0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03"),
							ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1133"))),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
							Name:     "Nouns",
							Symbol:   "NOUN",
							Standard: metadata.StandardERC721,
						},
					},
					{
						Type:     typex.CollectibleAuction,
						Platform: workerx.PlatformNouns.String(),
						From:     "0x830BD73E4184ceF73443C15111a1DF14e495C706",
						To:       "0x830BD73E4184ceF73443C15111a1DF14e495C706",
						Metadata: metadata.CollectibleAuction{
							Action: metadata.ActionCollectibleAuctionCreate,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03"),
								ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1133"))),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
								Name:     "Nouns",
								Symbol:   "NOUN",
								Standard: metadata.StandardERC721,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1717283063,
			},
			wantError: require.NoError,
		},
		{
			name: "Test Nouns Proposal",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xe0eafdef36c8689a22bce10e29e8f11f0064fe7d5b280b297425d3e9c62f07df"),
						ParentHash:   common.HexToHash("0xdbf429a079ad2d494e4c4abebb5b31156a4d1a4d12818208f022effc240ec7b4"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8"),
						Number:       lo.Must(new(big.Int).SetString("14584582", 0)),
						GasLimit:     30000000,
						GasUsed:      15758858,
						Timestamp:    1649953105,
						BaseFee:      lo.Must(new(big.Int).SetString("66704115205", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x1a9b6076a439ddd4003411ee723f3f2a6d6ef2761d7bf8e6b0d89f7ff178aa96"),
						From:      common.HexToAddress("0x83fCFe8Ba2FEce9578F0BbaFeD4Ebf5E915045B9"),
						Gas:       528496,
						GasPrice:  lo.Must(new(big.Int).SetString("68204115205", 10)),
						Hash:      common.HexToHash("0x1a9b6076a439ddd4003411ee723f3f2a6d6ef2761d7bf8e6b0d89f7ff178aa96"),
						Input:     hexutil.MustDecode("0xda95691a00000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000240000000000000000000000000000000000000000000000000000000000000032000000000000000000000000000000000000000000000000000000000000000020000000000000000000000006f3e6272a167e8accb32072d08e0957f9c79223d0000000000000000000000006f3e6272a167e8accb32072d08e0957f9c79223d00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000185f736574566f74696e6744656c61792875696e7432353629000000000000000000000000000000000000000000000000000000000000000000000000000000195f736574566f74696e67506572696f642875696e743235362900000000000000000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000067d9000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000081ce00000000000000000000000000000000000000000000000000000000000001782320496e63726561736520566f74696e672044656c617920616e6420566f74696e6720506572696f640a0a417320746865206e756d626572206f66206e6f756e65727320616e64206e6f756e73207375622d636f6d6d756e697469657320696e6372656173652c2077652077696c6c20736565206d6f72652070726f706f73616c7320617070656172206f6e2d636861696e20776974686f7574207072696f72206f66662d636861696e20636f6f7264696e6174696f6e2e205468697320697320612070726f706f73616c20746f20696e6372656173652074686520e28098566f74696e672044656c6179e2809920746f203420646179732c20616e642074686520e28098566f74696e6720506572696f64e2809920746f203520646179732c20736f2074686174206e6f756e6572732068617665206d6f72652074696d6520746f2064696c6967656e63652070726f706f73616c7320616e6420746f2064656369646520686f7720746865792077696c6c20766f74652e0000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x6f3E6272A167e8AcCb32072d08E0957F9c79223d")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xe0eafdef36c8689a22bce10e29e8f11f0064fe7d5b280b297425d3e9c62f07df"),
						BlockNumber:       lo.Must(new(big.Int).SetString("14584582", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 15612743,
						EffectiveGasPrice: hexutil.MustDecodeBig("0xfe1483505"),
						GasUsed:           520988,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x6f3E6272A167e8AcCb32072d08E0957F9c79223d"),
							Topics: []common.Hash{
								common.HexToHash("0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000003a00000000000000000000000083fcfe8ba2fece9578f0bbafed4ebf5e915045b90000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000001e000000000000000000000000000000000000000000000000000000000000002c00000000000000000000000000000000000000000000000000000000000debe5a0000000000000000000000000000000000000000000000000000000000df0b5800000000000000000000000000000000000000000000000000000000000003a000000000000000000000000000000000000000000000000000000000000000020000000000000000000000006f3e6272a167e8accb32072d08e0957f9c79223d0000000000000000000000006f3e6272a167e8accb32072d08e0957f9c79223d00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000185f736574566f74696e6744656c61792875696e7432353629000000000000000000000000000000000000000000000000000000000000000000000000000000195f736574566f74696e67506572696f642875696e743235362900000000000000000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000067d9000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000081ce00000000000000000000000000000000000000000000000000000000000001782320496e63726561736520566f74696e672044656c617920616e6420566f74696e6720506572696f640a0a417320746865206e756d626572206f66206e6f756e65727320616e64206e6f756e73207375622d636f6d6d756e697469657320696e6372656173652c2077652077696c6c20736565206d6f72652070726f706f73616c7320617070656172206f6e2d636861696e20776974686f7574207072696f72206f66662d636861696e20636f6f7264696e6174696f6e2e205468697320697320612070726f706f73616c20746f20696e6372656173652074686520e28098566f74696e672044656c6179e2809920746f203420646179732c20616e642074686520e28098566f74696e6720506572696f64e2809920746f203520646179732c20736f2074686174206e6f756e6572732068617665206d6f72652074696d6520746f2064696c6967656e63652070726f706f73616c7320616e6420746f2064656369646520686f7720746865792077696c6c20766f74652e0000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("14584582", 0)),
							TransactionHash: common.HexToHash("0x1a9b6076a439ddd4003411ee723f3f2a6d6ef2761d7bf8e6b0d89f7ff178aa96"),
							Index:           264,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x6f3E6272A167e8AcCb32072d08E0957F9c79223d"),
							Topics: []common.Hash{
								common.HexToHash("0x6af0134faa0f9290c1d686d55012aca80302d31d5c856e4bc7954f7613dc7f87"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000003a00000000000000000000000083fcfe8ba2fece9578f0bbafed4ebf5e915045b9000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001c0000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000debe5a0000000000000000000000000000000000000000000000000000000000df0b580000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001b00000000000000000000000000000000000000000000000000000000000003e000000000000000000000000000000000000000000000000000000000000000020000000000000000000000006f3e6272a167e8accb32072d08e0957f9c79223d0000000000000000000000006f3e6272a167e8accb32072d08e0957f9c79223d00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000185f736574566f74696e6744656c61792875696e7432353629000000000000000000000000000000000000000000000000000000000000000000000000000000195f736574566f74696e67506572696f642875696e743235362900000000000000000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000067d9000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000081ce00000000000000000000000000000000000000000000000000000000000001782320496e63726561736520566f74696e672044656c617920616e6420566f74696e6720506572696f640a0a417320746865206e756d626572206f66206e6f756e65727320616e64206e6f756e73207375622d636f6d6d756e697469657320696e6372656173652c2077652077696c6c20736565206d6f72652070726f706f73616c7320617070656172206f6e2d636861696e20776974686f7574207072696f72206f66662d636861696e20636f6f7264696e6174696f6e2e205468697320697320612070726f706f73616c20746f20696e6372656173652074686520e28098566f74696e672044656c6179e2809920746f203420646179732c20616e642074686520e28098566f74696e6720506572696f64e2809920746f203520646179732c20736f2074686174206e6f756e6572732068617665206d6f72652074696d6520746f2064696c6967656e63652070726f706f73616c7320616e6420746f2064656369646520686f7720746865792077696c6c20766f74652e0000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("14584582", 0)),
							TransactionHash: common.HexToHash("0x1a9b6076a439ddd4003411ee723f3f2a6d6ef2761d7bf8e6b0d89f7ff178aa96"),
							Index:           265,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x1a9b6076a439ddd4003411ee723f3f2a6d6ef2761d7bf8e6b0d89f7ff178aa96"),
						TransactionIndex: 176,
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
				ID:      "0x1a9b6076a439ddd4003411ee723f3f2a6d6ef2761d7bf8e6b0d89f7ff178aa96",
				Network: network.Ethereum,
				Index:   176,
				From:    "0x83fCFe8Ba2FEce9578F0BbaFeD4Ebf5E915045B9",
				To:      "0x6f3E6272A167e8AcCb32072d08E0957F9c79223d",
				Type:    typex.GovernanceProposal,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xda95691a",
				},
				Platform: workerx.PlatformNouns.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("35533525572422540")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.GovernanceProposal,
						Platform: workerx.PlatformNouns.String(),
						From:     "0x83fCFe8Ba2FEce9578F0BbaFeD4Ebf5E915045B9",
						To:       "0x6f3E6272A167e8AcCb32072d08E0957F9c79223d",
						Metadata: metadata.GovernanceProposal{
							ID:         "58",
							Body:       "# Increase Voting Delay and Voting Period\n\nAs the number of nouners and nouns sub-communities increase, we will see more proposals appear on-chain without prior off-chain coordination. This is a proposal to increase the ‘Voting Delay’ to 4 days, and the ‘Voting Period’ to 5 days, so that nouners have more time to diligence proposals and to decide how they will vote.",
							StartBlock: "14597722",
							EndBlock:   "14617432",
							Options:    []string{},
							Link:       "",
						},
					},
				},
				Status:    true,
				Timestamp: 1649953105,
			},
			wantError: require.NoError,
		},
		{
			name: "Test Auction Bid",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x106a1e42d66563f95d759fb1d3bc8a4d6db7e2cbfe18e7458ac999420840871b"),
						ParentHash:   common.HexToHash("0x27dcd5e378746539bb247818b85e02b12d6f8f7f38da6243733e7ecfc3a7d293"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xa09F34312545f1ac7Bf7680527F79c20dc146E59"),
						Number:       lo.Must(new(big.Int).SetString("20670837", 0)),
						GasLimit:     30000000,
						GasUsed:      6293656,
						Timestamp:    1725375995,
						BaseFee:      lo.Must(new(big.Int).SetString("5036299740", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x3dddab2a5fb8db6627905d8eabd3acfa12df233a5ff567ab77918681e3605867"),
						From:      common.HexToAddress("0x087c1135031db541Ff3Df2E2512e7493cCeE0d69"),
						Gas:       101662,
						GasPrice:  lo.Must(new(big.Int).SetString("6461561800", 10)),
						Hash:      common.HexToHash("0x3dddab2a5fb8db6627905d8eabd3acfa12df233a5ff567ab77918681e3605867"),
						Input:     hexutil.MustDecode("0x6a761202000000000000000000000000830bd73e4184cef73443c15111a1df14e495c7060000000000000000000000000000000000000000000000002195912da4720000000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001a00000000000000000000000000000000000000000000000000000000000000024659dd2b400000000000000000000000000000000000000000000000000000000000004d300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008266664498f7e9b49988ea9f8c216722e6fa506fb09c62a8fb58e9ad927d238fd76aa6bc6c78b93ea494121c25cda202216d29aa14e9ccd56fb991699cf98b3f631cf111cdb468c5caa5d29b56c9085a69607952430a63915f1d793cefa4a7f10327693bda1f71f06c70762e03ef6a49e1bba7d5504fcc9fd99a0ea149cf8cc64a6e1b000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xeA2BA25E59f84e87B50d3231077E777BD6879589")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x106a1e42d66563f95d759fb1d3bc8a4d6db7e2cbfe18e7458ac999420840871b"),
						BlockNumber:       lo.Must(new(big.Int).SetString("20670837", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 1637732,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x181239bc8"),
						GasUsed:           98735,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x830BD73E4184ceF73443C15111a1DF14e495C706"),
							Topics: []common.Hash{
								common.HexToHash("0x1159164c56f277e6fc99c11731bd380e0347deb969b75523398734c252706ea3"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000004d3"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000ea2ba25e59f84e87b50d3231077e777bd68795890000000000000000000000000000000000000000000000002195912da47200000000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20670837", 0)),
							TransactionHash: common.HexToHash("0x3dddab2a5fb8db6627905d8eabd3acfa12df233a5ff567ab77918681e3605867"),
							Index:           33,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xeA2BA25E59f84e87B50d3231077E777BD6879589"),
							Topics: []common.Hash{
								common.HexToHash("0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e"),
							},
							Data:            hexutil.MustDecode("0x7a8034ae8ca5b178dc413787b0fc5274cf82bf7a68fc68f7b90cce6d6958d5130000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20670837", 0)),
							TransactionHash: common.HexToHash("0x3dddab2a5fb8db6627905d8eabd3acfa12df233a5ff567ab77918681e3605867"),
							Index:           34,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x3dddab2a5fb8db6627905d8eabd3acfa12df233a5ff567ab77918681e3605867"),
						TransactionIndex: 30,
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
				ID:      "0x3dddab2a5fb8db6627905d8eabd3acfa12df233a5ff567ab77918681e3605867",
				Network: network.Ethereum,
				Index:   30,
				From:    "0x087c1135031db541Ff3Df2E2512e7493cCeE0d69",
				To:      "0xeA2BA25E59f84e87B50d3231077E777BD6879589",
				Type:    typex.CollectibleAuction,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x6a761202",
				},
				Platform: workerx.PlatformNouns.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("637982304323000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.CollectibleAuction,
						Platform: workerx.PlatformNouns.String(),
						From:     "0xeA2BA25E59f84e87B50d3231077E777BD6879589",
						To:       "0x830BD73E4184ceF73443C15111a1DF14e495C706",
						Metadata: metadata.CollectibleAuction{
							Action: metadata.ActionCollectibleAuctionBid,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03"),
								ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1235"))),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
								Name:     "Nouns",
								Symbol:   "NOUN",
								Standard: metadata.StandardERC721,
							},
							Cost: &metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2420000000000000000"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1725375995,
			},
			wantError: require.NoError,
		},
		{
			name: "Test Auction Bid",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x106a1e42d66563f95d759fb1d3bc8a4d6db7e2cbfe18e7458ac999420840871b"),
						ParentHash:   common.HexToHash("0x27dcd5e378746539bb247818b85e02b12d6f8f7f38da6243733e7ecfc3a7d293"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xa09F34312545f1ac7Bf7680527F79c20dc146E59"),
						Number:       lo.Must(new(big.Int).SetString("20670837", 0)),
						GasLimit:     30000000,
						GasUsed:      6293656,
						Timestamp:    1725375995,
						BaseFee:      lo.Must(new(big.Int).SetString("5036299740", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x3dddab2a5fb8db6627905d8eabd3acfa12df233a5ff567ab77918681e3605867"),
						From:      common.HexToAddress("0x087c1135031db541Ff3Df2E2512e7493cCeE0d69"),
						Gas:       101662,
						GasPrice:  lo.Must(new(big.Int).SetString("6461561800", 10)),
						Hash:      common.HexToHash("0x3dddab2a5fb8db6627905d8eabd3acfa12df233a5ff567ab77918681e3605867"),
						Input:     hexutil.MustDecode("0x6a761202000000000000000000000000830bd73e4184cef73443c15111a1df14e495c7060000000000000000000000000000000000000000000000002195912da4720000000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001a00000000000000000000000000000000000000000000000000000000000000024659dd2b400000000000000000000000000000000000000000000000000000000000004d300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008266664498f7e9b49988ea9f8c216722e6fa506fb09c62a8fb58e9ad927d238fd76aa6bc6c78b93ea494121c25cda202216d29aa14e9ccd56fb991699cf98b3f631cf111cdb468c5caa5d29b56c9085a69607952430a63915f1d793cefa4a7f10327693bda1f71f06c70762e03ef6a49e1bba7d5504fcc9fd99a0ea149cf8cc64a6e1b000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xeA2BA25E59f84e87B50d3231077E777BD6879589")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x106a1e42d66563f95d759fb1d3bc8a4d6db7e2cbfe18e7458ac999420840871b"),
						BlockNumber:       lo.Must(new(big.Int).SetString("20670837", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 1637732,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x181239bc8"),
						GasUsed:           98735,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x830BD73E4184ceF73443C15111a1DF14e495C706"),
							Topics: []common.Hash{
								common.HexToHash("0x1159164c56f277e6fc99c11731bd380e0347deb969b75523398734c252706ea3"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000004d3"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000ea2ba25e59f84e87b50d3231077e777bd68795890000000000000000000000000000000000000000000000002195912da47200000000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20670837", 0)),
							TransactionHash: common.HexToHash("0x3dddab2a5fb8db6627905d8eabd3acfa12df233a5ff567ab77918681e3605867"),
							Index:           33,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xeA2BA25E59f84e87B50d3231077E777BD6879589"),
							Topics: []common.Hash{
								common.HexToHash("0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e"),
							},
							Data:            hexutil.MustDecode("0x7a8034ae8ca5b178dc413787b0fc5274cf82bf7a68fc68f7b90cce6d6958d5130000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20670837", 0)),
							TransactionHash: common.HexToHash("0x3dddab2a5fb8db6627905d8eabd3acfa12df233a5ff567ab77918681e3605867"),
							Index:           34,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x3dddab2a5fb8db6627905d8eabd3acfa12df233a5ff567ab77918681e3605867"),
						TransactionIndex: 30,
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
				ID:      "0x3dddab2a5fb8db6627905d8eabd3acfa12df233a5ff567ab77918681e3605867",
				Network: network.Ethereum,
				Index:   30,
				From:    "0x087c1135031db541Ff3Df2E2512e7493cCeE0d69",
				To:      "0xeA2BA25E59f84e87B50d3231077E777BD6879589",
				Type:    typex.CollectibleAuction,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x6a761202",
				},
				Platform: workerx.PlatformNouns.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("637982304323000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.CollectibleAuction,
						Platform: workerx.PlatformNouns.String(),
						From:     "0xeA2BA25E59f84e87B50d3231077E777BD6879589",
						To:       "0x830BD73E4184ceF73443C15111a1DF14e495C706",
						Metadata: metadata.CollectibleAuction{
							Action: metadata.ActionCollectibleAuctionBid,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03"),
								ID:       lo.ToPtr(lo.Must(decimal.NewFromString("1235"))),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
								Name:     "Nouns",
								Symbol:   "NOUN",
								Standard: metadata.StandardERC721,
							},
							Cost: &metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2420000000000000000"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1725375995,
			},
			wantError: require.NoError,
		},
		{
			name: "Test Governance Vote",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x0ecb37379b269b0d0f76495e115dd1a7be2bc5fd1ab2a925789a52b967f2ba44"),
						ParentHash:   common.HexToHash("0x5c2bfa694d9f38186f1bd19ab757ca29410af46422866b49eb531a863da6ffc2"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("20676677", 0)),
						GasLimit:     30000000,
						GasUsed:      21256147,
						Timestamp:    1725446375,
						BaseFee:      lo.Must(new(big.Int).SetString("2033766393", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xf38432040c7487c6eb59b78d124bb3e65cfc6359f0016fa20e193357800356b1"),
						From:      common.HexToAddress("0x0a049e014999A489b3D7174B8f70D4200b0Ce79B"),
						Gas:       158488,
						GasPrice:  lo.Must(new(big.Int).SetString("2533766393", 10)),
						Hash:      common.HexToHash("0xf38432040c7487c6eb59b78d124bb3e65cfc6359f0016fa20e193357800356b1"),
						Input:     hexutil.MustDecode("0x8136730f000000000000000000000000000000000000000000000000000000000000026e00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000c2466f723a2033207c20416761696e73743a2031207c204162737461696e3a20300a0a2b666f7220e2809420403837626f6e65730a0a2b666f7220e280942040726f7862790a0a2b666f7220416c736f20686f70696e6720746f2073656520736f6d6520636f6f6c20e28c90e297a82de297a820696e737069726564206675726e69747572652f696e66726120696e207468652067616c6c65727920e280942040776964656579656b61726c0a0a2b616761696e737420e28094204062697862697465000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x6f3E6272A167e8AcCb32072d08E0957F9c79223d")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x0ecb37379b269b0d0f76495e115dd1a7be2bc5fd1ab2a925789a52b967f2ba44"),
						BlockNumber:       lo.Must(new(big.Int).SetString("20676677", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 10636723,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x970634f9"),
						GasUsed:           109385,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x6f3E6272A167e8AcCb32072d08E0957F9c79223d"),
							Topics: []common.Hash{
								common.HexToHash("0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4"),
								common.HexToHash("0x0000000000000000000000000a049e014999a489b3d7174b8f70d4200b0ce79b"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000026e00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c2466f723a2033207c20416761696e73743a2031207c204162737461696e3a20300a0a2b666f7220e2809420403837626f6e65730a0a2b666f7220e280942040726f7862790a0a2b666f7220416c736f20686f70696e6720746f2073656520736f6d6520636f6f6c20e28c90e297a82de297a820696e737069726564206675726e69747572652f696e66726120696e207468652067616c6c65727920e280942040776964656579656b61726c0a0a2b616761696e737420e28094204062697862697465000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20676677", 0)),
							TransactionHash: common.HexToHash("0xf38432040c7487c6eb59b78d124bb3e65cfc6359f0016fa20e193357800356b1"),
							Index:           285,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x6f3E6272A167e8AcCb32072d08E0957F9c79223d"),
							Topics: []common.Hash{
								common.HexToHash("0x651cc9d78606507fdcfc4f37ec37a744d612b1d8f5a73564190577c4f0edb0b6"),
								common.HexToHash("0x0000000000000000000000000a049e014999a489b3d7174b8f70d4200b0ce79b"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000026e"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000006"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("20676677", 0)),
							TransactionHash: common.HexToHash("0xf38432040c7487c6eb59b78d124bb3e65cfc6359f0016fa20e193357800356b1"),
							Index:           286,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x6f3E6272A167e8AcCb32072d08E0957F9c79223d"),
							Topics: []common.Hash{
								common.HexToHash("0xfabef36fd46c4c3a6ad676521be5367a4dfdbf3faa68d8e826003b1752d68f4f"),
								common.HexToHash("0x0000000000000000000000000a049e014999a489b3d7174b8f70d4200b0ce79b"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000e99bd4de8d150000000000000000000000000000000000000000000000000000000000000001"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20676677", 0)),
							TransactionHash: common.HexToHash("0xf38432040c7487c6eb59b78d124bb3e65cfc6359f0016fa20e193357800356b1"),
							Index:           287,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xf38432040c7487c6eb59b78d124bb3e65cfc6359f0016fa20e193357800356b1"),
						TransactionIndex: 101,
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
				ID:      "0xf38432040c7487c6eb59b78d124bb3e65cfc6359f0016fa20e193357800356b1",
				Network: network.Ethereum,
				Index:   101,
				From:    "0x0a049e014999A489b3D7174B8f70D4200b0Ce79B",
				To:      "0x6f3E6272A167e8AcCb32072d08E0957F9c79223d",
				Type:    typex.GovernanceVote,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x8136730f",
				},
				Platform: workerx.PlatformNouns.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("277156036898305")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.GovernanceVote,
						Platform: workerx.PlatformNouns.String(),
						From:     "0x0a049e014999A489b3D7174B8f70D4200b0Ce79B",
						To:       "0x6f3E6272A167e8AcCb32072d08E0957F9c79223d",
						Metadata: metadata.GovernanceVote{
							Action: metadata.ActionGovernanceFor,
							Count:  2,
							Reason: "For: 3 | Against: 1 | Abstain: 0\n\n+for — @87bones\n\n+for — @roxby\n\n+for Also hoping to see some cool ⌐◨-◨ inspired furniture/infra in the gallery — @wideeyekarl\n\n+against — @bixbite",
							Proposal: metadata.GovernanceProposal{
								ID:   "622",
								Link: "https://nouns.wtf/vote/622",
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1725446375,
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
