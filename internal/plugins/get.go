package plugins

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
	"github.com/talwat/pap/internal/plugins/sources/bukkit"
	"github.com/talwat/pap/internal/plugins/sources/modrinth"
	"github.com/talwat/pap/internal/plugins/sources/paplug"
	"github.com/talwat/pap/internal/plugins/sources/spigotmc"
)

// Trim plugin name of a list of prefixes.
func trimPluginName(name string, prefixes []string) string {
	trimmedName := name

	for _, prefix := range prefixes {
		trimmedName = strings.TrimPrefix(trimmedName, prefix)
	}

	return trimmedName
}

// Get plugin info & strip out the prefix (modrinth:, spigotmc:, etc...).
func getPluginFromSource(
	name string,
	source string,
	prefixes []string,
	getPluginInfo func(name string) paplug.PluginInfo,
) paplug.PluginInfo {
	pluginName := trimPluginName(name, prefixes)

	info := getPluginInfo(pluginName)
	info.Source = source

	return info
}

// Get using a URL, eg: 'https://example.com/plugin.json'.
func getURL(name string, info *paplug.PluginInfo) {
	log.Debug("using url (%s)", name)
	net.Get(name, fmt.Sprintf("plugin at %s not found", name), &info)

	info.URL = name
}

// Get using a local file, eg: 'plugins/plugin.json'.
func getLocal(name string, info *paplug.PluginInfo) {
	log.Debug("using local json file (%s)", name)
	raw := fs.ReadFile(name)

	log.Debug("unmarshaling %s...", name)

	err := json.Unmarshal(raw, &info)
	log.Error(err, "an error occurred while parsing %s", name)

	info.Path = name
}

// Get from modrinth, eg: 'modrinth:plugin'.
func getModrinth(name string, info *paplug.PluginInfo) {
	log.Debug("using modrinth (%s)", name)
	*info = getPluginFromSource(
		name,
		"modrinth",
		[]string{"modrinth:"},
		modrinth.GetPluginInfo,
	)
}

// Get from spigotmc, eg: 'spigotmc:plugin'.
func getSpigotmc(name string, info *paplug.PluginInfo) {
	log.Debug("using spigotmc (%s)", name)
	*info = getPluginFromSource(
		strings.ReplaceAll(name, "_", " "),
		"spigotmc",
		[]string{"spigot:", "spigotmc:"},
		spigotmc.GetPluginInfo,
	)
}

// Get from bukkitdev, eg: 'bukkitdev:plugin'.
func getBukkitdev(name string, info *paplug.PluginInfo) {
	log.Debug("using bukkitdev (%s)", name)
	*info = getPluginFromSource(
		name,
		"bukkit",
		[]string{"bukkit:", "bukkitdev:"},
		bukkit.GetPluginInfo,
	)
}

// Get from the repositories, eg: 'plugin'
func getRepos(name string, info *paplug.PluginInfo) {
	log.Debug("using repos (%s)", name)
	net.Get(
		fmt.Sprintf(
			"https://raw.githubusercontent.com/talwat/pap/main/plugins/%s.json",
			name,
		),
		fmt.Sprintf("plugin %s not found", name),
		&info,
	)
}

// This function will call itself in case of an alias.
func GetPluginInfo(name string) paplug.PluginInfo {
	var info paplug.PluginInfo

	switch {
	// If it's a url using http then use this:
	case strings.HasPrefix(name, "https://") || strings.HasPrefix(name, "http://"):
		getURL(name, &info)

	// If it's file which ends in .json try reading it locally:
	case strings.HasSuffix(name, ".json"):
		getLocal(name, &info)

	// If it's a modrinth plugin try getting it from modrinth:
	case strings.HasPrefix(name, "modrinth:"):
		getModrinth(name, &info)

	// If it's a spigot plugin try getting it from spigotmc:
	case strings.HasPrefix(name, "spigot:"),
		strings.HasPrefix(name, "spigotmc:"):
		getSpigotmc(name, &info)

	// If it's a bukkit plugin try getting it from bukkit:
	case strings.HasPrefix(name, "bukkit:"),
		strings.HasPrefix(name, "bukkitdev:"):
		getBukkitdev(name, &info)

	// If it's none of the options above try getting it from the repos:
	default:
		getRepos(name, &info)
	}

	if info.Alias != "" {
		log.Warn("%s is an alias to %s", name, info.Alias)

		return GetPluginInfo(info.Alias)
	}

	return info
}

func GetManyPluginInfo(plugins []string) []paplug.PluginInfo {
	pluginInfo := []paplug.PluginInfo{}

	for _, plugin := range plugins {
		pluginInfo = append(pluginInfo, GetPluginInfo(plugin))
	}

	return pluginInfo
}
