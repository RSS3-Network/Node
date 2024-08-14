package errorsx

import (
	"github.com/samber/lo"
)

type Error struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

type ErrorOption func(e *Error) error

func WithErr(err error) ErrorOption {
	return func(e *Error) error {
		e.Message = err.Error()

		return nil
	}
}

func InvalidTask(options ...ErrorOption) (e *Error) {
	defer lo.ForEach(options, func(option ErrorOption, _ int) {
		_ = option(e)
	})

	return &Error{
		Type:    ErrorTypeInvalidTask,
		Message: ErrorTypeInvalidTask.String(),
	}
}

func BuildActivityFailed(options ...ErrorOption) (e *Error) {
	defer lo.ForEach(options, func(option ErrorOption, _ int) {
		_ = option(e)
	})

	return &Error{
		Type:    ErrorTypeBuildActivityFailed,
		Message: ErrorTypeBuildActivityFailed.String(),
	}
}

func HandleTransactionFailed(options ...ErrorOption) (e *Error) {
	defer lo.ForEach(options, func(option ErrorOption, _ int) {
		_ = option(e)
	})

	return &Error{
		Type:    ErrorTypeHandleTransactionFailed,
		Message: ErrorTypeHandleTransactionFailed.String(),
	}
}

func HandleLogFailed(options ...ErrorOption) (e *Error) {
	defer lo.ForEach(options, func(option ErrorOption, _ int) {
		_ = option(e)
	})

	return &Error{
		Type:    ErrorTypeHandleLogFailed,
		Message: ErrorTypeHandleLogFailed.String(),
	}
}

func BuildActionFailed(options ...ErrorOption) (e *Error) {
	defer lo.ForEach(options, func(option ErrorOption, _ int) {
		_ = option(e)
	})

	return &Error{
		Type:    ErrorTypeBuildActionFailed,
		Message: ErrorTypeBuildActionFailed.String(),
	}
}

func ParseEventFailed(options ...ErrorOption) (e *Error) {
	defer lo.ForEach(options, func(option ErrorOption, _ int) {
		_ = option(e)
	})

	return &Error{
		Type:    ErrorTypeParseEventFailed,
		Message: ErrorTypeParseEventFailed.String(),
	}
}

func UnsupportedEvent(options ...ErrorOption) (e *Error) {
	defer lo.ForEach(options, func(option ErrorOption, _ int) {
		_ = option(e)
	})

	return &Error{
		Type:    ErrorTypeUnsupportedEvent,
		Message: ErrorTypeUnsupportedEvent.String(),
	}
}

func InvalidNetwork(options ...ErrorOption) (e *Error) {
	defer lo.ForEach(options, func(option ErrorOption, _ int) {
		_ = option(e)
	})

	return &Error{
		Type:    ErrorTypeInvalidNetwork,
		Message: ErrorTypeInvalidNetwork.String(),
	}
}

func LookupTokenFailed(options ...ErrorOption) (e *Error) {
	defer lo.ForEach(options, func(option ErrorOption, _ int) {
		_ = option(e)
	})

	return &Error{
		Type:    ErrorTypeLookupTokenFailed,
		Message: ErrorTypeLookupTokenFailed.String(),
	}
}

func InitClientFailed(options ...ErrorOption) (e *Error) {
	defer lo.ForEach(options, func(option ErrorOption, _ int) {
		_ = option(e)
	})

	return &Error{
		Type:    ErrorTypeInitClientFailed,
		Message: ErrorTypeInitClientFailed.String(),
	}
}
