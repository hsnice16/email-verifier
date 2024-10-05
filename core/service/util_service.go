package service

import "regexp"

func VerifyEmail(email string) (bool, error) {
	validateEmailRegexString := `^[\w\.-]+@[a-zA-Z\d\.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(validateEmailRegexString, email)
	return matched, err
}
