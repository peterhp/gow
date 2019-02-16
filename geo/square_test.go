package geo_test

import (
	"math"
	"testing"

	"github.com/peterhp/gow/geo"
	"github.com/peterhp/gow/utils"
)

func TestSquare(t *testing.T) {
	square := geo.Square{Side: 1.0}

	t.Run("Square Diagonal", func(t *testing.T) {
		utils.AssertEquals(t, math.Sqrt(2.0), square.Diagonal())
	})

	t.Run("Square Perimeter", func(t *testing.T) {
		utils.AssertEquals(t, 4.0, square.Perimeter())
	})

	t.Run("Square Area", func(t *testing.T) {
		utils.AssertEquals(t, 1.0, square.Area())
	})
}
