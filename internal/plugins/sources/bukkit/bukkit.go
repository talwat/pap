package bukkit

import (
	"fmt"

	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
	"github.com/talwat/pap/internal/plugins/sources"
	"github.com/talwat/pap/internal/plugins/sources/paplug"
)

func ConvertToPlugin(bukkitProject Project) paplug.PluginInfo {
	plugin := paplug.PluginInfo{}
	plugin.Name = sources.FormatName(bukkitProject.Slug)
	plugin.Site = fmt.Sprintf("https://dev.bukkit.org/projects/%s", plugin.Name)

	plugin.Install.Type = "simple"

	// Unknown vars
	plugin.Note = []string{}
	plugin.Dependencies = []string{}
	plugin.OptionalDependencies = []string{}
	plugin.Version = sources.Undefined
	plugin.Description = sources.Undefined
	plugin.License = sources.Undefined
	plugin.Description = sources.Undefined

	// File & Download
	// TODO: Use latest release, because currently this just uses the latest regardless of whether it is stable.
	latestFile := bukkitProject.ResolvedFiles[len(bukkitProject.ResolvedFiles)-1]

	// File
	file := paplug.File{}
	file.Path = latestFile.FileName
	file.Type = "other"

	plugin.Uninstall.Files = append(plugin.Uninstall.Files, file)

	// Download
	download := paplug.Download{}
	download.URL = latestFile.DownloadURL
	download.Type = "url"
	download.Filename = latestFile.FileName

	plugin.Downloads = append(plugin.Downloads, download)

	return plugin
}

func Get(name string) Project {
	var projects []Project

	net.Get(
		fmt.Sprintf("https://api.curseforge.com/servermods/projects?search=%s", name),
		fmt.Sprintf("bukkitdev search %s not found", name),
		&projects,
	)

	if len(projects) == 0 {
		log.RawError("bukkitdev plugin %s not found", name)
	}

	project := projects[0]

	net.Get(
		fmt.Sprintf("https://api.curseforge.com/servermods/files?projectIds=%d", project.ID),
		fmt.Sprintf("bukkitdev versions for %s not found", project.Slug),
		&project.ResolvedFiles,
	)

	return project
}

// Gets & converts to the standard pap format.
func GetPluginInfo(name string) paplug.PluginInfo {
	return ConvertToPlugin(Get(name))
}
