package mastodon

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=MessageType --output type_message.go --trimprefix=MessageType
type MessageType int

const (
	ActivityStreamsContext       = "https://www.w3.org/ns/activitystreams"
	SecurityV1Context            = "https://w3id.org/security/v1"
	ActivityStreamsPublicContext = "https://www.w3.org/ns/activitystreams#Public"
)

const (
	MessageTypeNone     MessageType = iota // Invalid default value
	MessageTypeCreate                      // Create ActivityPub message
	MessageTypeAnnounce                    // Announce ActivityPub message
	MessageTypeLike                        // Like ActivityPub message
)

const (
	InReplyTo           = "inReplyTo"
	TagType             = "type"
	TagName             = "name"
	Attachment          = "attachment"
	AttachmentURL       = "url"
	AttachmentMediaType = "mediaType"
	Tag                 = "tag"
	TagTypeHashtag      = "Hashtag"
	TagTypeMention      = "Mention"
	KafkaTopic          = "kafka_topic"
)
