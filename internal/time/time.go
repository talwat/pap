// Time related utilities.
package time

import "time"

//nolint:gochecknoglobals // It's a constant which I want to be accessible through this package.
var MinecraftTimeFormat = "Mon Jan 02 15:04:05 MST 2006"

// Gets the time now in minecraft's strange format.
func MinecraftDateNow() string {
	return time.Now().Format(MinecraftTimeFormat)
}
