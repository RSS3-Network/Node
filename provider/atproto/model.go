package atproto

import (
	"time"

	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/atproto/syntax"
)

type Message struct {
	URI        string
	Did        syntax.DID
	Handle     string
	Collection string
	Rkey       string
	CreatedAt  time.Time

	Feed       *bsky.FeedPost
	Profile    *bsky.ActorProfile
	RefMessage *Message
}
