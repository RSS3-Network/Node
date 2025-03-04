package agentdata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	"go.uber.org/zap"
)

type Response struct {
	Data any `json:"data"`
}

func (h *Component) Handler(ctx echo.Context) error {
	path := ctx.Param("*")
	zap.L().Debug("handling AgentData request",
		zap.String("path", path),
		zap.String("request_uri", ctx.Request().RequestURI))

	go h.CollectTrace(ctx.Request().Context(), ctx.Request().RequestURI, path)

	go h.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, path)

	addRecentRequest(ctx.Request().RequestURI)

	reqURL, err := h.formatRequestURL(path, ctx.Request().URL)
	if err != nil {
		zap.L().Error("failed to format request", zap.Error(err), zap.String("path", path))

		return response.BadRequestError(ctx, err)
	}

	request, err := http.NewRequestWithContext(ctx.Request().Context(), http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return response.BadRequestError(ctx, err)
	}

	resp, err := h.httpClient.Do(request)
	if err != nil {
		zap.L().Error("failed to get response from agentdata", zap.Error(err), zap.String("path", path))

		return response.InternalError(ctx)
	}

	if resp.StatusCode != http.StatusOK {
		zap.L().Error("failed to get response from agentdata", zap.Int("status_code", resp.StatusCode), zap.String("path", path))

		return response.InternalError(ctx)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	var result any

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		zap.L().Error("failed to decode response from agentdata", zap.Error(err))

		return response.InternalError(ctx)
	}

	return ctx.JSON(http.StatusOK, result)
}

// formatRequestURL prepares the URL for the agentdata request with necessary parameters.
func (h *Component) formatRequestURL(path string, in *url.URL) (*url.URL, error) {
	out, err := url.Parse(h.config.Component.AI.Endpoint.URL)
	if err != nil {
		return nil, fmt.Errorf("parse agentdata endpoint: %w", err)
	}

	parameters := in.Query()

	out.RawQuery = parameters.Encode()
	out.Path = path

	return out, nil
}
