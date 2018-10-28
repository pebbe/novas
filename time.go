package novas

/*
#include "NOVAS_novas.h"
*/
import "C"

import (
	"time"
)

/*
Type Time holds time in a format needed by the C Novas functions.

Type Time is a struct that includes a time.Time object.
All methods of time.Time are available.
To set the time of an object Time t through assignement, use: t.Time =
*/
type Time struct {
	time.Time
	current                                time.Time
	jd_utc, jd_tt, jd_ut1, delta_t, jd_tdb float64
}

var (
	Leap_secs = 33
	UT1_UTC   = float64(-0.387845)
)

// Now returns the current local time.
func Now() Time {
	return Time{Time: time.Now()}
}

/*
Date returns the Time corresponding to

    yyyy-mm-dd hh:mm:ss + nsec nanoseconds

in the appropriate zone for that time in the given location.

See godoc on time.Date for details.
*/
func Date(year, month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	return Time{Time: time.Date(year, time.Month(month), day, hour, min, sec, nsec, loc)}
}

// Julian returns the Time corresponding to the given Julian date
func Julian(jd float64) Time {
	var year, month, day C.short
	var hour C.double
	C.cal_date(C.double(jd), &year, &month, &day, &hour)

	f := float64(hour)
	nano := int64(f * 3600 * 1e9)
	h := nano / (3600 * 1e9)
	nano %= 3600 * 1e9
	m := nano / (60 * 1e9)
	nano %= 60 * 1e9
	s := nano / 1e9
	nano %= 1e9

	return Time{Time: time.Date(int(year), time.Month(int(month)), int(day), int(h), int(m), int(s), int(nano), time.UTC)}
}

func (t *Time) ToJulian() float64 {
	t.update()
	return t.jd_utc
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
		C.double(float64(t1.Hour())+float64(t1.Minute())/60+float64(t1.Second())/3600+float64(t1.Nanosecond())/(3600*1e9))))
	t.jd_tt = t.jd_utc + (float64(Leap_secs)+32.184)/86400
	t.jd_ut1 = t.jd_utc + UT1_UTC/86400
	t.delta_t = 32.184 + float64(Leap_secs) - UT1_UTC
	t.jd_tdb = t.jd_tt /* Approximation good to 0.0017 seconds. */
}
