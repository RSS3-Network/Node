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

func (r *RSS) getRSSHubData(ctx context.Context, path string, rawQuery string) ([]byte, error) {
	request, err := url.Parse(r.rsshub.Endpoint)
	if err != nil {
		return nil, err
	}

	request.Path = path
	request.RawQuery = rawQuery

	if !strings.Contains(request.Path, RSSHubUMSPath) {
		request.Path += RSSHubUMSPath
	}

	// fill in authentication config
	if err := r.parseRSSHubAuthentication(ctx, request); err != nil {
		return nil, fmt.Errorf("parse rsshub authentication: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, request.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create rsshub request: %w", err)
	}

	// nolint:bodyclose // False positive
	resp, err := r.httpClient.Do(req)
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
func (r *RSS) parseRSSHubAuthentication(_ context.Context, request *url.URL) error {
	if r.rsshub.Parameters == nil {
		return nil
	}

	authentication, ok := r.rsshub.Parameters["authentication"].(map[string]interface{})
	if !ok {
		return nil
	}

	// validate authentication parameters
	for key, value := range authentication {
		if value == nil {
			continue
		}

		if _, ok := value.(string); !ok {
			return fmt.Errorf("invalid authentication config, key: %s, value: %v", key, value)
		}
	}

	username, password := authentication["username"], authentication["password"]
	if username != nil && password != nil {
		request.User = url.UserPassword(username.(string), password.(string))
	}

	accessKey, accessCode := authentication["access_key"], authentication["access_code"]
	if accessKey != nil && accessCode != nil {
		request.RawQuery = fmt.Sprintf("%s&%s=%s", request.RawQuery, accessKey, accessCode)
	}

	return nil
}
