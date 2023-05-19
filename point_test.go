package tyler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrInvalidPoint(t *testing.T) {
	p := Point{Lng: -190, Lat: 52}
	_, err := p.ToTile(15)

	if assert.Error(t, err) {
		assert.Equal(t, &ErrInvalidPoint{p: p}, err)
	}
}

func TestPointValid(t *testing.T) {
	tests := []struct {
		Point    Point
		Expected bool
	}{
		{
			Point:    Point{Lat: 52, Lng: 13},
			Expected: true,
		},
		{
			Point:    Point{Lat: 91, Lng: 13},
			Expected: false,
		},
		{
			Point:    Point{Lat: 52, Lng: 181},
			Expected: false,
		},
		{
			Point:    Point{Lat: -91, Lng: -181},
			Expected: false,
		},
	}

	for _, test := range tests {
		got := test.Point.Valid()
		assert.Equal(t, test.Expected, got)
	}
}

func TestPointToTile(t *testing.T) {
	var tests = []struct {
		Point     Point
		ExpectedZ int
		ExpectedX int
		ExpectedY int
	}{
		{
			Point:     Point{Lat: 41.84, Lng: -87.65},
			ExpectedZ: 3,
			ExpectedX: 2,
			ExpectedY: 2,
		},
		{
			Point:     Point{Lat: 52.44950563632098, Lng: 13.357951727129988},
			ExpectedZ: 15,
			ExpectedX: 17599,
			ExpectedY: 10756,
		},
	}

	for _, test := range tests {
		tile, _ := test.Point.ToTile(test.ExpectedZ)

		assert.Equal(t, test.ExpectedZ, tile.Z)
		assert.Equal(t, test.ExpectedX, tile.X)
		assert.Equal(t, test.ExpectedY, tile.Y)
	}
}

func TestPointInTile(t *testing.T) {
	var tests = []struct {
		Point    Point
		Tile     Tile
		Expected bool
	}{
		{
			Point:    Point{Lat: 41.84, Lng: -87.65},
			Tile:     Tile{Z: 3, X: 2, Y: 2},
			Expected: true,
		},
		{
			Point:    Point{Lat: 52.44950563632098, Lng: 13.357951727129988},
			Tile:     Tile{Z: 15, X: 17599, Y: 10756},
			Expected: true,
		},
		{
			Point:    Point{Lat: 55.676098, Lng: 12.568337},
			Tile:     Tile{Z: 11, X: 1095, Y: 641},
			Expected: true,
		},
		{
			Point:    Point{Lat: 52.25, Lng: 13.37},
			Tile:     Tile{Z: 11, X: 111, Y: 111},
			Expected: false,
		},
	}

	for _, test := range tests {
		isIntersecting, _ := test.Point.Intersects(test.Tile)

		assert.Equal(t, test.Expected, isIntersecting)
	}
}
