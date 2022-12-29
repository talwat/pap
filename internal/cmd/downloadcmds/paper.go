package downloadcmds

import (
	"crypto/sha256"

	"github.com/talwat/pap/internal/cmd"
	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/jarfiles/paper"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
	"github.com/urfave/cli/v2"
)

func validateOptions() {
	const latest = "latest"

	if global.VersionInput != latest {
		cmd.ValidateOption(global.VersionInput, `^\d\.\d{1,2}(\.\d)?(-pre\d|-SNAPSHOT\d)?$`, "version")
	}

	if global.BuildInput != latest {
		cmd.ValidateOption(global.BuildInput, `^\d+$`, "build")
	}
}

func DownloadPaperCommand(cCtx *cli.Context) error {
	validateOptions()

	url, build := paper.GetURL(global.VersionInput, global.BuildInput)

	checksum := net.Download(url, "paper.jar", "paper jarfile", sha256.New())

	log.Log("done downloading")
	paper.VerifyJarfile(checksum, build)

	return nil
}
