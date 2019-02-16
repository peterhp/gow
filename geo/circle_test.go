package geo_test

import (
	"math"
	"testing"

	"github.com/peterhp/gow/geo"
	"github.com/peterhp/gow/utils"
)

func TestCircle(t *testing.T) {
	circle := geo.Circle{Radius: 1.0}

	t.Run("Circle Diameter", func(t *testing.T) {
		utils.AssertEquals(t, 2.0, circle.Diameter())
	})

	t.Run("Circle Perimeter", func(t *testing.T) {
		utils.AssertEquals(t, 2.0*math.Pi, circle.Perimeter())
	})

	t.Run("Circle Area", func(t *testing.T) {
		utils.AssertEquals(t, math.Pi, circle.Area())
	})
}
