package model

type MastodonHandleTransformer interface {
	Import(handle *MastodonHandle) error
	Export() (*MastodonHandle, error)
}

type MastodonHandle struct {
	Handle      string `json:"handle"`
	LastUpdated uint64 `json:"last_updated"`
}
