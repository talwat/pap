// Interact with official mojang downloads api and verification of downloaded files.
package official

import (
	"os"

	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/log/color"
)

func GetURL(versionInput string) (string, Package) {
	pkg := GetPackage(versionInput)

	if pkg.Downloads.Server.URL == "" {
		log.Log("%serror%s: the server URL could not be found", color.Red, color.Reset)
		log.Log("%serror%s: this may be because server versions below 1.2.5 are not available", color.Red, color.Reset)
		os.Exit(1)
	}

	return pkg.Downloads.Server.URL, pkg
}
