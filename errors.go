package tilr

import "fmt"

type ErrInvalidPoint struct {
	p Point
}

func (e *ErrInvalidPoint) Error() string {
	return fmt.Sprintf("Point{Lat: %v, Lng: %v} - invalid point", e.p.Lat, e.p.Lng)
}
