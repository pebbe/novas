package novas

/*
#cgo LDFLAGS: -lm
#include "NOVAS_eph_manager.h"
#include "NOVAS_novas.h"
*/
import "C"

import (
	"fmt"
	"log"
	"unsafe"
)

const (
	AU = float64(149597870.700) // Astronomical Units to kilometers
)

var (
	eph_data EphData

	Accuracy = 0 // 0 = full acccuray; 1 = reduced acccuray
)

func init() {
	filename := jpleph()
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	var jd_beg, jd_end C.double
	var de_num C.short
	if err := C.ephem_open(cs, &jd_beg, &jd_end, &de_num); err != 0 {
		if err == 1 {
			log.Fatalf("JPL ephemeris file \"%s\" not found. Download it from http://pkleiweg.home.xs4all.nl/jpleph/\n", filename)
		} else {
			log.Fatalf("Error reading JPL ephemeris file header \"%s\"\n", filename)
		}
	}
	eph_data = EphData{DE_num: int(de_num), JD_beg: float64(jd_beg), JD_end: float64(jd_end)}
}

type EphData struct {
	DE_num int     // number of DE file
	JD_beg float64 // begin date of DE file in jd
	JD_end float64 // end date of DE file in jd
}

// Return info on current ephemeris file
func EphInfo() EphData {
	return eph_data
}

func (e EphData) String() string {
	return fmt.Sprintf("JPL ephemeris DE%d. Start JD = %10.2f  End JD = %10.2f", e.DE_num, e.JD_beg, e.JD_end)
}
