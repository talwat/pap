package paper

import (
	"fmt"

	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/jarfiles"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
)

// Gets the latest minecraft version in the list of versions.
func GetLatestVersion() string {
	var versions Versions

	log.Log("getting latest version information")
	net.Get(
		"https://api.papermc.io/v2/projects/paper",
		"version information not found, please report this to https://github.com/talwat/pap/issues",
		&versions,
	)

	version := versions.Versions[len(versions.Versions)-1]
	log.Debug("latest version: %s", version)

	return version
}

// Gets the latest build in a version.
func GetLatestBuild(version string) Build {
	var builds Builds

	log.Log("getting latest build information")

	url := fmt.Sprintf("https://api.papermc.io/v2/projects/paper/versions/%s/builds", version)
	statusCode := net.Get(url, fmt.Sprintf("build information for %s not found", version), &builds)

	jarfiles.APIError(builds.Error, statusCode)

	// latest build, can be experimental or stable
	latest := builds.Builds[len(builds.Builds)-1]

	if global.PaperExperimentalBuildInput {
		log.Debug("using latest build (%d) regardless", latest.Build)

		return latest
	}

	// Iterate through builds.Builds backwards
	for i := len(builds.Builds) - 1; i >= 0; i-- {
		if builds.Builds[i].Channel == "default" { // "default" usually means stable
			return builds.Builds[i] // Stable build found, return it
		}
	}

	log.Continue("no stable build found, would you like to use the latest experimental build?")

	return latest
}

// Gets a specific build in a version.
func GetSpecificBuild(version string, buildID string) Build {
	log.Log("getting build information for %s", buildID)

	var build Build

	url := fmt.Sprintf("https://api.papermc.io/v2/projects/paper/versions/%s/builds/%s", version, buildID)
	statusCode := net.Get(url, fmt.Sprintf("build %s of version %s not found", buildID, version), &build)

	jarfiles.APIError(build.Error, statusCode)

	return build
}

// Gets either the latest build or a specific one depending on buildID.
func GetBuild(version string, buildID string) Build {
	var build Build

	if buildID == jarfiles.Latest {
		build = GetLatestBuild(version)
	} else {
		build = GetSpecificBuild(version, buildID)
	}

	if build.Channel == "experimental" && !global.PaperExperimentalBuildInput {
		log.Continue(
			"build %d has been flagged as experimental, are you sure you would like to download it?",
			build.Build,
		)
	}

	return build
}

// Gets a specific version or the latest one depending on the input.
// Additionally, if using a specific one, it will not get the latest release of a specific minor release.
// For example if you put in 1.12, you will get 1.12 and not 1.12.2.
func GetVersion(version string) string {
	if version == jarfiles.Latest {
		return GetLatestVersion()
	}

	return version
}
