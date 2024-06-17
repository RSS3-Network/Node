package component

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Component interface {
	Name() string
	InitMeter() error
	CollectMetric(ctx context.Context, path, value string)
	CollectTrace(ctx context.Context, path, value string)
}

var _ echo.Validator = (*Validator)(nil)

type Validator struct {
	Validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}
