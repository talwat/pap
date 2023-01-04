package purpur

import (
	"fmt"

	"github.com/talwat/pap/internal/jarfiles"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
)

func GetLatestVersion() Version {
	log.Log("getting versions...")

	var versions Versions

	net.Get("https://api.purpurmc.org/v2/purpur", &versions)

	log.Log("getting latest version info...")

	var version Version

	net.Get(
		fmt.Sprintf(
			"https://api.purpurmc.org/v2/purpur/%s",
			versions.Versions[len(versions.Versions)-1],
		),
		&version,
	)

	return version
}

func GetSpecificVersion(versionID string) Version {
	log.Log("getting info for %s...", versionID)

	var version Version
	statusCode := net.Get(
		fmt.Sprintf(
			"https://api.purpurmc.org/v2/purpur/%s",
			versionID,
		),
		&version,
	)

	jarfiles.APIError(version.Error, statusCode)

	return version
}

func GetLatestBuild(version Version) Build {
	log.Log("getting latest build info...")

	buildID := version.Builds.Latest

	var build Build

	net.Get(
		fmt.Sprintf(
			"https://api.purpurmc.org/v2/purpur/%s/%s",
			version.Version,
			buildID,
		),
		&build,
	)

	return build
}

func GetSpecificBuild(version Version, buildID string) Build {
	log.Log("getting build info for %s...", buildID)

	var build Build
	statusCode := net.Get(
		fmt.Sprintf(
			"https://api.purpurmc.org/v2/purpur/%s/%s",
			version.Version,
			buildID,
		),
		&build,
	)

	jarfiles.APIError(build.Error, statusCode)

	return build
}

func GetBuild(version Version, buildInput string) Build {
	if buildInput == jarfiles.Latest {
		return GetLatestBuild(version)
	}

	return GetSpecificBuild(version, buildInput)
}

func GetVersion(versionInput string) Version {
	if versionInput == jarfiles.Latest {
		return GetLatestVersion()
	}

	return GetSpecificVersion(versionInput)
}
