package main

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"time"
)

const latest = "latest"

func ValidateDownloadOptions() {
	if VersionInput == latest {
		return
	}

	match, err := regexp.MatchString(`^\d\.\d{1,2}(\.\d)?(-pre\d|-SNAPSHOT\d)?$`, VersionInput)
	Error(err, "an error occurred while verifying version")

	if !match {
		CustomError("version %s is not valid", VersionInput)
	}

	if BuildInput == latest {
		return
	}

	match, err = regexp.MatchString(`^\d+$`, BuildInput)
	Error(err, "an error occurred while verifying build")

	if !match {
		CustomError("build %s is not valid", BuildInput)
	}
}

func VerifyJarfile(calculated []byte, build PaperBuild) {
	Log("verifying downloaded jarfile...")

	checksum := hex.EncodeToString(calculated)

	if checksum == build.Downloads.Application.Sha256 {
		Log("checksums match!")
	} else {
		CustomError(
			fmt.Sprintf("checksums (calculated: %s, proper: %s) don't match!",
				calculated,
				build.Downloads.Application.Sha256,
			),
		)
	}
}

// returns URL to build download, and the build information.
func GetURL(versionInput string, buildID string) (string, PaperBuild) {
	var (
		version string
		build   PaperBuild
	)

	if versionInput == latest {
		version = GetLatestVersion()
	} else {
		version = versionInput
	}

	build = GetBuild(version, buildID)

	Log("using paper version %s", version)

	time, err := time.Parse(time.RFC3339, build.Time)

	Error(err, "an error occurred while parsing date supplied by papermc api")

	Log("using paper build %d (%s), changes:", build.Build, time.Format("2006-01-02"))

	for _, change := range build.Changes {
		RawLog("  (%s) %s\n", change.Commit, change.Summary)
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
