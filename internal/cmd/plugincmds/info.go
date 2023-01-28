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

	output := "name: " + plugin.Name + "\n"
	output += "version: " + plugin.Version + "\n"

	if plugin.Site != "" {
		output += "site: " + plugin.Site + "\n"
	}

	output += "description: " + plugin.Description + "\n"
	output += "license: " + plugin.License + "\n"
	output += "authors: " + join(plugin.Authors) + "\n"

	if len(plugin.Dependencies) > 0 {
		output += "dependencies: " + join(plugin.Dependencies) + "\n"
	}

	if len(plugin.OptionalDependencies) > 0 {
		output += "optional dependencies:"
		for _, dependency := range plugin.OptionalDependencies {
			output += fmt.Sprintf("  %s: %s", dependency.Name, dependency.Purpose)
		}
	}

	log.OutputLog(strings.TrimSpace(output))

	return nil
}
