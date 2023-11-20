package client

import (
	"context"
	"io"

	"github.com/samber/lo"
)

var DefaultHTTPClient = lo.Must(NewHTTPClient())

type HTTPClient interface {
	Fetch(ctx context.Context, id string) (io.ReadCloser, error)
}
