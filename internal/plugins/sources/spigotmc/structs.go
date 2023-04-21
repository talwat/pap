package spigotmc

// https://api.spiget.org/v2/authors/
// The author of a plugin, only the name is needed.
type ResolvedAuthor struct {
	Name string
}

// https://api.spiget.org/v2/resources/<resource>/versions/
// The latest version of a plugin.
type ResolvedLatestVersion struct {
	Name string
}

// Resolved information, like the author and latest version.
// Resolved meaning it is in a separate endpoint.
type Resolved struct {
	Author        ResolvedAuthor
	LatestVersion ResolvedLatestVersion
}

// Main struct

// The author of a plugin, used if contributors is undefined.
type Author struct {
	ID int
}

// A file provided by the plugin.
type File struct {
	FileType string `json:"type"`
	URL      string
}

// Websites that are used to get a plugins website.
type Links struct {
	SourceCodeLink string
	DonationLink   string
}

// A version with an ID, pretty basic.
type Version struct {
	ID uint32
}

// Non-important metadata for a plugin.
type Metadata struct {
	Contributors string
	Tag          string
	Likes        int
	Author       Author

	Links
}

// Information used to check if a plugin is able to be downloaded.
type DownloadInfo struct {
	Premium bool
	File    File
}

// The plugin information itself.
type PluginInfo struct {
	Name    string
	Version Version
	ID      uint32

	Metadata
	DownloadInfo
	Resolved Resolved
}
