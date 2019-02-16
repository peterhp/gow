package main

import (
	"fmt"

	"github.com/peterhp/gow/geo"
)

func main() {
	fmt.Println("Welcome to GoLang world!")

	fmt.Println(geo.NewGeometry(geo.RECTANGLE, 4.0, 3.0).Area())
}
