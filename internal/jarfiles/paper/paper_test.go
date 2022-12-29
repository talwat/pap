package paper_test

import (
	"testing"

	"github.com/talwat/pap/internal/jarfiles/paper"
)

func TestLatestBuild(t *testing.T) {
	t.Parallel()

	build := paper.GetLatestBuild("1.12")

	if build.Build != 1169 {
		t.Errorf(`GetLatetstBuild("1.12") = %d; want 1169`, build.Build)
	}

	build = paper.GetLatestBuild("1.13")

	if build.Build != 173 {
		t.Errorf(`GetLatetstBuild("1.12") = %d; want 173`, build.Build)
	}
}

func TestSpecificBuild(t *testing.T) {
	t.Parallel()

	build := paper.GetSpecificBuild("1.12", "1160")

	if build.Build != 1160 {
		t.Errorf(`GetSpecificBuild("1.12", "1160") = %d; want 1160`, build.Build)
	}
}

func TestGetBuild(t *testing.T) {
	t.Parallel()

	build := paper.GetBuild("1.12", "1160")

	if build.Build != 1160 {
		t.Errorf(`GetSpecificBuild("1.12", "1160") = %d; want 1160`, build.Build)
	}
}

func TestFormatErrorMessage(t *testing.T) {
	t.Parallel()

	msg := paper.FormatErrorMessage("This is a test message.")

	if msg != "this is a test message" {
		t.Errorf(`FormatErrorMessage("This is a test message.") = "%s"; want "this is a test message"`, msg)
	}
}

func TestGetURL(t *testing.T) {
	t.Parallel()

	url, build := paper.GetURL("1.12", "1160")

	expected := "https://api.papermc.io/v2/projects/paper/versions/1.12/builds/1160/downloads/paper-1.12-1160.jar"

	if url != expected || build.Build != 1160 {
		t.Errorf(`GetURL("1.12", "1160") = "%s", %d; want "%s", 1160`, url, build.Build, expected)
	}
}
