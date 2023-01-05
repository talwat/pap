package official

// Note: Uneeded values have been omitted from the original API responses.

type Versions struct {
	Latest   Latest
	Versions []Version
}

type Latest struct {
	Release  string
	Snapshot string
}

type Version struct {
	ID          string
	VersionType string `json:"type"`
	URL         string
	Time        string
	ReleaseTime string `json:"releaseTime"`
}

// Note: Values have been omitted from this struct from the original mojang API response.
type Package struct {
	ID          string
	Time        string
	ReleaseTime string `json:"releaseTime"`
	Downloads   Downloads
}

type Downloads struct {
	Server Download
}

type Download struct {
	SHA1 string
	URL  string
}
