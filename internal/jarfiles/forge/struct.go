package forge

import (
	"fmt"
)

// Forge Minecraft versions use a faux semantic versioning scheme (most of the time).
// Major, minor and patch correspond to similar ones in semantic versioning: X.x.p.
// Prerelease versions contain a _preX suffix, such as 1.7.10_pre4.
type MinecraftVersion struct {
	Major int
	Minor int
	Patch int

	IsPrerelease      bool
	PrereleaseVersion int
}

type ByVersion []MinecraftVersion

func (a ByVersion) Len() int { return len(a) }

//nolint:varnamelen
func (a ByVersion) Less(i, j int) bool {
	switch {
	case a[i].Major < a[j].Major:
		return true
	case a[i].Minor < a[j].Minor:
		return true
	case a[i].Patch == a[j].Patch:
		return !a[i].IsPrerelease
	case a[i].Patch < a[j].Patch:
		if a[i].Minor > a[j].Minor {
			return false
		}

		return true
	default:
		return false
	}
}
func (a ByVersion) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (minecraft *MinecraftVersion) String() string {
	builder := fmt.Sprintf("%d.%d", minecraft.Major, minecraft.Minor)

	if minecraft.Patch != 0 {
		builder += fmt.Sprintf(".%d", minecraft.Patch)
	}

	if minecraft.IsPrerelease {
		builder += fmt.Sprintf("_pre%d", minecraft.PrereleaseVersion)
	}

	return builder
}

type InstallerVersion struct {
	Version string
	Type    string
}

type PromotionsSlim struct {
	Promos map[string]string `json:"promos"`
}
