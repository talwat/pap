package cmd

// The entire `script` command is defined here.
// It could if deemed useful be split up into several files,
// And put in it's own directory which is inside `internal`.

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/log"
	"github.com/urfave/cli/v2"
)

// Aikars defines aikars flags.
// See https://docs.papermc.io/paper/aikars-flags for more info.
//
//nolint:gochecknoglobals
var Aikars = []string{
	"-XX:+UseG1GC",
	"-XX:+ParallelRefProcEnabled",
	"-XX:MaxGCPauseMillis=200",
	"-XX:+UnlockExperimentalVMOptions",
	"-XX:+DisableExplicitGC",
	"-XX:+AlwaysPreTouch",
	"-XX:G1HeapWastePercent=5",
	"-XX:G1MixedGCCountTarget=4",
	"-XX:G1MixedGCLiveThresholdPercent=90",
	"-XX:G1RSetUpdatingPauseTimePercent=5",
	"-XX:SurvivorRatio=32",
	"-XX:+PerfDisableSharedMem",
	"-XX:MaxTenuringThreshold=1",
	"-Dusing.aikars.flags=https://mcflags.emc.gs",
	"-Daikars.new.flags=true",
}

// LargeMemFlags are used if allocated memory is bigger than 12 GB.
//
//nolint:gochecknoglobals
var LargeMemFlags = []string{
	"-XX:G1NewSizePercent=40",
	"-XX:G1MaxNewSizePercent=50",
	"-XX:G1HeapRegionSize=16M",
	"-XX:G1ReservePercent=15",
	"-XX:InitiatingHeapOccupancyPercent=20",
}

// SmallMemFlags are used if allocated memory is smaller than 12 GB.
//
//nolint:gochecknoglobals
var SmallMemFlags = []string{
	"-XX:G1NewSizePercent=30",
	"-XX:G1MaxNewSizePercent=40",
	"-XX:G1HeapRegionSize=8M",
	"-XX:G1ReservePercent=20",
	"-XX:InitiatingHeapOccupancyPercent=15",
}

func memInputToMegabytes(memInput string) int {
	switch {
	case strings.HasSuffix(global.MemoryInput, "G"): // Memory is specified in gigabytes (G)
		log.Debug("using gigabytes as memory unit")

		gigabytes, err := strconv.Atoi(strings.TrimSuffix(memInput, "G"))
		log.Error(err, "invalid memory amount")

		// How many megabytes are in one gigabyte
		const MBInGB = 1000

		megabytes := gigabytes * MBInGB
		log.Debug("memory amount in megabytes: %d", megabytes)

		return megabytes
	case strings.HasSuffix(global.MemoryInput, "M"): // Memory is specified in megabytes (M)
		log.Debug("using megabytes as memory unit")

		megabytes, err := strconv.Atoi(strings.TrimSuffix(memInput, "M"))
		log.Error(err, "invalid memory amount")

		log.Debug("memory amount in megabytes: %d", megabytes)

		return megabytes
	default:
		log.RawError("memory value does not end with M (megabytes) or G (gigabytes)")

		return 0
	}
}

func generateAikars() []string {
	flagsToUse := Aikars

	// Specified RAM in megabytes
	ram := memInputToMegabytes(global.MemoryInput)

	// What is considered a lot of ram
	const largeRAM = 12000

	if ram > largeRAM {
		log.Debug("there is more than 12G of ram, using large ram flags")

		flagsToUse = append(flagsToUse, LargeMemFlags...)
	} else {
		log.Debug("there is less than 12G of ram, using small ram flags")

		flagsToUse = append(flagsToUse, SmallMemFlags...)
	}

	return flagsToUse
}

func generateCommand() string {
	// Base includes the base flags
	base := []string{
		"-Xms" + global.MemoryInput,
		"-Xmx" + global.MemoryInput,
	}

	flagsToUse := base

	if global.AikarsInput {
		flagsToUse = append(flagsToUse, generateAikars()...)
	}

	flagsToUse = append(flagsToUse, fmt.Sprintf("-jar %s", global.JarInput))

	if !global.GUIInput {
		flagsToUse = append(flagsToUse, "--nogui")
	}

	return fmt.Sprintf("java %s", strings.Join(flagsToUse, " "))
}

func output(name string, text string) {
	if global.UseStdoutInput {
		log.OutputLog(text)
	} else {
		fs.WriteFile(name, text, fs.ExecutePerm)
	}

	log.Success("generated shell script!")
}

func ScriptCommand(cCtx *cli.Context) error {
	if global.JarInput == "" {
		log.RawError("the --jar option is required for the script command")
	}

	command := generateCommand()

	if runtime.GOOS == "windows" {
		output("run.bat", fmt.Sprintf("@ECHO OFF\n%s\npause", command))
	} else {
		output("run.sh", fmt.Sprintf("#!/bin/sh\n%s", command))
	}

	log.Log("go to aikars flags (https://docs.papermc.io/paper/aikars-flags) for more information on optimizing flags and tuning java") //nolint:lll

	return nil
}
