package novas

import (
	"time"
)

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
