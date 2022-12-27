// Global variables, mainly values set by command line flags which are needed by the whole application.
package global

//nolint:gochecknoglobals
var (
	AssumeDefaultInput     = false
	VersionInput           = "latest"
	BuildInput             = "latest"
	ExperimentalBuildInput = false
	NoFloodGateInput       = false
	MemoryInput            = "2G"
	AikarsInput            = false
	JarInput               = "paper.jar"
	GUIInput               = false
)
