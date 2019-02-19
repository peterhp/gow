package main

import (
	"fmt"
	"strings"

	"github.com/peterhp/gow/img"
)

func generateImage(opName string) {
	const imgWidth, imgHeight = 256, 256
	image := img.ImageFactory(img.Operators[opName])(imgWidth, imgHeight)

	img.SavePNG(strings.Join([]string{opName, ".png"}, ""), image)
}

func main() {
	fmt.Println("Welcome to GoLang world!")

	generateImage("square_add")
}
