package generatecmds

import (
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/plugins/modrinth"
	"github.com/urfave/cli/v2"
)

func GenerateModrinth(cCtx *cli.Context) error {
	args := cCtx.Args()

	log.Log("getting plugins to write...")

	pluginsToWrite := modrinth.GetManyPluginInfo(args.Slice())

	for _, plugin := range pluginsToWrite {
		WritePlugin(plugin)
	}

	log.Success("all plugins generated successfully!")

	return nil
}
