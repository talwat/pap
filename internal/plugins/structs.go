package plugins

type OptionalDependency struct {
	Name    string
	Purpose string
}

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

type PluginInfo struct {
	ID                   string
	Name                 string
	Version              string
	Description          string
	License              string
	Authors              []string
	Site                 string
	Dependencies         []string
	Downloads            []Download
	OptionalDependencies []OptionalDependency
	Install              Install
	Uninstall            Uninstall
	Path                 string
	URL                  string
}

type JenkinsArtifacts struct {
	FileName    string
	DisplayName string
}
type JenkinsBuild struct {
	Artifacts []JenkinsArtifacts
}
