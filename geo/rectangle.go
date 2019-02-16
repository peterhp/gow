package geo

import (
	"math"
)

type Rectangle struct {
	Width, Height float64
}

func (rect Rectangle) Diagonal() float64 {
	return math.Sqrt(rect.Width*rect.Width + rect.Height*rect.Height)
}

func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.Width + rect.Height)
}

func (rect Rectangle) Area() float64 {
	return rect.Width * rect.Height
}
