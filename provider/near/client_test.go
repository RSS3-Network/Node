package near_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/rss3-network/node/provider/near"
	"github.com/stretchr/testify/require"
)

func TestNearClient(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	endpoint := "https://rpc.mainnet.near.org/"

	client, err := near.Dial(ctx, endpoint)
	require.NoError(t, err)

	t.Run("BlockByNumber", func(t *testing.T) {
		t.Parallel()

		blockNumber := big.NewInt(127103648)
		block, err := client.BlockByNumber(ctx, blockNumber)

		require.NoError(t, err)
		require.NotNil(t, block)
		require.Equal(t, 127103648, block.Header.Height)
	})

	t.Run("ChunkByHash", func(t *testing.T) {
		t.Parallel()

		chunkID := "GWAab7TDaf5T8dfgJ2YrHHXgdEpJUdUS1Hmie2N2aWhw"
		chunk, err := client.ChunkByHash(ctx, chunkID)

		require.NoError(t, err)
		require.NotNil(t, chunk)
		require.Equal(t, chunkID, chunk.Header.ChunkHash)
	})

	t.Run("TransactionByHash", func(t *testing.T) {
		t.Parallel()

		txHash := "9JCqrdN53VXiAVo9feEHY7FigiV1QYSyYuoXeTuee4eA"
		senderID := "pseudoyu.tg"
		tx, err := client.TransactionByHash(ctx, txHash, senderID)

		require.NoError(t, err)
		require.NotNil(t, tx)
		require.Equal(t, txHash, tx.Transaction.Hash)
	})
}
