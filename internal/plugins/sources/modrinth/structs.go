package modrinth

// /project

type License struct {
	ID string
}

type Websites struct {
	IssuesURL  string `json:"issues_url"`
	SourceURL  string `json:"source_url"`
	WikiURL    string `json:"wiki_url"`
	DiscordURL string `json:"discord_url"`
}

type PluginInfo struct {
	Slug        string
	Description string

	License License
	Websites

	Versions        []string
	ResolvedVersion Version
}

// /version

type File struct {
	URL      string
	Filename string
}

type Version struct {
	VersionNumber string `json:"version_number"`
	Files         []File
}
