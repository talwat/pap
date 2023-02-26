package modrinth

// https://api.modrinth.com/v2/project/<project>

// The license of a project.
type License struct {
	ID string
}

// A list of links that could be used as the website.
type Websites struct {
	IssuesURL  string `json:"issues_url"`
	SourceURL  string `json:"source_url"`
	WikiURL    string `json:"wiki_url"`
	DiscordURL string `json:"discord_url"`
}

// Non-important metadata for a modrinth plugin.
type Metadata struct {
	Description string
	License     License
	Websites
}

// The modrinth plugin itself.
// Uses a slug instead of a name.
type PluginInfo struct {
	Slug string

	ResolvedVersion Version
	Versions        []string
	Metadata
}

// https://api.modrinth.com/v2/version/<version>

// A file in a version.
type File struct {
	URL      string
	Filename string
}

// A version, which has a number and a list of files to download.
type Version struct {
	VersionNumber string `json:"version_number"`
	Files         []File
}
