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

	fmt.Printf("Distance from earth's center:   %.7f AU\n", novas.Sun.App(now).Dis)

	data := novas.Sun.Topo(now, geo, novas.REFR_NONE)

	fmt.Printf("Distance from surface location: %.7f AU\n\n", data.Dis)

	fmt.Printf("Altitude: %7.3f°\nAzimuth:  %7.3f°\n\n", data.Alt, data.Az)

	fmt.Println("             Azimuth  Altitude")

	t0 := novas.Date(now.Year(), int(now.Month()), now.Day(), 0, 0, 0, 0, now.Location())

	t1, topo, err := novas.Sun.Rise(t0, geo, sundip, time.Second, novas.REFR_NONE)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Sun rise:  %8.2f°            %s\n", topo.Az, t1)
		t0 = t1
	}

	t1, topo, err = novas.Sun.High(t0, geo, time.Second, novas.REFR_NONE)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("High noon: %8.2f° %8.2f°  %s\n", topo.Az, topo.Alt, t1)
		t0 = t1
	}

	t1, topo, err = novas.Sun.Set(t0, geo, sundip, time.Second, novas.REFR_NONE)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Sun set:   %8.2f°            %s\n", topo.Az, t1)
		t0 = t1
	}

	t1, topo, err = novas.Sun.Low(t0, geo, time.Second, novas.REFR_NONE)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Deep night:%8.2f° %8.2f°  %s\n\n", topo.Az, topo.Alt, t1)
		t0 = t1
	}

	fmt.Println("Start of seasons in northern hempisphere")
	fmt.Println("Spring:", novas.Spring(now.Time.Year(), time.Second))
	fmt.Println("Summer:", novas.Summer(now.Time.Year(), time.Second))
	fmt.Println("Autumn:", novas.Autumn(now.Time.Year(), time.Second))
	fmt.Println("Winter:", novas.Winter(now.Time.Year(), time.Second))
}
