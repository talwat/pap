package jarfiles_test

import (
	"encoding/hex"
	"testing"

	"github.com/talwat/pap/internal/jarfiles"
)

func TestFormatErrorMessage(t *testing.T) {
	t.Parallel()

	msg := jarfiles.FormatErrorMessage("This is a test message.")

	if msg != "this is a test message" {
		t.Errorf(`FormatErrorMessage("This is a test message.") = "%s"; want "this is a test message"`, msg)
	}
}

func TestAPIError(t *testing.T) {
	t.Parallel()

	jarfiles.APIError("", 0)
}

func TestVerifyJarfile(t *testing.T) {
	t.Parallel()

	sum := "54d626e08c1c802b305dad30b7e54a82f102390cc92c7d4db112048935236e9c"
	calculated, _ := hex.DecodeString(sum)
	jarfiles.VerifyJarfile(calculated, sum)
}
