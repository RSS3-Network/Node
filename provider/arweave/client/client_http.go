package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var _ HTTPClient = (*httpClient)(nil)

type httpClient struct {
	httpClient   *http.Client
	endpointURLs []*url.URL
}

func (h *httpClient) Fetch(ctx context.Context, id string) (io.ReadCloser, error) {
	for _, endpoint := range h.endpointURLs {
		response, err := h.fetch(ctx, endpoint, id)
		if err == nil {
			return response, nil
		}
	}

	return nil, fmt.Errorf("failed to fetch %s", id)
}

func (h *httpClient) fetch(ctx context.Context, endpointURL *url.URL, id string) (io.ReadCloser, error) {
	requestURL := endpointURL
	requestURL.Path = id

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	return response.Body, nil
}

type HTTPClientOption func(*httpClient) error

func WithEndpointURLs(endpoints []string) HTTPClientOption {
	return func(h *httpClient) error {
		for _, endpoint := range endpoints {
			endpointURL, err := url.Parse(endpoint)
			if err != nil {
				return fmt.Errorf("invalid endpoint %s: %w", endpoint, err)
			}

			h.endpointURLs = append(h.endpointURLs, endpointURL)
		}

		return nil
	}
}

func NewHTTPClient(options ...HTTPClientOption) (HTTPClient, error) {
	client := httpClient{
		httpClient:   http.DefaultClient,
		endpointURLs: make([]*url.URL, 0),
	}

	if err := WithEndpointURLs(EndpointStrings())(&client); err != nil {
		return nil, fmt.Errorf("apply default option: %w", err)
	}

	for _, option := range options {
		if err := option(&client); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return &client, nil
}
