package main

// Management of the server.properties file

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func WritePropertiesFile(filename string, props map[string]interface{}) {
	keys := make([]string, 0)
	final := fmt.Sprintf("#Minecraft server properties\n#%s\n", MinecraftDateNow())

	for k := range props {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		final += fmt.Sprintf("%s=%s\n", k, props[k])
	}

	err := os.WriteFile(filename, []byte(final), ReadWritePerm)
	Error(err, "an error occurred while writing properties file")
}

func ReadPropertiesFile(filename string) map[string]interface{} {
	conf := map[string]interface{}{}

	if len(filename) == 0 {
		return conf
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

		val := ""

		if len(line) > equal {
			val = strings.TrimSpace(line[equal+1:])
		}

		conf[key] = val
	}

	err = scanner.Err()
	Error(err, "an error occurred while parsing the properties file")

	return conf
}

func EditProperty(prop string, val string) {
	Log("reading server properties...")

	props := ReadPropertiesFile("server.properties")

	Log("editing server properties...")

	props[prop] = val

	Log("writing server properties...")

	WritePropertiesFile("server.properties", props)

	Log("done")
}

func ResetProperties() {
	Log("this command is expected to be used with the latest minecraft version")
	Log("if you are using an older version, please manually delete the properties file and run the server")
	Continue("are you sure you would like to reset your server.properties file?")
	Download(
		"https://raw.githubusercontent.com/talwat/pap/main/assets/default.server.properties",
		"server.properties",
		"server properties file",
	)
}

func GetProperty(prop string) interface{} {
	props := ReadPropertiesFile("server.properties")

	val := props[prop]

	if val == nil {
		CustomError("property %s does not exist", prop)
	}

	return val
}
