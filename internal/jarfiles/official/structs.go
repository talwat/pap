package official

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
	ID        string
	Time      string
	Downloads Downloads
}

type Downloads struct {
	Client         Download
	ClientMappings Download `json:"client_mappings"`
	Server         Download
	ServerMappings Download `json:"server_mappings"`
}

type Download struct {
	SHA1 string
	Size int
	URL  string
}
