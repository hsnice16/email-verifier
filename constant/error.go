package constant

import "errors"

var (
	FailedRegexCheck       = errors.New("Invalid Email Address: Regex check failed")
	FailedMxRecordCheck    = errors.New("Invalid Email Address: Mx record not found")
	FailedSmtpRunningCheck = errors.New("Invalid Email Address: SMTP server is not running")

	FailedSmtpClose = errors.New("Failed to close the SMTP server")
)
