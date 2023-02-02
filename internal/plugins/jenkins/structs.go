package jenkins

type Artifact struct {
	FileName     string `json:"file_name"`
	DisplayName  string `json:"display_name"`
	RelativePath string `json:"relative_path"`
}

type Build struct {
	Artifacts []Artifact `json:"artifacts"`
}
