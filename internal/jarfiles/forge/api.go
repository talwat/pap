package forge

import (
	"fmt"

	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
)

func getPromotions() PromotionsSlim {
	var promotions PromotionsSlim
	net.Get(
		"https://files.minecraftforge.net/maven/net/minecraftforge/forge/promotions_slim.json",
		"could not retrieve promotions",
		&promotions,
	)

	return promotions
}

func getInstaller(mver string, useLatestInstaller bool) (MinecraftVersion, InstallerVersion) {
	promos := getPromotions()

	var mv MinecraftVersion
	var iv InstallerVersion

	if mver == "latest" {
		mv = getLatestMinecraftVersion(&promos)
	} else {
		mv = parseMinecraftVersion(mver)
	}

	if useLatestInstaller {
		iv = getVersion(&promos, &mv, "latest")
		goto ret
	}

	iv = getVersion(&promos, &mv, "recommended")
	if (iv == InstallerVersion{}) {
		log.Continue("no recommended installer found for version %s. use the latest version?", mver)
	}

	iv = getVersion(&promos, &mv, "latest")

ret:
	if (iv == InstallerVersion{}) {
		log.RawError("could not get a valid installer version")
	}

	return mv, iv
}

func getSpecificInstaller(iver string) InstallerVersion {
	return InstallerVersion{
		Version: iver,
	}
}

func getVersion(promos *PromotionsSlim, mv *MinecraftVersion, t string) InstallerVersion {
	promo, found := promos.Promos[fmt.Sprintf("%s-%s", mv.String(), t)]
	if found {
		return InstallerVersion{
			Version: promo,
			Type:    t,
		}
	}

	return InstallerVersion{}
}
