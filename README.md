# Email Verifier

[![PkgGoDev](https://pkg.go.dev/badge/github.com/hsnice16/email-verifier@v0.1.0/core/service#pkg-overview)](https://pkg.go.dev/github.com/hsnice16/email-verifier@v0.1.0/core/service)

> Enter an email and verify if it's a valid email or not.

## Installation

`email-verifier` exposes the core service to verify an email. You can get that using:

```go
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
  fmt.Println("Hello, Email Verifier")

  email := "xyz@"
  result, _ := service.VerifyEmail(email)
  fmt.Printf("Is %+v a Valid Email?: %v\n", email, result)

  email = "xyz@example.com"
  result, _ = service.VerifyEmail(email)
  fmt.Printf("Is %+v a Valid Email?: %v\n", email, result)
}
```

```text
Hello, Email Verifier
Is xyz@ a Valid Email?: false
Is xyz@example.com a Valid Email?: true
```

---

If you liked the project then give it a star (⭐️) and consider sponsoring it (❤️).
