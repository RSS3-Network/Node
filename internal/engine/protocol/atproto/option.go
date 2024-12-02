package atproto

import (
	"fmt"
	"time"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type Option struct {
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`

	TimestampStart int64 `json:"timestamp_start" mapstructure:"timestamp_start"`
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

	if lo.IsEmpty(option.TimestampStart) {
		if parameter.CurrentNetworkStartBlock[network.Bluesky] == nil {
			// Default to 90 days ago
			option.TimestampStart = time.Now().Add(-time.Hour * 24 * 90).Unix()

			return &option, nil
		}

		option.TimestampStart = parameter.CurrentNetworkStartBlock[network.Bluesky].Timestamp
	}

	return &option, nil
}
