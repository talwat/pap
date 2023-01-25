package plugins

import (
	"fmt"
	"runtime"

	"github.com/talwat/pap/internal/exec"
	"github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/log"
)

func PluginInstall(plugin PluginInfo) {
	name := plugin.Name

	log.Log("installing %s...", name)

	fs.MakeDirectory("plugins")
	PluginDownload(plugin)

	if plugin.Install.Type == "simple" {
		return
	}

	log.Log("running commands for %s...", name)

	var cmds []string

	if runtime.GOOS == "windows" {
		cmds = plugin.Install.Commands.Windows
	} else {
		cmds = plugin.Install.Commands.Unix
	}

	for _, cmd := range cmds {
		exec.Run("plugins", cmd)
		log.RawLog("\n")
	}

	log.Log("successfully installed %s", name)
}

func PluginUninstall(plugin PluginInfo) {
	name := plugin.Name

	log.Log("uninstalling %s...", name)

	for _, file := range plugin.Uninstall.Files {
		path := fmt.Sprintf("plugins/%s", file.Path)

		if file.Type == "" {
			file.Type = "other"
		}

		log.Log("deleting %s at %s", file.Type, path)
		fs.DeleteFile(path)
	}

	log.Log("successfully uninstalled %s", name)
}

func PluginDoMany(plugins []PluginInfo, operation func(plugin PluginInfo)) {
	for _, plugin := range plugins {
		operation(plugin)
	}
}
