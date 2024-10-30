package middlewarex

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

// DecodePathParamsMiddleware decodes path parameters.
func DecodePathParamsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Decode path parameters
		paramValues := c.ParamValues()
		for i, value := range paramValues {
			decodedValue, err := url.PathUnescape(value)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid path parameter: "+err.Error())
			}

			paramValues[i] = decodedValue
		}

		return next(c)
	}
}

// HeadToGetMiddleware Add HEAD to all GET methods.
func HeadToGetMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == http.MethodHead {
			// Set the method to GET temporarily to reuse the handler
			c.Request().Method = http.MethodGet

			defer func() {
				c.Request().Method = http.MethodHead
			}() // Restore method after

			// Call the next handler and then clear the response body
			if err := next(c); err != nil {
				return err
			}

			return c.NoContent(http.StatusOK)
		}

		return next(c)
	}
}
