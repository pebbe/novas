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
	AU = float64(149597870.700) // Astronomical Units to kilometers
)

var (
	init_done = false
	Accuracy  = 0 // 0 = full acccuray; 1 = reduced acccuray
)

/*
Initialise the package, pointing to the location of the ephemeris file, usually called JPLEPH.

If not called explicitly, this function will be called automaticly at the time the ephemeris file is first needed.
It will then search for JPLEPH in the current directory.

You can download the file from: http://pkleiweg.home.xs4all.nl/jpleph/
*/
func Init(filename string, verbose bool) (de_num int, jd_beg float64, jd_end float64) {

	if init_done {
		return
	}

	init_done = true

	if filename == "" {
		filename = "JPLEPH"
	}


	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))

	var jd_beg_c, jd_end_c C.double
	var de_num_c C.short

	if err := C.ephem_open(cs, &jd_beg_c, &jd_end_c, &de_num_c); err != 0 {
		if err == 1 {
			log.Fatalf("JPL ephemeris file \"%s\" not found. Download it from http://pkleiweg.home.xs4all.nl/jpleph/\n", filename)
		} else {
			log.Fatalf("Error reading JPL ephemeris file header \"%s\"\n", filename)
		}
	}
	de_num, jd_beg, jd_end = int(de_num_c), float64(jd_beg_c), float64(jd_end_c)
	if verbose {
		log.Printf("JPL ephemeris DE%d open. Start JD = %10.2f  End JD = %10.2f\n", de_num, jd_beg, jd_end)
	}

	return
}
