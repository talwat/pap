package plugins

import "github.com/talwat/pap/internal/plugins/paplug"

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
		if pluginExists(dep, append(*dest, installed...)) {
			return
		}

		*dest = append(*dest, dep)

		if len(dep.Dependencies) > 0 {
			getDependencyLevel(dep.Dependencies, dest, installed)
		}
	}
}

func GetDependencies(deps []string, installed []paplug.PluginInfo) []paplug.PluginInfo {
	finalDeps := []paplug.PluginInfo{}

	getDependencyLevel(deps, &finalDeps, installed)

	return finalDeps
}
