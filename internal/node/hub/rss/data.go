package rss

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/samber/lo"
)

const (
	RSSHubUMSPath = ".ums"
)

func (h *Hub) getRSSHubData(ctx context.Context, path string, rawQuery string) ([]byte, error) {
	request, err := url.Parse(h.rsshub.Endpoint)
	if err != nil {
		return nil, err
	}

	request.Path = path
	request.RawQuery = rawQuery

	if !strings.Contains(request.Path, RSSHubUMSPath) {
		request.Path += RSSHubUMSPath
	}

	// fill in authentication config
	if err := h.parseRSSHubAuthentication(ctx, request); err != nil {
		return nil, fmt.Errorf("parse rsshub authentication: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, request.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create rsshub request: %w", err)
	}

	// nolint:bodyclose // False positive
	resp, err := h.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("rsshub request: %w", err)
	}

	defer lo.Try(resp.Body.Close)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("rsshub response status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("rsshub response body: %w", err)
	}

	return data, nil
}

// parseRSSHubAuthentication parse authentication config and validate it, then fill in request
func (h *Hub) parseRSSHubAuthentication(_ context.Context, request *url.URL) error {
	if h.rsshub.Parameters == nil {
		return nil
	}

	option, err := NewOption(h.rsshub.Parameters)
	if err != nil {
		return fmt.Errorf("parse parmeters: %w", err)
	}

	if option.Authentication.Username != "" && option.Authentication.Password != "" {
		request.User = url.UserPassword(option.Authentication.Username, option.Authentication.Password)
	}

	if option.Authentication.AccessKey != "" && option.Authentication.AccessCode != "" {
		request.RawQuery = fmt.Sprintf("%s&%s=%s", request.RawQuery, option.Authentication.AccessKey, option.Authentication.AccessCode)
	}

	return nil
}
