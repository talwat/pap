package forge

import (
	"fmt"

	"github.com/talwat/pap/internal/log"
)

func GetLatestURL(versionInput string, installerInput string) string {
	installer := GetLatestInstaller(versionInput)

	log.Log("using minecraft version %s", versionInput)
	log.Log("using %s forge installer version %s", installer.Type, installer.Version)

	prefix := "https://maven.minecraftforge.net/net/minecraftforge/forge"
	return fmt.Sprintf("%s/%s-%s/forge-%s-%s-installer.jar", prefix, versionInput, installer.Version, versionInput, installer.Version)
}
