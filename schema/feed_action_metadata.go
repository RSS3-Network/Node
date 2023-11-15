package schema

import "github.com/naturalselectionlabs/rss3-node/schema/filter"

type Metadata interface {
	Type() filter.Type
}
