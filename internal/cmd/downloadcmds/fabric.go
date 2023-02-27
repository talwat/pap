package downloadcmds

import (
	"github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/jarfiles"
	"github.com/talwat/pap/internal/jarfiles/fabric"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
	"github.com/urfave/cli/v2"
)

func DownloadFabricCommand(cCtx *cli.Context) error {
	url := fabric.GetURL(
		global.MinecraftVersionInput,
		global.FabricLoaderVersion,
		global.FabricInstallerVersion,
	)

	net.Download(
		url,
		"resolved fabric jarfile not found",
		"fabric.jar",
		"fabric server jarfile",
		nil,
		fs.ReadWritePerm,
	)

	log.Success("done downloading")

	jarfiles.UnsupportedMessage()

	return nil
}
