package nearsocial

type FunctionCallArgs struct {
	Data map[string]UserContent `json:"data"`
}

type UserContent struct {
	Index map[string]string `json:"index,omitempty"`
	Post  map[string]string `json:"post,omitempty"`
}

type IndexData struct {
	Key   IndexKey   `json:"key"`
	Value IndexValue `json:"value"`
}

type IndexKey struct {
	Type        string `json:"type"`
	Path        string `json:"path,omitempty"`
	BlockHeight int64  `json:"blockHeight,omitempty"`
}

type IndexValue struct {
	Type string `json:"type"`
	Path string `json:"path,omitempty"`
	Item *Item  `json:"item,omitempty"`
}

type Item struct {
	Type        string `json:"type"`
	Path        string `json:"path"`
	BlockHeight int64  `json:"blockHeight,omitempty"`
}

type PostData struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
	Item *Item  `json:"item,omitempty"`
}

type HashtagData struct {
	Key   string       `json:"key"`
	Value HashtagValue `json:"value"`
}

type HashtagValue struct {
	Type string `json:"type"`
	Path string `json:"path"`
}
