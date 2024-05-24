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
