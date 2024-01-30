package hub

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/node/hub/decentralized"
	"github.com/rss3-network/node/internal/node/hub/rss"
)

type Hub struct {
	RSS           *rss.Hub
	Decentralized *decentralized.Hub
}

var _ echo.Validator = (*Validator)(nil)

type Validator struct {
	validate *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validate.Struct(i)
}

func NewHub(ctx context.Context, config *config.File, database database.Client) *Hub {
	return &Hub{
		RSS:           rss.NewHub(ctx, config),
		Decentralized: decentralized.NewHub(ctx, database),
	}
}
