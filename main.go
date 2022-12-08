package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

var version = "0.2.0"

//nolint:funlen,exhaustruct
func main() {
	app := &cli.App{
		Name:        "pap",
		Usage:       "a helper for papermc",
		Version:     version,
		HideHelp:    true,
		HideVersion: true,
		CommandNotFound: func(ctx *cli.Context, command string) {
			CustomError("command not found: %s", command)
		},
		OnUsageError: func(ctx *cli.Context, err error, isSubcommand bool) error {
			CustomError("%s", err)

			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "download",
				Aliases: []string{"d"},
				Usage:   "Download a papermc jarfile",
				Action: func(cCtx *cli.Context) error {
					DownloadCommand()

					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "paper-version",
						Value:       "latest",
						Usage:       "The papermc version to download",
						Aliases:     []string{"version"},
						Destination: &PaperVersionInput,
					},
					&cli.StringFlag{
						Name:        "paper-build",
						Value:       "latest",
						Usage:       "The papermc build to download",
						Aliases:     []string{"build"},
						Destination: &PaperBuildInput,
					},
				},
			},
			{
				Name:    "script",
				Aliases: []string{"sc"},
				Usage:   "Generates a script to run the jarfile",
				Action: func(cCtx *cli.Context) error {
					ScriptCommand()

					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "xms",
						Value:       "2G",
						Usage:       "The value for xms in the run command",
						Destination: &XMSInput,
					},
					&cli.StringFlag{
						Name:        "xmx",
						Value:       "2G",
						Usage:       "The value for xmx in the run command",
						Destination: &XMXInput,
					},
					&cli.StringFlag{
						Name:        "jar",
						Value:       "paper.jar",
						Usage:       "The name for the server jarfile",
						Destination: &JarInput,
					},
					&cli.BoolFlag{
						Name:        "use-gui",
						Aliases:     []string{"gui"},
						Usage:       "Whether to use the GUI or not",
						Destination: &GUIInput,
					},
				},
			},
			{
				Name:    "sign",
				Aliases: []string{"si"},
				Usage:   "Signs the EULA",
				Action: func(cCtx *cli.Context) error {
					EulaCommand()

					return nil
				},
			},
			{
				Name:    "help",
				Aliases: []string{"h"},
				Usage:   "Help menu",
				Action:  cli.ShowAppHelp,
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Shows version",
				Action: func(cCtx *cli.Context) error {
					cli.ShowVersion(cCtx)

					return nil
				},
			},
		},
	}

	//nolint:errcheck
	app.Run(os.Args)
}
