package propcmds

import (
	"strings"

	"github.com/talwat/pap/app/log"
	"github.com/talwat/pap/app/properties"
	"github.com/urfave/cli/v2"
)

func EditPropertyCommand(cCtx *cli.Context) error {
	prop := cCtx.Args().Get(0)
	val := cCtx.Args().Tail()

	if prop == "" {
		log.CustomError("property name is required")
	} else if len(val) == 0 {
		log.CustomError("value is required")
	}

	properties.EditProperty(prop, strings.Join(val, " "))

	return nil
}
