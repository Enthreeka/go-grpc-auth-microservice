package validation

import (
	"net/mail"
	"regexp"
)

var (
	passwordPattern = regexp.MustCompile(``)
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidPassword(password string) bool {
	return passwordPattern.MatchString(password)
}
