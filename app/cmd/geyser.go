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

	disableKeySigning := log.YesOrNo(
		"y",
		"floodgate and geyser do not support key signing yet, would you like to disable it (recommended)?",
	)

	if disableKeySigning {
		properties.EditProperty("enforce-secure-profile", "false")
	}

	return nil
}
