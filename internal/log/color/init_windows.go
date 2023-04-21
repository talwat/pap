package color

// Windows is weird, but this is supposed to enable color on the old fashioned CMD.

import (
	"os"

	"golang.org/x/sys/windows"
)

func init() {
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32

	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}
