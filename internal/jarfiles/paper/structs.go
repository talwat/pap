package paper

// Note: Uneeded values have been omitted from the original API responses.

type Builds struct {
	Builds []Build
	Error  string
}

type Build struct {
	Build     int
	Time      string
	Channel   string
	Changes   []Change
	Downloads Downloads
	Error     string
}

type Change struct {
	Commit  string
	Summary string
}

type Downloads struct {
	Application Application
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
	Versions []string
}
