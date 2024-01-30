package arweave

import (
	"github.com/rss3-network/node/internal/engine"
)

var _ engine.SourceFilter = (*Filter)(nil)

type Filter struct {
	OwnerAddresses []string `yaml:"owner_addresses"`
}
