package paplug

// If type is "jenkins".
type JenkinsDownload struct {
	Job      string `json:"job,omitempty"`
	Artifact string `json:"artifact,omitempty"`
}

// If type is "url".
type URLDownload struct {
	URL string `json:"url,omitempty"`
}

type Download struct {
	Type     string `json:"type"`
	Filename string `json:"filename"`

	JenkinsDownload
	URLDownload
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
	Authors     []string `json:"authors"`
	Note        []string `json:"note,omitempty"`

	LessImportantMetadata
}

// Metadata that is not very important.
type LessImportantMetadata struct {
	License string `json:"license"`
	Site    string `json:"site,omitempty"`
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
