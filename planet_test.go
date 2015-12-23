package novas_test

import (
	"github.com/pebbe/novas"

	"fmt"
	"testing"
	"time"
)

func TestAppPlanet(t *testing.T) {
	nt := novas.Date(2012, 12, 13, 12, 0, 0, 0, time.UTC)
	sun := novas.Sun()
	data := sun.App(nt)
	tests := [][3]string{
		{"Right ascension", "17.411478", fmt.Sprintf("%.6f", data.RA)},
		{"Declination", "-23.187681", fmt.Sprintf("%.6f", data.Dec)},
		{"Distance in AU", "0.984424", fmt.Sprintf("%.6f", data.Dis)},
		{"Ecliptic longitude", "261.890295", fmt.Sprintf("%.6f", data.ELon)},
		{"Ecliptic latitude", "0.001756", fmt.Sprintf("%.6f", data.ELat)},
	}
	for _, test := range tests {
		if test[1] != test[2] {
			t.Errorf("%s: expected %s, got %s", test[0], test[1], test[2])
		}
	}
}

func TestTopoPlanet(t *testing.T) {
	nt := novas.Date(2012, 12, 13, 12, 0, 0, 0, time.UTC)
	sun := novas.Sun()
	place := novas.NewPlace(53.21853, 6.5670, 0, 18, 1010)
	data := sun.Topo(nt, place, novas.REFR_NONE)
	tests := [][3]string{
		{"Distance in AU", "0.984414", fmt.Sprintf("%.6f", data.Dis)},
		{"Altitude", "13.278250", fmt.Sprintf("%.6f", data.Alt)},
		{"Azimuth", "187.524224", fmt.Sprintf("%.6f", data.Az)},
	}
	for _, test := range tests {
		if test[1] != test[2] {
			t.Errorf("%s: expected %s, got %s", test[0], test[1], test[2])
		}
	}
}

func TestDisc(t *testing.T) {
	tests := []string{
		"Moon 62",
		"Mercury 13",
		"Venus 48",
		"Mars 97",
		"Jupiter 100",
	}
	nt := novas.Date(2012, 4, 1, 0, 0, 0, 0, time.UTC)
	for i, n := range []*novas.Body{
		novas.Moon(),
		novas.Mercury(),
		novas.Venus(),
		novas.Mars(),
		novas.Jupiter(),
	} {
		got := fmt.Sprintf("%s %.0f", n.Name(), n.Disc(nt))
		if tests[i] != got {
			t.Errorf("expected %s, got %s", tests[i], got)
		}
	}
}
