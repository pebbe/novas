package novas

import (
	"math"
	"time"
)

// Return the percentage of the Moon disc that is illuminated.
func MoonDisc(t Time) float64 {
	data := sun.App(t)
	elon1 := data.ELon / 180 * math.Pi
	elat1 := data.ELat / 180 * math.Pi
	data = moon.App(t)
	elon2 := data.ELon / 180 * math.Pi
	elat2 := data.ELat / 180 * math.Pi

	x1 := math.Sin(elon1) * math.Cos(elat1)
	y1 := math.Cos(elon1) * math.Cos(elat1)
	z1 := math.Sin(elat1)

	x2 := math.Sin(elon2) * math.Cos(elat2)
	y2 := math.Cos(elon2) * math.Cos(elat2)
	z2 := math.Sin(elat2)

	h := x1*x2 + y1*y2 + z1*z2
	return 100 * 0.5 * (1 - h)
}

// Return the current phase of the moon as an angle.
func MoonPhase(t Time) float64 {
	h := moon.App(t).ELon - sun.App(t).ELon
	if h < 0 {
		h += 360
	}
	return h
}

// Return the time of the next lunar phase, expressed as an angle.
// 0 = new moon, 90 = first quarter, 180 = full moon, 270 = last quater.
func MoonPhaseNext(t Time, phase float64, precision time.Duration) Time {
	for phase < 0 {
		phase += 360
	}
	for phase >= 360 {
		phase -= 360
	}

	p := MoonPhase(t)
	diff := phase - p
	if diff < 0 {
		diff += 360
	}
	tt := Time{Time: t.Add(29.5 * 24 * time.Hour * time.Duration(diff) / 360)}
	t1 := Time{Time: tt.AddDate(0, 0, -2)}
	t2 := Time{Time: tt.AddDate(0, 0, 2)}

	for t2.Sub(t1.Time) > precision {
		tt.Time = t1.Add(t2.Sub(t1.Time) / 2)
		p := MoonPhase(tt)
		if p - phase > 180 {
			p -= 360
		}
		if p > phase {
			t2 = tt
		} else {
			t1 = tt
		}
	}
	return tt
}
