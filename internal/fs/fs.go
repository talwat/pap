// Filesystem Management
package fs

import (
	"io/fs"
	"os"

	"github.com/talwat/pap/internal/log"
)

const (
	ExecutePerm   fs.FileMode = 0o700
	ReadWritePerm fs.FileMode = 0o600
)

func WriteFile(name string, text string, perms fs.FileMode) {
	WriteFileByte(name, []byte(text), perms)
}

func WriteFileByte(name string, text []byte, perms fs.FileMode) {
	err := os.WriteFile(name, text, perms)
	log.Error(err, "an error occurred while writing %s", name)
}

func ReadFile(name string) []byte {
	raw, err := os.ReadFile(name)
	log.Error(err, "an error occurred while reading %s", name)

	return raw
}

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	} else if err != nil {
		log.Error(err, "an error occurred while checking if %s exists", filename)
	}

	return true
}

func MakeDirectory(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	log.Error(err, "an error occurred while creating %s", path)
}

func DeletePath(path string) {
	err := os.RemoveAll(path)
	log.Error(err, "an error occurred while deleting %s", path)
}

func MoveFile(oldpath string, newpath string) {
	err := os.Rename(oldpath, newpath)
	log.Error(err, "an error occurred while moving %s to %s", oldpath, newpath)
}
