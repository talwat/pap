package generatecmds

import (
	"github.com/talwat/pap/internal/plugins/sources/bukkit"
	"github.com/urfave/cli/v2"
)

func GenerateBukkit(cCtx *cli.Context) error {
	args := cCtx.Args().Slice()

	Generate(bukkit.GetPluginInfo, args)

	return nil
}
