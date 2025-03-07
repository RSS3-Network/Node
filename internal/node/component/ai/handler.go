package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	"go.uber.org/zap"
)

type Response struct {
	Data any `json:"data"`
}

func (h *Component) Handler(ctx echo.Context) error {
	zap.L().Debug("handling AgentData request",
		zap.String("request_uri", ctx.Request().RequestURI))

	addRecentRequest(ctx.Request().RequestURI)

	reqURL, err := h.formatRequestURL(ctx.Request().Context(), ctx.Request().URL)
	if err != nil {
		zap.L().Error("failed to format request", zap.Error(err), zap.String("uri", reqURL.String()))

		return response.BadRequestError(ctx, err)
	}

	request, err := http.NewRequestWithContext(ctx.Request().Context(), http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return response.BadRequestError(ctx, err)
	}

	resp, err := h.httpClient.Do(request)
	if err != nil {
		zap.L().Error("failed to get response from ai", zap.Error(err), zap.String("uri", reqURL.String()))

		return response.InternalError(ctx)
	}

	if resp.StatusCode != http.StatusOK {
		zap.L().Error("failed to get response from ai", zap.Int("status_code", resp.StatusCode), zap.String("uri", reqURL.String()))

		return response.InternalError(ctx)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	var result any

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		zap.L().Error("failed to decode response from ai", zap.Error(err))

		return response.InternalError(ctx)
	}

	return ctx.JSON(http.StatusOK, result)
}

// formatRequestURL prepares the URL for the ai request with necessary parameters.
func (h *Component) formatRequestURL(_ context.Context, in *url.URL) (*url.URL, error) {
	base, err := url.Parse(h.config.Component.AI.Endpoint.URL)
	if err != nil {
		return nil, fmt.Errorf("parse ai endpoint: %w", err)
	}

	trimmedPath := strings.TrimPrefix(in.Path, fmt.Sprintf("/%s/", Name))
	base.Path = path.Join(base.Path, trimmedPath)
	base.RawQuery = in.Query().Encode()

	return base, nil
}
