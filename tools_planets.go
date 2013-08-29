package novas

import (
	"math"
)

// Return the percentage of the planet disc that is illuminated.
// Useful for moon and inner planets.
func (b *Body) Disc(t Time) float64 {
	if b.class != clPLANET || b.name == "Sun" {
		panic("Object " + b.name + " is not a planet or the moon")
	}

	// vector body -> earth
	// - vector earth -> body
	data := b.App(t)
	bed := data.Dis
	bex := -1 * bed * math.Sin(data.ELon/180*math.Pi) * math.Cos(data.ELat/180*math.Pi)
	bey := -1 * bed * math.Cos(data.ELon/180*math.Pi) * math.Cos(data.ELat/180*math.Pi)
	bez := -1 * bed * math.Sin(data.ELat/180*math.Pi)

	// vector earth -> sun
	data = sun.App(t)
	esd := data.Dis
	esx := esd * math.Sin(data.ELon/180*math.Pi) * math.Cos(data.ELat/180*math.Pi)
	esy := esd * math.Cos(data.ELon/180*math.Pi) * math.Cos(data.ELat/180*math.Pi)
	esz := esd * math.Sin(data.ELat/180*math.Pi)

	// vector body -> sun
	// vector body -> earth + vector earth -> sun
	bsx := bex + esx
	bsy := bey + esy
	bsz := bez + esz
	bsd := math.Sqrt(bsx*bsx + bsy*bsy + bsz*bsz)

	// normalise to length 1
	bsx /= bsd
	bsy /= bsd
	bsz /= bsd

	bex /= bed
	bey /= bed
	bez /= bed

	cos_h := bsx*bex + bsy*bey + bsz*bez
	return 100 * (0.5 + 0.5*cos_h)
}
