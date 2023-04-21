// Logging and user input.
package log

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/log/color"
)

// Checks if err != nil and if so exits.
func Error(err error, msg string, params ...interface{}) {
	if err != nil {
		RawError("%s: %s", fmt.Sprintf(msg, params...), err)
		os.Exit(1)
	}
}

// Like Error but spits out a newline before.
func NewlineBeforeError(err error, msg string, params ...interface{}) {
	if err != nil {
		RawLog("\n")
		RawError("%s: %s", fmt.Sprintf(msg, params...), err)
		os.Exit(1)
	}
}

// Basic log message including the pap: prefix.
func Log(msg string, params ...interface{}) {
	RawLog("pap: %s\n", fmt.Sprintf(msg, params...))
}

// Like Log but without a newline at the end.
func NoNewline(msg string, params ...interface{}) {
	RawLog("pap: %s", fmt.Sprintf(msg, params...))
}

// Basically just fmt.Fprintf.
// Use this function in case pap decides to add a log file or whatever.
func RawLog(msg string, params ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, params...)
}

// Like RawLog, but prints to stdout instead.
//
// Note: This function outputs a trailing newline.
func OutputLog(msg string, params ...interface{}) {
	fmt.Fprintf(os.Stdout, "%s\n", fmt.Sprintf(msg, params...))
}

// Prints out an error message and exits regardless.
func RawError(msg string, params ...interface{}) {
	Log("%serror%s: %s", color.BrightRed, color.Reset, fmt.Sprintf(msg, params...))
	os.Exit(1)
}

// Prints out a warning.
func Warn(msg string, params ...interface{}) {
	Log("%swarning%s: %s", color.Yellow, color.Reset, fmt.Sprintf(msg, params...))
}

// Prints out a debug message.
// Debug info should be internal info that makes it easier to debug pap.
// Usually the information outputted here is completely useless to the end user.
func Debug(msg string, params ...interface{}) {
	if global.Debug {
		Log("%sdebug%s: %s", color.Magenta, color.Reset, fmt.Sprintf(msg, params...))
	}
}

// Like Debug, but with a newline before.
func NewlineBeforeDebug(msg string, params ...interface{}) {
	if global.Debug {
		RawLog("\n")
		Log("%sdebug%s: %s", color.Magenta, color.Reset, fmt.Sprintf(msg, params...))
	}
}

// Prints out a success message.
// Whenever a "major" operation finishes, you can use this.
func Success(msg string, params ...interface{}) {
	Log("%ssuccess%s: %s", color.Green, color.Reset, fmt.Sprintf(msg, params...))
}

// Scans standard input until it reaches a newline.
func RawScan() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	Error(err, "an error occurred while reading input")

	return strings.TrimSpace(text)
}

// Scan but it also handles the -y flag and has some nice looking logs.
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

// A yes or no prompt.
// Different from Continue because it won't exit if you say no.
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

// A continue prompt. If the user puts out anything that isn't "y" (case insensitive), then it will exit.
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
