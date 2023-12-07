package rss

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"io"
	"net/http"
	"net/url"
)

// GetRSSHub get rsshub data from rsshub node
func (h *Handler) GetRSSHub(c echo.Context) error {
	request, err := url.Parse(h.rsshub.Endpoint)
	if err != nil {
		return err
	}

	request.Path = c.Param("*")
	request.RawQuery = c.Request().URL.RawQuery

	h.parseRSSHubAuthentication(c.Request().Context(), request)

	req, err := http.NewRequestWithContext(c.Request().Context(), http.MethodGet, request.String(), nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// nolint:bodyclose // False positive
	resp, err := h.httpClient.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	defer lo.Try(resp.Body.Close)

	if resp.StatusCode != http.StatusOK {
		return c.JSON(http.StatusInternalServerError, resp.Body)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSONBlob(http.StatusOK, data)
}

func (h *Handler) parseRSSHubAuthentication(_ context.Context, request *url.URL) {
	if h.rsshub.Parameters == nil {
		return
	}

	authentication, ok := h.rsshub.Parameters["authentication"].(map[string]interface{})
	if !ok {
		return
	}

	if username, password := authentication["username"], authentication["password"]; username != nil && password != nil {
		request.User = url.UserPassword(username.(string), password.(string))
	}

	if accessKey, accessCode := authentication["access_key"], authentication["access_code"]; accessKey != nil && accessCode != nil {
		request.RawQuery = fmt.Sprintf("%s&%s=%s", request.RawQuery, accessKey, accessCode)
	}
}
