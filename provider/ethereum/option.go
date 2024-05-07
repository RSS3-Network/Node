package ethereum

import (
	"net/http"
	"strings"
)

// Option used to configure the client with additional options.
type Option = func(client *client) error

// WithHTTPHeader sets the HTTP header for the client.
// The header can be used to set the authorization token and other headers.
func WithHTTPHeader(httpHeader http.Header) Option {
	return func(client *client) error {
		for key, value := range httpHeader {
			client.rpcClient.SetHeader(key, strings.Join(value, ","))
		}

		return nil
	}
}
