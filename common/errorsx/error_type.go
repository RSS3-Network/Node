package errorsx

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=ErrorType --transform=snake --trimprefix=ErrorType --output error_type_string.go --json --sql
type ErrorType int

const (
	ErrorTypeUnknown ErrorType = iota + 1000
	ErrorTypeInvalidTask
	ErrorTypeBuildActivityFailed
	ErrorTypeHandleTransactionFailed
	ErrorTypeHandleLogFailed
	ErrorTypeBuildActionFailed
	ErrorTypeParseEventFailed
	ErrorTypeUnsupportedEvent
	ErrorTypeInvalidNetwork
	ErrorTypeLookupTokenFailed
	ErrorTypeInitClientFailed
)
