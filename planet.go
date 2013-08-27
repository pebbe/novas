package novas

/*
#include "novas.h"
*/
import "C"

import (
	"log"
	"unsafe"
)

type Planet struct {
	planet     C.object
	dummy_star C.cat_entry
	name       string
}

// Information returned by function (*Planet) App(Time)
type PlanetData struct {
	RA   float64 // Right ascension
	Dec  float64 // Declination
	Dis  float64 // Distance in AU
	ELon float64 // Ecliptic longitude
	ELat float64 // Ecliptic latitude
}

// Information returned by function (*Planet) Topo(Time, *Place, RefractType)
type PlanetTopoData struct {
	Dis float64 // Distance in AU
	Az  float64 // Azimuth
	Alt float64 // Altitude
}

type RefractType int

const (
	// Types of refraction correction
	REFR_NONE     = RefractType(0)
	REFR_STANDARD = RefractType(1)
	REFR_PLACE    = RefractType(2)
)

var (
	mercury, venus, mars, jupiter, saturn, uranus, neptune, pluto, sun, moon *Planet
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
func Mercury() *Planet {
	return mercury
}

// Create planet Venus.
func Venus() *Planet {
	return venus
}

// Create planet Mars.
func Mars() *Planet {
	return mars
}

// Create planet Jupiter.
func Jupiter() *Planet {
	return jupiter
}

// Create planet Saturn.
func Saturn() *Planet {
	return saturn
}

// Create planet Uranus.
func Uranus() *Planet {
	return uranus
}

// Create planet Neptune.
func Neptune() *Planet {
	return neptune
}

// Create planet Pluto.
func Pluto() *Planet {
	return pluto
}

// Create solar body Sun.
func Sun() *Planet {
	return sun
}

// Create solar body Moon.
func Moon() *Planet {
	return moon
}

func new_planet(id int, name string) *Planet {

	planet := &Planet{}

	cs1 := C.CString("DUMMY")
	cs2 := C.CString("xxx")
	cs3 := C.CString(name)
	defer C.free(unsafe.Pointer(cs1))
	defer C.free(unsafe.Pointer(cs2))
	defer C.free(unsafe.Pointer(cs3))

	C.make_cat_entry(cs1, cs2, 0, 0, 0, 0, 0, 0, 0, &planet.dummy_star)
	if err := C.make_object(0, C.short(id), cs3, &planet.dummy_star, &planet.planet); err != 0 {
		log.Fatalf("Error %d from make_object (%s)\n", int(err), name)
	}
	planet.name = name
	return planet
}

// Get the name of a solar system body.
func (p *Planet) Name() string {
	return p.name
}

// Compute the apparent place of a solar system body.
func (p *Planet) App(t Time) PlanetData {

	t.update()

	data := PlanetData{}

	var ra, dec, dis C.double
	if err := C.app_planet(C.double(t.jd_tt), &p.planet, C.short(Accuracy), &ra, &dec, &dis); err != 0 {
		log.Fatalf("Error %d from app_planet (%s)\n", int(err), p.name)
	}
	data.RA = float64(ra)
	data.Dec = float64(dec)
	data.Dis = float64(dis)

	var elon, elat C.double
	C.equ2ecl(C.double(t.jd_tt), 0, C.short(Accuracy), ra, dec, &elon, &elat)
	data.ELon = float64(elon)
	data.ELat = float64(elat)

	return data
}

// Compute the topocentric place of a solar system body.
func (p *Planet) Topo(t Time, geo *Place, refr RefractType) PlanetTopoData {

	t.update()

	data := PlanetTopoData{}

	var ra, dec, dis C.double

	if err := C.topo_planet(C.double(t.jd_tt), &p.planet, C.double(t.delta_t), &geo.place, C.short(Accuracy), &ra, &dec, &dis); err != 0 {
		log.Fatalf("Error %d from app_planet (%s)\n", int(err), p.name)
	}
	data.Dis = float64(dis)

	var elon, elat C.double
	C.equ2ecl(C.double(t.jd_tt), 0, C.short(Accuracy), ra, dec, &elon, &elat)

	var zd, az, rar, decr C.double
	C.equ2hor(C.double(t.jd_ut1), C.double(t.delta_t), C.short(Accuracy), 0, 0, &geo.place, ra, dec, C.short(refr), &zd, &az, &rar, &decr)
	data.Alt = 90 - float64(zd)
	data.Az = float64(az)

	return data
}
