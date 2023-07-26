package apperror

import (
	"errors"
	"fmt"
)

type statusMessage string

const (
	invalidRequest  statusMessage = "Invalid request"
	incorrectData   statusMessage = "Incorrect data"
	invalidPassword statusMessage = "Invalid password"
)

var (
	ErrHashPasswordsNotEqual = NewAppError(errors.New("hashes_not_equal"), invalidPassword)
)

var (
	ErrDataNotValid = NewAppError(errors.New("password_or_email_not_valid"), incorrectData)
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
