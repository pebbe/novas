package novas

/*
#cgo LDFLAGS: -lm
#include "eph_manager.h"
#include "novas.h"
*/
import "C"

import (
	"log"
	"unsafe"
)

const (
	AU = float64(149597870.700)
)

var (
	init_done = false
	Accuracy  = 0 // 0 = full acccuray; 1 = reduced acccuray
)

func Init(filename string, verbose bool) {

	if init_done {
		return
	}

	init_done = true

	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))

	var jd_beg, jd_end C.double
	var de_num C.short

	if err := C.ephem_open(cs, &jd_beg, &jd_end, &de_num); err != 0 {
		if err == 1 {
			log.Fatalf("JPL ephemeris file \"%s\" not found\n", filename)
		} else {
			log.Fatalf("Error reading JPL ephemeris file header \"%s\"\n", filename)
		}
	}
	if verbose {
		log.Printf("JPL ephemeris DE%d open. Start JD = %10.2f  End JD = %10.2f\n", int(de_num), float64(jd_beg), float64(jd_end))
	}
}
