package paplug

type Download struct {
	Type     string `json:"type"`
	Filename string `json:"filename"`

	// If type is "jenkins"

	Job      string `json:"job,omitempty"`
	Artifact string `json:"artifact,omitempty"`

	// If type is "url"

	URL string `json:"url,omitempty"`
}

type Commands struct {
	Windows []string `json:"windows"`
	Unix    []string `json:"unix"`
}

type File struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type Install struct {
	Type     string    `json:"type"`
	Commands *Commands `json:"commands,omitempty"`
}

type Uninstall struct {
	Files []File `json:"files"`
}

// Defined in pap, not in the json files themselves.
type DefinedLater struct {
	Path   string `json:"path,omitempty"`
	URL    string `json:"url,omitempty"`
	Source string `json:"source,omitempty"`
}

// Metadata that isn't used for core operations in pap.
type Metadata struct {
	Description string   `json:"description"`
	License     string   `json:"license"`
	Authors     []string `json:"authors"`
	Site        string   `json:"site,omitempty"`
	Note        []string `json:"note,omitempty"`
}

// Operation steps (installing and uninstalling).
type Steps struct {
	Install   Install   `json:"install"`
	Uninstall Uninstall `json:"uninstall"`
}

// All dependencies including optional dependencies.
type AllDependencies struct {
	Dependencies         []string `json:"dependencies,omitempty"`
	OptionalDependencies []string `json:"optional_dependencies,omitempty"`
}

type PluginInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Metadata
	AllDependencies

	Downloads []Download `json:"downloads"`
	Steps
	DefinedLater

	Alias string `json:"alias,omitempty"`
}
