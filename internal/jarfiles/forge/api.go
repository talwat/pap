package forge

import (
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

func getLatestPromotion(promos *PromotionsSlim, ver string) InstallerVersion {
	p, found := promos.Promos[ver+"-latest"]
	if !found {
		log.RawError("version %s does not have a latest installer", ver)
	}

	iv := InstallerVersion{
		Version: p,
		Type: "latest",
	}
	
	return iv
}

func getRecommendedPromotion(promos *PromotionsSlim, ver string) InstallerVersion {
	p, found := promos.Promos[ver+"-recommended"]
	if !found {
		log.RawError("version %s does not have a recommended installer", ver)
	}

	iv := InstallerVersion {
		Version: p,
		Type: "recommended",
	}

	return iv
} 

func GetLatestInstaller(ver string) InstallerVersion {
	promotions := loadPromotions()
	iv := getLatestPromotion(&promotions, ver)

	return iv
}

func GetRecommendedInstaller(ver string) InstallerVersion {
	promotions := loadPromotions()
	iv := getRecommendedPromotion(&promotions, ver)

	return iv
}
