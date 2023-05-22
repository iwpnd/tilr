package tilr

import "math"

// Extent ...
type Extent struct {
	MinLng, MinLat, MaxLng, MaxLat float64
}

// ToSlice ...
func (b Extent) ToSlice() []float64 {
	return []float64{b.MinLng, b.MinLat, b.MaxLng, b.MaxLat}
}

// FromBounds returns all tiles of a
// certain zoom level that intersect an input
// bounding box
func (b Extent) ToTiles(z int) ([]Tile, error) {
	Z := int(math.Pow(2, float64(z)))

	var bbs []Extent

	if b.MinLng > b.MaxLng {
		bbw := Extent{MaxLat: b.MaxLat, MaxLng: b.MaxLng, MinLng: -180.0, MinLat: b.MinLat}
		bbe := Extent{MaxLat: b.MaxLat, MaxLng: 180.0, MinLng: b.MinLng, MinLat: b.MinLat}
		bbs = []Extent{bbw, bbe}
	} else {
		bbs = []Extent{b}
	}

	var tiles []Tile

	for _, bb := range bbs {
		minlng := max(-180.0, bb.MinLng)
		minlat := max(-85.0, bb.MinLat)
		maxlng := min(180.0, bb.MaxLng)
		maxlat := min(85.0, bb.MaxLat)

		ult, err := Point{Lng: minlng, Lat: maxlat}.ToTile(z)

		if err != nil {
			return nil, err
		}

		lrt, err := Point{Lng: maxlng, Lat: minlat}.ToTile(z)

		if err != nil {
			return nil, err
		}

		for i := ult.X; i <= lrt.X; i++ {
			for j := ult.Y; j <= lrt.Y; j++ {
				// ignore coordinates >= 2 ** zoom
				if i >= Z {
					continue
				}

				if j >= Z {
					continue
				}

				tile := Tile{X: i, Y: j, Z: z}
				tiles = append(tiles, tile)
			}
		}
	}

	return tiles, nil
}
