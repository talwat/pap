// Global variables, mainly values set by command line flags which are needed by the whole application.
package global

//nolint:gochecknoglobals
var (
	Version = "0.11.0"

	// CLI Options.
	AssumeDefaultInput          = false
	MinecraftVersionInput       = "latest"
	JarBuildInput               = "latest"
	PaperExperimentalBuildInput = false
	OfficialUseSnapshotInput    = false
	NoFloodGateInput            = false
	MemoryInput                 = "2G"
	AikarsInput                 = false
	ScriptUseStdoutInput        = false
	JarInput                    = "paper.jar"
	GUIInput                    = false
)
