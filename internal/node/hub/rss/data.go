package rss

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/rss3-network/protocol-go/schema"
	"github.com/samber/lo"
)

const (
	RSSHubUMSPath = "format=ums"
)

func (h *Hub) getRSSHubData(ctx context.Context, path string, rawQuery string) ([]*schema.Feed, error) {
	request, err := url.Parse(h.rsshub.Endpoint)
	if err != nil {
		return nil, err
	}

	request.Path = path
	request.RawQuery = rawQuery

	if !strings.Contains(request.RawQuery, RSSHubUMSPath) {
		if request.RawQuery != "" {
			request.RawQuery += "&" + RSSHubUMSPath
		} else {
			request.RawQuery = RSSHubUMSPath
		}
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

	activities, err := h.transformRSSHubToActivities(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("transform rsshub response: %w", err)
	}

	return activities, nil
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

func (h *Hub) transformRSSHubToActivities(_ context.Context, data []byte) ([]*schema.Feed, error) {
	type response struct {
		Data schema.Feeds `json:"data"`
	}

	var resp response

	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal rsshub response: %w", err)
	}

	return resp.Data, nil
}
