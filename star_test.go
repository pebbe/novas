package novas

import (
	"fmt"
	"time"
)

func ExampleApp_Star() {
	t := Date(2008, 4, 24, 10, 36, 18, 0, time.UTC)
	star := NewStar("GMB 1830", "FK6", 1307, 11.88299133, 37.71867646, 4003.27, -5815.07, 109.21, -98.8)
	data := star.App(t)
	fmt.Printf("Right ascension:    %10.6f\n", data.RA)
	fmt.Printf("Declination:        %10.6f\n", data.Dec)
	fmt.Printf("Distance:           %10.6f\n", data.Dis)
	fmt.Printf("Ecliptic longitude: %10.6f\n", data.ELon)
	fmt.Printf("Ecliptic latitude:  %10.6f\n", data.ELat)
	// Output:
	// Right ascension:     11.891551
	// Declination:         37.658636
	// Distance:                  NaN
	// Ecliptic longitude: 161.574754
	// Ecliptic latitude:   33.477283
}

func ExampleTopo_Star() {
	t := Date(2008, 4, 24, 10, 36, 18, 0, time.UTC)
	star := NewStar("GMB 1830", "FK6", 1307, 11.88299133, 37.71867646, 4003.27, -5815.07, 109.21, -98.8)
	place := NewPlace(42.0, -70, 0, 18, 1010)
	data := star.Topo(t, place, REFR_NONE)
	fmt.Printf("Distance: %10.6f\n", data.Dis)
	fmt.Printf("Altitude: %10.6f\n", data.Alt)
	fmt.Printf("Azimuth:  %10.6f\n", data.Az)
	// Output:
	// Distance:        NaN
	// Altitude:   4.808569
	// Azimuth:  318.528197
}
