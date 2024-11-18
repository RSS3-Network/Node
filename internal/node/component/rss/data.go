package rss

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

const (
	// SUFFIX is the parameter used to request data in the RSS3 format from RSSHub
	// https://docs.rsshub.app/guide/parameters#output-formats
	SUFFIX = "rss3"
)

// getActivities fetches data from an RSSHub and returns the transformed response in the Activity format.
func (h *Component) getActivities(ctx context.Context, path string, url *url.URL) ([]*activityx.Activity, error) {
	request, err := h.formatRequest(path, url)
	if err != nil {
		return nil, fmt.Errorf("format request: %w", err)
	}

	zap.L().Debug("sending request to RSSHub",
		zap.String("url", request.String()))

	//nolint:bodyclose // False positive
	response, err := h.getResponse(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to get RSSHub response: %w", err)
	}

	zap.L().Debug("successfully received response from RSSHub",
		zap.Int("status_code", response.StatusCode))

	return h.formatResponse(ctx, response)
}

// formatRequest prepares the URL for the RSSHub request with necessary parameters and authentication.
func (h *Component) formatRequest(path string, in *url.URL) (*url.URL, error) {
	out, err := url.Parse(h.rsshub.endpoint)
	if err != nil {
		return nil, fmt.Errorf("parse RSSHub endpoint: %w", err)
	}

	parameters := in.Query()
	parameters.Set("format", SUFFIX)

	if h.rsshub.accessKey != "" {
		parameters.Set("key", h.rsshub.accessKey)
		in.RawQuery = parameters.Encode()
	}

	out.RawQuery = parameters.Encode()
	out.Path = path

	return out, nil
}

// getResponse sends an HTTP request to the RSSHub and returns the response.
func (h *Component) getResponse(ctx context.Context, url *url.URL) (*http.Response, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create HTTP request: %w", err)
	}

	response, err := h.httpClient.Do(request)

	if err != nil {
		return nil, fmt.Errorf("HTTP request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("RSSHub status code: %d", response.StatusCode)
	}

	return response, nil
}

// formatResponse parses the RSS response into activities.
func (h *Component) formatResponse(ctx context.Context, response *http.Response) ([]*activityx.Activity, error) {
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	defer lo.Try(response.Body.Close)

	activities, err := h.transformResponse(ctx, data)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

// transformResponse converts the RSS response into the Activity format.
func (h *Component) transformResponse(_ context.Context, data []byte) ([]*activityx.Activity, error) {
	type response struct {
		Data activityx.Activities `json:"data"`
	}

	var resp response

	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal RSSHub response: %w", err)
	}

	return resp.Data, nil
}
