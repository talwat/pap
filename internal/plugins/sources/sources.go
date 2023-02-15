// Utilities for various external sources of plugins (modrinth, spigotmc, and bukkit)
package sources

import (
	"regexp"
	"strings"

	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/plugins/sources/paplug"
)

const Undefined = "unknown"

// Gets info for many plugins & converts them to the standard pap format.
func GetManyPluginInfo(names []string, getInfo func(name string) paplug.PluginInfo) []paplug.PluginInfo {
	infos := []paplug.PluginInfo{}

	for _, name := range names {
		log.Debug("getting info for %s...", name)
		infos = append(infos, getInfo(name))
	}

	return infos
}

// Format a plugin name well enough so that it can be safely handled.
func FormatName(name string) string {
	re := regexp.MustCompile("[[:^ascii:]]")
	newName := re.ReplaceAllLiteralString(name, "")
	newName = strings.ToLower(newName)

	log.Debug("stripped name of ascii characters: %s -> %s", name, newName)

	return newName
}

// Parses the description of a plugin to include a period (.) at the end.
func FormatDesc(desc string) string {
	if strings.HasSuffix(desc, ".") {
		return desc
	}

	log.Debug("description does not have a trailing period (.), appending a period...")

	newDesc := desc
	newDesc += "."

	return newDesc
}
