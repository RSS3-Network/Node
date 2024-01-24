package metadata

import (
	"github.com/rss3-network/serving-node/schema/filter"
)

var _ Metadata = (*RSS)(nil)

type RSS struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	PubDate     string    `json:"pub_date,omitempty"`
	Authors     []Authors `json:"authors,omitempty"`
}

type Authors struct {
	Name string `json:"name,omitempty"`
}

func (r RSS) Type() filter.Type {
	return filter.TypeRSSFeed
}
