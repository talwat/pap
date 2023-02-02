package spigot

type File struct {
	FileType string `json:"type"`
	URL      string
}

type Links struct {
	SourceCodeLink string
	DonationLink   string
}

type PluginInfo struct {
	Name         string
	File         File
	Tag          string
	ID           int
	Contributors string
	Likes        int
	Premium      bool

	Links
}
