// Time related utilities.
package time

import "time"

const MinecraftTimeFormat = "Mon Jan 02 15:04:05 MST 2006"

// Gets the time now in minecraft's strange format.
func MinecraftDateNow() string {
	return time.Now().Format(MinecraftTimeFormat)
}
