// Utilities for various external sources of plugins (modrinth, spigotmc, and bukkit)
package sources

import (
	"regexp"
	"strings"

	"github.com/talwat/pap/internal/plugins/sources/paplug"
)

const Undefined = "unknown"

// Gets info for many plugins & converts them to the standard pap format.
func GetManyPluginInfo(names []string, getInfo func(name string) paplug.PluginInfo) []paplug.PluginInfo {
	infos := []paplug.PluginInfo{}

	for _, name := range names {
		infos = append(infos, getInfo(name))
	}

	return infos
}

// Parses a plugin name well enough so that it can be safely handled.
func ParseName(name string) string {
	re := regexp.MustCompile("[[:^ascii:]]")
	newName := re.ReplaceAllLiteralString(name, "")
	newName = strings.ToLower(newName)

	return newName
}
