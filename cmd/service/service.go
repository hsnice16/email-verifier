package service

import (
	"github.com/hsnice16/email-verifier/cmd/service/www"
	"github.com/urfave/cli/v2"
)

var (
	ServiceCommand = &cli.Command{
		Name:        "service",
		Usage:       "Start the service",
		Description: "Email verifier service sub command",
		Action:      service,
	}
)

func service(cli *cli.Context) error {
	www.Run()
	return nil
}
