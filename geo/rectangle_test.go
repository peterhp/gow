package geo_test

import (
	"testing"

	"github.com/peterhp/gow/geo"
	"github.com/peterhp/gow/utils"
)

func TestRectangle(t *testing.T) {
	rect := geo.Rectangle{Width: 3.0, Height: 4.0}

	t.Run("Rectangle Diagonal", func(t *testing.T) {
		utils.AssertEquals(t, 5.0, rect.Diagonal())
	})

	t.Run("Rectangle Perimeter", func(t *testing.T) {
		utils.AssertEquals(t, 14.0, rect.Perimeter())
	})

	t.Run("Rectangle Area", func(t *testing.T) {
		utils.AssertEquals(t, 12.0, rect.Area())
	})
}
