package purpur

import (
	"fmt"
	"strings"
	"time"

	"github.com/talwat/pap/internal/log"
)

func formatURL(version string, build Build) string {
	return fmt.Sprintf(
		"https://api.purpurmc.org/v2/purpur/%s/%s/download",
		version,
		build.Build,
	)
}

// Returns URL to build download, and the build information.
func GetURL(versionInput string, buildID string) (string, Build) {
	version := GetVersion(versionInput)
	build := GetBuild(version, buildID)

	log.Log("using purpur version %s", version.Version)

	time := time.UnixMilli(int64(build.Timestamp))
	formattedTime := time.Format("2006-01-02")

	log.Debug("raw timestamp: %d", build.Timestamp)
	log.Debug("unix time: %s", time)
	log.Debug("formatted time: %s", formattedTime)

	log.Log("using purpur build %s (%s), commits:", build.Build, formattedTime)

	for _, commit := range build.Commits {
		log.RawLog("  (%s) %s\n", commit.Hash, strings.Split(commit.Description, "\n")[0])
	}

	return formatURL(version.Version, build), build
}
