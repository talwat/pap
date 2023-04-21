package fabric

import (
	"fmt"

	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/jarfiles"
	"github.com/talwat/pap/internal/log"
)

func getLatestVersion(versions MinecraftVersions) MinecraftVersion {
	latest := versions.Game[0]

	if global.UseSnapshotInput {
		log.Debug("using experimental/snapshot minecraft version (%s) regardless", latest.Version)

		return latest
	}

	for _, version := range versions.Game {
		if version.Stable {
			return version
		}
	}

	return latest
}

func getLatestLoader(versions []LoaderVersion) LoaderVersion {
	latest := versions[0]

	if global.FabricExperimentalLoaderVersion {
		log.Debug("using experimental fabric loader version (%s) regardless", latest.Version)

		return latest
	}

	for _, version := range versions {
		if version.Stable {
			return version
		}
	}

	return latest
}

func getLatestInstaller(versions []InstallerVersion) InstallerVersion {
	latest := versions[0]

	for _, version := range versions {
		if version.Stable {
			return version
		}
	}

	return latest
}

func GetURL(versionInput string, loaderInput string, installerInput string) string {
	version := versionInput
	loader := loaderInput
	installer := installerInput

	// A bit of repetitive code, but I am not willing to use generics to do this.
	if version == jarfiles.Latest {
		versions := GetMinecraftVersions()
		version = getLatestVersion(versions).Version
	}

	if loader == jarfiles.Latest {
		versions := GetLoaderVersions()
		loader = getLatestLoader(versions).Version
	}

	if installer == jarfiles.Latest {
		versions := GetInstallerVersions()
		installer = getLatestInstaller(versions).Version
	}

	log.Log("using minecraft version %s", version)
	log.Log("using fabric loader version %s", loader)
	log.Log("using fabric installer version %s", installer)

	return fmt.Sprintf("https://meta.fabricmc.net/v2/versions/loader/%s/%s/%s/server/jar", version, loader, installer)
}
