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
	"github.com/talwat/pap/internal/plugins/modrinth"
	"github.com/talwat/pap/internal/plugins/paplug"
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

// This function will call itself in case of an alias.
func GetPluginInfo(name string) paplug.PluginInfo {
	var info paplug.PluginInfo

	switch {
	// If it's a url using http then use this:
	case strings.HasPrefix(name, "https://") || strings.HasPrefix(name, "http://"):
		net.Get(name, fmt.Sprintf("plugin at %s not found", name), &info)

		info.URL = name

	// If it's file which ends in .json try reading it locally:
	case strings.HasSuffix(name, ".json"):
		raw := fs.ReadFile(name)
		err := json.Unmarshal(raw, &info)

		log.Error(err, "an error occurred while parsing %s", name)

		info.Path = name

	// If it's a modrinth plugin try getting it from modrinth:
	case strings.HasPrefix(name, "modrinth:"):
		info = modrinth.GetPluginInfo(strings.TrimPrefix(name, "modrinth:"))

		info.Source = "modrinth"

	// If it's none of the options above try getting it from the repos:
	default:
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
