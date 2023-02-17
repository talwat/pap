package plugincmds

import (
	"fmt"
	"strings"

	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/plugins"
	"github.com/urfave/cli/v2"
)

func DisplayLineString(output *string, key string, value string) {
	if value == "" {
		return
	}

	*output += fmt.Sprintf("%s: %s\n", key, value)
}

func DisplayLineArray(output *string, key string, value []string) {
	if len(value) <= 0 {
		return
	}

	*output += fmt.Sprintf("%s: %s\n", key, strings.Join(value, ", "))
}

func InfoCommand(cCtx *cli.Context) error {
	args := cCtx.Args()
	name := args.Get(0)

	plugin := plugins.GetPluginInfo(name)
	output := ""

	// Not using a map because golang doesn't preserve order (annoyingly).

	DisplayLineString(&output, "name", plugin.Name)
	DisplayLineString(&output, "version", plugin.Version)
	DisplayLineString(&output, "site", plugin.Site)
	DisplayLineString(&output, "description", plugin.Description)
	DisplayLineString(&output, "license", plugin.License)

	DisplayLineArray(&output, "authors", plugin.Authors)
	DisplayLineArray(&output, "dependencies", plugin.Dependencies)
	DisplayLineArray(&output, "optional dependencies", plugin.OptionalDependencies)

	log.OutputLog(strings.TrimSpace(output))

	return nil
}
