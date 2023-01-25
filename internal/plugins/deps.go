package plugins

func dependencyExists(dep PluginInfo, deps []PluginInfo) bool {
	for _, depToCheck := range deps {
		if depToCheck.Name == dep.Name {
			return true
		}
	}

	return false
}

// Recursive function.
func getDependencyLevel(deps []string, dest *[]PluginInfo) {
	depsInfo := GetManyPluginInfo(deps)

	for _, dep := range depsInfo {
		if !dependencyExists(dep, *dest) {
			*dest = append(*dest, dep)
		}

		if len(dep.Dependencies) > 0 {
			getDependencyLevel(dep.Dependencies, dest)
		}
	}
}

func GetDependencies(plugin PluginInfo) []PluginInfo {
	finalDeps := []PluginInfo{}

	getDependencyLevel(plugin.Dependencies, &finalDeps)

	return finalDeps
}
