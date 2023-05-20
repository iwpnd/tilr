package tyler

import (
	"math"
	"reflect"
)

func radToDegree(rad float64) float64 {
	return rad * 180 / math.Pi
}

func degreeToRad(degree float64) float64 {
	return degree * math.Pi / 180
}

func tileToLng(x, z int) float64 {
	return float64(x)/math.Pow(2.0, float64(z))*360.0 - 180
}

func tileToLat(y, z int) float64 {
	n := math.Pi - (2.0*math.Pi*float64(y))/math.Pow(2.0, float64(z))
	return radToDegree(math.Atan(math.Sinh(n)))
}

func min(x, y float64) float64 {
	if x > y {
		return y
	}

	return x
}

func max(x, y float64) float64 {
	if x > y {
		return x
	}

	return y
}

func inDeep(x Tile, y []Tile) bool {
	for i := range y {
		if reflect.DeepEqual(x, y[i]) {
			return true
		}
	}
	return false
}

func ToFeatureCollection(tiles []Tile) ([]byte, error) {
	fc := []byte(`{"type":"FeatureCollection","features":[`)
	for i := range tiles {
		s, err := tiles[i].MarshallGeoJSON()
		if err != nil {
			return nil, err
		}
		if i != len(tiles)-1 {
			s = append(s, []byte(",")...)
		}
		fc = append(fc, s...)

	}
	end := []byte("]}")
	fc = append(fc, end...)

	return fc, nil
}
