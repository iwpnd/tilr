package tyler

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func inDeep(x Tile, y []Tile) bool {
	for i := range y {
		if reflect.DeepEqual(x, y[i]) {
			return true
		}
	}
	return false
}

func TestFromBounds(t *testing.T) {
	tests := []struct {
		Extent        Extent
		z             int
		ExpectedTiles []Tile
	}{
		{
			Extent: Extent{
				MinLng: 10.045,
				MinLat: 51.2114,
				MaxLng: 13.825,
				MaxLat: 53.575,
			},
			z: 7,
			ExpectedTiles: []Tile{
				{X: 67, Y: 41, Z: 7},
				{X: 68, Y: 41, Z: 7},
				{X: 67, Y: 42, Z: 7},
				{X: 68, Y: 42, Z: 7},
			},
		},
		{
			Extent: Extent{
				MinLng: 9.7061,
				MinLat: 53.3942,
				MaxLng: 10.3019,
				MaxLat: 53.763,
			},
			z:             7,
			ExpectedTiles: []Tile{{X: 67, Y: 41, Z: 7}},
		},
		{
			Extent: Extent{
				MinLng: 178.65,
				MinLat: 70.81,
				MaxLng: -177.58,
				MaxLat: 71.6,
			},
			z: 7,
			ExpectedTiles: []Tile{
				{X: 127, Y: 26, Z: 7},
				{X: 0, Y: 26, Z: 7},
				{X: 127, Y: 27, Z: 7},
				{X: 0, Y: 27, Z: 7},
			},
		},
	}

	for _, test := range tests {
		got, err := test.Extent.ToTiles(test.z)
		assert.Nil(t, err, "cannot get tiles from extent")
		assert.Equal(t, len(test.ExpectedTiles), len(got))

		for i := range got {
			assert.Equal(t, true, inDeep(got[i], test.ExpectedTiles))
		}
	}
}

func TestTileChildren(t *testing.T) {
	tile := Tile{Z: 0, X: 0, Y: 0}
	expectedX := []int{tile.X * 2, tile.X*2 + 1}
	expectedY := []int{tile.Y * 2, tile.Y*2 + 1}

	children := tile.Children()

	assert.Equal(t, 4, len(children))

	for _, e := range children {
		assert.Equal(t, tile.Z+1, e.Z)
		assert.Contains(t, expectedX, e.X)
		assert.Contains(t, expectedY, e.Y)
	}
}

func TestTileParent(t *testing.T) {
	var tests = []struct {
		Child        Tile
		ExpectedTile Tile
	}{
		{
			Child:        Tile{Z: 1, X: 1, Y: 1},
			ExpectedTile: Tile{Z: 0, X: 0, Y: 0},
		},
		{
			Child:        Tile{Z: 1, X: 1, Y: 0},
			ExpectedTile: Tile{Z: 0, X: 0, Y: 0},
		},
		{
			Child:        Tile{Z: 1, X: 0, Y: 0},
			ExpectedTile: Tile{Z: 0, X: 0, Y: 0},
		},
		{
			Child:        Tile{Z: 1, X: 0, Y: 1},
			ExpectedTile: Tile{Z: 0, X: 0, Y: 0},
		},
	}

	for _, test := range tests {
		got := test.Child.Parent()
		assert.Equal(t, test.ExpectedTile, got)
	}
}

func TestTileToExtent(t *testing.T) {
	delta := 0.000000001

	var tests = []struct {
		z              int
		x              int
		y              int
		ExpectedMinLng float64
		ExpectedMinLat float64
		ExpectedMaxLng float64
		ExpectedMaxLat float64
	}{
		{
			z:              11,
			x:              525,
			y:              761,
			ExpectedMinLng: -87.71484375,
			ExpectedMinLat: 41.77131167976407,
			ExpectedMaxLng: -87.5390625,
			ExpectedMaxLat: 41.9022770409637,
		},
		{
			z:              15,
			x:              17599,
			y:              10756,
			ExpectedMinLng: 13.348388671875,
			ExpectedMinLat: 52.44931414086969,
			ExpectedMaxLng: 13.359375,
			ExpectedMaxLat: 52.456009392640745,
		},
		{
			z:              11,
			x:              1095,
			y:              641,
			ExpectedMinLng: 12.48046875,
			ExpectedMinLat: 55.57834467218205,
			ExpectedMaxLng: 12.65625,
			ExpectedMaxLat: 55.67758441108952,
		},
	}

	for _, test := range tests {
		tile := Tile{Z: test.z, X: test.x, Y: test.y}
		extent := tile.Extent()

		assert.InDelta(t, test.ExpectedMaxLat, extent.MaxLat, delta)
		assert.InDelta(t, test.ExpectedMaxLng, extent.MaxLng, delta)
		assert.InDelta(t, test.ExpectedMinLat, extent.MinLat, delta)
		assert.InDelta(t, test.ExpectedMinLng, extent.MinLng, delta)
	}
}

func TestTileContainsPoint(t *testing.T) {
	var tests = []struct {
		Tile     Tile
		Point    Point
		expected bool
	}{
		{
			Tile:     Tile{Z: 11, X: 525, Y: 761},
			Point:    Point{Lng: -87.65, Lat: 41.84},
			expected: true,
		},
		{
			Tile:     Tile{Z: 11, X: 1099, Y: 641},
			Point:    Point{Lng: 12.568337, Lat: 55.67609},
			expected: false,
		},
	}

	for _, test := range tests {
		i, _ := test.Tile.Contains(test.Point)

		assert.Equal(t, test.expected, i)
	}

	tile2 := Tile{Z: 11, X: 1099, Y: 641}
	_, err := tile2.Contains(Point{Lng: 999, Lat: 999})

	if assert.Error(t, err) {
		assert.Equal(t, &ErrInvalidPoint{p: Point{999, 999}}, err)
	}
}

func TestTileCenter(t *testing.T) {
	var tests = []struct {
		Tile           Tile
		ExpectedCenter Point
	}{
		{
			Tile:           Tile{Z: 11, X: 525, Y: 761},
			ExpectedCenter: Point{Lng: -87.626953125, Lat: 41.83679436036388},
		},
		{
			Tile:           Tile{Z: 15, X: 17599, Y: 10756},
			ExpectedCenter: Point{13.3538818359375, 52.45266176675521},
		},
	}

	for _, test := range tests {
		p := test.Tile.Center()
		assert.Equal(t, test.ExpectedCenter.Lat, p.Lat)
		assert.Equal(t, test.ExpectedCenter.Lng, p.Lng)
	}
}
