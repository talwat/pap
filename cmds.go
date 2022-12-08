package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

var version = "0.1"

//nolint:gochecknoglobals
var (
	PaperVersionInput = "latest"
	PaperBuildInput   = "latest"
	XMSInput          = "2G"
	XMXInput          = "2G"
	JarInput          = "paper.jar"
)

func GlobalHelp() {
	RawLog(`pap help
commands:
  help        displays this menu
  version     displays the version
  run         generates a script to run the server
  download    downloads the latest paper.jar file
  sign        signs the eula
`)
	os.Exit(0)
}

func VersionCommand() {
	RawLog("pap %s %s\nby talwat\n", version, runtime.GOOS)
	os.Exit(0)
}

func RunCommand() {
	command := fmt.Sprintf("java -Xms%s -Xmx%s -jar %s --nogui", XMSInput, XMXInput, JarInput)

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
