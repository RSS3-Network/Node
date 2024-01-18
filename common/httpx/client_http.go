package httpx

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

var _ Client = (*client)(nil)

type client struct {
	httpClient *resty.Client
}

func (c *client) Get(ctx context.Context, url string, headers http.Header, params url.Values) (io.ReadCloser, error) {
	response, err := c.httpClient.R().SetHeaderMultiValues(headers).SetQueryParamsFromValues(params).SetContext(ctx).Get(url)
	if err != nil {
		return nil, fmt.Errorf("http get content error: %w", err)
	}

	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode())
	}

	return io.NopCloser(bytes.NewReader(response.Body())), nil
}

func (c *client) GetContentType(ctx context.Context, url string) (string, error) {
	response, err := c.httpClient.R().SetContext(ctx).SetHeader("User-Agent", "curl/7.86.0").Get(url)
	if err != nil {
		return "", fmt.Errorf("fetch content: %w", err)
	}

	if response.StatusCode() != 200 {
		return "", fmt.Errorf("unexpected status code: %d", response.StatusCode())
	}

	mimeType := mimetype.Detect(response.Body())

	if mimeType == nil {
		return "", fmt.Errorf("can not detect mime type %s", url)
	}

	return mimeType.String(), nil
}

type ClientOption func(*client) error

func WithRetryCount(retryCount int) ClientOption {
	return func(c *client) error {
		c.httpClient.SetRetryCount(retryCount)

		return nil
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *client) error {
		c.httpClient.SetTimeout(timeout)

		return nil
	}
}

func NewClient(options ...ClientOption) (Client, error) {
	httpClient := client{
		httpClient: resty.New().SetTimeout(DefaultTimeout).SetRetryCount(DefaultRetryCount).SetLogger(zap.L().Sugar()),
	}

	for _, option := range options {
		if err := option(&httpClient); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return &httpClient, nil
}
