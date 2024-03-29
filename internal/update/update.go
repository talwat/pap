package update

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
)

func parseVersion(rawVersion string) []string {
	noExtra := strings.Split(rawVersion, "-")[0]

	return strings.Split(noExtra, ".")
}

func checkIfNewUpdate() string {
	log.Log("checking for a new update...")

	rawLatest, statusCode := net.GetPlainText(
		"https://raw.githubusercontent.com/talwat/pap/main/version.txt",
		"latest version information not found, please report this to https://github.com/talwat/pap/issues",
	)

	if statusCode != http.StatusOK {
		log.RawError("http request to get latest version returned %d", statusCode)
	}

	log.Debug("raw latest version: %s", rawLatest)

	latest := parseVersion(rawLatest)
	log.Debug("parsed latest version: %s", latest)

	current := parseVersion(global.Version)
	log.Debug("parsed current version: %s", current)

	latestLen := len(latest)
	currentLen := len(current)

	if latestLen != currentLen {
		log.RawError(
			//nolint:lll
			"latest (%s) and current (%s) version are different lengths, please report this issue to https://github.com/talwat/pap/issues",
			latest,
			current,
		)
	}

	for idx := range latest {
		switch {
		case latest[idx] < current[idx]:
			log.Debug("%s > %s, assuming you are using a development version", current[idx], latest[idx])

			log.Log("pap is newer than the current latest version (you are a developer?)")
			os.Exit(0)
		case latest[idx] > current[idx]:
			log.Debug("%s > %s, out of date!", latest[idx], current[idx])
			log.Log("out of date! current version is %s, latest is %s", global.Version, rawLatest)

			return rawLatest
		default:
			continue
		}
	}

	if global.ReinstallInput {
		log.Warn("pap is up to date, but --reinstall is set, so continuing")

		return rawLatest
	}

	log.Log("pap is up to date")
	os.Exit(0)

	return ""
}

func getExePath() string {
	exe, err := os.Executable()
	log.Error(err, "an error occurred while finding location of currently installed executable")
	log.Debug("executable path: %s", exe)

	evaluatedExe, err := filepath.EvalSymlinks(exe)
	log.Error(err, "an error occurred while locating location of original executable, perhaps a broken symlink")
	log.Debug("evaluated exe: %s", evaluatedExe)

	if runtime.GOOS == "windows" {
		log.Debug("running on windows, skipping path check")

		return evaluatedExe
	}

	home, err := os.UserHomeDir()
	homePath := filepath.Join(home, "/.local/bin/pap")

	log.Error(err, "an error occurred while getting the user's home directory")
	log.Debug("pap local: %s", homePath)

	if evaluatedExe != "/usr/bin" && evaluatedExe != homePath {
		log.Warn("it seems like you installed pap in a location not specified by the install guide (%s)", evaluatedExe)
		log.Warn("if this is expected, you can ignore this.")
		log.Continue("would like to continue?")
	}

	return evaluatedExe
}

func Update() {
	latest := checkIfNewUpdate()

	log.Log("finding exe...")

	exe := getExePath()
	url := fmt.Sprintf(
		"https://github.com/talwat/pap/releases/download/v%s/pap_%s_%s_%s",
		latest,
		latest,
		runtime.GOOS,
		runtime.GOARCH,
	)

	tmpPath := fmt.Sprintf("/tmp/pap-update-%s", latest)

	net.Download(
		url,
		"release not found, your OS or architecture may not be supported",
		tmpPath,
		fmt.Sprintf("pap %s", latest),
		nil,
		fs.ExecutePerm,
	)

	log.Log("installing pap...")
	fs.MoveFile(tmpPath, exe)

	log.Success("done! updated pap to %s", latest)
}
