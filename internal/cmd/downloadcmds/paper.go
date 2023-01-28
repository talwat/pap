package downloadcmds

import (
	"crypto/sha256"

	"github.com/talwat/pap/internal/cmd"
	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/jarfiles"
	"github.com/talwat/pap/internal/jarfiles/paper"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
	"github.com/urfave/cli/v2"
)

func validatePaperOptions() {
	const latest = "latest"

	if global.MinecraftVersionInput != latest {
		cmd.ValidateOption(global.MinecraftVersionInput, `^\d\.\d{1,2}(\.\d)?(-pre\d|-SNAPSHOT\d)?$`, "version")
	}

	if global.JarBuildInput != latest {
		cmd.ValidateOption(global.JarBuildInput, `^\d+$`, "build")
	}
}

func DownloadPaperCommand(cCtx *cli.Context) error {
	validatePaperOptions()

	url, build := paper.GetURL(global.MinecraftVersionInput, global.JarBuildInput)

	checksum := net.Download(url, "paper.jar", "paper jarfile", sha256.New())

	log.Success("done downloading")
	jarfiles.VerifyJarfile(checksum, build.Downloads.Application.Sha256)

	return nil
}
