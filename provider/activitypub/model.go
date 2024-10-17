package activitypub

import "time"

// Object represents a general ActivityPub object or activity.
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
	Attachment interface{}            `json:"attachment,omitempty"`
	Tag        interface{}            `json:"tag,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// Attachment represents an attachment to an ActivityPub object.
type Attachment struct {
	Type      string `json:"type"`
	URL       string `json:"url"`
	MediaType string `json:"mediaType"`
}

// Tag represents a tag in an ActivityPub object.
type Tag struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Href string `json:"href"`
}

// Note represents a note object in ActivityPub.
type Note struct {
	ID        string   `json:"id"`
	Type      string   `json:"type"`
	Content   string   `json:"content"`
	Published string   `json:"published,omitempty"`
	To        []string `json:"to,omitempty"`
	CC        []string `json:"cc,omitempty"`
	Tag       []Tag    `json:"tag,omitempty"`
}

// StatusResult represents the result of a status request.
type StatusResult struct {
	Content     string
	Timestamp   string
	Attachments []Attachment
	Tags        []Tag
}

// StatusResponse represents a status response from an ActivityPub server.
type StatusResponse struct {
	Object struct {
		Published  time.Time    `json:"published"`
		Content    string       `json:"content"`
		Attachment []Attachment `json:"attachment"`
		Tag        []Tag        `json:"tag"`
	} `json:"object"`
}
