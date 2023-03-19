package plugincmds

import (
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/plugins"
	"github.com/urfave/cli/v2"
)

func UninstallCommand(cCtx *cli.Context) error {
	args := cCtx.Args().Slice()

	if len(args) == 0 {
		log.RawError("you must specify plugins to uninstall")
	}

	log.Log("fetching plugins...")

	info := plugins.GetManyPluginInfo(args, false, false)

	plugins.PluginList(info, "uninstalling")
	plugins.PluginDoMany(info, plugins.PluginUninstall)

	log.Success("successfully uninstalled all plugins")

	return nil
}
