package forge

import (
	"fmt"

	"github.com/talwat/pap/internal/jarfiles"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
)

func BuildURL(minecraft *MinecraftVersion, installer *InstallerVersion) string {
	var returnURL string

	returnURL += "https://maven.minecraftforge.net/net/minecraftforge/forge"
	returnURL += fmt.Sprintf("/%s-%s", minecraft.String(), installer.Version)

	if minecraft.Minor == 8 || minecraft.Minor == 9 {
		returnURL += fmt.Sprintf("-%d.%d.%d", minecraft.Major, minecraft.Minor, minecraft.Patch)
	} else if minecraft.IsPrerelease {
		returnURL += "-prerelease"
	}

	returnURL += fmt.Sprintf("/forge-%s-%s-", minecraft.String(), installer.Version)

	// Forge versioning scheme changes with these two specific versions:
	// Versions 1.3 -> 1.7 and 1.10 -> latest use
	// MinecraftVersion-InstallerVersion/forge-MinecraftVersion-InstallerVersion-installer.jar
	// 1.19.4              45.0.57                 1.19.4           45.0.57
	//
	// While 1.8 and 1.9 use a different scheme
	// MinecraftVersion-InstallerVersion-VersionTriple/forge-MinecraftVersion-InstallerVersion-VersionTriple-installer.jar
	//     1.9            12.16.1.1938       1.9.0                 1.9          12.16.1.1938      1.9.0

	if minecraft.Minor == 8 || minecraft.Minor == 9 {
		returnURL += fmt.Sprintf("%d.%d.%d-", minecraft.Major, minecraft.Minor, minecraft.Patch)
	} else if minecraft.IsPrerelease {
		returnURL += "prerelease-"
	}

	returnURL += "installer.jar"

	return returnURL
}

func getPromotions() PromotionsSlim {
	var promotions PromotionsSlim

	log.Log("getting promotions...")

	net.Get(
		"https://files.minecraftforge.net/maven/net/minecraftforge/forge/promotions_slim.json",
		"could not retrieve promotions",
		&promotions,
	)

	return promotions
}

func getInstaller(version string, useLatestInstaller bool) (MinecraftVersion, InstallerVersion) {
	if useLatestInstaller {
		log.Debug("using latest installer version")

		return getSpecificInstaller(version, jarfiles.Latest)
	}

	promos := getPromotions()

	var minecraft MinecraftVersion

	if version == jarfiles.Latest {
		log.Debug("using latest minecraft version")

		minecraft = getLatestMinecraftVersion(&promos)
	} else {
		minecraft = parseMinecraftVersion(version)
	}

	installer := getVersion(&promos, &minecraft, "recommended")

	if (installer == InstallerVersion{}) {
		log.Continue("no recommended installer found for version %s. use the latest version?", minecraft.String())
	}

	installer = getVersion(&promos, &minecraft, jarfiles.Latest)

	if (installer == InstallerVersion{}) {
		log.RawError("could not get a valid installer version")
	}

	return minecraft, installer
}

func getSpecificInstaller(version string, installer string) (MinecraftVersion, InstallerVersion) {
	promos := getPromotions()

	var minecraft MinecraftVersion

	if version == jarfiles.Latest {
		log.Debug("using latest minecraft version")
		
		minecraft = getLatestMinecraftVersion(&promos)
	} else {
		minecraft = parseMinecraftVersion(version)
	}

	if installer == jarfiles.Latest {
		log.Debug("using latest installer version")

		return minecraft, getVersion(&promos, &minecraft, "latest")
	}

	return minecraft, InstallerVersion{
		Version: installer,
	}
}

// `golangci-lint` complains about *MinecraftVersion because this function only uses the string value
// of the version. `interfacer` says we can use fmt.Stringer here, but that may lead to confusion.

//nolint:interfacer
func getVersion(promos *PromotionsSlim, minecraft *MinecraftVersion, installerType string) InstallerVersion {
	promo, found := promos.Promos[fmt.Sprintf("%s-%s", minecraft.String(), installerType)]

	if found {
		return InstallerVersion{
			Version: promo,
			Type:    installerType,
		}
	}

	return InstallerVersion{}
}
