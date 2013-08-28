package novas

/*
#include "NOVAS_novas.h"
*/
import "C"

import (
	"log"
	"unsafe"
)

var (
	mercury, venus, mars, jupiter, saturn, uranus, neptune, pluto, sun, moon *Body
)

func init() {
	mercury = new_planet(1, "Mercury")
	venus = new_planet(2, "Venus")
	mars = new_planet(4, "Mars")
	jupiter = new_planet(5, "Jupiter")
	saturn = new_planet(6, "Saturn")
	uranus = new_planet(7, "Uranus")
	neptune = new_planet(8, "Neptune")
	pluto = new_planet(9, "Pluto")
	sun = new_planet(10, "Sun")
	moon = new_planet(11, "Moon")
}

// Create planet Mercury.
func Mercury() *Body {
	return mercury
}

// Create planet Venus.
func Venus() *Body {
	return venus
}

// Create planet Mars.
func Mars() *Body {
	return mars
}

// Create planet Jupiter.
func Jupiter() *Body {
	return jupiter
}

// Create planet Saturn.
func Saturn() *Body {
	return saturn
}

// Create planet Uranus.
func Uranus() *Body {
	return uranus
}

// Create planet Neptune.
func Neptune() *Body {
	return neptune
}

// Create planet Pluto.
func Pluto() *Body {
	return pluto
}

// Create solar body Sun.
func Sun() *Body {
	return sun
}

// Create solar body Moon.
func Moon() *Body {
	return moon
}

func new_planet(id int, name string) *Body {

	planet := &Body{class: clPLANET, name: name}

	cs1 := C.CString("DUMMY")
	cs2 := C.CString("xxx")
	cs3 := C.CString(name)
	defer C.free(unsafe.Pointer(cs1))
	defer C.free(unsafe.Pointer(cs2))
	defer C.free(unsafe.Pointer(cs3))

	C.make_cat_entry(cs1, cs2, 0, 0, 0, 0, 0, 0, 0, &planet.cat_entry)
	if err := C.make_object(0, C.short(id), cs3, &planet.cat_entry, &planet.object); err != 0 {
		log.Fatalf("Error %d from make_object (%s)\n", int(err), name)
	}
	return planet
}
