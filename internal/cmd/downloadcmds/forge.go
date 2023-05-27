package downloadcmds

import (
	"fmt"

	"github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/jarfiles"
	"github.com/talwat/pap/internal/jarfiles/forge"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
	"github.com/urfave/cli/v2"
)

func DownloadForgeCommand(cCtx *cli.Context) error {
	url := forge.GetURL(
		global.MinecraftVersionInput,
		global.ForgeInstallerVersion,
		global.ForgeUseLatestInstaller,
	)

	net.Download(
		url,
		"resolved forge installer jarfile not found",
		fmt.Sprintf("forge-%s-%s-installer.jar", global.MinecraftVersionInput, global.ForgeInstallerVersion),
		"forge server installer jarfile",
		nil,
		fs.ReadWritePerm,
	)

	log.Success("done downloading")

	jarfiles.UnsupportedMessage()

	return nil
}
