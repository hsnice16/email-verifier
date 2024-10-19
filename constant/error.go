package constant

import "errors"

var (
	FailedRegexCheck    = errors.New("Invalid Email Address: Regex check failed")
	FailedMxRecordCheck = errors.New("Invalid Email Address: Mx record not found")
)
