package purpur

import (
	"fmt"

	"github.com/talwat/pap/internal/jarfiles"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
)

// Gets the latest version's info.
func GetLatestVersion() Version {
	log.Log("getting versions...")

	var versions Versions

	net.Get(
		"https://api.purpurmc.org/v2/purpur",
		"version list not found, please report this to https://github.com/talwat/pap/issues",
		&versions,
	)

	log.Log("getting latest version info...")

	var version Version

	versionID := versions.Versions[len(versions.Versions)-1]
	log.Debug("latest version: %s", versionID)

	net.Get(
		fmt.Sprintf(
			"https://api.purpurmc.org/v2/purpur/%s",
			versionID,
		),
		fmt.Sprintf(
			"version information for %s not found, please report this to https://github.com/talwat/pap/issues",
			versionID,
		),
		&version,
	)

	return version
}

// Gets a specific version's info.
func GetSpecificVersion(versionID string) Version {
	log.Log("getting info for %s...", versionID)

	var version Version

	statusCode := net.Get(
		fmt.Sprintf(
			"https://api.purpurmc.org/v2/purpur/%s",
			versionID,
		),
		fmt.Sprintf("version information for %s not found", versionID),
		&version,
	)

	jarfiles.APIError(version.Error, statusCode)

	return version
}

// Gets the latest build using the provided version information.
func GetLatestBuild(version Version) Build {
	log.Log("getting latest build info...")

	buildID := version.Builds.Latest
	log.Debug("latest build: %s", buildID)

	var build Build

	net.Get(
		fmt.Sprintf(
			"https://api.purpurmc.org/v2/purpur/%s/%s",
			version.Version,
			buildID,
		),
		fmt.Sprintf("build information for %s not found", buildID),
		&build,
	)

	return build
}

// Gets a specific build using the provided version.
func GetSpecificBuild(version Version, buildID string) Build {
	log.Log("getting build info for %s...", buildID)

	var build Build
	statusCode := net.Get(
		fmt.Sprintf(
			"https://api.purpurmc.org/v2/purpur/%s/%s",
			version.Version,
			buildID,
		),
		fmt.Sprintf("build information for %s not found", buildID),
		&build,
	)

	jarfiles.APIError(build.Error, statusCode)

	return build
}

// Gets a build. It will be the latest one depending on what the input is.
func GetBuild(version Version, buildInput string) Build {
	if buildInput == jarfiles.Latest {
		return GetLatestBuild(version)
	}

	return GetSpecificBuild(version, buildInput)
}

// Gets a version. It will be the latest one depending on what the input is.
func GetVersion(versionInput string) Version {
	if versionInput == jarfiles.Latest {
		return GetLatestVersion()
	}

	return GetSpecificVersion(versionInput)
}
