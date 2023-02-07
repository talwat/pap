package generatecmds

import (
	"github.com/talwat/pap/internal/plugins/sources/spigotmc"
	"github.com/urfave/cli/v2"
)

func GenerateSpigotMC(cCtx *cli.Context) error {
	args := cCtx.Args().Slice()

	Generate(spigotmc.GetPluginInfo, args)

	return nil
}
