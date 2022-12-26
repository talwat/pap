package cmd

import (
	"fmt"
	"runtime"

	"github.com/talwat/pap/app/fs"
	"github.com/talwat/pap/app/global"
	"github.com/talwat/pap/app/log"
	"github.com/urfave/cli/v2"
)

func ScriptCommand(cCtx *cli.Context) error {
	gui := " --nogui"

	if global.GUIInput {
		gui = ""
	}

	command := fmt.Sprintf("java -Xms%s -Xmx%s -jar %s%s", global.XMSInput, global.XMXInput, global.JarInput, gui)

	if runtime.GOOS == "windows" {
		fs.WriteFile("run.bat", fmt.Sprintf("@ECHO OFF\n%s\npause", command), fs.ExecutePerm)
	} else {
		fs.WriteFile("run.sh", fmt.Sprintf("#!/bin/sh\n%s", command), fs.ExecutePerm)
	}

	log.Log("generated shell script")
	log.Log("keep in mind, this script will not be the absolute most efficiencent it can be")
	log.Log("go to aikars flags (https://docs.papermc.io/paper/aikars-flags) for more information on optimizing flags and tuning java") //nolint:lll
	log.Log("or, if you're lazy, go to flags.sh (https://flags.sh/) for a generator")

	return nil
}
