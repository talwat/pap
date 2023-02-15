// Management of the server.properties file
package properties

import (
	"bufio"
	"fmt"
	"sort"
	"strings"

	"github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
	"github.com/talwat/pap/internal/time"
)

func WritePropertiesFile(filename string, props map[string]interface{}) {
	log.Debug("writing properties file...")

	keys := make([]string, 0)
	final := fmt.Sprintf("#Minecraft server properties\n#%s\n", time.MinecraftDateNow())

	for k := range props {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		final += fmt.Sprintf("%s=%s\n", k, props[k])
	}

	fs.WriteFile(filename, final, fs.ReadWritePerm)
}

func parsePropertiesLine(line string, conf map[string]interface{}) {
	equalIdx := strings.Index(line, "=")

	// If "=" is -1 (which means "=" isn't in the line), skip.
	if equalIdx == -1 {
		log.Debug("'%s' does not include an = sign, skipping...", line)

		return
	}

	// Set the key to everything before "=" using the equal index.
	key := strings.TrimSpace(line[:equalIdx])

	// If the key is empty, skip.
	if len(key) == 0 {
		log.Log("the key is empty, skiping...")

		return
	}

	val := ""

	// Check if there is anything after "=" in the line.
	if len(line) > equalIdx {
		// If there is, set it as the value.
		val = strings.TrimSpace(line[equalIdx+1:])
	}

	// Save the value to the key in the conf map.
	conf[key] = val

	log.Debug("parsed %s. %s=%s", line, key, val)
}

func ReadPropertiesFile(filename string) map[string]interface{} {
	log.Debug("reading properties file...")

	conf := map[string]interface{}{}

	if len(filename) == 0 {
		return conf
	}

	file := fs.OpenFile(filename, fs.ReadWritePerm)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parsePropertiesLine(line, conf)
	}

	err := scanner.Err()
	log.Error(err, "an error occurred while parsing the properties file")

	return conf
}

func SetProperty(prop string, val string) {
	log.Log("reading server properties...")

	props := ReadPropertiesFile("server.properties")

	log.Log("editing server properties...")

	props[prop] = val

	log.Log("writing server properties...")
	WritePropertiesFile("server.properties", props)

	log.Success("successfully set %s to %s", prop, val)
}

func ResetProperties() {
	log.Log("this command is expected to be used with the latest minecraft version")
	log.Log("if you are using an older version, please manually delete the properties file and run the server")
	log.Continue("are you sure you would like to reset your server.properties file?")
	net.SimpleDownload(
		"https://raw.githubusercontent.com/talwat/pap/main/assets/default.server.properties",
		"server properties file not found, please report this to https://github.com/talwat/pap/issues",
		"server.properties",
		"server properties file",
	)
	log.Success("successfully reset server properties file")
}

func GetProperty(prop string) interface{} {
	props := ReadPropertiesFile("server.properties")

	val := props[prop]

	if val == nil {
		log.RawError("property %s does not exist", prop)
	}

	return val
}
