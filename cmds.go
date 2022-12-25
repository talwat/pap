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
	VersionInput           = "latest"
	BuildInput             = "latest"
	ExperimentalBuildInput = false
	NoFloodGateInput       = false
	XMSInput               = "2G"
	XMXInput               = "2G"
	JarInput               = "paper.jar"
	GUIInput               = false
)

func DownloadCommand(cCtx *cli.Context) error {
	ValidateDownloadOptions()

	url, build := GetURL(VersionInput, BuildInput)

	checksum := Download(url, "paper.jar", "paper jarfile")
	VerifyJarfile(checksum, build)

	return nil
}

func ScriptCommand(cCtx *cli.Context) error {
	gui := " --nogui"

	if GUIInput {
		gui = ""
	}

	command := fmt.Sprintf("java -Xms%s -Xmx%s -jar %s%s", XMSInput, XMXInput, JarInput, gui)

	if runtime.GOOS == "windows" {
		WriteFile("run.bat", fmt.Sprintf("@ECHO OFF\n%s\npause", command), ExecutePerm)
	} else {
		WriteFile("run.sh", fmt.Sprintf("#!/bin/sh\n%s", command), ExecutePerm)
	}

	Log("generated shell script")
	Log("keep in mind, this script will not be the absolute most efficiencent it can be")
	Log("go to aikars flags (https://docs.papermc.io/paper/aikars-flags) for more information on optimizing flags and tuning java") //nolint:lll
	Log("or, if you're lazy, go to flags.sh (https://flags.sh/) for a generator")

	return nil
}

func EulaCommand(cCtx *cli.Context) error {
	WriteFile("eula.txt", fmt.Sprintf(
		`#By changing the setting below to TRUE you are indicating your agreement to our EULA (https://aka.ms/MinecraftEULA).
#%s
#Signed by pap
eula=true`,
		MinecraftDateNow(),
	), ReadWritePerm)
	Log("signed eula")

	return nil
}

func EditPropertyCommand(cCtx *cli.Context) error {
	prop := cCtx.Args().Get(0)
	val := cCtx.Args().Tail()

	if prop == "" {
		CustomError("property name is required")
	} else if len(val) == 0 {
		CustomError("value is required")
	}

	EditProperty(prop, strings.Join(val, " "))

	return nil
}

func GetPropertyCommand(cCtx *cli.Context) error {
	prop := cCtx.Args().Get(0)

	if prop == "" {
		CustomError("property name is required")
	}

	val := GetProperty(prop)
	RawLog("%s\n", val)

	return nil
}

func ResetPropertiesCommand(cCtx *cli.Context) error {
	ResetProperties()

	return nil
}

func GeyserCommand(cCtx *cli.Context) error {
	MakeDirectory("plugins")

	//nolint:lll
	Download(
		"https://ci.opencollab.dev/job/GeyserMC/job/Geyser/job/master/lastSuccessfulBuild/artifact/bootstrap/spigot/build/libs/Geyser-Spigot.jar",
		"plugins/Geyser-Spigot.jar",
		"geyser",
	)

	if !NoFloodGateInput {
		//nolint:lll
		Download(
			"https://ci.opencollab.dev/job/GeyserMC/job/Floodgate/job/master/lastSuccessfulBuild/artifact/spigot/build/libs/floodgate-spigot.jar",
			"plugins/floodgate-spigot.jar",
			"floodgate",
		)
	}

	disableKeySigning := YesOrNo(
		"y",
		"floodgate and geyser do not support key signing yet, would you like to disable it (recommended)?",
	)

	if disableKeySigning {
		EditProperty("enforce-secure-profile", "false")
	}

	return nil
}
