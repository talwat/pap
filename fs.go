package main

import (
	"io/fs"
	"os"
)

func WriteFile(name string, text string, perms fs.FileMode) {
	err := os.WriteFile(name, []byte(text), perms)
	Error(err, "an error occurred while writing %s", name)
}
