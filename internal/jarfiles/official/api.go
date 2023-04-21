package official

import (
	"fmt"

	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/jarfiles"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
)

// Finds the version in a list of versions.
func FindVersion(versions Versions, version string) Version {
	for i := range versions.Versions {
		if versions.Versions[i].ID == version {
			return versions.Versions[i]
		}
	}

	log.RawError("version %s does not exist", version)

	//nolint:exhaustruct // The process will be terminated by log.RawError before this ever gets run.
	return Version{}
}

func GetVersionManifest() Versions {
	log.Log("getting version manifest...")

	var versions Versions

	net.Get(
		"https://launchermeta.mojang.com/mc/game/version_manifest.json",
		"version manifest not found, please report this to https://github.com/talwat/pap/issues",
		&versions,
	)

	return versions
}

func GetSpecificPackage(version string) Package {
	versions := GetVersionManifest()

	log.Log("locating version %s...", version)
	versionInfo := FindVersion(versions, version)

	log.Log("getting package for %s...", version)

	var pkg Package

	net.Get(versionInfo.URL, fmt.Sprintf("package %s not found", version), &pkg)

	return pkg
}

func GetLatestPackage() Package {
	versions := GetVersionManifest()

	var version string

	if global.UseSnapshotInput {
		version = versions.Latest.Snapshot
	} else {
		version = versions.Latest.Release
	}

	log.Log("locating version %s...", version)
	versionInfo := FindVersion(versions, version)

	log.Log("getting package for %s...", version)

	var pkg Package

	net.Get(versionInfo.URL, "latest package not found", &pkg)

	return pkg
}

func GetPackage(versionInput string) Package {
	if versionInput == jarfiles.Latest {
		return GetLatestPackage()
	}

	return GetSpecificPackage(versionInput)
}
