// Logging and such
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Error(err error, message string, params ...interface{}) {
	if err != nil {
		Log("error: %s: %s", fmt.Sprintf(message, params...), err)
		os.Exit(1)
	}
}

func Log(message string, params ...interface{}) {
	RawLog(("pap: %s\n"), fmt.Sprintf(message, params...))
}

func LogNoNewline(message string, params ...interface{}) {
	RawLog(("pap: %s"), fmt.Sprintf(message, params...))
}

//nolint:forbidigo
func RawLog(message string, params ...interface{}) {
	fmt.Printf(message, params...)
}

func CustomError(message string, params ...interface{}) {
	Log("error: %s", fmt.Sprintf(message, params...))
	os.Exit(1)
}

func Warn(message string, params ...interface{}) {
	Log("warning: %s", fmt.Sprintf(message, params...))
}

func RawScan() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	Error(err, "an error occurred while reading input")

	return strings.TrimSpace(text)
}

func Scan(defaultValue string, prompt string, params ...interface{}) string {
	LogNoNewline("%s (default %s): ", fmt.Sprintf(prompt, params...), defaultValue)

	if AssumeDefaultInput {
		RawLog("\n")
		Log("continuing with value %s because assume-default is turned on", defaultValue)

		return defaultValue
	}

	input := RawScan()

	if input == "" {
		return defaultValue
	}

	return input
}

func YesOrNo(defaultValue string, prompt string, params ...interface{}) bool {
	LogNoNewline("%s [y/n]: ", fmt.Sprintf(prompt, params...))

	if AssumeDefaultInput {
		RawLog("\n")
		Log("choosing [%s] because assume-default is turned on", defaultValue)

		return true
	}

	input := strings.ToLower(RawScan())

	if input == "" {
		input = defaultValue
	}

	return input == "y"
}

func Continue(prompt string, params ...interface{}) {
	LogNoNewline("%s [y/n]: ", fmt.Sprintf(prompt, params...))

	if AssumeDefaultInput {
		RawLog("\n")
		Log("continuing because assume-default is turned on")

		return
	}

	input := strings.ToLower(RawScan())

	if input != "y" {
		Log("aborting...")
		os.Exit(1)
	}
}
