package tyler

// Tile ...
type Tile struct {
	Z, X, Y int
}

// Extent returns the bounding box of a given Tile
func (t Tile) Extent() Extent {
	return Extent{
		MaxLat: tileToLat(t.Y, t.Z),
		MinLng: tileToLng(t.X, t.Z),
		MinLat: tileToLat(t.Y+1, t.Z),
		MaxLng: tileToLng(t.X+1, t.Z),
	}
}

// Contains valides whether a given Point is within a given Tile
func (t Tile) Contains(p Point) (bool, error) {
	if !p.Valid() {
		return false, &ErrInvalidPoint{p: p}
	}
	extent := t.Extent()

	return p.Lat > extent.MinLat && p.Lat < extent.MaxLat && p.Lng > extent.MinLng && p.Lng < extent.MaxLng, nil
}

// Center returns the centerpoint of a given Tile
func (t Tile) Center() Point {
	extent := t.Extent()

	cLng := extent.MinLng + (extent.MaxLng-extent.MinLng)/2
	cLat := extent.MinLat + (extent.MaxLat-extent.MinLat)/2

	return Point{Lng: cLng, Lat: cLat}
}

// Children returns the four children tiles of the input tile
func (t Tile) Children() [4]Tile {
	x := t.X
	y := t.Y
	z := t.Z

	return [4]Tile{
		{X: x * 2, Y: y * 2, Z: z + 1},
		{X: x*2 + 1, Y: y * 2, Z: z + 1},
		{X: x * 2, Y: y*2 + 1, Z: z + 1},
		{X: x*2 + 1, Y: y*2 + 1, Z: z + 1},
	}
}

// Parent returns the parent tile of the input tile
func (t Tile) Parent() Tile {
	x := t.X
	y := t.Y
	z := t.Z

	if z == 0 {
		return t
	}

	if x%2 == 0 && y%2 == 0 {
		return Tile{Z: z - 1, X: x / 2, Y: y / 2}
	} else if x%2 == 0 {
		return Tile{Z: z - 1, X: x / 2, Y: (y - 1) / 2}
	} else if x%2 != 0 && y%2 == 0 {
		return Tile{Z: z - 1, X: (x - 1) / 2, Y: y / 2}
	} else {
		return Tile{Z: z - 1, X: (x - 1) / 2, Y: (y - 1) / 2}
	}
}
