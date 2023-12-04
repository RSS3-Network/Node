package arweave

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	syncx "github.com/naturalselectionlabs/rss3-node/common/sync"
	"github.com/samber/lo"
)

// DefaultClient is the default client.
var DefaultClient = lo.Must(NewClient())

const (
	DefaultTimeout = 3 * time.Second
)

// Client is the interface that wraps the Fetch method.
type Client interface {
	GetTransactionData(ctx context.Context, id string) (io.ReadCloser, error)
}

// Ensure that client implements Client.
var _ Client = (*client)(nil)

// client is the default implementation of Client.
type client struct {
	httpClient *http.Client
	gateways   []string
	locker     sync.RWMutex
}

// GetTransactionData fetches the transaction data of the given id from arweave network.
func (c *client) GetTransactionData(ctx context.Context, id string) (io.ReadCloser, error) {
	c.locker.RLock()
	defer c.locker.RUnlock()

	quickGroup := syncx.NewQuickGroup[io.ReadCloser](ctx)

	// Try to fetch from all gateways.
	for _, gateway := range c.gateways {
		gateway := gateway

		quickGroup.Go(func(ctx context.Context) (io.ReadCloser, error) {
			requestURL, err := url.JoinPath(gateway, id)
			if err != nil {
				return nil, fmt.Errorf("invalid gateway url: %w", err)
			}

			data, err := c.do(ctx, http.MethodGet, requestURL, nil)
			if err != nil {
				return nil, fmt.Errorf("fetch from gateway %s: %w", gateway, err)
			}

			return data, nil
		})
	}

	result, err := quickGroup.Wait()
	if err != nil {
		return nil, fmt.Errorf("fetch from all gateways: %w", err)
	}

	return result, nil
}

// do send an HTTP request with the given method, url and body.
func (c *client) do(ctx context.Context, method string, url string, body io.Reader) (io.ReadCloser, error) {
	request, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected status code %d", response.StatusCode)
	}

	return response.Body, nil
}

// ClientOption is the type of the options passed to NewClient.
type ClientOption func(client *client) error

// WithGateways sets the client with Arweave gateways.
func WithGateways(gateways []string) ClientOption {
	return func(c *client) error {
		c.gateways = gateways

		return nil
	}
}

// WithTimeout sets the client with the given timeout.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(h *client) error {
		h.httpClient.Timeout = timeout

		return nil
	}
}

// NewClient creates a new Client with the given options.
func NewClient(options ...ClientOption) (Client, error) {
	instance := client{
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		gateways: DefaultGateways,
	}

	for _, option := range options {
		if err := option(&instance); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return &instance, nil
}
