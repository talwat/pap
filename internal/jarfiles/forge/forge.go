package forge

import (
	"fmt"

	"github.com/talwat/pap/internal/log"
)

func GetURL(versionInput string, useLatest bool) string {
	minecraft := versionInput
	var installer InstallerVersion

	if versionInput != "latest" {
		if useLatest {
			installer = GetLatestInstaller(versionInput)
		} else {
			installer = GetInstaller(versionInput)
		}
	} else {
		mc, iv := GetLatest()
		minecraft = mc
		installer = InstallerVersion{Version: iv, Type: "latest"}
	}

	if installer.Version == "" {
		log.RawError("no installer found for version %s", versionInput)
	}

	log.Log("using minecraft version %s", minecraft)
	log.Log("using %s forge installer version %s", installer.Type, installer.Version)

	prefix := "https://maven.minecraftforge.net/net/minecraftforge/forge"
	return fmt.Sprintf("%s/%s-%s/forge-%s-%s-installer.jar", prefix, minecraft, installer.Version, minecraft, installer.Version)
}
