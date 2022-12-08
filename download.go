package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/schollz/progressbar/v3"
)

//nolint:gochecknoglobals
var (
	PaperBuild   PaperBuildStruct
	PaperVersion string
)

func DownloadCommand() {
	verify()

	url := getURL()

	calculatedChecksum := download(url)
	checksum(calculatedChecksum)
	os.Exit(0)
}

func verify() {
	if PaperVersionInput == "latest" {
		return
	}

	match, err := regexp.MatchString(`^\d\.\d{1,2}(\.\d)?(-pre\d)?$`, PaperVersionInput)
	Error(err, "an error occurred while verifying version")
	if !match {
		CustomError("version %s is not valid", PaperVersionInput)
	}

	if PaperBuildInput == "latest" {
		return
	}

	match, err = regexp.MatchString(`^\d+$`, PaperBuildInput)
	Error(err, "an error occurred while verifying build")
	if !match {
		CustomError("build %s is not valid", PaperBuildInput)
	}
}

func checksum(rawCalculatedChecksum []byte) {
	Log("verifying downloaded jarfile...")

	calculatedChecksum := hex.EncodeToString(rawCalculatedChecksum)

	if calculatedChecksum == PaperBuild.Downloads.Application.Sha256 {
		Log("checksums match!")
	} else {
		CustomError(
			fmt.Sprintf("checksums (calculated: %s, proper: %s) don't match!",
				calculatedChecksum,
				PaperBuild.Downloads.Application.Sha256,
			),
		)
	}
}

func getURL() string {
	if PaperVersionInput == "latest" {
		PaperVersion = GetLatestVersion()
	} else {
		PaperVersion = PaperVersionInput
	}

	PaperBuild = GetBuild(PaperVersion, PaperBuildInput)

	LogNoNewline("using paper version %s and build %d\n", PaperVersion, PaperBuild.Build)

	//nolint:lll
	return fmt.Sprintf("https://api.papermc.io/v2/projects/paper/versions/%s/builds/%d/downloads/paper-%s-%d.jar", PaperVersion, PaperBuild.Build, PaperVersion, PaperBuild.Build)
}

func download(url string) []byte {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	Error(err, "an error occurred while making a new request")

	resp, err := http.DefaultClient.Do(req)
	Error(err, "an error occurred while sending an http request")

	defer resp.Body.Close()

	file, err := os.OpenFile("paper.jar", os.O_CREATE|os.O_WRONLY, 0o644)
	Error(err, "an error occurred while opening server jar file")

	defer file.Close()

	//nolint:gomnd
	newBar := func(maxBytes int64, description string) *progressbar.ProgressBar {
		bar := progressbar.NewOptions64(
			maxBytes,
			progressbar.OptionSetDescription(description),
			progressbar.OptionSetWriter(os.Stderr),
			progressbar.OptionShowBytes(true),
			progressbar.OptionSetWidth(10),
			progressbar.OptionThrottle(65*time.Millisecond),
			progressbar.OptionOnCompletion(func() {
				RawLog("\npap: done downloading\n")
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
		Error(err, "an error occurred while rendering loading bar")

		return bar
	}

	bar := newBar(
		resp.ContentLength,
		"pap: downloading",
	)

	hash := sha256.New()
	_, err = io.Copy(io.MultiWriter(file, bar, hash), resp.Body)

	Error(err, "An error occurred while drawing loading bar")

	return hash.Sum(nil)
}
