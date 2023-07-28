package internal

import (
	"fmt"
)

type MyError struct {
	Message string
	Code    int
}

func (err *MyError) Error() string {
	return fmt.Sprintf("Code: %d\nMessage: %s\n", err.Code, err.Message)
}
