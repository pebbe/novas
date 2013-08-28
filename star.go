package novas

/*
#include "novas.h"
*/
import "C"

import (
	"unsafe"
)

func NewStar(starname, catalog string, starnumber int64, ra, dec, promora, promodec, parallax, radialvelocity float64) *Body {

	star := &Body{class: STAR, name: starname}

	cs1 := C.CString(starname)
	cs2 := C.CString(catalog)
	defer C.free(unsafe.Pointer(cs1))
	defer C.free(unsafe.Pointer(cs2))

	C.make_cat_entry(cs1, cs2, C.long(starnumber), C.double(ra), C.double(dec),
		C.double(promora), C.double(promodec), C.double(parallax), C.double(radialvelocity), &star.cat_entry)

	return star
}
