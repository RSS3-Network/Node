package atproto

type State struct {
	SubscribeCursor    int64  `json:"subscribe_cursor,omitempty"`
	SubscribeTimestamp int64  `json:"subscribe_timestamp,omitempty"`
	ListReposCursor    string `json:"list_repos_cursor,omitempty"`
}
