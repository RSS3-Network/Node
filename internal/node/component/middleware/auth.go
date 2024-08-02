package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// BearerAuth middleware for bearer token authentication
func BearerAuth(accessToken string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
			}

			// Check if the header starts with "Bearer "
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
			}

			// Extract the token
			token := strings.TrimPrefix(authHeader, "Bearer ")

			// Verify the token
			if token != accessToken {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid access token")
			}

			return next(c)
		}
	}
}
