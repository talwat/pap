package forge_test

import (
	"sort"
	"testing"

	"github.com/talwat/pap/internal/jarfiles/forge"
)

//nolint:lll
func TestURL(t *testing.T) {
	t.Parallel()

	installer := forge.GetURL("1.6.4", "", false)
	want := "https://maven.minecraftforge.net/net/minecraftforge/forge/1.6.4-9.11.1.1345/forge-1.6.4-9.11.1.1345-installer.jar"

	if installer != want {
		t.Errorf(`GetURL("1.6.4", "", false) = %s; want %s`, installer, want)
	}

	installer = forge.GetURL("1.9", "", true)
	want = "https://maven.minecraftforge.net/net/minecraftforge/forge/1.9-12.16.1.1938-1.9.0/forge-1.9-12.16.1.1938-1.9.0-installer.jar"

	if installer != want {
		t.Errorf(`GetURL("1.9", "", true) = %s; want %s`, installer, want)
	}

	installer = forge.GetURL("1.7.10_pre4", "", true)
	want = "https://maven.minecraftforge.net/net/minecraftforge/forge/1.7.10_pre4-10.12.2.1149-prerelease/forge-1.7.10_pre4-10.12.2.1149-prerelease-installer.jar"

	if installer != want {
		t.Errorf(`GetURL("1.7.10_pre4", "", true) = %s; want %s`, installer, want)
	}

	installer = forge.GetURL("1.19.3", "44.1.0", false)
	want = "https://maven.minecraftforge.net/net/minecraftforge/forge/1.19.3-44.1.0/forge-1.19.3-44.1.0-installer.jar"

	if installer != want {
		t.Errorf(`forge.GetURL("1.19.3", "44.1.0", false) = %s; want %s`, installer, want)
	}
}

func TestSort(t *testing.T) {
	t.Parallel()

	version1 := forge.MinecraftVersion{
		Major: 1,
		Minor: 16,
		Patch: 5,
	}
	version2 := forge.MinecraftVersion{
		Major: 1,
		Minor: 6,
		Patch: 4,
	}
	version3 := forge.MinecraftVersion{
		Major: 1,
		Minor: 19,
		Patch: 3,
	}
	version4 := forge.MinecraftVersion{
		Major:             1,
		Minor:             7,
		Patch:             10,
		IsPrerelease:      true,
		PrereleaseVersion: 4,
	}
	versions := []forge.MinecraftVersion{version1, version2, version3, version4}

	sort.Sort(forge.ByVersion(versions))

	correctVersion := 3

	if versions[correctVersion].Minor != version3.Minor {
		t.Errorf(`sort.Sort(forge.ByVersion(vs))[3].Minor = %d; want %d`, versions[correctVersion].Minor, 19)
	}
}

//nolint:lll
func TestBuildURL(t *testing.T) {
	t.Parallel()

	minecraft := forge.MinecraftVersion{
		Major: 1,
		Minor: 9,
		Patch: 0,
	}

	installer := forge.InstallerVersion{
		Version: "12.16.1.1938",
	}

	url := forge.BuildURL(&minecraft, &installer)
	want := "https://maven.minecraftforge.net/net/minecraftforge/forge/1.9-12.16.1.1938-1.9.0/forge-1.9-12.16.1.1938-1.9.0-installer.jar"

	if url != want {
		t.Errorf(`forge.BuildUrl(&minecraft, &installer) = %s; want %s`, url, want)
	}

	minecraft = forge.MinecraftVersion{
		Major: 1,
		Minor: 16,
		Patch: 5,
	}

	installer = forge.InstallerVersion{
		Version: "36.2.39",
	}

	url = forge.BuildURL(&minecraft, &installer)
	want = "https://maven.minecraftforge.net/net/minecraftforge/forge/1.16.5-36.2.39/forge-1.16.5-36.2.39-installer.jar"

	if url != want {
		t.Errorf(`forge.BuildUrl(&minecraft, &installer) = %s; want %s`, url, want)
	}

	minecraft = forge.MinecraftVersion{
		Major:             1,
		Minor:             7,
		Patch:             10,
		IsPrerelease:      true,
		PrereleaseVersion: 4,
	}

	installer = forge.InstallerVersion{
		Version: "10.12.2.1149",
	}

	url = forge.BuildURL(&minecraft, &installer)
	want = "https://maven.minecraftforge.net/net/minecraftforge/forge/1.7.10_pre4-10.12.2.1149-prerelease/forge-1.7.10_pre4-10.12.2.1149-prerelease-installer.jar"

	if url != want {
		t.Errorf(`forge.BuildUrl(&minecraft, &installer) = %s; want %s`, url, want)
	}
}
