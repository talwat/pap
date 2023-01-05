package purpur

// Note: Uneeded values have been omitted from the original API responses.

type Versions struct {
	Versions []string
}

type Version struct {
	Builds  BuildsList
	Version string
	Error   string
}

type BuildsList struct {
	All    []string
	Latest string
}

type Build struct {
	Build     string
	Commits   []Commit
	MD5       string
	Timestamp int
	Error     string
}

type Commit struct {
	Description string
	Hash        string
	Timestamp   int
}
