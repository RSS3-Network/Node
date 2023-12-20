package arweave

import (
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
)

var _ engine.SourceFilter = (*Filter)(nil)

type Filter struct {
	OwnerAddress []string `yaml:"owner_addresses"`
}
