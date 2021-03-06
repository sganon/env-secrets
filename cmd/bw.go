package cmd

import (
	"fmt"
	"os"

	"github.com/sganon/env-secrets/bitwarden"
	"github.com/sganon/env-secrets/common"
	"github.com/urfave/cli"
)

// NewBW returns the subcommand handling bitwarden vaults
func NewBW() cli.Command {
	cmd := cli.Command{
		Name:   "bw",
		Usage:  "Get secrets from your bitwarden vault",
		Before: beforeChecksBW,
		Action: func(c *cli.Context) (err error) {
			bw := bitwarden.BW{}
			if err = bw.SetFoldersIDs(c.Args()); err != nil {
				return err
			}
			if err = bw.FetchItems(); err != nil {
				return err
			}
			common.OutputEnv(bw)
			return err
		},
	}
	return cmd
}

func beforeChecksBW(c *cli.Context) error {
	if len(c.Args()) < 1 {
		return fmt.Errorf("env-secrets error: you need to provide at least one folder")
	}
	if os.Getenv("BW_SESSION") == "" {
		return fmt.Errorf("env-secrets error: bw session isn't set, bitwarden CLI must be installed and logged in")
	}
	return nil
}
