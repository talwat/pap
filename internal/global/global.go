// Global variables, mainly values set by command line flags which are needed by the whole application.
package global

//nolint:gochecknoglobals
var (
	// Global.
	Version = "0.11.0"

	// CLI Options.
	AssumeDefaultInput = false

	// Downloading Server Jarfiles.
	MinecraftVersionInput       = "latest"
	JarBuildInput               = "latest"
	PaperExperimentalBuildInput = false
	OfficialUseSnapshotInput    = false

	// Geyser.
	NoFloodGateInput = false

	// Script.
	MemoryInput          = "2G"
	AikarsInput          = false
	ScriptUseStdoutInput = false
	JarInput             = "paper.jar"
	GUIInput             = false

	// Plugin.
	NoDepsInput              = false
	InstallOptionalDepsInput = false
)
