// Logging and user input.
package log

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/talwat/pap/app/global"
)

const (
	Reset         = "\x1B[0m"
	Bold          = "\x1B[1m"
	Dim           = "\x1B[2m"
	Italic        = "\x1B[3m"
	URL           = "\x1B[4m"
	Blink         = "\x1B[5m"
	Blink2        = "\x1B[6m"
	Selected      = "\x1B[7m"
	Hidden        = "\x1B[8m"
	Strikethrough = "\x1B[9m"

	Black   = "\x1B[30m"
	Red     = "\x1B[31m"
	Green   = "\x1B[32m"
	Yellow  = "\x1B[33m"
	Blue    = "\x1B[34m"
	Magenta = "\x1B[35m"
	Cyan    = "\x1B[36m"
	White   = "\x1B[37m"

	BrightBlack   = "\x1B[30;1m"
	BrightRed     = "\x1B[31;1m"
	BrightGreen   = "\x1B[32;1m"
	BrightYellow  = "\x1B[33;1m"
	BrightBlue    = "\x1B[34;1m"
	BrightMagenta = "\x1B[35;1m"
	BrightCyan    = "\x1B[36;1m"
	BrightWhite   = "\x1B[37;1m"
)

func Error(err error, msg string, params ...interface{}) {
	if err != nil {
		RawError("%s: %s", fmt.Sprintf(msg, params...), err)
		os.Exit(1)
	}
}

func Log(msg string, params ...interface{}) {
	RawLog("pap: %s\n", fmt.Sprintf(msg, params...))
}

func NoNewline(msg string, params ...interface{}) {
	RawLog("pap: %s", fmt.Sprintf(msg, params...))
}

func RawLog(msg string, params ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, params...)
}

// Like RawLog, but prints to stdout instead.
// Note: This function outputs a trailing newline.
func OuptutLog(msg string, params ...interface{}) {
	fmt.Fprintf(os.Stdout, "%s\n", fmt.Sprintf(msg, params...))
}

func RawError(msg string, params ...interface{}) {
	Log("%serror%s: %s", Red, Reset, fmt.Sprintf(msg, params...))
	os.Exit(1)
}

func Warn(msg string, params ...interface{}) {
	Log("%swarning%s: %s", Yellow, Reset, fmt.Sprintf(msg, params...))
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
