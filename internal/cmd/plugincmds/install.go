package plugincmds

import (
	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/log/color"
	"github.com/talwat/pap/internal/plugins"
	"github.com/urfave/cli/v2"
)

func InstallCommand(cCtx *cli.Context) error {
	args := cCtx.Args()

	log.Log("fetching plugins...")

	pluginsToInstall := plugins.GetManyPluginInfo(args.Slice())
	dependencies := []plugins.PluginInfo{}

	if !global.NoDepsInput {
		log.Log("resolving dependencies...")

		for _, plugin := range pluginsToInstall {
			dependencies = append(dependencies, plugins.GetDependencies(plugin, pluginsToInstall)...)
		}
	}

	plugins.PluginList(pluginsToInstall, dependencies, "installing")
	pluginsToInstall = append(pluginsToInstall, dependencies...)
	plugins.PluginDoMany(pluginsToInstall, plugins.PluginInstall)

	log.Success("successfully installed all plugins")

	// Display notes
	for _, plugin := range pluginsToInstall {
		if len(plugin.Note) < 1 {
			continue
		}

		log.RawLog("\n")

		for _, line := range plugin.Note {
			log.Log("%simportant note%s from %s: %s", color.BrightBlue, color.Reset, plugin.Name, line)
		}
	}

	return nil
}
