package plugins

import (
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/log/color"
	"github.com/talwat/pap/internal/plugins/sources/paplug"
)

func DisplayAdditionalInfo(plugin paplug.PluginInfo) {
	if len(plugin.Note) < 1 && len(plugin.OptionalDependencies) < 1 {
		return
	}

	log.RawLog("\n")
	log.Log("additional information for %s%s%s", color.BrightBlue, plugin.Name, color.Reset)

	if len(plugin.Note) > 0 {
		log.Log("%simportant note%s from %s:", color.BrightBlue, color.Reset, plugin.Name)

		for _, line := range plugin.Note {
			log.RawLog("  %s\n", line)
		}
	}

	if len(plugin.OptionalDependencies) > 0 {
		log.Log("%soptional dependencies%s from %s:", color.BrightBlue, color.Reset, plugin.Name)

		for _, dep := range plugin.OptionalDependencies {
			log.RawLog("  %s\n", dep)
		}
	}
}

func displayPluginLine(name string, plugin paplug.PluginInfo) {
	switch {
	case plugin.Path != "":
		log.RawLog("  %s %s (%s)\n", name, plugin.Version, plugin.Path)
	case plugin.URL != "":
		log.RawLog("  %s %s (%s)\n", name, plugin.Version, plugin.URL)
	case plugin.Source != "":
		log.RawLog("  %s %s (%s)\n", name, plugin.Version, plugin.Source)
	default:
		log.RawLog("  %s %s\n", name, plugin.Version)
	}
}

func PluginList(plugins []paplug.PluginInfo, deps []paplug.PluginInfo, operation string) {
	log.Log("%s %d plugin(s):", operation, len(plugins))

	for _, plugin := range plugins {
		name := plugin.Name

		displayPluginLine(name, plugin)
	}

	for _, dep := range deps {
		name := dep.Name
		log.RawLog("  %s %s [dependency]\n", name, dep.Version)
	}

	log.Continue("would you like to continue?")
}
