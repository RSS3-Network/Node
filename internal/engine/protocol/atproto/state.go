package atproto

type State struct {
	Cursor    string `json:"cursor,omitempty"`
	Timestamp uint64 `json:"timestamp,omitempty"`
}
