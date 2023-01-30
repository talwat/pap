package downloadcmds

import (
	//nolint:gosec // Not being used for security, only checksumming. No clue why mojang still uses SHA1.
	"crypto/sha1"

	"github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/jarfiles"
	"github.com/talwat/pap/internal/jarfiles/official"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
	"github.com/urfave/cli/v2"
)

func DownloadOfficialCommand(cCtx *cli.Context) error {
	log.Warn("the official jarfile is much slower and less efficient than paper")
	log.Continue("are you sure you would like to continue?")

	url, pkg := official.GetURL(global.MinecraftVersionInput)

	//nolint:gosec // Not being used for security, only checksumming. No clue why mojang still uses SHA1.
	checksum := net.Download(url, "server.jar", "official server jarfile", sha1.New(), fs.ReadWritePerm)

	log.Success("done downloading")
	jarfiles.VerifyJarfile(checksum, pkg.Downloads.Server.SHA1)

	jarfiles.UnsupportedMessage()

	return nil
}
