package math_test

import (
	"fmt"
	"testing"

	"github.com/peterhp/gow/math"
	"github.com/peterhp/gow/utils"
)

func TestIntAdd(t *testing.T) {
	utils.AssertEquals(t, 10, math.IntAdd(2, 8))
}

func TestIntSub(t *testing.T) {
	utils.AssertEquals(t, -6, math.IntSub(2, 8))
}

func TestIntMul(t *testing.T) {
	utils.AssertEquals(t, 16, math.IntMul(2, 8))
}

func TestIntDiv(t *testing.T) {
	utils.AssertEquals(t, 4, math.IntDiv(8, 2))
}

func BenchmarkIntAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.IntAdd(2, 8)
	}
}

func BenchmarkIntSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.IntSub(2, 8)
	}
}

func BenchmarkIntMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.IntMul(2, 8)
	}
}

func BenchmarkIntDiv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.IntDiv(8, 2)
	}
}

func ExampleIntAdd() {
	fmt.Println(math.IntAdd(2, 8))

	// Output:
	// 10
}

func ExampleIntSub() {
	fmt.Println(math.IntSub(2, 8))

	// Output:
	// -6
}

func ExampleIntMul() {
	fmt.Println(math.IntMul(2, 8))

	// Output:
	// 16
}

func ExampleIntDiv() {
	fmt.Println(math.IntDiv(8, 2))

	// Output:
	// 4
}
