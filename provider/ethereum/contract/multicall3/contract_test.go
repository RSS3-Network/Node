package multicall3_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/lens"
	"github.com/rss3-network/node/provider/ethereum/contract/multicall3"
	"github.com/rss3-network/node/provider/ethereum/contract/weth"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestAggregate(t *testing.T) {
	t.Parallel()

	lensABI, err := lens.V1LensHubMetaData.GetAbi()
	require.NoError(t, err)

	wethABI, err := weth.WETH9MetaData.GetAbi()
	require.NoError(t, err)

	type arguments struct {
		chainID     filter.EthereumChainID
		blockNumber *big.Int
		calls       []multicall3.Multicall3Call3
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      require.ValueAssertionFunc
	}{
		{
			name: "Get Lens profile metadata",
			arguments: arguments{
				chainID: filter.EthereumChainIDPolygon,
				calls: []multicall3.Multicall3Call3{
					{
						Target:       lens.AddressLensProtocol,
						AllowFailure: false,
						CallData:     lo.Must(lensABI.Pack("ownerOf", big.NewInt(33956))), // ownerOf(kallydev.lens)
					},
					{
						Target:       lens.AddressLensProtocol,
						AllowFailure: false,
						CallData:     lo.Must(lensABI.Pack("ownerOf", big.NewInt(1))), // ownerOf(lensprotocol.lens)
					},
					{
						Target:       lens.AddressLensProtocol,
						AllowFailure: false,
						CallData:     lo.Must(lensABI.Pack("getGovernance")), // getGovernance()
					},
				},
			},
			want: func(t require.TestingT, i any, i2 ...any) {
				results, ok := i.([]*multicall3.Multicall3Result)
				require.True(t, ok)

				for _, result := range results {
					require.True(t, result.Success)
				}
			},
		},
		{
			name: "Get WETH token metadata",
			arguments: arguments{
				chainID:     filter.EthereumChainIDMainnet,
				blockNumber: new(big.Int).SetUint64(4719568),
				calls: []multicall3.Multicall3Call3{
					{
						Target:       weth.AddressMainnet,
						AllowFailure: false,
						CallData:     lo.Must(wethABI.Pack("name")), // ownerOf()
					},
					{
						Target:       weth.AddressMainnet,
						AllowFailure: false,
						CallData:     lo.Must(wethABI.Pack("symbol")), // symbol()
					},
					{
						Target:       weth.AddressMainnet,
						AllowFailure: false,
						CallData:     lo.Must(wethABI.Pack("decimals")), // decimals()
					},
				},
			},
			want: func(t require.TestingT, i any, i2 ...any) {
				results, ok := i.([]*multicall3.Multicall3Result)
				require.True(t, ok)

				for _, result := range results {
					require.True(t, result.Success)
				}
			},
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			network, err := filter.NetworkString(testcase.arguments.chainID.String())
			require.NoError(t, err)

			ethereumClient, err := ethereum.Dial(ctx, endpoint.MustGet(network))
			require.NoError(t, err)

			results, err := multicall3.Aggregate3(ctx, uint64(testcase.arguments.chainID), testcase.arguments.calls, testcase.arguments.blockNumber, ethereumClient)
			require.NoError(t, err)

			testcase.want(t, results)
		})
	}
}
