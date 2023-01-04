package jarfiles

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/talwat/pap/internal/log"
)

const Latest = "latest"

func VerifyJarfile(calculated []byte, proper string) {
	log.Log("verifying downloaded jarfile...")

	if checksum := hex.EncodeToString(calculated); checksum == proper {
		log.Log("checksums match!")
	} else {
		log.RawError(
			fmt.Sprintf("checksums (calculated: %s, proper: %s) don't match!",
				checksum,
				proper,
			),
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
