package mastodon

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=MessageType --output mastodon_message_type.go --transform=snake-upper
type MessageType int

const (
	MessageTypeNone     MessageType = iota // Invalid default value
	MessageTypeCreate                      // Create ActivityPub message
	MessageTypeAnnounce                    // Announce ActivityPub message
	MessageTypeLike                        // Like ActivityPub message
)
