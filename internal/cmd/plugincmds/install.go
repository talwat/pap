package plugincmds

import (
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/plugins"
	"github.com/urfave/cli/v2"
)

func InstallCommand(cCtx *cli.Context) error {
	args := cCtx.Args().Slice()

	if len(args) < 1 {
		log.RawError("you must specify plugins to install")
	}

	log.Log("fetching plugins...")

	// This will later also contain the dependencies.
	pluginsToInstall := plugins.GetManyPluginInfo(args)
	dependencies := plugins.ResolveDependencies(pluginsToInstall)

	plugins.PluginList(pluginsToInstall, dependencies, "installing")

	pluginsToInstall = append(pluginsToInstall, dependencies...)

	plugins.PluginDoMany(pluginsToInstall, plugins.PluginInstall)

	log.Success("successfully installed all plugins")

	// Display notes
	for _, plugin := range pluginsToInstall {
		plugins.DisplayAdditionalInfo(plugin)
	}

	return nil
}
