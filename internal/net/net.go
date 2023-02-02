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
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	req.Header.Set("User-Agent", fmt.Sprintf("talwat/pap/%s", global.Version))
	log.Error(err, "an error occurred while making request")

	resp, err := http.DefaultClient.Do(req)
	log.Error(err, "an error occurred while sending request")

	if resp.StatusCode == http.StatusNotFound {
		log.RawError("404: %s (%s)", notFoundMsg, url)
	}

	return resp
}

// Like get, but just returns plaintext.
func GetPlainText(url string, notFoundMsg string) (string, int) {
	resp := DoRequest(url, notFoundMsg)

	raw, err := io.ReadAll(resp.Body)
	log.Error(err, "an error occurred while reading request body")

	defer resp.Body.Close()

	return string(raw), resp.StatusCode
}

// Saves the decoded JSON data to the value of content.
func Get(url string, notFoundMsg string, content interface{}) int {
	resp := DoRequest(url, notFoundMsg)

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

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perms)
	log.Error(err, "an error occurred while opening %s", filename)

	defer file.Close()

	bar := newLoadingBar(
		resp.ContentLength,
		fmt.Sprintf("pap: downloading %s", fileDesc),
	)

	if hash == nil {
		_, err = io.Copy(io.MultiWriter(file, bar), resp.Body)
	} else {
		_, err = io.Copy(io.MultiWriter(file, bar, hash), resp.Body)
	}

	log.Error(err, "An error occurred while writing %s", fileDesc)

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

	file, err := os.Create(filename)
	log.Error(err, "an error occurred while opening %s", filename)

	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	log.Error(err, "An error occurred while writing %s", fileDesc)
	log.Log("done downloading")
}
