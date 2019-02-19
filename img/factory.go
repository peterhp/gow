package img

import (
	"image"
)

var Operators = map[string]func(int, int) uint8{
	"zero": func(x, y int) uint8 {
		return 0
	},
	"one": func(x, y int) uint8 {
		return 0xFF
	},

	"add": func(x, y int) uint8 {
		return uint8(x + y)
	},
	"sub": func(x, y int) uint8 {
		return uint8(x - y)
	},
	"mul": func(x, y int) uint8 {
		return uint8(x * y)
	},

	"square_add": func(x, y int) uint8 {
		return uint8((x * x) + (y * y))
	},
	"square_sub": func(x, y int) uint8 {
		return uint8((x * x) - (y * y))
	},
	"square_mul": func(x, y int) uint8 {
		return uint8((x * x) * (y * y))
	},

	"exponent": func(x, y int) uint8 {
		return uint8(x ^ y)
	},
	"mixed_exponent_add": func(x, y int) uint8 {
		return uint8((y ^ x) + (x ^ y))
	},
	"mixed_add_sub_exponent": func(x, y int) uint8 {
		return uint8((x + y) ^ (x - y))
	},

	"beautiful_triangle_pieces": func(x, y int) uint8 {
		return uint8((x - y) ^ y - x)
	},
}

func MatrixFactory(operator func(int, int) uint8) func(int, int) [][]uint8 {
	return func(dx, dy int) [][]uint8 {
		matrix := NewMatrix(dx, dy)
		matrix.Fill(operator)
		return matrix.Data()
	}
}

func ImageFactory(operator func(int, int) uint8) func(int, int) image.Image {
	return func(dx, dy int) image.Image {
		return Create(MatrixFactory(operator))(dx, dy)
	}
}
