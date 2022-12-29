package propcmds

import (
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/properties"
	"github.com/urfave/cli/v2"
)

func GetPropertyCommand(cCtx *cli.Context) error {
	prop := cCtx.Args().Get(0)

	if prop == "" {
		log.RawError("property name is required")
	}

	val := properties.GetProperty(prop)
	log.OuptutLog("%s", val)

	return nil
}
