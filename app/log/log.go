// Logging and user input.
package log

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/talwat/pap/app/global"
)

func Error(err error, msg string, params ...interface{}) {
	if err != nil {
		Log("error: %s: %s", fmt.Sprintf(msg, params...), err)
		os.Exit(1)
	}
}

func Log(msg string, params ...interface{}) {
	RawLog(("pap: %s\n"), fmt.Sprintf(msg, params...))
}

func NoNewline(msg string, params ...interface{}) {
	RawLog(("pap: %s"), fmt.Sprintf(msg, params...))
}

//nolint:forbidigo
func RawLog(msg string, params ...interface{}) {
	fmt.Printf(msg, params...)
}

func CustomError(msg string, params ...interface{}) {
	Log("error: %s", fmt.Sprintf(msg, params...))
	os.Exit(1)
}

func Warn(msg string, params ...interface{}) {
	Log("warning: %s", fmt.Sprintf(msg, params...))
}

func RawScan() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	Error(err, "an error occurred while reading input")

	return strings.TrimSpace(text)
}

func Scan(defaultVal string, prompt string, params ...interface{}) string {
	NoNewline("%s (default %s): ", fmt.Sprintf(prompt, params...), defaultVal)

	if global.AssumeDefaultInput {
		RawLog("\n")
		Log("continuing with value %s because assume-default is turned on", defaultVal)

		return defaultVal
	}

	input := RawScan()

	if input == "" {
		return defaultVal
	}

	return input
}

func YesOrNo(defaultVal string, prompt string, params ...interface{}) bool {
	NoNewline("%s [y/n]: ", fmt.Sprintf(prompt, params...))

	if global.AssumeDefaultInput {
		RawLog("\n")
		Log("choosing [%s] because assume-default is turned on", defaultVal)

		return true
	}

	input := strings.ToLower(RawScan())

	if input == "" {
		input = defaultVal
	}

	return input == "y"
}

func Continue(prompt string, params ...interface{}) {
	NoNewline("%s [y/n]: ", fmt.Sprintf(prompt, params...))

	if global.AssumeDefaultInput {
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
