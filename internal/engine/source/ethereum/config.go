package ethereum

import (
	"math/big"

	"gopkg.in/yaml.v3"
)

type Config struct {
	BlockNumberStart  *big.Int `yaml:"block_number_start"`
	BlockNumberTarget *big.Int `yaml:"block_number_target"`
}

func NewConfig(values map[string]any) (*Config, error) {
	data, err := yaml.Marshal(values)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
