package exec

import (
	"bufio"
	"os/exec"
	"runtime"
	"strings"

	"github.com/talwat/pap/internal/log"
)

func Run(workDir string, cmd string) {
	log.NoNewline("running command %s", cmd)

	var cmdObj *exec.Cmd

	if runtime.GOOS == "windows" {
		cmdObj = exec.Command("powershell", "-command", cmd)
	} else {
		cmdObj = exec.Command("sh", "-c", cmd)
	}

	cmdObj.Dir = workDir

	cmdReader, err := cmdObj.StdoutPipe()
	cmdObj.Stderr = cmdObj.Stdout

	log.NewlineBeforeError(err, "an error occurred while creating stdout pipe")

	err = cmdObj.Start()
	log.NewlineBeforeError(err, "an error occurred while starting command")

	output := ""
	scanner := bufio.NewScanner(cmdReader)

	for scanner.Scan() {
		output += scanner.Text() + "\n"

		log.RawLog(".")
	}

	output = strings.TrimSpace(output)
	err = cmdObj.Wait()

	log.NewlineBeforeError(err, "an error occurred while running command. output:\n%s", output)
}
