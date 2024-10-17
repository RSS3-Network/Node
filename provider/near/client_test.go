package near_test

// func TestNearClient(t *testing.T) {
//	t.Parallel()
//
//	ctx := context.Background()
//	endpoint := "https://archival-rpc.mainnet.near.org"
//
//	client, err := near.Dial(ctx, endpoint)
//	require.NoError(t, err)
//
//	t.Run("GetBlockHeight", func(t *testing.T) {
//		t.Parallel()
//
//		blockHeight, err := client.GetBlockHeight(ctx)
//
//		require.NoError(t, err)
//		require.Greater(t, blockHeight, int64(0))
//	})
//
//	t.Run("BlockByHeight", func(t *testing.T) {
//		t.Parallel()
//
//		blockHeight := big.NewInt(127103648)
//		block, err := client.BlockByHeight(ctx, blockHeight)
//
//		require.NoError(t, err)
//		require.NotNil(t, block)
//		require.Equal(t, 127103648, block.Header.Height)
//	})
//
//	t.Run("ChunkByHash", func(t *testing.T) {
//		t.Parallel()
//
//		chunkID := "GWAab7TDaf5T8dfgJ2YrHHXgdEpJUdUS1Hmie2N2aWhw"
//		chunk, err := client.ChunkByHash(ctx, chunkID)
//
//		require.NoError(t, err)
//		require.NotNil(t, chunk)
//		require.Equal(t, chunkID, chunk.Header.ChunkHash)
//	})
//
//	t.Run("ChunkByHeight", func(t *testing.T) {
//		t.Parallel()
//
//		blockHeight := big.NewInt(127170738)
//		shardID := 0
//		chunk, err := client.ChunkByHeight(ctx, blockHeight, shardID)
//
//		require.NoError(t, err)
//		require.NotNil(t, chunk)
//	})
//
//	t.Run("TransactionByHash", func(t *testing.T) {
//		t.Parallel()
//
//		txHash := "9JCqrdN53VXiAVo9feEHY7FigiV1QYSyYuoXeTuee4eA"
//		senderID := "pseudoyu.tg"
//		tx, err := client.TransactionByHash(ctx, txHash, senderID)
//
//		require.NoError(t, err)
//		require.NotNil(t, tx)
//		require.Equal(t, txHash, tx.Transaction.Hash)
//	})
//}
