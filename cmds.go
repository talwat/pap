package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/urfave/cli/v2"
)

//nolint:gochecknoglobals
var (
	AssumeDefaultInput     = false
	PaperVersionInput      = "latest"
	PaperBuildInput        = "latest"
	ExperimentalBuildInput = false
	XMSInput               = "2G"
	XMXInput               = "2G"
	JarInput               = "paper.jar"
	GUIInput               = false
)

func DownloadCommand(cCtx *cli.Context) error {
	verifyOptions()

	url := getURL()

	calculatedChecksum := download(url, "paper.jar")
	checksum(calculatedChecksum)

	return nil
}

func ScriptCommand(cCtx *cli.Context) error {
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

	return nil
}

func EulaCommand(cCtx *cli.Context) error {
	WriteFile("eula.txt", fmt.Sprintf(
		`#By changing the setting below to TRUE you are indicating your agreement to our EULA (https://aka.ms/MinecraftEULA).
#%s
#Signed by pap
eula=true`,
		MinecraftDateNow(),
	), 0o600)
	Log("signed eula")

	return nil
}

func EditPropertyCommand(cCtx *cli.Context) error {
	propertyInput := cCtx.Args().Get(0)
	valueInput := cCtx.Args().Tail()

	if propertyInput == "" {
		CustomError("property name is required")
	} else if len(valueInput) == 0 {
		CustomError("value is required")
	}

	EditProperty(cCtx.Args().Get(0), strings.Join(valueInput, " "))

	return nil
}

func GetPropertyCommand(cCtx *cli.Context) error {
	propertyInput := cCtx.Args().Get(0)

	if propertyInput == "" {
		CustomError("property name is required")
	}

	property := GetProperty(propertyInput)
	RawLog("%s\n", property)

	return nil
}

func ResetPropertiesCommand(cCtx *cli.Context) error {
	ResetProperties()

	return nil
}
