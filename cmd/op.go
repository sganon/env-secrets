package cmd

import (
	"fmt"
	"os"

	"github.com/sganon/env-secrets/common"
	"github.com/sganon/env-secrets/onepassword"
	"github.com/urfave/cli"
)

const (
	domainNameFlagName = "domain"
)

// NewOP returns the subcommand handling 1password vaults
func NewOP() cli.Command {
	cmd := cli.Command{
		Name:  "op",
		Usage: "Get secrets from your 1password vault",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  domainNameFlagName,
				Usage: "Define your 1password subdomain",
			},
		},
		Before: beforeChecksOP,
		Action: func(c *cli.Context) (err error) {
			op := onepassword.OP{
				Tags: c.Args(),
			}
			if err = op.SetSecretsIDs(); err != nil {
				return err
			}
			if err = op.SetValues(); err != nil {
				return err
			}
			common.OutputEnv(op)
			return err
		},
	}
	return cmd
}

func beforeChecksOP(c *cli.Context) error {
	domain := c.String(domainNameFlagName)
	if os.Getenv("OP_SESSION_"+domain) == "" {
		return fmt.Errorf("env-secrets error: op session isn't set, op CLI must be installed and logged in")
	}
	if len(c.Args()) < 1 {
		return fmt.Errorf("env-secrets error: you need to provide at least one tag where to fetch secrets")
	}
	return nil
}
