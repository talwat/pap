package plugins

// Check if a plugin exists in a list of plugins.
func pluginExists(plugin PluginInfo, plugins []PluginInfo) bool {
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
func getDependencyLevel(deps []string, dest *[]PluginInfo, installed []PluginInfo) {
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

func GetDependencies(deps []string, installed []PluginInfo) []PluginInfo {
	finalDeps := []PluginInfo{}

	getDependencyLevel(deps, &finalDeps, installed)

	return finalDeps
}
