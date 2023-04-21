package forge_test

import (
	"testing"

	"github.com/talwat/pap/internal/jarfiles/forge"
)

func TestLatestInstaller(t *testing.T) {
	t.Parallel()

	iv := forge.GetLatestInstaller("1.6.4")
	if iv.Version != "9.11.1.1345" {
		t.Errorf(`GetLatestInstaller("1.6.4") = %+v; want version 9.11.1.1345`, iv)
	}
}

func TestRecommendedInstaller(t *testing.T) {
	t.Parallel()

	iv := forge.GetRecommendedInstaller("1.8.9")
	if iv.Version != "11.15.1.2318" {
		t.Errorf(`GetRecommendedInstaller("1.8.9") = %+v; want version 11.15.1.2318`, iv)
	}
}

func TestInstaller(t *testing.T) {
	t.Parallel()

	iv := forge.GetInstaller("1.4.7")
	if iv.Type != "latest" {
		t.Errorf(`GetInstaller("1.4.7") = %+v; want version 6.6.2.534`, iv)
	}

	iv = forge.GetInstaller("1.12.2")
	if iv.Type != "recommended" {
		t.Errorf(`GetInstaller("1.12.2") = %+v; want recommended installer`, iv)
	}
}
