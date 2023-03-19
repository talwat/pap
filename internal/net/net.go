// Networking
package net

import (
	"context"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"io/fs"
	"net/http"

	"github.com/talwat/gobar"
	papfs "github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/log"
)

func DoRequest(url string, notFoundMsg string) *http.Response {
	log.Debug("making a new request to %s", url)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	log.Error(err, "an error occurred while making request")

	userAgent := fmt.Sprintf("talwat/pap/%s", global.Version)
	req.Header.Set("User-Agent", userAgent)

	log.Debug("using user-agent %s", userAgent)
	log.Debug("doing request to %s", url)

	resp, err := http.DefaultClient.Do(req)
	log.Error(err, "an error occurred while sending request")

	if resp.StatusCode == http.StatusNotFound {
		log.RawError("404: %s (%s)", notFoundMsg, url)
	}

	log.Debug("status code %d", resp.StatusCode)

	return resp
}

// Like get, but just returns plaintext.
func GetPlainText(url string, notFoundMsg string) (string, int) {
	resp := DoRequest(url, notFoundMsg)

	log.Debug("reading response body...")

	raw, err := io.ReadAll(resp.Body)
	log.Error(err, "an error occurred while reading request body")

	defer resp.Body.Close()

	return string(raw), resp.StatusCode
}

// Saves the decoded JSON data to the value of content.
func Get(url string, notFoundMsg string, content interface{}) int {
	resp := DoRequest(url, notFoundMsg)

	log.Debug("decoding json response")

	err := json.NewDecoder(resp.Body).Decode(&content)
	log.Error(err, "an error occurred while decoding response")

	defer resp.Body.Close()

	return resp.StatusCode
}

// Set hash to nil in order to disable checksumming.
func Download(
	url string,
	notFoundMsg string,
	filename string,
	filedesc string,
	hash hash.Hash,
	perms fs.FileMode,
) []byte {
	resp := DoRequest(url, notFoundMsg)
	defer resp.Body.Close()

	file := papfs.CreateFile(filename, perms)
	defer file.Close()

	log.Debug("content length: %d", resp.ContentLength)
	bar := gobar.NewBar(
		0,
		resp.ContentLength,
		fmt.Sprintf("pap: downloading %s", filedesc),
	)

	var err error

	if hash == nil {
		_, err = io.Copy(io.MultiWriter(file, bar), resp.Body)
	} else {
		_, err = io.Copy(io.MultiWriter(file, bar, hash), resp.Body)
	}

	log.Error(err, "An error occurred while writing %s", filedesc)

	if hash == nil {
		return nil
	}

	return hash.Sum(nil)
}

// Download a file without keeping the hash or making a loading bar.
func SimpleDownload(url string, notFoundMsg string, filename string, fileDesc string) {
	log.Log("downloading %s...", fileDesc)

	resp := DoRequest(url, notFoundMsg)
	defer resp.Body.Close()

	file := papfs.CreateFile(filename, papfs.ReadWritePerm)
	defer file.Close()

	log.Debug("reading response body...")

	_, err := io.Copy(file, resp.Body)

	log.Error(err, "An error occurred while writing %s", fileDesc)
}
