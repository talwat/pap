package purpur

// Note: Uneeded values have been omitted from the original API responses.

type Errorable struct {
	Error string
}
type Versions struct {
	Versions []string
}

type Version struct {
	Builds  BuildsList
	Version string

	Errorable
}

type BuildsList struct {
	All    []string
	Latest string
}

type Commit struct {
	Description string
	Hash        string
	Timestamp   uint64
}

type BuildMetadata struct {
	Timestamp int
	Commits   []Commit
}

type Build struct {
	Build string
	MD5   string

	Errorable
	BuildMetadata
}
