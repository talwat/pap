package main

import (
	"os"

	"github.com/talwat/pap/app/cmd"
	"github.com/talwat/pap/app/cmd/propcmds"
	"github.com/talwat/pap/app/global"
	"github.com/talwat/pap/app/log"
	"github.com/urfave/cli/v2"
)

var version = "0.8.0-beta"

//nolint:funlen,exhaustruct
func main() {
	app := &cli.App{
		Name:    "pap",
		Usage:   "a helper for papermc",
		Version: version,
		Authors: []*cli.Author{
			{
				Name: "talwat",
			},
		},
		HideHelp:    true,
		HideVersion: true,
		//nolint:lll
		CustomAppHelpTemplate: `NAME:
   {{template "helpNameTemplate" .}}

USAGE:
   {{if .UsageText}}{{wrap .UsageText 3}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}

VERSION:
   {{.Version}}{{if .Description}}

DESCRIPTION:
   {{template "descriptionTemplate" .}}{{end}}
{{- if len .Authors}}

AUTHOR{{template "authorsTemplate" .}}{{end}}{{if .VisibleCommands}}

COMMANDS:{{template "visibleCommandCategoryTemplate" .}}{{end}}{{if .VisibleFlagCategories}}

GLOBAL OPTIONS:{{template "visibleFlagCategoryTemplate" .}}{{else if .VisibleFlags}}

GLOBAL OPTIONS:{{template "visibleFlagTemplate" .}}{{end}}{{if .Copyright}}

COPYRIGHT:
   {{template "copyrightTemplate" .}}{{end}}
`,
		CommandNotFound: func(ctx *cli.Context, command string) {
			log.CustomError("command not found: %s", command)
		},
		OnUsageError: func(ctx *cli.Context, err error, isSubcommand bool) error {
			log.CustomError("%s", err)

			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "assume-default",
				Value:       false,
				Usage:       "assume the default answer in all prompts",
				Aliases:     []string{"y"},
				Destination: &global.AssumeDefaultInput,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "download",
				Aliases: []string{"d"},
				Usage:   "download a papermc jarfile",
				Action:  cmd.DownloadCommand,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "minecraft-version",
						Value:       "latest",
						Usage:       "the minecraft/paper version to download",
						Aliases:     []string{"version", "v"},
						Destination: &global.VersionInput,
					},
					&cli.StringFlag{
						Name:        "paper-build",
						Value:       "latest",
						Usage:       "the papermc build to download",
						Aliases:     []string{"build", "b"},
						Destination: &global.BuildInput,
					},
					&cli.BoolFlag{
						Name:        "paper-experimental",
						Value:       false,
						Usage:       "takes the latest build regardless. also bypasses warning prompt",
						Aliases:     []string{"experimental", "e"},
						Destination: &global.ExperimentalBuildInput,
					},
				},
			},
			{
				Name:    "geyser",
				Aliases: []string{"d"},
				Usage:   "downloads geyser",
				Action:  cmd.GeyserCommand,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "no-floodgate",
						Value:       false,
						Usage:       "do not download floodgate",
						Destination: &global.NoFloodGateInput,
					},
				},
			},
			{
				Name:    "script",
				Aliases: []string{"sc"},
				Usage:   "generate a script to run the jarfile",
				Action:  cmd.ScriptCommand,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "mem",
						Value:       "2G",
						Usage:       "the value for -Xms and -Xmx in the run command",
						Aliases:     []string{"memory", "m"},
						Destination: &global.MemoryInput,
					},
					&cli.BoolFlag{
						Name:        "aikars",
						Value:       false,
						Usage:       "whether to use aikars flags: https://docs.papermc.io/paper/aikars-flags",
						Aliases:     []string{"a"},
						Destination: &global.AikarsInput,
					},
					&cli.StringFlag{
						Name:        "jar",
						Value:       "paper.jar",
						Usage:       "the name for the server jarfile",
						Destination: &global.JarInput,
					},
					&cli.BoolFlag{
						Name:        "use-gui",
						Aliases:     []string{"gui"},
						Usage:       "whether to use the GUI or not",
						Destination: &global.GUIInput,
					},
				},
			},
			{
				Name:    "sign",
				Aliases: []string{"si"},
				Usage:   "sign the EULA",
				Action:  cmd.EulaCommand,
			},
			{
				Name:    "help",
				Aliases: []string{"h"},
				Usage:   "show help",
				Action:  cli.ShowAppHelp,
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "show version",
				Action: func(cCtx *cli.Context) error {
					cli.ShowVersion(cCtx)

					return nil
				},
			},
			{
				Name:      "properties",
				Aliases:   []string{"p"},
				Usage:     "manages the server.properties file",
				ArgsUsage: "[set|get] [property] [value]",
				Subcommands: []*cli.Command{
					{
						Name:    "set",
						Aliases: []string{"s"},
						Usage:   "set property",
						Action:  propcmds.EditPropertyCommand,
					},
					{
						Name:    "get",
						Aliases: []string{"g"},
						Usage:   "get property",
						Action:  propcmds.GetPropertyCommand,
					},
					{
						Name:    "reset",
						Aliases: []string{"r"},
						Usage:   "downloads the default server.properties",
						Action:  propcmds.ResetPropertiesCommand,
					},
				},
			},
		},
	}

	if app.Run(os.Args) != nil {
		os.Exit(1)
	}
}
