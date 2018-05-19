package main

import "fmt"

import "github.com/peterhp/gow/util"

func main() {
	fmt.Printf("Hello, GO!\n")

	fmt.Printf("%d + %d = %d\n", 1, 2, math.IntAdd(1, 2))
}
