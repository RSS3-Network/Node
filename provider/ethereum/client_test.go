package ethereum_test

import (
	"context"
	"math"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethereumx "github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	networkx "github.com/rss3-network/protocol-go/schema/network"
	"github.com/stretchr/testify/require"
)

func TestGetNoExistentData(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	testcases := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "Get non-existent block by number",
			test: func(t *testing.T) {
				ethereumClient, err := ethereumx.Dial(ctx, endpoint.MustGet(networkx.Ethereum))
				require.NoError(t, err)

				blockNumber := new(big.Int).SetUint64(math.MaxInt64 - 1)

				_, err = ethereumClient.BlockByNumber(ctx, blockNumber)
				require.ErrorIs(t, err, ethereum.NotFound)

				_, err = ethereumClient.BatchBlockByNumbers(ctx, []*big.Int{blockNumber})
				require.ErrorIs(t, err, ethereum.NotFound)
			},
		},
		{
			name: "Get non-existent transaction by hash",
			test: func(t *testing.T) {
				ethereumClient, err := ethereumx.Dial(ctx, endpoint.MustGet(networkx.Ethereum))
				require.NoError(t, err)

				transactionHash := common.MaxHash

				_, err = ethereumClient.TransactionByHash(ctx, transactionHash)
				require.ErrorIs(t, err, ethereum.NotFound)
			},
		},
		{
			name: "Get non-existent transaction receipt by hash",
			test: func(t *testing.T) {
				ethereumClient, err := ethereumx.Dial(ctx, endpoint.MustGet(networkx.Ethereum))
				require.NoError(t, err)

				blockHash := common.MaxHash

				_, err = ethereumClient.TransactionReceipt(ctx, blockHash)
				require.ErrorIs(t, err, ethereum.NotFound)

				_, err = ethereumClient.BatchTransactionReceipt(ctx, []common.Hash{blockHash})
				require.ErrorIs(t, err, ethereum.NotFound)
			},
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			testcase.test(t)
		})
	}
}
