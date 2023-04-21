package forge

import (
	"sort"
	"strings"

	"github.com/hashicorp/go-version"
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

func GetLatestInstaller(ver string) InstallerVersion {
	promotions := getPromotions()
	p, found := promotions.Promos[ver+"-latest"]
	if !found {
		log.RawError("version %s does not have a latest installer", ver)
	}

	iv := InstallerVersion{
		Version: p,
		Type:    "latest",
	}

	return iv
}

func GetInstaller(ver string) InstallerVersion {
	promotions := getPromotions()

	promo, found := promotions.Promos[ver+"-recommended"]
	if found {
		return InstallerVersion{
			Version: promo,
			Type:    "recommended",
		}
	}

	promo, found = promotions.Promos[ver+"-latest"]
	if found {
		return InstallerVersion{
			Version: promo,
			Type:    "latest",
		}
	}

	log.RawError("no installer found for version %s", ver)
	return InstallerVersion{}
}

func GetLatest() (string, string) {
	promotions := getPromotions()

	var mvs []string
	for k := range promotions.Promos {
		splitVersion := strings.Split(k, "-")
		mvs = append(mvs, splitVersion[0])
	}

	sort.Slice(mvs, func(i, j int) bool {
		is := strings.Replace(mvs[i], "_pre", "0", -1)
		js := strings.Replace(mvs[j], "_pre", "0", -1)

		iv, err := version.NewVersion(is)
		log.Error(err, "could not parse version "+mvs[i])

		jv, err := version.NewVersion(js)
		log.Error(err, "could not parse version "+mvs[j])

		return jv.GreaterThan(iv)
	})

	lm := mvs[len(mvs)-1]
	iv := promotions.Promos[lm+"-latest"]

	if lm == "" || iv == "" {
		log.Log(lm + " " + iv)
		log.RawError("failed to get latest forge installer")
	}

	return lm, iv
}
