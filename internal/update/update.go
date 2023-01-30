package update

import (
	"fmt"
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
	rawLatest, statusCode := net.GetPlainText("https://raw.githubusercontent.com/talwat/pap/self-updater/version.txt")

	if statusCode != 200 {
		log.RawError("http request to get latest version returned %s", statusCode)
	}

	latest := parseVersion(rawLatest)
	current := parseVersion(global.Version)

	for i := len(latest) - 1; i >= 0; i-- {
		switch {
		case latest[i] > current[i]:
			log.Log("out of date! current version is %s, latest is %s", global.Version, rawLatest)
			return rawLatest
		case latest[i] == current[i]:
			continue
		}
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
	log.Error(err, "an error occured while getting the user's error")

	if evaluatedExe != "/usr/bin" || evaluatedExe != fmt.Sprintf("%s/.local/bin/pap", home) {
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
	url := fmt.Sprintf("https://github.com/talwat/pap/releases/download/v%s/pap_%s_macos_%s", version, version, runtime.GOARCH)
	tmpPath := fmt.Sprintf("/tmp/pap-update-%s", version)

	net.Download(url, tmpPath, fmt.Sprintf("pap %s", version), nil)
	log.Log("installing pap...")
	fs.MoveFile(tmpPath, exe)

	log.Success("done!")
}
