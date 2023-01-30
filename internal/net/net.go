// Networking
package net

import (
	"context"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
	"github.com/talwat/pap/internal/log"
)

//nolint:gomnd // Nolint because most numbers are configuration options
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

// Like get, but just returns plaintext
func GetPlainText(url string) (string, int) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	log.Error(err, "an error occurred while making request")

	resp, err := http.DefaultClient.Do(req)
	log.Error(err, "an error occurred while sending request")

	raw, err := io.ReadAll(resp.Body)
	log.Error(err, "an error occurred while reading request body")

	defer resp.Body.Close()

	return string(raw), resp.StatusCode
}

// saves the decoded JSON data to the value of content.
func Get(url string, content interface{}) int {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	log.Error(err, "an error occurred while making request")

	resp, err := http.DefaultClient.Do(req)
	log.Error(err, "an error occurred while sending request")

	err = json.NewDecoder(resp.Body).Decode(&content)
	log.Error(err, "an error occurred while decoding response")

	defer resp.Body.Close()

	return resp.StatusCode
}

func Download(url string, filename string, fileDesc string, hash hash.Hash) []byte {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	log.Error(err, "an error occurred while making a new request")

	resp, err := http.DefaultClient.Do(req)
	log.Error(err, "an error occurred while sending an http request")

	defer resp.Body.Close()

	file, err := os.Create(filename)
	log.Error(err, "an error occurred while opening %s", filename)

	defer file.Close()

	bar := newLoadingBar(
		resp.ContentLength,
		fmt.Sprintf("pap: downloading %s", fileDesc),
	)

	_, err = io.Copy(io.MultiWriter(file, bar, hash), resp.Body)

	log.Error(err, "An error occurred while writing %s", fileDesc)

	return hash.Sum(nil)
}

// Download a file without keeping the hash or making a loading bar.
func SimpleDownload(url string, filename string, fileDesc string) {
	log.Log("downloading %s...", fileDesc)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	log.Error(err, "an error occurred while making a new request")

	resp, err := http.DefaultClient.Do(req)
	log.Error(err, "an error occurred while sending an http request")

	defer resp.Body.Close()

	file, err := os.Create(filename)
	log.Error(err, "an error occurred while opening %s", filename)

	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	log.Error(err, "An error occurred while writing %s", fileDesc)
	log.Log("done downloading")
}
