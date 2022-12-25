package main

import (
	"fmt"
	"strings"
)

func GetLatestVersion() string {
	var versions PaperVersions

	Log("getting latest version information")
	Get("https://api.papermc.io/v2/projects/paper", &versions)

	return versions.Versions[len(versions.Versions)-1]
}

func FormatErrorMessage(msg string) string {
	return strings.ToLower(strings.TrimSuffix(msg, "."))
}

func GetLatestBuild(version string) PaperBuild {
	var builds PaperBuilds

	Log("getting latest build information")

	url := fmt.Sprintf("https://api.papermc.io/v2/projects/paper/versions/%s/builds", version)
	status := Get(url, &builds)

	if builds.Error != "" {
		CustomError("api returned an error with status code %d: %s", status, FormatErrorMessage(builds.Error))
	}

	// latest build, can be experimental or stable
	latest := builds.Builds[len(builds.Builds)-1]

	if ExperimentalBuildInput {
		return latest
	}

	// iterate through builds.Builds backwards
	for i := len(builds.Builds) - 1; i >= 0; i-- {
		if builds.Builds[i].Channel == "default" { // default = stable
			return builds.Builds[i] // stable build found, return it
		}
	}

	Continue("warning: no stable build found, would you like to use the latest experimental build?")

	return latest
}

func GetSpecificBuild(version string, buildID string) PaperBuild {
	Log("getting build information for %s", buildID)

	var build PaperBuild

	url := fmt.Sprintf("https://api.papermc.io/v2/projects/paper/versions/%s/builds/%s", version, buildID)
	statusCode := Get(url, &build)

	if build.Error != "" {
		CustomError("api returned an error with status code %d: %s", statusCode, FormatErrorMessage(build.Error))
	}

	return build
}

func GetBuild(version string, buildID string) PaperBuild {
	var build PaperBuild

	if buildID == "latest" {
		build = GetLatestBuild(version)
	} else {
		build = GetSpecificBuild(version, buildID)
	}

	if build.Channel == "experimental" && !ExperimentalBuildInput {
		Continue(
			"warning: build %d has been flagged as experimental, are you sure you would like to download it?",
			build.Build,
		)
	}

	return build
}
