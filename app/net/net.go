// Networking
package net

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
	"github.com/talwat/pap/app/fs"
	"github.com/talwat/pap/app/log"
)

//nolint:gomnd
func newLoadingBar(maxBytes int64, description string) *progressbar.ProgressBar {
	bar := progressbar.NewOptions64(
		maxBytes,
		progressbar.OptionSetDescription(description),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(10),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionOnCompletion(func() {
			log.RawLog("\npap: done downloading\n")
		}),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionFullWidth(),

		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "=",
			SaucerHead:    ">",
			AltSaucerHead: ">",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)

	err := bar.RenderBlank()
	log.Error(err, "an error occurred while rendering loading bar")

	return bar
}

// saves the decoded JSON data to the value of content.
func Get(url string, content interface{}) int {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	log.Error(err, "an error occurred while making request")

	resp, err := http.DefaultClient.Do(req)
	log.Error(err, "an error occurred while sending request to papermc api")

	err = json.NewDecoder(resp.Body).Decode(&content)
	log.Error(err, "an error occurred while decoding response")

	defer resp.Body.Close()

	return resp.StatusCode
}

func Download(url string, filename string, fileDesc string) []byte {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	log.Error(err, "an error occurred while making a new request")

	resp, err := http.DefaultClient.Do(req)
	log.Error(err, "an error occurred while sending an http request")

	defer resp.Body.Close()

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, fs.ReadWritePerm)
	log.Error(err, "an error occurred while opening %s", filename)

	defer file.Close()

	bar := newLoadingBar(
		resp.ContentLength,
		fmt.Sprintf("pap: downloading %s", fileDesc),
	)

	hash := sha256.New()
	_, err = io.Copy(io.MultiWriter(file, bar, hash), resp.Body)

	log.Error(err, "An error occurred while drawing loading bar")

	return hash.Sum(nil)
}
