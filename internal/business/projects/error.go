package projects

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type ProjectError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ProjectError) Error() string {
	return fmt.Sprintf("ProjectError [Code: %d] %s]", e.Code, e.Message)
}

func (e *ProjectError) JSON(ctx echo.Context) error {
	return ctx.JSON(e.Code, toErrorResponse(e.Message))
}

func toError(code int, message string) *ProjectError {
	return &ProjectError{Code: code, Message: message}
}

func toErrorF(code int, format string, args ...interface{}) *ProjectError {
	return &ProjectError{Code: code, Message: fmt.Sprintf(format, args...)}
}
