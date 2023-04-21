package forge

import (
	"fmt"

	"github.com/talwat/pap/internal/log"
)

func GetURL(versionInput string, installerInput string) string {
	var installer InstallerVersion
	if installerInput == "latest" {
		installer = GetLatestInstaller(versionInput)
	} else if installerInput == "recommended" {
		installer = GetRecommendedInstaller(versionInput)
	} else {
		installer = GetInstaller(versionInput)
	}

	log.Log("using minecraft version %s", versionInput)
	log.Log("using %s forge installer version %s", installer.Type, installer.Version)

	prefix := "https://maven.minecraftforge.net/net/minecraftforge/forge"
	return fmt.Sprintf("%s/%s-%s/forge-%s-%s-installer.jar", prefix, versionInput, installer.Version, versionInput, installer.Version)
}
