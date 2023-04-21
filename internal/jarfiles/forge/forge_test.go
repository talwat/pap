package forge_test

import (
	"testing"

	"github.com/talwat/pap/internal/jarfiles/forge"
)

func TestLatestInstaller(t *testing.T) {
	t.Parallel()

	iv := forge.GetLatestInstaller("1.6.4")
	if iv.Version != "9.11.1.1345" {
		t.Errorf(`GetLatestInstaller("1.6.4") = %s; want 9.11.1.1345`, iv.Version)
	}
}

func TestRecommendedInstalle(t *testing.T) {
	t.Parallel()

	iv := forge.GetRecommendedInstaller("1.8.9")
	if iv.Version != "11.15.1.2318" {
		t.Errorf(`GetRecommendedInstaller("1.8.9") = %s; want 11.15.1.2318`, iv.Version)
	}
}