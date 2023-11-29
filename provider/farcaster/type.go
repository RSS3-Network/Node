package farcaster

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=MessageType --output type_message.go --transform=snake-upper
type MessageType int

const (
	MessageTypeNone                      MessageType = 0  // Invalid default value
	MessageTypeCastAdd                   MessageType = 1  // Add a new Cast
	MessageTypeCastRemove                MessageType = 2  // Remove an existing Cast
	MessageTypeReactionAdd               MessageType = 3  // Add a Reaction to a Cast
	MessageTypeReactionRemove            MessageType = 4  // Remove a Reaction from a Cast
	MessageTypeLinkAdd                   MessageType = 5  // Add a Link to a target
	MessageTypeLinkRemove                MessageType = 6  // Remove a Link from a target
	MessageTypeVerificationAddEthAddress MessageType = 7  // Add a Verification of an Ethereum Address
	MessageTypeVerificationRemove        MessageType = 8  // Remove a Verification
	MessageTypeUserDataAdd               MessageType = 11 // Add metadata about a user
	MessageTypeUsernameProof             MessageType = 12 // Add or replace a username proof
)

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=UserDataType --output type_user_data.go --transform=snake-upper
type UserDataType int

const (
	UserDataTypeNone     UserDataType = 0 // Invalid default value
	UserDataTypePfp      UserDataType = 1 // Profile Picture for the user
	UserDataTypeDisplay  UserDataType = 2 // Display Name for the user
	UserDataTypeBio      UserDataType = 3 // Bio for the user
	UserDataTypeURL      UserDataType = 5 // URL of the user
	UserDataTypeUsername UserDataType = 6 // Preferred Farcaster Name for the user
)

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=ReactionType --output type_reaction.go --transform=snake-upper
type ReactionType int

const (
	ReactionTypeNone   ReactionType = 0 // Invalid default value
	ReactionTypeLike   ReactionType = 1 // Like the target cast
	ReactionTypeRecast ReactionType = 2 // Share target cast to the user's audience
)

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=HubEventType --output type_event.go --transform=snake-upper
type HubEventType int

const (
	HubEventTypeNone               HubEventType = 0 // Invalid default value
	HubEventTypeMergeMessage       HubEventType = 1 // A message was merged into the Hub
	HubEventTypePruneMessage       HubEventType = 2 // A message was pruned because a limit was exceeded
	HubEventTypeRevokeMessage      HubEventType = 3 // A message was revoked by a user
	HubEventTypeMergeUsernameProof HubEventType = 6
	HubEventTypeMergeOnChainEvent  HubEventType = 9
)

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=UsernameProofType --output type_username_proof.go --transform=snake-upper
type UsernameProofType int

const (
	UsernameTypeNone  UsernameProofType = 0
	UsernameTypeFname UsernameProofType = 1
	UsernameTypeEnsL1 UsernameProofType = 2
)
