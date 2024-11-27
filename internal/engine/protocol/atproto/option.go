package atproto

import (
	"fmt"

	"github.com/rss3-network/node/config"
	"go.uber.org/zap"
)

type Option struct {
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
}

func NewOption(parameters *config.Parameters) (*Option, error) {
	var option Option

	if parameters == nil {
		return &Option{}, nil
	}

	if err := parameters.Decode(&option); err != nil {
		zap.L().Error("decode parameters failed", zap.Error(err))

		return nil, fmt.Errorf("decode parameters failed: %w", err)
	}

	return &option, nil
}
