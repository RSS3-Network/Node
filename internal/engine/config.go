package engine

import (
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"gopkg.in/yaml.v3"
)

type Module struct {
	RSS           []*Config `yaml:"rss" validate:"dive"`
	Federated     []*Config `yaml:"federated" validate:"dive"`
	Decentralized []*Config `yaml:"decentralized" validate:"dive"`
}

var _ yaml.Unmarshaler = (*Config)(nil)

type Config struct {
	Chain      filter.Chain `yaml:"chain" validate:"required"`
	Endpoint   string       `yaml:"endpoint" validate:"required"`
	Worker     Name         `yaml:"worker"`
	Parameters *Options     `yaml:"parameters"`
}

func (c *Config) UnmarshalYAML(value *yaml.Node) error {
	type tmp struct {
		Network    filter.Network `yaml:"network"`
		Chain      string         `yaml:"chain"`
		Endpoint   string         `yaml:"endpoint"`
		Worker     Name           `yaml:"worker"`
		Parameters *Options       `yaml:"parameters"`
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
	c.Parameters = t.Parameters

	return nil
}

type Options struct {
	*yaml.Node
}

func (o *Options) UnmarshalYAML(node *yaml.Node) error {
	o.Node = node

	return nil
}
