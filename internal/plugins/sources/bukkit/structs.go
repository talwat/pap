package bukkit

// https://api.curseforge.com/servermods/projects?search=<plugin>

type Project struct {
	Slug string
	ID   int

	ResolvedFiles []File
}

// https://api.curseforge.com/servermods/files?projectIds=<plugin id>

type File struct {
	DownloadURL string
	FileName    string
	ReleaseType string
}