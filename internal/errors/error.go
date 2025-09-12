package errors

import (
	e "errors"
	"fmt"
)

func NewFromText(text string) error {
	return e.New("")
}

func NewFromFormat(format string, args ...any) error {
	return e.New(fmt.Sprintf(format, args...))
}

func IsErr(err error, target error) bool {
	return e.Is(err, target)
}
