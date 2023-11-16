package ipfs

import (
	"context"
	"io"
)

type Client interface {
	Fetch(ctx context.Context, path string, model Mode) (io.ReadCloser, error)
}
