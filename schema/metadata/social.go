package metadata

import (
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

var _ schema.Metadata = (*SocialPost)(nil)

type SocialPost struct {
	Handle        string   `json:"handle,omitempty"`
	Title         string   `json:"title,omitempty"`
	Summary       string   `json:"summary,omitempty"`
	Body          string   `json:"body,omitempty"`
	Media         []Media  `json:"media,omitempty"`
	ProfileID     string   `json:"profile_id,omitempty"`
	PublicationID string   `json:"publication_id,omitempty"`
	ContentURI    string   `json:"content_uri,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	AuthorURL     string   `json:"author_url,omitempty"`
	Reward        *Token   `json:"reward,omitempty"`

	Target    *SocialPost `json:"target,omitempty"`
	TargetURL string      `json:"target_url,omitempty"`
}

type Media struct {
	Address  string `json:"address"`
	MimeType string `json:"mime_type"`
}

func (s SocialPost) Type() filter.Type {
	return filter.TypeSocialPost
}
