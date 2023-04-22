package forge_test

import (
	"testing"

	"github.com/talwat/pap/internal/jarfiles/forge"
)

func TestURL(t *testing.T) {
	t.Parallel()

	iv := forge.GetURL("1.6.4", "", false)
	want := "https://maven.minecraftforge.net/net/minecraftforge/forge/1.6.4-9.11.1.1345/forge-1.6.4-9.11.1.1345-installer.jar"
	if iv != want {
		t.Errorf(`GetURL("1.6.4", "", false) = %s; want %s`, iv, want)
	}

	iv = forge.GetURL("1.9", "", true)
	want = "https://maven.minecraftforge.net/net/minecraftforge/forge/1.9-12.16.1.1938/forge-1.9-12.16.1.1938-installer.jar"
	if iv != want {
		t.Errorf(`GetURL("1.9", "", true) = %s; want %s`, iv, want)
	}

	iv = forge.GetURL("1.7.10_pre4", "", true)
	want = "https://maven.minecraftforge.net/net/minecraftforge/forge/1.7.10_pre4-10.12.2.1149/forge-1.7.10_pre4-10.12.2.1149-installer.jar"
	if iv != want {
		t.Errorf(`GetURL("1.7.10_pre4", "", true) = %s; want %s`, iv, want)
	}
}
