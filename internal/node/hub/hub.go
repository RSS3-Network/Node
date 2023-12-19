package hub

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"sync"

	"github.com/naturalselectionlabs/rss3-node/internal/config"
	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/node/hub/decentralized"
	"github.com/naturalselectionlabs/rss3-node/internal/node/hub/rss"
)

type Hub struct {
	RSS           *rss.Hub
	Decentralized *decentralized.Hub
}

var _ echo.Validator = &Validator{}

type Validator struct {
	validate *validator.Validate
	locker   sync.Mutex
}

func (v *Validator) Validate(i interface{}) error {
	if v.validate == nil {
		v.locker.Lock()
		defer v.locker.Unlock()

		if v.validate == nil {
			v.validate = validator.New()
		}
	}

	return v.validate.Struct(i)
}

func NewHub(ctx context.Context, config *config.File, database database.Client) *Hub {
	return &Hub{
		RSS:           rss.NewHub(ctx, config),
		Decentralized: decentralized.NewHub(ctx, database),
	}
}
