package plugins

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
)

func GetJenkinsURL(download Download) string {
	var jenkinsBuild JenkinsBuild

	log.Log("getting jenkins build information...")
	net.Get(fmt.Sprintf("%s/lastSuccessfulBuild/api/json", download.Job), &jenkinsBuild)

	log.Log("finding correct artifact...")

	for _, artifact := range jenkinsBuild.Artifacts {
		matched, err := regexp.MatchString(download.Artifact, artifact.FileName)
		log.Error(err, "an error occurred while checking if %s is the correct artifact", artifact.FileName)

		if matched {
			log.Log("using %s", artifact.FileName)

			return fmt.Sprintf("%s/lastSuccessfulBuild/artifact/%s", download.Job, artifact.RelativePath)
		}
	}

	log.RawError("no artifacts matched, please report this to https://github.com/talwat/pap/issues")

	return ""
}

func PluginDownload(plugin PluginInfo) {
	for _, download := range plugin.Downloads {
		var url string

		log.Log("getting download url...")

		if download.Type == "url" {
			url = download.URL
		} else if download.Type == "jenkins" {
			url = GetJenkinsURL(download)
		}

		url = SubstituteProps(plugin, url)

		path := fmt.Sprintf("plugins/%s", download.Filename)

		net.Download(url, path, plugin.Name, nil)

		if strings.HasSuffix(path, ".zip") {
			log.Log("unzipping %s...", path)
			fs.Unzip(path, "plugins/")

			log.Log("cleaning up...")
			fs.DeletePath(path)
		}
	}
}

func GetPluginInfo(name string) PluginInfo {
	var info PluginInfo

	switch {
	case strings.HasPrefix(name, "https://") || strings.HasPrefix(name, "http://"):
		net.Get(name, &info)

		info.URL = name
	case strings.HasSuffix(name, ".json"):
		raw := fs.ReadFile(name)
		err := json.Unmarshal(raw, &info)

		log.Error(err, "an error occurred while parsing %s", name)

		info.Path = name
	default:
		net.Get(
			fmt.Sprintf(
				"https://raw.githubusercontent.com/talwat/pap/plugin-manager/plugins/%s.json",
				name,
			), &info)
	}

	if info.Alias != "" {
		log.Warn("%s is an alias to %s", name, info.Alias)

		return GetPluginInfo(info.Alias)
	}

	return info
}

func GetManyPluginInfo(plugins []string) []PluginInfo {
	pluginInfo := []PluginInfo{}

	for _, plugin := range plugins {
		pluginInfo = append(pluginInfo, GetPluginInfo(plugin))
	}

	return pluginInfo
}
