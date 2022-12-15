package main

import "time"

func MinecraftDateNow() string {
	return time.Now().Format("Mon Jan 02 15:04:05 MST 2006")
}
