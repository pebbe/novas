package main

import (
	"github.com/pebbe/novas"

	"fmt"
	"time"
)

const (
	sundip = float64(-0.8)
)

func main() {

	jpleph := "/my/opt/novas/share/JPLEPH"

	latitude, longitude := 53.21853, 6.5670 // Groningen, The Netherlands

	// END OF USER SETTINGS

	novas.Init(jpleph, false)

	now := novas.Now()
	fmt.Println(now)

	geo := novas.NewPlace(latitude, longitude, 0, 20, 1010)
	fmt.Println("\nLocation:", geo, "\n")

	moon := novas.Moon()

	fmt.Printf("Distance from earth's center:   %.0f km\n", moon.App(now).Dis * novas.AU)

	data := moon.Topo(now, geo, novas.REFR_NONE)
    fmt.Printf("Distance from surface location: %.0f km\n\n", data.Dis * novas.AU)
	fmt.Printf("Altitude: %7.3f°\nAzimuth:  %7.3f°\n\n", data.Alt, data.Az)

	fmt.Printf("Phase of the moon: %.2f°, illuminated: %.0f%%\n\n", novas.MoonPhase(now), novas.MoonDisc(now))

	fmt.Println("Next new moon:      ", novas.MoonPhaseNext(now, 0, 30 * time.Second))
	fmt.Println("Next first quarter: ", novas.MoonPhaseNext(now, 90, 30 * time.Second))
	fmt.Println("Next full moon:     ", novas.MoonPhaseNext(now, 180, 30 * time.Second))
	fmt.Println("Next third quarter: ", novas.MoonPhaseNext(now, 270, 30 * time.Second))
}
