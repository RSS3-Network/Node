package model

import "time"

type MastodonHandleTransformer interface {
	Import(handle *MastodonHandle) error
	Export() (*MastodonHandle, error)
}

type MastodonHandle struct {
	Handle      string    `json:"handle"`
	LastUpdated time.Time `json:"last_updated"`
}
