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

func GetLatestBuild(paperVersion string) PaperBuildStruct {
	var paperBuilds PaperBuildsStruct

	Log("getting latest build information")

	url := fmt.Sprintf("https://api.papermc.io/v2/projects/paper/versions/%s/builds", paperVersion)
	statusCode := Get(url, &paperBuilds)

	if paperBuilds.Error != "" {
		CustomError("api returned an error with status code %d: %s", statusCode, paperBuilds.Error)
	}

	// latest build, can be experimental or stable
	latestBuild := paperBuilds.Builds[len(paperBuilds.Builds)-1]

	if ExperimentalBuildInput {
		return latestBuild
	}

	// iterate through paperBuilds.Builds backwards
	for i := len(paperBuilds.Builds) - 1; i >= 0; i-- {
		if paperBuilds.Builds[i].Channel == "default" { // default = stable
			return paperBuilds.Builds[i]
		}
	}

	Continue("warning: no stable build found, would you like to use the latest experimental build?")

	return latestBuild
}

func GetBuild(paperVersion string, paperBuildID string) PaperBuildStruct {
	var paperBuild PaperBuildStruct

	var statusCode int

	var url string

	if paperBuildID == "latest" {
		paperBuild = GetLatestBuild(paperVersion)
	} else {
		url = fmt.Sprintf("https://api.papermc.io/v2/projects/paper/versions/%s/builds/%s", paperVersion, PaperBuildInput)
		statusCode = Get(url, &paperBuild)

		if paperBuild.Error != "" {
			CustomError("api returned an error with status code %d: %s", statusCode, paperBuild.Error)
		}
	}

	if paperBuild.Channel == "experimental" && !ExperimentalBuildInput {
		Continue("warning: build %d has been flagged as experimental, are you sure you would like to download it?", paperBuild.Build)
	}

	return paperBuild
}
