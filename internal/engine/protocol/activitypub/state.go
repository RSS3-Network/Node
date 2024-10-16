package activitypub

// State represents the state of the ActivityPub data protocol.
type State struct {
	LastOffset int64 `json:"last_offset"`
}
