package propcmds

import (
	"strings"

	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/properties"
	"github.com/urfave/cli/v2"
)

func EditPropertyCommand(cCtx *cli.Context) error {
	args := cCtx.Args()
	prop := args.Get(0)
	val := args.Tail()

	if prop == "" {
		log.RawError("property name is required")
	} else if len(val) == 0 {
		log.RawError("value is required")
	}

	properties.EditProperty(prop, strings.Join(val, " "))

	return nil
}
