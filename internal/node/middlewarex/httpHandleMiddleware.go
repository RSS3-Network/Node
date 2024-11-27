package middlewarex

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// DecodePathParamsMiddleware decodes path parameters.
func DecodePathParamsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Decode path parameters
		paramValues := c.ParamValues()
		zap.L().Debug("decoding path parameters",
			zap.Strings("param_values", paramValues))

		for i, value := range paramValues {
			decodedValue, err := url.PathUnescape(value)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid path parameter: "+err.Error())
			}

			paramValues[i] = decodedValue
		}

		zap.L().Debug("successfully decoded path parameters",
			zap.Strings("decoded_values", paramValues))

		return next(c)
	}
}

// HeadToGetMiddleware Add HEAD to all GET methods.
func HeadToGetMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == http.MethodHead {
			zap.L().Debug("converting head request to get",
				zap.String("path", c.Request().URL.Path))

			// Set the method to GET temporarily to reuse the handler
			c.Request().Method = http.MethodGet

			defer func() {
				c.Request().Method = http.MethodHead
				zap.L().Debug("restored request method back to head")
			}() // Restore method after

			// Call the next handler and then clear the response body
			if err := next(c); err != nil {
				if err.Error() == echo.ErrMethodNotAllowed.Error() {
					zap.L().Debug("method not allowed, returning empty response for head request")
					c.NoContent(http.StatusOK) //nolint:errcheck

					return nil
				}

				zap.L().Error("error handling head request",
					zap.Error(err))

				return err
			}
		}

		return next(c)
	}
}
