package novas_test

import (
	"github.com/pebbe/novas"

	"testing"
)

func TestEphInfo(t *testing.T) {
	expected := "JPL ephemeris DE405. Start JD = 2305424.50  End JD = 2525008.50"
	got := novas.EphInfo().String()
	if expected != got {
		t.Error("Expected %q, got %q", expected, got)
	}
}
