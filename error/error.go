package error

import (
    `fmt`
)

func New(message string, params ...interface{}) *Error {
    return &Error{
        message: fmt.Sprintf(message, params...),
    }
}

type Error struct {
    message string
}

func (e *Error) Error() string {
    return e.message
}
