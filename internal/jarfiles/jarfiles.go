package jarfiles

import (
	"encoding/hex"
	"strings"

	"github.com/talwat/pap/internal/log"
)

const Latest = "latest"

func UnsupportedMessage() {
	log.Warn("because you are using a jarfile which is not by papermc, please do not use 'pap script' with --aikars")
}

func VerifyJarfile(calculated []byte, proper string) {
	log.Log("verifying downloaded jarfile...")

	if checksum := hex.EncodeToString(calculated); checksum == proper {
		log.Success("checksums match!")
	} else {
		log.RawError(
			"checksums (calculated: %s, proper: %s) don't match!",
			checksum,
			proper,
		)
	}
}

func APIError(err string, statusCode int) {
	if err != "" {
		log.RawError("api returned an error with status code %d: %s", statusCode, FormatErrorMessage(err))
	}
}

func FormatErrorMessage(msg string) string {
	return strings.ToLower(strings.TrimSuffix(msg, "."))
}
