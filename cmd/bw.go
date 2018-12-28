package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// NewBW returns the subcommand handling bitwarden vaults
func NewBW() cli.Command {
	cmd := cli.Command{
		Name:   "bw",
		Usage:  "Get secrets from your bitwarden vault",
		Before: beforeChecks,
		Action: func(c *cli.Context) (err error) {
			log.Debugf("Calling bw with category: %s\n", c.Args().Get(0))
			return err
		},
	}
	return cmd
}

func beforeChecks(c *cli.Context) error {
	if len(c.Args()) != 1 {
		return fmt.Errorf("env-secrets error: bw command accepts only 1 arguments but got %d", len(c.Args()))
	}
	if os.Getenv("BW_SESSION") == "" {
		return fmt.Errorf("env-secrets error: bw session isn't set, bitwarden CLI must be installed and logged in")
	}
	return nil
}
