package ens_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/testserver"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/database/dialer"
	"github.com/rss3-network/node/internal/database/model"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/contract/ens"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/ens"
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
			name: "ENS NameRegistered V1",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xff066467303b45a198cbfc148166ac1e00779a4e6feaf265fd77b021c13f678b"),
						ParentHash:   common.HexToHash("0xd0577760dddbee2a696c4c1738e3a5a955c28877f596186c6689a51e7f9bc77f"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xB3E84B6C6409826DC45432B655D8C9489A14A0D7"),
						Number:       lo.Must(new(big.Int).SetString("17475260", 0)),
						GasLimit:     30000000,
						GasUsed:      11832216,
						Timestamp:    1686710135,
						BaseFee:      lo.Must(new(big.Int).SetString("17948559167", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x85cfc1b0641425e7aefeef65f9fe274deff13154e14394b0c9fd316746a98d10"),
						From:      common.HexToAddress("0x7Da1E64e92094075A6D9b803d948015C993Bb58D"),
						Gas:       398440,
						GasPrice:  lo.Must(new(big.Int).SetString("18042559167", 10)),
						Hash:      common.HexToHash("0x85cfc1b0641425e7aefeef65f9fe274deff13154e14394b0c9fd316746a98d10"),
						Input:     hexutil.MustDecode("0xf7a1696300000000000000000000000000000000000000000000000000000000000000c00000000000000000000000007da1e64e92094075a6d9b803d948015c993bb58d0000000000000000000000000000000000000000000000000000000001e18558b459e752b076da16d21578c19b8aaa6f03a35cd942de31126d4b4947b877b7c0000000000000000000000000231b0ee14048e9dccd1d247744d114a4eb5e8e630000000000000000000000007da1e64e92094075a6d9b803d948015c993bb58d000000000000000000000000000000000000000000000000000000000000000c64617a75697169616e6764610000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x1d6552e8F46fD509f3918A174FE62C34b42564aE")),
						Value:     lo.Must(new(big.Int).SetString("3151952355793017", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xff066467303b45a198cbfc148166ac1e00779a4e6feaf265fd77b021c13f678b"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17475260", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 11239975,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x4336b9abf"),
						GasUsed:           300291,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000283af0b28c62c092c9727f1ee09c02ca627eb7f5"),
								common.HexToHash("0x23913293c04ada44918795ed0efdcb432ed227577f69dfc432395af6d0551173"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("17475260", 0)),
							TransactionHash: common.HexToHash("0x85cfc1b0641425e7aefeef65f9fe274deff13154e14394b0c9fd316746a98d10"),
							Index:           263,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e"),
							Topics: []common.Hash{
								common.HexToHash("0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82"),
								common.HexToHash("0x93cdeb708b7545dc668eb9280176169d1c33cfd8ed6f04690a0bcc88a93fc4ae"),
								common.HexToHash("0x23913293c04ada44918795ed0efdcb432ed227577f69dfc432395af6d0551173"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000283af0b28c62c092c9727f1ee09c02ca627eb7f5"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17475260", 0)),
							TransactionHash: common.HexToHash("0x85cfc1b0641425e7aefeef65f9fe274deff13154e14394b0c9fd316746a98d10"),
							Index:           264,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85"),
							Topics: []common.Hash{
								common.HexToHash("0xb3d987963d01b2f68493b4bdb130988f157ea43070d4ad840fee0466ed9370d9"),
								common.HexToHash("0x23913293c04ada44918795ed0efdcb432ed227577f69dfc432395af6d0551173"),
								common.HexToHash("0x000000000000000000000000283af0b28c62c092c9727f1ee09c02ca627eb7f5"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000666aaccf"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17475260", 0)),
							TransactionHash: common.HexToHash("0x85cfc1b0641425e7aefeef65f9fe274deff13154e14394b0c9fd316746a98d10"),
							Index:           265,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e"),
							Topics: []common.Hash{
								common.HexToHash("0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0"),
								common.HexToHash("0x730e1dbcdaeb9ad6b4e6a29170a4acce635b5f30f45d81f9b590a007bf08e9d7"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000231b0ee14048e9dccd1d247744d114a4eb5e8e63"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17475260", 0)),
							TransactionHash: common.HexToHash("0x85cfc1b0641425e7aefeef65f9fe274deff13154e14394b0c9fd316746a98d10"),
							Index:           266,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x231b0Ee14048e9dCcD1d247744d114a4EB5E8E63"),
							Topics: []common.Hash{
								common.HexToHash("0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752"),
								common.HexToHash("0x730e1dbcdaeb9ad6b4e6a29170a4acce635b5f30f45d81f9b590a007bf08e9d7"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000147da1e64e92094075a6d9b803d948015c993bb58d000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17475260", 0)),
							TransactionHash: common.HexToHash("0x85cfc1b0641425e7aefeef65f9fe274deff13154e14394b0c9fd316746a98d10"),
							Index:           267,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x231b0Ee14048e9dCcD1d247744d114a4EB5E8E63"),
							Topics: []common.Hash{
								common.HexToHash("0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2"),
								common.HexToHash("0x730e1dbcdaeb9ad6b4e6a29170a4acce635b5f30f45d81f9b590a007bf08e9d7"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000007da1e64e92094075a6d9b803d948015c993bb58d"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17475260", 0)),
							TransactionHash: common.HexToHash("0x85cfc1b0641425e7aefeef65f9fe274deff13154e14394b0c9fd316746a98d10"),
							Index:           268,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e"),
							Topics: []common.Hash{
								common.HexToHash("0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82"),
								common.HexToHash("0x93cdeb708b7545dc668eb9280176169d1c33cfd8ed6f04690a0bcc88a93fc4ae"),
								common.HexToHash("0x23913293c04ada44918795ed0efdcb432ed227577f69dfc432395af6d0551173"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000007da1e64e92094075a6d9b803d948015c993bb58d"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17475260", 0)),
							TransactionHash: common.HexToHash("0x85cfc1b0641425e7aefeef65f9fe274deff13154e14394b0c9fd316746a98d10"),
							Index:           269,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000283af0b28c62c092c9727f1ee09c02ca627eb7f5"),
								common.HexToHash("0x0000000000000000000000007da1e64e92094075a6d9b803d948015c993bb58d"),
								common.HexToHash("0x23913293c04ada44918795ed0efdcb432ed227577f69dfc432395af6d0551173"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("17475260", 0)),
							TransactionHash: common.HexToHash("0x85cfc1b0641425e7aefeef65f9fe274deff13154e14394b0c9fd316746a98d10"),
							Index:           270,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x283Af0B28c62C092C9727F1Ee09c02CA627EB7F5"),
							Topics: []common.Hash{
								common.HexToHash("0xca6abbe9d7f11422cb6ca7629fbf6fe9efb1c621f71ce8f02b9f2a230097404f"),
								common.HexToHash("0x23913293c04ada44918795ed0efdcb432ed227577f69dfc432395af6d0551173"),
								common.HexToHash("0x0000000000000000000000007da1e64e92094075a6d9b803d948015c993bb58d"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000a2e138aa6a1e300000000000000000000000000000000000000000000000000000000666aaccf000000000000000000000000000000000000000000000000000000000000000c64617a75697169616e6764610000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17475260", 0)),
							TransactionHash: common.HexToHash("0x85cfc1b0641425e7aefeef65f9fe274deff13154e14394b0c9fd316746a98d10"),
							Index:           271,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x85cfc1b0641425e7aefeef65f9fe274deff13154e14394b0c9fd316746a98d10"),
						TransactionIndex: 123,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:       "0x85cfc1b0641425e7aefeef65f9fe274deff13154e14394b0c9fd316746a98d10",
				Network:  filter.NetworkEthereum,
				Index:    123,
				From:     "0x7Da1E64e92094075A6D9b803d948015C993Bb58D",
				To:       "0x1d6552e8F46fD509f3918A174FE62C34b42564aE",
				Platform: lo.ToPtr(filter.PlatformENS),
				Type:     filter.TypeCollectibleTrade,
				Calldata: &schema.Calldata{
					FunctionHash: "0xf7a16963",
				},
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("5418018134817597")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeCollectibleTrade,
						Platform: filter.PlatformENS.String(),
						From:     ethereum.AddressGenesis.String(),
						To:       "0x7Da1E64e92094075A6D9b803d948015C993Bb58D",
						Metadata: metadata.CollectibleTrade{
							Action: metadata.ActionCollectibleTradeBuy,
							Token: metadata.Token{
								Address:  lo.ToPtr(ens.AddressBaseRegistrarImplementation.String()),
								ID:       lo.ToPtr(decimal.NewFromBigInt(common.HexToHash("0x23913293c04ada44918795ed0efdcb432ed227577f69dfc432395af6d0551173").Big(), 0)),
								Value:    lo.ToPtr(decimal.NewFromBigInt(big.NewInt(1), 0)),
								Standard: metadata.StandardERC721,
								Symbol:   "ENS",
								URI:      "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/16087491595487441497767710101977939215480045234077591418499274871761116926323",
							},
							Cost: &metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2865411232539107"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1686710135,
			},
			wantError: require.NoError,
		},
		{
			name: "ENS NameRegistered V2",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xa498498d8a10735e50b6a3b1f7b8497bb3a4cfa770dce4719c91a8d5b15c1ad6"),
						ParentHash:   common.HexToHash("0xd2dc925edbf121458e95b1bf52149d20439a8ccbe632dca08981fc35119e3516"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x5124fcC2B3F99F571AD67D075643C743F38f1C34"),
						Number:       lo.Must(new(big.Int).SetString("17476948", 0)),
						GasLimit:     30000000,
						GasUsed:      8629294,
						Timestamp:    1686730691,
						BaseFee:      lo.Must(new(big.Int).SetString("16116025849", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x2bbf9d6f2644c79a8983a532e2488e046c49caeea8b07e6a42f3c9741c16640d"),
						From:      common.HexToAddress("0x4265D230d2D54010d853b107848FC6e0B64c9c24"),
						Gas:       307289,
						GasPrice:  lo.Must(new(big.Int).SetString("16416025849", 10)),
						Hash:      common.HexToHash("0x2bbf9d6f2644c79a8983a532e2488e046c49caeea8b07e6a42f3c9741c16640d"),
						Input:     hexutil.MustDecode("0x74694a2b00000000000000000000000000000000000000000000000000000000000001000000000000000000000000004265d230d2d54010d853b107848fc6e0b64c9c2400000000000000000000000000000000000000000000000000000000096601809923eb94000000039d7f23fefb349133efd9485428adc89c065ca282574a40ae000000000000000000000000231b0ee14048e9dccd1d247744d114a4eb5e8e63000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e616c6c6168756d6d61626172696b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000a48b95dd718f3687d5085fb5971ea0e3d23d25d60a0175eb981d6c65e9e9c9c898640fa888000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000144265d230d2d54010d853b107848fc6e0b64c9c2400000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x253553366Da8546fC250F225fe3d25d0C782303b")),
						Value:     lo.Must(new(big.Int).SetString("14644878074514677", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xa498498d8a10735e50b6a3b1f7b8497bb3a4cfa770dce4719c91a8d5b15c1ad6"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17476948", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 4425493,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3d278acf9"),
						GasUsed:           294623,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000d4416b13d2b3a9abae7acd5d6c2bbdbe25686401"),
								common.HexToHash("0xda8c6e91c89ffabbbc6d14bac2c59c9c8a55fd24063001bbc5cb3a592f82ac2a"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("17476948", 0)),
							TransactionHash: common.HexToHash("0x2bbf9d6f2644c79a8983a532e2488e046c49caeea8b07e6a42f3c9741c16640d"),
							Index:           88,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e"),
							Topics: []common.Hash{
								common.HexToHash("0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82"),
								common.HexToHash("0x93cdeb708b7545dc668eb9280176169d1c33cfd8ed6f04690a0bcc88a93fc4ae"),
								common.HexToHash("0xda8c6e91c89ffabbbc6d14bac2c59c9c8a55fd24063001bbc5cb3a592f82ac2a"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000d4416b13d2b3a9abae7acd5d6c2bbdbe25686401"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17476948", 0)),
							TransactionHash: common.HexToHash("0x2bbf9d6f2644c79a8983a532e2488e046c49caeea8b07e6a42f3c9741c16640d"),
							Index:           89,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85"),
							Topics: []common.Hash{
								common.HexToHash("0xb3d987963d01b2f68493b4bdb130988f157ea43070d4ad840fee0466ed9370d9"),
								common.HexToHash("0xda8c6e91c89ffabbbc6d14bac2c59c9c8a55fd24063001bbc5cb3a592f82ac2a"),
								common.HexToHash("0x000000000000000000000000d4416b13d2b3a9abae7acd5d6c2bbdbe25686401"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000006def7943"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17476948", 0)),
							TransactionHash: common.HexToHash("0x2bbf9d6f2644c79a8983a532e2488e046c49caeea8b07e6a42f3c9741c16640d"),
							Index:           90,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xD4416b13d2b3a9aBae7AcD5D6C2BbDBE25686401"),
							Topics: []common.Hash{
								common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"),
								common.HexToHash("0x000000000000000000000000253553366da8546fc250f225fe3d25d0c782303b"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x0000000000000000000000004265d230d2d54010d853b107848fc6e0b64c9c24"),
							},
							Data:            hexutil.MustDecode("0x8f3687d5085fb5971ea0e3d23d25d60a0175eb981d6c65e9e9c9c898640fa8880000000000000000000000000000000000000000000000000000000000000001"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17476948", 0)),
							TransactionHash: common.HexToHash("0x2bbf9d6f2644c79a8983a532e2488e046c49caeea8b07e6a42f3c9741c16640d"),
							Index:           91,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xD4416b13d2b3a9aBae7AcD5D6C2BbDBE25686401"),
							Topics: []common.Hash{
								common.HexToHash("0x8ce7013e8abebc55c3890a68f5a27c67c3f7efa64e584de5fb22363c606fd340"),
								common.HexToHash("0x8f3687d5085fb5971ea0e3d23d25d60a0175eb981d6c65e9e9c9c898640fa888"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000800000000000000000000000004265d230d2d54010d853b107848fc6e0b64c9c240000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000006e66204300000000000000000000000000000000000000000000000000000000000000140e616c6c6168756d6d61626172696b0365746800000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17476948", 0)),
							TransactionHash: common.HexToHash("0x2bbf9d6f2644c79a8983a532e2488e046c49caeea8b07e6a42f3c9741c16640d"),
							Index:           92,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e"),
							Topics: []common.Hash{
								common.HexToHash("0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0"),
								common.HexToHash("0x8f3687d5085fb5971ea0e3d23d25d60a0175eb981d6c65e9e9c9c898640fa888"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000231b0ee14048e9dccd1d247744d114a4eb5e8e63"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17476948", 0)),
							TransactionHash: common.HexToHash("0x2bbf9d6f2644c79a8983a532e2488e046c49caeea8b07e6a42f3c9741c16640d"),
							Index:           93,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x231b0Ee14048e9dCcD1d247744d114a4EB5E8E63"),
							Topics: []common.Hash{
								common.HexToHash("0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752"),
								common.HexToHash("0x8f3687d5085fb5971ea0e3d23d25d60a0175eb981d6c65e9e9c9c898640fa888"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000144265d230d2d54010d853b107848fc6e0b64c9c24000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17476948", 0)),
							TransactionHash: common.HexToHash("0x2bbf9d6f2644c79a8983a532e2488e046c49caeea8b07e6a42f3c9741c16640d"),
							Index:           94,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x231b0Ee14048e9dCcD1d247744d114a4EB5E8E63"),
							Topics: []common.Hash{
								common.HexToHash("0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2"),
								common.HexToHash("0x8f3687d5085fb5971ea0e3d23d25d60a0175eb981d6c65e9e9c9c898640fa888"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000004265d230d2d54010d853b107848fc6e0b64c9c24"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17476948", 0)),
							TransactionHash: common.HexToHash("0x2bbf9d6f2644c79a8983a532e2488e046c49caeea8b07e6a42f3c9741c16640d"),
							Index:           95,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x253553366Da8546fC250F225fe3d25d0C782303b"),
							Topics: []common.Hash{
								common.HexToHash("0x69e37f151eb98a09618ddaa80c8cfaf1ce5996867c489f45b555b412271ebf27"),
								common.HexToHash("0xda8c6e91c89ffabbbc6d14bac2c59c9c8a55fd24063001bbc5cb3a592f82ac2a"),
								common.HexToHash("0x0000000000000000000000004265d230d2d54010d853b107848fc6e0b64c9c24"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000330246071bb25a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006def7943000000000000000000000000000000000000000000000000000000000000000e616c6c6168756d6d61626172696b000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17476948", 0)),
							TransactionHash: common.HexToHash("0x2bbf9d6f2644c79a8983a532e2488e046c49caeea8b07e6a42f3c9741c16640d"),
							Index:           96,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x2bbf9d6f2644c79a8983a532e2488e046c49caeea8b07e6a42f3c9741c16640d"),
						TransactionIndex: 68,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x2bbf9d6f2644c79a8983a532e2488e046c49caeea8b07e6a42f3c9741c16640d",
				Network: filter.NetworkEthereum,
				Index:   68,
				From:    "0x4265D230d2D54010d853b107848FC6e0B64c9c24",
				To:      "0x253553366Da8546fC250F225fe3d25d0C782303b",
				Type:    filter.TypeCollectibleTrade,
				Calldata: &schema.Calldata{
					FunctionHash: "0x74694a2b",
				},
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("4836538783709927")),
					Decimal: 18,
				},
				Platform: lo.ToPtr(filter.PlatformENS),
				Actions: []*schema.Action{
					{
						Type:     filter.TypeCollectibleTrade,
						Platform: filter.PlatformENS.String(),
						From:     ethereum.AddressGenesis.String(),
						To:       "0x4265D230d2D54010d853b107848FC6e0B64c9c24",
						Metadata: metadata.CollectibleTrade{
							Action: metadata.ActionCollectibleTradeBuy,
							Token: metadata.Token{
								Address:  lo.ToPtr(ens.AddressBaseRegistrarImplementation.String()),
								ID:       lo.ToPtr(decimal.NewFromBigInt(common.HexToHash("0xda8c6e91c89ffabbbc6d14bac2c59c9c8a55fd24063001bbc5cb3a592f82ac2a").Big(), 0)),
								Value:    lo.ToPtr(decimal.NewFromBigInt(big.NewInt(1), 0)),
								Standard: metadata.StandardERC721,
								Symbol:   "ENS",
								URI:      "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/98852322702639730223199955634700678434407562961828698565072639218093587082282",
							},
							Cost: &metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("14357723602465370"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1686730691,
			},
			wantError: require.NoError,
		},
		{
			name: "ENS NameRenewed V1",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xc06cd55ff53bff781a54917734bec5759859b4ab8e9b52e5745a945d7706f1b0"),
						ParentHash:   common.HexToHash("0x8d6016e8559f1efa08a0a7e407ae39728a8fe136ded31eb5013b804d4bf5ae6c"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"),
						Number:       lo.Must(new(big.Int).SetString("17476816", 0)),
						GasLimit:     30000000,
						GasUsed:      29043463,
						Timestamp:    1686729107,
						BaseFee:      lo.Must(new(big.Int).SetString("14378482122", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x228e0614701d5c365928387ad5ef0fa61d517eed1efe658827275f870674a8a4"),
						From:      common.HexToAddress("0xBaDacd66C71448047A3ebf0314DDb5dba046FF53"),
						Gas:       148738,
						GasPrice:  lo.Must(new(big.Int).SetString("14478482122", 10)),
						Hash:      common.HexToHash("0x228e0614701d5c365928387ad5ef0fa61d517eed1efe658827275f870674a8a4"),
						Input:     hexutil.MustDecode("0xe8d6dbb400000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000005a490080000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000083333323232323232000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xfF252725f6122A92551A5FA9a6b6bf10eb0Be035")),
						Value:     lo.Must(new(big.Int).SetString("9489688397133431", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xc06cd55ff53bff781a54917734bec5759859b4ab8e9b52e5745a945d7706f1b0"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17476816", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 24443489,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x35efc1aca"),
						GasUsed:           125491,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85"),
							Topics: []common.Hash{
								common.HexToHash("0x9b87a00e30f1ac65d898f070f8a3488fe60517182d0a2098e1b4b93a54aa9bd6"),
								common.HexToHash("0x5616336fd4453779836ba79a8886448a8c1ac025f588cf378c1e16648c996e9a"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000006a37385d"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17476816", 0)),
							TransactionHash: common.HexToHash("0x228e0614701d5c365928387ad5ef0fa61d517eed1efe658827275f870674a8a4"),
							Index:           255,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x283Af0B28c62C092C9727F1Ee09c02CA627EB7F5"),
							Topics: []common.Hash{
								common.HexToHash("0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae"),
								common.HexToHash("0x5616336fd4453779836ba79a8886448a8c1ac025f588cf378c1e16648c996e9a"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000001ea63381c38e9b000000000000000000000000000000000000000000000000000000006a37385d00000000000000000000000000000000000000000000000000000000000000083333323232323232000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17476816", 0)),
							TransactionHash: common.HexToHash("0x228e0614701d5c365928387ad5ef0fa61d517eed1efe658827275f870674a8a4"),
							Index:           256,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x228e0614701d5c365928387ad5ef0fa61d517eed1efe658827275f870674a8a4"),
						TransactionIndex: 64,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x228e0614701d5c365928387ad5ef0fa61d517eed1efe658827275f870674a8a4",
				Network: filter.NetworkEthereum,
				Index:   64,
				From:    "0xBaDacd66C71448047A3ebf0314DDb5dba046FF53",
				To:      "0xfF252725f6122A92551A5FA9a6b6bf10eb0Be035",
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("1816919199971902")),
					Decimal: 18,
				},
				Type: filter.TypeSocialProfile,
				Calldata: &schema.Calldata{
					FunctionHash: "0xe8d6dbb4",
				},
				Platform: lo.ToPtr(filter.PlatformENS),
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformENS.String(),
						From:     "0xBaDacd66C71448047A3ebf0314DDb5dba046FF53",
						To:       "0x283Af0B28c62C092C9727F1Ee09c02CA627EB7F5",
						Metadata: metadata.SocialProfile{
							Action:    metadata.ActionSocialProfileRenew,
							ProfileID: "38938130617558480159522967237952523372715276171406776489065168147179709755034",
							Handle:    "33222222.eth",
							Address:   ens.AddressBaseRegistrarImplementation,
							ImageURI:  "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/38938130617558480159522967237952523372715276171406776489065168147179709755034/image",
							Expiry:    lo.ToPtr(time.Unix(1782003805, 0)),
						},
					},
				},
				Status:    true,
				Timestamp: 1686729107,
			},
			wantError: require.NoError,
		},
		{
			name: "ENS NameRenewed V2",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x171d24a3a17e753fa7941294775bb4d89800729f8ddb8a863a5e93632a7a29bd"),
						ParentHash:   common.HexToHash("0xe823957d6a56d7f3357fc049818703bf6ee72e176e52b62a608fed7c0e900e74"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"),
						Number:       lo.Must(new(big.Int).SetString("17476895", 0)),
						GasLimit:     30000000,
						GasUsed:      13590017,
						Timestamp:    1686730055,
						BaseFee:      lo.Must(new(big.Int).SetString("16029216215", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xeb861ac3443943528e05530281e168abac3cec9d95583ee77a4df2d3e81b14d8"),
						From:      common.HexToAddress("0x6F4644485226276868658cC467700e104f2f9689"),
						Gas:       93264,
						GasPrice:  lo.Must(new(big.Int).SetString("16129216215", 10)),
						Hash:      common.HexToHash("0xeb861ac3443943528e05530281e168abac3cec9d95583ee77a4df2d3e81b14d8"),
						Input:     hexutil.MustDecode("0xacf1a84100000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000001e133800000000000000000000000000000000000000000000000000000000000000009736372616c746163680000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x253553366Da8546fC250F225fe3d25d0C782303b")),
						Value:     lo.Must(new(big.Int).SetString("2928975614902935", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x171d24a3a17e753fa7941294775bb4d89800729f8ddb8a863a5e93632a7a29bd"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17476895", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 7447388,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3c1604ed7"),
						GasUsed:           93264,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85"),
							Topics: []common.Hash{
								common.HexToHash("0x9b87a00e30f1ac65d898f070f8a3488fe60517182d0a2098e1b4b93a54aa9bd6"),
								common.HexToHash("0x5fc51659cd7ea1a38b30300d29d906e9ffb797cec104038bc653adcd796e26b1"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000665bdcc0"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17476895", 0)),
							TransactionHash: common.HexToHash("0xeb861ac3443943528e05530281e168abac3cec9d95583ee77a4df2d3e81b14d8"),
							Index:           176,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x253553366Da8546fC250F225fe3d25d0C782303b"),
							Topics: []common.Hash{
								common.HexToHash("0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae"),
								common.HexToHash("0x5fc51659cd7ea1a38b30300d29d906e9ffb797cec104038bc653adcd796e26b1"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000a67e34715029700000000000000000000000000000000000000000000000000000000665bdcc00000000000000000000000000000000000000000000000000000000000000009736372616c746163680000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17476895", 0)),
							TransactionHash: common.HexToHash("0xeb861ac3443943528e05530281e168abac3cec9d95583ee77a4df2d3e81b14d8"),
							Index:           177,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xeb861ac3443943528e05530281e168abac3cec9d95583ee77a4df2d3e81b14d8"),
						TransactionIndex: 107,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0xeb861ac3443943528e05530281e168abac3cec9d95583ee77a4df2d3e81b14d8",
				Network: filter.NetworkEthereum,
				Index:   107,
				From:    "0x6F4644485226276868658cC467700e104f2f9689",
				To:      "0x253553366Da8546fC250F225fe3d25d0C782303b",
				Type:    filter.TypeSocialProfile,
				Calldata: &schema.Calldata{
					FunctionHash: "0xacf1a841",
				},
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("1504275221075760")),
					Decimal: 18,
				},
				Platform: lo.ToPtr(filter.PlatformENS),
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformENS.String(),
						From:     "0x6F4644485226276868658cC467700e104f2f9689",
						To:       "0x253553366Da8546fC250F225fe3d25d0C782303b",
						Metadata: metadata.SocialProfile{
							Action:    metadata.ActionSocialProfileRenew,
							ProfileID: "43317943746667644306419405492250733837995152364233006308340930267579834836657",
							Handle:    "scraltach.eth",
							Address:   ens.AddressBaseRegistrarImplementation,
							ImageURI:  "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/43317943746667644306419405492250733837995152364233006308340930267579834836657/image",
							Expiry:    lo.ToPtr(time.Unix(1717296320, 0)),
						},
					},
				},
				Status:    true,
				Timestamp: 1686730055,
			},
			wantError: require.NoError,
		},
		{
			name: "ENS Public Resolver TextChanged",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xdc046f63ff0e6943715d065c87ce5ff281d90cde0ba15eef20fc084408eb1787"),
						ParentHash:   common.HexToHash("0x2a6b5eb5f3b1453547f0e0abe819cfdf51f3dda15fadc4678745317e2958d8f6"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("17477413", 0)),
						GasLimit:     30000000,
						GasUsed:      13448631,
						Timestamp:    1686736367,
						BaseFee:      lo.Must(new(big.Int).SetString("13876010942", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x452a0fa17e761353f1a311fb01a53a1a573f06b2348955a383d5e7a968aee9d6"),
						From:      common.HexToAddress("0xC4eE38B534CfbD26cB94e282A390eCa0B7e3e7b2"),
						Gas:       114983,
						GasPrice:  lo.Must(new(big.Int).SetString("13976010942", 10)),
						Hash:      common.HexToHash("0x452a0fa17e761353f1a311fb01a53a1a573f06b2348955a383d5e7a968aee9d6"),
						Input:     hexutil.MustDecode("0xac9650d80000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000002a000000000000000000000000000000000000000000000000000000000000003a0000000000000000000000000000000000000000000000000000000000000048000000000000000000000000000000000000000000000000000000000000000c410f13a8c602ef79a9a2a460da2ce5136e8a1000b242b1c672103af0a1aa63e440e23d0d2000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000375726c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c410f13a8c602ef79a9a2a460da2ce5136e8a1000b242b1c672103af0a1aa63e440e23d0d2000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000b636f6d2e7477697474657200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c410f13a8c602ef79a9a2a460da2ce5136e8a1000b242b1c672103af0a1aa63e440e23d0d2000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000005656d61696c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a48b95dd71602ef79a9a2a460da2ce5136e8a1000b242b1c672103af0a1aa63e440e23d0d2000000000000000000000000000000000000000000000000000000000000003c0000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000848b95dd71602ef79a9a2a460da2ce5136e8a1000b242b1c672103af0a1aa63e440e23d0d200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xdc046f63ff0e6943715d065c87ce5ff281d90cde0ba15eef20fc084408eb1787"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17477413", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 8619996,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3410900be"),
						GasUsed:           100583,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41"),
							Topics: []common.Hash{
								common.HexToHash("0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550"),
								common.HexToHash("0x602ef79a9a2a460da2ce5136e8a1000b242b1c672103af0a1aa63e440e23d0d2"),
								common.HexToHash("0xb68b5f5089998f2978a1dcc681e8ef27962b90d5c26c4c0b9c1945814ffa5ef0"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000375726c0000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17477413", 0)),
							TransactionHash: common.HexToHash("0x452a0fa17e761353f1a311fb01a53a1a573f06b2348955a383d5e7a968aee9d6"),
							Index:           192,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41"),
							Topics: []common.Hash{
								common.HexToHash("0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550"),
								common.HexToHash("0x602ef79a9a2a460da2ce5136e8a1000b242b1c672103af0a1aa63e440e23d0d2"),
								common.HexToHash("0x7b8e740dad11ad11abceebf798f458a149d945f7c4448784733725f3fa21225e"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000b636f6d2e74776974746572000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17477413", 0)),
							TransactionHash: common.HexToHash("0x452a0fa17e761353f1a311fb01a53a1a573f06b2348955a383d5e7a968aee9d6"),
							Index:           193,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41"),
							Topics: []common.Hash{
								common.HexToHash("0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550"),
								common.HexToHash("0x602ef79a9a2a460da2ce5136e8a1000b242b1c672103af0a1aa63e440e23d0d2"),
								common.HexToHash("0xa28419e839eda1f8810510d0b578872b1165326edcf6e8575613087bd6b77e35"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000005656d61696c000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17477413", 0)),
							TransactionHash: common.HexToHash("0x452a0fa17e761353f1a311fb01a53a1a573f06b2348955a383d5e7a968aee9d6"),
							Index:           194,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41"),
							Topics: []common.Hash{
								common.HexToHash("0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752"),
								common.HexToHash("0x602ef79a9a2a460da2ce5136e8a1000b242b1c672103af0a1aa63e440e23d0d2"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17477413", 0)),
							TransactionHash: common.HexToHash("0x452a0fa17e761353f1a311fb01a53a1a573f06b2348955a383d5e7a968aee9d6"),
							Index:           195,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41"),
							Topics: []common.Hash{
								common.HexToHash("0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2"),
								common.HexToHash("0x602ef79a9a2a460da2ce5136e8a1000b242b1c672103af0a1aa63e440e23d0d2"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17477413", 0)),
							TransactionHash: common.HexToHash("0x452a0fa17e761353f1a311fb01a53a1a573f06b2348955a383d5e7a968aee9d6"),
							Index:           196,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41"),
							Topics: []common.Hash{
								common.HexToHash("0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752"),
								common.HexToHash("0x602ef79a9a2a460da2ce5136e8a1000b242b1c672103af0a1aa63e440e23d0d2"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17477413", 0)),
							TransactionHash: common.HexToHash("0x452a0fa17e761353f1a311fb01a53a1a573f06b2348955a383d5e7a968aee9d6"),
							Index:           197,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x452a0fa17e761353f1a311fb01a53a1a573f06b2348955a383d5e7a968aee9d6"),
						TransactionIndex: 103,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x452a0fa17e761353f1a311fb01a53a1a573f06b2348955a383d5e7a968aee9d6",
				Network: filter.NetworkEthereum,
				Index:   103,
				From:    "0xC4eE38B534CfbD26cB94e282A390eCa0B7e3e7b2",
				To:      "0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41",
				Type:    filter.TypeSocialProfile,
				Calldata: &schema.Calldata{
					FunctionHash: "0xac9650d8",
				},
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("1405749108579186")),
					Decimal: 18,
				},
				Platform: lo.ToPtr(filter.PlatformENS),
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformENS.String(),
						From:     "0xC4eE38B534CfbD26cB94e282A390eCa0B7e3e7b2",
						To:       "0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41",
						Metadata: metadata.SocialProfile{
							Action:    metadata.ActionSocialProfileUpdate,
							ProfileID: "84724132240243453947957943856441614830553665754416461585799729112088148679789",
							Handle:    "cryptopaycheck.eth",
							Address:   ens.AddressBaseRegistrarImplementation,
							ImageURI:  "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/84724132240243453947957943856441614830553665754416461585799729112088148679789/image",
							Key:       "url",
						},
					},
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformENS.String(),
						From:     "0xC4eE38B534CfbD26cB94e282A390eCa0B7e3e7b2",
						To:       "0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41",
						Metadata: metadata.SocialProfile{
							Action:    metadata.ActionSocialProfileUpdate,
							ProfileID: "84724132240243453947957943856441614830553665754416461585799729112088148679789",
							Handle:    "cryptopaycheck.eth",
							Address:   ens.AddressBaseRegistrarImplementation,
							ImageURI:  "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/84724132240243453947957943856441614830553665754416461585799729112088148679789/image",
							Key:       "com.twitter",
						},
					},
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformENS.String(),
						From:     "0xC4eE38B534CfbD26cB94e282A390eCa0B7e3e7b2",
						To:       "0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41",
						Metadata: metadata.SocialProfile{
							Action:    metadata.ActionSocialProfileUpdate,
							ProfileID: "84724132240243453947957943856441614830553665754416461585799729112088148679789",
							Handle:    "cryptopaycheck.eth",
							Address:   ens.AddressBaseRegistrarImplementation,
							ImageURI:  "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/84724132240243453947957943856441614830553665754416461585799729112088148679789/image",
							Key:       "email",
						},
					},
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformENS.String(),
						From:     "0xC4eE38B534CfbD26cB94e282A390eCa0B7e3e7b2",
						To:       "0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41",
						Metadata: metadata.SocialProfile{
							Action:    metadata.ActionSocialProfileUpdate,
							ProfileID: "84724132240243453947957943856441614830553665754416461585799729112088148679789",
							Handle:    "cryptopaycheck.eth",
							Address:   ens.AddressBaseRegistrarImplementation,
							ImageURI:  "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/84724132240243453947957943856441614830553665754416461585799729112088148679789/image",
							Key:       "Address",
							Value:     "0x0000000000000000000000000000000000000000",
						},
					},
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformENS.String(),
						From:     "0xC4eE38B534CfbD26cB94e282A390eCa0B7e3e7b2",
						To:       "0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41",
						Metadata: metadata.SocialProfile{
							Action:    metadata.ActionSocialProfileUpdate,
							ProfileID: "84724132240243453947957943856441614830553665754416461585799729112088148679789",
							Handle:    "cryptopaycheck.eth",
							Address:   ens.AddressBaseRegistrarImplementation,
							ImageURI:  "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/84724132240243453947957943856441614830553665754416461585799729112088148679789/image",
							Key:       "Address",
							Value:     "0x0000000000000000000000000000000000000000",
						},
					},
				},
				Status:    true,
				Timestamp: 1686736367,
			},
			wantError: require.NoError,
		},
		{
			name: "ENS Public Resolver TextChanged With Value",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xd977d57f212b6e3e10dcdeffd445ec64b349c9bd854c76d92c6ecd5d617ba561"),
						ParentHash:   common.HexToHash("0x132123fb94b39c474fb7f6c49f8a7d420f54e60b25fdddb47485b556a44c4da8"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xDAFEA492D9c6733ae3d56b7Ed1ADB60692c98Bc5"),
						Number:       lo.Must(new(big.Int).SetString("17134061", 0)),
						GasLimit:     30000000,
						GasUsed:      17549824,
						Timestamp:    1682556563,
						BaseFee:      lo.Must(new(big.Int).SetString("33195723988", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x7364fd343c669e29af6710ebc55c4e6588e8acb7d2b63afe2a00f09f1528ada8"),
						From:      common.HexToAddress("0xA60e522c5517B05526eE0F7f3885b82b37CeeB2d"),
						Gas:       106531,
						GasPrice:  lo.Must(new(big.Int).SetString("33282173401", 10)),
						Hash:      common.HexToHash("0x7364fd343c669e29af6710ebc55c4e6588e8acb7d2b63afe2a00f09f1528ada8"),
						Input:     hexutil.MustDecode("0xac9650d8000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000010410f13a8c995b39dc62bd369ecd91aa1ea7f466b0a5b3ec55d4c24e8deafccb7023d9164a000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000066176617461720000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002468747470733a2f2f656e732e78797a2f706570657472696c6c696f6e616972652e6574680000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x231b0Ee14048e9dCcD1d247744d114a4EB5E8E63")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xd977d57f212b6e3e10dcdeffd445ec64b349c9bd854c76d92c6ecd5d617ba561"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17134061", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 13704628,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x7bfc5a9d9"),
						GasUsed:           106038,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x231b0Ee14048e9dCcD1d247744d114a4EB5E8E63"),
							Topics: []common.Hash{
								common.HexToHash("0x448bc014f1536726cf8d54ff3d6481ed3cbc683c2591ca204274009afa09b1a1"),
								common.HexToHash("0x995b39dc62bd369ecd91aa1ea7f466b0a5b3ec55d4c24e8deafccb7023d9164a"),
								common.HexToHash("0xd1f86c93d831119ad98fe983e643a7431e4ac992e3ead6e3007f4dd1adf66343"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000066176617461720000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002468747470733a2f2f656e732e78797a2f706570657472696c6c696f6e616972652e65746800000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17134061", 0)),
							TransactionHash: common.HexToHash("0x7364fd343c669e29af6710ebc55c4e6588e8acb7d2b63afe2a00f09f1528ada8"),
							Index:           249,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x7364fd343c669e29af6710ebc55c4e6588e8acb7d2b63afe2a00f09f1528ada8"),
						TransactionIndex: 124,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x7364fd343c669e29af6710ebc55c4e6588e8acb7d2b63afe2a00f09f1528ada8",
				Network: filter.NetworkEthereum,
				Index:   124,
				From:    "0xA60e522c5517B05526eE0F7f3885b82b37CeeB2d",
				To:      "0x231b0Ee14048e9dCcD1d247744d114a4EB5E8E63",
				Type:    filter.TypeSocialProfile,
				Calldata: &schema.Calldata{
					FunctionHash: "0xac9650d8",
				},
				Platform: lo.ToPtr(filter.PlatformENS),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("3529175103095238")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformENS.String(),
						From:     "0xA60e522c5517B05526eE0F7f3885b82b37CeeB2d",
						To:       "0x231b0Ee14048e9dCcD1d247744d114a4EB5E8E63",
						Metadata: metadata.SocialProfile{
							Action:    metadata.ActionSocialProfileUpdate,
							ProfileID: "82212315406362059018592037832286297744830726028007896928355368193891050262687",
							Handle:    "pepetrillionaire.eth",
							Address:   ens.AddressBaseRegistrarImplementation,
							ImageURI:  "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/82212315406362059018592037832286297744830726028007896928355368193891050262687/image",
							Key:       "avatar",
							Value:     "https://ens.xyz/pepetrillionaire.eth",
						},
					},
				},
				Status:    true,
				Timestamp: 1682556563,
			},
			wantError: require.NoError,
		},
		{
			name: "ENS Name Wrapped",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xbe094ed33b2dfeaecbe6695e0c1a557229643856ea8dde52aba1c51b9cc809f7"),
						ParentHash:   common.HexToHash("0xe5703a7034ffd533ffa4a21540ba7679684c47ef91ed9381e567ee00ad719129"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xDAFEA492D9c6733ae3d56b7Ed1ADB60692c98Bc5"),
						Number:       lo.Must(new(big.Int).SetString("17932511", 0)),
						GasLimit:     30000000,
						GasUsed:      9204557,
						Timestamp:    1692250823,
						BaseFee:      lo.Must(new(big.Int).SetString("16316428969", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x5597682570383f1a57a82b3b77673a4561d472d0fdc6ba324d8e687e789c9df9"),
						From:      common.HexToAddress("0xC59dc5B9906728A19070BeD06F10E31da2313AC6"),
						Gas:       136437,
						GasPrice:  lo.Must(new(big.Int).SetString("16416428969", 10)),
						Hash:      common.HexToHash("0x5597682570383f1a57a82b3b77673a4561d472d0fdc6ba324d8e687e789c9df9"),
						Input:     hexutil.MustDecode("0xb88d4fde000000000000000000000000c59dc5b9906728a19070bed06f10e31da2313ac6000000000000000000000000d4416b13d2b3a9abae7acd5d6c2bbdbe25686401461f3abf83be61e334dffeb9e05f99bf9561cb5df2b9dc5f03b9600f0ad02f23000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000080000000000000000000000000c59dc5b9906728a19070bed06f10e31da2313ac60000000000000000000000000000000000000000000000000000000000000000000000000000000000000000231b0ee14048e9dccd1d247744d114a4eb5e8e630000000000000000000000000000000000000000000000000000000000000005302d373730000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xbe094ed33b2dfeaecbe6695e0c1a557229643856ea8dde52aba1c51b9cc809f7"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17932511", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 7184084,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3d27ed3a9"),
						GasUsed:           135324,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000c59dc5b9906728a19070bed06f10e31da2313ac6"),
								common.HexToHash("0x000000000000000000000000d4416b13d2b3a9abae7acd5d6c2bbdbe25686401"),
								common.HexToHash("0x461f3abf83be61e334dffeb9e05f99bf9561cb5df2b9dc5f03b9600f0ad02f23"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("17932511", 0)),
							TransactionHash: common.HexToHash("0x5597682570383f1a57a82b3b77673a4561d472d0fdc6ba324d8e687e789c9df9"),
							Index:           175,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e"),
							Topics: []common.Hash{
								common.HexToHash("0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82"),
								common.HexToHash("0x93cdeb708b7545dc668eb9280176169d1c33cfd8ed6f04690a0bcc88a93fc4ae"),
								common.HexToHash("0x461f3abf83be61e334dffeb9e05f99bf9561cb5df2b9dc5f03b9600f0ad02f23"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000d4416b13d2b3a9abae7acd5d6c2bbdbe25686401"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17932511", 0)),
							TransactionHash: common.HexToHash("0x5597682570383f1a57a82b3b77673a4561d472d0fdc6ba324d8e687e789c9df9"),
							Index:           176,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xD4416b13d2b3a9aBae7AcD5D6C2BbDBE25686401"),
							Topics: []common.Hash{
								common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"),
								common.HexToHash("0x00000000000000000000000057f1887a8bf19b14fc0df6fd9b2acc9af147ea85"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000c59dc5b9906728a19070bed06f10e31da2313ac6"),
							},
							Data:            hexutil.MustDecode("0xb9a59c0f1c96d0c72d582725718e72fc107328ad0ddba6ef735bc2fce9d85d920000000000000000000000000000000000000000000000000000000000000001"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17932511", 0)),
							TransactionHash: common.HexToHash("0x5597682570383f1a57a82b3b77673a4561d472d0fdc6ba324d8e687e789c9df9"),
							Index:           177,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xD4416b13d2b3a9aBae7AcD5D6C2BbDBE25686401"),
							Topics: []common.Hash{
								common.HexToHash("0x8ce7013e8abebc55c3890a68f5a27c67c3f7efa64e584de5fb22363c606fd340"),
								common.HexToHash("0xb9a59c0f1c96d0c72d582725718e72fc107328ad0ddba6ef735bc2fce9d85d92"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000c59dc5b9906728a19070bed06f10e31da2313ac600000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000067447f18000000000000000000000000000000000000000000000000000000000000000b05302d3737300365746800000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17932511", 0)),
							TransactionHash: common.HexToHash("0x5597682570383f1a57a82b3b77673a4561d472d0fdc6ba324d8e687e789c9df9"),
							Index:           178,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e"),
							Topics: []common.Hash{
								common.HexToHash("0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0"),
								common.HexToHash("0xb9a59c0f1c96d0c72d582725718e72fc107328ad0ddba6ef735bc2fce9d85d92"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000231b0ee14048e9dccd1d247744d114a4eb5e8e63"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17932511", 0)),
							TransactionHash: common.HexToHash("0x5597682570383f1a57a82b3b77673a4561d472d0fdc6ba324d8e687e789c9df9"),
							Index:           179,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x5597682570383f1a57a82b3b77673a4561d472d0fdc6ba324d8e687e789c9df9"),
						TransactionIndex: 64,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x5597682570383f1a57a82b3b77673a4561d472d0fdc6ba324d8e687e789c9df9",
				Network: filter.NetworkEthereum,
				Index:   64,
				From:    "0xC59dc5B9906728A19070BeD06F10E31da2313AC6",
				To:      "0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85",
				Type:    filter.TypeSocialProfile,
				Calldata: &schema.Calldata{
					FunctionHash: "0xb88d4fde",
				},
				Platform: lo.ToPtr(filter.PlatformENS),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("2221536833800956")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformENS.String(),
						From:     "0xC59dc5B9906728A19070BeD06F10E31da2313AC6",
						To:       "0xD4416b13d2b3a9aBae7AcD5D6C2BbDBE25686401",
						Metadata: metadata.SocialProfile{
							Action:    metadata.ActionSocialProfileWrap,
							ProfileID: "31717077124348933353064271983281360053420871390651458975650880342288344559395",
							Handle:    "0-770.eth",
							Address:   ens.AddressBaseRegistrarImplementation,
							ImageURI:  "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/31717077124348933353064271983281360053420871390651458975650880342288344559395/image",
						},
					},
				},
				Status:    true,
				Timestamp: 1692250823,
			},
			wantError: require.NoError,
		},
		{
			name: "ENS Name Unwrapped",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x3de45fda3a4b278c4e5ee85a91a3cf5a2888d1788aac562ae028a8fb7df33a01"),
						ParentHash:   common.HexToHash("0xa479946603a37092cced087556ccd913cdc5ba48400ce4d7b133495c22ecf7a6"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("17919301", 0)),
						GasLimit:     30000000,
						GasUsed:      13339216,
						Timestamp:    1692091115,
						BaseFee:      lo.Must(new(big.Int).SetString("14232372221", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xe2e6f42795b4bbff284d4d68b68e9099ddb7dcb4dcdbb21add936f0e63e01fa7"),
						From:      common.HexToAddress("0x4015e9865cb268E7939220edfbbf623C6A41DaC2"),
						Gas:       90790,
						GasPrice:  lo.Must(new(big.Int).SetString("14332372221", 10)),
						Hash:      common.HexToHash("0xe2e6f42795b4bbff284d4d68b68e9099ddb7dcb4dcdbb21add936f0e63e01fa7"),
						Input:     hexutil.MustDecode("0x8b4dfa753bfe92f24206a87e9e9d273d8903e907a8d16512a9576a5fdc916c3c60387ad90000000000000000000000004015e9865cb268e7939220edfbbf623c6a41dac20000000000000000000000004015e9865cb268e7939220edfbbf623c6a41dac2"),
						To:        lo.ToPtr(common.HexToAddress("0xD4416b13d2b3a9aBae7AcD5D6C2BbDBE25686401")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x3de45fda3a4b278c4e5ee85a91a3cf5a2888d1788aac562ae028a8fb7df33a01"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17919301", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 6540181,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x35646a4fd"),
						GasUsed:           90185,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xD4416b13d2b3a9aBae7AcD5D6C2BbDBE25686401"),
							Topics: []common.Hash{
								common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"),
								common.HexToHash("0x0000000000000000000000004015e9865cb268e7939220edfbbf623c6a41dac2"),
								common.HexToHash("0x0000000000000000000000004015e9865cb268e7939220edfbbf623c6a41dac2"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0xa39a56400e2f4e46ac8bae2f36e167d15321abddcaa11c6b7f191754674b32610000000000000000000000000000000000000000000000000000000000000001"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17919301", 0)),
							TransactionHash: common.HexToHash("0xe2e6f42795b4bbff284d4d68b68e9099ddb7dcb4dcdbb21add936f0e63e01fa7"),
							Index:           207,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e"),
							Topics: []common.Hash{
								common.HexToHash("0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266"),
								common.HexToHash("0xa39a56400e2f4e46ac8bae2f36e167d15321abddcaa11c6b7f191754674b3261"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000004015e9865cb268e7939220edfbbf623c6a41dac2"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17919301", 0)),
							TransactionHash: common.HexToHash("0xe2e6f42795b4bbff284d4d68b68e9099ddb7dcb4dcdbb21add936f0e63e01fa7"),
							Index:           208,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xD4416b13d2b3a9aBae7AcD5D6C2BbDBE25686401"),
							Topics: []common.Hash{
								common.HexToHash("0xee2ba1195c65bcf218a83d874335c6bf9d9067b4c672f3c3bf16cf40de7586c4"),
								common.HexToHash("0xa39a56400e2f4e46ac8bae2f36e167d15321abddcaa11c6b7f191754674b3261"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000004015e9865cb268e7939220edfbbf623c6a41dac2"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17919301", 0)),
							TransactionHash: common.HexToHash("0xe2e6f42795b4bbff284d4d68b68e9099ddb7dcb4dcdbb21add936f0e63e01fa7"),
							Index:           209,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000d4416b13d2b3a9abae7acd5d6c2bbdbe25686401"),
								common.HexToHash("0x0000000000000000000000004015e9865cb268e7939220edfbbf623c6a41dac2"),
								common.HexToHash("0x3bfe92f24206a87e9e9d273d8903e907a8d16512a9576a5fdc916c3c60387ad9"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("17919301", 0)),
							TransactionHash: common.HexToHash("0xe2e6f42795b4bbff284d4d68b68e9099ddb7dcb4dcdbb21add936f0e63e01fa7"),
							Index:           210,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xe2e6f42795b4bbff284d4d68b68e9099ddb7dcb4dcdbb21add936f0e63e01fa7"),
						TransactionIndex: 84,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0xe2e6f42795b4bbff284d4d68b68e9099ddb7dcb4dcdbb21add936f0e63e01fa7",
				Network: filter.NetworkEthereum,
				Index:   84,
				From:    "0x4015e9865cb268E7939220edfbbf623C6A41DaC2",
				To:      "0xD4416b13d2b3a9aBae7AcD5D6C2BbDBE25686401",
				Type:    filter.TypeSocialProfile,
				Calldata: &schema.Calldata{
					FunctionHash: "0x8b4dfa75",
				},
				Platform: lo.ToPtr(filter.PlatformENS),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("1292564988750885")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformENS.String(),
						From:     "0x4015e9865cb268E7939220edfbbf623C6A41DaC2",
						To:       "0xD4416b13d2b3a9aBae7AcD5D6C2BbDBE25686401",
						Metadata: metadata.SocialProfile{
							Action:    metadata.ActionSocialProfileUnwrap,
							ProfileID: "27136251407093501673864290121767055139262057508067012720705335482940619520729",
							Handle:    "notantoine.eth",
							Address:   ens.AddressBaseRegistrarImplementation,
							ImageURI:  "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/27136251407093501673864290121767055139262057508067012720705335482940619520729/image",
						},
					},
				},
				Status:    true,
				Timestamp: 1692091115,
			},
			wantError: require.NoError,
		},
		{
			name: "ENS FusesSet",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xbca740d04a185ceb2dca75a08c0c9610aa086a6048cfbff594885e03f7f56590"),
						ParentHash:   common.HexToHash("0x8064f75f2a43afb795944e6fae909f1e6997dd424088c9d6475b783525da9908"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("17914181", 0)),
						GasLimit:     29970705,
						GasUsed:      19647045,
						Timestamp:    1692029303,
						BaseFee:      lo.Must(new(big.Int).SetString("41539233483", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xaa33107f3cf26f828a955970f7765724aadf71fe0f8a77a494ab92130cfc649c"),
						From:      common.HexToAddress("0x6d5601E90220C989111939d9317FCbba27c015ab"),
						Gas:       30414,
						GasPrice:  lo.Must(new(big.Int).SetString("41639233483", 10)),
						Hash:      common.HexToHash("0xaa33107f3cf26f828a955970f7765724aadf71fe0f8a77a494ab92130cfc649c"),
						Input:     hexutil.MustDecode("0x402906fc41c498a8085c28662e7bf05fefda3167d23b3e8ab7914a6ee41ea1e90f38b6c90000000000000000000000000000000000000000000000000000000000000001"),
						To:        lo.ToPtr(common.HexToAddress("0xD4416b13d2b3a9aBae7AcD5D6C2BbDBE25686401")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xbca740d04a185ceb2dca75a08c0c9610aa086a6048cfbff594885e03f7f56590"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17914181", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 17910919,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x9b1e447cb"),
						GasUsed:           30414,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xD4416b13d2b3a9aBae7AcD5D6C2BbDBE25686401"),
							Topics: []common.Hash{
								common.HexToHash("0x39873f00c80f4f94b7bd1594aebcf650f003545b74824d57ddf4939e3ff3a34b"),
								common.HexToHash("0x41c498a8085c28662e7bf05fefda3167d23b3e8ab7914a6ee41ea1e90f38b6c9"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000030001"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17914181", 0)),
							TransactionHash: common.HexToHash("0xaa33107f3cf26f828a955970f7765724aadf71fe0f8a77a494ab92130cfc649c"),
							Index:           275,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xaa33107f3cf26f828a955970f7765724aadf71fe0f8a77a494ab92130cfc649c"),
						TransactionIndex: 105,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0xaa33107f3cf26f828a955970f7765724aadf71fe0f8a77a494ab92130cfc649c",
				Network: filter.NetworkEthereum,
				Index:   105,
				From:    "0x6d5601E90220C989111939d9317FCbba27c015ab",
				To:      "0xD4416b13d2b3a9aBae7AcD5D6C2BbDBE25686401",
				Type:    filter.TypeSocialProfile,
				Calldata: &schema.Calldata{
					FunctionHash: "0x402906fc",
				},
				Platform: lo.ToPtr(filter.PlatformENS),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("1266415647151962")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformENS.String(),
						From:     "0x6d5601E90220C989111939d9317FCbba27c015ab",
						To:       "0xD4416b13d2b3a9aBae7AcD5D6C2BbDBE25686401",
						Metadata: metadata.SocialProfile{
							Action:    metadata.ActionSocialProfileUpdate,
							ProfileID: "33672763687308778590322826700630451429751767233736265975971535021200371750237",
							Handle:    "aintgottadollar.eth",
							Address:   ens.AddressBaseRegistrarImplementation,
							ImageURI:  "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/33672763687308778590322826700630451429751767233736265975971535021200371750237/image",
							Key:       "Fuses",
							Value:     "196609",
						},
					},
				},
				Status:    true,
				Timestamp: 1692029303,
			},
			wantError: require.NoError,
		},
		{
			name: "ENS ContentHash Changed",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xb2a142d10c1ed9a886c4267eec7d80e6b21558ef276c2750da003253b98362c3"),
						ParentHash:   common.HexToHash("0x38652eb11ab6bcd456b6b3e23fc3379410e7a6cb787f86699154a2abc4f7903b"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("17932571", 0)),
						GasLimit:     30000000,
						GasUsed:      19818297,
						Timestamp:    1692251543,
						BaseFee:      lo.Must(new(big.Int).SetString("16076742490", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x01ce07608c77865004f8a6ae8139b14b2e10e304ab3f214cf12bf79a6410e6b9"),
						From:      common.HexToAddress("0xDFF917ab602e8508b6907dE1b038dd52B24A2379"),
						Gas:       47097,
						GasPrice:  lo.Must(new(big.Int).SetString("16176742490", 10)),
						Hash:      common.HexToHash("0x01ce07608c77865004f8a6ae8139b14b2e10e304ab3f214cf12bf79a6410e6b9"),
						Input:     hexutil.MustDecode("0x304e6ade724a552bea37ebe5b9c5bc0cc9eef73b60fd1b227756e6cf1612ef91652fa7e400000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000026e301017012205e573e98f03fa345cdd904931e95cd81435d4928c45a0fd7f7f715d8ad30132e0000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xb2a142d10c1ed9a886c4267eec7d80e6b21558ef276c2750da003253b98362c3"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17932571", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 16666340,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3c435805a"),
						GasUsed:           47097,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41"),
							Topics: []common.Hash{
								common.HexToHash("0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578"),
								common.HexToHash("0x724a552bea37ebe5b9c5bc0cc9eef73b60fd1b227756e6cf1612ef91652fa7e4"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000026e301017012205e573e98f03fa345cdd904931e95cd81435d4928c45a0fd7f7f715d8ad30132e0000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17932571", 0)),
							TransactionHash: common.HexToHash("0x01ce07608c77865004f8a6ae8139b14b2e10e304ab3f214cf12bf79a6410e6b9"),
							Index:           385,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x01ce07608c77865004f8a6ae8139b14b2e10e304ab3f214cf12bf79a6410e6b9"),
						TransactionIndex: 126,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x01ce07608c77865004f8a6ae8139b14b2e10e304ab3f214cf12bf79a6410e6b9",
				Network: filter.NetworkEthereum,
				Index:   126,
				From:    "0xDFF917ab602e8508b6907dE1b038dd52B24A2379",
				To:      "0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41",
				Type:    filter.TypeSocialProfile,
				Calldata: &schema.Calldata{
					FunctionHash: "0x304e6ade",
				},
				Platform: lo.ToPtr(filter.PlatformENS),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("761876041051530")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformENS.String(),
						From:     "0xDFF917ab602e8508b6907dE1b038dd52B24A2379",
						To:       "0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41",
						Metadata: metadata.SocialProfile{
							Action:    metadata.ActionSocialProfileUpdate,
							ProfileID: "114786103065688359762889232472579792936878517910599069657674525697915765984556",
							Handle:    "legal🦅.eth",
							Address:   ens.AddressBaseRegistrarImplementation,
							ImageURI:  "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/114786103065688359762889232472579792936878517910599069657674525697915765984556/image",
							Key:       "ContentHash",
							Value:     "0x5e573e98f03fa345cdd904931e95cd81435d4928c45a0fd7f7f715d8ad30132e",
						},
					},
				},
				Status:    true,
				Timestamp: 1692251543,
			},
			wantError: require.NoError,
		},
		{
			name: "ENS Name Changed",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xba0387a7bbba87f7b7fe9f9a8c1fa416fad711a96d2621bd35f1b602f3dcd2ee"),
						ParentHash:   common.HexToHash("0xb90fb4a737ebd34fe0fafd8c39557b60abe06bcefaa8ac04eae9f4f81fe199e8"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xab3B229eB4BcFF881275E7EA2F0FD24eeaC8C83a"),
						Number:       lo.Must(new(big.Int).SetString("14849713", 0)),
						GasLimit:     29970705,
						GasUsed:      8219168,
						Timestamp:    1653592227,
						BaseFee:      lo.Must(new(big.Int).SetString("74071025981", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x6b2b172406ace420b9c43f91bb9ae0b2948ee85eea91ba1a61e29c5003e57379"),
						From:      common.HexToAddress("0x63A2368F4B509438ca90186cb1C15156713D5834"),
						Gas:       84072,
						GasPrice:  lo.Must(new(big.Int).SetString("75071025981", 10)),
						Hash:      common.HexToHash("0x6b2b172406ace420b9c43f91bb9ae0b2948ee85eea91ba1a61e29c5003e57379"),
						Input:     hexutil.MustDecode("0x77372213d090701b54d48eadf461c98976d657a3d75e6d108dc67434922aa4e230f518ad000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000046173646600000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xba0387a7bbba87f7b7fe9f9a8c1fa416fad711a96d2621bd35f1b602f3dcd2ee"),
						BlockNumber:       lo.Must(new(big.Int).SetString("14849713", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 6472962,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x117a94f33d"),
						GasUsed:           56048,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41"),
							Topics: []common.Hash{
								common.HexToHash("0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7"),
								common.HexToHash("0xd090701b54d48eadf461c98976d657a3d75e6d108dc67434922aa4e230f518ad"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000046173646600000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("14849713", 0)),
							TransactionHash: common.HexToHash("0x6b2b172406ace420b9c43f91bb9ae0b2948ee85eea91ba1a61e29c5003e57379"),
							Index:           87,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x6b2b172406ace420b9c43f91bb9ae0b2948ee85eea91ba1a61e29c5003e57379"),
						TransactionIndex: 58,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x6b2b172406ace420b9c43f91bb9ae0b2948ee85eea91ba1a61e29c5003e57379",
				Network: filter.NetworkEthereum,
				Index:   58,
				From:    "0x63A2368F4B509438ca90186cb1C15156713D5834",
				To:      "0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41",
				Type:    filter.TypeSocialProfile,
				Calldata: &schema.Calldata{
					FunctionHash: "0x77372213",
				},
				Platform: lo.ToPtr(filter.PlatformENS),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("4207580864183088")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformENS.String(),
						From:     "0x63A2368F4B509438ca90186cb1C15156713D5834",
						To:       "0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41",
						Metadata: metadata.SocialProfile{
							Action:    metadata.ActionSocialProfileUpdate,
							ProfileID: "60387918586951520976322447866692493653835638235765373089736745843064604088600",
							Handle:    "testinginprod.eth",
							Address:   ens.AddressBaseRegistrarImplementation,
							ImageURI:  "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/60387918586951520976322447866692493653835638235765373089736745843064604088600/image",
							Key:       "Name",
							Value:     "asdf",
						},
					},
				},
				Status:    true,
				Timestamp: 1653592227,
			},
			wantError: require.NoError,
		},
		{
			name: "ENS Address Changed",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xc174922af411ebe8b9aa6ceff6110d225c06801ed9f921ef697c32fe62040d3e"),
						ParentHash:   common.HexToHash("0x9924c0676058914a7c5a7e1dfedc806faa1cf8128a6d217b85bcb52482b4b528"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("17932149", 0)),
						GasLimit:     30000000,
						GasUsed:      19877173,
						Timestamp:    1692246443,
						BaseFee:      lo.Must(new(big.Int).SetString("16297491278", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x41ac65d8b98e28587ce4836a192f2416a2f11bb96895fcea29b86a479e8cd360"),
						From:      common.HexToAddress("0x07bd403d0E4Cd0f2cF5e4b1eF44D8Fb18CF6eCad"),
						Gas:       47789,
						GasPrice:  lo.Must(new(big.Int).SetString("16397491278", 10)),
						Hash:      common.HexToHash("0x41ac65d8b98e28587ce4836a192f2416a2f11bb96895fcea29b86a479e8cd360"),
						Input:     hexutil.MustDecode("0xac9650d800000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000a48b95dd711b24fd1914601309cd6e8fa33e342b11abb2fa2ccef6ccc9c77597f5e1b8b242000000000000000000000000000000000000000000000000000000000000003c0000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000001407bd403d0e4cd0f2cf5e4b1ef44d8fb18cf6ecad00000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xc174922af411ebe8b9aa6ceff6110d225c06801ed9f921ef697c32fe62040d3e"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17932149", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 13576294,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3d15ddc4e"),
						GasUsed:           47789,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41"),
							Topics: []common.Hash{
								common.HexToHash("0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752"),
								common.HexToHash("0x1b24fd1914601309cd6e8fa33e342b11abb2fa2ccef6ccc9c77597f5e1b8b242"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000003c0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000001407bd403d0e4cd0f2cf5e4b1ef44d8fb18cf6ecad000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17932149", 0)),
							TransactionHash: common.HexToHash("0x41ac65d8b98e28587ce4836a192f2416a2f11bb96895fcea29b86a479e8cd360"),
							Index:           364,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41"),
							Topics: []common.Hash{
								common.HexToHash("0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2"),
								common.HexToHash("0x1b24fd1914601309cd6e8fa33e342b11abb2fa2ccef6ccc9c77597f5e1b8b242"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000007bd403d0e4cd0f2cf5e4b1ef44d8fb18cf6ecad"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17932149", 0)),
							TransactionHash: common.HexToHash("0x41ac65d8b98e28587ce4836a192f2416a2f11bb96895fcea29b86a479e8cd360"),
							Index:           365,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x41ac65d8b98e28587ce4836a192f2416a2f11bb96895fcea29b86a479e8cd360"),
						TransactionIndex: 138,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x41ac65d8b98e28587ce4836a192f2416a2f11bb96895fcea29b86a479e8cd360",
				Network: filter.NetworkEthereum,
				Index:   138,
				From:    "0x07bd403d0E4Cd0f2cF5e4b1eF44D8Fb18CF6eCad",
				To:      "0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41",
				Type:    filter.TypeSocialProfile,
				Calldata: &schema.Calldata{
					FunctionHash: "0xac9650d8",
				},
				Platform: lo.ToPtr(filter.PlatformENS),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("783619710684342")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformENS.String(),
						From:     "0x07bd403d0E4Cd0f2cF5e4b1eF44D8Fb18CF6eCad",
						To:       "0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41",
						Metadata: metadata.SocialProfile{
							Action:    metadata.ActionSocialProfileUpdate,
							ProfileID: "102116724323029317809332311195807824478919292408436900635638808938719422634466",
							Handle:    "rifler.eth",
							Address:   ens.AddressBaseRegistrarImplementation,
							ImageURI:  "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/102116724323029317809332311195807824478919292408436900635638808938719422634466/image",
							Key:       "Address",
							Value:     "0x07bd403d0E4Cd0f2cF5e4b1eF44D8Fb18CF6eCad",
						},
					},
				},
				Status:    true,
				Timestamp: 1692246443,
			},
			wantError: require.NoError,
		},
		{
			name: "ENS Public Key Changed",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x16971d16aa0abb60a0b2fdee0719c53c7299bc9d4b74177102b7a5fce75162bb"),
						ParentHash:   common.HexToHash("0xe36fb88b8ced846710fbec2cbe05fb3c358115fb141a60818f2f7d3efa185ce7"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xDAFEA492D9c6733ae3d56b7Ed1ADB60692c98Bc5"),
						Number:       lo.Must(new(big.Int).SetString("17558923", 0)),
						GasLimit:     30000000,
						GasUsed:      10986541,
						Timestamp:    1687725635,
						BaseFee:      lo.Must(new(big.Int).SetString("11803834461", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x6794d34bd533740716b19658e4c957aa1e39cbfb1d34dc56aed50b1cca31fbdf"),
						From:      common.HexToAddress("0x790BEd7B93e14235d8EB153Eb7CF4497906260F4"),
						Gas:       93982,
						GasPrice:  lo.Must(new(big.Int).SetString("11880831964", 10)),
						Hash:      common.HexToHash("0x6794d34bd533740716b19658e4c957aa1e39cbfb1d34dc56aed50b1cca31fbdf"),
						Input:     hexutil.MustDecode("0x29cd62ea9e0a38b09e9ca677ff14b4dc22f8cffd1d89cacff6797c553210b8966906a627edf13e710c0ab6b6e5fb662a2cb4ec8aa92a72b6facdb531ab75576e4de9482d7ec161c9fb865fe6f29974973de8a964f3e37cfb172e2f671ff886b1820ded39"),
						To:        lo.ToPtr(common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x16971d16aa0abb60a0b2fdee0719c53c7299bc9d4b74177102b7a5fce75162bb"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17558923", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 10129327,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x2c4271bdc"),
						GasUsed:           78319,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41"),
							Topics: []common.Hash{
								common.HexToHash("0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46"),
								common.HexToHash("0x9e0a38b09e9ca677ff14b4dc22f8cffd1d89cacff6797c553210b8966906a627"),
							},
							Data:            hexutil.MustDecode("0xedf13e710c0ab6b6e5fb662a2cb4ec8aa92a72b6facdb531ab75576e4de9482d7ec161c9fb865fe6f29974973de8a964f3e37cfb172e2f671ff886b1820ded39"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17558923", 0)),
							TransactionHash: common.HexToHash("0x6794d34bd533740716b19658e4c957aa1e39cbfb1d34dc56aed50b1cca31fbdf"),
							Index:           222,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x6794d34bd533740716b19658e4c957aa1e39cbfb1d34dc56aed50b1cca31fbdf"),
						TransactionIndex: 91,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x6794d34bd533740716b19658e4c957aa1e39cbfb1d34dc56aed50b1cca31fbdf",
				Network: filter.NetworkEthereum,
				Index:   91,
				From:    "0x790BEd7B93e14235d8EB153Eb7CF4497906260F4",
				To:      "0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41",
				Type:    filter.TypeSocialProfile,
				Calldata: &schema.Calldata{
					FunctionHash: "0x29cd62ea",
				},
				Platform: lo.ToPtr(filter.PlatformENS),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("930494878588516")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformENS.String(),
						From:     "0x790BEd7B93e14235d8EB153Eb7CF4497906260F4",
						To:       "0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41",
						Metadata: metadata.SocialProfile{
							Action:    metadata.ActionSocialProfileUpdate,
							ProfileID: "61943364761146994600684238183074484629471563196685701445153204449159877985562",
							Handle:    "alexkubica.eth",
							Address:   ens.AddressBaseRegistrarImplementation,
							ImageURI:  "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/61943364761146994600684238183074484629471563196685701445153204449159877985562/image",
							Key:       "Pubkey",
							Value:     "03edf13e710c0ab6b6e5fb662a2cb4ec8aa92a72b6facdb531ab75576e4de9482d",
						},
					},
				},
				Status:    true,
				Timestamp: 1687725635,
			},
			wantError: require.NoError,
		},
	}

	driver := database.DriverCockroachDB
	partition := true

	testDB, err := testserver.NewTestServer(testserver.CustomVersionOpt("v23.1.8"))
	//container, dataSourceName, err := createContainer(context.Background(), driver, partition)
	require.NoError(t, err)

	//t.Cleanup(func() {
	//	require.NoError(t, gnomock.Stop(container))
	//})

	// Dial the database.
	databaseClient, err := dialer.Dial(context.Background(), &config.Database{
		Driver:    driver,
		URI:       testDB.PGURL().String(),
		Partition: &partition,
	})
	require.NoError(t, err)

	err = databaseClient.Migrate(context.Background())
	require.NoError(t, err)

	// Pre-initialize some name hash
	requiredNames := []string{
		// Skip register events user - they should be saved automatically into dataset
		"33222222.eth",         // ENS NameRenewed V1
		"scraltach.eth",        // ENS NameRenewed V2
		"cryptopaycheck.eth",   // ENS Public Resolver TextChanged
		"pepetrillionaire.eth", // ENS Public Resolver TextChanged With Value
		"0-770.eth",            // ENS Name Wrapped
		"notantoine.eth",       // ENS Name Unwrapped
		"aintgottadollar.eth",  // ENS FusesSet
		"legal🦅.eth",           // ENS ContentHash Changed
		"testinginprod.eth",    // ENS Name Changed
		"rifler.eth",           // ENS Address Changed
		"alexkubica.eth",       // ENS Public Key Changed
	}
	for _, fullName := range requiredNames {
		err = databaseClient.SaveDatasetENSNamehash(context.Background(), &model.ENSNamehash{
			Name: fullName,
			Hash: ens.NameHash(fullName),
		})
		require.NoError(t, err)
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			instance, err := worker.NewWorker(testcase.arguments.config, databaseClient)
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

//func createContainer(ctx context.Context, driver database.Driver, partition bool) (container *gnomock.Container, dataSourceName string, err error) {
//	cfg := config.Database{
//		Driver:    driver,
//		Partition: &partition,
//	}
//
//	switch driver {
//	case database.DriverCockroachDB:
//		preset := cockroachdb.Preset(
//			cockroachdb.WithDatabase("test"),
//			cockroachdb.WithVersion("v23.1.8"),
//		)
//
//		// Use a health check function to wait for the database to be ready.
//		healthcheckFunc := func(ctx context.Context, container *gnomock.Container) error {
//			cfg.URI = formatContainerURI(container)
//
//			client, err := dialer.Dial(ctx, &cfg)
//			if err != nil {
//				return err
//			}
//
//			transaction, err := client.Begin(ctx)
//			if err != nil {
//				return err
//			}
//
//			defer lo.Try(transaction.Rollback)
//
//			return nil
//		}
//
//		container, err = gnomock.Start(preset, gnomock.WithContext(ctx), gnomock.WithHealthCheck(healthcheckFunc))
//		if err != nil {
//			return nil, "", err
//		}
//
//		return container, formatContainerURI(container), nil
//	default:
//		return nil, "", fmt.Errorf("unsupported driver: %s", driver)
//	}
//}
//
//func formatContainerURI(container *gnomock.Container) string {
//	return fmt.Sprintf(
//		"postgres://root@%s:%d/%s?sslmode=disable",
//		container.Host,
//		container.DefaultPort(),
//		"test",
//	)
//}
