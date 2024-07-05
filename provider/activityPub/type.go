package activitypub

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=ActivityType --output type_activity.go --transform=snake-upper
type ActivityType int

const (
	ActivityTypeNone     ActivityType = iota // Invalid default value
	ActivityTypeCreate                       // Create a new object
	ActivityTypeDelete                       // Delete an existing object
	ActivityTypeFollow                       // Follow a user
	ActivityTypeUnfollow                     // Unfollow a user
	ActivityTypeLike                         // Like an object
	ActivityTypeAnnounce                     // Announce (boost/share) an object
	ActivityTypeUpdate                       // Update an object
	ActivityTypeAdd                          // Add an object to a collection
	ActivityTypeRemove                       // Remove an object from a collection
	ActivityTypeBlock                        // Block a user
	ActivityTypeUndo                         // Undo a previous action
)

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=ObjectType --output type_object.go --transform=snake-upper
type ObjectType int

const (
	ObjectTypeNone       ObjectType = iota // Invalid default value
	ObjectTypeNote                         // A Note object
	ObjectTypeImage                        // An Image object
	ObjectTypeVideo                        // A Video object
	ObjectTypeAudio                        // An Audio object
	ObjectTypeEvent                        // An Event object
	ObjectTypePerson                       // A Person object
	ObjectTypeArticle                      // An Article object
	ObjectTypePlace                        // A Place object
	ObjectTypeDocument                     // A Document object
	ObjectTypeCollection                   // A Collection object
	ObjectTypeService                      // A Service object
)
