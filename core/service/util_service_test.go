package service

import "testing"

// TestVerifyEmailCorrect calls service.VerifyEmail with a correct value
func TestVerifyEmailCorrect(t *testing.T) {
	email := "xyz@example.abc"
	isValid, err := VerifyEmail(email)

	if isValid == false {
		t.Fatalf("String %+v is not a valid email.", email)
	}

	if err != nil {
		t.Fatalf("Err %+v", err)
	}

	t.Logf("String %+v is a valid email.", email)
}

// TestVerifyEmailEmpty calls service.VerifyEmail with an empty value
func TestVerifyEmailEmpty(t *testing.T) {
	email := ""
	isValid, err := VerifyEmail(email)

	if isValid == true {
		t.Fatalf("String %+v is a valid email.", email)
	}

	if err != nil {
		t.Fatalf("Err %+v", err)
	}

	t.Logf("String `%+v` is not a valid email. Empty String.", email)
}

// TestVerifyEmailIncorrect calls service.VerifyEmail with an incorrect value
func TestVerifyEmailIncorrect(t *testing.T) {
	email := "a@xyz"
	isValid, err := VerifyEmail(email)

	if isValid == true {
		t.Fatalf("String %+v is valid email.", email)
	}

	if err != nil {
		t.Fatalf("Err %+v", err)
	}

	t.Logf("String %+v is an invalid email.", email)
}
