package httpx_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/naturalselectionlabs/rss3-node/provider/httpx"
	"github.com/naturalselectionlabs/rss3-node/provider/ipfs"
	"github.com/stretchr/testify/require"
)

var (
	setupOnce  sync.Once
	httpClient httpx.Client
)

func setup(t *testing.T) {
	setupOnce.Do(func() {
		var err error

		httpClient, err = httpx.NewHTTPClient()
		require.NoError(t, err)
	})
}

func TestHTTPClient_Fetch(t *testing.T) {
	t.Parallel()

	setup(t)

	type arguments struct {
		url string
	}

	testcases := []struct {
		name      string
		arguments arguments
	}{
		{
			name: "Fetch Arweave",
			arguments: arguments{
				url: "https://arweave.net/aMAYipJXf9rVHnwRYnNF7eUCxBc1zfkaopBt5TJwLWw",
			},
		},
		{
			name: "Fetch External Api",
			arguments: arguments{
				url: "https://data.lens.phaver.com/api/lens/posts/1fdcc7ce-91a7-4af7-8022-13132842a5ec",
			},
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			_, err := httpClient.Fetch(context.TODO(), testcase.arguments.url)
			require.NoError(t, err)
		})
	}
}

func TestHttpClient_GetContentType(t *testing.T) {
	t.Parallel()

	type arguments struct {
		ctx     context.Context
		options []ipfs.HTTPClientOption
		uri     string
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      string
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "mime type jpeg",
			arguments: arguments{
				ctx: context.Background(),
				options: []ipfs.HTTPClientOption{
					ipfs.WithTimeout(10 * time.Second),
				},
				uri: "https://i.imgur.com/DsrBswx.jpg",
			},
			want:      "image/jpeg",
			wantError: require.NoError,
		},
		{
			name: "mime type text/html",
			arguments: arguments{
				ctx: context.Background(),
				options: []ipfs.HTTPClientOption{
					ipfs.WithTimeout(10 * time.Second),
				},
				uri: "http://rss3.io",
			},
			want:      "text/html; charset=utf-8",
			wantError: require.NoError,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			httpClient, err := ipfs.NewHTTPClient(testcase.arguments.options...)
			testcase.wantError(t, err)

			result, err := httpClient.GetContentType(testcase.arguments.ctx, testcase.arguments.uri)
			testcase.wantError(t, err)
			require.Equal(t, testcase.want, result)
		})
	}
}
