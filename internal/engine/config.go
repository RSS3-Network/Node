package engine

import (
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

type Config struct {
	Name     string         `mapstructure:"name"`
	Type     string         `mapstructure:"type"`
	Endpoint string         `mapstructure:"endpoint"`
	Worker   Name           `mapstructure:"worker"`
	Network  filter.Network `mapstructure:"network"`
	Chain    filter.Chain   `mapstructure:"chain"`
}
