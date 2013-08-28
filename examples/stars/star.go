package main

import (
	"github.com/pebbe/novas"

	"fmt"
	"time"
)

func main() {

	t := novas.Date(2008, 4, 24, 10, 36, 18, 0, time.UTC)

	star := novas.NewStar("GMB 1830", "FK6", 1307, 11.88299133, 37.71867646, 4003.27, -5815.07, 109.21, -98.8)
	data := star.App(t)

	fmt.Println("Results for star", star.Name(), "at", t.Local())

	fmt.Printf("Right ascension:    %10.6f\n", data.RA)
	fmt.Printf("Declination:        %10.6f\n", data.Dec)
	fmt.Printf("Ecliptic longitude: %10.6f\n", data.ELon)
	fmt.Printf("Ecliptic latitude:  %10.6f\n", data.ELat)

	place := novas.NewPlace(42.0, -70, 0, 18, 1010)

	data2 := star.Topo(t, place, novas.REFR_NONE)
	fmt.Printf("Altitude:           %10.6f\n", data2.Alt)
	fmt.Printf("Azimuth:            %10.6f\n", data2.Az)
}
