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
	log.Debug("writing to %s", name)

	err := os.WriteFile(name, text, perms)
	log.Error(err, "an error occurred while writing %s", name)
}

func ReadFile(name string) []byte {
	log.Debug("reading %s", name)

	raw, err := os.ReadFile(name)
	log.Error(err, "an error occurred while reading %s", name)

	return raw
}

func FileExists(filename string) bool {
	log.Debug("checking if %s exists", filename)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	} else if err != nil {
		log.Error(err, "an error occurred while checking if %s exists", filename)
	}

	return true
}

func MakeDirectory(path string) {
	log.Debug("making directory at %s", path)

	err := os.MkdirAll(path, os.ModePerm)
	log.Error(err, "an error occurred while creating %s", path)
}

func DeletePath(path string) {
	log.Debug("deleting %s", path)

	err := os.RemoveAll(path)
	log.Error(err, "an error occurred while deleting %s", path)
}

func MoveFile(oldpath string, newpath string) {
	log.Debug("moving %s to %s", oldpath, newpath)

	err := os.Rename(oldpath, newpath)
	log.Error(err, "an error occurred while moving %s to %s", oldpath, newpath)
}

func OpenFile(filename string, perms fs.FileMode) *os.File {
	log.Debug("opening %s...", filename)

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perms)
	log.Error(err, "an error occurred while opening %s", filename)

	return file
}
