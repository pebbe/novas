package novas

/*
#include <novas.h>
*/
import "C"

type Place C.on_surface

func NewPlace(latitude, longitude, height, temperature, pressure float64) *Place {
	p := new(Place)
	C.make_on_surface(
		C.double(latitude),
		C.double(longitude),
		C.double(height),
		C.double(temperature),
		C.double(pressure),
		(*C.on_surface)(p))
	return p
}
