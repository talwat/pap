package official

import (
	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
)

func findVersion(versions Versions, version string) Version {
	for i := range versions.Versions {
		if versions.Versions[i].ID == version {
			return versions.Versions[i]
		}
	}

	log.RawError("version %s does not exist", version)

	return Version{}
}

func GetVersionManifest() Versions {
	log.Log("getting version manifest...")

	var versions Versions
	net.Get("https://launchermeta.mojang.com/mc/game/version_manifest.json", &versions)

	return versions
}

func GetSpecificPackage(version string) Package {
	versions := GetVersionManifest()

	log.Log("locating version %s...", version)
	versionInfo := findVersion(versions, version)

	log.Log("getting package for %s...", version)

	var pkg Package
	net.Get(versionInfo.URL, &pkg)

	return pkg
}

func GetLatestPackage() Package {
	versions := GetVersionManifest()

	var versionToGet string

	if global.OfficialUseSnapshotInput {
		versionToGet = versions.Latest.Snapshot
	} else {
		versionToGet = versions.Latest.Release
	}

	log.Log("locating version %s...", versionToGet)
	versionInfo := findVersion(versions, versionToGet)

	log.Log("getting package for %s...", versionToGet)

	var pkg Package
	net.Get(versionInfo.URL, &pkg)

	return pkg
}

func GetPackage(versionInput string) Package {
	if versionInput == "latest" {
		return GetLatestPackage()
	}

	return GetSpecificPackage(versionInput)
}
