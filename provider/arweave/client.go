package arweave

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/samber/lo"
)

// DefaultClient is the default client.
var DefaultClient = lo.Must(NewClient())

var ErrorNotFound = errors.New("not found")

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

	var latestError error

	for _, gateway := range c.gateways {
		requestURL, err := url.JoinPath(gateway, path)
		if err != nil {
			return nil, fmt.Errorf("invalid gateway url: %w", err)
		}

		readCloser, code, err := c.do(ctx, http.MethodGet, requestURL, nil)
		if err == nil {
			// Successful response received, return the data.
			return readCloser, err
		}

		// Record the latest error.
		latestError = err

		if gateway == GatewayArweave.String() && code == http.StatusNotFound {
			// If the gateway is the official Arweave gateway and the response is 404, return error.
			return nil, ErrorNotFound
		}
	}

	// If all gateways have been tried and none succeeded, return the last error.
	return nil, fmt.Errorf("fetch from all gateways failed: %w", latestError)
}

// do send an HTTP request with the given method, url and body.
func (c *client) do(ctx context.Context, method string, url string, body io.Reader) (io.ReadCloser, int, error) {
	request, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, 0, fmt.Errorf("create request: %w", err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, 0, fmt.Errorf("send request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		marshal, _ := json.Marshal(response.Body)

		return nil, response.StatusCode, fmt.Errorf("unexpected status code: %d, response: %s, url: %s", response.StatusCode, string(marshal), url)
	}

	return response.Body, 0, nil
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
