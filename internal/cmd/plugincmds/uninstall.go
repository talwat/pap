package plugincmds

import (
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/plugins"
	"github.com/urfave/cli/v2"
)

func UninstallCommand(cCtx *cli.Context) error {
	args := cCtx.Args().Slice()

	if len(args) < 1 {
		log.RawError("you must specify plugins to uninstall")
	}

	log.Log("fetching plugins...")

	info := plugins.GetManyPluginInfo(args)

	plugins.PluginList(info, nil, "uninstalling")
	plugins.PluginDoMany(info, plugins.PluginUninstall)

	log.Success("successfully uninstalled all plugins")

	return nil
}
