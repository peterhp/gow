package geo

import (
	"math"
)

type Circle struct {
	Radius float64
}

func (circle Circle) Diameter() float64 {
	return 2 * circle.Radius
}

func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.Radius
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}
