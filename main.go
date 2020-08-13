package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// how to run go
// func main() {
// 	fmt.Println("Hello, Go")
// 	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
// 	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
// }

var c, python, java bool

var i1, j int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

// tour
func main() {
	// Package
	fmt.Println("My favorite number is", rand.Intn(10))

	// Import
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// Exported names = function name begin with capital letter
	fmt.Println(math.Pi)

	// Functions
	fmt.Println(add(42, 13))

	// Multiple results
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// Named return values
	fmt.Println(split(17))

	// Variables
	var i int
	fmt.Println(i, c, python, java)

	// Variables with initializers
	var c1, python, java = true, false, "no!"
	fmt.Println(i1, j, c1, python, java)

	// Short variable declarations
	var i2, j2 int = 1, 2
	k := 3
	c2, python1, java1 := true, false, "no!"
	fmt.Println(i2, j2, k, c2, python1, java1)

	// Basic types
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Zero values
	var i3 int
	var f float64
	var b3 bool
	var s string
	fmt.Printf("%v %v %v %q\n", i3, f, b3, s)

	// Type conversions
	var x, y int = 3, 4
	var f1 float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f1)
	fmt.Println(x, y, z)

	// Type inference
	v := 42.2 // change me!
	fmt.Printf("v is of type %T\n", v)

	// Constants
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
