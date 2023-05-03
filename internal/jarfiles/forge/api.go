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

	net.Get(
		"https://files.minecraftforge.net/maven/net/minecraftforge/forge/promotions_slim.json",
		"could not retrieve promotions",
		&promotions,
	)

	return promotions
}

func getInstaller(version string, useLatestInstaller bool) (MinecraftVersion, InstallerVersion) {
	var minecraft MinecraftVersion

	var installer InstallerVersion

	promos := getPromotions()

	if version == jarfiles.Latest {
		minecraft = getLatestMinecraftVersion(&promos)
	} else {
		minecraft = parseMinecraftVersion(version)
	}

	if useLatestInstaller {
		installer = getVersion(&promos, &minecraft, jarfiles.Latest)
	} else {
		installer = getVersion(&promos, &minecraft, "recommended")

		if (installer == InstallerVersion{}) {
			log.Continue("no recommended installer found for version %s. use the latest version?", minecraft.String())
		}

		installer = getVersion(&promos, &minecraft, jarfiles.Latest)
	}

	if (installer == InstallerVersion{}) {
		log.RawError("could not get a valid installer version")
	}

	return minecraft, installer
}

func getSpecificInstaller(version string, installer string) (MinecraftVersion, InstallerVersion) {
	promos := getPromotions()

	var minecraft MinecraftVersion

	if version == "latest" {
		minecraft = getLatestMinecraftVersion(&promos)
	} else {
		minecraft = parseMinecraftVersion(version)
	}

	if installer == "latest" {
		return minecraft, getVersion(&promos, &minecraft, "latest")
	}

	return minecraft, InstallerVersion{
		Version: installer,
	}
}

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
