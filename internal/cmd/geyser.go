package cmd

// This should be purged after the 1.0 release including the entry in the main.go file.

import (
	"github.com/talwat/pap/internal/log"
	"github.com/urfave/cli/v2"
)

func GeyserCommand(cCtx *cli.Context) error {
	log.RawError("this command has been replaced by: pap plugin install --optional geyser")

	return nil
}
