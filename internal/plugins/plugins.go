package plugins

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/talwat/pap/internal/exec"
	"github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/plugins/sources/paplug"
)

// Substitutes parts of a string like {version} with their proper counterpart.
func SubstituteProps(plugin paplug.PluginInfo, str string) string {
	toReplace := map[string]string{
		"version": plugin.Version,
		"name":    plugin.Name,
	}

	final := str

	for key, value := range toReplace {
		log.Debug("substituting %s with %s", key, value)
		final = strings.ReplaceAll(final, fmt.Sprintf("{%s}", key), value)
	}

	return final
}

func PluginInstall(plugin paplug.PluginInfo) {
	name := plugin.Name

	log.Log("installing %s...", name)

	log.Log("making plugins directory...")
	fs.MakeDirectory("plugins")

	log.Log("checking if plugin is already installed...")

	for _, file := range plugin.Uninstall.Files {
		if file.Type != "main" || !fs.FileExists(filepath.Join("plugins", file.Path)) {
			continue
		}

		log.Warn("%s may already be installed. if it is not installed, try uninstalling it first and then reinstalling", name)
		log.Warn("skipping %s...", name)

		return
	}

	PluginDownload(plugin)

	if plugin.Install.Type == "simple" {
		log.Success("successfully installed %s (simple)", name)

		return
	}

	log.Log("running commands for %s...", name)

	var cmds []string

	if runtime.GOOS == "windows" {
		log.Debug("using windows commands...")

		cmds = plugin.Install.Commands.Windows
	} else {
		log.Debug("using unix commands...")

		cmds = plugin.Install.Commands.Unix
	}

	for _, cmd := range cmds {
		exec.Run("plugins", SubstituteProps(plugin, cmd))
		log.RawLog("\n")
	}

	log.Success("successfully installed %s", name)
}

func PluginUninstall(plugin paplug.PluginInfo) {
	name := plugin.Name

	log.Log("uninstalling %s...", name)

	for _, file := range plugin.Uninstall.Files {
		path := filepath.Join("plugins", SubstituteProps(plugin, file.Path))

		if file.Type == "" {
			file.Type = "other"
		}

		log.Log("deleting %s at %s", file.Type, path)
		fs.DeletePath(path)
	}

	log.Success("successfully uninstalled %s", name)
}

func PluginDoMany(plugins []paplug.PluginInfo, operation func(plugin paplug.PluginInfo)) {
	for _, plugin := range plugins {
		operation(plugin)
	}
}
