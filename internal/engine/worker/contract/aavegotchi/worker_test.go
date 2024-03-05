package aavegotchi_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	"github.com/rss3-network/node/internal/engine/worker/contract/aavegotchi"
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
		task *source.Task
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      *schema.Feed
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "List ERC1155",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkPolygon,
					ChainID: 137,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x46708a88a6f844dfa9e4497324f5d21a1b466a718d7db3c841f80a537bf41688"),
						ParentHash:   common.HexToHash("0x76d0fb11caecdbdf501c8c9b69b5763678884a3dc110ef8fe76a4eff973b26cd"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x7c7379531b2aEE82e4Ca06D4175D13b9CBEafd49"),
						Number:       lo.Must(new(big.Int).SetString("43673506", 0)),
						GasLimit:     30000000,
						GasUsed:      24868693,
						Timestamp:    1686225179,
						BaseFee:      lo.Must(new(big.Int).SetString("126883313209", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xad7aec56979b5f4bb6e2d216504058b3768c4d8c9a34504e240116b78f993fe9"),
						From:      common.HexToAddress("0xA2faa3405a734c04aE713AAa837E6cEcC2cAee9F"),
						Gas:       466714,
						GasPrice:  lo.Must(new(big.Int).SetString("172918381771", 10)),
						Hash:      common.HexToHash("0xad7aec56979b5f4bb6e2d216504058b3768c4d8c9a34504e240116b78f993fe9"),
						Input:     hexutil.MustDecode("0xe8da2bfa00000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d00000000000000000000000000000000000000000000000000000000000000f000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000049b9ca9a694340000"),
						To:        lo.ToPtr(common.HexToAddress("0x86935F11C86623deC8a25696E1C19a8659CbF95d")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   lo.Must(new(big.Int).SetString("137", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x46708a88a6f844dfa9e4497324f5d21a1b466a718d7db3c841f80a537bf41688"),
						BlockNumber:       lo.Must(new(big.Int).SetString("43673506", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 5002005,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x2842bd1ccb"),
						GasUsed:           452446,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x86935F11C86623deC8a25696E1C19a8659CbF95d"),
							Topics: []common.Hash{
								common.HexToHash("0xf99a07edc22c88f73b19df09fffa9b11c9ab5c5dd6dfc2f6e103a95f5b40e2fc"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000056e59"),
								common.HexToHash("0x000000000000000000000000a2faa3405a734c04ae713aaa837e6cecc2caee9f"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d00000000000000000000000000000000000000000000000000000000000000f000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000049b9ca9a694340000000000000000000000000000000000000000000000000000000000006481c11b"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43673506", 0)),
							TransactionHash: common.HexToHash("0xad7aec56979b5f4bb6e2d216504058b3768c4d8c9a34504e240116b78f993fe9"),
							Index:           103,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000a2faa3405a734c04ae713aaa837e6cecc2caee9f"),
								common.HexToHash("0x000000000000000000000000ffffffffffffffffffffffffffffffffffffffff"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000016345785d8a0000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43673506", 0)),
							TransactionHash: common.HexToHash("0xad7aec56979b5f4bb6e2d216504058b3768c4d8c9a34504e240116b78f993fe9"),
							Index:           104,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000a2faa3405a734c04ae713aaa837e6cecc2caee9f"),
								common.HexToHash("0x00000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d"),
							},
							Data:            hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffffffbec63ee6d3a9ffff"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43673506", 0)),
							TransactionHash: common.HexToHash("0xad7aec56979b5f4bb6e2d216504058b3768c4d8c9a34504e240116b78f993fe9"),
							Index:           105,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x0000000000000000000000000000000000001010"),
							Topics: []common.Hash{
								common.HexToHash("0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000001010"),
								common.HexToHash("0x000000000000000000000000a2faa3405a734c04ae713aaa837e6cecc2caee9f"),
								common.HexToHash("0x0000000000000000000000007c7379531b2aee82e4ca06d4175d13b9cbeafd49"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000049ff4dbbf1679c000000000000000000000000000000000000000000000002754a1530a99036bc000000000000000000000000000000000000000000021e817cfcbd8af5afd291000000000000000000000000000000000000000000000002750015e2ed9ecf20000000000000000000000000000000000000000000021e817d46bcd8b1a13a2d"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43673506", 0)),
							TransactionHash: common.HexToHash("0xad7aec56979b5f4bb6e2d216504058b3768c4d8c9a34504e240116b78f993fe9"),
							Index:           106,
							Removed:         false,
						},
						},
						Status:           1,
						TransactionHash:  common.HexToHash("0xad7aec56979b5f4bb6e2d216504058b3768c4d8c9a34504e240116b78f993fe9"),
						TransactionIndex: 27,
					},
				},
			},
			want: &schema.Feed{
				ID:      common.HexToHash("0xad7aec56979b5f4bb6e2d216504058b3768c4d8c9a34504e240116b78f993fe9").String(),
				Network: filter.NetworkPolygon,
				Index:   27,
				From:    common.HexToAddress("0xa2faa3405a734c04ae713aaa837e6cecc2caee9f").String(),
				To:      common.HexToAddress("0x86935f11c86623dec8a25696e1c19a8659cbf95d").String(),
				Type:    filter.TypeMetaverseTrade,
				Calldata: &schema.Calldata{
					FunctionHash: "0xe8da2bfa",
				},
				Platform: lo.ToPtr(filter.PlatformAavegotchi),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("78236230158761866")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeMetaverseTrade,
						Platform: filter.PlatformAavegotchi.String(),
						From:     common.HexToAddress("0xa2faa3405a734c04ae713aaa837e6cecc2caee9f").String(),
						To:       common.HexToAddress("0x86935f11c86623dec8a25696e1c19a8659cbf95d").String(),
						Metadata: metadata.MetaverseTrade{
							Action: metadata.ActionMetaverseTradeList,
							Token: metadata.Token{
								Name:     "Aavegotchi",
								Symbol:   "GOTCHI",
								Standard: 721,
								URI:      "https://app.aavegotchi.com/metadata/aavegotchis/240",
								Address:  lo.ToPtr(common.HexToAddress("0x86935F11C86623deC8a25696E1C19a8659CbF95d").String()),
								Value:    lo.ToPtr(decimal.NewFromInt(1)),
								ID:       lo.ToPtr(decimal.NewFromInt(240)),
							},
							Cost: metadata.Token{
								Name:     "Aavegotchi GHST Token (PoS)",
								Symbol:   "GHST",
								Decimals: 18,
								Standard: 20,
								Address:  lo.ToPtr(common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7").String()),
								Value:    lo.ToPtr(decimal.NewFromInt(100000000000000000)),
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1686225179,
			},
			wantError: require.NoError,
		},
		{
			name: "Trade ERC1155",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkPolygon,
					ChainID: 137,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xc4ae62923cbbee8c757490bac85812244f31829d5c53f1d3a65a9dbc8d4d6cbf"),
						ParentHash:   common.HexToHash("0xbdbe9a8258da772e17252fdcc21f4df4c3f52055d7cc8a039adbbf00e79a94f5"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1D25c827AbD466387BDA00b429fE728627D6EEE6"),
						Number:       lo.Must(new(big.Int).SetString("43675511", 0)),
						GasLimit:     29765755,
						GasUsed:      7874288,
						Timestamp:    1686229535,
						BaseFee:      lo.Must(new(big.Int).SetString("148881512536", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x05c2ae65d685af45d9e7cac72543931ecc4f6e62a437d804b6e63a3aecb2f40d"),
						From:      common.HexToAddress("0xb9AD10B590bCd2b0dB23d0005b2DB0d53d9a1cF0"),
						Gas:       734327,
						GasPrice:  lo.Must(new(big.Int).SetString("179685181888", 10)),
						Hash:      common.HexToHash("0x05c2ae65d685af45d9e7cac72543931ecc4f6e62a437d804b6e63a3aecb2f40d"),
						Input:     hexutil.MustDecode("0x575ae8760000000000000000000000000000000000000000000000000000000000056e530000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000001e5b8fa8fe2ac0000"),
						To:        lo.ToPtr(common.HexToAddress("0x86935F11C86623deC8a25696E1C19a8659CbF95d")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("137", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xc4ae62923cbbee8c757490bac85812244f31829d5c53f1d3a65a9dbc8d4d6cbf"),
						BlockNumber:       lo.Must(new(big.Int).SetString("43675511", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 6897068,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x29d61249c0"),
						GasUsed:           675305,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x000000000000000000000000d4151c984e6cf33e04ffaaf06c3374b2926ecc64"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000009b6e64a8ec60000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675511", 0)),
							TransactionHash: common.HexToHash("0x05c2ae65d685af45d9e7cac72543931ecc4f6e62a437d804b6e63a3aecb2f40d"),
							Index:           150,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x00000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d"),
							},
							Data:            hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffffb8bc42990c9f7bffff"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675511", 0)),
							TransactionHash: common.HexToHash("0x05c2ae65d685af45d9e7cac72543931ecc4f6e62a437d804b6e63a3aecb2f40d"),
							Index:           151,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x000000000000000000000000b208f8bb431f580cc4b216826affb128cd1431ab"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000004db732547630000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675511", 0)),
							TransactionHash: common.HexToHash("0x05c2ae65d685af45d9e7cac72543931ecc4f6e62a437d804b6e63a3aecb2f40d"),
							Index:           152,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x00000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d"),
							},
							Data:            hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffffb8b76725e75818ffff"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675511", 0)),
							TransactionHash: common.HexToHash("0x05c2ae65d685af45d9e7cac72543931ecc4f6e62a437d804b6e63a3aecb2f40d"),
							Index:           153,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x00000000000000000000000027df5c6dcd360f372e23d5e63645ec0072d0c098"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000026db992a3b18000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675511", 0)),
							TransactionHash: common.HexToHash("0x05c2ae65d685af45d9e7cac72543931ecc4f6e62a437d804b6e63a3aecb2f40d"),
							Index:           154,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x00000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d"),
							},
							Data:            hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffffb8b4f96c54b4677fff"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675511", 0)),
							TransactionHash: common.HexToHash("0x05c2ae65d685af45d9e7cac72543931ecc4f6e62a437d804b6e63a3aecb2f40d"),
							Index:           155,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x000000000000000000000000cbef46a7cbe1f46a94ab77501eaa32596ab3c538"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000001d4b8e78d68d18000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675511", 0)),
							TransactionHash: common.HexToHash("0x05c2ae65d685af45d9e7cac72543931ecc4f6e62a437d804b6e63a3aecb2f40d"),
							Index:           156,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x00000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d"),
							},
							Data:            hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffffb6e04084c74b95ffff"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675511", 0)),
							TransactionHash: common.HexToHash("0x05c2ae65d685af45d9e7cac72543931ecc4f6e62a437d804b6e63a3aecb2f40d"),
							Index:           157,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x86935F11C86623deC8a25696E1C19a8659CbF95d"),
							Topics: []common.Hash{
								common.HexToHash("0x276ac39b0f98215592b4487434c1618e3e527c57ca2f471ac93eb4eab1694091"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000056e53"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006481d21f"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675511", 0)),
							TransactionHash: common.HexToHash("0x05c2ae65d685af45d9e7cac72543931ecc4f6e62a437d804b6e63a3aecb2f40d"),
							Index:           158,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x58de9AaBCaeEC0f69883C94318810ad79Cc6a44f"),
							Topics: []common.Hash{
								common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"),
								common.HexToHash("0x00000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d"),
								common.HexToHash("0x000000000000000000000000cbef46a7cbe1f46a94ab77501eaa32596ab3c538"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000008d0000000000000000000000000000000000000000000000000000000000000001"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675511", 0)),
							TransactionHash: common.HexToHash("0x05c2ae65d685af45d9e7cac72543931ecc4f6e62a437d804b6e63a3aecb2f40d"),
							Index:           159,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x86935F11C86623deC8a25696E1C19a8659CbF95d"),
							Topics: []common.Hash{
								common.HexToHash("0x4004185f45cf5f331fe63afc9aa5aa0a0d8cfa1d6bc5d8c6cb0304136c49515c"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000056e53"),
								common.HexToHash("0x000000000000000000000000cbef46a7cbe1f46a94ab77501eaa32596ab3c538"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf000000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d000000000000000000000000000000000000000000000000000000000000008d0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000001e5b8fa8fe2ac0000000000000000000000000000000000000000000000000000000000006481d21f"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675511", 0)),
							TransactionHash: common.HexToHash("0x05c2ae65d685af45d9e7cac72543931ecc4f6e62a437d804b6e63a3aecb2f40d"),
							Index:           160,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x0000000000000000000000000000000000001010"),
							Topics: []common.Hash{
								common.HexToHash("0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000001010"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x0000000000000000000000001d25c827abd466387bda00b429fe728627d6eee6"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000049e7313b27bfa8000000000000000000000000000000000000000000000000b85459f3e481dace000000000000000000000000000000000000000000000012983b9d20e48881aa000000000000000000000000000000000000000000000000b80a72c2a95a1b26000000000000000000000000000000000000000000000012988584521fb04152"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675511", 0)),
							TransactionHash: common.HexToHash("0x05c2ae65d685af45d9e7cac72543931ecc4f6e62a437d804b6e63a3aecb2f40d"),
							Index:           161,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x05c2ae65d685af45d9e7cac72543931ecc4f6e62a437d804b6e63a3aecb2f40d"),
						TransactionIndex: 39,
					},
				},
			},
			want: &schema.Feed{
				ID:      common.HexToHash("0x05c2ae65d685af45d9e7cac72543931ecc4f6e62a437d804b6e63a3aecb2f40d").String(),
				Network: filter.NetworkPolygon,
				Index:   39,
				From:    common.HexToAddress("0xb9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0").String(),
				To:      common.HexToAddress("0x86935f11c86623dec8a25696e1c19a8659cbf95d").String(),
				Type:    filter.TypeMetaverseTrade,
				Calldata: &schema.Calldata{
					FunctionHash: "0x575ae876",
				},
				Platform: lo.ToPtr(filter.PlatformAavegotchi),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("121342301754875840")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeMetaverseTrade,
						Platform: filter.PlatformAavegotchi.String(),
						From:     common.HexToAddress("0xcbef46a7cbe1f46a94ab77501eaa32596ab3c538").String(),
						To:       common.HexToAddress("0xb9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0").String(),
						Metadata: metadata.MetaverseTrade{
							Action: metadata.ActionMetaverseTradeBuy,
							Token: metadata.Token{
								Name:     "Aavegotchi",
								Symbol:   "GOTCHI",
								Standard: 721,
								URI:      "https://app.aavegotchi.com/metadata/aavegotchis/141",
								Address:  lo.ToPtr(common.HexToAddress("0x86935F11C86623deC8a25696E1C19a8659CbF95d").String()),
								Value:    lo.ToPtr(decimal.NewFromInt(1)),
								ID:       lo.ToPtr(decimal.NewFromInt(141)),
							},
							Cost: metadata.Token{
								Name:     "Aavegotchi GHST Token (PoS)",
								Symbol:   "GHST",
								Decimals: 18,
								Standard: 20,
								Address:  lo.ToPtr(common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7").String()),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("35000000000000000000"))),
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1686229535,
			},
			wantError: require.NoError,
		},
		{
			name: "List ERC721",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkPolygon,
					ChainID: 137,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xcb95b74b6e727c8991bed5351f1a5d78d0c5251089380e7d6e34be0f74181dca"),
						ParentHash:   common.HexToHash("0x0a60c11c8512344d092f6c1a25d0b013959dbd72433e3e25f9e25b6203e61848"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x67B94473D81D0cd00849D563C94d0432Ac988B49"),
						Number:       lo.Must(new(big.Int).SetString("43702826", 0)),
						GasLimit:     30000000,
						GasUsed:      17693748,
						Timestamp:    1686291411,
						BaseFee:      lo.Must(new(big.Int).SetString("149606312586", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x4e94b27356abebab22b676fc7c9a4018da8dfa102cfdf93434f80529c066c88e"),
						From:      common.HexToAddress("0xB7039fbd301CfAa7C66BA43fc10e2c60aa18F785"),
						Gas:       413959,
						GasPrice:  lo.Must(new(big.Int).SetString("464568926610", 10)),
						Hash:      common.HexToHash("0x4e94b27356abebab22b676fc7c9a4018da8dfa102cfdf93434f80529c066c88e"),
						Input:     hexutil.MustDecode("0x31b0b51400000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d0000000000000000000000000000000000000000000000000000000000001c86000000000000000000000000000000000000000000000010357563d481cc0000"),
						To:        lo.ToPtr(common.HexToAddress("0x86935F11C86623deC8a25696E1C19a8659CbF95d")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("137", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xcb95b74b6e727c8991bed5351f1a5d78d0c5251089380e7d6e34be0f74181dca"),
						BlockNumber:       lo.Must(new(big.Int).SetString("43702826", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 400522,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x6c2a774192"),
						GasUsed:           400522,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x86935F11C86623deC8a25696E1C19a8659CbF95d"),
							Topics: []common.Hash{
								common.HexToHash("0xd9b541cd1d9650dbbcb9b244efe4aa9455a9e7bb329f80cd30ab25275534c72d"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000041a22"),
								common.HexToHash("0x000000000000000000000000b7039fbd301cfaa7c66ba43fc10e2c60aa18f785"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000003"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d0000000000000000000000000000000000000000000000000000000000001c86000000000000000000000000000000000000000000000010357563d481cc0000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43702826", 0)),
							TransactionHash: common.HexToHash("0x4e94b27356abebab22b676fc7c9a4018da8dfa102cfdf93434f80529c066c88e"),
							Index:           0,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000b7039fbd301cfaa7c66ba43fc10e2c60aa18f785"),
								common.HexToHash("0x000000000000000000000000ffffffffffffffffffffffffffffffffffffffff"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000016345785d8a0000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43702826", 0)),
							TransactionHash: common.HexToHash("0x4e94b27356abebab22b676fc7c9a4018da8dfa102cfdf93434f80529c066c88e"),
							Index:           1,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000b7039fbd301cfaa7c66ba43fc10e2c60aa18f785"),
								common.HexToHash("0x00000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d"),
							},
							Data:            hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffffdb82955e57b50fffff"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43702826", 0)),
							TransactionHash: common.HexToHash("0x4e94b27356abebab22b676fc7c9a4018da8dfa102cfdf93434f80529c066c88e"),
							Index:           2,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x0000000000000000000000000000000000001010"),
							Topics: []common.Hash{
								common.HexToHash("0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000001010"),
								common.HexToHash("0x000000000000000000000000b7039fbd301cfaa7c66ba43fc10e2c60aa18f785"),
								common.HexToHash("0x00000000000000000000000067b94473d81d0cd00849d563c94d0432ac988b49"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000001c02c430f194a5000000000000000000000000000000000000000000000000014419e12483759d4000000000000000000000000000000000000000000000af1e43ec1f065576c3d000000000000000000000000000000000000000000000000128171cf391e0f84000000000000000000000000000000000000000000000af1e5feee337470b68d"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43702826", 0)),
							TransactionHash: common.HexToHash("0x4e94b27356abebab22b676fc7c9a4018da8dfa102cfdf93434f80529c066c88e"),
							Index:           3,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x4e94b27356abebab22b676fc7c9a4018da8dfa102cfdf93434f80529c066c88e"),
						TransactionIndex: 0,
					},
				},
			},
			want: &schema.Feed{
				ID:      common.HexToHash("0x4e94b27356abebab22b676fc7c9a4018da8dfa102cfdf93434f80529c066c88e").String(),
				Network: filter.NetworkPolygon,
				Index:   0,
				From:    common.HexToAddress("0xb7039fbd301cfaa7c66ba43fc10e2c60aa18f785").String(),
				To:      common.HexToAddress("0x86935f11c86623dec8a25696e1c19a8659cbf95d").String(),
				Type:    filter.TypeMetaverseTrade,
				Calldata: &schema.Calldata{
					FunctionHash: "0x31b0b514",
				},
				Platform: lo.ToPtr(filter.PlatformAavegotchi),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("186070075623690420")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeMetaverseTrade,
						Platform: filter.PlatformAavegotchi.String(),
						From:     common.HexToAddress("0xb7039fbd301cfaa7c66ba43fc10e2c60aa18f785").String(),
						To:       common.HexToAddress("0x86935f11c86623dec8a25696e1c19a8659cbf95d").String(),
						Metadata: metadata.MetaverseTrade{
							Action: metadata.ActionMetaverseTradeList,
							Token: metadata.Token{
								Name:     "Aavegotchi",
								Symbol:   "GOTCHI",
								Standard: 721,
								URI:      "https://app.aavegotchi.com/metadata/aavegotchis/7302",
								Address:  lo.ToPtr(common.HexToAddress("0x86935F11C86623deC8a25696E1C19a8659CbF95d").String()),
								ID:       lo.ToPtr(decimal.NewFromInt(7302)),
							},
							Cost: metadata.Token{
								Name:     "Aavegotchi GHST Token (PoS)",
								Symbol:   "GHST",
								Decimals: 18,
								Standard: 20,
								Address:  lo.ToPtr(common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7").String()),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("100000000000000000"))),
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1686291411,
			},
			wantError: require.NoError,
		},
		{
			name: "Buy ERC721",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkPolygon,
					ChainID: 137,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x914f2743c948677f68a969587831db27d357492d11c6a4600b2380a065a71906"),
						ParentHash:   common.HexToHash("0x2138ca4330f23c209b268dd49eb4448d04ab0e2c86c7e9fc627277252016826c"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x7c7379531b2aEE82e4Ca06D4175D13b9CBEafd49"),
						Number:       lo.Must(new(big.Int).SetString("43675427", 0)),
						GasLimit:     29649772,
						GasUsed:      20016523,
						Timestamp:    1686229357,
						BaseFee:      lo.Must(new(big.Int).SetString("168437875378", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x23d7f6c49c11e947066b7cd65c384a63e34715aaca4f0ed6eb2b876b8b3f74f9"),
						From:      common.HexToAddress("0xb9AD10B590bCd2b0dB23d0005b2DB0d53d9a1cF0"),
						Gas:       356961,
						GasPrice:  lo.Must(new(big.Int).SetString("463325087112", 10)),
						Hash:      common.HexToHash("0x23d7f6c49c11e947066b7cd65c384a63e34715aaca4f0ed6eb2b876b8b3f74f9"),
						Input:     hexutil.MustDecode("0x66609c8800000000000000000000000000000000000000000000000000000000000419f000000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d0000000000000000000000000000000000000000000000252248deb6e69400000000000000000000000000000000000000000000000000000000000000000463000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
						To:        lo.ToPtr(common.HexToAddress("0x86935F11C86623deC8a25696E1C19a8659CbF95d")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("137", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x914f2743c948677f68a969587831db27d357492d11c6a4600b2380a065a71906"),
						BlockNumber:       lo.Must(new(big.Int).SetString("43675427", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 2704826,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x6be053c588"),
						GasUsed:           327048,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x86935F11C86623deC8a25696E1C19a8659CbF95d"),
							Topics: []common.Hash{
								common.HexToHash("0xd309bba073dd463d9d78e92766d06da7cde7d0e6e83149168a7bc913c9ecc442"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000419f0"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000006481d16d"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675427", 0)),
							TransactionHash: common.HexToHash("0x23d7f6c49c11e947066b7cd65c384a63e34715aaca4f0ed6eb2b876b8b3f74f9"),
							Index:           60,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x000000000000000000000000d4151c984e6cf33e04ffaaf06c3374b2926ecc64"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000be202d6a0eda0000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675427", 0)),
							TransactionHash: common.HexToHash("0x23d7f6c49c11e947066b7cd65c384a63e34715aaca4f0ed6eb2b876b8b3f74f9"),
							Index:           61,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x00000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d"),
							},
							Data:            hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffffdd2a2230a405fbffff"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675427", 0)),
							TransactionHash: common.HexToHash("0x23d7f6c49c11e947066b7cd65c384a63e34715aaca4f0ed6eb2b876b8b3f74f9"),
							Index:           62,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x000000000000000000000000b208f8bb431f580cc4b216826affb128cd1431ab"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000005f1016b5076d0000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675427", 0)),
							TransactionHash: common.HexToHash("0x23d7f6c49c11e947066b7cd65c384a63e34715aaca4f0ed6eb2b876b8b3f74f9"),
							Index:           63,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x00000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d"),
							},
							Data:            hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffffdccb1219eefe8effff"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675427", 0)),
							TransactionHash: common.HexToHash("0x23d7f6c49c11e947066b7cd65c384a63e34715aaca4f0ed6eb2b876b8b3f74f9"),
							Index:           64,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x00000000000000000000000027df5c6dcd360f372e23d5e63645ec0072d0c098"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000002f880b5a83b68000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675427", 0)),
							TransactionHash: common.HexToHash("0x23d7f6c49c11e947066b7cd65c384a63e34715aaca4f0ed6eb2b876b8b3f74f9"),
							Index:           65,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x00000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d"),
							},
							Data:            hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffffdc9b8a0e947ad87fff"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675427", 0)),
							TransactionHash: common.HexToHash("0x23d7f6c49c11e947066b7cd65c384a63e34715aaca4f0ed6eb2b876b8b3f74f9"),
							Index:           66,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x000000000000000000000000a2faa3405a734c04ae713aaa837e6cecc2caee9f"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000023d5908f3d4c968000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675427", 0)),
							TransactionHash: common.HexToHash("0x23d7f6c49c11e947066b7cd65c384a63e34715aaca4f0ed6eb2b876b8b3f74f9"),
							Index:           67,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x00000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d"),
							},
							Data:            hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffffb8c5f97f572e41ffff"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675427", 0)),
							TransactionHash: common.HexToHash("0x23d7f6c49c11e947066b7cd65c384a63e34715aaca4f0ed6eb2b876b8b3f74f9"),
							Index:           68,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x86935F11C86623deC8a25696E1C19a8659CbF95d"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000a2faa3405a734c04ae713aaa837e6cecc2caee9f"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000463"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("43675427", 0)),
							TransactionHash: common.HexToHash("0x23d7f6c49c11e947066b7cd65c384a63e34715aaca4f0ed6eb2b876b8b3f74f9"),
							Index:           69,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x86935F11C86623deC8a25696E1C19a8659CbF95d"),
							Topics: []common.Hash{
								common.HexToHash("0x8abb27e53fc4ab514e7e5521910bdf2d1e48a69f150dbb671efa7ccc8f44bb28"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000419f0"),
								common.HexToHash("0x000000000000000000000000a2faa3405a734c04ae713aaa837e6cecc2caee9f"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000003"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf000000000000000000000000086935f11c86623dec8a25696e1c19a8659cbf95d00000000000000000000000000000000000000000000000000000000000004630000000000000000000000000000000000000000000000252248deb6e6940000000000000000000000000000000000000000000000000000000000006481d16d"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675427", 0)),
							TransactionHash: common.HexToHash("0x23d7f6c49c11e947066b7cd65c384a63e34715aaca4f0ed6eb2b876b8b3f74f9"),
							Index:           70,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x0000000000000000000000000000000000001010"),
							Topics: []common.Hash{
								common.HexToHash("0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000001010"),
								common.HexToHash("0x000000000000000000000000b9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0"),
								common.HexToHash("0x0000000000000000000000007c7379531b2aee82e4ca06d4175d13b9cbeafd49"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000156a1bc74f4ffb000000000000000000000000000000000000000000000000111add5d440af7d2a000000000000000000000000000000000000000000021e8c326aaa96317b2e3500000000000000000000000000000000000000000000000110573417cbba7d7a000000000000000000000000000000000000000000021e8c33c14c52a6702de5"),
							BlockNumber:     lo.Must(new(big.Int).SetString("43675427", 0)),
							TransactionHash: common.HexToHash("0x23d7f6c49c11e947066b7cd65c384a63e34715aaca4f0ed6eb2b876b8b3f74f9"),
							Index:           71,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x23d7f6c49c11e947066b7cd65c384a63e34715aaca4f0ed6eb2b876b8b3f74f9"),
						TransactionIndex: 14,
					},
				},
			},
			want: &schema.Feed{
				ID:      common.HexToHash("0x23d7f6c49c11e947066b7cd65c384a63e34715aaca4f0ed6eb2b876b8b3f74f9").String(),
				Network: filter.NetworkPolygon,
				Index:   14,
				From:    common.HexToAddress("0xb9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0").String(),
				To:      common.HexToAddress("0x86935f11c86623dec8a25696e1c19a8659cbf95d").String(),
				Type:    filter.TypeMetaverseTrade,
				Calldata: &schema.Calldata{
					FunctionHash: "0x66609c88",
				},
				Platform: lo.ToPtr(filter.PlatformAavegotchi),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("151529543089805376")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeMetaverseTrade,
						Platform: filter.PlatformAavegotchi.String(),
						From:     common.HexToAddress("0xa2faa3405a734c04ae713aaa837e6cecc2caee9f").String(),
						To:       common.HexToAddress("0xb9ad10b590bcd2b0db23d0005b2db0d53d9a1cf0").String(),
						Metadata: metadata.MetaverseTrade{
							Action: metadata.ActionMetaverseTradeBuy,
							Token: metadata.Token{
								Name:     "Aavegotchi",
								Symbol:   "GOTCHI",
								Standard: 721,
								URI:      "https://app.aavegotchi.com/metadata/aavegotchis/1123",
								Address:  lo.ToPtr(common.HexToAddress("0x86935F11C86623deC8a25696E1C19a8659CbF95d").String()),
								ID:       lo.ToPtr(decimal.NewFromInt(1123)),
							},
							Cost: metadata.Token{
								Name:     "Aavegotchi GHST Token (PoS)",
								Symbol:   "GHST",
								Decimals: 18,
								Standard: 20,
								Address:  lo.ToPtr(common.HexToAddress("0x385Eeac5cB85A38A9a07A70c73e0a3271CfB54A7").String()),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("685000000000000000000"))),
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1686229357,
			},
			wantError: require.NoError,
		},
	}

	config := &config.Module{
		Network:  filter.NetworkPolygon,
		Endpoint: endpoint.MustGet(filter.NetworkPolygon),
		Worker:   filter.Aavegotchi,
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			instance, err := aavegotchi.NewWorker(config)
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
