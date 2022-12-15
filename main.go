package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

var version = "0.5.0"

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
			CustomError("command not found: %s", command)
		},
		OnUsageError: func(ctx *cli.Context, err error, isSubcommand bool) error {
			CustomError("%s", err)

			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "assume-default",
				Value:       false,
				Usage:       "assume the default answer in all prompts",
				Aliases:     []string{"y"},
				Destination: &AssumeDefaultInput,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "download",
				Aliases: []string{"d"},
				Usage:   "download a papermc jarfile",
				Action:  DownloadCommand,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "minecraft-version",
						Value:       "latest",
						Usage:       "the minecraft version to download",
						Aliases:     []string{"version"},
						Destination: &PaperVersionInput,
					},
					&cli.StringFlag{
						Name:        "paper-build",
						Value:       "latest",
						Usage:       "the papermc build to download",
						Aliases:     []string{"build"},
						Destination: &PaperBuildInput,
					},
					&cli.BoolFlag{
						Name:        "paper-experimental",
						Value:       false,
						Usage:       "takes the latest build regardless. also bypasses warning prompt",
						Aliases:     []string{"experimental"},
						Destination: &ExperimentalBuildInput,
					},
				},
			},
			{
				Name:    "geyser",
				Aliases: []string{"d"},
				Usage:   "downloads geyser",
				Action:  GeyserCommand,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "no-floodgate",
						Value:       false,
						Usage:       "do not download floodgate",
						Destination: &NoFloodGate,
					},
				},
			},
			{
				Name:    "script",
				Aliases: []string{"sc"},
				Usage:   "generate a script to run the jarfile",
				Action:  ScriptCommand,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "xms",
						Value:       "2G",
						Usage:       "the value for xms in the run command",
						Destination: &XMSInput,
					},
					&cli.StringFlag{
						Name:        "xmx",
						Value:       "2G",
						Usage:       "the value for xmx in the run command",
						Destination: &XMXInput,
					},
					&cli.StringFlag{
						Name:        "jar",
						Value:       "paper.jar",
						Usage:       "the name for the server jarfile",
						Destination: &JarInput,
					},
					&cli.BoolFlag{
						Name:        "use-gui",
						Aliases:     []string{"gui"},
						Usage:       "whether to use the GUI or not",
						Destination: &GUIInput,
					},
				},
			},
			{
				Name:    "sign",
				Aliases: []string{"si"},
				Usage:   "sign the EULA",
				Action:  EulaCommand,
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
						Action:  EditPropertyCommand,
					},
					{
						Name:    "get",
						Aliases: []string{"g"},
						Usage:   "get property",
						Action:  GetPropertyCommand,
					},
					{
						Name:    "reset",
						Aliases: []string{"r"},
						Usage:   "downloads the default server.properties",
						Action:  ResetPropertiesCommand,
					},
				},
			},
		},
	}

	//nolint:errcheck
	app.Run(os.Args)
}
