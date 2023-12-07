package engine

import (
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"gopkg.in/yaml.v3"
)

type Module struct {
	RSS           []*Config `yaml:"rss"`
	Federated     []*Config `yaml:"federated"`
	Decentralized []*Config `yaml:"decentralized"`
}

type Config struct {
	Chain    filter.Chain `yaml:"chain"`
	Endpoint string       `yaml:"endpoint" validate:"required"`
	Worker   Name         `yaml:"worker"`
}

var _ yaml.Unmarshaler = (*Config)(nil)

func (c *Config) UnmarshalYAML(value *yaml.Node) error {
	type tmp struct {
		Network  filter.Network `yaml:"network" validate:"required"`
		Chain    string         `yaml:"chain"`
		Endpoint string         `yaml:"endpoint" validate:"required"`
		Worker   Name           `yaml:"worker"`
	}

	var t tmp

	err := value.Decode(&t)
	if err != nil {
		return err
	}

	c.Chain, err = filter.ChainString(t.Network, t.Chain)
	if err != nil {
		return err
	}

	c.Endpoint = t.Endpoint
	c.Worker = t.Worker

	return nil
}
