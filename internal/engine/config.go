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

type Config struct {
	Network    filter.Network `yaml:"network" validate:"required"`
	Endpoint   string         `yaml:"endpoint" validate:"required"`
	Worker     Name           `yaml:"worker"`
	Parameters *Options       `yaml:"parameters"`
}

type Options struct {
	*yaml.Node
}

func (o *Options) UnmarshalYAML(node *yaml.Node) error {
	o.Node = node

	return nil
}
