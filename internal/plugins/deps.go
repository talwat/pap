package plugins

func pluginExists(plugin PluginInfo, plugins []PluginInfo) bool {
	for _, pluginToCheck := range plugins {
		if pluginToCheck.Name == plugin.Name {
			return true
		}
	}

	return false
}

// Recursive function.
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
