package novas_test

import (
	"github.com/pebbe/novas"

	"testing"
)

func TestEphInfo(t *testing.T) {
	expected := "JPL ephemeris DE430. Start JD = 2287184.50  End JD = 2688976.50"
	got := novas.EphInfo().String()
	if expected != got {
		t.Errorf("Expected %q, got %q", expected, got)
	}
}
