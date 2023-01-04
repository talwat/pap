// Interact with papermc downloads api and verification of downloaded files.
package paper

// This is the only file which is accessed from other packages.
// If you would like to add compatibility for other jarfile types, you only need a GetURL() function.

import (
	"fmt"
	"time"

	"github.com/talwat/pap/internal/log"
)

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
