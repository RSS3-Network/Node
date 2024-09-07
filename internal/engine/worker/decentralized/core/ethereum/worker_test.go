package ethereum_test

import (
	"context"
	"math/big"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/redis"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/decentralized/core/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/rss3"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	redisx "github.com/rss3-network/node/provider/redis"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

var (
	setupOnce   sync.Once
	redisClient rueidis.Client
)

func setup(t *testing.T) {
	setupOnce.Do(func() {
		var err error

		// Start Redis container with TLS
		preset := redis.Preset(
			redis.WithVersion("6.0.9"),
		)

		container, err := gnomock.Start(preset)
		require.NoError(t, err)

		t.Cleanup(func() {
			require.NoError(t, gnomock.Stop(container))
		})

		// Connect to Redis without TLS
		redisClient, err = redisx.NewClient(config.Redis{
			Endpoint: container.DefaultAddress(),
			TLS: config.RedisTLS{
				Enabled:            false,
				CAFile:             "/path/to/ca.crt",
				CertFile:           "/path/to/client.crt",
				KeyFile:            "/path/to/client.key",
				InsecureSkipVerify: false,
			},
		})
		require.NoError(t, err)
	})
}

func TestWorker_Ethereum(t *testing.T) {
	t.Parallel()
	setup(t)

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
			name: "Transfer native token on Ethereum",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
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
					Network: network.Ethereum,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Ethereum),
					},
				},
			},
			want: &activityx.Activity{
				ID:       "0x0c2f413efbc243f3bb8edac7e70bdc21936e01401a21b0d63e97732aa80f5d99",
				Network:  network.Ethereum,
				Index:    244,
				From:     "0x000000A52a03835517E9d193B3c27626e1Bc96b1",
				To:       "0xA1b2DCAC834117F38FB0356b5176B5693E165c90",
				Type:     typex.TransactionTransfer,
				Calldata: &activityx.Calldata{},
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("335686667463000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type: typex.TransactionTransfer,
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
					Network: network.Ethereum,
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
					Network: network.Ethereum,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Ethereum),
					},
				},
			},
			want: &activityx.Activity{
				ID:       "0x228cab09172d5c01cdb0ad311ff384bfbbc9dc5cd149f9057bc9bd18ae7b491c",
				Network:  network.Ethereum,
				Index:    293,
				From:     "0xc5750031497c6ae819578016Da83E035964418B8",
				To:       "0x0000000000000000000000000000000000000000",
				Type:     typex.TransactionBurn,
				Calldata: &activityx.Calldata{},
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("1155455364252000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type: typex.TransactionBurn,
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
					Network: network.Ethereum,
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
					Network: network.Ethereum,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Ethereum),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x4de22d893cd30c987ec54f9d550d9e5fa7e1ce76767a82bc732d49a81cf4e408",
				Network: network.Ethereum,
				Index:   22,
				From:    "0xE93381fB4c4F14bDa253907b18faD305D799241a",
				To:      "0xc98D64DA73a6616c42117b582e832812e7B8D57F",
				Type:    typex.TransactionTransfer,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xa9059cbb",
				},
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("811508239167366")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type: typex.TransactionTransfer,
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
					Network: network.Ethereum,
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
					Network: network.Ethereum,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Ethereum),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x164b78e31e14e08c2e445050592ae58f1d72b26d541761c2f12f7c511ef89c8a",
				Network: network.Ethereum,
				Index:   271,
				From:    "0x123456D72E335b4415b7A5665C254578575a8A0e",
				To:      "0x5452C7fB99D99fAb3Cc1875E9DA9829Cb50F7A13",
				Type:    typex.CollectibleTransfer,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x23b872dd",
				},
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("941262359203425")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type: typex.CollectibleApproval,
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
							},
						},
					},
					{
						Type: typex.CollectibleTransfer,
						From: "0x123456D72E335b4415b7A5665C254578575a8A0e",
						To:   "0x000000A52a03835517E9d193B3c27626e1Bc96b1",
						Metadata: metadata.CollectibleTransfer{
							Address:  lo.ToPtr("0x5452C7fB99D99fAb3Cc1875E9DA9829Cb50F7A13"),
							ID:       lo.ToPtr(decimal.NewFromInt(1032)),
							Value:    lo.ToPtr(decimal.New(1, 0)),
							Name:     "The Genesis RSS3 Avatar NFT",
							Symbol:   "The Genesis RSS3 Avatar NFT",
							Standard: metadata.StandardERC721,
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
		// 			Network: network.Ethereum,
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
		// 			Network:  network.Ethereum,
		// 			Endpoint: config.Endpoint{
		//				URL: endpoint.MustGet(network.Ethereum),
		//			},
		// 		},
		// 	},
		// 	want: &activityx.Activity{
		// 		ID:      "0x344c70bb46d4ce1cc56a7aacd67092a17e854bf141818c8f0c66df60c61aea92",
		// 		Network: network.Ethereum,
		// 		Index:   135,
		// 		From:    "0x6C9AE8e43AF92BB4CA6E040A610c6eF42C0245C6",
		// 		To:      "0x8442864d6AB62a9193be2F16580c08E0D7BCda2f",
		// 		Type:    typex.CollectibleTransfer,
		// 		Fee: &activityx.Fee{
		// 			Amount:  lo.Must(decimal.NewFromString("1018675320043040")),
		// 			Decimal: 18,
		// 		},
		// 		Actions: []*activityx.Action{
		// 			{
		// 				Type: typex.CollectibleTransfer,
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
					Network: network.Ethereum,
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
					Network: network.Ethereum,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Ethereum),
					},
				},
			},
			want: &activityx.Activity{
				ID:       "0xc5bb7e79738b494b9742ccc78451e4884304ad22471dbbaaf6fc70d5572ae284",
				Network:  network.Ethereum,
				Index:    8,
				From:     "0x1d97d5c7d68E03BAe6FBb1ec6E5887d6eAaaAA7d",
				To:       "0xdd9176eA3E7559D6B68b537eF555D3e89403f742",
				Type:     typex.TransactionTransfer,
				Calldata: &activityx.Calldata{},
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("50290295402717")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type: typex.TransactionTransfer,
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
					Network: network.Ethereum,
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
					Network: network.Ethereum,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Ethereum),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0xded5fd30829d08dcf608539e993c5916f23c19c48ee0ebac8b748e9aec1e6ce5",
				Network: network.Ethereum,
				Index:   160,
				From:    "0x38d004aea23c1DC86F9cFF25dcda840Ae9dF8dAe",
				To:      "0xF047ab4c75cebf0eB9ed34Ae2c186f3611aEAfa6",
				Type:    typex.Unknown,
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("1995855251931234")),
					Decimal: 18,
				},
				Calldata: &activityx.Calldata{
					FunctionHash: "0xf6203e35",
				},
				Actions:   []*activityx.Action{},
				Status:    true,
				Timestamp: 1708946831,
			},
			wantError: require.NoError,
		},
		{
			name: "Transfer RSS3 on VSL",
			arguments: arguments{
				task: &source.Task{
					Network: network.VSL,
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
					Network: network.VSL,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.VSL),
					},
				},
			},
			want: &activityx.Activity{
				ID:       "0xa2e19137d6164d8a73e031c6ad34dd44f4c6244996e5250485e4c4dbb08e6d06",
				Network:  network.VSL,
				Index:    1,
				From:     "0x27005E931Ea9ff482a3e5deb4c6eFa980305d74a",
				To:       "0xDBEe696fA4398649DFD6a6B9F474B7Dfd4CAA7e0",
				Type:     typex.TransactionTransfer,
				Calldata: &activityx.Calldata{},
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("654191108761866000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type: typex.TransactionTransfer,
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
			name: "Transfer WETH tokens on VSL",
			arguments: arguments{
				task: &source.Task{
					Network: network.VSL,
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
					Network: network.VSL,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.VSL),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0xdfd0b84e7a7e076f97f69e5027117edc7bc34f849165d9b2adc95f7cae44455b",
				Network: network.VSL,
				Index:   1,
				From:    "0xF209b7Bbadf8d9518a822aEaa7119B38b17377A7",
				To:      "0xB659A97D8ae3c43E91Eafe5dba110a7e799157c4",
				Type:    typex.TransactionTransfer,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xa9059cbb",
				},
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("1013690438583071400")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type: typex.TransactionTransfer,
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
		//{
		//	name: "Mint Open Chips tokens on VSL",
		//	arguments: arguments{
		//		task: &source.Task{
		//			Network: network.VSL,
		//			ChainID: 12553,
		//			Header: &ethereum.Header{
		//				Hash:         common.HexToHash("0xc358a8dc0da6cac588e5a98b805d60cd5d2746eced7606cb153155fccd6ea460"),
		//				ParentHash:   common.HexToHash("0xcebd5a7ab6f8fc7b396225f6c8a9a3b4963da19ac4165629cb73ab1012dde314"),
		//				UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
		//				Coinbase:     common.HexToAddress("0x4200000000000000000000000000000000000011"),
		//				Number:       lo.Must(new(big.Int).SetString("184967", 0)),
		//				GasLimit:     30000000,
		//				GasUsed:      184758,
		//				Timestamp:    1710228453,
		//				BaseFee:      lo.Must(new(big.Int).SetString("50", 0)),
		//				Transactions: nil,
		//			},
		//			Transaction: &ethereum.Transaction{
		//				BlockHash: common.HexToHash("0x26fbdff8a7af1a1917b6dd8a54c636151cb7186a6b3c3a5922df3a98edda54a8"),
		//				From:      common.HexToAddress("0x30286DD245338292F319809935a1037CcD4573Ea"),
		//				Gas:       140094,
		//				GasPrice:  lo.Must(new(big.Int).SetString("1000050", 10)),
		//				Hash:      common.HexToHash("0x26fbdff8a7af1a1917b6dd8a54c636151cb7186a6b3c3a5922df3a98edda54a8"),
		//				Input:     hexutil.MustDecode("0x2647620400000000000000000000000039f9e912c1f696f533e7a2267ea233aec9742b35"),
		//				To:        lo.ToPtr(common.HexToAddress("0x28F14d917fddbA0c1f2923C406952478DfDA5578")),
		//				Value:     lo.Must(new(big.Int).SetString("500000000000000000000", 0)),
		//				Type:      2,
		//				ChainID:   lo.Must(new(big.Int).SetString("12553", 0)),
		//			},
		//			Receipt: &ethereum.Receipt{
		//				BlockHash:         common.HexToHash("0xc358a8dc0da6cac588e5a98b805d60cd5d2746eced7606cb153155fccd6ea460"),
		//				BlockNumber:       lo.Must(new(big.Int).SetString("184967", 0)),
		//				ContractAddress:   nil,
		//				CumulativeGasUsed: 184758,
		//				EffectiveGasPrice: hexutil.MustDecodeBig("0xf4272"),
		//				GasUsed:           137833,
		//				L1GasPrice:        lo.Must(new(big.Int).SetString("51479747982", 0)),
		//				L1GasUsed:         lo.Must(new(big.Int).SetString("2500", 0)),
		//				L1Fee:             lo.Must(new(big.Int).SetString("772196219730000000", 0)),
		//				FeeScalar:         lo.Must(new(big.Float).SetString("6000")),
		//				Logs: []*ethereum.Log{{
		//					Address: common.HexToAddress("0x849f8F55078dCc69dD857b58Cc04631EBA54E4DE"),
		//					Topics: []common.Hash{
		//						common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
		//						common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		//						common.HexToHash("0x00000000000000000000000030286dd245338292f319809935a1037ccd4573ea"),
		//						common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000a2"),
		//					},
		//					Data:            nil,
		//					BlockNumber:     lo.Must(new(big.Int).SetString("184967", 0)),
		//					TransactionHash: common.HexToHash("0x26fbdff8a7af1a1917b6dd8a54c636151cb7186a6b3c3a5922df3a98edda54a8"),
		//					Index:           0,
		//					Removed:         false,
		//				}, {
		//					Address: common.HexToAddress("0x28F14d917fddbA0c1f2923C406952478DfDA5578"),
		//					Topics: []common.Hash{
		//						common.HexToHash("0xad3fa07f4195b47e64892eb944ecbfc253384053c119852bb2bcae484c2fcb69"),
		//						common.HexToHash("0x00000000000000000000000030286dd245338292f319809935a1037ccd4573ea"),
		//						common.HexToHash("0x00000000000000000000000039f9e912c1f696f533e7a2267ea233aec9742b35"),
		//						common.HexToHash("0x00000000000000000000000000000000000000000000001b1ae4d6e2ef500000"),
		//					},
		//					Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000a200000000000000000000000000000000000000000000000000000000000000a2"),
		//					BlockNumber:     lo.Must(new(big.Int).SetString("184967", 0)),
		//					TransactionHash: common.HexToHash("0x26fbdff8a7af1a1917b6dd8a54c636151cb7186a6b3c3a5922df3a98edda54a8"),
		//					Index:           1,
		//					Removed:         false,
		//				}},
		//				Status:           1,
		//				TransactionHash:  common.HexToHash("0x26fbdff8a7af1a1917b6dd8a54c636151cb7186a6b3c3a5922df3a98edda54a8"),
		//				TransactionIndex: 1,
		//			},
		//		},
		//		config: &config.Module{
		//			Network: network.VSL,
		//			Endpoint: config.Endpoint{
		//				URL: endpoint.MustGet(network.VSL),
		//			},
		//		},
		//	},
		//	want: &activityx.Activity{
		//		ID:      "0x26fbdff8a7af1a1917b6dd8a54c636151cb7186a6b3c3a5922df3a98edda54a8",
		//		Network: network.VSL,
		//		Index:   1,
		//		From:    "0x30286DD245338292F319809935a1037CcD4573Ea",
		//		To:      "0x28F14d917fddbA0c1f2923C406952478DfDA5578",
		//		Type:    typex.CollectibleMint,
		//		Calldata: &activityx.Calldata{
		//			FunctionHash: "0x26476204",
		//		},
		//		Fee: &activityx.Fee{
		//			Amount:  lo.Must(decimal.NewFromString("772196357569891650")),
		//			Decimal: 18,
		//		},
		//		Actions: []*activityx.Action{
		//			{
		//				Type: typex.CollectibleMint,
		//				From: "0x0000000000000000000000000000000000000000",
		//				To:   "0x30286DD245338292F319809935a1037CcD4573Ea",
		//				Metadata: metadata.CollectibleTransfer{
		//					Address:  lo.ToPtr("0x849f8F55078dCc69dD857b58Cc04631EBA54E4DE"),
		//					ID:       lo.ToPtr(decimal.NewFromInt(162)),
		//					Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
		//					Name:     "Open Chips",
		//					Symbol:   "Chips",
		//					Standard: metadata.StandardERC721,
		//				},
		//			},
		//		},
		//		Status:    true,
		//		Timestamp: 1710228453,
		//	},
		//	wantError: require.NoError,
		// },
		{
			name: "Deposit $RSS3 to the VSL Staking contract",
			arguments: arguments{
				task: &source.Task{
					Network: network.VSL,
					ChainID: 12553,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x65c5c7eebca0f8bad09de7c5ceadaa369878db3d7488a881155611e9bb18d820"),
						ParentHash:   common.HexToHash("0xfffd9b8f39aff5b2c99c3f2d3e0b3c885f3b823207d9bb7dfbdf10e856d9e102"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4200000000000000000000000000000000000011"),
						Number:       lo.Must(new(big.Int).SetString("182670", 0)),
						GasLimit:     30000000,
						GasUsed:      260163,
						Timestamp:    1710223859,
						BaseFee:      lo.Must(new(big.Int).SetString("50", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xb49208a4fb936cf3d73da0b92af0776945a09062af8983cf72c8bd4fffb7f624"),
						From:      common.HexToAddress("0x39F9e912C1F696F533e7A2267Ea233AeC9742b35"),
						Gas:       218959,
						GasPrice:  lo.Must(new(big.Int).SetString("1500000050", 10)),
						Hash:      common.HexToHash("0xb49208a4fb936cf3d73da0b92af0776945a09062af8983cf72c8bd4fffb7f624"),
						Input:     hexutil.MustDecode("0x96531623000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000003200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000667657869616f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000015466f72204f70656e20496e666f726d6174696f6e2e0000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x28F14d917fddbA0c1f2923C406952478DfDA5578")),
						Value:     lo.Must(new(big.Int).SetString("10000000000000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("12553", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x65c5c7eebca0f8bad09de7c5ceadaa369878db3d7488a881155611e9bb18d820"),
						BlockNumber:       lo.Must(new(big.Int).SetString("182670", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 260163,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x59682f32"),
						GasUsed:           213238,
						L1GasPrice:        lo.Must(new(big.Int).SetString("51211146219", 0)),
						L1GasUsed:         lo.Must(new(big.Int).SetString("3636", 0)),
						L1Fee:             lo.Must(new(big.Int).SetString("1117222365913704000", 0)),
						FeeScalar:         lo.Must(new(big.Float).SetString("6000")),
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x28F14d917fddbA0c1f2923C406952478DfDA5578"),
							Topics: []common.Hash{
								common.HexToHash("0x37570f68d94fd46cd4009b3823da2b2bc1a9a7e38f824f311ede9e876816e321"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000008"),
								common.HexToHash("0x00000000000000000000000039f9e912c1f696f533e7a2267ea233aec9742b35"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000032000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000667657869616f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000015466f72204f70656e20496e666f726d6174696f6e2e0000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("182670", 0)),
							TransactionHash: common.HexToHash("0xb49208a4fb936cf3d73da0b92af0776945a09062af8983cf72c8bd4fffb7f624"),
							Index:           0,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x28F14d917fddbA0c1f2923C406952478DfDA5578"),
							Topics: []common.Hash{
								common.HexToHash("0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4"),
								common.HexToHash("0x00000000000000000000000039f9e912c1f696f533e7a2267ea233aec9742b35"),
								common.HexToHash("0x00000000000000000000000000000000000000000000021e19e0c9bab2400000"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("182670", 0)),
							TransactionHash: common.HexToHash("0xb49208a4fb936cf3d73da0b92af0776945a09062af8983cf72c8bd4fffb7f624"),
							Index:           1,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xb49208a4fb936cf3d73da0b92af0776945a09062af8983cf72c8bd4fffb7f624"),
						TransactionIndex: 1,
					},
				},
				config: &config.Module{
					Network: network.VSL,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.VSL),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0xb49208a4fb936cf3d73da0b92af0776945a09062af8983cf72c8bd4fffb7f624",
				Network: network.VSL,
				Index:   1,
				From:    "0x39F9e912C1F696F533e7A2267Ea233AeC9742b35",
				To:      rss3.AddressStakingVSL.String(),
				Type:    typex.ExchangeStaking,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x96531623",
				},
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("1117542222924365900")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type: typex.ExchangeStaking,
						From: "0x39F9e912C1F696F533e7A2267Ea233AeC9742b35",
						To:   "0x39F9e912C1F696F533e7A2267Ea233AeC9742b35",
						Metadata: metadata.ExchangeStaking{
							Action: metadata.ActionExchangeStakingStake,
							Token: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("10000000000000000000000"))),
								Name:     "RSS3",
								Symbol:   "RSS3",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1710223859,
			},
			wantError: require.NoError,
		},
		{
			name: "Staking $RSS3 to the VSL Nodes",
			arguments: arguments{
				task: &source.Task{
					Network: network.VSL,
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
					Network: network.VSL,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.VSL),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x26fbdff8a7af1a1917b6dd8a54c636151cb7186a6b3c3a5922df3a98edda54a8",
				Network: network.VSL,
				Index:   1,
				From:    "0x30286DD245338292F319809935a1037CcD4573Ea",
				To:      rss3.AddressStakingVSL.String(),
				Type:    typex.ExchangeStaking,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x26476204",
				},
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("772196357569891650")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type: typex.CollectibleMint,
						From: ethereum.AddressGenesis.String(),
						To:   common.HexToAddress("0x30286DD245338292F319809935a1037CcD4573Ea").String(),
						Metadata: metadata.CollectibleTransfer{
							Address:  lo.ToPtr("0x849f8F55078dCc69dD857b58Cc04631EBA54E4DE"),
							ID:       lo.ToPtr(decimal.NewFromBigInt(big.NewInt(162), 0)),
							Value:    lo.ToPtr(decimal.NewFromBigInt(big.NewInt(1), 0)),
							Name:     "Open Chips",
							Symbol:   "Chips",
							Standard: metadata.StandardERC721,
						},
					},
					{
						Type: typex.ExchangeStaking,
						From: "0x30286DD245338292F319809935a1037CcD4573Ea",
						To:   "0x39F9e912C1F696F533e7A2267Ea233AeC9742b35",
						Metadata: metadata.ExchangeStaking{
							Action: metadata.ActionExchangeStakingStake,
							Token: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("500000000000000000000"))),
								Name:     "RSS3",
								Symbol:   "RSS3",
								Decimals: 18,
							},
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
			name: "Approve SAVM tokens on SAVM Alpha Mainnet",
			arguments: arguments{
				task: &source.Task{
					Network: network.SatoshiVM,
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
					Network: network.SatoshiVM,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.SatoshiVM),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0xf28475ceeb1fc4857012fa9b26d415a137c7d6ccc1e842ebae4aca6d1cf83de9",
				Network: network.SatoshiVM,
				Index:   0,
				From:    "0x5481300cF3632225E833868f120826680E3688d8",
				To:      "0x0E02765992f946397E6d2e65642eABb9cc674928",
				Type:    typex.TransactionApproval,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x095ea7b3",
				},
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("3491250000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type: typex.TransactionApproval,
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
		//	X Layer
		{
			name: "Mint WETH tokens on X Layer",
			arguments: arguments{
				task: &source.Task{
					Network: network.XLayer,
					ChainID: 196,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x951702861dbd144773cba9b9a89e48dbcc61c8c38abc884b06e0bbfd52235c42"),
						ParentHash:   common.HexToHash("0xf52ce2c55ad400604a59bba05c819c608c4f136f2fdefb8465f84eafb06b5232"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xAF9d27ffe4d51eD54AC8eEc78f2785D7E11E5ab1"),
						Number:       lo.Must(new(big.Int).SetString("4176026", 0)),
						GasLimit:     1125899906842624,
						GasUsed:      133554,
						Timestamp:    1724379264,
						BaseFee:      nil,
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xbf2d3e7b416ba49de2f1d0924634fd4415d1ccf1abad6b0f4785baff3050e732"),
						From:      common.HexToAddress("0x835a51c47E853b074524355Ff05508aE097ed0C7"),
						Gas:       106405,
						GasPrice:  lo.Must(new(big.Int).SetString("0", 10)),
						Hash:      common.HexToHash("0xbf2d3e7b416ba49de2f1d0924634fd4415d1ccf1abad6b0f4785baff3050e732"),
						Input:     hexutil.MustDecode("0xccaa2d110000000000000000000000000000000000000000000000000000000000000000ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5b4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3021ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba85e58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a193440eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a196868c6a3b4b2ebd6088b8114c20110d1b37a2e476286f80b1827790355fddae1a79867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756afba318d857ba121d5265a492d4f8c403ed2174dc8ca163920d21f38cee38cdc3df9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a521d1f66d8c0c948214c0d25da64ddca79af6d64ac8443a3dc5b7682387931dd23490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c7e27d444a8231d418ea8f26b97574a0eac4c61c4644dc682e77f0c4747d531005c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8beccda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d277a0f8b0e0b4e5a57f5e381b3892bb41a0bcdbfdf3c7d591fae02081159b594d361122b4b1d18ab577f2aeb6632c690713456a66a5670649ceb2c0a31e43ab465a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a0b46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa0c65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e2f4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd95a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3774df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee652cdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618db8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d0838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea32293237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7358448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a9000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000032a80e2407cc34e64d8a662ca7bd9b37aae728d8ba0d8060da5a46dcd70216c4b971063787901aee8d835e3004d52b39d2678a2a9a598fa8d763b00f29d112db6cbb700000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000300000000000000000000000064d22295d0ce3beb947fc13e7256361e2cfb06dd000000000000000000000000000000000000000000000000014a6701dc1c800000000000000000000000000000000000000000000000000000000000000009200000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x2a3DD3EB832aF982ec71669E178424b10Dca2EDe")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   lo.Must(new(big.Int).SetString("196", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x951702861dbd144773cba9b9a89e48dbcc61c8c38abc884b06e0bbfd52235c42"),
						BlockNumber:       lo.Must(new(big.Int).SetString("4176026", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 133554,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x0"),
						GasUsed:           105579,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x5A77f1443D16ee5761d310e38b62f77f726bC71c"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x00000000000000000000000064d22295d0ce3beb947fc13e7256361e2cfb06dd"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000014a6701dc1c8000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("4176026", 0)),
							TransactionHash: common.HexToHash("0xbf2d3e7b416ba49de2f1d0924634fd4415d1ccf1abad6b0f4785baff3050e732"),
							Index:           1,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x2a3DD3EB832aF982ec71669E178424b10Dca2EDe"),
							Topics: []common.Hash{
								common.HexToHash("0x1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000010000000000032a800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000064d22295d0ce3beb947fc13e7256361e2cfb06dd000000000000000000000000000000000000000000000000014a6701dc1c8000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("4176026", 0)),
							TransactionHash: common.HexToHash("0xbf2d3e7b416ba49de2f1d0924634fd4415d1ccf1abad6b0f4785baff3050e732"),
							Index:           2,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xbf2d3e7b416ba49de2f1d0924634fd4415d1ccf1abad6b0f4785baff3050e732"),
						TransactionIndex: 1,
					},
				},
				config: &config.Module{
					Network: network.XLayer,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.XLayer),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0xbf2d3e7b416ba49de2f1d0924634fd4415d1ccf1abad6b0f4785baff3050e732",
				Network: network.XLayer,
				Index:   1,
				From:    "0x835a51c47E853b074524355Ff05508aE097ed0C7",
				To:      "0x2a3DD3EB832aF982ec71669E178424b10Dca2EDe",
				Type:    typex.TransactionMint,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xccaa2d11",
				},
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("0")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type: typex.TransactionMint,
						From: "0x0000000000000000000000000000000000000000",
						To:   "0x64d22295D0CE3beb947Fc13e7256361E2CFb06dD",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0x5A77f1443D16ee5761d310e38b62f77f726bC71c"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("93000000000000000"))),
							Name:     "Wrapped Ether",
							Symbol:   "WETH",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
				},
				Status:    true,
				Timestamp: 1724379264,
			},
			wantError: require.NoError,
		},
		{
			name: "Transfer WETH tokens on X Layer",
			arguments: arguments{
				task: &source.Task{
					Network: network.XLayer,
					ChainID: 196,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x6baa590dd26db5a869a9fedd88dc1cb99a6430778d2b41b2b0b0d87cd8da155e"),
						ParentHash:   common.HexToHash("0x4658efb8b88f194721b1d1aea62e69c8b0c27009cf700156ec01bbec969990ed"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xAF9d27ffe4d51eD54AC8eEc78f2785D7E11E5ab1"),
						Number:       lo.Must(new(big.Int).SetString("4186444", 0)),
						GasLimit:     1125899906842624,
						GasUsed:      58198,
						Timestamp:    1724410763,
						BaseFee:      nil,
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x60a199b81e086dd8fdc03e5daf0a9c16ea57459a334bc77925f95b75b27889dd"),
						From:      common.HexToAddress("0xE4eDb277e41dc89aB076a1F049f4a3EfA700bCE8"),
						Gas:       37717,
						GasPrice:  lo.Must(new(big.Int).SetString("5500000000", 10)),
						Hash:      common.HexToHash("0x60a199b81e086dd8fdc03e5daf0a9c16ea57459a334bc77925f95b75b27889dd"),
						Input:     hexutil.MustDecode("0xa9059cbb0000000000000000000000009a018e5dfbbf40f6102ac0d707f8b68166bd568500000000000000000000000000000000000000000000000000071afd498d0040"),
						To:        lo.ToPtr(common.HexToAddress("0x5A77f1443D16ee5761d310e38b62f77f726bC71c")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   lo.Must(new(big.Int).SetString("196", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x6baa590dd26db5a869a9fedd88dc1cb99a6430778d2b41b2b0b0d87cd8da155e"),
						BlockNumber:       lo.Must(new(big.Int).SetString("4186444", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 34289,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x147d35700"),
						GasUsed:           34289,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x5A77f1443D16ee5761d310e38b62f77f726bC71c"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000e4edb277e41dc89ab076a1f049f4a3efa700bce8"),
								common.HexToHash("0x0000000000000000000000009a018e5dfbbf40f6102ac0d707f8b68166bd5685"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000071afd498d0040"),
							BlockNumber:     lo.Must(new(big.Int).SetString("4186444", 0)),
							TransactionHash: common.HexToHash("0x60a199b81e086dd8fdc03e5daf0a9c16ea57459a334bc77925f95b75b27889dd"),
							Index:           0,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x60a199b81e086dd8fdc03e5daf0a9c16ea57459a334bc77925f95b75b27889dd"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network: network.XLayer,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.XLayer),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x60a199b81e086dd8fdc03e5daf0a9c16ea57459a334bc77925f95b75b27889dd",
				Network: network.XLayer,
				Index:   0,
				From:    "0xE4eDb277e41dc89aB076a1F049f4a3EfA700bCE8",
				To:      "0x5A77f1443D16ee5761d310e38b62f77f726bC71c",
				Type:    typex.TransactionTransfer,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xa9059cbb",
				},
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("188589500000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type: typex.TransactionTransfer,
						From: "0xE4eDb277e41dc89aB076a1F049f4a3EfA700bCE8",
						To:   "0x9A018E5dfBbF40F6102ac0d707F8b68166Bd5685",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0x5A77f1443D16ee5761d310e38b62f77f726bC71c"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2000000000000064"))),
							Name:     "Wrapped Ether",
							Symbol:   "WETH",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
				},
				Status:    true,
				Timestamp: 1724410763,
			},
			wantError: require.NoError,
		},
		{
			name: "Burn WETH tokens on X Layer",
			arguments: arguments{
				task: &source.Task{
					Network: network.XLayer,
					ChainID: 196,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x488966882718029e467763448d65f487d203d05e98caacd7e22206833c48df9f"),
						ParentHash:   common.HexToHash("0x3c4125b1e84f80a36fe6e809f5106cdc2f01774b6e6ba9806ade55c2a5e07891"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xAF9d27ffe4d51eD54AC8eEc78f2785D7E11E5ab1"),
						Number:       lo.Must(new(big.Int).SetString("3930575", 0)),
						GasLimit:     1125899906842624,
						GasUsed:      207514,
						Timestamp:    1723636165,
						BaseFee:      nil,
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x06ea5dcbaeac9f61b7dbac2f59a2ded44a20e0831d4279f98f55085d3f2575b1"),
						From:      common.HexToAddress("0xEEc2fC175C458EC8501b905A0097694Ad93E54D0"),
						Gas:       760890,
						GasPrice:  lo.Must(new(big.Int).SetString("5000000000", 10)),
						Hash:      common.HexToHash("0x06ea5dcbaeac9f61b7dbac2f59a2ded44a20e0831d4279f98f55085d3f2575b1"),
						Input:     hexutil.MustDecode("0x972250fe0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000001d000000000000000000000000eec2fc175c458ec8501b905a0097694ad93e54d00000000000000000000000001e4a5963abfd975d8c9021ce480b42188849d41d0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000098968000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002a30786565633266633137356334353865633835303162393035613030393736393461643933653534643000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x5965851f21DAE82eA7C62f87fb7C57172E9F2adD")),
						Value:     lo.Must(new(big.Int).SetString("37427638103441460", 0)),
						Type:      0,
						ChainID:   lo.Must(new(big.Int).SetString("196", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x488966882718029e467763448d65f487d203d05e98caacd7e22206833c48df9f"),
						BlockNumber:       lo.Must(new(big.Int).SetString("3930575", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 207514,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x12a05f200"),
						GasUsed:           207514,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x1E4a5963aBFD975d8c9021ce480b42188849D41d"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000eec2fc175c458ec8501b905a0097694ad93e54d0"),
								common.HexToHash("0x0000000000000000000000008b773d83bc66be128c60e07e17c8901f7a64f000"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3930575", 0)),
							TransactionHash: common.HexToHash("0x06ea5dcbaeac9f61b7dbac2f59a2ded44a20e0831d4279f98f55085d3f2575b1"),
							Index:           0,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x1E4a5963aBFD975d8c9021ce480b42188849D41d"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000eec2fc175c458ec8501b905a0097694ad93e54d0"),
								common.HexToHash("0x000000000000000000000000ace2b97d84dea5c202d9aa538610a9845fdfd244"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000989680"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3930575", 0)),
							TransactionHash: common.HexToHash("0x06ea5dcbaeac9f61b7dbac2f59a2ded44a20e0831d4279f98f55085d3f2575b1"),
							Index:           1,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x1E4a5963aBFD975d8c9021ce480b42188849D41d"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000ace2b97d84dea5c202d9aa538610a9845fdfd244"),
								common.HexToHash("0x000000000000000000000000a28062bd708ce49e9311d6293def7df63f2b0816"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3930575", 0)),
							TransactionHash: common.HexToHash("0x06ea5dcbaeac9f61b7dbac2f59a2ded44a20e0831d4279f98f55085d3f2575b1"),
							Index:           2,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x1E4a5963aBFD975d8c9021ce480b42188849D41d"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000ace2b97d84dea5c202d9aa538610a9845fdfd244"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000989680"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3930575", 0)),
							TransactionHash: common.HexToHash("0x06ea5dcbaeac9f61b7dbac2f59a2ded44a20e0831d4279f98f55085d3f2575b1"),
							Index:           3,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x2a3DD3EB832aF982ec71669E178424b10Dca2EDe"),
							Topics: []common.Hash{
								common.HexToHash("0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec70000000000000000000000000000000000000000000000000000000000000000000000000000000000000000eec2fc175c458ec8501b905a0097694ad93e54d0000000000000000000000000000000000000000000000000000000000098968000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000e5500000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a546574686572205553440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000045553445400000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3930575", 0)),
							TransactionHash: common.HexToHash("0x06ea5dcbaeac9f61b7dbac2f59a2ded44a20e0831d4279f98f55085d3f2575b1"),
							Index:           4,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xAcE2b97D84DEA5c202D9Aa538610A9845fdFD244"),
							Topics: []common.Hash{
								common.HexToHash("0xf45e63959fcd4f065d8bf61ed6ff6c58093327df820aab9a1cbf1d44e2532af7"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000084f83d11f000340000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3930575", 0)),
							TransactionHash: common.HexToHash("0x06ea5dcbaeac9f61b7dbac2f59a2ded44a20e0831d4279f98f55085d3f2575b1"),
							Index:           5,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xAcE2b97D84DEA5c202D9Aa538610A9845fdFD244"),
							Topics: []common.Hash{
								common.HexToHash("0x6ea4114cf93cff6c9ff22aef8d246e78a0368fa48033adc49edde5fa273f69a1"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000eec2fc175c458ec8501b905a0097694ad93e54d0000000000000000000000000eec2fc175c458ec8501b905a0097694ad93e54d00000000000000000000000001e4a5963abfd975d8c9021ce480b42188849d41d00000000000000000000000000000000000000000000000000000000009896800000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3930575", 0)),
							TransactionHash: common.HexToHash("0x06ea5dcbaeac9f61b7dbac2f59a2ded44a20e0831d4279f98f55085d3f2575b1"),
							Index:           6,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x5965851f21DAE82eA7C62f87fb7C57172E9F2adD"),
							Topics: []common.Hash{
								common.HexToHash("0xf6481cbc1da19356c5cb6b884be507da735b89f21dc4bbb7c9b7cc0968b03b7a"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000001d"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000eec2fc175c458ec8501b905a0097694ad93e54d0000000000000000000000000eec2fc175c458ec8501b905a0097694ad93e54d00000000000000000000000001e4a5963abfd975d8c9021ce480b42188849d41d0000000000000000000000000000000000000000000000000000000000989680000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030786565633266633137356334353865633835303162393035613030393736393461643933653534643000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3930575", 0)),
							TransactionHash: common.HexToHash("0x06ea5dcbaeac9f61b7dbac2f59a2ded44a20e0831d4279f98f55085d3f2575b1"),
							Index:           7,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x06ea5dcbaeac9f61b7dbac2f59a2ded44a20e0831d4279f98f55085d3f2575b1"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network: network.XLayer,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.XLayer),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x06ea5dcbaeac9f61b7dbac2f59a2ded44a20e0831d4279f98f55085d3f2575b1",
				Network: network.XLayer,
				Index:   0,
				From:    "0xEEc2fC175C458EC8501b905A0097694Ad93E54D0",
				To:      "0x5965851f21DAE82eA7C62f87fb7C57172E9F2adD",
				Type:    typex.TransactionBurn,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x972250fe",
				},
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("1037570000000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type: typex.TransactionApproval,
						From: "0xEEc2fC175C458EC8501b905A0097694Ad93E54D0",
						To:   "0x8b773D83bc66Be128c60e07E17C8901f7a64F000",
						Metadata: metadata.TransactionApproval{
							Action: metadata.ActionTransactionRevoke,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x1E4a5963aBFD975d8c9021ce480b42188849D41d"),
								Name:     "Tether USD",
								Symbol:   "USDT",
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("0"))),
								Decimals: 6,
								Standard: metadata.StandardERC20,
							},
						},
					},
					{
						Type: typex.TransactionTransfer,
						From: "0xEEc2fC175C458EC8501b905A0097694Ad93E54D0",
						To:   "0xAcE2b97D84DEA5c202D9Aa538610A9845fdFD244",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0x1E4a5963aBFD975d8c9021ce480b42188849D41d"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("10000000"))),
							Name:     "Tether USD",
							Symbol:   "USDT",
							Decimals: 6,
							Standard: metadata.StandardERC20,
						},
					},
					{
						Type: typex.TransactionTransfer,
						From: "0xAcE2b97D84DEA5c202D9Aa538610A9845fdFD244",
						To:   "0xA28062Bd708Ce49e9311d6293dEf7DF63f2b0816",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0x1E4a5963aBFD975d8c9021ce480b42188849D41d"),
							Name:     "Tether USD",
							Symbol:   "USDT",
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("0"))),
							Decimals: 6,
							Standard: metadata.StandardERC20,
						},
					},
					{
						Type: typex.TransactionBurn,
						From: "0xAcE2b97D84DEA5c202D9Aa538610A9845fdFD244",
						To:   "0x0000000000000000000000000000000000000000",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0x1E4a5963aBFD975d8c9021ce480b42188849D41d"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("10000000"))),
							Name:     "Tether USD",
							Symbol:   "USDT",
							Decimals: 6,
							Standard: metadata.StandardERC20,
						},
					},
				},
				Status:    true,
				Timestamp: 1723636165,
			},
			wantError: require.NoError,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			instance, err := worker.NewWorker(testcase.arguments.config, redisClient)
			require.NoError(t, err)

			activity, err := instance.Transform(ctx, testcase.arguments.task)
			testcase.wantError(t, err)

			// t.Log(string(lo.Must(json.MarshalIndent(activity, "", "\x20\x20"))))

			require.Equal(t, testcase.want, activity)
		})
	}
}
