// Filesystem Management
package fs

import (
	"io/fs"
	"os"

	"github.com/talwat/pap/app/log"
)

const (
	ExecutePerm   fs.FileMode = 0o700
	ReadWritePerm fs.FileMode = 0o600
)

func WriteFile(name string, text string, perms fs.FileMode) {
	err := os.WriteFile(name, []byte(text), perms)
	log.Error(err, "an error occurred while writing %s", name)
}

func MakeDirectory(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	log.Error(err, "an error occurred while creating %s", path)
}
