package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Error(err error, message string, params ...interface{}) {
	if err != nil {
		Log("pap: %s: %s", fmt.Sprintf(message, params...), err)
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
	Log(message, params...)
	os.Exit(1)
}

func Scan(prompt string, defaultValue string) string {
	LogNoNewline("%s (default %s): ", prompt, defaultValue)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	Error(err, "an error occurred while reading input")

	trimmed := strings.TrimSpace(text)

	if trimmed == "" {
		return defaultValue
	}

	return trimmed
}
