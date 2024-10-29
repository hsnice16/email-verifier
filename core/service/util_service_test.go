package service

import "testing"

// TestVerifyEmailCorrect calls service.VerifyEmail
// to validate Regex with a correct value
func TestVerifyEmailCorrect(t *testing.T) {
	email := "xyz@example.abc"
	isValid, err := VerifyEmail(email, VerifyEmailOptions{ValidateRegex: true})

	if isValid == false {
		t.Fatalf("String %+v is not a valid email.", email)
	}

	if err != nil {
		t.Fatal(err)
		t.Logf("String %+v is not a valid email.", email)
	}

	t.Logf("String %+v is a valid email.", email)
}

// TestVerifyEmailCorrectWithSpecialChars calls service.VerifyEmail
// to validate Regex with a correct value
func TestVerifyEmailCorrectWithSpecialChars(t *testing.T) {
	email := "x+y.-z@example.abc.in"
	isValid, err := VerifyEmail(email, VerifyEmailOptions{ValidateRegex: true})

	if isValid == false {
		t.Fatalf("String %+v is not a valid email.", email)
	}

	if err != nil {
		t.Fatal(err)
		t.Logf("String %+v is not a valid email.", email)
	}

	t.Logf("String %+v is a valid email.", email)
}

// TestVerifyEmailEmpty calls service.VerifyEmail
// to validate Regex with an empty value
func TestVerifyEmailEmpty(t *testing.T) {
	email := ""
	isValid, err := VerifyEmail(email, VerifyEmailOptions{ValidateRegex: true})

	if isValid == true {
		t.Fatalf("String %+v is a valid email.", email)
	}

	t.Log(err)
	t.Logf("String %+v is not a valid email.", email)
}

// TestVerifyEmailIncorrect calls service.VerifyEmail
// to validate Regex with an incorrect value
func TestVerifyEmailIncorrect(t *testing.T) {
	email := "a@xyz"
	isValid, err := VerifyEmail(email, VerifyEmailOptions{ValidateRegex: true})

	if isValid == true {
		t.Fatalf("String %+v is a valid email.", email)
	}

	t.Log(err)
	t.Logf("String %+v is not a valid email.", email)
}

// TestVerifyEmailIncorrectWithoutAtTheRate calls service.VerifyEmail
// to validate Regex with an incorrect value
func TestVerifyEmailIncorrectWithoutAtTheRate(t *testing.T) {
	email := "axyzgmail.com"
	isValid, err := VerifyEmail(email, VerifyEmailOptions{ValidateRegex: true})

	if isValid == true {
		t.Fatalf("String %+v is a valid email.", email)
	}

	t.Log(err)
	t.Logf("String %+v is not a valid email.", email)
}

// TestVerifyEmailMxRecordCorrect calls service.VerifyEmail
// to validate MX Record with a correct value
func TestVerifyEmailMxRecordCorrect(t *testing.T) {
	email := "xyz@gmail.com"
	isValid, err := VerifyEmail(email, VerifyEmailOptions{ValidateMxRecord: true})

	if isValid == false {
		t.Fatalf("String %+v is not a valid email.", email)
	}

	if err != nil {
		t.Fatal(err)
		t.Logf("String %+v is not a valid email.", email)
	}

	t.Logf("String %+v is a valid email.", email)
}

// TestVerifyEmailMxRecordIncorrect calls service.VerifyEmail
// to validate MX Record with an incorrect value
func TestVerifyEmailMxRecordInCorrect(t *testing.T) {
	email := "xyz@foo.bar"
	isValid, err := VerifyEmail(email, VerifyEmailOptions{ValidateMxRecord: true})

	if isValid == true {
		t.Fatalf("String %+v is a valid email.", email)
	}

	t.Log(err)
	t.Logf("String %+v is not a valid email.", email)
}

// TestVerifyEmailSmtpServerRunningCorrect calls service.VerifyEmail
// to validate SMTP Server Running with a correct value
func TestVerifyEmailSmtpServerRunningCorrect(t *testing.T) {
	email := "xyz@gmail.com"
	isValid, err := VerifyEmail(email, VerifyEmailOptions{ValidateSmtpRunning: true})

	if isValid == false {
		t.Fatalf("String %+v is not a valid email.", email)

	}

	if err != nil {
		t.Fatal(err)
		t.Fatalf("String %+v is not a valid email.", email)
	}

	t.Logf("String %+v is a valid email.", email)
}

// TestVerifyEmailSmtpServerRunningIncorrect calls service.VerifyEmail
// to validate SMTP Server Running with an incorrect value
func TestVerifyEmailSmtpServerRunningInCorrect(t *testing.T) {
	email := "xyz@foo.bar"
	isValid, err := VerifyEmail(email, VerifyEmailOptions{ValidateSmtpRunning: true})

	if isValid == true {
		t.Fatalf("String %+v is a valid email.", email)
	}

	t.Log(err)
	t.Logf("String %+v is not a valid email.", email)
}

// TestVerifyEmailDisposableCorrect calls service.VerifyEmail
// to validate Disposable Email with a correct value
func TestVerifyEmailDisposableCorrect(t *testing.T) {
	email := "xyz@gmail.com"
	isValid, err := VerifyEmail(email, VerifyEmailOptions{ValidateDisposableEmail: true})

	if isValid == false {
		t.Fatalf("String %+v is not a valid email.", email)
	}

	if err != nil {
		t.Fatal(err)
		t.Fatalf("String %+v is not a valid email.", email)
	}

	t.Logf("String %+v is a valid email.", email)
}

func TestVerifyEmailDisposableInCorrect(t *testing.T) {
	email := "abc@risu.be"
	isValid, err := VerifyEmail(email, VerifyEmailOptions{ValidateDisposableEmail: true})

	if isValid == true {
		t.Fatalf("String %+v is a valid email.", email)
	}

	t.Log(err)
	t.Logf("String %+v is not a valid email.", email)
}
