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

			return fmt.Sprintf("%s/lastSuccessfulBuild/artifact/artifacts/%s", download.Job, artifact.FileName)
		}
	}

	log.RawError("no artifacts matched, please report this to https://github.com/talwat/pap/issues")

	return ""
}

func PluginDownload(plugin PluginInfo) {
	for _, download := range plugin.Downloads {
		var url string

		if download.Type == "url" {
			url = download.URL
		} else if download.Type == "jenkins" {
			url = GetJenkinsURL(download)
		}

		toReplace := map[string]string{
			"version": plugin.Version,
			"name":    plugin.Name,
		}

		log.Log("parsing download URL...")

		for key, value := range toReplace {
			url = strings.ReplaceAll(url, fmt.Sprintf("{%s}", key), value)
		}

		net.Download(url, fmt.Sprintf("plugins/%s", download.Filename), plugin.Name, nil)
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
				strings.ToLower(name),
			), &info)
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
