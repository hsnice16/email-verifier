# Email Verifier

[![PkgGoDev](https://pkg.go.dev/badge/github.com/hsnice16/email-verifier@v0.2.0/core/service#pkg-overview)](https://pkg.go.dev/github.com/hsnice16/email-verifier@v0.1.0/core/service)

> Enter an email and verify if it's a valid email or not.

## Installation

`email-verifier` exposes the core service to verify an email. You can get that using:

```shell
go get github.com/hsnice16/email-verifier/core/service
```

## Quickstart

```go
package main

import (
  "fmt"

  "github.com/hsnice16/email-verifier/core/service"
)

func main() {
  fmt.Printf("Hello, Email Verifier\n\n")

  email := "xyz@"
  result, err := service.VerifyEmail(email, service.VerifyEmailOptions{ValidateRegex: true})
  fmt.Printf("Is %+v a Valid Email?: %v\n", email, result)
  fmt.Printf("%+v\n\n", err)

  email = "xyz@example.com"
  result, err = service.VerifyEmail(email, service.VerifyEmailOptions{
    ValidateRegex:       true,
    ValidateMxRecord:    true,
    ValidateSmtpRunning: true,
  },
  )
  fmt.Printf("Is %+v a Valid Email?: %v\n", email, result)
  fmt.Println(err)
}
```

```text
Hello, Email Verifier

Is xyz@ a Valid Email?: false
Invalid Email Address: Regex check failed

Is xyz@example.com a Valid Email?: false
Invalid Email Address: Mx record not found
```

---

If you liked the project then give it a star (⭐️) and consider sponsoring it (❤️).
