package atproto

type State struct {
	SubscribeCursor    string `json:"subscribe_cursor,omitempty"`
	SubscribeTimestamp uint64 `json:"subscribe_timestamp,omitempty"`
	ListReposCursor    string `json:"list_repos_cursor,omitempty"`
}
