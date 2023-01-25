package plugincmds

import (
	"fmt"
	"strings"

	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/plugins"
	"github.com/urfave/cli/v2"
)

func join(value []string) string {
	return strings.Join(value, ", ")
}

func InfoCommand(cCtx *cli.Context) error {
	args := cCtx.Args()
	name := args.Get(0)

	plugin := plugins.GetPluginInfo(name)

	log.OutputLog("name: " + plugin.Name)
	log.OutputLog("version: " + plugin.Version)

	if plugin.Site != "" {
		log.OutputLog("site: " + plugin.Site)
	}

	log.OutputLog("description: " + plugin.Description)
	log.OutputLog("license: " + plugin.License)
	log.OutputLog("authors: " + join(plugin.Authors))

	if len(plugin.Dependencies) > 0 {
		log.OutputLog("dependencies " + join(plugin.Dependencies))
	}

	if len(plugin.OptionalDependencies) > 0 {
		log.OutputLog("optional dependencies:")
		for _, dependency := range plugin.OptionalDependencies {
			log.OutputLog(fmt.Sprintf("  %s: %s", dependency.Name, dependency.Purpose))
		}
	}

	return nil
}
