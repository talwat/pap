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

	rawLatest, statusCode := net.GetPlainText("https://raw.githubusercontent.com/talwat/pap/main/version.txt")

	if statusCode != http.StatusOK {
		log.RawError("http request to get latest version returned %d", statusCode)
	}

	latest := parseVersion(rawLatest)
	current := parseVersion(global.Version)

	for idx := len(latest) - 1; idx >= 0; idx-- {
		switch {
		case latest[idx] > current[idx]:
			log.Log("out of date! current version is %s, latest is %s", global.Version, rawLatest)

			return rawLatest
		case latest[idx] == current[idx]:
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

	evaluatedExe, err := filepath.EvalSymlinks(exe)
	log.Error(err, "an error occurred while locating location of original executable, perhaps a broken symlink")

	if runtime.GOOS == "windows" {
		return evaluatedExe
	}

	home, err := os.UserHomeDir()
	homePath := fmt.Sprintf("%s/.local/bin/pap", home)

	log.Error(err, "an error occurred while getting the user's home directory")

	if evaluatedExe != "/usr/bin" && evaluatedExe != homePath {
		log.Warn("it seems like you installed pap in a location not specified by the install guide (%s)", evaluatedExe)
		log.Warn("if this is expected, you can ignore this.")
		log.Continue("would like to continue?")
	}

	return evaluatedExe
}

func Update() {
	version := checkIfNewUpdate()

	log.Log("finding exe...")

	exe := getExePath()
	url := fmt.Sprintf(
		"https://github.com/talwat/pap/releases/download/v%s/pap_%s_%s_%s",
		version,
		version,
		runtime.GOOS,
		runtime.GOARCH,
	)
	tmpPath := fmt.Sprintf("/tmp/pap-update-%s", version)

	net.Download(url, tmpPath, fmt.Sprintf("pap %s", version), nil, fs.ExecutePerm)
	log.Log("installing pap...")
	fs.MoveFile(tmpPath, exe)

	log.Success("done!")
}
