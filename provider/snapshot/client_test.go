package snapshot_test

import (
	"context"
	"os"
	"testing"

	"github.com/rss3-network/node/provider/snapshot"
	"github.com/stretchr/testify/require"
)

func Test_Client(t *testing.T) {
	t.Parallel()

	options := make([]snapshot.Option, 0)

	if apiKey := os.Getenv("SNAPSHOT_API_KEY"); apiKey != "" {
		options = append(options, snapshot.WithAPIKey(apiKey))
	}

	client, err := snapshot.NewClient(snapshot.EndpointMainnet, options...)
	require.NoError(t, err, "build snapshot client")

	proposalsResponse, err := snapshot.Proposals(context.Background(), client, 0)
	require.NoError(t, err)
	require.NotEmpty(t, proposalsResponse.Proposals)

	votesResponse, err := snapshot.Votes(context.Background(), client, "0x25a196ede90368c66da7a5d8de1edd6dac49337a33861215647a4478db91a4ad", 0)
	require.NoError(t, err)
	require.NotEmpty(t, votesResponse.Votes)
}
