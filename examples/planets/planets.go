package main

import (
	"github.com/pebbe/novas"

	"fmt"
)

func main() {

	jpleph := "/my/opt/novas/share/JPLEPH"

	longitude, latitude := 6.5670, 53.21853 // Groningen, The Netherlands

	novas.Init(jpleph, false)

	t := novas.Now()
	fmt.Println(t)

	geo := novas.NewPlace(latitude, longitude, 0, 20, 1010)

	fmt.Println("\n            Distance   Altitude   Azimuth")
	for _, i := range []novas.PlanetType{
		novas.MERCURY,
		novas.VENUS,
		novas.MARS,
		novas.JUPITER,
		novas.SATURN,
		novas.URANUS,
		novas.NEPTUNE,
		novas.PLUTO,
	} {
		obj := novas.NewPlanet(i)
		data := obj.Topo(t, geo, novas.REFR_NONE)
		fmt.Printf("%-8s%12.6f%11.2f%10.2f\n", obj.Name(), data.Dis, data.Alt, data.Az)
	}
}
