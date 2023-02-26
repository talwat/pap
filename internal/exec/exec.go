package exec

import (
	"bufio"
	"errors"
	"os/exec"
	"runtime"
	"strings"

	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/log"
)

func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)

	return !errors.Is(err, exec.ErrNotFound)
}

func Run(workDir string, cmd string) int {
	log.NoNewline("running command %s", cmd)

	var cmdObj *exec.Cmd

	if runtime.GOOS == "windows" {
		log.NewlineBeforeDebug("using powershell")

		cmdObj = exec.Command("powershell", "-command", cmd)
	} else {
		log.NewlineBeforeDebug("using sh")

		cmdObj = exec.Command("sh", "-c", cmd)
	}

	log.Debug("using working directory %s", workDir)
	cmdObj.Dir = workDir

	cmdReader, err := cmdObj.StdoutPipe()
	cmdObj.Stderr = cmdObj.Stdout

	log.NewlineBeforeError(err, "an error occurred while creating stdout pipe")

	err = cmdObj.Start()
	log.NewlineBeforeError(err, "an error occurred while starting command")

	output := ""
	scanner := bufio.NewScanner(cmdReader)

	for scanner.Scan() {
		output += scanner.Text()

		if global.Debug {
			log.RawLog("  %s\n", output)
		} else {
			log.RawLog(".")
		}
	}

	output = strings.TrimSpace(output)
	err = cmdObj.Wait()

	log.NewlineBeforeError(err, "an error occurred while running command. output:\n%s", output)

	return cmdObj.ProcessState.ExitCode()
}
