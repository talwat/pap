package plugins

import (
	"github.com/talwat/pap/internal/log"
)

func PluginList(plugins []PluginInfo, deps []PluginInfo, operation string) {
	log.Log("%s %d plugin(s):", operation, len(plugins))

	for _, plugin := range plugins {
		name := plugin.Name

		switch {
		case plugin.Path != "":
			log.RawLog("  %s %s (%s)\n", name, plugin.Version, plugin.Path)
		case plugin.URL != "":
			log.RawLog("  %s %s (%s)\n", name, plugin.Version, plugin.URL)
		default:
			log.RawLog("  %s %s\n", name, plugin.Version)
		}
	}

	for _, dep := range deps {
		name := dep.Name
		log.RawLog("  %s %s [dependency]\n", name, dep.Version)
	}

	log.Continue("would you like to continue?")
}
