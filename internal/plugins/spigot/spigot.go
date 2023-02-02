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
		return paplug.Undefined
	}
}

// TODO: Also convert version.
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
	plugin.Authors = strings.Split(spigotPlugin.Contributors, ", ")

	// Unknown vars
	plugin.Note = []string{}
	plugin.Dependencies = []string{}
	plugin.OptionalDependencies = []string{}
	plugin.Version = paplug.Undefined
	plugin.License = paplug.Undefined

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
	download.URL = paplug.Undefined

	plugin.Downloads = append(plugin.Downloads, download)

	return plugin
}

func Get(name string) PluginInfo {
	var spigotPlugin []PluginInfo

	net.Get(
		//nolint:lll
		fmt.Sprintf("https://api.spiget.org/v2/search/resources/%s?field=name&size=1&page=0&sort=-likes&fields=file%%2Ccontributors%%2Clikes%%2Cname%%2Ctag%%2CsourceCodeLink%%2CdonationLink%%2Cpremium", name),
		fmt.Sprintf("spigot plugin %s not found", name),
		&spigotPlugin,
	)

	return spigotPlugin[0]
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
