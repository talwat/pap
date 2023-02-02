package modrinth

import (
	"fmt"
	"strings"

	"github.com/talwat/pap/internal/net"
	"github.com/talwat/pap/internal/plugins/paplug"
)

func getWebsite(plugin PluginInfo) string {
	switch {
	case plugin.SourceURL != "":
		return plugin.SourceURL
	case plugin.WikiURL != "":
		return plugin.WikiURL
	case plugin.IssuesURL != "":
		return plugin.IssuesURL
	case plugin.DiscordURL != "":
		return plugin.DiscordURL
	default:
		return paplug.Undefined
	}
}

func ConvertToPlugin(modrinthPlugin PluginInfo) paplug.PluginInfo {
	plugin := paplug.PluginInfo{}

	plugin.Name = modrinthPlugin.Slug
	plugin.Description = modrinthPlugin.Description
	plugin.License = modrinthPlugin.License.ID
	plugin.Site = getWebsite(modrinthPlugin)

	if !strings.HasSuffix(plugin.Description, ".") {
		plugin.Description += "."
	}

	plugin.Install.Type = "simple"

	// Unknown vars
	plugin.Authors = []string{}
	plugin.Note = []string{}
	plugin.Dependencies = []string{}
	plugin.OptionalDependencies = []string{}

	var version Version

	net.Get(
		fmt.Sprintf("https://api.modrinth.com/v2/version/%s", modrinthPlugin.Versions[0]),
		fmt.Sprintf("version %s not found", modrinthPlugin.Versions[0]),
		&version,
	)

	plugin.Version = version.VersionNumber

	for _, file := range version.Files {
		download := paplug.Download{}

		download.Type = "url"
		download.URL = file.URL
		download.Filename = file.Filename
		plugin.Downloads = append(plugin.Downloads, download)

		uninstallFile := paplug.File{}

		uninstallFile.Path = file.Filename
		uninstallFile.Type = "other"

		plugin.Uninstall.Files = append(plugin.Uninstall.Files, uninstallFile)
	}

	return plugin
}

func Get(name string) PluginInfo {
	var modrinthPlugin PluginInfo

	net.Get(
		fmt.Sprintf("https://api.modrinth.com/v2/project/%s", name),
		fmt.Sprintf("modrinth plugin %s not found", name),
		&modrinthPlugin,
	)

	return modrinthPlugin
}

// Gets & converts to the standard pap format.
func GetPluginInfo(name string) paplug.PluginInfo {
	return ConvertToPlugin(Get(name))
}

// Note: This also converts the plugins to the standard pap format.
func GetManyPluginInfo(names []string) []paplug.PluginInfo {
	infos := []paplug.PluginInfo{}

	for _, name := range names {
		infos = append(infos, GetPluginInfo(name))
	}

	return infos
}
