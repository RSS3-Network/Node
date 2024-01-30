package proxy_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	"github.com/rss3-network/node/provider/ethereum/proxy"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/stretchr/testify/require"
)

func TestGetImplementation(t *testing.T) {
	t.Parallel()

	type arguments struct {
		ctx         context.Context
		network     filter.Network
		address     common.Address
		blockNumber *big.Int
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      require.ValueAssertionFunc
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "Ethereum USD Coin",
			arguments: arguments{
				ctx:     context.Background(),
				network: filter.NetworkEthereum,
				address: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
			},
			want: func(t require.TestingT, value interface{}, msgAndArgs ...interface{}) {
				address, ok := value.(*common.Address)
				require.True(t, ok)

				require.Equal(t, common.HexToAddress("0x43506849D7C04F9138D1A2050bbF3A0c054402dd"), *address)
			},
			wantError: require.NoError,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			ethereumClient, err := ethereum.Dial(testcase.arguments.ctx, endpoint.MustGet(testcase.arguments.network))
			testcase.wantError(t, err)

			result, err := proxy.GetImplementation(testcase.arguments.ctx, testcase.arguments.address, testcase.arguments.blockNumber, ethereumClient)
			testcase.wantError(t, err)

			testcase.want(t, result)
		})
	}
}
