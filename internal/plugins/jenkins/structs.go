package jenkins

type Artifact struct {
	FileName     string `json:"fileName"`
	DisplayName  string `json:"displayPath"`
	RelativePath string `json:"relativePath"`
}

type Build struct {
	Artifacts []Artifact `json:"artifacts"`
}
