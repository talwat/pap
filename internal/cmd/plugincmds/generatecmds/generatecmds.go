package generatecmds

import (
	"encoding/json"
	"fmt"

	"github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/plugins/paplug"
)

func WritePlugin(plugin paplug.PluginInfo) {
	unmarshaled, err := json.MarshalIndent(plugin, "", "    ")
	log.Error(err, "an error occurred while converting plugin back into json")

	log.Log("writing %s...", plugin.Name)
	fs.WriteFileByte(fmt.Sprintf("%s.json", plugin.Name), unmarshaled, fs.ReadWritePerm)

	log.Success("generated %s!", plugin.Name)
}
