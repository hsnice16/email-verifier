// Package service provides core services for verifying an email address.
package service

import "regexp"

// VerifyEmail verifies if the passed email string argument is
// a valid email or not. It returns the bool status and any
// error encountered.
func VerifyEmail(email string) (bool, error) {
	validateEmailRegexString := `^[\w\.-]+@[a-zA-Z\d\.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(validateEmailRegexString, email)
	return matched, err
}
