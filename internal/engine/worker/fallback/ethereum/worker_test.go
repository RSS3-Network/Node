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
		{
			name: "Transfer RSS3 on RSS3 Mainnet",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkVSL,
					ChainID: 12553,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xa77316fd6775123ef7eb4edaa8dd9604227b62c6ea2970fd813ebd38deff3ed7"),
						ParentHash:   common.HexToHash("0x97f36887c0a528f944b93d9c929a686d2fb603c6eec5e7add9d96053faf1c4f5"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4200000000000000000000000000000000000011"),
						Number:       lo.Must(new(big.Int).SetString("182448", 0)),
						GasLimit:     30000000,
						GasUsed:      67925,
						Timestamp:    1710223415,
						BaseFee:      lo.Must(new(big.Int).SetString("50", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xa2e19137d6164d8a73e031c6ad34dd44f4c6244996e5250485e4c4dbb08e6d06"),
						From:      common.HexToAddress("0x27005E931Ea9ff482a3e5deb4c6eFa980305d74a"),
						Gas:       21000,
						GasPrice:  lo.Must(new(big.Int).SetString("1500000050", 10)),
						Hash:      common.HexToHash("0xa2e19137d6164d8a73e031c6ad34dd44f4c6244996e5250485e4c4dbb08e6d06"),
						Input:     nil,
						To:        lo.ToPtr(common.HexToAddress("0xDBEe696fA4398649DFD6a6B9F474B7Dfd4CAA7e0")),
						Value:     lo.Must(new(big.Int).SetString("10000000000000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("12553", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xa77316fd6775123ef7eb4edaa8dd9604227b62c6ea2970fd813ebd38deff3ed7"),
						BlockNumber:       lo.Must(new(big.Int).SetString("182448", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 67925,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x59682f32"),
						GasUsed:           21000,
						L1GasPrice:        lo.Must(new(big.Int).SetString("52215805297", 0)),
						L1GasUsed:         lo.Must(new(big.Int).SetString("2088", 0)),
						L1Fee:             lo.Must(new(big.Int).SetString("654159608760816000", 0)),
						FeeScalar:         lo.Must(new(big.Float).SetString("6000")),
						Logs:              []*ethereum.Log{},
						Status:            1,
						TransactionHash:   common.HexToHash("0xa2e19137d6164d8a73e031c6ad34dd44f4c6244996e5250485e4c4dbb08e6d06"),
						TransactionIndex:  1,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkVSL,
					Endpoint: endpoint.MustGet(filter.NetworkVSL),
				},
			},
			want: &schema.Feed{
				ID:       "0xa2e19137d6164d8a73e031c6ad34dd44f4c6244996e5250485e4c4dbb08e6d06",
				Network:  filter.NetworkVSL,
				Index:    1,
				From:     "0x27005E931Ea9ff482a3e5deb4c6eFa980305d74a",
				To:       "0xDBEe696fA4398649DFD6a6B9F474B7Dfd4CAA7e0",
				Type:     filter.TypeTransactionTransfer,
				Calldata: &schema.Calldata{},
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("654191108761866000")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type: filter.TypeTransactionTransfer,
						From: "0x27005E931Ea9ff482a3e5deb4c6eFa980305d74a",
						To:   "0xDBEe696fA4398649DFD6a6B9F474B7Dfd4CAA7e0",
						Metadata: metadata.TransactionTransfer{
							Name:     "RSS3",
							Symbol:   "RSS3",
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("10000000000000000000000"))),
							Decimals: 18,
						},
					},
				},
				Status:    true,
				Timestamp: 1710223415,
			},
			wantError: require.NoError,
		},
		{
			name: "Transfer WETH tokens on RSS3 Mainnet",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkVSL,
					ChainID: 12553,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x94150fd1dff8f4353e66e81f48a3e03c2f341d03bd1c3297f30e96b27ed73954"),
						ParentHash:   common.HexToHash("0x448461a64af441570cc09aed6b1e4484d239be133f6cc5c7d6c34b2eb7c6035f"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4200000000000000000000000000000000000011"),
						Number:       lo.Must(new(big.Int).SetString("221260", 0)),
						GasLimit:     30000000,
						GasUsed:      115533,
						Timestamp:    1710301039,
						BaseFee:      lo.Must(new(big.Int).SetString("50", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xdfd0b84e7a7e076f97f69e5027117edc7bc34f849165d9b2adc95f7cae44455b"),
						From:      common.HexToAddress("0xF209b7Bbadf8d9518a822aEaa7119B38b17377A7"),
						Gas:       77851,
						GasPrice:  lo.Must(new(big.Int).SetString("100050", 10)),
						Hash:      common.HexToHash("0xdfd0b84e7a7e076f97f69e5027117edc7bc34f849165d9b2adc95f7cae44455b"),
						Input:     hexutil.MustDecode("0xa9059cbb00000000000000000000000001d1843cbcfff37b298a09a03939c90dab191622000000000000000000000000000000000000000000000000016345785d8a0000"),
						To:        lo.ToPtr(common.HexToAddress("0xB659A97D8ae3c43E91Eafe5dba110a7e799157c4")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("12553", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x94150fd1dff8f4353e66e81f48a3e03c2f341d03bd1c3297f30e96b27ed73954"),
						BlockNumber:       lo.Must(new(big.Int).SetString("221260", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 115533,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x186d2"),
						GasUsed:           51508,
						L1GasPrice:        lo.Must(new(big.Int).SetString("48810209622", 0)),
						L1GasUsed:         lo.Must(new(big.Int).SetString("2596", 0)),
						L1Fee:             lo.Must(new(big.Int).SetString("1013690433429696000", 0)),
						FeeScalar:         lo.Must(new(big.Float).SetString("8000")),
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xB659A97D8ae3c43E91Eafe5dba110a7e799157c4"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000f209b7bbadf8d9518a822aeaa7119b38b17377a7"),
								common.HexToHash("0x00000000000000000000000001d1843cbcfff37b298a09a03939c90dab191622"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000016345785d8a0000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("221260", 0)),
							TransactionHash: common.HexToHash("0xdfd0b84e7a7e076f97f69e5027117edc7bc34f849165d9b2adc95f7cae44455b"),
							Index:           0,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xdfd0b84e7a7e076f97f69e5027117edc7bc34f849165d9b2adc95f7cae44455b"),
						TransactionIndex: 1,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkVSL,
					Endpoint: endpoint.MustGet(filter.NetworkVSL),
				},
			},
			want: &schema.Feed{
				ID:      "0xdfd0b84e7a7e076f97f69e5027117edc7bc34f849165d9b2adc95f7cae44455b",
				Network: filter.NetworkVSL,
				Index:   1,
				From:    "0xF209b7Bbadf8d9518a822aEaa7119B38b17377A7",
				To:      "0xB659A97D8ae3c43E91Eafe5dba110a7e799157c4",
				Type:    filter.TypeTransactionTransfer,
				Calldata: &schema.Calldata{
					FunctionHash:   "0xa9059cbb",
					ParsedFunction: "transfer",
				},
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("1013690438583071400")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type: filter.TypeTransactionTransfer,
						From: "0xF209b7Bbadf8d9518a822aEaa7119B38b17377A7",
						To:   "0x01D1843CbCFfF37B298a09A03939c90dab191622",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0xB659A97D8ae3c43E91Eafe5dba110a7e799157c4"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("100000000000000000"))),
							Name:     "WETH",
							Symbol:   "WETH",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
				},
				Status:    true,
				Timestamp: 1710301039,
			},
			wantError: require.NoError,
		},
		{
			name: "Mint Open Chips tokens on Ethereum",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkVSL,
					ChainID: 12553,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xc358a8dc0da6cac588e5a98b805d60cd5d2746eced7606cb153155fccd6ea460"),
						ParentHash:   common.HexToHash("0xcebd5a7ab6f8fc7b396225f6c8a9a3b4963da19ac4165629cb73ab1012dde314"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4200000000000000000000000000000000000011"),
						Number:       lo.Must(new(big.Int).SetString("184967", 0)),
						GasLimit:     30000000,
						GasUsed:      184758,
						Timestamp:    1710228453,
						BaseFee:      lo.Must(new(big.Int).SetString("50", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x26fbdff8a7af1a1917b6dd8a54c636151cb7186a6b3c3a5922df3a98edda54a8"),
						From:      common.HexToAddress("0x30286DD245338292F319809935a1037CcD4573Ea"),
						Gas:       140094,
						GasPrice:  lo.Must(new(big.Int).SetString("1000050", 10)),
						Hash:      common.HexToHash("0x26fbdff8a7af1a1917b6dd8a54c636151cb7186a6b3c3a5922df3a98edda54a8"),
						Input:     hexutil.MustDecode("0x2647620400000000000000000000000039f9e912c1f696f533e7a2267ea233aec9742b35"),
						To:        lo.ToPtr(common.HexToAddress("0x28F14d917fddbA0c1f2923C406952478DfDA5578")),
						Value:     lo.Must(new(big.Int).SetString("500000000000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("12553", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xc358a8dc0da6cac588e5a98b805d60cd5d2746eced7606cb153155fccd6ea460"),
						BlockNumber:       lo.Must(new(big.Int).SetString("184967", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 184758,
						EffectiveGasPrice: hexutil.MustDecodeBig("0xf4272"),
						GasUsed:           137833,
						L1GasPrice:        lo.Must(new(big.Int).SetString("51479747982", 0)),
						L1GasUsed:         lo.Must(new(big.Int).SetString("2500", 0)),
						L1Fee:             lo.Must(new(big.Int).SetString("772196219730000000", 0)),
						FeeScalar:         lo.Must(new(big.Float).SetString("6000")),
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x849f8F55078dCc69dD857b58Cc04631EBA54E4DE"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x00000000000000000000000030286dd245338292f319809935a1037ccd4573ea"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000a2"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("184967", 0)),
							TransactionHash: common.HexToHash("0x26fbdff8a7af1a1917b6dd8a54c636151cb7186a6b3c3a5922df3a98edda54a8"),
							Index:           0,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x28F14d917fddbA0c1f2923C406952478DfDA5578"),
							Topics: []common.Hash{
								common.HexToHash("0xad3fa07f4195b47e64892eb944ecbfc253384053c119852bb2bcae484c2fcb69"),
								common.HexToHash("0x00000000000000000000000030286dd245338292f319809935a1037ccd4573ea"),
								common.HexToHash("0x00000000000000000000000039f9e912c1f696f533e7a2267ea233aec9742b35"),
								common.HexToHash("0x00000000000000000000000000000000000000000000001b1ae4d6e2ef500000"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000a200000000000000000000000000000000000000000000000000000000000000a2"),
							BlockNumber:     lo.Must(new(big.Int).SetString("184967", 0)),
							TransactionHash: common.HexToHash("0x26fbdff8a7af1a1917b6dd8a54c636151cb7186a6b3c3a5922df3a98edda54a8"),
							Index:           1,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x26fbdff8a7af1a1917b6dd8a54c636151cb7186a6b3c3a5922df3a98edda54a8"),
						TransactionIndex: 1,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkVSL,
					Endpoint: endpoint.MustGet(filter.NetworkVSL),
				},
			},
			want: &schema.Feed{
				ID:      "0x26fbdff8a7af1a1917b6dd8a54c636151cb7186a6b3c3a5922df3a98edda54a8",
				Network: filter.NetworkVSL,
				Index:   1,
				From:    "0x30286DD245338292F319809935a1037CcD4573Ea",
				To:      "0x28F14d917fddbA0c1f2923C406952478DfDA5578",
				Type:    filter.TypeCollectibleMint,
				Calldata: &schema.Calldata{
					FunctionHash:   "0x26476204",
					ParsedFunction: "stake",
				},
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("772196357569891650")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type: filter.TypeCollectibleMint,
						From: "0x0000000000000000000000000000000000000000",
						To:   "0x30286DD245338292F319809935a1037CcD4573Ea",
						Metadata: metadata.CollectibleTransfer{
							Address:  lo.ToPtr("0x849f8F55078dCc69dD857b58Cc04631EBA54E4DE"),
							ID:       lo.ToPtr(decimal.NewFromInt(162)),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
							Name:     "Open Chips",
							Symbol:   "Chips",
							Standard: metadata.StandardERC721,
							URI:      "data:application/json;base64,eyJuYW1lIjogIk9wZW4gQ2hpcHMgIzE2MiIsICJkZXNjcmlwdGlvbiI6ICJDaGlwIE1vbnN0ZXJzIGFyZSB1bmlxdWUgY3JlYXR1cmVzIGxpdmluZyBpbiB0aGUgUlNTMyBOZXR3b3JrLCBlYWNoIG9uZSBzcGVjaWFsIGJlY2F1c2Ugb2Ygd2hlcmUgaXQgd2FzIGJvcm4uIFRoZXkgcmVwcmVzZW50IHRoZSBpZGVhIG9mIEZSRUUgYW5kIE9QRU4gSU5GT1JNQVRJT04sIHRocml2aW5nIGluIGEgd29ybGQgdGhhdCB2YWx1ZXMgc2hhcmluZyBhbmQgYmVpbmcgZGlmZmVyZW50LiBUaGVzZSBDaGlwIE1vbnN0ZXJzIGFyZSBtb3JlIHRoYW4ganVzdCBkaWdpdGFsOyB0aGV5IHN5bWJvbGl6ZSB0aGUgZXhjaXRlbWVudCBhbmQgaW1wb3J0YW5jZSBvZiBiZWluZyB1bmlxdWUgaW4gYSBjb25uZWN0ZWQgZGlnaXRhbCB3b3JsZC4iLCJpbWFnZSI6ImRhdGE6aW1hZ2Uvc3ZnK3htbDtiYXNlNjQsUEQ5NGJXd2dkbVZ5YzJsdmJqMGlNUzR3SWlCbGJtTnZaR2x1WnowaWRYUm1MVGdpUHo0OGMzWm5JSFpsY25OcGIyNDlJakV1TVNJZ2FXUTlJa3hoZVdWeVh6RWlJSGh0Ykc1elBTSm9kSFJ3T2k4dmQzZDNMbmN6TG05eVp5OHlNREF3TDNOMlp5SWdlRzFzYm5NNmVHeHBibXM5SW1oMGRIQTZMeTkzZDNjdWR6TXViM0puTHpFNU9Ua3ZlR3hwYm1zaUlIZzlJakJ3ZUNJZ2VUMGlNSEI0SWlCMmFXVjNRbTk0UFNJd0lEQWdNVEF3SURFd01DSWdjM1I1YkdVOUltVnVZV0pzWlMxaVlXTnJaM0p2ZFc1a09tNWxkeUF3SURBZ01UQXdJREV3TUR0aVlXTnJaM0p2ZFc1a0xXTnZiRzl5T21Kc1lXTnJPeUlnZUcxc09uTndZV05sUFNKd2NtVnpaWEoyWlNJK1BITjBlV3hsSUhSNWNHVTlJblJsZUhRdlkzTnpJajR1Wm50bWFXeHNPaU14TkRjM1JrSTdmUzVqZTJacGJHdzZJekUwTnpkR1FqdDlMbUo3Wm1sc2JEb2pSa0l4TkRZM08zMHVhSHRtYVd4c09pTXhORGMzUmtJN2ZTNWhiSEJvWVh0bWFXeHNPaU14TkRjM1JrSTdmUzV3WjN0bWFXeHNPaU5HUWpFME5qYzdmUzVsZTJacGJHd3RjblZzWlRwbGRtVnViMlJrTzJOc2FYQXRjblZzWlRwbGRtVnViMlJrTzMwOEwzTjBlV3hsUGp4d1lYUm9JR05zWVhOelBTSm1JaUJrUFNKTk1UWWdNVEJvTmpoMk1rZ3hObnB0TkMwMGFEUjJNbWd0TkhwdE9DQXdhRFIyTW1ndE5IcHRPQ0F3YURSMk1tZ3ROSHB0T0NBd2FEUjJNbWd0TkhwdE9DQXdhRFIyTW1ndE5IcHRPQ0F3YURSMk1tZ3ROSHB0T0NBd2FEUjJNbWd0TkhwdE9DQXdhRFIyTW1ndE5IcE5NVFlnT0Rob05qaDJNa2d4Tm5wdE5DQTBhRFIyTW1ndE5IcHRPQ0F3YURSMk1tZ3ROSHB0T0NBd2FEUjJNbWd0TkhwdE9DQXdhRFIyTW1ndE5IcHRPQ0F3YURSMk1tZ3ROSHB0T0NBd2FEUjJNbWd0TkhwdE9DQXdhRFIyTW1ndE5IcHRPQ0F3YURSMk1tZ3ROSHB0TVRRdE56WjJOamhvTFRKV01UWjZiVFFnTkhZMGFDMHlkaTAwZW0wd0lEaDJOR2d0TW5ZdE5IcHRNQ0E0ZGpSb0xUSjJMVFI2YlRBZ09IWTBhQzB5ZGkwMGVtMHdJRGgyTkdndE1uWXROSHB0TUNBNGRqUm9MVEoyTFRSNmJUQWdPSFkwYUMweWRpMDBlbTB3SURoMk5HZ3RNbll0TkhwTk1USWdNVFoyTmpob0xUSldNVFo2YlMwMElEUjJORWcyZGkwMGVtMHdJRGgyTkVnMmRpMDBlbTB3SURoMk5FZzJkaTAwZW0wd0lEaDJORWcyZGkwMGVtMHdJRGgyTkVnMmRpMDBlbTB3SURoMk5FZzJkaTAwZW0wd0lEaDJORWcyZGkwMGVtMHdJRGgyTkVnMmRpMDBlaUl2UGp4emRtY2dZMnhoYzNNOUltTWlQanh3WVhSb0lHUTlJazB4T0NBNE1tZ3lWakU0YUMweWVtMHdMVFl5YUMweWRpMHlhREo2YlRBZ05HZ3RNbll0TW1neWVtMHdJRFJvTFRKMkxUSm9NbnB0TUNBMGFDMHlkaTB5YURKNmJUQWdOR2d0TW5ZdE1tZ3llbTB3SURSb0xUSjJMVEpvTW5wdE1DQTBhQzB5ZGkweWFESjZiVEFnTkdndE1uWXRNbWd5ZW0wd0lEUm9MVEoyTFRKb01ucHRNQ0EwYUMweWRpMHlhREo2YlRBZ05HZ3RNbll0TW1neWVtMHdJRFJvTFRKMkxUSm9NbnB0TUNBMGFDMHlkaTB5YURKNmJUQWdOR2d0TW5ZdE1tZ3llbTB3SURSb0xUSjJMVEpvTW5wdE1DQTBhQzB5ZGkweWFESjZiVFkySURKMkxUSklNakIyTW5wdExUWXlJREIyTW1ndE1uWXRNbnB0TkNBd2RqSm9MVEoyTFRKNmJUUWdNSFl5YUMweWRpMHllbTAwSURCMk1tZ3RNbll0TW5wdE5DQXdkakpvTFRKMkxUSjZiVFFnTUhZeWFDMHlkaTB5ZW0wMElEQjJNbWd0TW5ZdE1ucHROQ0F3ZGpKb0xUSjJMVEo2YlRRZ01IWXlhQzB5ZGkweWVtMDBJREIyTW1ndE1uWXRNbnB0TkNBd2RqSm9MVEoyTFRKNmJUUWdNSFl5YUMweWRpMHllbTAwSURCMk1tZ3RNbll0TW5wdE5DQXdkakpvTFRKMkxUSjZiVFFnTUhZeWFDMHlkaTB5ZW0wMElEQjJNbWd0TW5ZdE1ub2lMejQ4Y0dGMGFDQmtQU0pOT0RJZ01UaG9MVEoyTmpSb01ucHRNQ0EyTW1neWRqSm9MVEo2YlRBdE5HZ3lkakpvTFRKNmJUQXROR2d5ZGpKb0xUSjZiVEF0TkdneWRqSm9MVEo2YlRBdE5HZ3lkakpvTFRKNmJUQXROR2d5ZGpKb0xUSjZiVEF0TkdneWRqSm9MVEo2YlRBdE5HZ3lkakpvTFRKNmJUQXROR2d5ZGpKb0xUSjZiVEF0TkdneWRqSm9MVEo2YlRBdE5HZ3lkakpvTFRKNmJUQXROR2d5ZGpKb0xUSjZiVEF0TkdneWRqSm9MVEo2YlRBdE5HZ3lkakpvTFRKNmJUQXROR2d5ZGpKb0xUSjZiVEF0TkdneWRqSm9MVEo2SWk4K1BIQmhkR2dnWkQwaVRUZ3lJREU0ZGpKSU1UaDJMVEo2YlMwMk1pQXdkaTB5YUMweWRqSjZiVFFnTUhZdE1tZ3RNbll5ZW0wMElEQjJMVEpvTFRKMk1ucHROQ0F3ZGkweWFDMHlkako2YlRRZ01IWXRNbWd0TW5ZeWVtMDBJREIyTFRKb0xUSjJNbnB0TkNBd2RpMHlhQzB5ZGpKNmJUUWdNSFl0TW1ndE1uWXllbTAwSURCMkxUSm9MVEoyTW5wdE5DQXdkaTB5YUMweWRqSjZiVFFnTUhZdE1tZ3RNbll5ZW0wMElEQjJMVEpvTFRKMk1ucHROQ0F3ZGkweWFDMHlkako2YlRRZ01IWXRNbWd0TW5ZeWVtMDBJREIyTFRKb0xUSjJNbnB0TkNBd2RpMHlhQzB5ZGpKNmJTMDJJREpvTW5ZeWFDMHllbTB5SURKb01uWXlhQzB5ZW0weUlESm9Nbll5YUMweWVtMHROQ0ExTm1neWRpMHlhQzB5ZW0weUxUSm9Nbll0TW1ndE1ucHRNaTB5YURKMkxUSm9MVEo2VFRJMklESXdhQzB5ZGpKb01ucHRMVElnTW1ndE1uWXlhREo2YlMweUlESm9MVEoyTW1neWVtMDBJRFUyYUMweWRpMHlhREo2YlMweUxUSm9MVEoyTFRKb01ucHRMVEl0TW1ndE1uWXRNbWd5ZWlJdlBqd3ZjM1puUGp4d1lYUm9JR05zWVhOelBTSmhiSEJvWVNJZ1pEMGlUVFlnTmxZMGFEUjJNbnB0TkNBd2RqSklObFkyU0RSMk5tZ3lkaTB5YURSMk1tZ3lWalo2YlRnMExUSjJNbWd0TkZZMGVtMHdJREoyTW1ndE5GWTJhQzB5ZGpab01uWXRNbWcwZGpKb01sWTJlazB4TUNBNE9IWXlTRFoyTFRKNmJUQWdNbll5U0RaMkxUSklOSFkyYURKMkxUSm9OSFl5YURKMkxUWjZiVGcwTFRKMk1tZ3ROSFl0TW5wdE1DQXlkakpvTFRSMkxUSm9MVEoyTm1neWRpMHlhRFIyTW1neWRpMDJlaUl2UGp4d2IyeDVaMjl1SUdOc1lYTnpQU0ppSWlCd2IybHVkSE05SWpjeUxEUTRJRGN5TERZMElEY3dMRFkwSURjd0xEWTJJRFk0TERZMklEWTRMRFk0SURZMExEWTRJRFkwTERjd0lEWXdMRGN3SURZd0xEY3lJRFF3TERjeUlEUXdMRGN3SURNMkxEY3dJRE0yTERZNElETXlMRFk0SURNeUxEWTJJQ0FnTXpBc05qWWdNekFzTmpRZ01qZ3NOalFnTWpnc05EZ2dNekFzTkRnZ016QXNORFlnTXpJc05EWWdNeklzTkRRZ05qZ3NORFFnTmpnc05EWWdOekFzTkRZZ056QXNORGdnSWk4K1BIQmhkR2dnWkQwaVRUWTJJRFV3U0RVMGRqSm9NVEo2YlMweUxUSklOVFIyTW1neE1IcHRMVEl0TW1ndE5uWXlhRFo2YlRRZ05rZzFOSFl5YURFeWVtMHdJREpJTlRSMk1tZ3hNbnB0TFRJZ01tZ3RPSFl5YURoNklpQm1hV3hzUFNJak1EQXdJaTgrUEhCaGRHZ2daRDBpVFRZeElEVTBhQzAwZGpKb05Ib2lJR1pwYkd3OUlpTkVSVVUxUkRraUx6NDhjR0YwYUNCa1BTSk5NelFnTlRCb01USjJNa2d6TkhwdE1pMHlhREV3ZGpKSU16WjZiVEl0TW1nMmRqSm9MVFo2YlMwMElEWm9NVEoyTWtnek5IcHRNQ0F5YURFeWRqSklNelI2YlRJZ01tZzRkakpvTFRoNklpQm1hV3hzUFNJak1EQXdJaTgrUEhCaGRHZ2daRDBpVFRRd0lEVTBhRFIyTW1ndE5Ib2lJR1pwYkd3OUlpTkVSVVUxUkRraUx6NDhjR0YwYUNCa1BTSk5ORFlnTlRab09IWXlhQzA0ZW0weUlESm9OSFl5YUMwMGVpSWdabWxzYkQwaUl6QXdNQ0l2UGp4d1lYUm9JR1E5SWswek9DQTJNV2d5ZGpKb0xUSjZUVFF3SURZemFEUjJNbWd0TkhwTk5EUWdOakZvTkhZeWFDMDBlazAwT0NBMk0yZzBkakpvTFRSNlRUVXlJRFl4YURSMk1tZ3ROSHBOTlRZZ05qTm9OSFl5YUMwMGVrMDJNQ0EyTVdneWRqSm9MVEo2VFRZMElEWTFhREoyTW1ndE1ucE5NelFnTmpWb01uWXlhQzB5ZWswek5pQTJNMmd5ZGpKb0xUSjZUVFl5SURZemFESjJNbWd0TW5vaUlHWnBiR3c5SWlNd01EQWlMejQ4Y0dGMGFDQmpiR0Z6Y3owaWFDSWdaRDBpVFRRM0lESTBhRFIyTW1ndE5Ib2lMejQ4Y0dGMGFDQmpiR0Z6Y3owaWFDQmxJaUJrUFNKTk56Z2dOREIyTWtnMk0zWXRObWd0TW5ZdE1tZ3RNbll0TW1ndE1uWXRNbWd0TW5ZdE1tZzBkakpvTW5ZeWFESjJNbWd6ZGpKb01uWXlhREoyTW5vaUx6NDhjR0YwYUNCamJHRnpjejBpYUNCbElpQmtQU0pOTlRrZ016WjJMVEpvTFRKMkxUSm9MVEoyTFRKb0xUSjJMVEpvTFRSMk1tZ3RNbll5YUMweWRqSm9MVEoyTFRKb01uWXRNbWd5ZGkweWFDMDRkakpvTFRKMk1tZ3RNbll5YUMwemRqSm9MVEoyTW1ndE1uWTBhRE16ZGkwMmVtMHRPU0EwU0RNNGRpMHlhREoyTFRKb09IWXlhREo2SWk4K1BDOXpkbWMrIiwgImF0dHJpYnV0ZXMiOiBbeyJ0cmFpdF90eXBlIjogIkZyYW1lIiwgInZhbHVlIjogIkZyYW1lIDMifSwgeyJ0cmFpdF90eXBlIjogIkNoaXAgRGV0YWlsIiwgInZhbHVlIjogIiMxNDc3RkIifSwgeyJ0cmFpdF90eXBlIjogIkNoaXAgRGV0YWlsIiwgInZhbHVlIjogIkRldGFpbCAxIn0sIHsidHJhaXRfdHlwZSI6ICJDaGlwIERldGFpbCBDb2xvciIsICJ2YWx1ZSI6ICIjMTQ3N0ZCIn0sIHsidHJhaXRfdHlwZSI6ICJDb3JuZXIiLCAidmFsdWUiOiAiQWxwaGEgTm9kZSJ9LHsidHJhaXRfdHlwZSI6ICJIZWFkIFNoYXBlIiwgInZhbHVlIjogInJvdW5kIn0sIHsidHJhaXRfdHlwZSI6ICJIZWFkIFNoYXBlIENvbG9yIiwgInZhbHVlIjogIiNGQjE0NjcifSx7InRyYWl0X3R5cGUiOiAiRXllcyIsICJ2YWx1ZSI6ICJQYW5kYSJ9LHsidHJhaXRfdHlwZSI6ICJNb3V0aCIsICJ2YWx1ZSI6ICJTY2FyZWQifSx7InRyYWl0X3R5cGUiOiAiSGVhZCBEZXRhaWwiLCAidmFsdWUiOiAiSGF0In0sIHsidHJhaXRfdHlwZSI6ICJIZWFkIERldGFpbCBDb2xvciIsICJ2YWx1ZSI6ICIjMTQ3N0ZCIn1dfQ==",
						},
					},
				},
				Status:    true,
				Timestamp: 1710228453,
			},
			wantError: require.NoError,
		},
		// SAVM
		{
			name: "Transfer WBTC tokens on SAVM Alpha Mainnet",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkSatoshiVM,
					ChainID: 3109,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x67cdef9e2558ebd2d278035a010abd52b617e90c1e588c5a762c2e7b5a837f38"),
						ParentHash:   common.HexToHash("0x14fff53d924e94e557a93ade1e460303c65fee97dbdc4ed5aa14e142d56bed2b"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x5210cAa3ff0033B7e3566cC83e8F49D7D4297656"),
						Number:       lo.Must(new(big.Int).SetString("112532", 0)),
						GasLimit:     50000000,
						GasUsed:      30322,
						Timestamp:    1710570057,
						BaseFee:      lo.Must(new(big.Int).SetString("75000000", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xf0db1d72d78fadc2740d6ab75a8f561ebd81ddaa3a78628afd59320f945150fc"),
						From:      common.HexToAddress("0x9Fc5B67633C82607e2e6ab02BA4fd0c4CF8cbFfb"),
						Gas:       113850,
						GasPrice:  lo.Must(new(big.Int).SetString("1575000000", 10)),
						Hash:      common.HexToHash("0xf0db1d72d78fadc2740d6ab75a8f561ebd81ddaa3a78628afd59320f945150fc"),
						Input:     hexutil.MustDecode("0x2e1a7d4d0000000000000000000000000000000000000000000000000001fd5b4b7fbb59"),
						To:        lo.ToPtr(common.HexToAddress("0x5db252ead05C54B08A83414adCAbF46Eaa9E0337")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("3109", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x67cdef9e2558ebd2d278035a010abd52b617e90c1e588c5a762c2e7b5a837f38"),
						BlockNumber:       lo.Must(new(big.Int).SetString("112532", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 30322,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x5de097c0"),
						GasUsed:           30322,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x5db252ead05C54B08A83414adCAbF46Eaa9E0337"),
							Topics: []common.Hash{
								common.HexToHash("0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65"),
								common.HexToHash("0x0000000000000000000000009fc5b67633c82607e2e6ab02ba4fd0c4cf8cbffb"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000001fd5b4b7fbb59"),
							BlockNumber:     lo.Must(new(big.Int).SetString("112532", 0)),
							TransactionHash: common.HexToHash("0xf0db1d72d78fadc2740d6ab75a8f561ebd81ddaa3a78628afd59320f945150fc"),
							Index:           0,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xf0db1d72d78fadc2740d6ab75a8f561ebd81ddaa3a78628afd59320f945150fc"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkSatoshiVM,
					Endpoint: endpoint.MustGet(filter.NetworkSatoshiVM),
				},
			},
			want: &schema.Feed{
				ID:      "0xf0db1d72d78fadc2740d6ab75a8f561ebd81ddaa3a78628afd59320f945150fc",
				Network: filter.NetworkSatoshiVM,
				Index:   0,
				From:    "0x9Fc5B67633C82607e2e6ab02BA4fd0c4CF8cbFfb",
				To:      "0x5db252ead05C54B08A83414adCAbF46Eaa9E0337",
				Type:    filter.TypeUnknown,
				Calldata: &schema.Calldata{
					FunctionHash:   "0x2e1a7d4d",
					ParsedFunction: "withdraw",
				},
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("47757150000000")),
					Decimal: 18,
				},
				Actions:   []*schema.Action{},
				Status:    true,
				Timestamp: 1710570057,
			},
			wantError: require.NoError,
		},
		{
			name: "Approve SAVM tokens on SAVM Alpha Mainnet",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkSatoshiVM,
					ChainID: 3109,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x98e90ef0ac2c35830a73294d35787697ee321993fa21d8a715770a68a7ce7011"),
						ParentHash:   common.HexToHash("0x69a9c5025ad2b65fd8cee9155edee317369b61ffff9ec013215aa158908c15a1"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x5210cAa3ff0033B7e3566cC83e8F49D7D4297656"),
						Number:       lo.Must(new(big.Int).SetString("172442", 0)),
						GasLimit:     50000000,
						GasUsed:      46550,
						Timestamp:    1710749793,
						BaseFee:      lo.Must(new(big.Int).SetString("75000000", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xf28475ceeb1fc4857012fa9b26d415a137c7d6ccc1e842ebae4aca6d1cf83de9"),
						From:      common.HexToAddress("0x5481300cF3632225E833868f120826680E3688d8"),
						Gas:       51649,
						GasPrice:  lo.Must(new(big.Int).SetString("75000000", 10)),
						Hash:      common.HexToHash("0xf28475ceeb1fc4857012fa9b26d415a137c7d6ccc1e842ebae4aca6d1cf83de9"),
						Input:     hexutil.MustDecode("0x095ea7b300000000000000000000000014743ab346b36a3d48be6842f178deadb8a71623ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"),
						To:        lo.ToPtr(common.HexToAddress("0x0E02765992f946397E6d2e65642eABb9cc674928")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("3109", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x98e90ef0ac2c35830a73294d35787697ee321993fa21d8a715770a68a7ce7011"),
						BlockNumber:       lo.Must(new(big.Int).SetString("172442", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 46550,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x47868c0"),
						GasUsed:           46550,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x0E02765992f946397E6d2e65642eABb9cc674928"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x0000000000000000000000005481300cf3632225e833868f120826680e3688d8"),
								common.HexToHash("0x00000000000000000000000014743ab346b36a3d48be6842f178deadb8a71623"),
							},
							Data:            hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"),
							BlockNumber:     lo.Must(new(big.Int).SetString("172442", 0)),
							TransactionHash: common.HexToHash("0xf28475ceeb1fc4857012fa9b26d415a137c7d6ccc1e842ebae4aca6d1cf83de9"),
							Index:           0,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xf28475ceeb1fc4857012fa9b26d415a137c7d6ccc1e842ebae4aca6d1cf83de9"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkSatoshiVM,
					Endpoint: endpoint.MustGet(filter.NetworkSatoshiVM),
				},
			},
			want: &schema.Feed{
				ID:      "0xf28475ceeb1fc4857012fa9b26d415a137c7d6ccc1e842ebae4aca6d1cf83de9",
				Network: filter.NetworkSatoshiVM,
				Index:   0,
				From:    "0x5481300cF3632225E833868f120826680E3688d8",
				To:      "0x0E02765992f946397E6d2e65642eABb9cc674928",
				Type:    filter.TypeTransactionApproval,
				Calldata: &schema.Calldata{
					FunctionHash:   "0x095ea7b3",
					ParsedFunction: "approve",
				},
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("3491250000000")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type: filter.TypeTransactionApproval,
						From: "0x5481300cF3632225E833868f120826680E3688d8",
						To:   "0x14743ab346b36a3d48bE6842f178dEAdB8A71623",
						Metadata: metadata.TransactionApproval{
							Action: metadata.ActionTransactionApprove,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x0E02765992f946397E6d2e65642eABb9cc674928"),
								Name:     "SatoshiVM",
								Symbol:   "SAVM",
								Decimals: 18,
								Standard: metadata.StandardERC20,
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("115792089237316195423570985008687907853269984665640564039457584007913129639935"))),
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1710749793,
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
