package forge

import (
	"fmt"

	"github.com/talwat/pap/internal/log"
)

func GetURL(versionInput string, installerInput string) string {
	var minecraft MinecraftVersion
	var installer InstallerVersion

	if versionInput == "latest" {
		versions := GetMinecraftVersions()
		minecraft = getLatestVersion(versions)
	}

	if installerInput == "latest" {
		versions := GetInstallerVersions()
		installer = getLatestInstaller(versions)
	}

	log.Log("using minecraft version %s", minecraft.Version)
	log.Log("using %s forge installer version %s", installer.Type, installer.Version)

	prefix := "https://maven.minecraftforge.net/net/minecraftforge/forge"
	return fmt.Sprintf("%s/%s-%s/forge-%s-%s-installer.jar", prefix, minecraft.Version, installer.Version, minecraft.Version, installer.Version)
}
