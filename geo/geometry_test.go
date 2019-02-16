package geo_test

import (
	"math"
	"testing"

	"github.com/peterhp/gow/geo"
	"github.com/peterhp/gow/utils"
)

func TestGeometryFactory(t *testing.T) {
	t.Run("New Circle", func(t *testing.T) {
		g := geo.NewGeometry(geo.CIRCLE, 1.0)

		utils.AssertEquals(t, 2.0*math.Pi, g.Perimeter())
		utils.AssertEquals(t, math.Pi, g.Area())
	})

	t.Run("New Rectangle", func(t *testing.T) {
		g := geo.NewGeometry(geo.RECTANGLE, 4.0, 3.0)

		utils.AssertEquals(t, 14.0, g.Perimeter())
		utils.AssertEquals(t, 12.0, g.Area())
	})

	t.Run("New Square", func(t *testing.T) {
		g := geo.NewGeometry(geo.SQUARE, 1.0)

		utils.AssertEquals(t, 4.0, g.Perimeter())
		utils.AssertEquals(t, 1.0, g.Area())
	})

	t.Run("New Rectangle Failed", func(t *testing.T) {
		g := geo.NewGeometry(geo.RECTANGLE, 1.0)
		if g != nil {
			t.Errorf("Geometry factory has issues.")
		}
	})
}
