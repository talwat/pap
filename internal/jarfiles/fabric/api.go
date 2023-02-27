package fabric

import (
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
)

func GetMinecraftVersions() MinecraftVersions {
	log.Log("getting version list...")

	var versions MinecraftVersions

	net.Get(
		"https://meta.fabricmc.net/v2/versions",
		"minecraft version information not found, please report this to https://github.com/talwat/pap/issues",
		&versions,
	)

	return versions
}

func GetLoaderVersions() []LoaderVersion {
	log.Log("getting loader list...")

	var versions []LoaderVersion

	net.Get(
		"https://meta.fabricmc.net/v2/versions/loader",
		"loader version information not found, please report this to https://github.com/talwat/pap/issues",
		&versions,
	)

	return versions
}

func GetInstallerVersions() []InstallerVersion {
	log.Log("getting installer list...")

	var versions []InstallerVersion

	net.Get(
		"https://meta.fabricmc.net/v2/versions/installer/",
		"installer version information not found, please report this to https://github.com/talwat/pap/issues",
		&versions,
	)

	return versions
}
