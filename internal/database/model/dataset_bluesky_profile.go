package model

import "time"

type BlueskyProfile struct {
	DID       string    `json:"did"`
	Handle    string    `json:"handle"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type QueryBlueskyProfiles struct {
	Since  *uint64
	Limit  *int
	Cursor *string
}
