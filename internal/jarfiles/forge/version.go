package forge

import (
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/talwat/pap/internal/log"
	"golang.org/x/exp/maps"
)

var (
	preRegex  = regexp.MustCompile(`_pre[0-9]`)
	typeRegex = regexp.MustCompile(`-[^"]*`)
)

func cleanMinecraftVersionString(ver string, mv *MinecraftVersion) string {
	pver := preRegex.FindString(ver)
	if pver != "" {
		mv.IsPrerelease = true
		pver = strings.Replace(pver, "_pre", "", 1)

		var err error
		mv.PrereleaseVersion, err = strconv.Atoi(pver)
		log.Error(err, "failed to parse prerelease version number")

		ver = preRegex.ReplaceAllString(ver, "")
	}

	ver = typeRegex.ReplaceAllString(ver, "")

	return ver
}

func parseMinecraftVersion(ver string) MinecraftVersion {
	var mv MinecraftVersion

	cmv := cleanMinecraftVersionString(ver, &mv)
	smv := strings.Split(cmv, ".")

	var err error
	mv.Major, err = strconv.Atoi(smv[0])
	log.Error(err, "failed to parse major version")

	mv.Minor, err = strconv.Atoi(smv[1])
	log.Error(err, "failed to parse minor version")

	if len(smv) > 2 {
		mv.Patch, err = strconv.Atoi(smv[2])
		log.Error(err, "failed to parse minor version")
	}

	return mv
}

func getLatestMinecraftVersion(promotions *PromotionsSlim) MinecraftVersion {
	svers := maps.Keys(promotions.Promos)

	check := make(map[string]bool, 0)

	for _, val := range svers {
		s := strings.Split(val, "-")[0]
		check[s] = true
	}

	mvers := make([]MinecraftVersion, len(svers))

	i := 0

	for ver := range check {
		i++
		mvers[i] = parseMinecraftVersion(ver)
	}

	sort.Stable(ByVersion(mvers))

	return mvers[len(mvers)-1]
}
