package fs

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/talwat/pap/internal/exec"
	"github.com/talwat/pap/internal/log"
)

const (
	Success  = 1
	Fail     = 2
	NotFound = 3
)

// Try using a specific command to unzip a file.
func tryUnzipCommand(program string, src string, dest string, cmd string, params ...interface{}) {
	log.Log("using %s to unzip %s to %s...", program, src, dest)

	exitCode := exec.Run(".", fmt.Sprintf(cmd, params...))

	log.RawLog("\n")

	if exitCode != 0 {
		log.RawError("%s failed with exit code: %d", program, exitCode)
	}
}

// Unzip using commands provided by the OS.
func commandUnzip(src string, dest string) int {
	switch {
	case exec.CommandExists("unzip"):
		tryUnzipCommand("unzip", src, dest, "unzip -o %s -d %s", src, dest)

		return Success
	case exec.CommandExists("7z"):
		tryUnzipCommand("7z", src, dest, "7z %s -vd %s", src, dest)

		return Success
	case exec.CommandExists("bsdtar"):
		tryUnzipCommand("bstdar", src, dest, "bsdtar -xvf %s -C %s", src, dest)

		return Success
	default:
		log.Log("using golang method to unzip %s...", src)

		return NotFound
	}
}

// Full golang implementation, refactoring and editing is needed.
// This function is avoided if possible, but kept just in case the user doesn't have basic utilities.
// I mean seriously, who doesn't have unzip?
//
//nolint:goerr113,funlen,wrapcheck // I have no idea how to shorten this mess.
func unsafeUnzip(src string, dest string) {
	zipReader, err := zip.OpenReader(src)
	log.Error(err, "an error occurred while opening zip reader")

	defer func() {
		err := zipReader.Close()
		log.Error(err, "a critical error occurred while closing zip reader")
	}()

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(zipfile *zip.File) error {
		readCloser, err := zipfile.Open()
		if err != nil {
			return err
		}

		defer func() {
			err := readCloser.Close()
			log.Error(err, "a critical error occurred while closing the readcloser")
		}()

		//nolint:gosec // Checked by very next statement
		path := filepath.Join(dest, zipfile.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		// If it's a directory that is specified, make the directory and then exit
		if zipfile.FileInfo().IsDir() {
			MakeDirectory(path)

			return nil
		}

		MakeDirectory(filepath.Dir(path))

		file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zipfile.Mode())
		if err != nil {
			return err
		}

		defer func() {
			err := file.Close()

			log.Error(err, "a critical error occurred while closing the file")
		}()

		//nolint:gosec // This file is trusted, so we don't need to worry about a decompression bomb.
		_, err = io.Copy(file, readCloser)
		if err != nil {
			return err
		}

		return nil
	}

	for _, f := range zipReader.File {
		err := extractAndWriteFile(f)
		log.Error(err, "an error occurred while extracting zip file")
	}
}

func Unzip(src string, dest string) {
	MakeDirectory(dest)

	status := commandUnzip(src, dest)

	if status == Success {
		return
	}

	unsafeUnzip(src, dest)
}
