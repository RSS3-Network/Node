package engine

import (
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

type Config struct {
	Name     string         `mapstructure:"name" validate:"required"`
	Type     string         `mapstructure:"type" validate:"required"`
	Endpoint string         `mapstructure:"endpoint" validate:"required"`
	Worker   Name           `mapstructure:"worker"`
	Network  filter.Network `mapstructure:"network"`
	Chain    filter.Chain   `mapstructure:"chain"`
}
