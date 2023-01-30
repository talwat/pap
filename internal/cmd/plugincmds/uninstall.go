package plugincmds

import (
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/plugins"
	"github.com/urfave/cli/v2"
)

func UninstallCommand(cCtx *cli.Context) error {
	args := cCtx.Args()

	info := plugins.GetManyPluginInfo(args.Slice())

	log.Log("fetching plugins...")
	plugins.PluginList(info, nil, "uninstalling")
	plugins.PluginDoMany(info, plugins.PluginUninstall)

	log.Success("successfully uninstalled all plugins")

	return nil
}
