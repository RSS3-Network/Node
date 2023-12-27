package arweave

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

// DefaultClient is the default client.
var DefaultClient = lo.Must(NewClient())

const (
	DefaultTimeout = 3 * time.Second
)

// Client is the interface that wraps the Fetch method.
type Client interface {
	GetTransactionData(ctx context.Context, id string) (io.ReadCloser, error)
	GetBlockHeight(ctx context.Context) (blockHeight int64, err error)
	GetBlockByHeight(ctx context.Context, height int64) (block *Block, err error)
	GetTransactionByID(ctx context.Context, id string) (transaction *Transaction, err error)
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
	return c.queryArweaveByRoute(ctx, id)
}

// GetBlockHeight returns the current block height of the arweave network.
func (c *client) GetBlockHeight(ctx context.Context) (blockHeight int64, err error) {
	data, err := c.queryArweaveByRoute(ctx, "info")
	if err != nil {
		return 0, fmt.Errorf("query arweave by route: %w", err)
	}

	// close the response body when the function returns.
	defer lo.Try(data.Close)

	network := Network{}

	content, err := io.ReadAll(data)
	if err != nil {
		return 0, err
	}

	err = json.Unmarshal(content, &network)
	if err != nil {
		return 0, fmt.Errorf("unmarshal response body: %w", err)
	}

	return network.Blocks, nil
}

// GetBlockByHeight returns block by specific block height.
func (c *client) GetBlockByHeight(ctx context.Context, height int64) (block *Block, err error) {
	data, err := c.queryArweaveByRoute(ctx, fmt.Sprintf("block/height/%d", height))
	if err != nil {
		return nil, fmt.Errorf("query arweave by route: %w", err)
	}

	// close the response body when the function returns.
	defer lo.Try(data.Close)

	content, err := io.ReadAll(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &block)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response body: %w", err)
	}

	return block, nil
}

// GetTransactionByID returns transaction by specific transaction id
func (c *client) GetTransactionByID(ctx context.Context, id string) (transaction *Transaction, err error) {
	data, err := c.queryArweaveByRoute(ctx, fmt.Sprintf("tx/%s", id))
	if err != nil {
		return nil, fmt.Errorf("query arweave by route: %w", err)
	}

	// close the response body when the function returns.
	defer lo.Try(data.Close)

	content, err := io.ReadAll(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &transaction)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response body: %w", err)
	}

	return transaction, nil
}

// queryArweaveByRoute queries the arweave network by the given route.
func (c *client) queryArweaveByRoute(ctx context.Context, path string) (io.ReadCloser, error) {
	c.locker.RLock()
	defer c.locker.RUnlock()

	var data io.ReadCloser

	var err error

	retryableFunc := func() error {
		for _, gateway := range c.gateways {
			requestURL, err := url.JoinPath(gateway, path)
			if err != nil {
				return fmt.Errorf("invalid gateway url: %w", err)
			}

			data, err = c.do(ctx, http.MethodGet, requestURL, nil)
			if err == nil {
				// Successful response received, return the data.
				return nil
			}
		}

		// If all gateways have been tried and none succeeded, return the last error.
		return fmt.Errorf("fetch from all gateways failed: %w", err)
	}

	onRetryFunc := func(err error, duration time.Duration) {
		// Log the error and the backoff duration.
		zap.L().Error("retry failed", zap.Error(err), zap.Duration("duration", duration))
	}

	// Create a new exponential backoff policy
	backoffPolicy := backoff.NewExponentialBackOff()
	backoffPolicy.InitialInterval = 500 * time.Millisecond

	err = backoff.RetryNotify(retryableFunc, backoffPolicy, onRetryFunc)

	if err != nil {
		// Log the error after all retries have failed.
		zap.L().Error("retry failed", zap.Error(err))
		return nil, err
	}

	return data, nil
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
