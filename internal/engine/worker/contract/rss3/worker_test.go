package rss3_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/contract/rss3"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/rss3"
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
			name: "Deposit $RSS3 to the Staking contract",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x9fa20b204527b2bb325ab65062c8f035a32ee0a6b42351e349d77cbdec11a225"),
						ParentHash:   common.HexToHash("0xe0e05ac73328a38f02e3b3b1c42f512ba111fef3f406252f7a54d1c9083b65d3"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("16577896", 0)),
						GasLimit:     30000000,
						GasUsed:      14433408,
						Timestamp:    1675784915,
						BaseFee:      lo.Must(new(big.Int).SetString("33788878088", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x85ded9199c3c8a46bcb1e431e3e05f62fe014930c0a13a029a78e3a18061011f"),
						From:      common.HexToAddress("0x1B861760AdE296aBE523C594118EF812208194CE"),
						Gas:       300000,
						Hash:      common.HexToHash("0x85ded9199c3c8a46bcb1e431e3e05f62fe014930c0a13a029a78e3a18061011f"),
						Input:     hexutil.MustDecode("0x383c7d8700000000000000000000000000000000000000000000043c33c1937564800000"),
						To:        lo.ToPtr(common.HexToAddress("0x5301CbBeDc048AbaC7e213184132cf982d593563")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x9fa20b204527b2bb325ab65062c8f035a32ee0a6b42351e349d77cbdec11a225"),
						BlockNumber:       lo.Must(new(big.Int).SetString("16577896", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 11526815,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x819942708"),
						GasUsed:           157507,
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0x5301CbBeDc048AbaC7e213184132cf982d593563"),
								Topics: []common.Hash{
									common.HexToHash("0xdf29796aad820e4bb192f3a8d631b76519bcd2cbe77cc85af20e9df53cece086"),
									common.HexToHash("0x0000000000000000000000001b861760ade296abe523c594118ef812208194ce"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000002c8e66a3de3fbf1f200"),
								BlockNumber:     lo.Must(new(big.Int).SetString("16577896", 0)),
								TransactionHash: common.HexToHash("0x85ded9199c3c8a46bcb1e431e3e05f62fe014930c0a13a029a78e3a18061011f"),
								Index:           262,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0xc98D64DA73a6616c42117b582e832812e7B8D57F"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x0000000000000000000000001b861760ade296abe523c594118ef812208194ce"),
									common.HexToHash("0x0000000000000000000000005301cbbedc048abac7e213184132cf982d593563"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000043c33c1937564800000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("16577896", 0)),
								TransactionHash: common.HexToHash("0x85ded9199c3c8a46bcb1e431e3e05f62fe014930c0a13a029a78e3a18061011f"),
								Index:           263,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0xc98D64DA73a6616c42117b582e832812e7B8D57F"),
								Topics: []common.Hash{
									common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
									common.HexToHash("0x0000000000000000000000001b861760ade296abe523c594118ef812208194ce"),
									common.HexToHash("0x0000000000000000000000005301cbbedc048abac7e213184132cf982d593563"),
								},
								Data:            hexutil.MustDecode("0xfffffffffffffffffffffffffffffffffffffffffffff9a5b25da2cfe93fffff"),
								BlockNumber:     lo.Must(new(big.Int).SetString("16577896", 0)),
								TransactionHash: common.HexToHash("0x85ded9199c3c8a46bcb1e431e3e05f62fe014930c0a13a029a78e3a18061011f"),
								Index:           264,
								Removed:         false,
							}, {
								Address: common.HexToAddress("0x5301CbBeDc048AbaC7e213184132cf982d593563"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
									common.HexToHash("0x0000000000000000000000001b861760ade296abe523c594118ef812208194ce"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000043c33c1937564800000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("16577896", 0)),
								TransactionHash: common.HexToHash("0x85ded9199c3c8a46bcb1e431e3e05f62fe014930c0a13a029a78e3a18061011f"),
								Index:           265,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x5301CbBeDc048AbaC7e213184132cf982d593563"),
								Topics: []common.Hash{
									common.HexToHash("0x91ede45f04a37a7c170f5c1207df3b6bc748dc1e04ad5e917a241d0f52feada3"),
									common.HexToHash("0x0000000000000000000000001b861760ade296abe523c594118ef812208194ce"),
									common.HexToHash("0x00000000000000000000000000000000000000000000043c33c1937564800000"),
									common.HexToHash("0x000000000000000000000000000000000000000000000000000000000076a700"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000063e272d3"),
								BlockNumber:     lo.Must(new(big.Int).SetString("16577896", 0)),
								TransactionHash: common.HexToHash("0x85ded9199c3c8a46bcb1e431e3e05f62fe014930c0a13a029a78e3a18061011f"),
								Index:           266,
								Removed:         false,
							},
						},
						Status:           1,
						TransactionHash:  common.HexToHash("0x85ded9199c3c8a46bcb1e431e3e05f62fe014930c0a13a029a78e3a18061011f"),
						TransactionIndex: 140,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x85ded9199c3c8a46bcb1e431e3e05f62fe014930c0a13a029a78e3a18061011f",
				Network: filter.NetworkEthereum,
				Index:   140,
				From:    "0x1B861760AdE296aBE523C594118EF812208194CE",
				To:      rss3.AddressStaking.String(),
				Type:    filter.TypeExchangeStaking,
				Calldata: &schema.Calldata{
					FunctionHash: "0x383c7d87",
				},
				Platform: lo.ToPtr(filter.PlatformRSS3),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("5479491821006616")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeStaking,
						Platform: filter.PlatformRSS3.String(),
						From:     "0x1B861760AdE296aBE523C594118EF812208194CE",
						To:       "0x1B861760AdE296aBE523C594118EF812208194CE",
						Metadata: metadata.ExchangeStaking{
							Action: metadata.ActionExchangeStakingStake,
							Token: metadata.Token{
								Address:  lo.ToPtr("0xc98D64DA73a6616c42117b582e832812e7B8D57F"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("20000000000000000000000"))),
								Name:     "RSS3",
								Symbol:   "RSS3",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
							Period: &metadata.ExchangeStakingPeriod{
								Start: time.Unix(1675784915, 0),
								End:   time.Unix(1675784915+7776000, 0),
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1675784915,
			},
			wantError: require.NoError,
		},
		{
			name: "Withdraw $RSS3 from the Staking contract",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x1064d670c369cf48e032bd03a9c0d9b7db9eb6929bab811c7f6d79f83a14ab63"),
						ParentHash:   common.HexToHash("0x14c46f9ba345a1945cbbd5b313991c2a1de1042185544cd27e1dd87e5a2f04a4"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x388C818CA8B9251b393131C08a736A67ccB19297"),
						Number:       lo.Must(new(big.Int).SetString("17032966", 0)),
						GasLimit:     30000000,
						GasUsed:      29989768,
						Timestamp:    1681315199,
						BaseFee:      lo.Must(new(big.Int).SetString("24548648256", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xb3a2c0760f29dba991b8d225a705dbb1785fba7b0e04f4a54c3e91560a1b3227"),
						From:      common.HexToAddress("0xE0BA908Be2f52063B0bD210544e67FCD76bd0b56"),
						Gas:       371432,
						Hash:      common.HexToHash("0xb3a2c0760f29dba991b8d225a705dbb1785fba7b0e04f4a54c3e91560a1b3227"),
						Input:     hexutil.MustDecode("0x00f714ce0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e0ba908be2f52063b0bd210544e67fcd76bd0b56"),
						To:        lo.ToPtr(common.HexToAddress("0x5301CbBeDc048AbaC7e213184132cf982d593563")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x1064d670c369cf48e032bd03a9c0d9b7db9eb6929bab811c7f6d79f83a14ab63"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17032966", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 24000114,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x5d21dba00"),
						GasUsed:           92858,
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0x5301CbBeDc048AbaC7e213184132cf982d593563"),
								Topics: []common.Hash{
									common.HexToHash("0xdf29796aad820e4bb192f3a8d631b76519bcd2cbe77cc85af20e9df53cece086"),
									common.HexToHash("0x000000000000000000000000e0ba908be2f52063b0bd210544e67fcd76bd0b56"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000ef4898fe4143613560"),
								BlockNumber:     lo.Must(new(big.Int).SetString("17032966", 0)),
								TransactionHash: common.HexToHash("0xb3a2c0760f29dba991b8d225a705dbb1785fba7b0e04f4a54c3e91560a1b3227"),
								Index:           623,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x5301CbBeDc048AbaC7e213184132cf982d593563"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x000000000000000000000000e0ba908be2f52063b0bd210544e67fcd76bd0b56"),
									common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000001cd2a9e51e0578e0000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("17032966", 0)),
								TransactionHash: common.HexToHash("0xb3a2c0760f29dba991b8d225a705dbb1785fba7b0e04f4a54c3e91560a1b3227"),
								Index:           624,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0xc98D64DA73a6616c42117b582e832812e7B8D57F"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x0000000000000000000000005301cbbedc048abac7e213184132cf982d593563"),
									common.HexToHash("0x000000000000000000000000e0ba908be2f52063b0bd210544e67fcd76bd0b56"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000e6954f28f02bc70000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("17032966", 0)),
								TransactionHash: common.HexToHash("0xb3a2c0760f29dba991b8d225a705dbb1785fba7b0e04f4a54c3e91560a1b3227"),
								Index:           625,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x5301CbBeDc048AbaC7e213184132cf982d593563"),
								Topics: []common.Hash{
									common.HexToHash("0xe5df19de43c8c04fd192bc68e484b2593570925fbb6ad8c07ccafbc2aa5c37a1"),
									common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
									common.HexToHash("0x000000000000000000000000e0ba908be2f52063b0bd210544e67fcd76bd0b56"),
									common.HexToHash("0x000000000000000000000000e0ba908be2f52063b0bd210544e67fcd76bd0b56"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000e6954f28f02bc70000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("17032966", 0)),
								TransactionHash: common.HexToHash("0xb3a2c0760f29dba991b8d225a705dbb1785fba7b0e04f4a54c3e91560a1b3227"),
								Index:           626,
								Removed:         false,
							},
						},
						Status:           1,
						TransactionHash:  common.HexToHash("0xb3a2c0760f29dba991b8d225a705dbb1785fba7b0e04f4a54c3e91560a1b3227"),
						TransactionIndex: 305,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0xb3a2c0760f29dba991b8d225a705dbb1785fba7b0e04f4a54c3e91560a1b3227",
				Network: filter.NetworkEthereum,
				Index:   305,
				From:    "0xE0BA908Be2f52063B0bD210544e67FCD76bd0b56",
				To:      rss3.AddressStaking.String(),
				Type:    filter.TypeExchangeStaking,
				Calldata: &schema.Calldata{
					FunctionHash: "0x00f714ce",
				},
				Platform: lo.ToPtr(filter.PlatformRSS3),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("2321450000000000")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeStaking,
						Platform: filter.PlatformRSS3.String(),
						From:     "0xE0BA908Be2f52063B0bD210544e67FCD76bd0b56",
						To:       "0xE0BA908Be2f52063B0bD210544e67FCD76bd0b56",
						Metadata: metadata.ExchangeStaking{
							Action: metadata.ActionExchangeStakingUnstake,
							Token: metadata.Token{
								Address:  lo.ToPtr("0xc98D64DA73a6616c42117b582e832812e7B8D57F"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("4253510000000000000000"))),
								Name:     "RSS3",
								Symbol:   "RSS3",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1681315199,
			},
			wantError: require.NoError,
		},
		{
			name: "Claim $RSS3 from the Staking contract",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xe685eebd34c2e1a1f7e941a2a5c09281849234a01cd868f99b8596c31ba2e5ef"),
						ParentHash:   common.HexToHash("0x8e774d3084dea206d1e0a396cadfc9e1f170324adf8a04008a1221954e6113f5"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x388C818CA8B9251b393131C08a736A67ccB19297"),
						Number:       lo.Must(new(big.Int).SetString("16617713", 0)),
						GasLimit:     30000000,
						GasUsed:      24570190,
						Timestamp:    1676265671,
						BaseFee:      lo.Must(new(big.Int).SetString("14357381025", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x46cf455d0500f0b0979401ef1192af508890c2f888bc0dc6c006b49f65a37b6b"),
						From:      common.HexToAddress("0x2a03278590cd1962De28F9AbC855CF3774fe3eb6"),
						Gas:       334112,
						Hash:      common.HexToHash("0x46cf455d0500f0b0979401ef1192af508890c2f888bc0dc6c006b49f65a37b6b"),
						Input:     hexutil.MustDecode("0xef5cfb8c0000000000000000000000002a03278590cd1962de28f9abc855cf3774fe3eb6"),
						To:        lo.ToPtr(common.HexToAddress("0x5301CbBeDc048AbaC7e213184132cf982d593563")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xe685eebd34c2e1a1f7e941a2a5c09281849234a01cd868f99b8596c31ba2e5ef"),
						BlockNumber:       lo.Must(new(big.Int).SetString("16617713", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 2617691,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3b7224fa1"),
						GasUsed:           83528,
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0x5301CbBeDc048AbaC7e213184132cf982d593563"),
								Topics: []common.Hash{
									common.HexToHash("0xdf29796aad820e4bb192f3a8d631b76519bcd2cbe77cc85af20e9df53cece086"),
									common.HexToHash("0x0000000000000000000000002a03278590cd1962de28f9abc855cf3774fe3eb6"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000d33aca254d0847b0"),
								BlockNumber:     lo.Must(new(big.Int).SetString("16617713", 0)),
								TransactionHash: common.HexToHash("0x46cf455d0500f0b0979401ef1192af508890c2f888bc0dc6c006b49f65a37b6b"),
								Index:           105,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x5301CbBeDc048AbaC7e213184132cf982d593563"),
								Topics: []common.Hash{
									common.HexToHash("0x8a43c4352486ec339f487f64af78ca5cbf06cd47833f073d3baf3a193e503161"),
									common.HexToHash("0x0000000000000000000000002a03278590cd1962de28f9abc855cf3774fe3eb6"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000001acccc7dcc31ff79e4"),
								BlockNumber:     lo.Must(new(big.Int).SetString("16617713", 0)),
								TransactionHash: common.HexToHash("0x46cf455d0500f0b0979401ef1192af508890c2f888bc0dc6c006b49f65a37b6b"),
								Index:           106,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0xc98D64DA73a6616c42117b582e832812e7B8D57F"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x0000000000000000000000005301cbbedc048abac7e213184132cf982d593563"),
									common.HexToHash("0x0000000000000000000000002a03278590cd1962de28f9abc855cf3774fe3eb6"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000001acccc7dcc31ff79e4"),
								BlockNumber:     lo.Must(new(big.Int).SetString("16617713", 0)),
								TransactionHash: common.HexToHash("0x46cf455d0500f0b0979401ef1192af508890c2f888bc0dc6c006b49f65a37b6b"),
								Index:           107,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x5301CbBeDc048AbaC7e213184132cf982d593563"),
								Topics: []common.Hash{
									common.HexToHash("0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7"),
									common.HexToHash("0x0000000000000000000000002a03278590cd1962de28f9abc855cf3774fe3eb6"),
									common.HexToHash("0x0000000000000000000000002a03278590cd1962de28f9abc855cf3774fe3eb6"),
									common.HexToHash("0x00000000000000000000000000000000000000000000001acccc7dcc31ff79e4"),
								},
								Data:            nil,
								BlockNumber:     lo.Must(new(big.Int).SetString("16617713", 0)),
								TransactionHash: common.HexToHash("0x46cf455d0500f0b0979401ef1192af508890c2f888bc0dc6c006b49f65a37b6b"),
								Index:           108,
								Removed:         false,
							},
						},
						Status:           1,
						TransactionHash:  common.HexToHash("0x46cf455d0500f0b0979401ef1192af508890c2f888bc0dc6c006b49f65a37b6b"),
						TransactionIndex: 49,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x46cf455d0500f0b0979401ef1192af508890c2f888bc0dc6c006b49f65a37b6b",
				Network: filter.NetworkEthereum,
				Index:   49,
				From:    "0x2a03278590cd1962De28F9AbC855CF3774fe3eb6",
				To:      rss3.AddressStaking.String(),
				Type:    filter.TypeExchangeStaking,
				Calldata: &schema.Calldata{
					FunctionHash: "0xef5cfb8c",
				},
				Platform: lo.ToPtr(filter.PlatformRSS3),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("1332888122256200")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeStaking,
						Platform: filter.PlatformRSS3.String(),
						From:     "0x2a03278590cd1962De28F9AbC855CF3774fe3eb6",
						To:       "0x2a03278590cd1962De28F9AbC855CF3774fe3eb6",
						Metadata: metadata.ExchangeStaking{
							Action: metadata.ActionExchangeStakingClaim,
							Token: metadata.Token{
								Address:  lo.ToPtr("0xc98D64DA73a6616c42117b582e832812e7B8D57F"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("494372654311400241636"))),
								Name:     "RSS3",
								Symbol:   "RSS3",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1676265671,
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
