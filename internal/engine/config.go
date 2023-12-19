package engine

import (
	"encoding/json"
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/samber/lo"
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

// ID returns the unique identifier of the configuration.
func (c *Config) ID() string {
	id := fmt.Sprintf("%s.%s", c.Network, c.Worker)

	if c.Parameters != nil {
		var buffer map[string]any

		lo.Must0(c.Parameters.Decode(&buffer))

		if buffer != nil {
			id = fmt.Sprintf("%s.%s", id, string(lo.Must(json.Marshal(buffer))))
		}
	}

	return id
}

type Options struct {
	*yaml.Node
}

func (o *Options) UnmarshalYAML(node *yaml.Node) error {
	o.Node = node

	return nil
}
