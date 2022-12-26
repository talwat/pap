package propcmds

import (
	"github.com/talwat/pap/app/properties"
	"github.com/urfave/cli/v2"
)

func ResetPropertiesCommand(cCtx *cli.Context) error {
	properties.ResetProperties()

	return nil
}
