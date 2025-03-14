package atproto

import (
	"fmt"
	"time"

	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/config/parameter"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type Option struct {
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`

	TimestampStart time.Time `json:"timestamp_start" mapstructure:"timestamp_start"`
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
			option.TimestampStart = time.Now().Add(-time.Hour * 24 * 90)

			return &option, nil
		}

		option.TimestampStart = time.Unix(parameter.CurrentNetworkStartBlock[network.Bluesky].Timestamp, 0)
	}

	return &option, nil
}
