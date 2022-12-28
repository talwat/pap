// Interact with papermc downloads api and verify of downloaded files.
package paper

// This is the only file which is accessed from other packages.
// If you would like to add compatibility for other jarfile types, you need:
// - A GetURL(version string, build string) function.
// - A function to verify the jarfile.
// - And most likely a struct for each build as well as all of the builds.

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/talwat/pap/app/log"
)

const latest = "latest"

func VerifyJarfile(calculated []byte, build Build) {
	log.Log("verifying downloaded jarfile...")

	if checksum := hex.EncodeToString(calculated); checksum == build.Downloads.Application.Sha256 {
		log.Log("checksums match!")
	} else {
		log.RawError(
			fmt.Sprintf("checksums (calculated: %s, proper: %s) don't match!",
				checksum,
				build.Downloads.Application.Sha256,
			),
		)
	}
}

func formatURL(version string, build Build) string {
	return fmt.Sprintf(
		"https://api.papermc.io/v2/projects/paper/versions/%s/builds/%d/downloads/paper-%s-%d.jar",
		version,
		build.Build,
		version,
		build.Build,
	)
}

// Returns URL to build download, and the build information.
func GetURL(versionInput string, buildID string) (string, Build) {
	version := GetVersion(versionInput)
	build := GetBuild(version, buildID)

	log.Log("using paper version %s", version)

	time, err := time.Parse(time.RFC3339, build.Time)
	log.Error(err, "an error occurred while parsing date supplied by papermc api")

	log.Log("using paper build %d (%s), changes:", build.Build, time.Format("2006-01-02"))

	for _, change := range build.Changes {
		log.RawLog("  (%s) %s\n", change.Commit, change.Summary)
	}

	return formatURL(version, build), build
}
