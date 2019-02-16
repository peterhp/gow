package geo

import (
	"math"
)

type Square struct {
	Side float64
}

func (square Square) Diagonal() float64 {
	return math.Sqrt(2) * square.Side
}

func (square Square) Perimeter() float64 {
	return 4 * square.Side
}

func (square Square) Area() float64 {
	return square.Side * square.Side
}
