package cmd_test

import (
	"testing"

	"github.com/sganon/env-secrets/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type testCLI struct {
	Args             []string
	ExpectedLogLevel string
	ShouldErr        bool
}

var testsCLI []testCLI = []testCLI{
	{
		Args:             []string{"./env-secrets"},
		ExpectedLogLevel: "info",
		ShouldErr:        false,
	},
	{
		Args:             []string{"./env-secrets", "--logLevel", "debug"},
		ExpectedLogLevel: "debug",
		ShouldErr:        false,
	},
	{
		Args:             []string{"./env-secrets", "--logLevel", "42"},
		ShouldErr:        true,
		ExpectedLogLevel: "info",
	},
}

func TestCLI(t *testing.T) {
	for i, te := range testsCLI {
		cli := cmd.NewCLI()
		err := cli.Run(te.Args)
		if te.ShouldErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, te.ExpectedLogLevel, log.GetLevel().String())
		t.Logf("Test at index %d passed\n", i)
		// Reset logLevel to default
		log.SetLevel(log.InfoLevel)
	}
}
