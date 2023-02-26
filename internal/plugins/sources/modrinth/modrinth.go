package modrinth

import (
	"fmt"
	"strings"

	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
	"github.com/talwat/pap/internal/plugins/sources"
	"github.com/talwat/pap/internal/plugins/sources/paplug"
)

// Gets a website for a modrinth plugin.
func getWebsite(plugin PluginInfo) string {
	switch {
	case plugin.SourceURL != "":
		log.Debug("source url isn't empty, using it (%s)", plugin.SourceURL)

		return plugin.SourceURL
	case plugin.WikiURL != "":
		log.Debug("wiki url isn't empty, using it (%s)", plugin.WikiURL)

		return plugin.WikiURL
	case plugin.IssuesURL != "":
		log.Debug("issues url isn't empty, using it (%s)", plugin.IssuesURL)

		return plugin.IssuesURL
	case plugin.DiscordURL != "":
		log.Debug("discord url isn't empty, using it (%s)", plugin.DiscordURL)

		return plugin.DiscordURL
	default:
		url := fmt.Sprintf("https://modrinth.com/mod/%s", plugin.Slug)
		log.Debug("no links defined, falling back to modrinth page (%s)", url)

		return fmt.Sprintf("https://modrinth.com/mod/%s", plugin.Slug)
	}
}

// Converts a modrinth plugin into the paplug format.
func ConvertToPlugin(modrinthPlugin PluginInfo) paplug.PluginInfo {
	plugin := paplug.PluginInfo{}

	plugin.Name = sources.FormatName(modrinthPlugin.Slug)
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

	plugin.Version = modrinthPlugin.ResolvedVersion.VersionNumber

	for _, file := range modrinthPlugin.ResolvedVersion.Files {
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

// Gets a raw modrinth plugin.
func Get(name string) PluginInfo {
	var modrinthPlugin PluginInfo

	net.Get(
		fmt.Sprintf("https://api.modrinth.com/v2/project/%s", name),
		fmt.Sprintf("modrinth plugin %s not found", name),
		&modrinthPlugin,
	)

	// This may take a beta version.
	// But to get only a stable one it would require sending a GET request for potentially every version.
	// Which is very slow.
	version := modrinthPlugin.Versions[0]

	net.Get(
		fmt.Sprintf("https://api.modrinth.com/v2/version/%s", version),
		fmt.Sprintf("version %s not found", version),
		&modrinthPlugin.ResolvedVersion,
	)

	return modrinthPlugin
}

// Gets & converts to the standard pap format.
func GetPluginInfo(name string) paplug.PluginInfo {
	return ConvertToPlugin(Get(name))
}
