package engine

import (
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

type Module struct {
	RSS           []*Config `mapstructure:"rss"`
	Federated     []*Config `mapstructure:"federated"`
	Decentralized []*Config `mapstructure:"decentralized"`
}

type Config struct {
	Network  filter.Network `mapstructure:"network" validate:"required"`
	Chain    string         `mapstructure:"chain"`
	Endpoint string         `mapstructure:"endpoint" validate:"required"`
	Worker   Name           `mapstructure:"worker"`
}
