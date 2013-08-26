package novas

import "time"

func Spring(year int, precision time.Duration) Time {
	t1 := Date(year, 3, 15, 0, 0, 0, 0, time.Local)
	t2 := Time{Time: t1.AddDate(0, 0, 15)}
	tt := Time{}
	for t2.Sub(t1.Time) > precision {
		tt.Time = t1.Add(t2.Sub(t1.Time) / 2)
		if Sun.App(tt).Dec > 0 {
			t2.Time = tt.Time
		} else {
			t1.Time = tt.Time
		}
	}
	return tt
}

func Summer(year int, precision time.Duration) Time {
	t1 := Date(year, 6, 15, 0, 0, 0, 0, time.Local)
	t2 := Time{Time: t1.AddDate(0, 0, 15)}
	tt := Time{}
	dec1 := Sun.App(t1).Dec
	dec2 := Sun.App(t2).Dec
	for t2.Sub(t1.Time) > precision {
		tt.Time = t1.Add(t2.Sub(t1.Time) / 2)
		dec := Sun.App(tt).Dec
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

func Autumn(year int, precision time.Duration) Time {
	t1 := Date(year, 9, 15, 0, 0, 0, 0, time.Local)
	t2 := Time{Time: t1.AddDate(0, 0, 15)}
	tt := Time{}
	for t2.Sub(t1.Time) > precision {
		tt.Time = t1.Add(t2.Sub(t1.Time) / 2)
		if Sun.App(tt).Dec < 0 {
			t2.Time = tt.Time
		} else {
			t1.Time = tt.Time
		}
	}
	return tt
}

func Winter(year int, precision time.Duration) Time {
	t1 := Date(year, 12, 15, 0, 0, 0, 0, time.Local)
	t2 := Time{Time: t1.AddDate(0, 0, 15)}
	tt := Time{}
	dec1 := Sun.App(t1).Dec
	dec2 := Sun.App(t2).Dec
	for t2.Sub(t1.Time) > precision {
		tt.Time = t1.Add(t2.Sub(t1.Time) / 2)
		dec := Sun.App(tt).Dec
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

