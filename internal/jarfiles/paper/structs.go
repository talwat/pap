package paper

// Note: Uneeded values have been omitted from the original API responses.

// Structs that are set directly to the API response and may have an 'error' attribute.
type Errorable struct {
	Error string
}

type Builds struct {
	Builds []Build
	Errorable
}

type BuildMetadata struct {
	Time    string
	Changes []Change
}

type Build struct {
	Build     int
	Channel   string
	Downloads Downloads

	Errorable
	BuildMetadata
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

type Versions struct {
	Versions []string
}
