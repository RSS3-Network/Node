package response

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:generate go run --mod=mod github.com/dmarkham/enumer --type=ErrorCode --transform=snake --values --trimprefix=ErrorCode --json --output error_code.go
type ErrorCode int

const (
	ErrorCodeBadRequest ErrorCode = iota + 1
	ErrorCodeValidationFailed
	ErrorCodeBadParams
	ErrorCodeInternalError
)

type ErrorResponse struct {
	Message string    `json:"message"`
	Code    ErrorCode `json:"code"`
}

func BadRequestError(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, &ErrorResponse{
		Code:    ErrorCodeBadRequest,
		Message: fmt.Sprintf("Please check your request and try again, %s", err),
	})
}

func ValidationFailedError(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, &ErrorResponse{
		Code:    ErrorCodeValidationFailed,
		Message: fmt.Sprintf("Please check your request validation and try again, %s", err),
	})
}

// InternalError should not return the details of the error to the client for safety reasons.
func InternalError(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, &ErrorResponse{
		Code:    ErrorCodeInternalError,
		Message: "An internal error has occurred, please try again later",
	})
}
