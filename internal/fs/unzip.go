package fs

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/talwat/pap/internal/log"
)

// This function is a bit of a complicated mess, and if you can improve it please do.
// Zip operations in go are very complex, and this function might be replaces with a simple invoking of tar.
//
//nolint:wrapcheck,goerr113,funlen,cyclop
func Unzip(src, dest string) {
	zipReader, err := zip.OpenReader(src)
	log.Error(err, "an error occurred while opening zip reader")

	defer func() {
		if err := zipReader.Close(); err != nil {
			log.Error(err, "a critical error occurred while closing zip reader")
		}
	}()

	MakeDirectory(dest)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(file *zip.File) error {
		readCloser, err := file.Open()
		if err != nil {
			return err
		}

		defer func() {
			if err := readCloser.Close(); err != nil {
				panic(err)
			}
		}()

		//nolint:gosec // Checked by very next statement
		path := filepath.Join(dest, file.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if file.FileInfo().IsDir() {
			MakeDirectory(path)
		} else {
			MakeDirectory(filepath.Dir(path))
			file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := file.Close(); err != nil {
					panic(err)
				}
			}()

			//nolint:gosec // This file is trusted, so we don't need to worry about a decompression bomb.
			_, err = io.Copy(file, readCloser)
			if err != nil {
				return err
			}
		}

		return nil
	}

	for _, f := range zipReader.File {
		err := extractAndWriteFile(f)
		if err != nil {
			log.Error(err, "an error occurred while extracting zip file")
		}
	}
}
