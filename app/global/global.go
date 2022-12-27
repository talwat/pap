// Global variables, mainly values set by command line flags which are needed by the whole application.
package global

//nolint:gochecknoglobals
var (
	AssumeDefaultInput     = false
	VersionInput           = "latest"
	BuildInput             = "latest"
	ExperimentalBuildInput = false
	NoFloodGateInput       = false
	XMSInput               = "2G"
	XMXInput               = "2G"
	JarInput               = "paper.jar"
	GUIInput               = false
)
