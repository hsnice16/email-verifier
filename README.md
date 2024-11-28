<div align="center">

# ✨ Email Verifier

![Forks](https://img.shields.io/github/forks/hsnice16/email-verifier)
![Stars](https://img.shields.io/github/stars/hsnice16/email-verifier)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hsnice16/email-verifier@v0.2.1/core/service#pkg-overview)](https://pkg.go.dev/github.com/hsnice16/email-verifier@v0.2.1/core/service)

> Enter an email and verify if it's a valid email or not.

</div>

---

## ⬇️ Installation

`email-verifier` exposes the core service to verify an email. You can get that using:

```shell
go get github.com/hsnice16/email-verifier/core/service
```

---

## 🎬 Quickstart

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

## ⚙️ Develop Locally

### Pull the code

```shell
git clone git@github.com:hsnice16/email-verifier.git
cd email-verifier
```

### Run in a Docker Container or start the server manually

#### ➡️ Run in a Docker Container

You will need [docker](https://www.docker.com/get-started/) on your local machine for this.

##### • Build the image

```shell
docker build -t hsnice16/email-verifier:1.0 .
```

##### • Run container

```shell
docker run -p 8080:8080 hsnice16/email-verifier:1.0 -- service
```

<div align="center">
• • •
</div>

#### ➡️ Run manually

##### • Install our command to the `$GOPATH/bin` directory

```shell
go install
```

<!--

`go build` vs `go install`:

`go build` just compiles the executable file and moves it to the destination.
`go install` does a little bit more. It moves the executable file to `$GOPATH/bin`
and caches all non-main packages which are imported to `$GOPATH/pkg`. The cache
will be used during the next compilation provided the source did not change yet.

-->

##### • Run our new command

```shell
cmd
```

##### • Start the core service

```shell
cmd service
```

---

## 💚 Sponsor

If you found this project helpful, then do consider sponsoring it - [Sponsor](https://github.com/sponsors/hsnice16)

And, give it a star 🌟
