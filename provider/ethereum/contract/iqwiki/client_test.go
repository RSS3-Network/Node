package iqwiki_test

import (
	"context"
	"testing"

	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/iqwiki"
	"github.com/stretchr/testify/require"
)

func Test_Client(t *testing.T) {
	t.Parallel()

	client, err := iqwiki.NewClient(iqwiki.Endpoint)

	require.NoError(t, err, "build IQ.Wiki client")

	wikiResponse, err := iqwiki.ActivityByWikiIdAndBlock(context.Background(), client, 46938791, "buffer")

	require.NoError(t, err)

	t.Log(wikiResponse.ActivityByWikiIdAndBlock.Ipfs)
}
