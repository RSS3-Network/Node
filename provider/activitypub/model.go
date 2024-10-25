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

type Actor struct {
	Context   []string  `json:"@context"`
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Inbox     string    `json:"inbox"`
	Outbox    string    `json:"outbox"`
	PublicKey PublicKey `json:"publicKey"`
	Endpoints EndPoint  `json:"endpoints"`
}

type PublicKey struct {
	ID           string `json:"id"`
	Owner        string `json:"owner"`
	PublicKeyPem string `json:"publicKeyPem"`
}

type EndPoint struct {
	SharedInbox string `json:"sharedInbox"`
}

// NodeInfo represents the node information response structure
type NodeInfo struct {
	Links []NodeInfoLink `json:"links"`
}

// NodeInfoLink represents a link in NodeInfo.
type NodeInfoLink struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

// NodeInfoDetails represents detailed node information.
type NodeInfoDetails struct {
	Version           string        `json:"version"`
	Software          SoftwareInfo  `json:"software"`
	Protocols         []string      `json:"protocols"`
	Services          ServicesInfo  `json:"services"`
	OpenRegistrations bool          `json:"openRegistrations"`
	Usage             NodeUsageInfo `json:"usage"`
}

// SoftwareInfo provides software details in NodeInfo.
type SoftwareInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ServicesInfo represents the inbound and outbound services of a node.
type ServicesInfo struct {
	Inbound  []string `json:"inbound"`
	Outbound []string `json:"outbound"`
}

// NodeUsageInfo contains usage statistics for a node.
type NodeUsageInfo struct {
	Users UsersInfo `json:"users"`
}

// UsersInfo provides user statistics.
type UsersInfo struct {
	Total int `json:"total"`
}

// InstanceInfo represents instance-specific information response.
type InstanceInfo struct {
	URI         string `json:"uri"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
