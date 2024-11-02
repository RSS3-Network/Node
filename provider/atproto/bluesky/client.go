package bluesky

import (
	"context"
	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/xrpc"
	"go.uber.org/zap"
	"net/http"
)

const (
	BskyEndpoint = "https://bsky.social"
)

type Client struct {
	xrpcClient *xrpc.Client
}

func (c *Client) GetEvents(ctx context.Context) {

}

func NewClient(ctx context.Context, username string, password string) (*Client, error) {
	xrpcClient := &xrpc.Client{
		Host:   BskyEndpoint,
		Client: http.DefaultClient,
	}

	auth, err := atproto.ServerCreateSession(ctx, xrpcClient, &atproto.ServerCreateSession_Input{
		Identifier: username,
		Password:   password,
	})
	if err != nil {
		zap.L().Error("create session failed", zap.Error(err), zap.String("username", username))

		return nil, err
	}

	xrpcClient.Auth = &xrpc.AuthInfo{
		AccessJwt:  auth.AccessJwt,
		RefreshJwt: auth.RefreshJwt,
		Handle:     auth.Handle,
		Did:        auth.Did,
	}

	return &Client{
		xrpcClient: xrpcClient,
	}, nil
}
