package novas

/*
#include "novas.h"
*/
import "C"

import (
	"errors"
	"log"
	"time"
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

type RefractType int

const (
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

func Mercury() *Planet {
	return mercury
}

func Venus() *Planet {
	return venus
}

func Mars() *Planet {
	return mars
}

func Jupiter() *Planet {
	return jupiter
}

func Saturn() *Planet {
	return saturn
}

func Uranus() *Planet {
	return uranus
}

func Neptune() *Planet {
	return neptune
}

func Pluto() *Planet {
	return pluto
}

func Sun() *Planet {
	return sun
}

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

func (p *Planet) Rise(t Time, geo *Place, dip float64, precision time.Duration, refr RefractType) (Time, PlanetTopoData, error) {

	alt1 := p.Topo(t, geo, refr).Alt
	t1 := t
	t2 := t
	found := false
	for i := 0; i < 48; i++ {
		t2.Time = t1.Add(30 * time.Minute)
		alt2 := p.Topo(t2, geo, refr).Alt
		if alt1 < dip && alt2 > dip {
			found = true
			break
		}
		t1 = t2
		alt1 = alt2
	}
	if !found {
		return Time{}, PlanetTopoData{}, errors.New("No rise above dip in the next 24 hours")
	}

	var topo PlanetTopoData
	tt := t1
	for t2.Sub(t1.Time) > precision {
		tt.Time = t1.Add(t2.Sub(t1.Time) / 2)
		topo = p.Topo(tt, geo, refr)
		if topo.Alt > dip {
			t2 = tt
		} else {
			t1 = tt
		}
	}
	return tt, topo, nil
}

func (p *Planet) Set(t Time, geo *Place, dip float64, precision time.Duration, refr RefractType) (Time, PlanetTopoData, error) {

	alt1 := p.Topo(t, geo, refr).Alt
	t1 := t
	t2 := t
	found := false
	for i := 0; i < 48; i++ {
		t2.Time = t1.Add(30 * time.Minute)
		alt2 := p.Topo(t2, geo, refr).Alt
		if alt1 > dip && alt2 < dip {
			found = true
			break
		}
		t1 = t2
		alt1 = alt2
	}
	if !found {
		return Time{}, PlanetTopoData{}, errors.New("No set below dip in the next 24 hours")
	}

	var topo PlanetTopoData
	tt := t1
	for t2.Sub(t1.Time) > precision {
		tt.Time = t1.Add(t2.Sub(t1.Time) / 2)
		topo = p.Topo(tt, geo, refr)
		if topo.Alt < dip {
			t2 = tt
		} else {
			t1 = tt
		}
	}
	return tt, topo, nil
}

func (p *Planet) High(t Time, geo *Place, precision time.Duration, refr RefractType) (Time, PlanetTopoData, error) {

	alt1 := p.Topo(t, geo, refr).Alt
	alt2 := alt1
	alt3 := alt1
	t1 := t
	t2 := t
	t3 := t
	found := false
	for i := 0; i < 48; i++ {
		t3.Time = t2.Add(30 * time.Minute)
		alt3 := p.Topo(t3, geo, refr).Alt
		if alt3 < alt2 && alt2 > alt1 {
			found = true
			break
		}
		t1, t2 = t2, t3
		alt1, alt2 = alt2, alt3
	}

	if !found {
		return Time{}, PlanetTopoData{}, errors.New("No high point in the next 24 hours")
	}

	var topo PlanetTopoData
	tt := t1
	for t3.Sub(t1.Time) > precision {
		tt.Time = t1.Add(t3.Sub(t1.Time) / 2)
		topo = p.Topo(tt, geo, refr)
		if alt1 < alt3 {
			t1 = tt
			alt1 = topo.Alt
		} else {
			t3 = tt
			alt3 = topo.Alt
		}
	}
	return tt, topo, nil
}

func (p *Planet) Low(t Time, geo *Place, precision time.Duration, refr RefractType) (Time, PlanetTopoData, error) {

	alt1 := p.Topo(t, geo, refr).Alt
	alt2 := alt1
	alt3 := alt1
	t1 := t
	t2 := t
	t3 := t
	found := false
	for i := 0; i < 48; i++ {
		t3.Time = t2.Add(30 * time.Minute)
		alt3 := p.Topo(t3, geo, refr).Alt
		if alt3 > alt2 && alt2 < alt1 {
			found = true
			break
		}
		t1, t2 = t2, t3
		alt1, alt2 = alt2, alt3
	}

	if !found {
		return Time{}, PlanetTopoData{}, errors.New("No low point in the next 24 hours")
	}

	var topo PlanetTopoData
	tt := t1
	for t3.Sub(t1.Time) > precision {
		tt.Time = t1.Add(t3.Sub(t1.Time) / 2)
		topo = p.Topo(tt, geo, refr)
		if alt1 > alt3 {
			t1 = tt
			alt1 = topo.Alt
		} else {
			t3 = tt
			alt3 = topo.Alt
		}
	}
	return tt, topo, nil
}
