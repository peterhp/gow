package geo

type Geometry interface {
	Perimeter() float64
	Area() float64
}

type GeoType int

const (
	CIRCLE    GeoType = 0x01
	RECTANGLE GeoType = 0x31
	SQUARE    GeoType = 0x32
)

func NewGeometry(t GeoType, params ...float64) Geometry {
	switch {
	case t == CIRCLE && len(params) == 1:
		return &Circle{Radius: params[0]}

	case t == RECTANGLE && len(params) == 2:
		return &Rectangle{Width: params[0], Height: params[1]}

	case t == SQUARE && len(params) == 1:
		return &Square{Side: params[0]}

	default:
		return nil
	}
}
