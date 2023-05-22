package tilr

import (
	"math"
)

// Point ...
type Point struct {
	Lng, Lat float64
}

// Valid validates a Point
func (p Point) Valid() bool {
	return p.Lng >= -180 && p.Lng <= 180 && p.Lat >= -90 && p.Lat <= 90
}

// ToTile calculates Tile coordinates Z/X/Y from a given Point
func (p Point) ToTile(z int) (Tile, error) {
	if !p.Valid() {
		return Tile{}, &ErrInvalidPoint{p: p}
	}

	latRad := degreeToRad(p.Lat)
	n := math.Pow(2, float64(z))

	xtile := int((p.Lng + 180) / 360 * n)
	ytile := int((1 - math.Asinh(math.Tan(latRad))/math.Pi) / 2 * n)

	return Tile{z, xtile, ytile}, nil
}

// Intersects validates if Point is in a given Tile bounding box
func (p Point) Intersects(t Tile) (bool, error) {
	if !p.Valid() {
		return false, &ErrInvalidPoint{p: p}
	}
	extent := t.Extent()

	return p.Lng > extent.MinLng && p.Lat > extent.MinLat && p.Lng < extent.MaxLng && p.Lat < extent.MaxLat, nil
}
