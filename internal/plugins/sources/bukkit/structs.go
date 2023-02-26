package bukkit

// https://api.curseforge.com/servermods/projects?search=<plugin>
// A bukkit project.
// Called PluginInfo for the sake of consistency.
type PluginInfo struct {
	Slug string
	ID   uint32

	ResolvedFiles []File
}

// https://api.curseforge.com/servermods/files?projectIds=<plugin id>
// A bukkit file.
type File struct {
	DownloadURL string
	FileName    string
	ReleaseType string
}
