package crossbell

import (
	"encoding/json"
	"time"
)

type NoteContent struct {
	Type          string                  `json:"type"`
	Tags          []string                `json:"tags"`
	Authors       []string                `json:"authors"`
	Title         string                  `json:"title"`
	Content       string                  `json:"content"`
	Attachments   []NoteContentAttachment `json:"attachments"`
	Sources       []string                `json:"sources"`
	ExternalUrls  []string                `json:"external_urls"`
	DatePublished time.Time               `json:"date_published"`
}

type NoteContentAttachment struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	MimeType    string `json:"mime_type"`
	SizeInBytes int    `json:"size_in_bytes"`
	Alt         string `json:"alt"`
	Width       int    `json:"width"`
}

type CharacterURIContent struct {
	Avatars           []string        `json:"avatars"`
	Bio               string          `json:"bio"`
	ConnectedAccounts json.RawMessage `json:"connected_accounts"`
	Name              string          `json:"name"`
}

type ProfileURIContent struct {
	ConnectedAccounts json.RawMessage `json:"connected_accounts"`
	ConnectedAvatars  []string        `json:"connected_avatars"`
	Name              string          `json:"name"`
	Bio               string          `json:"bio"`
	Type              string          `json:"type"`
}
