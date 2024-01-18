package httpx

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/samber/lo"
)

const (
	DefaultTimeout    = 3 * time.Second
	DefaultRetryCount = 1
)

var DefaultClient = lo.Must(NewClient())

type Client interface {
	Get(ctx context.Context, url string, headers http.Header, params url.Values) (io.ReadCloser, error)
	GetContentType(ctx context.Context, path string) (string, error)
}
