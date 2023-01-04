package paper

import (
	"fmt"

	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/jarfiles"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
)

func GetLatestVersion() string {
	var versions Versions

	log.Log("getting latest version information")
	net.Get("https://api.papermc.io/v2/projects/paper", &versions)

	return versions.Versions[len(versions.Versions)-1]
}

func GetLatestBuild(version string) Build {
	var builds Builds

	log.Log("getting latest build information")

	url := fmt.Sprintf("https://api.papermc.io/v2/projects/paper/versions/%s/builds", version)
	statusCode := net.Get(url, &builds)

	jarfiles.APIError(builds.Error, statusCode)

	// latest build, can be experimental or stable
	latest := builds.Builds[len(builds.Builds)-1]

	if global.PaperExperimentalBuildInput {
		return latest
	}

	// Iterate through builds.Builds backwards
	for i := len(builds.Builds) - 1; i >= 0; i-- {
		if builds.Builds[i].Channel == "default" { // "default" usually means stable
			return builds.Builds[i] // Stable build found, return it
		}
	}

	log.Continue("warning: no stable build found, would you like to use the latest experimental build?")

	return latest
}

func GetSpecificBuild(version string, buildID string) Build {
	log.Log("getting build information for %s", buildID)

	var build Build

	url := fmt.Sprintf("https://api.papermc.io/v2/projects/paper/versions/%s/builds/%s", version, buildID)
	statusCode := net.Get(url, &build)

	jarfiles.APIError(build.Error, statusCode)

	return build
}

func GetBuild(version string, buildID string) Build {
	var build Build

	if buildID == "latest" {
		build = GetLatestBuild(version)
	} else {
		build = GetSpecificBuild(version, buildID)
	}

	if build.Channel == "experimental" && !global.PaperExperimentalBuildInput {
		log.Continue(
			"warning: build %d has been flagged as experimental, are you sure you would like to download it?",
			build.Build,
		)
	}

	return build
}

func GetVersion(versionInput string) string {
	if versionInput == jarfiles.Latest {
		return GetLatestVersion()
	}

	return versionInput
}
