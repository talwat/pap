package cmd

import (
	"crypto/sha256"

	"github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
	"github.com/talwat/pap/internal/properties"
	"github.com/urfave/cli/v2"
)

func GeyserCommand(cCtx *cli.Context) error {
	log.Warn("this command is deprecated, please install geyser & floodgate with pap plugin install geyser")
	log.Warn("this downloads the spigot version of geyser, and thus requires spigot or paper")
	fs.MakeDirectory("plugins")

	//nolint:lll
	net.Download(
		"https://ci.opencollab.dev/job/GeyserMC/job/Geyser/job/master/lastSuccessfulBuild/artifact/bootstrap/spigot/build/libs/Geyser-Spigot.jar",
		"plugins/Geyser-Spigot.jar",
		"geyser",
		sha256.New(),
	)

	if !global.NoFloodGateInput {
		//nolint:lll
		net.Download(
			"https://ci.opencollab.dev/job/GeyserMC/job/Floodgate/job/master/lastSuccessfulBuild/artifact/spigot/build/libs/floodgate-spigot.jar",
			"plugins/floodgate-spigot.jar",
			"floodgate",
			sha256.New(),
		)
	}

	log.Log("floodgate and geyser do not support key signing yet for chat messages")
	log.Log("this feature was introduced in 1.19.1, so you do not have to disable it if your version is below that")

	if log.YesOrNo(
		"y",
		"would you like to disable it?",
	) {
		properties.SetProperty("enforce-secure-profile", "false")
	}

	return nil
}
