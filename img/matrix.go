package img

type Matrix struct {
	data [][]uint8
}

func NewMatrix(dx, dy int) *Matrix {
	m := make([][]uint8, dy)
	for y := range m {
		m[y] = make([]uint8, dx)
	}
	return &Matrix{m}
}

func (m Matrix) Fill(operator func(int, int) uint8) {
	for y := range m.data {
		for x := range m.data[y] {
			m.data[y][x] = operator(x, y)
		}
	}
}

func (m Matrix) Zeros() {
	m.Fill(func(x, y int) uint8 {
		return 0
	})
}

func (m Matrix) Ones() {
	m.Fill(func(x, y int) uint8 {
		return 0xFF
	})
}

func (m Matrix) Data() [][]uint8 {
	return m.data
}
