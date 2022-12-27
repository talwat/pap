package cmd

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/talwat/pap/app/fs"
	"github.com/talwat/pap/app/global"
	"github.com/talwat/pap/app/log"
	"github.com/urfave/cli/v2"
)

func memInputToMegabytes(memInput string) int {
	switch {
	case strings.HasSuffix(global.MemoryInput, "G"): // Memory is specified in gigabytes (G)
		gigabytes, err := strconv.Atoi(strings.TrimSuffix(memInput, "G"))
		log.Error(err, "invalid memory amount")

		// How many megabytes are in one gigabyte
		const MBInGB = 1000

		return gigabytes * MBInGB
	case strings.HasSuffix(global.MemoryInput, "M"): // Memory is specified in megabytes (M)
		megabytes, err := strconv.Atoi(strings.TrimSuffix(memInput, "M"))
		log.Error(err, "invalid memory amount")

		return megabytes
	default:
		log.CustomError("memory value does not end with M (megabytes) or G (gigabytes)")

		return 0
	}
}

//nolint:funlen // Ignoring because most of the length comes from the flag definitions
func generateCommand() string {
	// Specified RAM in megabytes
	ram := memInputToMegabytes(global.MemoryInput)

	flags := []string{
		"-Xms" + global.MemoryInput,
		"-Xmx" + global.MemoryInput,
	}

	aikars := []string{
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

	// If allocated memory is bigger than 12 GB
	largeMemFlags := []string{
		"-XX:G1NewSizePercent=40",
		"-XX:G1MaxNewSizePercent=50",
		"-XX:G1HeapRegionSize=16M",
		"-XX:G1ReservePercent=15",
		"-XX:InitiatingHeapOccupancyPercent=20",
	}

	// If allocated memory is smaller than 12 GB
	smallMemFlags := []string{
		"-XX:G1NewSizePercent=30",
		"-XX:G1MaxNewSizePercent=40",
		"-XX:G1HeapRegionSize=8M",
		"-XX:G1ReservePercent=20",
		"-XX:InitiatingHeapOccupancyPercent=15",
	}

	// What is considered a lot of ram
	const largeRAM = 12000

	if ram > largeRAM {
		aikars = append(aikars, largeMemFlags...)
	} else {
		aikars = append(aikars, smallMemFlags...)
	}

	if global.AikarsInput {
		flags = append(flags, aikars...)
	}

	flags = append(flags, "-jar "+global.JarInput)

	if !global.GUIInput {
		flags = append(flags, "--nogui")
	}

	return fmt.Sprintf("java %s", strings.Join(flags, " "))
}

func ScriptCommand(cCtx *cli.Context) error {
	command := generateCommand()

	if runtime.GOOS == "windows" {
		fs.WriteFile("run.bat", fmt.Sprintf("@ECHO OFF\n%s\npause", command), fs.ExecutePerm)
	} else {
		fs.WriteFile("run.sh", fmt.Sprintf("#!/bin/sh\n%s", command), fs.ExecutePerm)
	}

	log.Log("generated shell script")
	log.Log("go to aikars flags (https://docs.papermc.io/paper/aikars-flags) for more information on optimizing flags and tuning java") //nolint:lll

	return nil
}
