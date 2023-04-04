package generatecmds

import (
	"encoding/json"
	"fmt"

	"github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/plugins/sources"
	"github.com/talwat/pap/internal/plugins/sources/paplug"
)

func Generate(getPluginInfo func(plugin string) paplug.PluginInfo, plugins []string) {
	if len(plugins) == 0 {
		log.RawError("you must specify plugins to generate")
	}

	log.Log("getting plugins to write...")

	pluginsToWrite := sources.GetManyPluginInfo(plugins, getPluginInfo)

	for _, plugin := range pluginsToWrite {
		WritePlugin(plugin)
	}

	log.Success("all plugins generated successfully!")
}

func WritePlugin(plugin paplug.PluginInfo) {
	unmarshaled, err := json.MarshalIndent(plugin, "", "    ")
	log.Error(err, "an error occurred while converting plugin back into json")

	if global.UseStdoutInput {
		log.OutputLog(string(unmarshaled))
	} else {
		log.Log("writing %s...", plugin.Name)
		fs.WriteFileByte(fmt.Sprintf("%s.json", plugin.Name), unmarshaled, fs.ReadWritePerm)
	}

	log.Success("generated %s!", plugin.Name)
}
