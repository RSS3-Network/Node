package arweave

type State struct {
	BlockHeight uint64 `json:"block_height,omitempty"`
	Cursor      string `json:"cursor,omitempty"`
}
