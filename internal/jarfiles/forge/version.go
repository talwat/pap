package forge

import (
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/talwat/pap/internal/log"
	"golang.org/x/exp/maps"
)

var preRegex = regexp.MustCompile(`_pre[0-9]`)
var typeRegex = regexp.MustCompile(`-[^"]*`)

func cleanMinecraftVersionString(ver string, mv *MinecraftVersion) string {
	var err error

	pver := preRegex.FindString(ver)
	if pver != "" {
		mv.IsPrerelease = true
		pver = strings.Replace(pver, "_pre", "", 1)

		mv.PrereleaseVersion, err = strconv.Atoi(pver)
		log.Error(err, "failed to parse prerelease version number")

		ver = preRegex.ReplaceAllString(ver, "")
	}

	ver = typeRegex.ReplaceAllString(ver, "")

	return ver
}

func parseMinecraftVersion(ver string) MinecraftVersion {
	var mv MinecraftVersion
	var smv []string
	var err error

	cmv := cleanMinecraftVersionString(ver, &mv)
	smv = strings.Split(cmv, ".")

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

func sortMinecraftVersionStrings(mvers []string) {
	sort.Slice(mvers, func(i int, j int) bool {
		iv := parseMinecraftVersion(mvers[i])
		jv := parseMinecraftVersion(mvers[j])

		return jv.GreaterThan(&iv)
	})
}

func getLatestMinecraftVersion(promotions *PromotionsSlim) MinecraftVersion {
	mvers := maps.Keys(promotions.Promos)
	sortMinecraftVersionStrings(mvers)

	parsed := parseMinecraftVersion(mvers[len(mvers)-1])

	return parsed
}
