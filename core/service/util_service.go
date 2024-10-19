// Package service provides core services for verifying an email address.
package service

import (
	"errors"
	"regexp"
)

type VerifyEmailOptions struct {
	ValidateRegex bool
}

// VerifyEmail verifies if the passed email string argument is
// a valid email or not. It returns the bool status and any
// error encountered.
func VerifyEmail(email string, options VerifyEmailOptions) (bool, error) {
	if options.ValidateRegex {
		err := validateRegex(email)
		
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func validateRegex(email string) error {
	validateEmailRegexString := `^[\w\.-]+@[a-zA-Z\d\.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(validateEmailRegexString, email)

	if(err != nil) {
		return err
	}

	if(matched == false) {
		return errors.New("Invalid Email Address")
	}

	return nil
}