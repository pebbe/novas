package novas_test

import (
	"github.com/pebbe/novas"

	"fmt"
	"testing"
	"time"
)

func TestAppStar(t *testing.T) {
	nt := novas.Date(2008, 4, 24, 10, 36, 18, 0, time.UTC)
	star := novas.NewStar("GMB 1830", "FK6", 1307, 11.88299133, 37.71867646, 4003.27, -5815.07, 109.21, -98.8)
	data := star.App(nt)
	tests := [][3]string{
		{"Right ascension", "11.891551", fmt.Sprintf("%.6f", data.RA)},
		{"Declination", "37.658636", fmt.Sprintf("%.6f", data.Dec)},
		{"Distance", "NaN", fmt.Sprintf("%.6f", data.Dis)},
		{"Ecliptic longitude", "161.574754", fmt.Sprintf("%.6f", data.ELon)},
		{"Ecliptic latitude", "33.477283", fmt.Sprintf("%.6f", data.ELat)},
	}
	for _, test := range tests {
		if test[1] != test[2] {
			t.Errorf("%s: expected %s, got %s", test[0], test[1], test[2])
		}
	}
}

func TestTopoStar(t *testing.T) {
	nt := novas.Date(2008, 4, 24, 10, 36, 18, 0, time.UTC)
	star := novas.NewStar("GMB 1830", "FK6", 1307, 11.88299133, 37.71867646, 4003.27, -5815.07, 109.21, -98.8)
	place := novas.NewPlace(42.0, -70, 0, 18, 1010)
	data := star.Topo(nt, place, novas.REFR_NONE)
	tests := [][3]string{
		{"Distance", "NaN", fmt.Sprintf("%.6f", data.Dis)},
		{"Altitude", "4.808569", fmt.Sprintf("%.6f", data.Alt)},
		{"Azimuth", "318.528197", fmt.Sprintf("%.6f", data.Az)},
	}
	for _, test := range tests {
		if test[1] != test[2] {
			t.Errorf("%s: expected %s, got %s", test[0], test[1], test[2])
		}
	}
}
