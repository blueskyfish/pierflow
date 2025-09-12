package errors

import "fmt"

type ErrorResponse struct {
	Message string `json:"message"`
}

func ToErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{Message: message}
}

func ToErrorResponseF(format string, args ...any) *ErrorResponse {
	return ToErrorResponse(fmt.Sprintf(format, args...))
}
