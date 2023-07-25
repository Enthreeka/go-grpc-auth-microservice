package apperror

import (
	"errors"
	"fmt"
)

type statusMessage string

const (
	invalidRequest statusMessage = "Invalid request"
)

var (
	ErrUserExist    = NewAppError(errors.New("user_exist"), invalidRequest)
	ErrUserNotExist = NewAppError(errors.New("user_not_exist"), invalidRequest)
)

type AppError struct {
	Err error
	Msg statusMessage
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Msg, e.Err)
}

func NewAppError(err error, msg statusMessage) *AppError {
	return &AppError{
		Err: err,
		Msg: msg,
	}
}

// Check user id. If id == zero value then function
func ValidateID(id string) error {
	if id == "" {
		return NewAppError(errors.New("id_empty"), invalidRequest)
	}
	return nil
}
