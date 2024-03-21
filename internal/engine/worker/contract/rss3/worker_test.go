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
		// Ethereum Mainnet
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
		// VSL Mainnet
		{
			name: "Deposit $RSS3 to the VSL Staking contract",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkVSL,
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
					Network:  filter.NetworkVSL,
					Endpoint: endpoint.MustGet(filter.NetworkVSL),
				},
			},
			want: &schema.Feed{
				ID:      "0xb49208a4fb936cf3d73da0b92af0776945a09062af8983cf72c8bd4fffb7f624",
				Network: filter.NetworkVSL,
				Index:   1,
				From:    "0x39F9e912C1F696F533e7A2267Ea233AeC9742b35",
				To:      rss3.AddressStakingVSL.String(),
				Type:    filter.TypeExchangeStaking,
				Calldata: &schema.Calldata{
					FunctionHash: "0x96531623",
				},
				Platform: lo.ToPtr(filter.PlatformRSS3),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("1117542222924365900")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeStaking,
						Platform: filter.PlatformRSS3.String(),
						From:     "0x39F9e912C1F696F533e7A2267Ea233AeC9742b35",
						To:       "0x39F9e912C1F696F533e7A2267Ea233AeC9742b35",
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
				To:      rss3.AddressStakingVSL.String(),
				Type:    filter.TypeExchangeStaking,
				Calldata: &schema.Calldata{
					FunctionHash: "0x26476204",
				},
				Platform: lo.ToPtr(filter.PlatformRSS3),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("772196357569891650")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeCollectibleMint,
						Platform: filter.PlatformRSS3.String(),
						From:     ethereum.AddressGenesis.String(),
						To:       common.HexToAddress("0x30286DD245338292F319809935a1037CcD4573Ea").String(),
						Metadata: metadata.CollectibleTransfer{
							Address:  lo.ToPtr("0x849f8F55078dCc69dD857b58Cc04631EBA54E4DE"),
							ID:       lo.ToPtr(decimal.NewFromBigInt(big.NewInt(162), 0)),
							Value:    lo.ToPtr(decimal.NewFromBigInt(big.NewInt(1), 0)),
							Name:     "Open Chips",
							Symbol:   "Chips",
							Standard: metadata.StandardERC721,
							URI:      "data:application/json;base64,eyJuYW1lIjogIk9wZW4gQ2hpcHMgIzE2MiIsICJkZXNjcmlwdGlvbiI6ICJDaGlwIE1vbnN0ZXJzIGFyZSB1bmlxdWUgY3JlYXR1cmVzIGxpdmluZyBpbiB0aGUgUlNTMyBOZXR3b3JrLCBlYWNoIG9uZSBzcGVjaWFsIGJlY2F1c2Ugb2Ygd2hlcmUgaXQgd2FzIGJvcm4uIFRoZXkgcmVwcmVzZW50IHRoZSBpZGVhIG9mIEZSRUUgYW5kIE9QRU4gSU5GT1JNQVRJT04sIHRocml2aW5nIGluIGEgd29ybGQgdGhhdCB2YWx1ZXMgc2hhcmluZyBhbmQgYmVpbmcgZGlmZmVyZW50LiBUaGVzZSBDaGlwIE1vbnN0ZXJzIGFyZSBtb3JlIHRoYW4ganVzdCBkaWdpdGFsOyB0aGV5IHN5bWJvbGl6ZSB0aGUgZXhjaXRlbWVudCBhbmQgaW1wb3J0YW5jZSBvZiBiZWluZyB1bmlxdWUgaW4gYSBjb25uZWN0ZWQgZGlnaXRhbCB3b3JsZC4iLCJpbWFnZSI6ImRhdGE6aW1hZ2Uvc3ZnK3htbDtiYXNlNjQsUEQ5NGJXd2dkbVZ5YzJsdmJqMGlNUzR3SWlCbGJtTnZaR2x1WnowaWRYUm1MVGdpUHo0OGMzWm5JSFpsY25OcGIyNDlJakV1TVNJZ2FXUTlJa3hoZVdWeVh6RWlJSGh0Ykc1elBTSm9kSFJ3T2k4dmQzZDNMbmN6TG05eVp5OHlNREF3TDNOMlp5SWdlRzFzYm5NNmVHeHBibXM5SW1oMGRIQTZMeTkzZDNjdWR6TXViM0puTHpFNU9Ua3ZlR3hwYm1zaUlIZzlJakJ3ZUNJZ2VUMGlNSEI0SWlCMmFXVjNRbTk0UFNJd0lEQWdNVEF3SURFd01DSWdjM1I1YkdVOUltVnVZV0pzWlMxaVlXTnJaM0p2ZFc1a09tNWxkeUF3SURBZ01UQXdJREV3TUR0aVlXTnJaM0p2ZFc1a0xXTnZiRzl5T21Kc1lXTnJPeUlnZUcxc09uTndZV05sUFNKd2NtVnpaWEoyWlNJK1BITjBlV3hsSUhSNWNHVTlJblJsZUhRdlkzTnpJajR1Wm50bWFXeHNPaU14TkRjM1JrSTdmUzVqZTJacGJHdzZJekUwTnpkR1FqdDlMbUo3Wm1sc2JEb2pSa0l4TkRZM08zMHVhSHRtYVd4c09pTXhORGMzUmtJN2ZTNWhiSEJvWVh0bWFXeHNPaU14TkRjM1JrSTdmUzV3WjN0bWFXeHNPaU5HUWpFME5qYzdmUzVsZTJacGJHd3RjblZzWlRwbGRtVnViMlJrTzJOc2FYQXRjblZzWlRwbGRtVnViMlJrTzMwOEwzTjBlV3hsUGp4d1lYUm9JR05zWVhOelBTSm1JaUJrUFNKTk1UWWdNVEJvTmpoMk1rZ3hObnB0TkMwMGFEUjJNbWd0TkhwdE9DQXdhRFIyTW1ndE5IcHRPQ0F3YURSMk1tZ3ROSHB0T0NBd2FEUjJNbWd0TkhwdE9DQXdhRFIyTW1ndE5IcHRPQ0F3YURSMk1tZ3ROSHB0T0NBd2FEUjJNbWd0TkhwdE9DQXdhRFIyTW1ndE5IcE5NVFlnT0Rob05qaDJNa2d4Tm5wdE5DQTBhRFIyTW1ndE5IcHRPQ0F3YURSMk1tZ3ROSHB0T0NBd2FEUjJNbWd0TkhwdE9DQXdhRFIyTW1ndE5IcHRPQ0F3YURSMk1tZ3ROSHB0T0NBd2FEUjJNbWd0TkhwdE9DQXdhRFIyTW1ndE5IcHRPQ0F3YURSMk1tZ3ROSHB0TVRRdE56WjJOamhvTFRKV01UWjZiVFFnTkhZMGFDMHlkaTAwZW0wd0lEaDJOR2d0TW5ZdE5IcHRNQ0E0ZGpSb0xUSjJMVFI2YlRBZ09IWTBhQzB5ZGkwMGVtMHdJRGgyTkdndE1uWXROSHB0TUNBNGRqUm9MVEoyTFRSNmJUQWdPSFkwYUMweWRpMDBlbTB3SURoMk5HZ3RNbll0TkhwTk1USWdNVFoyTmpob0xUSldNVFo2YlMwMElEUjJORWcyZGkwMGVtMHdJRGgyTkVnMmRpMDBlbTB3SURoMk5FZzJkaTAwZW0wd0lEaDJORWcyZGkwMGVtMHdJRGgyTkVnMmRpMDBlbTB3SURoMk5FZzJkaTAwZW0wd0lEaDJORWcyZGkwMGVtMHdJRGgyTkVnMmRpMDBlaUl2UGp4emRtY2dZMnhoYzNNOUltTWlQanh3WVhSb0lHUTlJazB4T0NBNE1tZ3lWakU0YUMweWVtMHdMVFl5YUMweWRpMHlhREo2YlRBZ05HZ3RNbll0TW1neWVtMHdJRFJvTFRKMkxUSm9NbnB0TUNBMGFDMHlkaTB5YURKNmJUQWdOR2d0TW5ZdE1tZ3llbTB3SURSb0xUSjJMVEpvTW5wdE1DQTBhQzB5ZGkweWFESjZiVEFnTkdndE1uWXRNbWd5ZW0wd0lEUm9MVEoyTFRKb01ucHRNQ0EwYUMweWRpMHlhREo2YlRBZ05HZ3RNbll0TW1neWVtMHdJRFJvTFRKMkxUSm9NbnB0TUNBMGFDMHlkaTB5YURKNmJUQWdOR2d0TW5ZdE1tZ3llbTB3SURSb0xUSjJMVEpvTW5wdE1DQTBhQzB5ZGkweWFESjZiVFkySURKMkxUSklNakIyTW5wdExUWXlJREIyTW1ndE1uWXRNbnB0TkNBd2RqSm9MVEoyTFRKNmJUUWdNSFl5YUMweWRpMHllbTAwSURCMk1tZ3RNbll0TW5wdE5DQXdkakpvTFRKMkxUSjZiVFFnTUhZeWFDMHlkaTB5ZW0wMElEQjJNbWd0TW5ZdE1ucHROQ0F3ZGpKb0xUSjJMVEo2YlRRZ01IWXlhQzB5ZGkweWVtMDBJREIyTW1ndE1uWXRNbnB0TkNBd2RqSm9MVEoyTFRKNmJUUWdNSFl5YUMweWRpMHllbTAwSURCMk1tZ3RNbll0TW5wdE5DQXdkakpvTFRKMkxUSjZiVFFnTUhZeWFDMHlkaTB5ZW0wMElEQjJNbWd0TW5ZdE1ub2lMejQ4Y0dGMGFDQmtQU0pOT0RJZ01UaG9MVEoyTmpSb01ucHRNQ0EyTW1neWRqSm9MVEo2YlRBdE5HZ3lkakpvTFRKNmJUQXROR2d5ZGpKb0xUSjZiVEF0TkdneWRqSm9MVEo2YlRBdE5HZ3lkakpvTFRKNmJUQXROR2d5ZGpKb0xUSjZiVEF0TkdneWRqSm9MVEo2YlRBdE5HZ3lkakpvTFRKNmJUQXROR2d5ZGpKb0xUSjZiVEF0TkdneWRqSm9MVEo2YlRBdE5HZ3lkakpvTFRKNmJUQXROR2d5ZGpKb0xUSjZiVEF0TkdneWRqSm9MVEo2YlRBdE5HZ3lkakpvTFRKNmJUQXROR2d5ZGpKb0xUSjZiVEF0TkdneWRqSm9MVEo2SWk4K1BIQmhkR2dnWkQwaVRUZ3lJREU0ZGpKSU1UaDJMVEo2YlMwMk1pQXdkaTB5YUMweWRqSjZiVFFnTUhZdE1tZ3RNbll5ZW0wMElEQjJMVEpvTFRKMk1ucHROQ0F3ZGkweWFDMHlkako2YlRRZ01IWXRNbWd0TW5ZeWVtMDBJREIyTFRKb0xUSjJNbnB0TkNBd2RpMHlhQzB5ZGpKNmJUUWdNSFl0TW1ndE1uWXllbTAwSURCMkxUSm9MVEoyTW5wdE5DQXdkaTB5YUMweWRqSjZiVFFnTUhZdE1tZ3RNbll5ZW0wMElEQjJMVEpvTFRKMk1ucHROQ0F3ZGkweWFDMHlkako2YlRRZ01IWXRNbWd0TW5ZeWVtMDBJREIyTFRKb0xUSjJNbnB0TkNBd2RpMHlhQzB5ZGpKNmJTMDJJREpvTW5ZeWFDMHllbTB5SURKb01uWXlhQzB5ZW0weUlESm9Nbll5YUMweWVtMHROQ0ExTm1neWRpMHlhQzB5ZW0weUxUSm9Nbll0TW1ndE1ucHRNaTB5YURKMkxUSm9MVEo2VFRJMklESXdhQzB5ZGpKb01ucHRMVElnTW1ndE1uWXlhREo2YlMweUlESm9MVEoyTW1neWVtMDBJRFUyYUMweWRpMHlhREo2YlMweUxUSm9MVEoyTFRKb01ucHRMVEl0TW1ndE1uWXRNbWd5ZWlJdlBqd3ZjM1puUGp4d1lYUm9JR05zWVhOelBTSmhiSEJvWVNJZ1pEMGlUVFlnTmxZMGFEUjJNbnB0TkNBd2RqSklObFkyU0RSMk5tZ3lkaTB5YURSMk1tZ3lWalo2YlRnMExUSjJNbWd0TkZZMGVtMHdJREoyTW1ndE5GWTJhQzB5ZGpab01uWXRNbWcwZGpKb01sWTJlazB4TUNBNE9IWXlTRFoyTFRKNmJUQWdNbll5U0RaMkxUSklOSFkyYURKMkxUSm9OSFl5YURKMkxUWjZiVGcwTFRKMk1tZ3ROSFl0TW5wdE1DQXlkakpvTFRSMkxUSm9MVEoyTm1neWRpMHlhRFIyTW1neWRpMDJlaUl2UGp4d2IyeDVaMjl1SUdOc1lYTnpQU0ppSWlCd2IybHVkSE05SWpjeUxEUTRJRGN5TERZMElEY3dMRFkwSURjd0xEWTJJRFk0TERZMklEWTRMRFk0SURZMExEWTRJRFkwTERjd0lEWXdMRGN3SURZd0xEY3lJRFF3TERjeUlEUXdMRGN3SURNMkxEY3dJRE0yTERZNElETXlMRFk0SURNeUxEWTJJQ0FnTXpBc05qWWdNekFzTmpRZ01qZ3NOalFnTWpnc05EZ2dNekFzTkRnZ016QXNORFlnTXpJc05EWWdNeklzTkRRZ05qZ3NORFFnTmpnc05EWWdOekFzTkRZZ056QXNORGdnSWk4K1BIQmhkR2dnWkQwaVRUWTJJRFV3U0RVMGRqSm9NVEo2YlMweUxUSklOVFIyTW1neE1IcHRMVEl0TW1ndE5uWXlhRFo2YlRRZ05rZzFOSFl5YURFeWVtMHdJREpJTlRSMk1tZ3hNbnB0TFRJZ01tZ3RPSFl5YURoNklpQm1hV3hzUFNJak1EQXdJaTgrUEhCaGRHZ2daRDBpVFRZeElEVTBhQzAwZGpKb05Ib2lJR1pwYkd3OUlpTkVSVVUxUkRraUx6NDhjR0YwYUNCa1BTSk5NelFnTlRCb01USjJNa2d6TkhwdE1pMHlhREV3ZGpKSU16WjZiVEl0TW1nMmRqSm9MVFo2YlMwMElEWm9NVEoyTWtnek5IcHRNQ0F5YURFeWRqSklNelI2YlRJZ01tZzRkakpvTFRoNklpQm1hV3hzUFNJak1EQXdJaTgrUEhCaGRHZ2daRDBpVFRRd0lEVTBhRFIyTW1ndE5Ib2lJR1pwYkd3OUlpTkVSVVUxUkRraUx6NDhjR0YwYUNCa1BTSk5ORFlnTlRab09IWXlhQzA0ZW0weUlESm9OSFl5YUMwMGVpSWdabWxzYkQwaUl6QXdNQ0l2UGp4d1lYUm9JR1E5SWswek9DQTJNV2d5ZGpKb0xUSjZUVFF3SURZemFEUjJNbWd0TkhwTk5EUWdOakZvTkhZeWFDMDBlazAwT0NBMk0yZzBkakpvTFRSNlRUVXlJRFl4YURSMk1tZ3ROSHBOTlRZZ05qTm9OSFl5YUMwMGVrMDJNQ0EyTVdneWRqSm9MVEo2VFRZMElEWTFhREoyTW1ndE1ucE5NelFnTmpWb01uWXlhQzB5ZWswek5pQTJNMmd5ZGpKb0xUSjZUVFl5SURZemFESjJNbWd0TW5vaUlHWnBiR3c5SWlNd01EQWlMejQ4Y0dGMGFDQmpiR0Z6Y3owaWFDSWdaRDBpVFRRM0lESTBhRFIyTW1ndE5Ib2lMejQ4Y0dGMGFDQmpiR0Z6Y3owaWFDQmxJaUJrUFNKTk56Z2dOREIyTWtnMk0zWXRObWd0TW5ZdE1tZ3RNbll0TW1ndE1uWXRNbWd0TW5ZdE1tZzBkakpvTW5ZeWFESjJNbWd6ZGpKb01uWXlhREoyTW5vaUx6NDhjR0YwYUNCamJHRnpjejBpYUNCbElpQmtQU0pOTlRrZ016WjJMVEpvTFRKMkxUSm9MVEoyTFRKb0xUSjJMVEpvTFRSMk1tZ3RNbll5YUMweWRqSm9MVEoyTFRKb01uWXRNbWd5ZGkweWFDMDRkakpvTFRKMk1tZ3RNbll5YUMwemRqSm9MVEoyTW1ndE1uWTBhRE16ZGkwMmVtMHRPU0EwU0RNNGRpMHlhREoyTFRKb09IWXlhREo2SWk4K1BDOXpkbWMrIiwgImF0dHJpYnV0ZXMiOiBbeyJ0cmFpdF90eXBlIjogIkZyYW1lIiwgInZhbHVlIjogIkZyYW1lIDMifSwgeyJ0cmFpdF90eXBlIjogIkNoaXAgRGV0YWlsIiwgInZhbHVlIjogIiMxNDc3RkIifSwgeyJ0cmFpdF90eXBlIjogIkNoaXAgRGV0YWlsIiwgInZhbHVlIjogIkRldGFpbCAxIn0sIHsidHJhaXRfdHlwZSI6ICJDaGlwIERldGFpbCBDb2xvciIsICJ2YWx1ZSI6ICIjMTQ3N0ZCIn0sIHsidHJhaXRfdHlwZSI6ICJDb3JuZXIiLCAidmFsdWUiOiAiQWxwaGEgTm9kZSJ9LHsidHJhaXRfdHlwZSI6ICJIZWFkIFNoYXBlIiwgInZhbHVlIjogInJvdW5kIn0sIHsidHJhaXRfdHlwZSI6ICJIZWFkIFNoYXBlIENvbG9yIiwgInZhbHVlIjogIiNGQjE0NjcifSx7InRyYWl0X3R5cGUiOiAiRXllcyIsICJ2YWx1ZSI6ICJQYW5kYSJ9LHsidHJhaXRfdHlwZSI6ICJNb3V0aCIsICJ2YWx1ZSI6ICJTY2FyZWQifSx7InRyYWl0X3R5cGUiOiAiSGVhZCBEZXRhaWwiLCAidmFsdWUiOiAiSGF0In0sIHsidHJhaXRfdHlwZSI6ICJIZWFkIERldGFpbCBDb2xvciIsICJ2YWx1ZSI6ICIjMTQ3N0ZCIn1dfQ==",
						},
					},
					{
						Type:     filter.TypeExchangeStaking,
						Platform: filter.PlatformRSS3.String(),
						From:     "0x30286DD245338292F319809935a1037CcD4573Ea",
						To:       "0x39F9e912C1F696F533e7A2267Ea233AeC9742b35",
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
