// pap :)
package main

import (
	"os"

	"github.com/talwat/pap/internal/cmd"
	"github.com/talwat/pap/internal/cmd/downloadcmds"
	"github.com/talwat/pap/internal/cmd/plugincmds"
	"github.com/talwat/pap/internal/cmd/plugincmds/generatecmds"
	"github.com/talwat/pap/internal/cmd/propcmds"
	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/log"
	"github.com/urfave/cli/v2"
)

//nolint:funlen,exhaustruct,maintidx // Ignoring these issues because this file only serves to define commands.
func main() {
	app := &cli.App{
		Name:    "pap",
		Usage:   "a swiss army knife for minecraft servers",
		Version: global.Version,
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
			log.RawError("command not found: %s", command)
		},
		OnUsageError: func(ctx *cli.Context, err error, isSubcommand bool) error {
			log.RawError("%s", err)

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
			&cli.BoolFlag{
				Name:        "debug",
				Value:       false,
				Usage:       "print extra information for debugging or troubleshooting",
				Aliases:     []string{"d"},
				Destination: &global.Debug,
			},
		},
		Commands: []*cli.Command{
			{
				Name:      "download",
				Aliases:   []string{"d"},
				Usage:     "download a jarfile",
				ArgsUsage: "[type]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "minecraft-version",
						Value:       "latest",
						Usage:       "the minecraft version to download",
						Aliases:     []string{"version", "v"},
						Destination: &global.MinecraftVersionInput,
					},
				},
				Subcommands: []*cli.Command{
					{
						Name:    "paper",
						Aliases: []string{"pa"},
						Usage:   "download a paper jarfile",
						Action:  downloadcmds.DownloadPaperCommand,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "paper-build",
								Value:       "latest",
								Usage:       "the papermc build to download",
								Aliases:     []string{"build", "b"},
								Destination: &global.JarBuildInput,
							},
							&cli.BoolFlag{
								Name:        "paper-experimental",
								Value:       false,
								Usage:       "takes the latest build regardless",
								Aliases:     []string{"experimental", "e"},
								Destination: &global.PaperExperimentalBuildInput,
							},
							&cli.StringFlag{
								Name:        "minecraft-version",
								Value:       "latest",
								Usage:       "the minecraft version to download",
								Aliases:     []string{"version", "v"},
								Destination: &global.MinecraftVersionInput,
							},
						},
					},
					{
						Name:    "fabric",
						Aliases: []string{"f"},
						Usage:   "download a fabric jarfile",
						Action:  downloadcmds.DownloadFabricCommand,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "fabric-loader",
								Value:       "latest",
								Usage:       "the fabric loader version to use",
								Aliases:     []string{"loader", "l"},
								Destination: &global.FabricLoaderVersion,
							},
							&cli.StringFlag{
								Name:        "fabric-installer",
								Value:       "latest",
								Usage:       "the fabric installer version to use",
								Aliases:     []string{"installer", "i"},
								Destination: &global.FabricInstallerVersion,
							},
							&cli.BoolFlag{
								Name:        "fabric-loader-experimental",
								Value:       false,
								Usage:       "takes the latest loader version regardless",
								Aliases:     []string{"experimental", "e"},
								Destination: &global.FabricExperimentalLoaderVersion,
							},
							&cli.BoolFlag{
								Name:        "minecraft-snapshot",
								Value:       false,
								Usage:       "takes the latest snapshot instead of the latest release",
								Aliases:     []string{"snapshot", "s"},
								Destination: &global.UseSnapshotInput,
							},
							&cli.StringFlag{
								Name:        "minecraft-version",
								Value:       "latest",
								Usage:       "the minecraft version to download",
								Aliases:     []string{"version", "v"},
								Destination: &global.MinecraftVersionInput,
							},
						},
					},
					{
						Name:    "purpur",
						Aliases: []string{"pu"},
						Usage:   "download a purpur jarfile",
						Action:  downloadcmds.DownloadPurpurCommand,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "purpur-build",
								Value:       "latest",
								Usage:       "the papermc build to download",
								Aliases:     []string{"build", "b"},
								Destination: &global.JarBuildInput,
							},
							&cli.StringFlag{
								Name:        "minecraft-version",
								Value:       "latest",
								Usage:       "the minecraft version to download",
								Aliases:     []string{"version", "v"},
								Destination: &global.MinecraftVersionInput,
							},
						},
					},
					{
						Name:    "official",
						Aliases: []string{"o"},
						Usage:   "download an official mojang jarfile",
						Action:  downloadcmds.DownloadOfficialCommand,
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:        "minecraft-snapshot",
								Value:       false,
								Usage:       "takes the latest snapshot instead of the latest release",
								Aliases:     []string{"snapshot", "s"},
								Destination: &global.UseSnapshotInput,
							},
							&cli.StringFlag{
								Name:        "minecraft-version",
								Value:       "latest",
								Usage:       "the minecraft version to download",
								Aliases:     []string{"version", "v"},
								Destination: &global.MinecraftVersionInput,
							},
						},
					},
				},
			},
			{
				Name:    "geyser",
				Aliases: []string{"g"},
				Usage:   "downloads geyser",
				Action:  cmd.GeyserCommand,
			},
			{
				Name:      "plugin",
				Aliases:   []string{"pl"},
				Usage:     "manages plugins",
				ArgsUsage: "[install|uninstall] [plugin]",
				Subcommands: []*cli.Command{
					{
						Name:    "install",
						Aliases: []string{"i"},
						Usage:   "installs a plugin",
						Action:  plugincmds.InstallCommand,
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:        "no-deps",
								Value:       false,
								Usage:       "whether to install and calculate dependencies",
								Aliases:     []string{"no-dependencies"},
								Destination: &global.NoDepsInput,
							},
							&cli.BoolFlag{
								Name:        "install-optional-deps",
								Value:       false,
								Usage:       "whether to install optional dependencies",
								Aliases:     []string{"optional"},
								Destination: &global.InstallOptionalDepsInput,
							},
							&cli.BoolFlag{
								Name:        "plugin-experimental",
								Value:       false,
								Usage:       "takes the latest version regardless",
								Aliases:     []string{"experimental"},
								Destination: &global.PluginExperimentalInput,
							},
						},
					},
					{
						Name:    "uninstall",
						Aliases: []string{"u", "remove", "r"},
						Usage:   "get property",
						Action:  plugincmds.UninstallCommand,
					},
					{
						Name:    "info",
						Aliases: []string{"in"},
						Usage:   "get information about a plugin",
						Action:  plugincmds.InfoCommand,
					},
					{
						Name:    "generate",
						Aliases: []string{"gen", "g"},
						Usage:   "generate a plugin json file using a 3rd party plugin library",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:        "use-stdout",
								Aliases:     []string{"stdout", "s"},
								Usage:       "output to stdout instead of writing it to the disk",
								Destination: &global.UseStdoutInput,
							},
						},
						Subcommands: []*cli.Command{
							{
								Name:    "modrinth",
								Aliases: []string{"m"},
								Usage:   "generate a plugin json file using modrinth",
								Action:  generatecmds.GenerateModrinth,
							},
							{
								Name:    "spigotmc",
								Aliases: []string{"s"},
								Usage:   "generate a plugin json file using spigotmc",
								Action:  generatecmds.GenerateSpigotMC,
							},
							{
								Name:    "bukkit",
								Aliases: []string{"b"},
								Usage:   "generate a plugin json file using bukkit",
								Action:  generatecmds.GenerateBukkit,
							},
						},
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
						Usage:       "the name for the server jarfile",
						Aliases:     []string{"j"},
						Destination: &global.JarInput,
					},
					&cli.BoolFlag{
						Name:        "use-gui",
						Aliases:     []string{"gui"},
						Usage:       "whether to use the GUI or not",
						Destination: &global.GUIInput,
					},
					&cli.BoolFlag{
						Name:        "use-stdout",
						Aliases:     []string{"stdout", "s"},
						Usage:       "output to stdout instead of writing it to the disk",
						Destination: &global.UseStdoutInput,
					},
				},
			},
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "updates pap if there is a new version available",
				Action:  cmd.UpdateCommand,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "reinstall",
						Aliases:     []string{"r"},
						Usage:       "reinstalls even if pap is up to date",
						Destination: &global.ReinstallInput,
					},
				},
			},
			{
				Name:    "sign",
				Aliases: []string{"si", "eula", "e"},
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
				Aliases:   []string{"pr"},
				Usage:     "manages the server.properties file",
				ArgsUsage: "[set|get] [property] [value]",
				Subcommands: []*cli.Command{
					{
						Name:    "set",
						Aliases: []string{"s"},
						Usage:   "set property",
						Action:  propcmds.SetPropertyCommand,
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
