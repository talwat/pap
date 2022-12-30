package jarfiles

import (
	"encoding/hex"
	"fmt"

	"github.com/talwat/pap/internal/log"
)

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
