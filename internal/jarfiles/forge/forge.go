package forge

import (
	"github.com/talwat/pap/internal/log"
)

func GetURL(mverInput, iverInput string, useLatest bool) string {
	var minecraft MinecraftVersion

	var installer InstallerVersion

	if iverInput != "" {
		minecraft, installer = getSpecificInstaller(mverInput, iverInput)
	} else {
		minecraft, installer = getInstaller(mverInput, useLatest)
	}

	log.Log("using minecraft version %s", minecraft.String())
	log.Log("using %s forge installer version %s", installer.Type, installer.Version)

	url := BuildUrl(&minecraft, &installer)

	return url
}
