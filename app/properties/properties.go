// Management of the server.properties file
package properties

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/talwat/pap/app/fs"
	"github.com/talwat/pap/app/log"
	"github.com/talwat/pap/app/net"
	"github.com/talwat/pap/app/time"
)

func WritePropertiesFile(filename string, props map[string]interface{}) {
	keys := make([]string, 0)
	final := fmt.Sprintf("#Minecraft server properties\n#%s\n", time.MinecraftDateNow())

	for k := range props {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		final += fmt.Sprintf("%s=%s\n", k, props[k])
	}

	err := os.WriteFile(filename, []byte(final), fs.ReadWritePerm)
	log.Error(err, "an error occurred while writing properties file")
}

func ReadPropertiesFile(filename string) map[string]interface{} {
	conf := map[string]interface{}{}

	if len(filename) == 0 {
		return conf
	}

	file, err := os.Open(filename)
	log.Error(err, "an error occurred while opening properties file")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		equal := strings.Index(line, "=")

		if equal < 0 {
			continue
		}

		key := strings.TrimSpace(line[:equal])

		if len(key) == 0 {
			continue
		}

		val := ""

		if len(line) > equal {
			val = strings.TrimSpace(line[equal+1:])
		}

		conf[key] = val
	}

	err = scanner.Err()
	log.Error(err, "an error occurred while parsing the properties file")

	return conf
}

func EditProperty(prop string, val string) {
	log.Log("reading server properties...")

	props := ReadPropertiesFile("server.properties")

	log.Log("editing server properties...")

	props[prop] = val

	log.Log("writing server properties...")

	WritePropertiesFile("server.properties", props)

	log.Log("done")
}

func ResetProperties() {
	log.Log("this command is expected to be used with the latest minecraft version")
	log.Log("if you are using an older version, please manually delete the properties file and run the server")
	log.Continue("are you sure you would like to reset your server.properties file?")
	net.Download(
		"https://raw.githubusercontent.com/talwat/pap/main/assets/default.server.properties",
		"server.properties",
		"server properties file",
	)
}

func GetProperty(prop string) interface{} {
	props := ReadPropertiesFile("server.properties")

	val := props[prop]

	if val == nil {
		log.CustomError("property %s does not exist", prop)
	}

	return val
}
