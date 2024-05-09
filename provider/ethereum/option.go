package ethereum

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/ethereum/go-ethereum/rpc"
)

// Option used to configure the client with additional options.
type Option = func(ctx context.Context, client *client) error

// WithHTTP2Disabled sets the HTTP2 client to be disabled.
func WithHTTP2Disabled() Option {
	return func(ctx context.Context, client *client) error {
		tr := http.DefaultTransport.(*http.Transport).Clone()
		tr.ForceAttemptHTTP2 = false
		tr.TLSNextProto = make(map[string]func(authority string, c *tls.Conn) http.RoundTripper)
		tr.TLSClientConfig = new(tls.Config)

		httpClient := &http.Client{
			Transport: tr,
		}

		var err error

		client.rpcClient, err = rpc.DialOptions(ctx, client.endpoint, rpc.WithHTTPClient(httpClient))
		if err != nil {
			return err
		}

		return nil
	}
}

// WithHTTPHeader sets the HTTP header for the client.
// The header can be used to set the authorization token and other headers.
func WithHTTPHeader(httpHeader map[string]string) Option {
	return func(_ context.Context, client *client) error {
		for key, value := range httpHeader {
			client.rpcClient.SetHeader(key, value)
		}

		return nil
	}
}
