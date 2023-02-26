package spigotmc

import (
	"fmt"
	"strings"

	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
	"github.com/talwat/pap/internal/plugins/sources"
	"github.com/talwat/pap/internal/plugins/sources/paplug"
)

// Gets a plugins website by checking it's links.
func getWebsite(plugin PluginInfo) string {
	switch {
	case plugin.SourceCodeLink != "":
		log.Debug("source code link isn't empty, using it (%s)", plugin.SourceCodeLink)

		return plugin.SourceCodeLink
	case plugin.DonationLink != "":
		log.Debug("donation link isn't empty, using it (%s)", plugin.DonationLink)

		return plugin.DonationLink
	default:
		url := fmt.Sprintf("https://www.spigotmc.org/resources/%d", plugin.ID)
		log.Debug("both donation link and plugin link are empty, falling back to the spigotmc page (%s)", url)

		return url
	}
}

// Gets a plugins author(s).
func getAuthors(plugin PluginInfo) []string {
	if plugin.Contributors == "" {
		log.Debug("contributors is empty, using authors information (%s)", plugin.Resolved.Author.Name)

		return []string{plugin.Resolved.Author.Name}
	}

	log.Debug("contributors is not empty, splitting it by ', ' (%s)", plugin.Contributors)

	return strings.Split(plugin.Contributors, ", ")
}

// Gets a plugins license by checking if it has a source code link.
func getLicense(plugin PluginInfo) string {
	if plugin.SourceCodeLink != "" {
		log.Debug("source code link is not empty, using unknown license")

		return sources.Undefined
	}

	log.Debug("source code link is empty, assuming app is proprietary")

	return "proprietary"
}

// Converts a spigotmc file into a paplug download.
// Path is the parsed filename for the plugin jarfile.
// NOTE: Any plugin that has an "external" download source cannot be downloaded.
func ConvertDownload(plugin PluginInfo, path string) paplug.Download {
	download := paplug.Download{}
	download.Type = "url"
	download.Filename = path

	if !plugin.Premium && plugin.File.FileType == ".jar" {
		log.Debug("%s has a direct download and isn't premium, adding download", plugin.Name)
		download.URL = fmt.Sprintf("https://api.spiget.org/v2/resources/%d/download", plugin.ID)
	} else {
		log.Debug("%s is either premium or doesn't have a .jar filetype", plugin.Name)
		download.URL = sources.Undefined
		log.Warn(
			"%s does not support downloading. if you are downloading %s as a plugin, you will get an error",
			plugin.Name,
			plugin.Name,
		)
	}

	return download
}

// Converts a spigotmc plugin into the paplug format.
func ConvertToPlugin(spigotPlugin PluginInfo) paplug.PluginInfo {
	plugin := paplug.PluginInfo{}

	plugin.Name = sources.FormatName(spigotPlugin.Name)
	plugin.Description = sources.FormatDesc(spigotPlugin.Tag)
	plugin.Site = getWebsite(spigotPlugin)
	plugin.Authors = getAuthors(spigotPlugin)
	plugin.License = getLicense(spigotPlugin)

	plugin.Install.Type = "simple"

	plugin.Version = spigotPlugin.Resolved.LatestVersion.Name

	// Unknown vars
	plugin.Note = []string{}
	plugin.Dependencies = []string{}
	plugin.OptionalDependencies = []string{}

	// File & Download
	path := fmt.Sprintf("%s.jar", plugin.Name)

	log.Debug("plugin jarfile path: %s", path)

	// File
	log.Debug("adding uninstall file...")

	file := paplug.File{}
	file.Path = path
	file.Type = "other"

	plugin.Uninstall.Files = append(plugin.Uninstall.Files, file)

	// Download
	log.Debug("adding download...")

	plugin.Downloads = append(plugin.Downloads, ConvertDownload(spigotPlugin, path))

	return plugin
}

// Gets a raw spigotmc plugin.
func Get(name string) PluginInfo {
	var plugins []PluginInfo

	net.Get(
		//nolint:lll
		fmt.Sprintf("https://api.spiget.org/v2/search/resources/%s?field=name&size=1&page=0&sort=-likes&fields=file,contributors,likes,name,tag,sourceCodeLink,donationLink,premium,id,version,author", name),
		fmt.Sprintf("spigot plugin %s not found", name),
		&plugins,
	)

	// Gets the first plugin because it's sorted by likes, so it should be fine.
	// Also spigotmc doesn't really have slugs.
	plugin := plugins[0]

	if plugin.Contributors == "" {
		net.Get(
			fmt.Sprintf("https://api.spiget.org/v2/authors/%d?fields=name", plugin.Author.ID),
			fmt.Sprintf("spigot author %d not found", plugin.Author.ID),
			&plugin.Resolved.Author,
		)
	}

	version := plugin.Version.ID

	net.Get(
		fmt.Sprintf("https://api.spiget.org/v2/resources/%d/versions/%d?fields=name", plugin.ID, version),
		fmt.Sprintf("spigot version %d not found", version),
		&plugin.Resolved.LatestVersion,
	)

	return plugin
}

// Gets & converts to the standard pap format.
func GetPluginInfo(name string) paplug.PluginInfo {
	return ConvertToPlugin(Get(name))
}
