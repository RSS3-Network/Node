package mastodon

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=MessageType --output type_message.go --trimprefix=MessageType
type MessageType int

const (
	MessageTypeNone     MessageType = iota // Invalid default value
	MessageTypeCreate                      // Create ActivityPub message
	MessageTypeAnnounce                    // Announce ActivityPub message
	MessageTypeLike                        // Like ActivityPub message
)

// ActivityPub standard contexts and public addressing
const (
	ActivityStreamsContext       = "https://www.w3.org/ns/activitystreams"
	SecurityV1Context            = "https://w3id.org/security/v1"
	ActivityStreamsPublicContext = "https://www.w3.org/ns/activitystreams#Public"
)

// Default Configuration Values

// DefaultServerPort is the server port
const DefaultServerPort = 8181
const DefaultMonitorServerPort = 9191

// DefaultRelayURLList is a list of hardcoded relay URLs to follow
var DefaultRelayURLList = []string{
	"https://relay.toot.io/inbox",
	"https://relay.infosec.exchange/inbox",
	"https://relay.intahnet.co.uk/inbox",
	"https://relay.an.exchange/inbox",
}

// ActivityPub message properties and types
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
)

// HTTP paths and headers
const (
	ActivitySuffix    = "/activity"
	inBoxPath         = "/inbox"
	outBoxPath        = "/outbox"
	actorPath         = "/actor"
	headerHost        = "Host"
	headerDigest      = "Digest"
	headerDate        = "Date"
	headerContentType = "Content-Type"
	activityJSONType  = "application/activity+json"
)
