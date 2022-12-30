package downloadcmds

import (
	"crypto/sha1"

	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/jarfiles"
	"github.com/talwat/pap/internal/jarfiles/official"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
	"github.com/urfave/cli/v2"
)

func DownloadOfficialCommand(cCtx *cli.Context) error {
	url, pkg := official.GetURL(global.MinecraftVersionInput)

	checksum := net.Download(url, "server.jar", "official server jarfile", sha1.New())

	log.Log("done downloading")
	jarfiles.VerifyJarfile(checksum, pkg.Downloads.Server.SHA1)

	return nil
}
