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

	longitude, latitude := 6.5670, 53.21853 // Groningen, The Netherlands

	// END USER SETTINGS

	novas.Init(jpleph, false)

	now := novas.Now()
	fmt.Println(now)

	if latitude < 0 {
		fmt.Printf("\nLocation: %7.3f°S\n", -latitude)
	} else {
		fmt.Printf("\nLocation: %7.3f°N\n", latitude)
	}
	if longitude < 0 {
		fmt.Printf("          %7.3f°W\n\n", -longitude)
	} else {
		fmt.Printf("          %7.3f°E\n\n", longitude)
	}

	geo := novas.NewPlace(latitude, longitude, 0, 20, 1010)

	sun := novas.NewPlanet(novas.SUN)

	fmt.Printf("Distance from earth's center:   %.7f AU\n", sun.App(now).Dis)

	data := sun.Topo(now, geo, novas.REFR_NONE)

	fmt.Printf("Distance from surface location: %.7f AU\n\n", data.Dis)

	fmt.Printf("Altitude: %7.3f°\nAzimuth:  %7.3f°\n\n", data.Alt, data.Az)



	fmt.Println("             Azimuth  Altitude")


	t0 := novas.Date(now.Year(), int(now.Month()), now.Day(), 0, 0, 0, 0, now.Location())

	t1 := t0
	t2 := novas.Time{Time: t1.Add(12 * time.Hour)}
	for t2.Sub(t1.Time) > time.Second {
		t0.Time = t1.Add(t2.Sub(t1.Time) / 2)
		data = sun.Topo(t0, geo, novas.REFR_NONE)
		if data.Alt > sundip {
			t2 = t0
		} else {
			t1 = t0
		}
	}
	fmt.Printf("Sun rise:  %8.2f°            %s\n", data.Az, t0)

	alt1 := data.Alt
	t2.Time = t1.Add(12 * time.Hour)
	data = sun.Topo(t2, geo, novas.REFR_NONE)
	alt2 :=  data.Alt
	for t2.Sub(t1.Time) > time.Second {
		t0.Time = t1.Add(t2.Sub(t1.Time) / 2)
		data = sun.Topo(t0, geo, novas.REFR_NONE)
		if alt1 < alt2 {
			t1 = t0
			alt1 = data.Alt
		} else {
			t2 = t0
			alt2 = data.Alt
		}
	}
	fmt.Printf("High noon: %8.2f° %8.2f°  %s\n", data.Az, data.Alt, t0)

	t2.Time = t1.Add(12 * time.Hour)
	for t2.Sub(t1.Time) > time.Second {
		t0.Time = t1.Add(t2.Sub(t1.Time) / 2)
		data = sun.Topo(t0, geo, novas.REFR_NONE)
		if data.Alt < sundip {
			t2 = t0
		} else {
			t1 = t0
		}
	}
	fmt.Printf("Sun set:   %8.2f°            %s\n\n", data.Az, t0)

	fmt.Println("Start of seasons in northern hempisphere")
	fmt.Println("Spring:", novas.Spring(now.Time.Year(), time.Second))
	fmt.Println("Summer:", novas.Summer(now.Time.Year(), time.Second))
	fmt.Println("Autumn:", novas.Autumn(now.Time.Year(), time.Second))
	fmt.Println("Winter:", novas.Winter(now.Time.Year(), time.Second))

}
