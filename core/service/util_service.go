// Package service provides core services for verifying an email address.
package service

import (
	"net"
	"net/smtp"
	"regexp"
	"strings"

	"github.com/hsnice16/email-verifier/constant"
)

// VerifyEmailOptions struct to hold VerifyEmail options
type VerifyEmailOptions struct {
	ValidateRegex       bool // Validates email address using regex
	ValidateMxRecord    bool // Validates MX records are present on DNS
	ValidateSmtpRunning bool // Validates SMTP server is running
}

// VerifyEmail verifies if the passed email string argument is
// a valid email or not. It returns the bool status and any
// error encountered.
func VerifyEmail(email string, options VerifyEmailOptions) (bool, error) {
	if options.ValidateRegex {
		err := ValidateRegex(email)

		if err != nil {
			return false, err
		}
	}

	if options.ValidateMxRecord {
		err := ValidateMxRecord(email)

		if err != nil {
			return false, err
		}
	}

	if options.ValidateSmtpRunning {
		err := ValidateSmtpRunning(email)

		if err != nil {
			return false, err
		}
	}

	return true, nil
}

// ValidateRegex validates the email address using regex
func ValidateRegex(email string) error {
	validateEmailRegexString := `^[\w\+\.-]+@[a-zA-Z\d\.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(validateEmailRegexString, email)

	if matched == false || err != nil {
		return constant.FailedRegexCheck
	}

	return nil
}

// ValidateMxRecord validates the email address by checking
// if the MX records are present on DNS
func ValidateMxRecord(email string) error {
	domain := strings.Split(email, "@")[1]
	_, err := net.LookupMX(domain)

	if err != nil {
		return constant.FailedMxRecordCheck
	}

	return nil
}

// ValidateSmtpRunning validates the email address by checking
// if the SMTP server is running
func ValidateSmtpRunning(email string) error {
	domain := strings.Split(email, "@")[1]
	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		return constant.FailedMxRecordCheck
	}

	bestMxRecord := mxRecords[0]
	client, err := smtp.Dial(bestMxRecord.Host + ":smtp")
	if err != nil {
		return constant.FailedSmtpRunningCheck
	}

	err = client.Close()
	if err != nil {
		return constant.FailedSmtpClose
	}

	return nil
}
