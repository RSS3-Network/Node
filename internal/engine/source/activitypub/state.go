package activitypub

// State represents the state of the ActivityPub data source.
type State struct {
	LastOffset int64 `json:"last_offset"`
}
