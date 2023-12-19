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
	ErrorCodeValidateFailed
	ErrorCodeBadParams
	ErrorCodeInternalError
)

type ErrorResponse struct {
	Message string    `json:"error"`
	Code    ErrorCode `json:"error_code"`
}

func BadRequestError(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, &ErrorResponse{
		Code:    ErrorCodeBadRequest,
		Message: fmt.Sprintf("Please check your request and try again, %s", err),
	})
}

func ValidateFailedError(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, &ErrorResponse{
		Code:    ErrorCodeValidateFailed,
		Message: fmt.Sprintf("Please check your request validation and try again, %s", err),
	})
}

func BadParamsError(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, &ErrorResponse{
		Code:    ErrorCodeBadParams,
		Message: fmt.Sprintf("Please check your param combination and try again, %s", err),
	})
}

func InternalError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, &ErrorResponse{
		Code:    ErrorCodeInternalError,
		Message: fmt.Sprintf("An internal error has occurred, please try again later, %s", err),
	})
}
