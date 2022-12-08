package main

type PaperBuildsStruct struct {
	ProjectID   string `json:"project_id"`
	ProjectName string `json:"name"`
	Version     string
	Builds      []PaperBuildStruct
	Error       string
}

type PaperBuildStruct struct {
	Build     int
	Time      string
	Channel   string
	Promoted  bool
	Changes   []ChangeStruct
	Downloads DownloadsStruct
	Error     string
}

type ChangeStruct struct {
	Commit  string
	Summary string
	Message string
}

type DownloadsStruct struct {
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

type PaperVersions struct {
	ProjectID     string   `json:"project_id"`
	ProjectName   string   `json:"project_name"`
	VersionGroups []string `json:"version_groups"`
	Versions      []string
}
