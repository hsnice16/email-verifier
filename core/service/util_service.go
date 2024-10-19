// Package service provides core services for verifying an email address.
package service

import (
	"net"
	"regexp"
	"strings"

	"github.com/hsnice16/email-verifier/constant"
)

// VerifyEmailOptions struct to hold VerifyEmail options
type VerifyEmailOptions struct {
	ValidateRegex    bool // Validates email address using regex
	ValidateMxRecord bool // Validates MX records are present on DNS
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

	if options.ValidateMxRecord {
		err := validateMxRecord(email)

		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func validateRegex(email string) error {
	validateEmailRegexString := `^[\w\+\.-]+@[a-zA-Z\d\.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(validateEmailRegexString, email)

	if matched == false || err != nil {
		return constant.FailedRegexCheck
	}

	return nil
}

func validateMxRecord(email string) error {
	domain := strings.Split(email, "@")[1]
	_, err := net.LookupMX(domain)

	if err != nil {
		return constant.FailedMxRecordCheck
	}

	// bestMxRecord := mxRecords[0]
	return nil
}
