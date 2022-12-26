package paper

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/talwat/pap/app/log"
)

const latest = "latest"

func VerifyJarfile(calculated []byte, build Build) {
	log.Log("verifying downloaded jarfile...")

	checksum := hex.EncodeToString(calculated)

	if checksum == build.Downloads.Application.Sha256 {
		log.Log("checksums match!")
	} else {
		log.CustomError(
			fmt.Sprintf("checksums (calculated: %s, proper: %s) don't match!",
				calculated,
				build.Downloads.Application.Sha256,
			),
		)
	}
}

// returns URL to build download, and the build information.
func GetURL(versionInput string, buildID string) (string, Build) {
	var (
		version string
		build   Build
	)

	if versionInput == latest {
		version = GetLatestVersion()
	} else {
		version = versionInput
	}

	build = GetBuild(version, buildID)

	log.Log("using paper version %s", version)

	time, err := time.Parse(time.RFC3339, build.Time)

	log.Error(err, "an error occurred while parsing date supplied by papermc api")

	log.Log("using paper build %d (%s), changes:", build.Build, time.Format("2006-01-02"))

	for _, change := range build.Changes {
		log.RawLog("  (%s) %s\n", change.Commit, change.Summary)
	}

	url := fmt.Sprintf(
		"https://api.papermc.io/v2/projects/paper/versions/%s/builds/%d/downloads/paper-%s-%d.jar",
		version,
		build.Build,
		version,
		build.Build,
	)

	return url, build
}
