package fabric

type MinecraftVersion struct {
	Version string
	Stable  bool
}

type MinecraftVersions struct {
	Game []MinecraftVersion
}

type LoaderVersion struct {
	Separator string
	Version   string
	Build     uint16
	Stable    bool
}

type InstallerVersion struct {
	Stable  bool
	Version string
}

type Versions interface {
	~[]struct{ Stable bool }
}
