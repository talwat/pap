package plugins

import (
	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/plugins/sources/paplug"
)

// Check if a plugin exists in a list of plugins.
func pluginExists(plugin paplug.PluginInfo, plugins []paplug.PluginInfo) bool {
	for _, pluginToCheck := range plugins {
		// Just check the name which should normally be unique.
		if pluginToCheck.Name == plugin.Name {
			return true
		}
	}

	return false
}

// Recursive function.
// Gets a plugins dependencies and then calls itself to get that dependencies own dependencies.
// This happens until it's done.
func getDependencyLevel(deps []string, dest *[]paplug.PluginInfo, installed []paplug.PluginInfo) {
	depsInfo := GetManyPluginInfo(deps)

	for _, dep := range depsInfo {
		log.Debug("checking if %s already marked for installation...", dep.Name)

		*dest = append(*dest, dep)

		if pluginExists(dep, append(*dest, installed...)) {
			return
		}

		log.Debug("checking if %s has subdependencies...", dep.Name)

		if len(dep.Dependencies) != 0 {
			getDependencyLevel(dep.Dependencies, dest, installed)
		}
	}
}

// Gets the dependencies for one plugin.
// This avoids using dependencies that are already specified for installation.
func getDependencies(deps []string, installed []paplug.PluginInfo) []paplug.PluginInfo {
	finalDeps := []paplug.PluginInfo{}

	getDependencyLevel(deps, &finalDeps, installed)

	return finalDeps
}

// Resolves all of the dependencies for a list of plugins.
func ResolveDependencies(plugins []paplug.PluginInfo) []paplug.PluginInfo {
	deps := []paplug.PluginInfo{}

	if global.NoDepsInput {
		return deps
	}

	log.Log("resolving dependencies...")

	for _, plugin := range plugins {
		deps = append(deps, getDependencies(plugin.Dependencies, plugins)...)

		// Append optional dependencies aswell
		if global.InstallOptionalDepsInput {
			log.Debug("appending optional dependencies: %s...", plugin.OptionalDependencies)

			deps = append(deps, getDependencies(plugin.OptionalDependencies, deps)...)
		}
	}

	return deps
}
