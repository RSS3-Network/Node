package arweave

import (
	"github.com/rss3-network/node/internal/engine"
)

var _ engine.DataSourceFilter = (*Filter)(nil)

type Filter struct {
	OwnerAddresses  []string `yaml:"owner_addresses"`
	BundlrAddresses []string `yaml:"bundlr_addresses"`
}
