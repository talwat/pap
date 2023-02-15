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
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
	papfs "github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/global"
	"github.com/talwat/pap/internal/log"
)

//nolint:gomnd // Nolint because most numbers are configuration options.
func newLoadingBar(maxBytes int64, desc string) *progressbar.ProgressBar {
	bar := progressbar.NewOptions64(
		maxBytes,
		progressbar.OptionSetDescription(desc),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionShowBytes(true),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionOnCompletion(func() {
			log.RawLog("\n")
		}),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionFullWidth(),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionSetPredictTime(false),

		progressbar.OptionEnableColorCodes(false),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "=",
			SaucerHead:    ">",
			AltSaucerHead: ">",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)

	return bar
}

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
	fileDesc string,
	hash hash.Hash,
	perms fs.FileMode,
) []byte {
	resp := DoRequest(url, notFoundMsg)
	defer resp.Body.Close()

	file := papfs.OpenFile(filename, perms)
	defer file.Close()

	bar := newLoadingBar(
		resp.ContentLength,
		fmt.Sprintf("pap: downloading %s", fileDesc),
	)

	log.Debug("reading response body...")

	var err error

	if hash == nil {
		log.Debug("hash is nil, not writing to hash")

		_, err = io.Copy(io.MultiWriter(file, bar), resp.Body)
	} else {
		_, err = io.Copy(io.MultiWriter(file, bar, hash), resp.Body)
	}

	log.Error(err, "An error occurred while writing %s", fileDesc)

	if hash == nil {
		return nil
	}

	log.Debug("calculating hash...")

	return hash.Sum(nil)
}

// Download a file without keeping the hash or making a loading bar.
func SimpleDownload(url string, notFoundMsg string, filename string, fileDesc string) {
	log.Log("downloading %s...", fileDesc)

	resp := DoRequest(url, notFoundMsg)
	defer resp.Body.Close()

	file := papfs.OpenFile(filename, papfs.ReadWritePerm)
	defer file.Close()

	log.Debug("reading response body...")

	_, err := io.Copy(file, resp.Body)

	log.Error(err, "An error occurred while writing %s", fileDesc)
	log.Log("done downloading")
}
