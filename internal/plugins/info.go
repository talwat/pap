package plugins

import (
	"fmt"

	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/log/color"
	"github.com/talwat/pap/internal/plugins/sources/paplug"
)

func DisplayNote(plugin paplug.PluginInfo) {
	if len(plugin.Note) == 0 {
		return
	}

	log.Log("%simportant note%s from %s:", color.BrightBlue, color.Reset, plugin.Name)

	for _, line := range plugin.Note {
		log.RawLog("  %s\n", line)
	}
}

func DisplayOptionalDependencies(plugin paplug.PluginInfo) {
	if len(plugin.OptionalDependencies) == 0 || global.InstallOptionalDepsInput {
		return
	}

	log.Log("%soptional dependencies%s from %s:", color.BrightBlue, color.Reset, plugin.Name)

	for _, dep := range plugin.OptionalDependencies {
		log.RawLog("  %s\n", dep)
	}
}

func DisplayAdditionalInfo(plugin paplug.PluginInfo) {
	if len(plugin.Note) == 0 && len(plugin.OptionalDependencies) == 0 {
		return
	}

	log.RawLog("\n")
	log.Log("additional information for %s%s%s", color.BrightBlue, plugin.Name, color.Reset)

	DisplayNote(plugin)
	DisplayOptionalDependencies(plugin)
}

func displayPluginLine(plugin paplug.PluginInfo) {
	pluginLine := fmt.Sprintf("  %s %s", plugin.Name, plugin.Version)

	switch {
	case plugin.Path != "":
		pluginLine += fmt.Sprintf(" (%s)", plugin.Path)
	case plugin.URL != "":
		pluginLine += fmt.Sprintf(" (%s)", plugin.URL)
	case plugin.Source != "":
		pluginLine += fmt.Sprintf(" (%s)", plugin.Source)
	}

	if plugin.IsDependency {
		pluginLine += " [dependency]"
	}

	if plugin.IsOptionalDependency {
		pluginLine += " [optional dependency]"
	}

	log.RawLog("%s\n", pluginLine)
}

// List out plugins.
func PluginList(plugins []paplug.PluginInfo, operation string) {
	log.Log("%s %d plugin(s):", operation, len(plugins))

	for _, plugin := range plugins {
		displayPluginLine(plugin)
	}

	log.Continue("would you like to continue?")
}
