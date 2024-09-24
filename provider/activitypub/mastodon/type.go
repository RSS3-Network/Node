package mastodon

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=MessageType --output type_message.go --trimprefix=MessageType
type MessageType int

const (
	MessageTypeNone     MessageType = iota // Invalid default value
	MessageTypeCreate                      // Create ActivityPub message
	MessageTypeAnnounce                    // Announce ActivityPub message
	MessageTypeLike                        // Like ActivityPub message
)

const (
	// ActivityPub object types
	ObjectTypeNote = "Note"

	InReplyTo = "inReplyTo"
	To        = "to"
	Type      = "type"

	// ActivityPub tag types
	TagTypeMention = "Mention"
	TagTypeHashtag = "Hashtag"

	KafkaTopic = "kafka_topic"
)
