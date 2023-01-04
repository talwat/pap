package purpur_test

import (
	"testing"

	"github.com/talwat/pap/internal/jarfiles/purpur"
)

func TestLatestBuild(t *testing.T) {
	t.Parallel()

	version := purpur.GetSpecificVersion("1.14.2")

	if version.Version != "1.14.2" {
		t.Errorf(`GetSpecificVersion("1.14.2") = %s; want 1.14.2`, version.Version)
	}

	build := purpur.GetLatestBuild(version)

	if build.Build != "126" {
		t.Errorf(`GetLatetstBuild(version) = %s; want 126`, build.Build)
	}
}

func TestSpecificBuild(t *testing.T) {
	t.Parallel()

	version := purpur.GetSpecificVersion("1.14.2")

	if version.Version != "1.14.2" {
		t.Errorf(`GetSpecificVersion("1.14.2") = %s; want 1.14.2`, version.Version)
	}

	build := purpur.GetSpecificBuild(version, "124")

	if build.Build != "124" {
		t.Errorf(`GetSpecificBuild("1.14.2", "124") = %s; want 124`, build.Build)
	}
}

func TestGetURL(t *testing.T) {
	t.Parallel()

	url, build := purpur.GetURL("1.14.2", "124")

	expected := "https://api.purpurmc.org/v2/purpur/1.14.2/124/download"

	if url != expected || build.Build != "124" {
		t.Errorf(`GetURL("1.14.2", "124") = "%s", %s; want "%s", 124`, url, build.Build, expected)
	}
}
