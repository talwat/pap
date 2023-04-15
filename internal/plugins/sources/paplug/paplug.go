// The pap plugin manager format
package paplug

// Check if a plugin exists in a list of plugins.
func PluginExists(plugin PluginInfo, plugins []PluginInfo) bool {
	for _, pluginToCheck := range plugins {
		// Just check the name which should normally be unique.
		if pluginToCheck.Name == plugin.Name {
			return true
		}
	}

	return false
}
