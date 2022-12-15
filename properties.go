// Management of the server.properties file
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func WritePropertiesFile(filename string, properties map[string]interface{}) {
	keys := make([]string, 0)
	final := fmt.Sprintf("#Minecraft server properties\n#%s\n", MinecraftDateNow())

	for k := range properties {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		final += fmt.Sprintf("%s=%s\n", k, properties[k])
	}

	err := os.WriteFile(filename, []byte(final), 0o600)
	Error(err, "an error occurred while writing properties file")
}

func ReadPropertiesFile(filename string) map[string]interface{} {
	config := map[string]interface{}{}

	if len(filename) == 0 {
		return config
	}

	file, err := os.Open(filename)
	Error(err, "an error occurred while opening properties file")

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

		value := ""

		if len(line) > equal {
			value = strings.TrimSpace(line[equal+1:])
		}

		config[key] = value
	}

	err = scanner.Err()
	Error(err, "an error occurred while parsing the properties file")

	return config
}

func EditProperty(property string, value string) {
	Log("reading server properties...")

	properties := ReadPropertiesFile("server.properties")

	Log("editing server properties...")

	properties[property] = value

	Log("writing server properties...")

	WritePropertiesFile("server.properties", properties)

	Log("done")
}

func ResetProperties() {
	Continue("are you sure you would like to reset your server.properties file?")
	Download(
		"https://raw.githubusercontent.com/talwat/pap/main/assets/default.server.properties",
		"server.properties",
		"server properties file",
	)
}

func GetProperty(propertyInput string) interface{} {
	properties := ReadPropertiesFile("server.properties")

	property := properties[propertyInput]

	if property == nil {
		CustomError("property %s does not exist", propertyInput)
	}

	return property
}
