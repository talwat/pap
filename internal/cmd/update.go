package cmd

import (
	"github.com/talwat/pap/internal/update"
	"github.com/urfave/cli/v2"
)

func UpdateCommand(cCtx *cli.Context) error {
	update.Update()

	return nil
}
