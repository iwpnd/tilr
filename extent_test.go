package tilr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtentToTiles(t *testing.T) {
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
