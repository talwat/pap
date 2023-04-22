package forge

import "fmt"

type MinecraftVersion struct {
	Major int
	Minor int
	Patch int

	IsPrerelease      bool
	PrereleaseVersion int
}

func (jver *MinecraftVersion) GreaterThan(iver *MinecraftVersion) bool {
	if jver.Major > iver.Major {
		return true
	}

	if jver.Minor > iver.Minor {
		return true
	}

	if jver.Patch >= iver.Patch {
		return !jver.IsPrerelease
	}

	return false
}

func (mver *MinecraftVersion) String() string {
	mvs := fmt.Sprintf("%d.%d", mver.Major, mver.Minor)

	if mver.Patch != 0 {
		mvs += fmt.Sprintf(".%d", mver.Patch)
	}

	if mver.IsPrerelease {
		mvs += fmt.Sprintf("_pre%d", mver.PrereleaseVersion)
	}

	return mvs
}

type InstallerVersion struct {
	Version string
	Type    string
}

type PromotionsSlim struct {
	Promos map[string]string `json:"promos"`
}
