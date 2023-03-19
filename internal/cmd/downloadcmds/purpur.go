package downloadcmds

//nolint:gosec // Not being used for security, only checksumming. Why does purpur use MD5?
import (
	"crypto/md5"

	"github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/jarfiles"
	"github.com/talwat/pap/internal/jarfiles/purpur"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
	"github.com/urfave/cli/v2"
)

//nolint:revive // cCtx kept for consistency with other commands.
func DownloadPurpurCommand(cCtx *cli.Context) error {
	url, build := purpur.GetURL(global.MinecraftVersionInput, global.JarBuildInput)

	//nolint:gosec // Not being used for security, only checksumming. Why does purpur use MD5?
	checksum := net.Download(
		url,
		"resolved purpur jarfile not found",
		"purpur.jar",
		"purpur jarfile",
		md5.New(),
		fs.ReadWritePerm,
	)

	log.Success("done downloading")
	jarfiles.VerifyJarfile(checksum, build.MD5)

	return nil
}
