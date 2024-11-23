package main

import (
	"log"
	"os"

	"github.com/hsnice16/email-verifier/cmd/service"
	"github.com/urfave/cli/v2"
)

// @externalDocs.description	pkg.go.dev
// @externalDocs.url			https://pkg.go.dev/github.com/hsnice16/email-verifier@v0.2.1/core/service
func main() {
	app := &cli.App{
		Name:                 "cmd",
		Usage:                "The email verifier command line interface",
		Copyright:            "(c) 2024 Email Verifier Authors",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			service.ServiceCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
