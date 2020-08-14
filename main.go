package main

import (
	"fmt"
)

// HOW TO RUN =======================================================

// func main() {
// 	fmt.Println("Hello, Go")
// 	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
// 	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
// }

// BASIC =============================================================

// var c, python, java bool

// var i1, j int = 1, 2

// var (
// 	ToBe   bool       = false
// 	MaxInt uint64     = 1<<64 - 1
// 	z      complex128 = cmplx.Sqrt(-5 + 12i)
// )

// const Pi = 3.14

// const (
// 	// Create a huge number by shifting a 1 bit left 100 places.
// 	// In other words, the binary number that is 1 followed by 100 zeroes.
// 	Big = 1 << 100
// 	// Shift it right again 99 places, so we end up with 1<<1, or 2.
// 	Small = Big >> 99
// )

// // tour
// func main() {
// 	// Package
// 	fmt.Println("My favorite number is", rand.Intn(10))

// 	// Import
// 	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

// 	// Exported names = function name begin with capital letter
// 	fmt.Println(math.Pi)

// 	// Functions
// 	fmt.Println(add(42, 13))

// 	// Multiple results
// 	a, b := swap("hello", "world")
// 	fmt.Println(a, b)

// 	// Named return values
// 	fmt.Println(split(17))

// 	// Variables
// 	var i int
// 	fmt.Println(i, c, python, java)

// 	// Variables with initializers
// 	var c1, python, java = true, false, "no!"
// 	fmt.Println(i1, j, c1, python, java)

// 	// Short variable declarations
// 	var i2, j2 int = 1, 2
// 	k := 3
// 	c2, python1, java1 := true, false, "no!"
// 	fmt.Println(i2, j2, k, c2, python1, java1)

// 	// Basic types
// 	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
// 	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
// 	fmt.Printf("Type: %T Value: %v\n", z, z)

// 	// Zero values
// 	var i3 int
// 	var f float64
// 	var b3 bool
// 	var s string
// 	fmt.Printf("%v %v %v %q\n", i3, f, b3, s)

// 	// Type conversions
// 	var x, y int = 3, 4
// 	var f1 float64 = math.Sqrt(float64(x*x + y*y))
// 	var z uint = uint(f1)
// 	fmt.Println(x, y, z)

// 	// Type inference
// 	v := 42.2 // change me!
// 	fmt.Printf("v is of type %T\n", v)

// 	// Constants
// 	const World = "世界"
// 	fmt.Println("Hello", World)
// 	fmt.Println("Happy", Pi, "Day")
// 	const Truth = true
// 	fmt.Println("Go rules?", Truth)

// 	fmt.Println(needInt(Small))
// 	fmt.Println(needFloat(Small))
// 	fmt.Println(needFloat(Big))
// }

// func add(x, y int) int {
// 	return x + y
// }

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }

// func needInt(x int) int { return x*10 + 1 }
// func needFloat(x float64) float64 {
// 	return x * 0.1
// }

// FLOW CONTROL ==================================================

func main() {

	// For

	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// For continued

	// sum := 1
	// for sum < 1000 {
	// 	sum += sum
	// }
	// fmt.Println(sum)

	// Forever

	// for {
	// 	fmt.Println("wtf")
	// }

	// If

	// fmt.Println(sqrt(2), sqrt(-4))

	// If with a short statement

	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )

	// Exercise: Loops and Functions

	// fmt.Println(sqrt(2))
	// fmt.Println(math.Sqrt(2))

	// Switch

	// fmt.Print("Go runs on ")
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("OS X.")
	// case "linux":
	// 	fmt.Println("Linux.")
	// default:
	// 	// freebsd, openbsd,
	// 	// plan9, windows...
	// 	fmt.Printf("%s.\n", os)
	// }

	// Switch evaluation order

	// today := time.Now().Weekday()
	// fmt.Println("Today is", today)
	// fmt.Println("time.Saturday", time.Saturday)
	// fmt.Println("today + 0", today+0)
	// fmt.Println("When's Saturday?")
	// switch time.Saturday {
	// case today + 0:
	// 	fmt.Println("Today.")
	// case today + 1:
	// 	fmt.Println("Tomorrow.")
	// case today + 2:
	// 	fmt.Println("In two days.")
	// default:
	// 	fmt.Println("Too far away.")
	// }

	// Switch with no condition

	// t := time.Now()
	// fmt.Println("now is", t)
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("Good morning!")
	// case t.Hour() < 17:
	// 	fmt.Println("Good afternoon.")
	// default:
	// 	fmt.Println("Good evening.")
	// }

	// Defer

	// defer fmt.Println("world")

	// fmt.Println("hello")

	// Stacking defers

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// func sqrt(x float64) float64 {
// 	z := 1.0
// 	for i := 0; i < 10; i++ {
// 		z -= (z*z - x) / (2 * z)
// 	}
// 	return z
// }

// func pow(x, n, lim float64) float64 {
// 	if v := math.Pow(x, n); v < lim {
// 		return v
// 	} else {
// 		fmt.Printf("%g >= %g\n", v, lim)
// 	}
// 	// can't use v here, though
// 	return lim
// }

// func sqrt(x float64) string {
// 	if x < 0 {
// 		return sqrt(-x) + "i"
// 	}
// 	return fmt.Sprint(math.Sqrt(x))
// }
