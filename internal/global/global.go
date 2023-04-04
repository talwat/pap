// Global variables, mainly values set by command line flags which are needed by the whole application.
package global

//nolint:gochecknoglobals
var (
	Version = "0.14.2"

	// Global options.

	AssumeDefaultInput = false
	Debug              = false

	// Downloading Server Jarfiles.

	MinecraftVersionInput              = "latest"
	JarBuildInput                      = "latest"
	PaperExperimentalBuildInput        = false
	UseSnapshotInput                   = false
	FabricExperimentalMinecraftVersion = false
	FabricExperimentalLoaderVersion    = false
	FabricLoaderVersion                = "latest"
	FabricInstallerVersion             = "latest"

	// Script.

	MemoryInput          = "2G"
	AikarsInput          = false
	ScriptUseStdoutInput = false
	JarInput             = ""
	GUIInput             = false

	// Plugin.

	NoDepsInput              = false
	InstallOptionalDepsInput = false
	PluginExperimentalInput  = false

	// Update.

	ReinstallInput = false
)
