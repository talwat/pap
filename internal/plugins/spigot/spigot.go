package spigot

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/talwat/pap/internal/net"
	"github.com/talwat/pap/internal/plugins/paplug"
)

func getWebsite(plugin PluginInfo) string {
	switch {
	case plugin.SourceCodeLink != "":
		return plugin.SourceCodeLink
	case plugin.DonationLink != "":
		return plugin.DonationLink
	default:
		return fmt.Sprintf("https://www.spigotmc.org/resources/%d", plugin.ID)
	}
}

func ConvertToPlugin(spigotPlugin PluginInfo) paplug.PluginInfo {
	plugin := paplug.PluginInfo{}

	re := regexp.MustCompile("[[:^ascii:]]")
	plugin.Name = re.ReplaceAllLiteralString(spigotPlugin.Name, "")
	plugin.Name = strings.ToLower(plugin.Name)

	plugin.Description = spigotPlugin.Tag
	plugin.Site = getWebsite(spigotPlugin)

	if !strings.HasSuffix(plugin.Description, ".") {
		plugin.Description += "."
	}

	plugin.Install.Type = "simple"

	if spigotPlugin.Contributors == "" {
		plugin.Authors = append(plugin.Authors, spigotPlugin.Resolved.Author.Name)
	} else {
		plugin.Authors = strings.Split(spigotPlugin.Contributors, ", ")
	}

	plugin.Version = spigotPlugin.Resolved.LatestVersion.Name

	// Unknown vars
	plugin.Note = []string{}
	plugin.Dependencies = []string{}
	plugin.OptionalDependencies = []string{}

	if spigotPlugin.SourceCodeLink != "" {
		plugin.License = paplug.Undefined
	} else {
		plugin.License = "proprietary"
	}

	// File & Download
	path := fmt.Sprintf("%s.jar", plugin.Name)

	// File
	file := paplug.File{}
	file.Path = path
	file.Type = "other"

	plugin.Uninstall.Files = append(plugin.Uninstall.Files, file)

	// Download
	download := paplug.Download{}
	download.Type = "url"
	download.Filename = path

	if !spigotPlugin.Premium && spigotPlugin.File.FileType == ".jar" {
		download.URL = fmt.Sprintf("https://api.spiget.org/v2/resources/%d/download", spigotPlugin.ID)
	} else {
		download.URL = paplug.Undefined
	}

	plugin.Downloads = append(plugin.Downloads, download)

	return plugin
}

func Get(name string) PluginInfo {
	var plugins []PluginInfo

	net.Get(
		//nolint:lll
		fmt.Sprintf("https://api.spiget.org/v2/search/resources/%s?field=name&size=1&page=0&sort=-likes&fields=file,contributors,likes,name,tag,sourceCodeLink,donationLink,premium,id,version,author", name),
		fmt.Sprintf("spigot plugin %s not found", name),
		&plugins,
	)

	plugin := plugins[0]

	net.Get(
		fmt.Sprintf("https://api.spiget.org/v2/authors/%d?fields=name", plugin.Author.ID),
		fmt.Sprintf("spigot author %d not found", plugin.Author.ID),
		&plugin.Resolved.Author,
	)

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

// Note: This also converts the plugins to the standard pap format.
func GetManyPluginInfo(names []string) []paplug.PluginInfo {
	infos := []paplug.PluginInfo{}

	for _, name := range names {
		infos = append(infos, GetPluginInfo(name))
	}

	return infos
}
