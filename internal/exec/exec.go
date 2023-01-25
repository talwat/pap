package exec

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/talwat/pap/internal/log"
)

func Run(workDir string, cmd string) {
	log.NoNewline("running command %s", cmd)

	args := strings.Split(cmd, " ")

	//nolint:gosec // gosec is complaining because of the custom args, but those custom args are necessary.
	cmdObj := exec.Command(args[0], args[1:]...)

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
