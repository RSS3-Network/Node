package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/serving-node/internal/engine"
)

var _ engine.SourceFilter = (*Filter)(nil)

type Filter struct {
	LogAddresses []common.Address `yaml:"log_addresses"`
	LogTopics    []common.Hash    `yaml:"log_topics"`
}
