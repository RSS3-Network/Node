package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/v2/internal/engine"
)

var _ engine.DataSourceFilter = (*Filter)(nil)

type Filter struct {
	LogAddresses []common.Address `yaml:"log_addresses"`
	LogTopics    []common.Hash    `yaml:"log_topics"`
}
