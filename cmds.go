package main

import (
	"fmt"
	"runtime"
	"time"
)

//nolint:gochecknoglobals
var (
	PaperVersionInput = "latest"
	PaperBuildInput   = "latest"
	XMSInput          = "2G"
	XMXInput          = "2G"
	JarInput          = "paper.jar"
	GUIInput          = false
)

func ScriptCommand() {
	gui := " --nogui"

	if GUIInput {
		gui = ""
	}

	command := fmt.Sprintf("java -Xms%s -Xmx%s -jar %s%s", XMSInput, XMXInput, JarInput, gui)

	if runtime.GOOS == "windows" {
		WriteFile("run.bat", command, 0o700)
	} else {
		WriteFile("run.sh", fmt.Sprintf("#!/bin/sh\n%s", command), 0o700)
	}

	Log("generated shell script.")
}

func EulaCommand() {
	WriteFile("eula.txt", fmt.Sprintf(
		`#By changing the setting below to TRUE you are indicating your agreement to our EULA (https://aka.ms/MinecraftEULA).
#%s
#Signed by pap
eula=true`,
		time.Now().Format("Mon Jan 02 15:04:05 MST 2006"),
	), 0o600)
	Log("signed eula")
}
