package novas

/*
#include "novas.h"
*/
import "C"

import (
	"log"
	"math"
)

type bodyClass int

const (
	clSTAR = bodyClass(iota)
	clPLANET
)

type Body struct {
	class     bodyClass
	object    C.object
	cat_entry C.cat_entry
	name      string
}

// Information returned by function (*Body) App(Time)
type BodyData struct {
	RA   float64 // Right ascension
	Dec  float64 // Declination
	Dis  float64 // Distance in AU
	ELon float64 // Ecliptic longitude
	ELat float64 // Ecliptic latitude
}

// Information returned by function (*Body) Topo(Time, *Place, RefractType)
type BodyTopoData struct {
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

// Get the name of a body.
func (b *Body) Name() string {
	return b.name
}

// Compute the apparent place of an object.
func (p *Body) App(t Time) BodyData {

	t.update()

	data := BodyData{}

	var ra, dec, dis C.double
	switch p.class {
	case clPLANET:
		if err := C.app_planet(C.double(t.jd_tt), &p.object, C.short(Accuracy), &ra, &dec, &dis); err != 0 {
			log.Fatalf("Error %d from app_planet (%s)\n", int(err), p.name)
		}
		data.Dis = float64(dis)
	case clSTAR:
		if err := C.app_star(C.double(t.jd_tt), &p.cat_entry, C.short(Accuracy), &ra, &dec); err != 0 {
			log.Fatalf("Error %d from app_star (%s)\n", int(err), p.name)
		}
		data.Dis = math.NaN()
	}
	data.RA = float64(ra)
	data.Dec = float64(dec)

	var elon, elat C.double
	C.equ2ecl(C.double(t.jd_tt), 0, C.short(Accuracy), ra, dec, &elon, &elat)
	data.ELon = float64(elon)
	data.ELat = float64(elat)

	return data
}

// Compute the topocentric place of a solar system body.
func (p *Body) Topo(t Time, geo *Place, refr RefractType) BodyTopoData {

	t.update()

	data := BodyTopoData{}

	var ra, dec, dis C.double

	switch p.class {
	case clPLANET:
		if err := C.topo_planet(C.double(t.jd_tt), &p.object, C.double(t.delta_t), &geo.place, C.short(Accuracy), &ra, &dec, &dis); err != 0 {
			log.Fatalf("Error %d from app_planet (%s)\n", int(err), p.name)
		}
		data.Dis = float64(dis)
	case clSTAR:
		if err := C.topo_star(C.double(t.jd_tt), C.double(t.delta_t), &p.cat_entry, &geo.place, C.short(Accuracy), &ra, &dec); err != 0 {
			log.Fatalf("Error %d from app_planet (%s)\n", int(err), p.name)
		}
		data.Dis = math.NaN()
	}

	var elon, elat C.double
	C.equ2ecl(C.double(t.jd_tt), 0, C.short(Accuracy), ra, dec, &elon, &elat)

	var zd, az, rar, decr C.double
	C.equ2hor(C.double(t.jd_ut1), C.double(t.delta_t), C.short(Accuracy), 0, 0, &geo.place, ra, dec, C.short(refr), &zd, &az, &rar, &decr)
	data.Alt = 90 - float64(zd)
	data.Az = float64(az)

	return data
}
