package handler

import (
	"context"

	"github.com/naturalselectionlabs/rss3-node/internal/config"
	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/node/explorer/handler/rss"
)

type Handler struct {
	RSSHandler *rss.Handler
}

func NewHandler(ctx context.Context, config *config.File, _ database.Client) *Handler {
	return &Handler{
		RSSHandler: rss.NewHandler(ctx, config),
	}
}
