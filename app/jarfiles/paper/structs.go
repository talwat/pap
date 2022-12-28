package paper

type Builds struct {
	ProjectID   string `json:"project_id"`
	ProjectName string `json:"name"`
	Version     string
	Builds      []Build
	Error       string
}

type Build struct {
	Build     int
	Time      string
	Channel   string
	Promoted  bool
	Changes   []Change
	Downloads Downloads
	Error     string
}

type Change struct {
	Commit  string
	Summary string
	Message string
}

type Downloads struct {
	Application    Application
	MojangMappings MojangMappings `json:"mojang-mappings"`
}

type Application struct {
	Name   string
	Sha256 string `json:"sha256"`
}

type MojangMappings struct {
	Name   string
	Sha256 string `json:"sha256"`
}

type Versions struct {
	ProjectID     string   `json:"project_id"`
	ProjectName   string   `json:"project_name"`
	VersionGroups []string `json:"version_groups"`
	Versions      []string
}
