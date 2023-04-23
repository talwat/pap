package forge

import (
	"fmt"

	"github.com/talwat/pap/internal/log"
)

func GetURL(mverInput, iverInput string, useLatest bool) string {
	var minecraft MinecraftVersion
	var installer InstallerVersion

	if iverInput != "" {
		minecraft, _ = getInstaller(mverInput, useLatest)
		installer = getSpecificInstaller(&minecraft, iverInput)
	} else {
		minecraft, installer = getInstaller(mverInput, useLatest)
	}

	log.Log("using minecraft version %s", minecraft.String())
	log.Log("using %s forge installer version %s", installer.Type, installer.Version)

	prefix := "https://maven.minecraftforge.net/net/minecraftforge/forge"
	return fmt.Sprintf("%s/%s-%s/forge-%s-%s-installer.jar", prefix, minecraft.String(), installer.Version, minecraft.String(), installer.Version)
}
