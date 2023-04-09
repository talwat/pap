package jenkins

import (
	"fmt"
	"regexp"

	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
	"github.com/talwat/pap/internal/plugins/sources/paplug"
)

func GetJenkinsURL(download paplug.Download) string {
	var jenkinsBuild Build

	log.Debug("getting jenkins build information...")

	url := fmt.Sprintf("%s/lastSuccessfulBuild/api/json", download.Job)

	net.Get(
		url,
		"jenkins build not found, please report this to https://github.com/talwat/pap/issues",
		&jenkinsBuild,
	)

	log.Debug("finding correct artifact...")

	for _, artifact := range jenkinsBuild.Artifacts {
		log.Debug("checking if %s matches %s...", artifact.FileName, download.Artifact)

		matched, err := regexp.MatchString(download.Artifact, artifact.FileName)
		log.Error(err, "an error occurred while checking if %s is the correct artifact", artifact.FileName)

		if matched {
			log.Debug("using %s", artifact.FileName)

			return fmt.Sprintf("%s/lastSuccessfulBuild/artifact/%s", download.Job, artifact.RelativePath)
		}
	}

	log.RawError("no artifacts matched, please report this to https://github.com/talwat/pap/issues")

	return ""
}
