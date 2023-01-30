package plugincmds

import (
	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/plugins"
	"github.com/urfave/cli/v2"
)

func InstallCommand(cCtx *cli.Context) error {
	args := cCtx.Args()

	log.Log("fetching plugins...")

	// This will later also contain the dependencies.
	pluginsToInstall := plugins.GetManyPluginInfo(args.Slice())

	// The plugins instructed to be installed, this does not include dependencies.
	pluginsNoDeps := pluginsToInstall

	dependencies := []plugins.PluginInfo{}

	if !global.NoDepsInput {
		log.Log("resolving dependencies...")

		for _, plugin := range pluginsToInstall {
			dependencies = append(dependencies, plugins.GetDependencies(plugin.Dependencies, pluginsToInstall)...)

			// Append optional dependencies aswell
			if global.InstallOptionalDepsInput {
				dependencies = append(dependencies, plugins.GetDependencies(plugin.OptionalDependencies, pluginsToInstall)...)
			}
		}
	}

	plugins.PluginList(pluginsToInstall, dependencies, "installing")
	pluginsToInstall = append(pluginsToInstall, dependencies...)
	plugins.PluginDoMany(pluginsToInstall, plugins.PluginInstall)

	log.Success("successfully installed all plugins")

	// Display notes
	for _, plugin := range pluginsNoDeps {
		plugins.DisplayAdditionalInfo(plugin)
	}

	return nil
}
