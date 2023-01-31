package plugins

// Jenkins.
type JenkinsArtifact struct {
	FileName     string
	DisplayName  string
	RelativePath string
}

type JenkinsBuild struct {
	Artifacts []JenkinsArtifact
}

// Part of PluginInfo.

type Download struct {
	Type     string
	Filename string

	// If type is "jenkins"
	Job      string
	Artifact string

	// If type is "url"
	URL string
}

type Commands struct {
	Windows []string
	Unix    []string
}

type File struct {
	Type string
	Path string
}

type Install struct {
	Type     string
	Commands Commands
}

type Uninstall struct {
	Files []File
}

// Defined in pap, not in the json files themselves.
type DefinedLater struct {
	Path  string
	URL   string
	Alias string
}

// Metadata that isn't used for core operations in pap.
type Metadata struct {
	Description string
	License     string
	Authors     []string
	Site        string
	Note        []string
}

// Operation steps (installing and uninstalling).
type Steps struct {
	Install   Install
	Uninstall Uninstall
}

// All dependencies including optional dependencies.
type AllDependencies struct {
	Dependencies         []string
	OptionalDependencies []string
}

type PluginInfo struct {
	Name    string
	Version string

	Downloads []Download

	Metadata
	AllDependencies
	Steps
	DefinedLater
}
