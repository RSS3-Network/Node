package engine

import (
	"fmt"
	"strings"

	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

type Config struct {
	Name     string         `mapstructure:"name" validate:"required"`
	Endpoint string         `mapstructure:"endpoint" validate:"required"`
	Worker   Name           `mapstructure:"-"`
	Network  filter.Network `mapstructure:"-"`
	Chain    filter.Chain   `mapstructure:"-"`
}

// Parse parses the config name.
func (c *Config) Parse() (err error) {
	data := strings.Split(c.Name, ":")
	if len(data) < 1 {
		return fmt.Errorf("invalid name %s", c.Name)
	}

	engineConfigs := strings.Split(data[0], ".")
	if len(engineConfigs) < 1 {
		return fmt.Errorf("invalid name %s", c.Name)
	}

	if c.Network, err = filter.NetworkString(engineConfigs[0]); err != nil {
		return fmt.Errorf("invalid network: %w", err)
	}

	switch c.Network {
	case filter.NetworkRSSHub:
	case filter.NetworkFarcaster:
	default:
		if len(engineConfigs) < 3 {
			return fmt.Errorf("invalid name %s", c.Name)
		}

		if c.Chain, err = filter.ChainString(c.Network, engineConfigs[1]); err != nil {
			return fmt.Errorf("invalid chain: %w", err)
		}

		if c.Worker, err = NameString(engineConfigs[2]); err != nil {
			return fmt.Errorf("invalid worker: %w", err)
		}
	}

	return nil
}
