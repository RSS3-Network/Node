package snapshot_test

import (
	"context"
	"os"
	"testing"

	"github.com/naturalselectionlabs/rss3-node/provider/snapshot"
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

	for _, proposal := range proposalsResponse.Proposals {
		t.Log("proposal", proposal)
	}

	votesResponse, err := snapshot.Votes(context.Background(), client, "0xd8e77a881d93bdcc7df454f24d350068895d90cc4c1bd37dc48e1c9a285c24a4", 0)
	require.NoError(t, err)

	for _, vote := range votesResponse.Votes {
		t.Log("vote", vote)
	}
}
