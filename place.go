package novas

/*
#include <novas.h>
*/
import "C"

import (
	"fmt"
)

type Place struct {
	place C.on_surface
	lat, long, h, t, p float64
}

func NewPlace(latitude, longitude, height, temperature, pressure float64) *Place {
	p := Place{lat: latitude, long: longitude, h:height, t: temperature, p: pressure}
	C.make_on_surface(
		C.double(latitude),
		C.double(longitude),
		C.double(height),
		C.double(temperature),
		C.double(pressure),
		&p.place)
	return &p
}

func (p Place) String() string {
	var s string
	if p.lat < 0 {
		s = fmt.Sprintf("%.3f°S", -p.lat)
	} else {
		s = fmt.Sprintf("%.3f°N", p.lat)
	}
	if p.long < 0 {
		s += fmt.Sprintf("  %.3f°W", -p.long)
	} else {
		s += fmt.Sprintf("  %.3f°E", p.long)
	}
	return s
}
