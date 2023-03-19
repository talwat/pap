package cmd

import (
	"github.com/talwat/pap/internal/update"
	"github.com/urfave/cli/v2"
)

//nolint:revive // cCtx kept for consistency with other commands.
func UpdateCommand(cCtx *cli.Context) error {
	update.Update()

	return nil
}
