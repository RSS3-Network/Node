package response

import (
	"fmt"

	"github.com/NaturalSelectionLabs/goapi"
	"github.com/NaturalSelectionLabs/goapi/lib/openapi"
)

//go:generate go run --mod=mod github.com/dmarkham/enumer --type=ErrorCode --transform=snake --values --trimprefix=ErrorCode --json --output error_code.go
type ErrorCode int

const (
	ErrorCodeBadRequest ErrorCode = iota + 1
	ErrorCodeValidateFailed
	ErrorCodeBadParams
	ErrorCodeAddressIsEmpty
	ErrorCodeAddressIsInvalid
	ErrorCodeInternalError
	ErrorCodeNotFound
)

type Error openapi.CommonError[ErrorCode]

func (e Error) Error() string {
	return e.Message
}

type BadRequest struct {
	goapi.StatusBadRequest
	Error Error
}

type InternalErrorResp struct {
	goapi.StatusInternalServerError
	Error Error
}

type NotFound struct {
	goapi.StatusNotFound
	Error Error
}

func BadRequestError() Error {
	return Error{
		Code:    ErrorCodeBadRequest,
		Message: "Please check your request and try again.",
	}
}

func ValidateFailedError(param string) Error {
	return Error{
		Code:    ErrorCodeValidateFailed,
		Message: fmt.Sprintf("Please check your request validation and try again, parameter: %s is invalid", param),
	}
}

func BadParamsError(tag []string, typeX []string) Error {
	return Error{
		Code:    ErrorCodeBadParams,
		Message: fmt.Sprintf("Please check your param combination and try again. Tag: %s. Type: %s", tag, typeX),
	}
}

func AddressIsEmptyError() Error {
	return Error{
		Code:    ErrorCodeAddressIsEmpty,
		Message: "At least one address is required",
	}
}

func ResolveDomainError() Error {
	return Error{
		Code:    ErrorCodeAddressIsInvalid,
		Message: "An error occurred while resolving the domain.",
	}
}

func InternalError() Error {
	return Error{
		Code:    ErrorCodeInternalError,
		Message: "An internal error has occurred, please try again later.",
	}
}

func NotFoundError(id string) Error {
	return Error{
		Code:    ErrorCodeNotFound,
		Message: fmt.Sprintf("Not found id: %s", id),
	}
}

func BadRequestErrorResponse() BadRequest {
	return BadRequest{Error: BadRequestError()}
}

func ValidateFailedErrorResponse(param string) BadRequest {
	return BadRequest{Error: ValidateFailedError(param)}
}

func BadParamsErrorResponse(tag []string, typeX []string) BadRequest {
	return BadRequest{Error: BadParamsError(tag, typeX)}
}

func ResolveDomainErrorResponse() InternalErrorResp {
	return InternalErrorResp{Error: ResolveDomainError()}
}

func InternalErrorResponse() InternalErrorResp {
	return InternalErrorResp{Error: InternalError()}
}

func NotFoundErrorResponse(id string) NotFound {
	return NotFound{Error: NotFoundError(id)}
}
