package cmd

import (
	"fmt"

	"github.com/talwat/pap/app/fs"
	"github.com/talwat/pap/app/log"
	"github.com/talwat/pap/app/time"
	"github.com/urfave/cli/v2"
)

func EulaCommand(cCtx *cli.Context) error {
	fs.WriteFile("eula.txt", fmt.Sprintf(
		`#By changing the setting below to TRUE you are indicating your agreement to our EULA (https://aka.ms/MinecraftEULA).
#%s
#Signed by pap
eula=true`,
		time.MinecraftDateNow(),
	), fs.ReadWritePerm)
	log.Log("signed eula")

	return nil
}
