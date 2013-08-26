package novas

/*
#include <novas.h>
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

type PlanetData struct {
	RA   float64 // Right ascension
	Dec  float64 // Declination
	Dis  float64 // Distance in AU
	ELon float64 // Ecliptic longitude
	ELat float64 // Ecliptic latitude
}

type PlanetTopoData struct {
	Dis float64 // Distance in AU
	Az  float64 // Azimuth
	Alt float64 // Altitude
}

type PlanetType struct {
	id   int
	name string
}

type RefractType int

var (
	MERCURY = PlanetType{1, "Mercury"}
	VENUS   = PlanetType{2, "Venus"}
	MARS    = PlanetType{4, "Mars"}
	JUPITER = PlanetType{5, "Jupiter"}
	SATURN  = PlanetType{6, "Saturn"}
	URANUS  = PlanetType{7, "Uranus"}
	NEPTUNE = PlanetType{8, "Neptune"}
	PLUTO   = PlanetType{9, "Pluto"}
	SUN     = PlanetType{10, "Sun"}
	MOON    = PlanetType{11, "Moon"}

	REFR_NONE     = RefractType(0)
	REFR_STANDARD = RefractType(1)
	REFR_PLACE    = RefractType(2)
)

func NewPlanet(p PlanetType) *Planet {

	planet := &Planet{}

	cs1 := C.CString("DUMMY")
	cs2 := C.CString("xxx")
	cs3 := C.CString(p.name)
	defer C.free(unsafe.Pointer(cs1))
	defer C.free(unsafe.Pointer(cs2))
	defer C.free(unsafe.Pointer(cs3))

	C.make_cat_entry(cs1, cs2, 0, 0, 0, 0, 0, 0, 0, &planet.dummy_star)
	if err := C.make_object(0, C.short(p.id), cs3, &planet.dummy_star, &planet.planet); err != 0 {
		log.Fatalf("Error %d from make_object (%s)\n", int(err), p.name)
	}
	planet.name = p.name
	return planet
}

func (p *Planet) Name() string {
	return p.name
}

func (p *Planet) App(t Time) PlanetData {

	Init("JPLEPH", true)

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

func (p *Planet) Topo(t Time, geo *Place, refr RefractType) PlanetTopoData {

	Init("JPLEPH", true)

	t.update()

	data := PlanetTopoData{}

	var ra, dec, dis C.double

	if err := C.topo_planet(C.double(t.jd_tt), &p.planet, C.double(t.delta_t), (*C.on_surface)(geo), C.short(Accuracy), &ra, &dec, &dis); err != 0 {
		log.Fatalf("Error %d from app_planet (%s)\n", int(err), p.name)
	}
	data.Dis = float64(dis)

	var elon, elat C.double
	C.equ2ecl(C.double(t.jd_tt), 0, C.short(Accuracy), ra, dec, &elon, &elat)

	var zd, az, rar, decr C.double
	C.equ2hor(C.double(t.jd_ut1), C.double(t.delta_t), C.short(Accuracy), 0, 0, (*C.on_surface)(geo), ra, dec, C.short(refr), &zd, &az, &rar, &decr)
	data.Alt = 90 - float64(zd)
	data.Az = float64(az)

	return data
}
