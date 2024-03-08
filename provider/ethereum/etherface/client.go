package etherface

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/samber/lo"
	"github.com/tidwall/gjson"
)

var (
	ErrorNoResults = errors.New("no results")
)

const (
	EtherfaceEndpoint = "https://api.etherface.io/v1"
	DefaultTimeout    = 3 * time.Second
	DefaultAttempts   = 3
)

type Client interface {
	Lookup(ctx context.Context, hash string) (string, error)
}

var _ Client = (*etherfaceClient)(nil)

type etherfaceClient struct {
	httpClient *http.Client
	attempts   uint
}

func (h *etherfaceClient) Lookup(ctx context.Context, hash string) (functionName string, err error) {
	retryableFunc := func() error {
		functionName, err = h.fetch(ctx, hash)
		return err
	}

	retryIfFunc := func(err error) bool {
		return !errors.Is(err, ErrorNoResults)
	}

	if err := retry.Do(retryableFunc, retry.Attempts(h.attempts), retry.RetryIf(retryIfFunc)); err != nil {
		return "", fmt.Errorf("retry attempts: %w", err)
	}

	return functionName, nil
}

func (h *etherfaceClient) fetch(ctx context.Context, hash string) (string, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/signatures/hash/function/%s/1", EtherfaceEndpoint, hash), nil)
	if err != nil {
		return "", fmt.Errorf("new request: %w", err)
	}

	// nolint:bodyclose // False positive
	response, err := h.httpClient.Do(request)
	if err != nil {
		return "", fmt.Errorf("send request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		defer lo.Try(response.Body.Close)

		return "", fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	defer lo.Try(response.Body.Close)

	respData, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("read all: %w", err)
	}

	etherfaceData := gjson.ParseBytes(respData)

	items := etherfaceData.Get("items").Array()

	if len(items) > 0 {
		return extractFunctionName(items[0].Get("text").String()), nil
	}

	return "", nil
}

// extractFunctionName extracts the function name from a string before the first bracket
func extractFunctionName(str string) string {
	index := strings.Index(str, "(")
	if index != -1 {
		return str[:index]
	}

	return ""
}

func NewEtherfaceClient(options ...ClientOption) (Client, error) {
	instance := etherfaceClient{
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		attempts: DefaultAttempts,
	}

	for _, option := range options {
		if err := option(&instance); err != nil {
			return nil, fmt.Errorf("apply options: %w", err)
		}
	}

	return &instance, nil
}

type ClientOption func(client *etherfaceClient) error

func WithAttempts(attempts uint) ClientOption {
	return func(h *etherfaceClient) error {
		h.attempts = attempts

		return nil
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(h *etherfaceClient) error {
		h.httpClient.Timeout = timeout

		return nil
	}
}
