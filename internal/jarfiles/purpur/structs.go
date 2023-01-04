package purpur

type Versions struct {
	Project  string
	Versions []string
}

type Version struct {
	Builds  BuildsList
	Project string
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
	Duration  int
	MD5       string
	Project   string
	Result    string
	Timestamp int
	Version   string
	Error     string
}

type Commit struct {
	Author      string
	Description string
	Email       string
	Hash        string
	Timestamp   int
}
