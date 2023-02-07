package generatecmds

import (
	"github.com/talwat/pap/internal/plugins/sources/modrinth"
	"github.com/urfave/cli/v2"
)

func GenerateModrinth(cCtx *cli.Context) error {
	args := cCtx.Args().Slice()

	Generate(modrinth.GetPluginInfo, args)

	return nil
}
