package img

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func Create(creator func(int, int) [][]uint8) func(int, int) image.Image {
	return func(dx, dy int) image.Image {
		img := image.NewRGBA(image.Rect(0, 0, dx, dy))

		data := creator(dx, dy)
		for y := 0; y < dy; y++ {
			for x := 0; x < dx; x++ {
				img.SetRGBA(x, y, color.RGBA{
					R: data[y][x],
					G: data[y][x],
					B: 0xFF,
					A: 0xFF,
				})
			}
		}

		return img
	}
}

func SavePNG(filename string, img image.Image) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return err
	}

	return nil
}
