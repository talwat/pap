package forge

type MinecraftVersion struct {
	Version string
	Stable  bool
}

type InstallerVersion struct {
	Version string
	Type    string
}

type PromotionsSlim struct {
	Promos map[string]string `json:"promos"`
}