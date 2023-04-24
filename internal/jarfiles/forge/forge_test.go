package forge_test

import (
	"sort"
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

	iv = forge.GetURL("1.19.3", "44.1.0", false)
	want = "https://maven.minecraftforge.net/net/minecraftforge/forge/1.19.3-44.1.0/forge-1.19.3-44.1.0-installer.jar"
	
	if iv != want {
		t.Errorf(`forge.GetURL("1.19.3", "44.1.0", false) = %s; want %s`, iv, want)
	}
}

func TestSort(t *testing.T) {
	t.Parallel()

	v1 := forge.MinecraftVersion{
		Major: 1,
		Minor: 16,
		Patch: 5,
	}
	v2 := forge.MinecraftVersion{
		Major: 1,
		Minor: 6,
		Patch: 4,
	}
	v3 := forge.MinecraftVersion{
		Major: 1,
		Minor: 19,
		Patch: 3,
	}
	v4 := forge.MinecraftVersion{
		Major:             1,
		Minor:             7,
		Patch:             10,
		IsPrerelease:      true,
		PrereleaseVersion: 4,
	}
	vs := []forge.MinecraftVersion{v1, v2, v3, v4}

	sort.Sort(forge.ByVersion(vs))

	if vs[3].Minor != v3.Minor {
		t.Errorf(`sort.Sort(forge.ByVersion(vs))[3].Minor = %d; want %d`, vs[3].Minor, 19)
	}
}
