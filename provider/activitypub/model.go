package activitypub

type MastodonQuery struct {
	Limit *int64 `form:"limit,omitempty"`
}

type MessageResponse struct {
}

// This struct represents a general ActivityPub object or activity.
type Object struct {
	Context    interface{}            `json:"@context,omitempty"`
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Actor      string                 `json:"actor,omitempty"`
	Object     interface{}            `json:"object,omitempty"`
	Target     interface{}            `json:"target,omitempty"`
	Result     interface{}            `json:"result,omitempty"`
	Origin     interface{}            `json:"origin,omitempty"`
	Instrument interface{}            `json:"instrument,omitempty"`
	Published  string                 `json:"published,omitempty"`
	To         []string               `json:"to,omitempty"`
	CC         []string               `json:"cc,omitempty"`
	Bto        []string               `json:"bto,omitempty"`
	Bcc        []string               `json:"bcc,omitempty"`
	Attachment []Attachment           `json:"attachment,omitempty"`
	Tag        []Tag                  `json:"tag,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// this struct represents an attachment to an ActivityPub object.
type Attachment struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// this struct represents a tag in an ActivityPub object.
type Tag struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// this struct represents a Note object in ActivityPub.
type Note struct {
	ID        string   `json:"id"`
	Type      string   `json:"type"`
	Content   string   `json:"content"`
	Published string   `json:"published,omitempty"`
	To        []string `json:"to,omitempty"`
	CC        []string `json:"cc,omitempty"`
}
