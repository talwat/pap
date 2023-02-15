package plugins

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
	"github.com/talwat/pap/internal/plugins/jenkins"
	"github.com/talwat/pap/internal/plugins/sources/bukkit"
	"github.com/talwat/pap/internal/plugins/sources/modrinth"
	"github.com/talwat/pap/internal/plugins/sources/paplug"
	"github.com/talwat/pap/internal/plugins/sources/spigotmc"
)

func PluginDownload(plugin paplug.PluginInfo) {
	for _, download := range plugin.Downloads {
		var url string

		log.Log("getting download url...")

		if download.Type == "url" {
			url = download.URL
		} else if download.Type == "jenkins" {
			url = jenkins.GetJenkinsURL(download)
		}

		url = SubstituteProps(plugin, url)
		path := filepath.Join("plugins", download.Filename)

		net.Download(
			url,
			fmt.Sprintf("%s not found, please report this to https://github.com/talwat/pap/issues", url),
			path,
			download.Filename,
			nil,
			fs.ReadWritePerm,
		)

		if strings.HasSuffix(path, ".zip") {
			log.Log("unzipping %s...", path)
			fs.Unzip(path, "plugins/")

			log.Log("cleaning up...")
			fs.DeletePath(path)
		}
	}
}

func trimPluginName(name string, prefixes []string) string {
	trimmedName := name

	for _, prefix := range prefixes {
		trimmedName = strings.TrimPrefix(trimmedName, prefix)
	}

	return trimmedName
}

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

// This function will call itself in case of an alias.
//
//nolint:funlen
func GetPluginInfo(name string) paplug.PluginInfo {
	var info paplug.PluginInfo

	switch {
	// If it's a url using http then use this:
	case strings.HasPrefix(name, "https://") || strings.HasPrefix(name, "http://"):
		log.Debug("using url (%s)", name)
		net.Get(name, fmt.Sprintf("plugin at %s not found", name), &info)

		info.URL = name

	// If it's file which ends in .json try reading it locally:
	case strings.HasSuffix(name, ".json"):
		log.Debug("using local json file (%s)", name)
		raw := fs.ReadFile(name)

		log.Debug("unmarshaling %s...", name)

		err := json.Unmarshal(raw, &info)
		log.Error(err, "an error occurred while parsing %s", name)

		info.Path = name

	// If it's a modrinth plugin try getting it from modrinth:
	case strings.HasPrefix(name, "modrinth:"):
		log.Debug("using modrinth (%s)", name)
		info = getPluginFromSource(
			name,
			"modrinth",
			[]string{"modrinth:"},
			modrinth.GetPluginInfo,
		)

	// If it's a spigot plugin try getting it from spigotmc:
	case strings.HasPrefix(name, "spigot:"),
		strings.HasPrefix(name, "spigotmc:"):
		log.Debug("using spigotmc (%s)", name)
		info = getPluginFromSource(
			strings.ReplaceAll(name, "_", " "),
			"spigotmc",
			[]string{"spigot:", "spigotmc:"},
			spigotmc.GetPluginInfo,
		)

	// If it's a bukkit plugin try getting it from bukkit:
	case strings.HasPrefix(name, "bukkit:"),
		strings.HasPrefix(name, "bukkitdev:"):
		log.Debug("using bukkitdev (%s)", name)
		info = getPluginFromSource(
			name,
			"bukkit",
			[]string{"bukkit:", "bukkitdev:"},
			bukkit.GetPluginInfo,
		)

	// If it's none of the options above try getting it from the repos:
	default:
		log.Debug("using repos (%s)", name)
		net.Get(
			fmt.Sprintf(
				"https://raw.githubusercontent.com/talwat/pap/main/plugins/%s.json",
				name,
			),
			fmt.Sprintf("package %s not found", name),
			&info,
		)
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
