package forge

import (
	"strings"

	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
)

func loadPromotions() PromotionsSlim {
	var promotions PromotionsSlim
	net.Get(
		"https://files.minecraftforge.net/maven/net/minecraftforge/forge/promotions_slim.json",
		"could not retrieve promotions",
		&promotions,
	)

	return promotions
}

func getPromotion(promos *PromotionsSlim, ver string) (MinecraftVersion, InstallerVersion) {
	if ver == "latest" {
		var mvs []MinecraftVersion
		var ivs []InstallerVersion

		for mp, ip := range promos.Promos {
			mvParts := strings.Split(mp, "-")

			mvs = append(mvs, MinecraftVersion{
				Version: mvParts[0],
				Stable: !strings.Contains(mvParts[0], "pre"),
			})

			ivs = append(ivs, InstallerVersion{
				Version: ip,
				Type: mvParts[1],
			})
		}

		return mvs[len(mvs)-1], ivs[len(ivs)-1]
	}

	mp, found := promos.Promos["ver"]
	if !found {
		log.RawError("version %s not found in forge promotions", ver)
	}

	mpParts := strings.Split(mp, "-")
	
	mv := MinecraftVersion {
		Version: mpParts[0],
		Stable: !strings.Contains(mpParts[0], "pre"),
	}
	iv := InstallerVersion {
		Version: promos.Promos["ver"],
		Type: mpParts[1],
	}

	return mv, iv
} 

func getLatestVersion([]MinecraftVersion) MinecraftVersion {
	promotions := loadPromotions()
	mv, _ := getPromotion(&promotions, "latest")

	return mv
}

func getLatestInstaller([]InstallerVersion) InstallerVersion {
	promotions := loadPromotions()
	_, iv := getPromotion(&promotions, "latest")

	return iv
}

func GetMinecraftVersions() []MinecraftVersion {
	return nil
}

func GetInstallerVersions() []InstallerVersion {
	return nil
}
