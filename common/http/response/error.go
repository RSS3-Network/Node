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
	Error     string    `json:"error"`
	ErrorCode ErrorCode `json:"error_code"`
}

func BadRequestError(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, &ErrorResponse{
		ErrorCode: ErrorCodeBadRequest,
		Error:     fmt.Sprintf("Please check your request and try again, %s", err),
	})
}

func ValidateFailedError(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, &ErrorResponse{
		ErrorCode: ErrorCodeValidateFailed,
		Error:     fmt.Sprintf("Please check your request validation and try again, %s", err),
	})
}

func BadParamsError(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, &ErrorResponse{
		ErrorCode: ErrorCodeBadParams,
		Error:     fmt.Sprintf("Please check your param combination and try again, %s", err),
	})
}

func InternalError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, &ErrorResponse{
		ErrorCode: ErrorCodeInternalError,
		Error:     fmt.Sprintf("An internal error has occurred, please try again later, %s", err),
	})
}
