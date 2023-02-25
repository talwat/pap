// Global variables, mainly values set by command line flags which are needed by the whole application.
package global

//nolint:gochecknoglobals
var (
	Version = "0.13.0"

	// Global options.
	AssumeDefaultInput = false
	Debug              = false

	// Downloading Server Jarfiles.
	MinecraftVersionInput       = "latest"
	JarBuildInput               = "latest"
	PaperExperimentalBuildInput = false
	OfficialUseSnapshotInput    = false

	// Script.
	MemoryInput          = "2G"
	AikarsInput          = false
	ScriptUseStdoutInput = false
	JarInput             = "paper.jar"
	GUIInput             = false

	// Plugin.
	NoDepsInput              = false
	InstallOptionalDepsInput = false

	// Update.
	ReinstallInput = false
)
