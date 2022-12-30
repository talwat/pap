package official_test

import (
	"testing"

	"github.com/talwat/pap/internal/jarfiles/official"
)

func TestGetURL(t *testing.T) {
	t.Parallel()

	url, pkg := official.GetURL("1.12")

	expected := "https://launcher.mojang.com/v1/objects/8494e844e911ea0d63878f64da9dcc21f53a3463/server.jar"
	expectedSha1 := "8494e844e911ea0d63878f64da9dcc21f53a3463"

	if url != expected {
		t.Errorf(`GetURL("1.12") = "%s"; want "%s"`, url, expected)
	} else if pkg.Downloads.Server.SHA1 != expectedSha1 {
		t.Errorf(`GetURL("1.12").sha1 = "%s"; want "%s"`, pkg.Downloads.Server.SHA1, expectedSha1)
	}
}

func TestGetPackage(t *testing.T) {
	t.Parallel()

	pkg := official.GetPackage("1.12")

	expected := "1.12"

	if pkg.ID != expected {
		t.Errorf(`GetPackage("1.12") = "%s"; want "%s"`, pkg.ID, expected)
	}
}
