package plugincmds

import (
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/plugins"
	"github.com/urfave/cli/v2"
)

func InstallCommand(cCtx *cli.Context) error {
	args := cCtx.Args()

	log.Log("fetching packages...")

	pluginsToInstall := plugins.GetManyPluginInfo(args.Slice())
	dependencies := []plugins.PluginInfo{}

	log.Log("resolving dependencies...")

	for _, plugin := range pluginsToInstall {
		dependencies = append(dependencies, plugins.GetDependencies(plugin, pluginsToInstall)...)
	}

	plugins.PluginList(pluginsToInstall, dependencies, "installing")
	pluginsToInstall = append(pluginsToInstall, dependencies...)
	plugins.PluginDoMany(pluginsToInstall, plugins.PluginInstall)

	return nil
}
