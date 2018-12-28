package main

import (
	"os"

	"github.com/sganon/env-secrets/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	cli := cmd.NewCLI()
	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
