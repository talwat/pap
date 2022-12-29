// CLI Commands
package cmd

import (
	"regexp"

	"github.com/talwat/pap/internal/log"
)

func ValidateOption(value string, pattern string, name string) {
	match, err := regexp.MatchString(pattern, value)
	log.Error(err, "an error occurred while verifying %s", name)

	if !match {
		log.RawError("%s '%s' is not valid", name, value)
	}
}
