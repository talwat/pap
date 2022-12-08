package main

import (
	"fmt"
)

func GetLatestVersion() string {
	var paperVersions PaperVersions

	Log("getting paper version information")
	Get("https://api.papermc.io/v2/projects/paper", &paperVersions)

	return paperVersions.Versions[len(paperVersions.Versions)-1]
}

func GetBuild(paperVersion string, paperBuildID string) PaperBuildStruct {
	var paperBuilds PaperBuildsStruct

	var paperBuild PaperBuildStruct

	var statusCode int

	var url string

	if paperBuildID == "latest" {
		Log("getting latest build information")

		url = fmt.Sprintf("https://api.papermc.io/v2/projects/paper/versions/%s/builds", paperVersion)
		statusCode = Get(url, &paperBuilds)

		if paperBuilds.Error != "" {
			CustomError("api returned an error with status code %d: %s", statusCode, paperBuilds.Error)
		}

		paperBuild = paperBuilds.Builds[len(paperBuilds.Builds)-1]
	} else {
		url = fmt.Sprintf("https://api.papermc.io/v2/projects/paper/versions/%s/builds/%s", paperVersion, PaperBuildInput)
		statusCode = Get(url, &paperBuild)

		if paperBuild.Error != "" {
			CustomError("api returned an error with status code %d: %s", statusCode, paperBuild.Error)
		}
	}

	return paperBuild
}
