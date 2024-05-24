package middlewarex_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/internal/node/middlewarex"
	"github.com/stretchr/testify/assert"
)

func TestPathParameterDecodeMiddleware_DecodesValidPathParameters(t *testing.T) {
	t.Parallel()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/test/36kr%2Fnewsflashes+", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/test/:path")
	c.SetParamNames("path")
	c.SetParamValues("36kr%2Fnewsflashes+")

	handler := middlewarex.DecodePathParamsMiddleware(func(c echo.Context) error {
		assert.Equal(t, "36kr/newsflashes+", c.Param("path"))
		return nil
	})

	err := handler(c)
	assert.NoError(t, err)
}
