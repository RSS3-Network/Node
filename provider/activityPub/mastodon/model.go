package mastodon

type mastodonQuery struct {
	Limit *int64 `form:"limit,omitempty"`
}

type MessageResponse struct {
}
