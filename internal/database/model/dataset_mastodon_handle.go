package model

import "time"

type MastodonHandle struct {
	Handle    string    `json:"handle"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PaginatedMastodonHandles struct {
	Handles    []string
	TotalCount int64
	NextCursor string
}

type QueryMastodonHandles struct {
	Since  *uint64
	Limit  *int
	Cursor *string
}
