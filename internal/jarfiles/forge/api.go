package forge

import (
	"fmt"

	"github.com/talwat/pap/internal/jarfiles"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
)

func BuildUrl(minecraft *MinecraftVersion, installer *InstallerVersion) string {
	var returnUrl string

	returnUrl += "https://maven.minecraftforge.net/net/minecraftforge/forge"
	returnUrl += fmt.Sprintf("/%s-%s", minecraft.String(), installer.Version)

	if minecraft.Minor == 8 || minecraft.Minor == 9 {
		returnUrl += fmt.Sprintf("-%d.%d.%d", minecraft.Major, minecraft.Minor, minecraft.Patch)
	} else if minecraft.IsPrerelease {
		returnUrl += "-prerelease"
	}

	returnUrl += fmt.Sprintf("/forge-%s-%s-", minecraft.String(), installer.Version)

	if minecraft.Minor == 8 || minecraft.Minor == 9 {
		returnUrl += fmt.Sprintf("%d.%d.%d-", minecraft.Major, minecraft.Minor, minecraft.Patch)
	} else if minecraft.IsPrerelease {
		returnUrl += "prerelease-"
	}

	returnUrl += "installer.jar"

	return returnUrl
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
