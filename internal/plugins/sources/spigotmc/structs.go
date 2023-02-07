package spigotmc

// Structs that are defined by different api routes

// https://api.spiget.org/v2/authors/
type ResolvedAuthor struct {
	Name string
}

// https://api.spiget.org/v2/resources/<resource>/versions/
type ResolvedLatestVersion struct {
	Name string
}
type Resolved struct {
	Author        ResolvedAuthor
	LatestVersion ResolvedLatestVersion
}

// Main struct

type Author struct {
	ID int
}

type File struct {
	FileType string `json:"type"`
	URL      string
}

type Links struct {
	SourceCodeLink string
	DonationLink   string
}

type Version struct {
	ID int
}

// Metadata that isn't used for core operations in pap.
type Metadata struct {
	Contributors string
	Tag          string
	Likes        int
	Author       Author

	Version Version
	Links
}

type PluginInfo struct {
	Name    string
	File    File
	ID      int
	Premium bool

	Metadata
	Resolved Resolved
}
