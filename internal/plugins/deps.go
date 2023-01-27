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
func getDependencyLevel(deps []string, dest *[]PluginInfo, plugins []PluginInfo) {
	depsInfo := GetManyPluginInfo(deps)

	for _, dep := range depsInfo {
		if pluginExists(dep, append(*dest, plugins...)) {
			return
		}

		*dest = append(*dest, dep)

		if len(dep.Dependencies) > 0 {
			getDependencyLevel(dep.Dependencies, dest, plugins)
		}
	}
}

func GetDependencies(plugin PluginInfo, plugins []PluginInfo) []PluginInfo {
	finalDeps := []PluginInfo{}

	getDependencyLevel(plugin.Dependencies, &finalDeps, plugins)

	return finalDeps
}
