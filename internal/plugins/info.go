package plugins

import (
	"github.com/talwat/pap/internal/log"
)

func PluginList(plugins []PluginInfo, dependencies []PluginInfo, operation string) {
	log.Log("%s %d plugin(s):", operation, len(plugins))

	for _, plugin := range plugins {
		name := plugin.Name

		switch {
		case plugin.Path != "":
			log.RawLog("  %s (%s)\n", name, plugin.Path)
		case plugin.URL != "":
			log.RawLog("  %s (%s)\n", name, plugin.URL)
		default:
			log.RawLog("  %s\n", name)
		}
	}

	for _, dependency := range dependencies {
		name := dependency.Name
		log.RawLog("  %s [dependency]\n", name)
	}

	log.Continue("would you like to continue?")
}
