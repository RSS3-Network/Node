package schema

import "github.com/naturalselectionlabs/rss3-node/schema/filter"

type Action struct {
	Tag      filter.Tag  `json:"tag"`
	Type     filter.Type `json:"type"`
	Platform string      `json:"platform,omitempty"`
	From     string      `json:"from"`
	To       string      `json:"to"`
	Metadata Metadata    `json:"metadata"`
}
