package log_test

import (
	"testing"

	"github.com/talwat/pap/internal/log"
)

func TestLog(t *testing.T) {
	t.Parallel()

	log.Log("test")
	log.Warn("test")
	log.Success("test")
	log.Error(nil, "test")
	log.RawLog("test\n")
	log.OutputLog("test")
}
