package cmd

import (
	"crypto/sha256"

	"github.com/talwat/pap/app/global"
	"github.com/talwat/pap/app/jarfiles/paper"
	"github.com/talwat/pap/app/log"
	"github.com/talwat/pap/app/net"
	"github.com/urfave/cli/v2"
)

func validateOptions() {
	const latest = "latest"

	if global.VersionInput != latest {
		ValidateOption(global.VersionInput, `^\d\.\d{1,2}(\.\d)?(-pre\d|-SNAPSHOT\d)?$`, "version")
	}

	if global.BuildInput != latest {
		ValidateOption(global.BuildInput, `^\d+$`, "build")
	}
}

func DownloadCommand(cCtx *cli.Context) error {
	validateOptions()

	url, build := paper.GetURL(global.VersionInput, global.BuildInput)

	checksum := net.Download(url, "paper.jar", "paper jarfile", sha256.New())

	log.Log("done downloading")
	paper.VerifyJarfile(checksum, build)

	return nil
}
