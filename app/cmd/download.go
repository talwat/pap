package cmd

import (
	"regexp"

	"github.com/talwat/pap/app/global"
	"github.com/talwat/pap/app/log"
	"github.com/talwat/pap/app/net"
	"github.com/talwat/pap/app/paper"
	"github.com/urfave/cli/v2"
)

func validateOptions() {
	const latest = "latest"

	if global.VersionInput == latest {
		return
	}

	match, err := regexp.MatchString(`^\d\.\d{1,2}(\.\d)?(-pre\d|-SNAPSHOT\d)?$`, global.VersionInput)
	log.Error(err, "an error occurred while verifying version")

	if !match {
		log.CustomError("version %s is not valid", global.VersionInput)
	}

	if global.BuildInput == latest {
		return
	}

	match, err = regexp.MatchString(`^\d+$`, global.BuildInput)
	log.Error(err, "an error occurred while verifying build")

	if !match {
		log.CustomError("build %s is not valid", global.BuildInput)
	}
}

func DownloadCommand(cCtx *cli.Context) error {
	validateOptions()

	url, build := paper.GetURL(global.VersionInput, global.BuildInput)

	checksum := net.Download(url, "paper.jar", "paper jarfile")
	paper.VerifyJarfile(checksum, build)

	return nil
}
