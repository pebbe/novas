package novas

/*
#include "NOVAS_novas.h"
*/
import "C"

import (
	"unsafe"
)

/*
Create a star
    starname:       name of celestial object
    catalog:        catalog designator (e.g., HIP)
    starnumber:     integer identifier assigned to object
    ra:             ICRS right ascension (hours)
    dec:            ICRS declination (degrees)
    promora:        ICRS proper motion in right ascension (milliarcseconds/year)
    promodec:       ICRS proper motion in declination (milliarcseconds/year)
    parallax:       parallax (milliarcseconds)
    radialvelocity: radial velocity (km/s)
*/
func NewStar(starname, catalog string, starnumber int64, ra, dec, promora, promodec, parallax, radialvelocity float64) *Body {

	if len(starname) >= int(C.SIZE_OF_OBJ_NAME) {
		panic("Length of starname \"" + starname + "\" too long")
	}

	if len(catalog) >= int(C.SIZE_OF_CAT_NAME) {
		panic("Length of catalog \"" + catalog + "\" too long")
	}

	star := &Body{class: clSTAR, name: starname}

	cs1 := C.CString(starname)
	cs2 := C.CString(catalog)
	defer C.free(unsafe.Pointer(cs1))
	defer C.free(unsafe.Pointer(cs2))

	C.make_cat_entry(cs1, cs2, C.long(starnumber), C.double(ra), C.double(dec),
		C.double(promora), C.double(promodec), C.double(parallax), C.double(radialvelocity), &star.cat_entry)

	return star
}
