package cmd

import (
	"github.com/talwat/pap/app/fs"
	"github.com/talwat/pap/app/global"
	"github.com/talwat/pap/app/log"
	"github.com/talwat/pap/app/net"
	"github.com/talwat/pap/app/properties"
	"github.com/urfave/cli/v2"
)

func GeyserCommand(cCtx *cli.Context) error {
	fs.MakeDirectory("plugins")

	//nolint:lll
	net.Download(
		"https://ci.opencollab.dev/job/GeyserMC/job/Geyser/job/master/lastSuccessfulBuild/artifact/bootstrap/spigot/build/libs/Geyser-Spigot.jar",
		"plugins/Geyser-Spigot.jar",
		"geyser",
	)

	if !global.NoFloodGateInput {
		//nolint:lll
		net.Download(
			"https://ci.opencollab.dev/job/GeyserMC/job/Floodgate/job/master/lastSuccessfulBuild/artifact/spigot/build/libs/floodgate-spigot.jar",
			"plugins/floodgate-spigot.jar",
			"floodgate",
		)
	}

	log.Log("floodgate and geyser do not support key signing yet for chat messages")
	log.Log("this feature was introduced in 1.19.1, so you do not have to disable it if your version is below that")

	if log.YesOrNo(
		"y",
		"would you like to disable it?",
	) {
		properties.EditProperty("enforce-secure-profile", "false")
	}

	return nil
}
