package near

import (
	"github.com/rss3-network/node/internal/engine"
)

var _ engine.DataSourceFilter = (*Filter)(nil)

type Filter struct {
	ReceiverIDs []string `yaml:"receiver_ids"`
	RelayerIDs  []string `yaml:"relayer_ids"`
}
