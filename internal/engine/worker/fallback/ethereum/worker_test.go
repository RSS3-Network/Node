package ethereum_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/fallback/ethereum"
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
		{
			name: "Transfer native token on Ethereum",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xea9d0ecd7a085aa998789e8e9c017a7d45f199873380ecb568218525171165b0"),
						ParentHash:   common.HexToHash("0x5eaec1d0cb27184353b58d38ee2d1c1fdabdde060b781af03e68fc4fb2e5af12"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8"),
						Number:       lo.Must(new(big.Int).SetString("14422928", 0)),
						GasLimit:     29999972,
						GasUsed:      29944698,
						Timestamp:    1647774927,
						BaseFee:      lo.Must(new(big.Int).SetString("15564031841", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x0c2f413efbc243f3bb8edac7e70bdc21936e01401a21b0d63e97732aa80f5d99"),
						From:      common.HexToAddress("0x000000A52a03835517E9d193B3c27626e1Bc96b1"),
						Gas:       21000,
						Hash:      common.HexToHash("0x0c2f413efbc243f3bb8edac7e70bdc21936e01401a21b0d63e97732aa80f5d99"),
						Input:     nil,
						To:        lo.ToPtr(common.HexToAddress("0xA1b2DCAC834117F38FB0356b5176B5693E165c90")),
						Value:     lo.Must(new(big.Int).SetString("3739360016119348", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xea9d0ecd7a085aa998789e8e9c017a7d45f199873380ecb568218525171165b0"),
						BlockNumber:       lo.Must(new(big.Int).SetString("14422928", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 25895390,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3b8c8f46b"),
						GasUsed:           21000,
						Logs:              []*ethereum.Log{},
						Status:            1,
						TransactionHash:   common.HexToHash("0x0c2f413efbc243f3bb8edac7e70bdc21936e01401a21b0d63e97732aa80f5d99"),
						TransactionIndex:  244,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:       "0x0c2f413efbc243f3bb8edac7e70bdc21936e01401a21b0d63e97732aa80f5d99",
				Network:  filter.NetworkEthereum,
				Index:    244,
				From:     "0x000000A52a03835517E9d193B3c27626e1Bc96b1",
				To:       "0xA1b2DCAC834117F38FB0356b5176B5693E165c90",
				Type:     filter.TypeTransactionTransfer,
				Calldata: &schema.Calldata{},
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("335686667463000")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type: filter.TypeTransactionTransfer,
						From: "0x000000A52a03835517E9d193B3c27626e1Bc96b1",
						To:   "0xA1b2DCAC834117F38FB0356b5176B5693E165c90",
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("3739360016119348"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
				},
				Status:    true,
				Timestamp: 1647774927,
			},
			wantError: require.NoError,
		},
		{
			name: "Burn native token on Ethereum",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xd03bf8222ec4fca9dcaf36af91eeb7d97e3693d9560acc90379bee31390cbf83"),
						ParentHash:   common.HexToHash("0xbf025fc6b39082d77f90a4a241c98d8aa3f66b06bbb0ed47a0d4a26a14eb2cb0"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xeBec795c9c8bBD61FFc14A6662944748F299cAcf"),
						Number:       lo.Must(new(big.Int).SetString("17337080", 0)),
						GasLimit:     30000000,
						GasUsed:      26304079,
						Timestamp:    1685028899,
						BaseFee:      lo.Must(new(big.Int).SetString("54021684012", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x228cab09172d5c01cdb0ad311ff384bfbbc9dc5cd149f9057bc9bd18ae7b491c"),
						From:      common.HexToAddress("0xc5750031497c6ae819578016Da83E035964418B8"),
						Gas:       500000,
						Hash:      common.HexToHash("0x228cab09172d5c01cdb0ad311ff384bfbbc9dc5cd149f9057bc9bd18ae7b491c"),
						Input:     nil,
						To:        lo.ToPtr(common.HexToAddress("0x0000000000000000000000000000000000000000")),
						Value:     lo.Must(new(big.Int).SetString("1000000000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xd03bf8222ec4fca9dcaf36af91eeb7d97e3693d9560acc90379bee31390cbf83"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17337080", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 24054688,
						EffectiveGasPrice: hexutil.MustDecodeBig("0xccf8c452c"),
						GasUsed:           21000,
						Logs:              []*ethereum.Log{},
						Status:            1,
						TransactionHash:   common.HexToHash("0x228cab09172d5c01cdb0ad311ff384bfbbc9dc5cd149f9057bc9bd18ae7b491c"),
						TransactionIndex:  293,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:       "0x228cab09172d5c01cdb0ad311ff384bfbbc9dc5cd149f9057bc9bd18ae7b491c",
				Network:  filter.NetworkEthereum,
				Index:    293,
				From:     "0xc5750031497c6ae819578016Da83E035964418B8",
				To:       "0x0000000000000000000000000000000000000000",
				Type:     filter.TypeTransactionBurn,
				Calldata: &schema.Calldata{},
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("1155455364252000")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type: filter.TypeTransactionBurn,
						From: "0xc5750031497c6ae819578016Da83E035964418B8",
						To:   "0x0000000000000000000000000000000000000000",
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1000000000000000000"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
				},
				Status:    true,
				Timestamp: 1685028899,
			},
			wantError: require.NoError,
		},
		{
			name: "Transfer ERC-20 tokens on Ethereum",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x136e5166e4a3bf7545a6681d711cb661f3e5091d6bcdd30b4462c6bbb78ecd9f"),
						ParentHash:   common.HexToHash("0x11c607312ba6d4d6698f4c1dc0ded6c518c390ef2f3d53b72411193379c94333"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x690B9A9E9aa1C9dB991C7721a92d351Db4FaC990"),
						Number:       lo.Must(new(big.Int).SetString("17981387", 0)),
						GasLimit:     29970705,
						GasUsed:      20043643,
						Timestamp:    1692841715,
						BaseFee:      lo.Must(new(big.Int).SetString("13480409622", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x4de22d893cd30c987ec54f9d550d9e5fa7e1ce76767a82bc732d49a81cf4e408"),
						From:      common.HexToAddress("0xE93381fB4c4F14bDa253907b18faD305D799241a"),
						Gas:       90000,
						Hash:      common.HexToHash("0x4de22d893cd30c987ec54f9d550d9e5fa7e1ce76767a82bc732d49a81cf4e408"),
						Input:     hexutil.MustDecode("0xa9059cbb0000000000000000000000001b961e6604a6c830a8d907447f6fc7084dc61328000000000000000000000000000000000000000000001fc3842bd1f071c00000"),
						To:        lo.ToPtr(common.HexToAddress("0xc98D64DA73a6616c42117b582e832812e7B8D57F")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x136e5166e4a3bf7545a6681d711cb661f3e5091d6bcdd30b4462c6bbb78ecd9f"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17981387", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 4743550,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3a6a01016"),
						GasUsed:           51753,
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0xc98D64DA73a6616c42117b582e832812e7B8D57F"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x000000000000000000000000e93381fb4c4f14bda253907b18fad305d799241a"),
									common.HexToHash("0x0000000000000000000000001b961e6604a6c830a8d907447f6fc7084dc61328"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000001fc3842bd1f071c00000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("17981387", 0)),
								TransactionHash: common.HexToHash("0x4de22d893cd30c987ec54f9d550d9e5fa7e1ce76767a82bc732d49a81cf4e408"),
								Index:           183,
								Removed:         false,
							},
						},
						Status:           1,
						TransactionHash:  common.HexToHash("0x4de22d893cd30c987ec54f9d550d9e5fa7e1ce76767a82bc732d49a81cf4e408"),
						TransactionIndex: 22,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x4de22d893cd30c987ec54f9d550d9e5fa7e1ce76767a82bc732d49a81cf4e408",
				Network: filter.NetworkEthereum,
				Index:   22,
				From:    "0xE93381fB4c4F14bDa253907b18faD305D799241a",
				To:      "0xc98D64DA73a6616c42117b582e832812e7B8D57F",
				Type:    filter.TypeTransactionTransfer,
				Calldata: &schema.Calldata{
					FunctionHash:   "0xa9059cbb",
					ParsedFunction: "transfer",
				},
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("811508239167366")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type: filter.TypeTransactionTransfer,
						From: "0xE93381fB4c4F14bDa253907b18faD305D799241a",
						To:   "0x1b961e6604A6c830a8D907447F6fc7084dc61328",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0xc98D64DA73a6616c42117b582e832812e7B8D57F"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("150000000000000000000000"))),
							Name:     "RSS3",
							Symbol:   "RSS3",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
				},
				Status:    true,
				Timestamp: 1692841715,
			},
			wantError: require.NoError,
		},
		{
			name: "Transfer ERC-721 tokens on Ethereum",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xab6414e152d967ec1d7c32799367b7d9c550aac5984707dc4ee891d90fc3f096"),
						ParentHash:   common.HexToHash("0xeaed29c23149acaf07790ac5d4ca007d653ed7b60511e6635dbd798bdb753cd6"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8"),
						Number:       lo.Must(new(big.Int).SetString("14555313", 0)),
						GasLimit:     30107510,
						GasUsed:      30102075,
						Timestamp:    1649557790,
						BaseFee:      lo.Must(new(big.Int).SetString("21156453465", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x164b78e31e14e08c2e445050592ae58f1d72b26d541761c2f12f7c511ef89c8a"),
						From:      common.HexToAddress("0x123456D72E335b4415b7A5665C254578575a8A0e"),
						Gas:       69517,
						Hash:      common.HexToHash("0x164b78e31e14e08c2e445050592ae58f1d72b26d541761c2f12f7c511ef89c8a"),
						Input:     hexutil.MustDecode("0x23b872dd000000000000000000000000123456d72e335b4415b7a5665c254578575a8a0e000000000000000000000000000000a52a03835517e9d193b3c27626e1bc96b10000000000000000000000000000000000000000000000000000000000000408"),
						To:        lo.ToPtr(common.HexToAddress("0x5452C7fB99D99fAb3Cc1875E9DA9829Cb50F7A13")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xab6414e152d967ec1d7c32799367b7d9c550aac5984707dc4ee891d90fc3f096"),
						BlockNumber:       lo.Must(new(big.Int).SetString("14555313", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 28195670,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x5466e0b59"),
						GasUsed:           41545,
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0x5452C7fB99D99fAb3Cc1875E9DA9829Cb50F7A13"),
								Topics: []common.Hash{
									common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
									common.HexToHash("0x000000000000000000000000123456d72e335b4415b7a5665c254578575a8a0e"),
									common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
									common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000408"),
								},
								Data:            nil,
								BlockNumber:     lo.Must(new(big.Int).SetString("14555313", 0)),
								TransactionHash: common.HexToHash("0x164b78e31e14e08c2e445050592ae58f1d72b26d541761c2f12f7c511ef89c8a"),
								Index:           533,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x5452C7fB99D99fAb3Cc1875E9DA9829Cb50F7A13"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x000000000000000000000000123456d72e335b4415b7a5665c254578575a8a0e"),
									common.HexToHash("0x000000000000000000000000000000a52a03835517e9d193b3c27626e1bc96b1"),
									common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000408"),
								},
								Data:            nil,
								BlockNumber:     lo.Must(new(big.Int).SetString("14555313", 0)),
								TransactionHash: common.HexToHash("0x164b78e31e14e08c2e445050592ae58f1d72b26d541761c2f12f7c511ef89c8a"),
								Index:           534,
								Removed:         false,
							},
						},
						Status:           1,
						TransactionHash:  common.HexToHash("0x164b78e31e14e08c2e445050592ae58f1d72b26d541761c2f12f7c511ef89c8a"),
						TransactionIndex: 271,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x164b78e31e14e08c2e445050592ae58f1d72b26d541761c2f12f7c511ef89c8a",
				Network: filter.NetworkEthereum,
				Index:   271,
				From:    "0x123456D72E335b4415b7A5665C254578575a8A0e",
				To:      "0x5452C7fB99D99fAb3Cc1875E9DA9829Cb50F7A13",
				Type:    filter.TypeCollectibleTransfer,
				Calldata: &schema.Calldata{
					FunctionHash:   "0x23b872dd",
					ParsedFunction: "transferFrom",
				},
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("941262359203425")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type: filter.TypeCollectibleApproval,
						From: "0x123456D72E335b4415b7A5665C254578575a8A0e",
						To:   ethereum.AddressGenesis.String(),
						Metadata: metadata.CollectibleApproval{
							Action: metadata.ActionCollectibleApprovalRevoke,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x5452C7fB99D99fAb3Cc1875E9DA9829Cb50F7A13"),
								ID:       lo.ToPtr(decimal.NewFromInt(1032)),
								Name:     "The Genesis RSS3 Avatar NFT",
								Symbol:   "The Genesis RSS3 Avatar NFT",
								Standard: metadata.StandardERC721,
								URI:      "data:application/json;base64,eyJuYW1lIjogIlRoZSBHZW5lc2lzIFJTUzMgQXZhdGFyIE5GVCAjMTAzMiIsICJkZXNjcmlwdGlvbiI6ICJUaGUgR2VuZXNpcyBSU1MzIEF2YXRhciBORlQgaXMgYSBjb2xsZWN0aW9uIG9mIDEwLDAwMCB1bmlxdWUgYXZhdGFycyBtZXRpY3Vsb3VzbHkgZGVzaWduZWQgdG8gaWRlbnRpZnkgUlNTMyBjb21tdW5pdHkgbWVtYmVycy4iLCAiaW1hZ2UiOiAiaXBmczovL1FtU1g5UWl3alRHQms1bTIyVXNjVGczdnJiTXdVZkZzbXhWek1INTdoa1BENVUvMTAzMi5wbmcifQ==",
							},
						},
					},
					{
						Type: filter.TypeCollectibleTransfer,
						From: "0x123456D72E335b4415b7A5665C254578575a8A0e",
						To:   "0x000000A52a03835517E9d193B3c27626e1Bc96b1",
						Metadata: metadata.CollectibleTransfer{
							Address:  lo.ToPtr("0x5452C7fB99D99fAb3Cc1875E9DA9829Cb50F7A13"),
							ID:       lo.ToPtr(decimal.NewFromInt(1032)),
							Value:    lo.ToPtr(decimal.New(1, 0)),
							Name:     "The Genesis RSS3 Avatar NFT",
							Symbol:   "The Genesis RSS3 Avatar NFT",
							Standard: metadata.StandardERC721,
							URI:      "data:application/json;base64,eyJuYW1lIjogIlRoZSBHZW5lc2lzIFJTUzMgQXZhdGFyIE5GVCAjMTAzMiIsICJkZXNjcmlwdGlvbiI6ICJUaGUgR2VuZXNpcyBSU1MzIEF2YXRhciBORlQgaXMgYSBjb2xsZWN0aW9uIG9mIDEwLDAwMCB1bmlxdWUgYXZhdGFycyBtZXRpY3Vsb3VzbHkgZGVzaWduZWQgdG8gaWRlbnRpZnkgUlNTMyBjb21tdW5pdHkgbWVtYmVycy4iLCAiaW1hZ2UiOiAiaXBmczovL1FtU1g5UWl3alRHQms1bTIyVXNjVGczdnJiTXdVZkZzbXhWek1INTdoa1BENVUvMTAzMi5wbmcifQ==",
						},
					},
				},
				Status:    true,
				Timestamp: 1649557790,
			},
			wantError: require.NoError,
		},
		// {
		// 	name: "Transfer ERC-1155 tokens on Ethereum",
		// 	arguments: arguments{
		// 		task: &source.Task{
		// 			Network: filter.NetworkEthereum,
		// 			ChainID: 1,
		// 			Header: &ethereum.Header{
		// 				Hash:         common.HexToHash("0xee79fe890d57acb46271930ea27018fb2701be815d8f2f4c10564b8df520dd76"),
		// 				ParentHash:   common.HexToHash("0xb8e5494fb5bd0f42fa3c7b4027f28b0878877ee14d9673a65dc4028ea818b3f2"),
		// 				UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
		// 				Coinbase:     common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"),
		// 				Number:       lo.Must(new(big.Int).SetString("16814688", 0)),
		// 				GasLimit:     30000000,
		// 				GasUsed:      12119148,
		// 				Timestamp:    1678657307,
		// 				BaseFee:      lo.Must(new(big.Int).SetString("19583648036", 0)),
		// 				Transactions: nil,
		// 			},
		// 			Transaction: &ethereum.Transaction{
		// 				BlockHash: common.HexToHash("0x344c70bb46d4ce1cc56a7aacd67092a17e854bf141818c8f0c66df60c61aea92"),
		// 				From:      common.HexToAddress("0x6C9AE8e43AF92BB4CA6E040A610c6eF42C0245C6"),
		// 				Gas:       77990,
		// 				Hash:      common.HexToHash("0x344c70bb46d4ce1cc56a7aacd67092a17e854bf141818c8f0c66df60c61aea92"),
		// 				Input:     hexutil.MustDecode("0xf242432a0000000000000000000000006c9ae8e43af92bb4ca6e040a610c6ef42c0245c6000000000000000000000000be39b72d425bae21e547dbe5718b14692011a9080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000"),
		// 				To:        lo.ToPtr(common.HexToAddress("0x8442864d6AB62a9193be2F16580c08E0D7BCda2f")),
		// 				Value:     lo.Must(new(big.Int).SetString("0", 0)),
		// 				Type:      2,
		// 				ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
		// 			},
		// 			Receipt: &ethereum.Receipt{
		// 				BlockHash:         common.HexToHash("0xee79fe890d57acb46271930ea27018fb2701be815d8f2f4c10564b8df520dd76"),
		// 				BlockNumber:       lo.Must(new(big.Int).SetString("16814688", 0)),
		// 				ContractAddress:   nil,
		// 				CumulativeGasUsed: 12013037,
		// 				EffectiveGasPrice: hexutil.MustDecodeBig("0x49172e32a"),
		// 				GasUsed:           51920,
		// 				Logs: []*ethereum.Log{
		// 					{
		// 						Address: common.HexToAddress("0x8442864d6AB62a9193be2F16580c08E0D7BCda2f"),
		// 						Topics: []common.Hash{
		// 							common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"),
		// 							common.HexToHash("0x0000000000000000000000006c9ae8e43af92bb4ca6e040a610c6ef42c0245c6"),
		// 							common.HexToHash("0x0000000000000000000000006c9ae8e43af92bb4ca6e040a610c6ef42c0245c6"),
		// 							common.HexToHash("0x000000000000000000000000be39b72d425bae21e547dbe5718b14692011a908"),
		// 						},
		// 						Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001"),
		// 						BlockNumber:     lo.Must(new(big.Int).SetString("16814688", 0)),
		// 						TransactionHash: common.HexToHash("0x344c70bb46d4ce1cc56a7aacd67092a17e854bf141818c8f0c66df60c61aea92"),
		// 						Index:           304,
		// 						Removed:         false,
		// 					},
		// 				},
		// 				Status:           1,
		// 				TransactionHash:  common.HexToHash("0x344c70bb46d4ce1cc56a7aacd67092a17e854bf141818c8f0c66df60c61aea92"),
		// 				TransactionIndex: 135,
		// 			},
		// 		},
		// 		config: &config.Module{
		// 			Network:  filter.NetworkEthereum,
		// 			Endpoint: endpoint.MustGet(filter.NetworkEthereum),
		// 		},
		// 	},
		// 	want: &schema.Feed{
		// 		ID:      "0x344c70bb46d4ce1cc56a7aacd67092a17e854bf141818c8f0c66df60c61aea92",
		// 		Network: filter.NetworkEthereum,
		// 		Index:   135,
		// 		From:    "0x6C9AE8e43AF92BB4CA6E040A610c6eF42C0245C6",
		// 		To:      "0x8442864d6AB62a9193be2F16580c08E0D7BCda2f",
		// 		Type:    filter.TypeCollectibleTransfer,
		// 		Fee: &schema.Fee{
		// 			Amount:  lo.Must(decimal.NewFromString("1018675320043040")),
		// 			Decimal: 18,
		// 		},
		// 		Actions: []*schema.Action{
		// 			{
		// 				Type: filter.TypeCollectibleTransfer,
		// 				From: "0x6C9AE8e43AF92BB4CA6E040A610c6eF42C0245C6",
		// 				To:   "0xBE39b72D425BAe21e547dBE5718b14692011a908",
		// 				Metadata: metadata.CollectibleTransfer{
		// 					Address:  lo.ToPtr("0x8442864d6AB62a9193be2F16580c08E0D7BCda2f"),
		// 					ID:       lo.ToPtr(lo.Must(decimal.NewFromString("0"))),
		// 					Value:    lo.ToPtr(decimal.NewFromInt(1)),
		// 					Name:     "TIME NFT Special Issues",
		// 					Standard: metadata.StandardERC1155,
		// 					URI:      "https://timepieces.mypinata.cloud/ipfs/QmbZtT3KVHYL6QAuNbrLGKp4sQyaWkkzTDRvHFCmq7Py6y",
		// 				},
		// 			},
		// 		},
		// 		Status:    true,
		// 		Timestamp: 1678657307,
		// 	},
		// 	wantError: require.NoError,
		// },
		{
			name: "Transfer ETH on Base",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 8453,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x24c8559fb47d3a9320c868b30c257255d6ca3430ddb18264f46554578b9665e4"),
						ParentHash:   common.HexToHash("0xa988a9126cfc8739bfac483e62bff291af7103c6225b9f5a4e9e0f245ae42fb5"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4200000000000000000000000000000000000011"),
						Number:       lo.Must(new(big.Int).SetString("2430365", 0)),
						GasLimit:     30000000,
						GasUsed:      2745325,
						Timestamp:    1691650077,
						BaseFee:      lo.Must(new(big.Int).SetString("55", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xc5bb7e79738b494b9742ccc78451e4884304ad22471dbbaaf6fc70d5572ae284"),
						From:      common.HexToAddress("0x1d97d5c7d68E03BAe6FBb1ec6E5887d6eAaaAA7d"),
						Gas:       21000,
						GasPrice:  lo.Must(new(big.Int).SetString("1500000055", 10)),
						Hash:      common.HexToHash("0xc5bb7e79738b494b9742ccc78451e4884304ad22471dbbaaf6fc70d5572ae284"),
						Input:     nil,
						To:        lo.ToPtr(common.HexToAddress("0xdd9176eA3E7559D6B68b537eF555D3e89403f742")),
						Value:     lo.Must(new(big.Int).SetString("550000000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("8453", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x24c8559fb47d3a9320c868b30c257255d6ca3430ddb18264f46554578b9665e4"),
						BlockNumber:       lo.Must(new(big.Int).SetString("2430365", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 513567,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x59682f37"),
						GasUsed:           21000,
						L1GasPrice:        lo.Must(new(big.Int).SetString("13361473940", 0)),
						L1GasUsed:         lo.Must(new(big.Int).SetString("2056", 0)),
						L1Fee:             lo.Must(new(big.Int).SetString("18790294247717", 0)),
						FeeScalar:         lo.Must(new(big.Float).SetString("0.684")),
						Logs:              []*ethereum.Log{},
						Status:            1,
						TransactionHash:   common.HexToHash("0xc5bb7e79738b494b9742ccc78451e4884304ad22471dbbaaf6fc70d5572ae284"),
						TransactionIndex:  8,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:       "0xc5bb7e79738b494b9742ccc78451e4884304ad22471dbbaaf6fc70d5572ae284",
				Network:  filter.NetworkEthereum,
				Index:    8,
				From:     "0x1d97d5c7d68E03BAe6FBb1ec6E5887d6eAaaAA7d",
				To:       "0xdd9176eA3E7559D6B68b537eF555D3e89403f742",
				Type:     filter.TypeTransactionTransfer,
				Calldata: &schema.Calldata{},
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("50290295402717")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type: filter.TypeTransactionTransfer,
						From: "0x1d97d5c7d68E03BAe6FBb1ec6E5887d6eAaaAA7d",
						To:   "0xdd9176eA3E7559D6B68b537eF555D3e89403f742",
						Metadata: metadata.TransactionTransfer{
							Name:     "Ethereum",
							Symbol:   "ETH",
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("550000000000000000"))),
							Decimals: 18,
						},
					},
				},
				Status:    true,
				Timestamp: 1691650077,
			},
			wantError: require.NoError,
		},
		{
			name: "Ethernal Inscriptions",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x279b9741d69fc624dfaee27b21c014732732751ee4c3283e6e605da40a136af1"),
						ParentHash:   common.HexToHash("0xd1c532ebe5a54be1205201f43648af9c65ec90993caba97cb9081eb7c71e18c7"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"),
						Number:       lo.Must(new(big.Int).SetString("17577091", 0)),
						GasLimit:     30000000,
						GasUsed:      29034573,
						Timestamp:    1687946543,
						BaseFee:      lo.Must(new(big.Int).SetString("40411201078", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x6a1780cad0c68e9500bd9375b3b8b3b5f84ad82a042f2f19aa55a63f4d60e9bf"),
						From:      common.HexToAddress("0x9FA545098dc66E4A2A0fBfE0388830ef51F70150"),
						Gas:       91960,
						GasPrice:  lo.Must(new(big.Int).SetString("40630040231", 10)),
						Hash:      common.HexToHash("0x6a1780cad0c68e9500bd9375b3b8b3b5f84ad82a042f2f19aa55a63f4d60e9bf"),
						Input:     hexutil.MustDecode("0x657468010d696d6167652f7376672b786d6c000001d83c7376672077696474683d223130302522206865696768743d2231303025222076657273696f6e3d22312e31222076696577426f783d223020302032342032342220786d6c6e733d22687474703a2f2f7777772e77332e6f72672f323030302f737667223e0a3c696d61676520687265663d222f636f6e74656e742f3078313562363930303736313839636566383633316361336662663865623135336338386335363536353131616538316236336565653139313065303531346537362220696d6167652d72656e646572696e673d22706978656c6174656422202f3e0a3c696d61676520687265663d222f636f6e74656e742f3078326263386233643238323131646535333331373536306335363730333432346439366430343237666436353137356532313761393937646161623232366133362220696d6167652d72656e646572696e673d22706978656c6174656422202f3e0a3c696d61676520687265663d222f636f6e74656e742f3078313366653061623463326561616561333537343536363262626438613136343731656466333533646139343139333533643766383836336366336131323132322220696d6167652d72656e646572696e673d22706978656c6174656422202f3e0a3c2f7376673e0a"),
						To:        lo.ToPtr(common.HexToAddress("0xA3C66393049fAB4830C330Dfe658f94A4de0A122")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x279b9741d69fc624dfaee27b21c014732732751ee4c3283e6e605da40a136af1"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17577091", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 27869117,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x975bd36a7"),
						GasUsed:           61307,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xA3C66393049fAB4830C330Dfe658f94A4de0A122"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x0000000000000000000000009fa545098dc66e4a2a0fbfe0388830ef51f70150"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000081ef"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("17577091", 0)),
							TransactionHash: common.HexToHash("0x6a1780cad0c68e9500bd9375b3b8b3b5f84ad82a042f2f19aa55a63f4d60e9bf"),
							Index:           130,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x6a1780cad0c68e9500bd9375b3b8b3b5f84ad82a042f2f19aa55a63f4d60e9bf"),
						TransactionIndex: 1160,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x6a1780cad0c68e9500bd9375b3b8b3b5f84ad82a042f2f19aa55a63f4d60e9bf",
				Network: filter.NetworkEthereum,
				Index:   1160,
				From:    "0x9FA545098dc66E4A2A0fBfE0388830ef51F70150",
				To:      "0xA3C66393049fAB4830C330Dfe658f94A4de0A122",
				Type:    filter.TypeUnknown,
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("2490905876441917")),
					Decimal: 18,
				},
				Actions:   nil,
				Status:    true,
				Timestamp: 1687946543,
			},
			wantError: require.NoError,
		},
		{
			name: "Get Function Name From Calldata",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x44b99ccd1008d85a47683a2cdb428751f41441eeed3ad91ec4f621c390f58335"),
						ParentHash:   common.HexToHash("0x4c81f84dc2b2a329b274f2214e85c158234627858e25987ff1e7e0e0fd6c1425"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xDccA982701a264e8d629A6E8CFBa9C1a27912623"),
						Number:       lo.Must(new(big.Int).SetString("19311439", 0)),
						GasLimit:     29999972,
						GasUsed:      12142826,
						Timestamp:    1708946831,
						BaseFee:      lo.Must(new(big.Int).SetString("27571507876", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xded5fd30829d08dcf608539e993c5916f23c19c48ee0ebac8b748e9aec1e6ce5"),
						From:      common.HexToAddress("0x38d004aea23c1DC86F9cFF25dcda840Ae9dF8dAe"),
						Gas:       109884,
						GasPrice:  lo.Must(new(big.Int).SetString("27572393169", 10)),
						Hash:      common.HexToHash("0xded5fd30829d08dcf608539e993c5916f23c19c48ee0ebac8b748e9aec1e6ce5"),
						Input:     hexutil.MustDecode("0xf6203e3500000000000000000000000038d004aea23c1dc86f9cff25dcda840ae9df8dae"),
						To:        lo.ToPtr(common.HexToAddress("0xF047ab4c75cebf0eB9ed34Ae2c186f3611aEAfa6")),
						Value:     lo.Must(new(big.Int).SetString("50000000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x44b99ccd1008d85a47683a2cdb428751f41441eeed3ad91ec4f621c390f58335"),
						BlockNumber:       lo.Must(new(big.Int).SetString("19311439", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 10143994,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x66b7154d1"),
						GasUsed:           72386,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xF047ab4c75cebf0eB9ed34Ae2c186f3611aEAfa6"),
							Topics: []common.Hash{
								common.HexToHash("0x2c0f148b435140de488c1b34647f1511c646f7077e87007bacf22ef9977a16d8"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000001ebd"),
								common.HexToHash("0x00000000000000000000000038d004aea23c1dc86f9cff25dcda840ae9df8dae"),
								common.HexToHash("0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000b1a2bc2ec50000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19311439", 0)),
							TransactionHash: common.HexToHash("0xded5fd30829d08dcf608539e993c5916f23c19c48ee0ebac8b748e9aec1e6ce5"),
							Index:           227,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"),
								common.HexToHash("0x000000000000000000000000f047ab4c75cebf0eb9ed34ae2c186f3611aeafa6"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000b1a2bc2ec50000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19311439", 0)),
							TransactionHash: common.HexToHash("0xded5fd30829d08dcf608539e993c5916f23c19c48ee0ebac8b748e9aec1e6ce5"),
							Index:           228,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xded5fd30829d08dcf608539e993c5916f23c19c48ee0ebac8b748e9aec1e6ce5"),
						TransactionIndex: 160,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0xded5fd30829d08dcf608539e993c5916f23c19c48ee0ebac8b748e9aec1e6ce5",
				Network: filter.NetworkEthereum,
				Index:   160,
				From:    "0x38d004aea23c1DC86F9cFF25dcda840Ae9dF8dAe",
				To:      "0xF047ab4c75cebf0eB9ed34Ae2c186f3611aEAfa6",
				Type:    filter.TypeUnknown,
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("1995855251931234")),
					Decimal: 18,
				},
				Calldata: &schema.Calldata{
					FunctionHash:   "0xf6203e35",
					ParsedFunction: "depositETHFor",
				},
				Actions:   []*schema.Action{},
				Status:    true,
				Timestamp: 1708946831,
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
