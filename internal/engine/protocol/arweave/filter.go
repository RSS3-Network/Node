package arweave

import (
	"github.com/rss3-network/node/internal/engine"
)

var _ engine.DataSourceFilter = (*Filter)(nil)

type Filter struct {
	OwnerAddresses []string `yaml:"owner_addresses"`

	// BundlrOnly is a tag, indicating that only transactions of Bundlr nodes are pulled.
	BundlrOnly bool `yaml:"bundlr_only"`

	// BundlrAddresses is a list of addresses that are considered Bundlr nodes.
	// If BundlrOnly is true, the field needs to be ignored.
	BundlrAddresses []string `yaml:"bundlr_addresses"`
}
