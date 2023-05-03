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

func cleanMinecraftVersionString(version string, minecraft *MinecraftVersion) string {
	preVersion := preRegex.FindString(version)
	if preVersion != "" {
		minecraft.IsPrerelease = true
		preVersion = strings.Replace(preVersion, "_pre", "", 1)

		var err error
		minecraft.PrereleaseVersion, err = strconv.Atoi(preVersion)
		log.Error(err, "failed to parse prerelease version number")

		version = preRegex.ReplaceAllString(version, "")
	}

	version = typeRegex.ReplaceAllString(version, "")

	return version
}

func parseMinecraftVersion(ver string) MinecraftVersion {
	var minecraft MinecraftVersion

	cleanVersion := cleanMinecraftVersionString(ver, &minecraft)
	splitVersion := strings.Split(cleanVersion, ".")

	var err error
	minecraft.Major, err = strconv.Atoi(splitVersion[0])
	log.Error(err, "failed to parse major version")

	minecraft.Minor, err = strconv.Atoi(splitVersion[1])
	log.Error(err, "failed to parse minor version")

	// to avoid magic numbers
	lenMajorAndMinor := 2

	if len(splitVersion) > lenMajorAndMinor {
		minecraft.Patch, err = strconv.Atoi(splitVersion[2])
		log.Error(err, "failed to parse minor version")
	}

	return minecraft
}

func getLatestMinecraftVersion(promotions *PromotionsSlim) MinecraftVersion {
	promoKeys := maps.Keys(promotions.Promos)

	keymap := make(map[string]bool, len(promoKeys))

	for _, val := range promoKeys {
		s := strings.Split(val, "-")[0]
		_, exists := keymap[s]

		if !exists {
			keymap[s] = true
		}
	}

	keys := maps.Keys(keymap)
	minecraftVersions := make([]MinecraftVersion, len(keys))

	for i := 0; i < len(keys); i++ {
		minecraftVersions[i] = parseMinecraftVersion(keys[i])
	}

	sort.Sort(ByVersion(minecraftVersions))

	last := minecraftVersions[len(minecraftVersions)-1]

	return last
}
