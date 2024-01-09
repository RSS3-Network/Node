package ethereum_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/naturalselectionlabs/rss3-node/config"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	worker "github.com/naturalselectionlabs/rss3-node/internal/engine/worker/fallback/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/endpoint"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
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
				ID:      "0x0c2f413efbc243f3bb8edac7e70bdc21936e01401a21b0d63e97732aa80f5d99",
				Network: filter.NetworkEthereum,
				Index:   244,
				From:    "0x000000A52a03835517E9d193B3c27626e1Bc96b1",
				To:      "0xA1b2DCAC834117F38FB0356b5176B5693E165c90",
				Type:    filter.TypeTransactionTransfer,
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
				ID:      "0x228cab09172d5c01cdb0ad311ff384bfbbc9dc5cd149f9057bc9bd18ae7b491c",
				Network: filter.NetworkEthereum,
				Index:   293,
				From:    "0xc5750031497c6ae819578016Da83E035964418B8",
				To:      "0x0000000000000000000000000000000000000000",
				Type:    filter.TypeTransactionBurn,
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
							Standard: contract.StandardERC20,
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
								Standard: contract.StandardERC721,
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
							Standard: contract.StandardERC721,
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
		// 					Standard: contract.StandardERC1155,
		// 					URI:      "https://timepieces.mypinata.cloud/ipfs/QmbZtT3KVHYL6QAuNbrLGKp4sQyaWkkzTDRvHFCmq7Py6y",
		// 				},
		// 			},
		// 		},
		// 		Status:    true,
		// 		Timestamp: 1678657307,
		// 	},
		// 	wantError: require.NoError,
		// },
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
			require.Equal(t, testcase.want, feed)

			t.Log(feed)
		})
	}
}
