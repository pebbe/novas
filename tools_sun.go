package novas

import "time"

// Computes the time of the start of spring in the given year.
func Spring(year int, precision time.Duration) Time {
	t1 := Date(year, 3, 15, 0, 0, 0, 0, time.Local)
	t2 := Time{Time: t1.AddDate(0, 0, 15)}
	tt := Time{}
	for t2.Sub(t1.Time) > precision {
		tt.Time = t1.Add(t2.Sub(t1.Time) / 2)
		if sun.App(tt).Dec > 0 {
			t2.Time = tt.Time
		} else {
			t1.Time = tt.Time
		}
	}
	return tt
}

// Computes the time of the start of summer in the given year.
func Summer(year int, precision time.Duration) Time {
	t1 := Date(year, 6, 15, 0, 0, 0, 0, time.Local)
	t2 := Time{Time: t1.AddDate(0, 0, 15)}
	tt := Time{}
	dec1 := sun.App(t1).Dec
	dec2 := sun.App(t2).Dec
	for t2.Sub(t1.Time) > precision {
		tt.Time = t1.Add(t2.Sub(t1.Time) / 2)
		dec := sun.App(tt).Dec
		if dec1 < dec2 {
			dec1 = dec
			t1 = tt
		} else {
			dec2 = dec
			t2 = tt
		}
	}
	return tt
}

// Computes the time of the start of autumn in the given year.
func Autumn(year int, precision time.Duration) Time {
	t1 := Date(year, 9, 15, 0, 0, 0, 0, time.Local)
	t2 := Time{Time: t1.AddDate(0, 0, 15)}
	tt := Time{}
	for t2.Sub(t1.Time) > precision {
		tt.Time = t1.Add(t2.Sub(t1.Time) / 2)
		if sun.App(tt).Dec < 0 {
			t2.Time = tt.Time
		} else {
			t1.Time = tt.Time
		}
	}
	return tt
}

// Computes the time of the start of winter in the given year.
func Winter(year int, precision time.Duration) Time {
	t1 := Date(year, 12, 15, 0, 0, 0, 0, time.Local)
	t2 := Time{Time: t1.AddDate(0, 0, 15)}
	tt := Time{}
	dec1 := sun.App(t1).Dec
	dec2 := sun.App(t2).Dec
	for t2.Sub(t1.Time) > precision {
		tt.Time = t1.Add(t2.Sub(t1.Time) / 2)
		dec := sun.App(tt).Dec
		if dec1 > dec2 {
			dec1 = dec
			t1 = tt
		} else {
			dec2 = dec
			t2 = tt
		}
	}
	return tt
}

