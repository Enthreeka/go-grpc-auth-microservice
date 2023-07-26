package validation

import "regexp"

var (
	emailPattern    = regexp.MustCompile(`^.*(?=.{8,})(?=.*[a-zA-Z])(?=.*\d)(?=.*[!#$%&? "]).*$`)
	passwordPattern = regexp.MustCompile(``)
)

func IsValidEmail(email string) bool {
	return emailPattern.MatchString(email)
}

func IsValidPassword(password string) bool {
	return passwordPattern.MatchString(password)
}
