package novas

import (
	"fmt"
	"time"
)

func ExampleApp_Planet() {
	t := Date(2012, 12, 13, 12, 0, 0, 0, time.UTC)
	sun := Sun()
	data := sun.App(t)
	fmt.Printf("Right ascension:    %10.6f\n", data.RA)
	fmt.Printf("Declination:        %10.6f\n", data.Dec)
	fmt.Printf("Distance in AU:     %10.6f\n", data.Dis)
	fmt.Printf("Ecliptic longitude: %10.6f\n", data.ELon)
	fmt.Printf("Ecliptic latitude:  %10.6f\n", data.ELat)
	// Output:
	// Right ascension:     17.411478
	// Declination:        -23.187681
	// Distance in AU:       0.984424
	// Ecliptic longitude: 261.890295
	// Ecliptic latitude:    0.001756
}

func ExampleTopo_Planet() {
	t := Date(2012, 12, 13, 12, 0, 0, 0, time.UTC)
	sun := Sun()
	place := NewPlace(53.21853, 6.5670, 0, 18, 1010)
	data := sun.Topo(t, place, REFR_NONE)
	fmt.Printf("Distance in AU: %10.6f\n", data.Dis)
	fmt.Printf("Altitude:       %10.6f\n", data.Alt)
	fmt.Printf("Azimuth:        %10.6f\n", data.Az)
	// Output:
	// Distance in AU:   0.984414
	// Altitude:        13.278250
	// Azimuth:        187.524224
}
