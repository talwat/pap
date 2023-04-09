package plugincmds

import (
	"github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/plugins"
	"github.com/urfave/cli/v2"
)

func InstallCommand(cCtx *cli.Context) error {
	args := cCtx.Args().Slice()

	if len(args) == 0 {
		log.RawError("you must specify plugins to install")
	}

	log.Debug("making plugins directory...")
	fs.MakeDirectory("plugins")

	log.Log("fetching plugins...")

	pluginsToInstall := plugins.GetManyPluginInfo(args, false, false, true)
	dependencies := plugins.ResolveDependencies(pluginsToInstall)

	// Append dependencies
	pluginsToInstall = append(pluginsToInstall, dependencies...)

	plugins.PluginList(pluginsToInstall, "installing")

	plugins.PluginDoMany(pluginsToInstall, plugins.PluginDownload)
	plugins.PluginDoMany(pluginsToInstall, plugins.PluginInstall)

	log.Success("successfully installed all plugins")

	// Display notes
	for _, plugin := range pluginsToInstall {
		plugins.DisplayAdditionalInfo(plugin)
	}

	return nil
}
