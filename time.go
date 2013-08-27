package novas

/*
#include "novas.h"
*/
import "C"

import (
	"time"
)

type Time struct {
	time.Time
	current                                time.Time
	jd_utc, jd_tt, jd_ut1, delta_t, jd_tdb float64
}

var (
	Leap_secs = 33
	UT1_UTC   = float64(-0.387845)
)

func Now() Time {
	return Time{Time: time.Now()}
}

func Date(year, month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	return Time{Time: time.Date(year, time.Month(month), day, hour, min, sec, nsec, loc)}
}

func (t *Time) update() {
	if t.Time.Equal(t.current) {
		return
	}
	t.current = t.Time
	t1 := t.current.UTC()
	t.jd_utc = float64(C.julian_date(
		C.short(t1.Year()),
		C.short(t1.Month()),
		C.short(t1.Day()),
		C.double(float64(t1.Hour())+float64(t1.Minute())/60+float64(t1.Second())/3600+float64(t1.Nanosecond())/3600/1e9)))
	t.jd_tt = t.jd_utc + (float64(Leap_secs)+32.184)/86400
	t.jd_ut1 = t.jd_utc + UT1_UTC/86400
	t.delta_t = 32.184 + float64(Leap_secs) - UT1_UTC
	t.jd_tdb = t.jd_tt /* Approximation good to 0.0017 seconds. */
}
